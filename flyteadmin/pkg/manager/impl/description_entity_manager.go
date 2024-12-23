package impl

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"

	"github.com/flyteorg/flyte/flyteadmin/pkg/common"
	"github.com/flyteorg/flyte/flyteadmin/pkg/errors"
	"github.com/flyteorg/flyte/flyteadmin/pkg/manager/impl/util"
	"github.com/flyteorg/flyte/flyteadmin/pkg/manager/impl/validation"
	"github.com/flyteorg/flyte/flyteadmin/pkg/manager/interfaces"
	repoInterfaces "github.com/flyteorg/flyte/flyteadmin/pkg/repositories/interfaces"
	"github.com/flyteorg/flyte/flyteadmin/pkg/repositories/models"
	"github.com/flyteorg/flyte/flyteadmin/pkg/repositories/transformers"
	runtimeInterfaces "github.com/flyteorg/flyte/flyteadmin/pkg/runtime/interfaces"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flytestdlib/contextutils"
	"github.com/flyteorg/flyte/flytestdlib/logger"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
)

type DescriptionEntityMetrics struct {
	Scope promutils.Scope
}

type DescriptionEntityManager struct {
	db      repoInterfaces.Repository
	config  runtimeInterfaces.Configuration
	metrics DescriptionEntityMetrics
}

func (d *DescriptionEntityManager) GetDescriptionEntity(ctx context.Context, request *admin.ObjectGetRequest) (
	*admin.DescriptionEntity, error) {
	if err := validation.ValidateDescriptionEntityGetRequest(request); err != nil {
		logger.Errorf(ctx, "invalid request [%+v]: %v", request, err)
		return nil, err
	}
	ctx = contextutils.WithProjectDomain(ctx, request.GetId().GetProject(), request.GetId().GetDomain())
	return util.GetDescriptionEntity(ctx, d.db, request.GetId())
}

func (d *DescriptionEntityManager) ListDescriptionEntity(ctx context.Context, request *admin.DescriptionEntityListRequest) (*admin.DescriptionEntityList, error) {
	// Check required fields
	if err := validation.ValidateDescriptionEntityListRequest(request); err != nil {
		return nil, err
	}
	ctx = contextutils.WithProjectDomain(ctx, request.GetId().GetProject(), request.GetId().GetDomain())

	if request.GetResourceType() == core.ResourceType_WORKFLOW {
		ctx = contextutils.WithWorkflowID(ctx, request.GetId().GetName())
	} else {
		ctx = contextutils.WithTaskID(ctx, request.GetId().GetName())
	}

	filters, err := util.GetDbFilters(util.FilterSpec{
		Project:        request.GetId().GetProject(),
		Domain:         request.GetId().GetDomain(),
		Name:           request.GetId().GetName(),
		RequestFilters: request.GetFilters(),
	}, common.ResourceTypeToEntity[request.GetResourceType()])
	if err != nil {
		logger.Error(ctx, "failed to get database filter")
		return nil, err
	}

	sortParameter, err := common.NewSortParameter(request.GetSortBy(), models.DescriptionEntityColumns)
	if err != nil {
		return nil, err
	}

	offset, err := validation.ValidateToken(request.GetToken())
	if err != nil {
		return nil, errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"invalid pagination token %s for ListWorkflows", request.GetToken())
	}
	listDescriptionEntitiesInput := repoInterfaces.ListResourceInput{
		Limit:         int(request.GetLimit()),
		Offset:        offset,
		InlineFilters: filters,
		SortParameter: sortParameter,
	}
	output, err := d.db.DescriptionEntityRepo().List(ctx, listDescriptionEntitiesInput)
	if err != nil {
		logger.Debugf(ctx, "Failed to list workflows with [%+v] with err %v", request.GetId(), err)
		return nil, err
	}
	descriptionEntityList, err := transformers.FromDescriptionEntityModels(output.Entities)
	if err != nil {
		logger.Errorf(ctx,
			"Failed to transform workflow models [%+v] with err: %v", output.Entities, err)
		return nil, err
	}
	var token string
	if len(output.Entities) == int(request.GetLimit()) {
		token = strconv.Itoa(offset + len(output.Entities))
	}
	return &admin.DescriptionEntityList{
		DescriptionEntities: descriptionEntityList,
		Token:               token,
	}, nil
}

func NewDescriptionEntityManager(
	db repoInterfaces.Repository,
	config runtimeInterfaces.Configuration,
	scope promutils.Scope) interfaces.DescriptionEntityInterface {

	metrics := DescriptionEntityMetrics{
		Scope: scope,
	}
	return &DescriptionEntityManager{
		db:      db,
		config:  config,
		metrics: metrics,
	}
}

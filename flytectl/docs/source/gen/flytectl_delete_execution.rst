.. _flytectl_delete_execution:

flytectl delete execution
-------------------------

Terminates/deletes execution resources.

Synopsis
~~~~~~~~



Task executions can be aborted only if they are in non-terminal state. If they are FAILED, ABORTED, or SUCCEEDED, calling terminate on them has no effect.
Terminate a single execution with its name:

::

 flytectl delete execution c6a51x2l9e  -d development  -p flytesnacks

.. note::
    The terms execution/executions are interchangeable in these commands.

Get an execution to check its state:

::

 flytectl get execution  -d development  -p flytesnacks
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | NAME (7)   | WORKFLOW NAME                                                           | TYPE     | PHASE     | STARTED                        | ELAPSED TIME  |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | c6a51x2l9e | recipes.core.basic.lp.go_greet                                          | WORKFLOW | ABORTED   | 2021-02-17T08:13:04.680476300Z | 15.540361300s |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 

Terminate multiple executions with their names:
::

 flytectl delete execution eeam9s8sny p4wv4hwgc4  -d development  -p flytesnacks

Get an execution to find the state of previously terminated executions:

::

 flytectl get execution  -d development  -p flytesnacks
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | NAME (7)   | WORKFLOW NAME                                                           | TYPE     | PHASE     | STARTED                        | ELAPSED TIME  |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | c6a51x2l9e | recipes.core.basic.lp.go_greet                                          | WORKFLOW | ABORTED   | 2021-02-17T08:13:04.680476300Z | 15.540361300s |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | eeam9s8sny | recipes.core.basic.lp.go_greet                                          | WORKFLOW | ABORTED   | 2021-02-17T08:14:04.803084100Z | 42.306385500s |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 
 | p4wv4hwgc4 | recipes.core.basic.lp.go_greet                                          | WORKFLOW | ABORTED   | 2021-02-17T08:14:27.476307400Z | 19.727504400s |
  ------------ ------------------------------------------------------------------------- ---------- ----------- -------------------------------- --------------- 

Usage


::

  flytectl delete execution [flags]

Options
~~~~~~~

::

      --dryRun   execute command without making any modifications.
  -h, --help     help for execution

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --admin.audience string                        Audience to use when initiating OAuth2 authorization requests.
      --admin.authType string                        Type of OAuth2 flow used for communicating with admin.ClientSecret, Pkce, ExternalCommand are valid values (default "ClientSecret")
      --admin.authorizationHeader string             Custom metadata header to pass JWT
      --admin.authorizationServerUrl string          This is the URL to your IdP's authorization server. It'll default to Endpoint
      --admin.caCertFilePath string                  Use specified certificate file to verify the admin server peer.
      --admin.clientId string                        Client ID (default "flytepropeller")
      --admin.clientSecretEnvVar string              Environment variable containing the client secret
      --admin.clientSecretLocation string            File containing the client secret (default "/etc/secrets/client_secret")
      --admin.command strings                        Command for external authentication token generation
      --admin.defaultServiceConfig string            
      --admin.deviceFlowConfig.pollInterval string   amount of time the device flow would poll the token endpoint if auth server doesn't return a polling interval. Okta and google IDP do return an interval' (default "5s")
      --admin.deviceFlowConfig.refreshTime string    grace period from the token expiry after which it would refresh the token. (default "5m0s")
      --admin.deviceFlowConfig.timeout string        amount of time the device flow should complete or else it will be cancelled. (default "10m0s")
      --admin.endpoint string                        For admin types,  specify where the uri of the service is located.
      --admin.httpProxyURL string                    OPTIONAL: HTTP Proxy to be used for OAuth requests.
      --admin.insecure                               Use insecure connection.
      --admin.insecureSkipVerify                     InsecureSkipVerify controls whether a client verifies the server's certificate chain and host name. Caution : shouldn't be use for production usecases'
      --admin.maxBackoffDelay string                 Max delay for grpc backoff (default "8s")
      --admin.maxMessageSizeBytes int                The max size in bytes for incoming gRPC messages
      --admin.maxRetries int                         Max number of gRPC retries (default 4)
      --admin.perRetryTimeout string                 gRPC per retry timeout (default "15s")
      --admin.pkceConfig.refreshTime string          grace period from the token expiry after which it would refresh the token. (default "5m0s")
      --admin.pkceConfig.timeout string              Amount of time the browser session would be active for authentication from client app. (default "2m0s")
      --admin.proxyCommand strings                   Command for external proxy-authorization token generation
      --admin.scopes strings                         List of scopes to request
      --admin.tokenRefreshWindow string              Max duration between token refresh attempt and token expiry. (default "0s")
      --admin.tokenUrl string                        OPTIONAL: Your IdP's token endpoint. It'll be discovered from flyte admin's OAuth Metadata endpoint if not provided.
      --admin.useAudienceFromAdmin                   Use Audience configured from admins public endpoint config.
      --admin.useAuth                                Deprecated: Auth will be enabled/disabled based on admin's dynamically discovered information.
  -c, --config string                                config file (default is $HOME/.flyte/config.yaml)
      --console.endpoint string                      Endpoint of console,  if different than flyte admin
  -d, --domain string                                Specifies the Flyte project's domain.
      --files.archive                                Pass in archive file either an http link or local path.
      --files.assumableIamRole string                Custom assumable iam auth role to register launch plans with.
      --files.continueOnError                        Continue on error when registering files.
      --files.destinationDirectory string            Location of source code in container.
      --files.dryRun                                 Execute command without making any modifications.
      --files.enableSchedule                         Enable the schedule if the files contain schedulable launchplan.
      --files.force                                  Force use of version number on entities registered with flyte.
      --files.k8ServiceAccount string                Deprecated. Please use --K8sServiceAccount
      --files.k8sServiceAccount string               Custom kubernetes service account auth role to register launch plans with.
      --files.outputLocationPrefix string            Custom output location prefix for offloaded types (files/schemas).
      --files.sourceUploadPath string                Deprecated: Update flyte admin to avoid having to configure storage access from flytectl.
      --files.version string                         Version of the entity to be registered with flyte which are un-versioned after serialization.
  -i, --interactive                                  Set this flag to use an interactive CLI
      --logger.formatter.type string                 Sets logging format type. (default "json")
      --logger.level int                             Sets the minimum logging level. (default 3)
      --logger.mute                                  Mutes all logs regardless of severity. Intended for benchmarks/tests only.
      --logger.show-source                           Includes source code location in logs.
      --otel.file.filename string                    Filename to store exported telemetry traces (default "/tmp/trace.txt")
      --otel.jaeger.endpoint string                  Endpoint for the jaeger telemetry trace ingestor (default "http://localhost:14268/api/traces")
      --otel.otlpgrpc.endpoint string                Endpoint for the OTLP telemetry trace collector (default "http://localhost:4317")
      --otel.otlphttp.endpoint string                Endpoint for the OTLP telemetry trace collector (default "http://localhost:4318/v1/traces")
      --otel.sampler.parentSampler string            Sets the parent sampler to use for the tracer (default "always")
      --otel.type string                             Sets the type of exporter to configure [noop/file/jaeger/otlpgrpc/otlphttp]. (default "noop")
  -o, --output string                                Specifies the output type - supported formats [TABLE JSON YAML DOT DOTURL]. NOTE: dot, doturl are only supported for Workflow (default "TABLE")
  -p, --project string                               Specifies the Flyte project.
      --storage.cache.max_size_mbs int               Maximum size of the cache where the Blob store data is cached in-memory. If not specified or set to 0,  cache is not used
      --storage.cache.target_gc_percent int          Sets the garbage collection target percentage.
      --storage.connection.access-key string         Access key to use. Only required when authtype is set to accesskey.
      --storage.connection.auth-type string          Auth Type to use [iam, accesskey]. (default "iam")
      --storage.connection.disable-ssl               Disables SSL connection. Should only be used for development.
      --storage.connection.endpoint string           URL for storage client to connect to.
      --storage.connection.region string             Region to connect to. (default "us-east-1")
      --storage.connection.secret-key string         Secret to use when accesskey is set.
      --storage.container string                     Initial container (in s3 a bucket) to create -if it doesn't exist-.'
      --storage.defaultHttpClient.timeout string     Sets time out on the http client. (default "0s")
      --storage.enable-multicontainer                If this is true,  then the container argument is overlooked and redundant. This config will automatically open new connections to new containers/buckets as they are encountered
      --storage.limits.maxDownloadMBs int            Maximum allowed download size (in MBs) per call. (default 2)
      --storage.stow.config stringToString           Configuration for stow backend. Refer to github/flyteorg/stow (default [])
      --storage.stow.kind string                     Kind of Stow backend to use. Refer to github/flyteorg/stow
      --storage.type string                          Sets the type of storage to configure [s3/minio/local/mem/stow]. (default "s3")

SEE ALSO
~~~~~~~~

* :doc:`flytectl_delete` 	 - Terminates/deletes various Flyte resources such as executions and resource attributes.


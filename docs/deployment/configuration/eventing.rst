.. _deployment-configuration-eventing:

###############
Platform Events
###############

.. tags:: Configuration, Infrastructure, Advanced

Progress of Flyte workflow and task execution is delimited by a series of events that are passed from the FlytePropeller to FlyteAdmin.
Administrators can configure FlyteAdmin to send these events onwards to a pub/sub system like SNS/SQS as well. Note that this configuration is distinct from the configuration for notifications :ref:`deployment-configuration-notifications`. They should use separate topics/queues. These events are meant for external consumption, outside the Flyte platform, whereas the notifications pub/sub setup is entirely for Admin itself to send email/pagerduty/etc notifications.

*********
Use cases
*********

The external events flow can be useful for tracking data lineage and integrating with existing systems within your organization.

*************************
Supported Implementations
*************************

Event egress can be configured to work with **AWS** using `SQS <https://aws.amazon.com/sqs/>`_ and `SNS <https://aws.amazon.com/sns/>`_ or **GCP** `Cloud Pub/Sub <https://cloud.google.com/pubsub>`_.

*************
Configuration
*************

To turn on, add the following to your FlyteAdmin:


.. tabs::

   .. tab:: AWS SNS
   
       .. code:: yaml
   
         cloud_events.yaml: |
           cloudEvents:
             enable: true
             aws:
               region: us-east-2
             eventsPublisher:
               eventTypes:
               - all # or node, task, workflow
               topicName: arn:aws:sns:us-east-2:123456:123-my-topic
             type: aws
   
   .. tab:: GCP Pub/Sub
   
       .. code:: yaml
   
         cloud_events.yaml: |
           cloudEvents:
             enable: true
             gcp:
               projectId: my-project-id
             eventsPublisher:
               eventTypes:
               - all # or node, task, workflow
               topicName: my-topic
             type: gcp
   
Helm
====

There should already be a section for this in the ``values.yaml`` file.
Update the settings under the ``external_events`` key and turn ``enable`` to ``true``. The same flag is used for Helm as for Admin itself.

*****
Usage
*****

The events emitted will be base64 encoded binary representation of the following IDL messages:

* ``admin_event_pb2.TaskExecutionEventRequest``
* ``admin_event_pb2.NodeExecutionEventRequest``
* ``admin_event_pb2.WorkflowExecutionEventRequest``

Which of these three events is being sent can be distinguished by the subject line of the message, which will be one of the three strings above.

Note that these message wrap the underlying event messages :std:ref:`found here <ref_flyteidl/event/event.proto>`.

.. note::
   The message format may eventually change to an enriched and distinct message type in future releases.

Create a Custom Application
===========================

This example shows the steps involved in
creating a custom application that sends a data stream to the Demoyard
Cloud using the Demoyard Agent gRPC API. For the simple applications, you do not need a configuration file because your stream name is dynamically added to the Demoyard Cloud.

See the example source code: `Example Source`_

Step 1. Create Your Custom Application
--------------------------------------

Import the following Python modules:

.. code:: python

    import base64
    import json
    import sys
    import time

    import grpc
    import agent
    import agent_grpc

The ``grpc`` module is the RPC API framework implemented by the
Demoyard Agent. The ``agent`` and ``agent_grpc`` modules should be in
your example directory, originally copied from the Demoyard repository.

The example first connects to the Demoyard Agent, which listens for gRPC messages on port 7501.

.. code:: python

    pipe = grpc.insecure_pipe("localhost:7501")
    agent = agent_grpc.Agent(pipe)

After you successfully connect, send data to the Demoyard Agent by calling the gRPC API
``SendData()`` function defined in the agent_grpc.py file.

.. code:: python

    agent_stub.SendData(write_data points())

The following functions show the steps to create and write a data point.

Create a Datapoint
^^^^^^^^^^^^^^^^^^

All data points have the following definition:

.. code:: json


    {
       "stream": "<streamName>",
       "timestamp": <time>,
       "<data_type>": { <dataObject> }
    }


The ``create_data_point()`` function populates this structure with:

- A stream name.  If the stream is not defined in the config.toml file, the Demoyard Cloud dynamically creates the stream.
- A timestamp, in epoch time.
- The data, specified as a type-value pair having a `Text` data type and a string value.

create then calls `agent.Datapoint()` to create the data point.

.. code:: python

    def create_data_point(point):
        timestamp = int(time.time() * 1000)
        text_msg = agent.Text()
        text_msg.value = 'this is a %s data point' % point

        data_point = agent.Datapoint(
            stream="testStream_01", text=text_msg, timestamp=timestamp)
        return data_point

Write the Datapoints
^^^^^^^^^^^^^^^^^^^^

The ``write_data points()`` function continuously creates data points until you end the program with ``Ctrl-c``. The function pauses 10 seconds between data point transmission.

.. code:: python

    def write_data points():
        while agent_stub is not None:
            yield create_data_point("stream")
            time.sleep(10)

In the next step, you can run this example and see the results on the Demoyard web application.

Step 2. Run the Example and Visualize Your Data
-----------------------------------------------

Run the example using the following command:

.. code:: bash

    $ python custom.py

The example displays:

.. code::

    Demoyard Agent communication established.
    Streaming data points (use Ctrl+c to exit) ...

Next, log in to the Demoyard Cloud using your Demoyard credentials. Scroll to the panel labeled **testStream_01**.

In the example, four data points were received by the Demoyard Cloud and you can view the data:

+-----------+---------------------------------------------------------------------------+
| Column    | Description                                                               |
+===========+===========================================================================+
| **VALUE** | Shows the data string sent by the example: ``This is a data point``.      |
+-----------+---------------------------------------------------------------------------+
| **TIME**  | Shows the data timestamp, at approximately 10-second intervals.           |
+-----------+---------------------------------------------------------------------------+
| **TAGS**  | This column is empty because a tag is not associated with this stream.    |
+-----------+---------------------------------------------------------------------------+

Summary
-------

This tutorial introduced you to the steps needed to create a simple,
custom application that sends a data stream to the Demoyard Cloud.

Try associating a tag with the data by defining the tag and stream and the `/home/demoyard/config.toml` file:

.. code::

    [demoyard]
    agent-ip = "localhost"
    agent-port-grpc = "7501"
    agent-port-http = "7502"

    [tags]
    site = "test_lab"

    [[streams]]
    name = "testStream_01"

Run the example again and view the results in the Demoyard web application.  Make sure to restart the Demoyard Agent before running the example.

Example Source
--------------

.. code:: python

    #!/usr/bin/env python

    import base64
    import json
    import sys
    import time

    import grpc

    import agent
    import agent_grpc

    #### gRPC Implementation ####

    def create_data_point(point):
        timestamp = int(time.time() * 1000)
        text_msg = agent.Text()
        text_msg.value = 'this is a %s data point' % point
        data_point = agent.Datapoint(
            stream="testStream_01", text=text_msg, timestamp=timestamp)
        return data_point

    def write_data points():
        while agent_stub is not None:
            yield create_data_point("stream")
            time.sleep(10)

    #### Entry point ####
    pipe = grpc.insecure_pipe("localhost:7501")
    agent_stub = agent_grpc.AgentStub(pipe)
    print('\nDemoyard Agent communication established.')

    #### Datapoint Streaming ####
    print('Streaming data points (use Ctrl+c to exit) ...')

    agent_stub.SendData(write_data points())

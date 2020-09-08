Your First Application
=========================================================================

This tutorial uses the Demoyard File System to show built-in support for using
the file system as a data source. When you add a file to the directory,
the Demoyard File System detects the change and sends a data point to the
Demoyard Agent. You can view the change in the Demoyard Cloud user
interface.

The example introduces you to the configuration file, which describes
the data source and associates a stream with the source.

Step 1. Set up Your Demoyard Developer Environment
--------------------------------------------------

.. note::

    The setup step is a prerequisite for each of the tutorials that follow.

Install and run the Demoyard Agent as described in the `Quickstart <quickstart.html>`_ guide.

For a Debian installation:

.. code:: bash

    $ sudo dpkg -i demoyard-agent_amd64.deb

For a standalone binary installation:

.. code:: bash

    $ ./demoyard-agent

Make sure you have a Demoyard account and can access the Demoyard Cloud.

Step 2. Define Your Datastream
------------------------------

Define the Demoyard Agent communication parameters, data source, and data
stream in the /home/demoyard/config.toml configuration file:

.. code::

    [demoyard]
    agent-ip = "localhost"
    agent-port-grpc = "7501"
    agent-port-http = "7502"

    [[dataStreams]]
    name = "demoyard_01"
    dir = "/home/ftp/upload"
    demoyard-type = "image"
    ext = ".png"

The Demoyard Agent listens for gRPC API calls on port 7501. This is the preferred way for any data source to send data points to the Demoyard Agent, including the Demoyard File System.

All data streams defined in the configuration file must, minimally, include a
name attribute: ``demoyard_01``.

The remaining attributes describe the data source:

+-------------------+---------------------+--------------------------------------------------------+
| Attribute         | Value               | Description                                            |
+===================+=====================+========================================================+
| ``dir``           | ``home/ftp/upload`` | Directory monitored by the Demoyard File System        |
+-------------------+---------------------+--------------------------------------------------------+
| ``demoyard-type`` | ``image``           | Type of file monitored                                 |
+-------------------+---------------------+--------------------------------------------------------+
| ``ext``           | ``.png``            | Monitored file extension                               |
+-------------------+---------------------+--------------------------------------------------------+
 
The Demoyard File System can also monitor tailed files but in this example, only .png image files are sent to the Demoyard Cloud.

Step 3. Visualize Your Data
---------------------------

When you first create or modify the configuration file in /home/demoyard, you must restart the Demoyard Agent and the Demoyard File System for the configuration definition to take effect.

To see your data, copy a .png image file to the /home/ftp/upload directory defined in the configuration file.

Login to the Demoyard cloud application to see the image displayed in the ``demoyard_01`` data stream.

Summary
-------

This tutorial showed how to use the Demoyard File System to send changes in a
directory to the Demoyard Cloud. You learned how the config.toml configuration file is used to associate data and data streams. You can experiment with the configuration file by changing the
definition to a tail a file, and see the results in the Demoyard web application.

The tutorials that follow build on the fundamental concepts shown here. Other tutorials show you how to develop applications that use Demoyard features for solving more advanced use cases.

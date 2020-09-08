Quickstart
=========================================================================

Get started in three easy steps:

1. Download the agent kernel
2. Install the agent
3. Verify successful installation

Requirements
------------

+------------------+----------------------------------------------------+
| Item             | Requirement                                        |
+==================+====================================================+
| Package          | Debian (systemd-compatible) and binary             |
+------------------+----------------------------------------------------+
| Operating system | Linux                                              |
+------------------+----------------------------------------------------+
| Architecture     | AMD 64-bit                                         |
+------------------+----------------------------------------------------+

Prerequisites
-------------

- The kernel installation procedure requires you to authenticate with the cloud using your account credentials.

- The agent uses ports 7501 and 7502 to access the application API. Make sure these ports are available.

Step 1. Download the Agent Kernel
----------------------------------

1. Download the latest kernel release package from the release web site.

2. Extract the `kernel_<version>.tar.gz` file.

    .. code:: bash

        $ tar xvzf kernel_<version>.tar.gz

3. Change directories to the extracted files:

    .. code:: bash

        $ cd kernel_<version>

    Continue with the `Debian Installation`_ or `Standalone Binary Installation`_ depending on your environment.

Step 2. Install the Agent
-------------------------------

This step prompts for your account credentials. For installation on multiple machines, see the `Notes on Batch Installation`_.

After successful installation, the agent listens for gRPC requests on port 7501 and HTTP requests on port 7502.

Debian Installation
^^^^^^^^^^^^^^^^^^^

.. code:: bash

    $ sudo dpkg -i demoyard-agent_amd64.deb

The installer creates an ``Agent`` user and /home/agent directory where it installs the default config.toml configuration file.

The agent is installed in the /etc/agent directory so you can use systemd to manage the agent lifecycle.

Standalone Binary Installation
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code:: bash

    $ ./demoyard-agent

The installer creates the `$HOME/.agent` directory where it installs the default `config.toml` configuration file.

The agent binary is a standalone binary. You can daemonize it using a daemon service of your choice or manually manage its lifecycle.

Step 3. Verify Successful Installation
--------------------------------------

The download includes a Python example that generates data streams and sends them to the cloud using the agent.

Download and run the example to verify successful agent installation:

1. Verify that the agent service is running.

2. Clone the GitHub agent repository and locate the Python example in the `/public/examples/python` directory.

3. Install the dependencies:

    .. code:: bash

        $ pip install -r requirements.txt

4. Login to the web application using your cloud credentials and click the **Live** button, if needed, to be ready to view the telemetry and notifications sent by the Python example.

5. Run the example:

    .. code:: bash

        $ python main.py $PWD

6. The example displays progress in the command window, indicating what is being sent to the agent.  Verify that this data is also displayed in the web application.  The last output sequence is a periodic data stream at 10-second intervals. Enter ``Ctrl+c`` to terminate the continuous data stream and the example.

Notes on Batch Installation
---------------------------

For batch installation, set the following environment variables to match your cloud credentials. This automates cloud authentication.

- `AGENT_EMAIL`
- `AGENT_PASSWORD`

Make sure to unset these variables after the installation successfully completes.

Next Steps
----------

After installing an agent, read the Developer Guide and the Tutorials for a step-by-step guide to communicating with the agent from your application.

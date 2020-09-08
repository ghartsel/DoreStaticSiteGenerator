Make an API Call to the Demoyard Cloud
=========================================================================

The Demoyard Cloud implements management and data storage functions for a network of distributed, autonomous agents. Agents access cloud functionality using the Demoyard Cloud API.

This tutorial authenticates the agent with the cloud and gets the list of users associated with the agent. See the `Example`_ code for the complete implementation.

Authenticate the Agent with the Cloud
-------------------------------------

Agents must first authenticate with the cloud before transferring data. Demoyard Cloud uses JSON Web Token (JWT) authentication.

Pass your Demoyard account email address and password in the JWT request:

.. code:: bash

    curl -X POST -k -H 'Content-Type: application/json' -i 'https://api-dev.demoyard/v1/admin/auth/login' --data '{"email": "joe.demoyard@demoyard.com", "password": "changeMe123"}'

A JWT token is returned:

.. code:: json

    {"authentication":
        {"accessToken":"eyJraWQiOiJpcDJMUH ... v1NLf5Y70xtdtUyxRudUg",
        "refreshToken":"eyJjdHkiOiJKV1Q ... 3pvYhCXvtntwMQFyPUNzA",
        "userId":"5bc046f1-967a-4455-bae1-dc208f5e619a"}
    }

Pass this token in the header of subsequent API calls. The token is valid for one hour.

Get a List of Users Associated with the Agent
---------------------------------------------

Use the ``/admin/users`` URI to get information about the users associated with the agent:

.. code:: bash

    curl -X GET -k -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJraWQiOiJpcDJMUH ... v1NLf5Y70xtdtUyxRudUg' -i 'https://api-dev.demoyard/v1/admin/users'

Notice that the ``Authorization`` field has a value of ``Bearer`` and includes the previously returned JWT token.

Demoyard Cloud returns a JSON structure of user information for all the users associated with the calling agent:

.. code:: json

    {
      "items": [
        {
          "id": "174bb39d-bd58-450b-8614-f325c0525a7f",
          "created": "2018-09-12T22:02:15.748Z",
          "updated": "2018-11-16T23:22:54.937Z",
          "email": "admin@example.com",
          "name": "admin",
          "role": "administrator",
          "enabled": true
        },
        {
          "id": "5bc046f1-967a-4455-bae1-dc208f5e619a",
          "created": "2018-10-25T20:41:17.167Z",
          "updated": "2018-10-25T20:41:17.167Z",
          "email": "qa@demoyard.com",
          "name": "QA",
          "role": "user",
          "enabled": true
        },
                            .
                            . (elided)
                            .
        {
          "id": "c26637d8-3d7b-47e3-b2bb-19ce7d250289",
          "created": "2018-11-16T21:08:47.793Z",
          "updated": "2018-11-16T21:08:47.793Z",
          "email": "mmaynard@demoyard",
          "name": "Mark Maynard",
          "role": "user",
          "enabled": true
        }
      ]
    }

Example
-------

.. code:: python

    #!/usr/bin/env python

    import requests
    import json
    import getpass

    baseURI = 'https://api-dev.demoyard/v1/admin/'

    print '\nEnter your Demoyard login credentials ...'
    user_email = raw_input('Email: ')

    try:
        user_password = getpass.getpass()
    except Exception as error:
        print('ERROR', error)

    response = requests.post(baseURI + 'auth/login', json={"email": user_email, "password": user_password}, headers={'content-type':'application/json'})
    if (response.status_code == 200):

        jwt = json.loads (response.text)
        token = jwt['authentication']['accessToken']

        token_header = 'Bearer ' + token
        response = requests.get(baseURI + 'users', headers={'content-type':'application/json', 'Authorization': token_header})

        users_list = json.loads (response.text)
        for user in users_list['items']:
            print user['id']
            if user['id'] == jwt['authentication']['userId']:
                print "Your user information:"
                response = requests.get(baseURI + 'users/' + user['id'], headers={'content-type':'application/json', 'Authorization': token_header})
                parsed = json.loads (response.text)
                print(json.dumps(parsed, indent=4, sort_keys=True))

    elif (response.status_code == 401):
        print "Invalid login credentials provided"
        print parser.print_help()


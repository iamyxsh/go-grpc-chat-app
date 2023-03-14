# auth-api

This part of the application is responsible for authentication. It is written in Go.


It provides following calls for

    - `Signup`: allow users to register to the platform.
        The JSON object structure is following:
            ```json
            {
                "email": "email@abc.com",
                "password": "password",
            }
            ```

    - `Login`: allow users to login to the platform.
        The JSON object structure is following:
            ```json
            {
                "email": "email@abc.com",
                "password": "password",
            }
            ```
        The response will be something like this 
            ```json
            {
                "msg": {{token}}
            }
            ```
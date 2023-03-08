# line webhook

## build
~~~
$ docker compose up -d
$ ngrok http 8080
~~~

## API
* webhook
    ~~~
    uri : https://{nghrok url}/callback
    method : post
    ~~~
* push message
    ~~~
    uri : https://{nghrok url}/pushmessage/{userID}
    method : post
    post body : { "message":"test 20230308" }
    response body : { "message": "ok" }
    ~~~
* index messages
    ~~~
    uri : https://{nghrok url}/pushmessage/{userID}
    method : get
    response body : 
    {
        "Data": [
            {
                "user_id": "640837b256c6c297cf5552b6",
                "source": "line",
                "type": "recive",
                "message": "test",
                "Time": "2023-03-08T07:27:30.232Z"
            },
            ...
        ]
    }
    ~~~

## demo

![](demo/20230308153007.mp4)
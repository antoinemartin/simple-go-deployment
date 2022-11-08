Sample Go program
=================


To build the image issue the command: 

```console
$ docker build -t localhost:5000/gocmd:latest .
...
$ docker push localhost:5000/gocmd:latest
The push refers to repository [localhost:5000/gocmd]
1ba45b76e8ed: Pushed 
latest: digest: sha256:2b61c2d70f16c26cb85917e2ab1c71d3383e669f057f434710fd4a817f2e61cd size: 528
```

Then the deployment can be done with:

```console
$ kubectl apply -f deployment.yaml
```

And the service can be accessed with:

```console
$ curl  http://172.21.0.2:4000
```

[![Go Report Card](https://goreportcard.com/badge/pmorelli92/go-state-machine-two)](https://goreportcard.com/report/pmorelli92/go-state-machine-two)
[![Build Status](https://travis-ci.com/pmorelli92/go-state-machine-two.svg?branch=master)](https://travis-ci.com/pmorelli92/go-state-machine-two)
[![Coverage Status](https://coveralls.io/repos/github/pmorelli92/go-state-machine-two/badge.svg?branch=master)](https://coveralls.io/github/pmorelli92/go-state-machine-two?branch=master)

# GO State Machine Two
#### A Tech Task on GO

- Environment variables needed to run or debug the API can be found [here](https://github.com/pmorelli92/go-state-machine-two/blob/b769cafe1ffc3d98e21b41b5fccc41b648f96410/_kubernetes/app.yaml)

- The app is runnable on kubernetes just by doing  (needs Docker images to be built first)

  ```
  kubectl apply -f _kubernetes/
  ```

- Postman collection can be imported using this [url](https://www.getpostman.com/collections/d7e3bee8076474163ccc)

  - For debugging the APP the following env vars are needed:

    ```
    GoStateMachineApi http://localhost:8080
    ```

  - For running on kubernetes:

    ```
    GoStateMachineApi http://<REPLACE_WITH_MINIKUBE_OR_K8S_IP>:30704
    ```
    
---

![Imgur](https://i.imgur.com/FMJUjA7.png)

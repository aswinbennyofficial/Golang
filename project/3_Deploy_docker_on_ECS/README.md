## Deploy docker on ECS

- We will be deploying using Fargate
- Open ECS (elastic container service) in AWS console

### Task Definition
- Click on `Task definiREADME.md`
- On the container1 section
    - Give the container URI from ECR and give it a name
    - In port mapping open the port 8082, since we have setup our container to use this
- Leave everything default and click`create`


### Creating cluster
- `Click create cluster` button
- Provide a name for the cluster and choose AWS fargate in the infrastructure settings
- Click `create`
- After the cluster is created, add a new task by clicking `tasks` section and button `run a new task`
- Set everything default
- In the family section, choose the created task definition
- In networking create an existing group and allow TCP connections to 8082 (fargate doesnt support bridging)
- Click create
- Once the task is running click on the task and copy the public URL . Use in browser <public ip>:8082





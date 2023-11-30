## Pushing local docker img to ECR

- AWS ECR or Elastic Container Registry


### Create a container repository
- Search for `ECR` IN AWS and click on create a repository button
- Give a name, rest will be default and click create
- Open the repository in AWS console


### Setting up AWS CLI
- install AWS cli on your device. Use below command on RHEL based OS like fedora, centos, RHEL etc..
```bash
sudo dnf install awscli -y
```
- or it can be installed on below method
```bash
pip install awscli
```
- Configure the awscli using command
```bash
aws configure
```
- It asks for access key ID and access key. 
    - Go to IAM in aws
    - Click on users and create a user with necessary permission
    - Click on the user and `open security credential` tab
    - Click on `create access key` and choose CLI
- Paste the access key id, access key and a default region


### Pushing 
- Open the repo on the ECR in aws console
- Click on `view push commands` button
- follow the commands 

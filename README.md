# Setup Instructions

## Clone Repo and Start Servers

SSH onto a VM and clone code repo. For example,
```
ssh siyuanc3@fa22-cs425-8001.cs.illinois.edu

# In VM 8001
git clone git@gitlab.engr.illinois.edu:siyuan-ruiqi/mp1.git
cd mp1
```
Setup repos in all VM,  start servers with our python scripts. 
We have added public keys of vm 8001 to all the other VMs and public keys of all VMs to the github repo, 
so this process won't require any manual password input.
```
# Clone or pull repos on all VM
python3 setup_repo.py
# Require unzipped demo logs under this folder ahead of time
python3 scp_log.py

# run_servers.py by default start all 10 VMs
# flag --servers=x y z to only start certain servers
python3 run_servers.py
```

## Run grep

Feel free to ssh onto other machines, we stay on VM 8001 for simplicity
```
# IN VM 8001
cd ~/mp1

# Grep on frequent pattern
go run client/client.go -flags=c -pattern=http -prefix=vm
# Grep on rare pattern
go run client/client.go -flags=c -pattern=alexander.com -prefix=vm
# Grep on regular expression 
go run client/client.go -flags=Ec -pattern="Windows N[A-Za-z][A-Za-z0-9]*" -prefix=vm
```

## Test with Fault Tolerance

To stop some servers 
```
# In VM 8001
python3 stop_servers.py --servers 3 4 5

# Grep on frequent pattern
go run client/client.go -flags=c -pattern=http -prefix=vm
# Grep on rare pattern
go run client/client.go -flags=c -pattern=alexander.com -prefix=vm
```

Stop all servers
```
# Similar to run_servers.py, stop_servers stop all servers by default
python3 stop_servers.py
```

## Run Unit Tests or Reproduce Results in Report

```
# In VM 8001, under mp1 folder
# Run all tests under all subdirectories
go test ./...

# Reproduce report data
bash run_exp.sh
```
# Foodtraze Network Installation

Having successfully covered the prerequisites in the previous section, let's dive into the practical aspects of setting up Foodtraze Network Installation.This technical tutorial will guide you through the steps necessary to get your Foodtraze Network instance up and running smoothly.

# Step 1: Download the Source File

Clone the Git repository that contains Foodtraze project files from the path  https://github.com/hyperledger-foodtraze/foodtraze-network.git using "git clone" command

``` bash

git clone -b predev https://github.com/hyperledger-foodtraze/foodtraze-network.git

```
After succesful git clone,you should see a folder foodtraze-network in your directry path. 

# Step 2: Pull Binaries & Docker image

- Install the Hyperledger Fabric platform-specific binaries and config files for the version specified into the /bin and /config directories of fabric-samples
- Download the Hyperledger Fabric docker images for the version specified to your system.

- Get into the cloned directory path(foodtraze-network) and execute the command to download the binaries and images to your system.

``` bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- -- 1.5.6
```
The command above downloads and extract all the platform specific binaries and docker images from Docker Hub in current working directory.

# Step 3: Running the foodtraze network

Inside the directory execute the below command to stand up a foodtraze network.

``` bash
./startNetwork.sh
```
The Foodtraze network has four peer organizations with one peer each and a single node raft ordering service. You can also use the ./network.sh script to create channels and deploy chaincode.

For more information, see Using the Fabric test network. The test network is being introduced in Fabric v2.0 as the long term replacement for the first-network sample.

If you are planning to run the Foodtraze network with consensus type BFT then please pass -bft flag as input to the network.sh script when creating the channel. Note that currently this sample does not yet support the use of consensus type BFT and CA together. That is to create a network use:

``` bash
./network.sh up -bft
```

To create a channel use:

``` bash
./network.sh createChannel -bft
```

To restart a running network use:


``` bash
./network.sh restart -bft
```

Note that running the createChannel command will start the network, if it is not already running.


# Step 4: Using the Peer commands

The setOrgEnv.sh script can be used to set up the environment variables for the organizations, this will help to be able to use the peer commands directly.

First, ensure that the peer binaries are on your path, and the Fabric Config path is set assuming that you're in the foodtraze-network directory.


``` bash
export PATH=$PATH:$(realpath ../bin)
export FABRIC_CFG_PATH=$(realpath ../config)
```


You can then set up the environment variables for each organization. The ./setOrgEnv.sh command is designed to be run as follows.


``` bash
export $(./setOrgEnv.sh Org2 | xargs)
```


Note bash v4 is required for the scripts.

You will now be able to run the peer commands in the context of Org2. If a different command prompt, you can run the same command with Org1 instead. The setOrgEnv script outputs a series of <name>=<value> strings. These can then be fed into the export command for your current shell.



# Step 5 : Chaincode-as-a-service







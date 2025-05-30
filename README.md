[//]: # (SPDX-License-Identifier: CC-BY-4.0)

# FoodTraze

You can use Foodtraze to get started working with Hyperledger Fabric, explore important Fabric features, and learn how to build applications that can interact with blockchain networks using the Fabric SDKs.

## Getting started with the FoodTraze

To use the Fabric samples, you need to download the Fabric Docker images and the Fabric CLI tools. First, make sure that you have installed all of the [FoodTraze prerequisites]. You can then follow the instructions to [Install the Fabric Samples, Binaries, and Docker Images] in the Fabric documentation. In addition to downloading the Fabric images and tool binaries, the Foodtraze samples will also be cloned to your local machine.

## FoodTraze network

The [Foodtraze network]() in the samples repository provides a Docker Compose based foodtraze network with four
Organization peers and an ordering service node. You can use it on your local machine to run the samples listed below.
You can also use it to deploy and test your own Foodtraze chaincodes and applications.

## Asset transfer samples and tutorials

The asset transfer series provides a series of sample smart contracts and applications to demonstrate how to store and transfer assets using Foodtraze.
Each sample and associated tutorial in the series demonstrates a different core capability in Foodtraze. The **Ledger queries** sample demonstrates how to bring all the capabilities together to securely
transfer an asset in a more realistic transfer scenario.

|  **Smart Contract** | **Description** | **Smart contract languages** | **Application languages** |

| [State-Based Endorsement] | This sample demonstrates how to override the chaincode-level endorsement policy to set endorsement policies at the key-level (data/asset level). | [Using State-based endorsement] | golang |


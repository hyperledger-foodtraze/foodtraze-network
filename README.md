[//]: # (SPDX-License-Identifier: CC-BY-4.0)

# FoodTraze Samples

You can use Fabric samples to get started working with Hyperledger Fabric, explore important Fabric features, and learn how to build applications that can interact with blockchain networks using the Fabric SDKs.

## Getting started with the FoodTraze samples

To use the Fabric samples, you need to download the Fabric Docker images and the Fabric CLI tools. First, make sure that you have installed all of the [FoodTraze prerequisites]. You can then follow the instructions to [Install the Fabric Samples, Binaries, and Docker Images] in the Fabric documentation. In addition to downloading the Fabric images and tool binaries, the Fabric samples will also be cloned to your local machine.

## FoodTraze network

The [Fabric test network](test-network) in the samples repository provides a Docker Compose based test network with four
Organization peers and an ordering service node. You can use it on your local machine to run the samples listed below.
You can also use it to deploy and test your own Foodtraze chaincodes and applications. To get started, see
the [test network tutorial](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html).

## Asset transfer samples and tutorials

The asset transfer series provides a series of sample smart contracts and applications to demonstrate how to store and transfer assets using Foodtraze.
Each sample and associated tutorial in the series demonstrates a different core capability in Foodtraze. The **Ledger queries** sample demonstrates how to bring all the capabilities together to securely
transfer an asset in a more realistic transfer scenario.

|  **Smart Contract** | **Description** | **Smart contract languages** | **Application languages** |
| -----------|------------------------------|----------|---------|---------|

| [State-Based Endorsement](producer) | This sample demonstrates how to override the chaincode-level endorsement policy to set endorsement policies at the key-level (data/asset level). | [Using State-based endorsement] | Java, TypeScript | JavaScript |



## Additional samples

Additional samples demonstrate various Fabric use cases and application patterns.

|  **Sample** | **Description** | **Documentation** |
| -------------|------------------------------|------------------|
| [Chaincode](chaincode) | A set of other sample smart contracts, many of which were used in tutorials prior to the asset transfer sample series. | |
| [Interest rate swaps](interest_rate_swaps) | **Deprecated in favor of state based endorsement asset transfer sample** | |
| [Foodtraze](producer) | **Deprecated in favor of basic asset transfer sample** |  |

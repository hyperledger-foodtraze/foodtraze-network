
# FoodTraze

**FoodTraze** is a blockchain-based food traceability application built on Hyperledger Fabric. It enables transparent, tamper-proof tracking of food products across the entire supply chain — from production to distribution to retail.

---

## 📚 Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Why FoodTraze](#why-foodtraze)
4. [Network Participants](#network-participants)
   - [Food Producer](#food-producer)
   - [Food Processor](#food-processor)
   - [Food Distributor](#food-distributor)
   - [Food Retailer](#food-retailer)
5. [Principles, Vision, Goals and Strategy](#principles-vision-goals-and-strategy)
   - [Principles](#principles)
   - [Vision](#vision)
   - [Goals](#goals)
   - [Strategy](#strategy)
6. [Installation Guide](#installation-guide)
   - [Introduction](#introduction)
   - [Prerequisites](#prerequisites)
   - [Network Setup](#network-setup)
7. [Contact](#contact)

---

## Introduction

FoodTraze is an open-source blockchain-based food traceability platform built on Hyperledger Fabric. It enables transparent, tamper-proof tracking of food products across the entire supply chain — from production to distribution to retail. By leveraging blockchain's immutability and distributed ledger technology, FoodTraze ensures trust, safety, and accountability for all stakeholders.

---


## Features

### 🔗 Blockchain-Powered Traceability  
Built on Hyperledger Fabric, ensuring tamper-proof, decentralized recording of every event in the food supply chain.

### 🌾 End-to-End Lifecycle Tracking  
Captures all farm events — from sowing to harvesting, storage, processing, and distribution — with Key Data Elements (KDEs).

### 📱 Farmer-Friendly Mobile App  
Intuitive mobile interface for farmers to easily record real-time data with minimal effort, even in low-connectivity areas.

### 🌍 Multilingual & Global Ready  
Adaptable for different geographies, language requirements, and farming standards.

### 🧪 Ingredient & Lab Test Traceability  
Tracks each ingredient source and allows uploading of lab reports for pesticide, contamination, and quality checks.

### 📦 Consumer Transparency via QR Code  
Generates scannable QR codes showing the full story of the product, enhancing trust and enabling informed choices.

### 🛠️ Open Source & Modular  
Freely available, customizable platform that can scale for different crops, geographies, or sectors like poultry and processed foods.


---

## Why FoodTraze?

- **Immutable Tracking** – Every transaction and update is recorded permanently, creating a verifiable audit trail.
- **End-to-End Visibility** – Producers, distributors, retailers, and regulators can trace the full history of food items in real-time.
- **Decentralized Trust** – Built on Hyperledger Fabric, ensuring privacy, scalability, and security in a multi-party environment.
- **Developer-Friendly** – Modular codebase with Docker-based deployment, REST APIs, and smart contract SDKs for fast integration.

---

## Network Participants

FoodTraze models a real-world food supply chain using distinct participant roles on a Hyperledger Fabric blockchain.

### Food Producer
The starting point of the supply chain — farmers or primary producers who generate raw food materials (e.g., vegetables, dairy, grains). They log harvest details, batch IDs, origin data, and environmental conditions.

### Food Processor
Entities that transform raw materials into consumable or packaged goods. They record data on processing stages, lot combinations, quality checks, and certification (e.g., cleaning, milling, packaging).

### Food Distributor
Responsible for logistics and warehousing. They update the chain with shipment events, transit conditions, handling logs, and destination tracking for each batch of goods.

### Food Retailer
The final node before consumer access. Retailers record stock intake, shelf life, disposal (if any), and sales metadata, ensuring traceability extends to the point of sale.

---

## Principles, Vision, Goals and Strategy

### Principles
- **Blockchain-first Traceability** – Securely recorded transactions on Hyperledger Fabric.
- **Straightforward and Modular** – Simple design and easy extensibility.
- **Data Ownership & Privacy** – User-owned data with no vendor lock-in.
- **Reliable & Auditable** – Fully traceable, enterprise-ready architecture.

### Vision
To empower supply chains with an open, transparent, and trusted platform that enables real-time visibility and accountability for every food item — from farm to fork.

### Goals
- Extreme Traceability & Transparency
- Operational & Regulatory Efficiency
- Enhanced Trust & Market Access
- Promote Sustainable Practices

### Strategy
- Blockchain-Backed Record-keeping
- Inclusive & Collaborative Ecosystem
- QR Code & Consumer Interaction
- Scalable, Open-Source Architecture
- Certification & Compliance-Enabling Functionality

---

## Installation Guide

###  Introduction
FoodTraze is a blockchain-based food traceability application built on Hyperledger Fabric.

- Enables secure, transparent tracking of food products across the supply chain.
- Designed to improve food safety and ensure accountability.
- Each transaction is immutably recorded, enabling real-time traceability and auditability.
- Built with a modular, scalable architecture.

## Prerequisites

Before setting up the FoodTraze application, ensure the following tools are installed on your system:

### ✅ Step 1: Install GIT
Download and install Git:  
🔗 [https://enterprise.github.com/releases](https://enterprise.github.com/releases)

---

### ✅ Step 2: Install cURL
Download and install cURL:  
🔗 [https://curl.se/download.html](https://curl.se/download.html)

---

### ✅ Step 3: Install Docker and Docker Compose

- **Docker Version Required:** 17.06.2-ce or greater  
- **Supported Platforms:** macOS, Linux, Windows 10 (use Docker Toolbox for older versions)  
- Installing Docker Desktop or Docker Toolbox also installs Docker Compose.  
- **Docker Compose Version Required:** 1.14.0 or greater

---

### ✅ Step 4: Install Go

Download and install Go (v1.20.12):  
🔗 [https://go.dev/dl/](https://go.dev/dl/)

---
### ✅ Step 5: Install Node.js & npm

Use the following commands:
```bash
sudo apt-get install nodejs
npm install
```
Once these steps are completed, you are ready to set up and configure the FoodTraze application.


## Network Setup

### ✅ Step 1: Clone the Repository

Clone the official FoodTraze network repository:

```bash
git clone -b predev https://github.com/hyperledger-foodtraze/foodtraze-network.git
cd foodtraze-network
```


### ✅ Step 2: Download Hyperledger Fabric Binaries & Docker Images**
```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- -- 1.5.6
```

### ✅ Step 3: Start the Blockchain Network**
```bash
./network.sh up createChannel -ca -s couchdb
```

This script will:
- Start all Docker containers (peers, orderers, CA, CouchDB)
- Create and join channels
- Install and instantiate chaincode

---

## Contact

Have a question, idea, or want to collaborate? We’d love to hear from you!

- 🌍 **Website**: [www.foodtraze.com](https://www.foodtraze.com)
- 📧 **Email**: [support@foodtraze.com](mailto:support@foodtraze.com)
- 🐛 **Issues**: [GitHub Issues](https://github.com/hyperledger-foodtraze/foodtraze-network/issues)
- 💼 **LinkedIn**: [FoodTraze on LinkedIn](https://www.linkedin.com/company/foodtraze)

> We’re open to contributions, collaborations, and conversations. Reach out anytime!
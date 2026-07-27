# Kubernetes training repository for a "Multi-Tenancy" App

## What is the applicatoin about?

### This consists of a SPA for the frontend and a backend. The backend is connected with a MongoDB for the persistence layer. The backend is exposing a RESP API with basic endpoints to retrieve songs. The frontend implements features to list those songs.

## Goal of the training:

### Build a true Multi-Tenant application that shares infraestructure along all the customers hosted in the cluster. Both frontend and backend should be the same artifacts but know how to handle traffid from differente customers. The Database should be the same but divided into logical databases per tenants.

## Infraestructure Layer:

### This uses Istio as a Gateway to handle incoming traffic and route based on and internal traffic for service descovery and service mesh to monitor and observability.

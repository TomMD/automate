# Event Feed (Actions)

This document explains what a Chef Action is in the context of the Event Feed feature for A2.

## What is a Chef Action?
A Chef Action is a type of message that the [ingest-service](../components/ingest-service) can ingest through its pipeline. You can identify this kind of message when the request has a field that matches to `message_type=action`. (find [here](../components/ingest-service/examples/chef_action.json) an example)

## What type of Chef Actions we support (Entity Type)

There are a number of types of Chef Actions we support, the field in the request is `entity_type` and it can be any of the following:

The Config Management types are:
* Node (`node`)
* Cookbook (`cookbook`)
* Data Bag (`data_bag`)
* Environment (`environment`)
* Role (`role`)
* Policyfile (`policyfile`)

The Compliance types are:
* Profile (`profile`)
* Scan Jobs (`scan_jobs`)

## What specific actions a Chef Action can have (Task)

Every entity type can emit three different actions:
* Create (`create`)
* Delete (`delete`)
* Edit (`edit`)

The associated field in the request is `task`.

## Other fields
You can find other fields description at this [RFC#077](https://github.com/chef/chef-rfc/blob/master/rfc077-mode-agnostic-data-collection.md).

### id
The ID of the Action Message (UUID generated by chef-server)

### request_id
[**[OPTIONAL]** The ID of the request sent to the chef-server that then originates this Action Message

### remote_request_id
**[OPTIONAL]** The ID coming from the client side. (chef-client & knife)

### node_id
**[OPTIONAL]** Used only for Node Delete action

### message_version
The version of the chef-server message

### requestor_name
The name of the client or user that initiated the action

### requestor_type
Was the requestor a client or user? (`client` or `user`)

### recorded_at
The ISO timestamp when the action happened

### entity_name
The name of the entity

### user_agent
The User-Agent of the requestor

### remote_hostname
The remote hostname which initiated the action"

### service_hostname
The FQDN of the Chef server, if appropriate

### organization_name
The name of the org on which the run took place

### data
The payload containing the entire request data

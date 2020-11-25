#!/bin/bash

mockgen -source="pkg/environment/environmentiface/iface.go" -destination="mocks/environment/api.go"
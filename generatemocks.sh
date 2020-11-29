#!/bin/bash

mockgen -source="pkg/environment/environmentiface/iface.go" -destination="mocks/environment/api.go"
mockgen -source="pkg/objects/objectiface/iface.go" -destination="mocks/objects/api.go"
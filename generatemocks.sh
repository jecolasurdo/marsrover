#!/bin/bash

mockgen -source="pkg/environment/environmentiface/iface.go" -destination="mocks/environment/api.go"
mockgen -source="pkg/object/objectiface/iface.go" -destination="mocks/object/api.go"
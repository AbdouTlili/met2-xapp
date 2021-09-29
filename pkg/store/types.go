// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import "github.com/onosproject/onos-api/go/onos/topo"

// Entry is store entry
type Entry struct {
	Key   string
	Value interface{}
}

// StoreEvent is a store event
type EventType int

const (
	// None none cell event
	None EventType = iota
	// Created created store event
	Created
	// Updated updated store event
	Updated
	// Deleted deleted store event
	Deleted
)

func (e EventType) String() string {
	return [...]string{"None", "Created", "Updated", "Deleted"}[e]
}

// RsmE2Node has E2 node information
type RsmE2Node struct {
	RsmE2NodeID       string
	RsmNodeCapability []topo.RSMNodeSlicingCapabilityItem
	//RsmSliceList      []topo.RSMSlicingItem
}

type RsmUe struct {
}

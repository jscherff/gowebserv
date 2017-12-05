// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usbci

import (
	`encoding/json`
	`time`
	`github.com/jscherff/cmdbd/model`
	`github.com/jscherff/cmdbd/store`
)

var ds store.DataStore

func Init(storeName, queryFile string) (err error) {
	if ds, err = model.Prepare(storeName, queryFile); err != nil {
		return err
	}
	return nil
}

type Ident struct {
	Id		interface{}	`db:"id"`
	VendorID	string		`db:"vendor_id"`
	ProductID	string		`db:"product_id"`
	SerialNum	string		`db:"serial_number"`
	HostName	string		`db:"host_name"`
	RemoteAddr	string		`db:"remote_addr"`
}

type Common struct {
	VendorName	string		`db:"vendor_name"`
	ProductName	string		`db:"product_name"`
	ProductVer	string		`db:"product_ver"`
	FirmwareVer	string		`db:"firmware_ver"`
	SoftwareID	string		`db:"software_id"`
	PortNumber	int		`db:"port_number"`
	BusNumber	int		`db:"bus_number"`
	BusAddress	int		`db:"bus_address"`
	BufferSize	int		`db:"buffer_size"`
	MaxPktSize	int		`db:"max_pkt_size"`
	USBSpec		string		`db:"usb_spec"`
	USBClass	string		`db:"usb_class"`
	USBSubClass	string		`db:"usb_subclass"`
	USBProtocol	string		`db:"usb_protocol"`
	DeviceSpeed	string		`db:"device_speed"`
	DeviceVer	string		`db:"device_ver"`
	DeviceSN	string		`db:"device_sn"`
	FactorySN	string		`db:"factory_sn"`
	DescriptorSN	string		`db:"descriptor_sn"`
	ObjectType	string		`db:"object_type"`
	ObjectJSON	[]byte		`db:"object_json"`
}

type Custom struct {
	Custom01	string		`db:"custom_01,omitempty"`
	Custom02	string		`db:"custom_02,omitempty"`
	Custom03	string		`db:"custom_03,omitempty"`
	Custom04	string		`db:"custom_04,omitempty"`
	Custom05	string		`db:"custom_05,omitempty"`
	Custom06	string		`db:"custom_06,omitempty"`
	Custom07	string		`db:"custom_07,omitempty"`
	Custom08	string		`db:"custom_08,omitempty"`
	Custom09	string		`db:"custom_09,omitempty"`
	Custom10	string		`db:"custom_10,omitempty"`
}

type Checkin struct {
	*Ident
	*Common
	CheckinDate	time.Time	`db:"checkin_date"`
}

type SnRequest struct {
	*Ident
	*Common
	RequestDate	time.Time	`db:"request_date"`
}

type Serialized struct {
	*Ident
	*Common
	FirstSeen	time.Time	`db:"first_seen"`
	LastSeen	time.Time	`db:"last_seen"`
	Checkins	int		`db:"checkins"`
}

type Unserialized struct {
	*Ident
	*Common
	FirstSeen	time.Time	`db:"first_seen"`
	LastSeen	time.Time	`db:"last_seen"`
	Checkins	int		`db:"checkins"`
}

type Audit struct {
	*Ident
	Changes		[]byte		`db:"changes"`
	AuditDate	time.Time	`db:"audit_date"`
}

type Change struct {
	*Ident
	PropertyName	string		`db:"property_name"`
	PreviousValue	string		`db:"previous_value"`
	CurrentValue	string		`db:"current_value"`
	AuditDate	time.Time	`db:"audit_date"`
}

func (this *Audit) Create() (int64, error) {
	return ds.Insert(`InsertAudit`, this)
}

func (this *Change) Create() (int64, error) {
	return ds.Insert(`InsertChange`, this)
}

func (this *Checkin) Create() (int64, error) {
	return ds.Insert(`InsertCheckin`, this)
}

func (this *SnRequest) Create() (int64, error) {
	return ds.Insert(`InsertSnRequest`, this)
}

func (this *SnRequest) Update() (int64, error) {
	return ds.Exec(`UpdateSnRequest`, this)
}

func (this *Serialized) Read() (error) {
	return ds.Get(`SelectSerialized`, this)
}

func (this *Serialized) JSON() ([]byte, error) {
	return json.Marshal(this)
}

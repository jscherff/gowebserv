INSERT INTO usbci_checkins (
	host_name,
	vendor_id,
	product_id,
	serial_number,
	remote_addr,
	vendor_name,
	product_name,
	product_ver,
	firmware_ver,
	software_id,
	port_number,
	bus_number,
	bus_address,
	buffer_size,
	max_pkt_size,
	usb_spec,
	usb_class,
	usb_subclass,
	usb_protocol,
	device_speed,
	device_ver,
	device_sn,
	factory_sn,
	descriptor_sn,
	object_type,
	object_json
)
VALUES (
	:host_name,
	:vendor_id,
	:product_id,
	:serial_number,
	:remote_addr,
	:vendor_name,
	:product_name,
	:product_ver,
	:firmware_ver,
	:software_id,
	:port_number,
	:bus_number,
	:bus_address,
	:buffer_size,
	:max_pkt_size,
	:usb_spec,
	:usb_class,
	:usb_subclass,
	:usb_protocol,
	:device_speed,
	:device_ver,
	:device_sn,
	:factory_sn,
	:descriptor_sn,
	:object_type,
	:object_json
)

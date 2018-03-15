-- Dumping structure for procedure gocmdb.proc_usbci_insert_serialized
DROP PROCEDURE IF EXISTS `proc_usbci_insert_serialized`;
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_usbci_insert_serialized`(
	IN `host_name_in` VARCHAR(255),
	IN `vendor_id_in` VARCHAR(4),
	IN `product_id_in` VARCHAR(4),
	IN `serial_number_in` VARCHAR(127),
	IN `vendor_name_in` VARCHAR(127),
	IN `product_name_in` VARCHAR(127),
	IN `product_ver_in` VARCHAR(255),
	IN `firmware_ver_in` VARCHAR(255),
	IN `software_id_in` VARCHAR(255),
	IN `bus_number_in` INT(10),
	IN `bus_address_in` INT(10),
	IN `port_number_in` INT(10),
	IN `buffer_size_in` INT(10),
	IN `max_pkt_size_in` INT(10),
	IN `usb_spec_in` VARCHAR(5),
	IN `usb_class_in` VARCHAR(127),
	IN `usb_subclass_in` VARCHAR(127),
	IN `usb_protocol_in` VARCHAR(127),
	IN `device_speed_in` VARCHAR(127),
	IN `device_ver_in` VARCHAR(5),
	IN `device_sn_in` VARCHAR(127),
	IN `factory_sn_in` VARCHAR(127),
	IN `descriptor_sn_in` VARCHAR(127),
	IN `object_type_in` VARCHAR(255),
	IN `object_json_in` JSON,
	IN `remote_addr_in` VARCHAR(255),
	IN `checkin_date_in` DATETIME

)
    DETERMINISTIC
    SQL SECURITY INVOKER
BEGIN
  INSERT INTO usbci_serialized (
    host_name,
    vendor_id,
    product_id,
    serial_number,
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
    object_json,
    remote_addr,
    first_seen,
    last_seen
  )
  VALUES (
    host_name_in,
    vendor_id_in,
    product_id_in,
    serial_number_in,
    vendor_name_in,
    product_name_in,
    product_ver_in,
    firmware_ver_in,
    software_id_in,
    port_number_in,
    bus_number_in,
    bus_address_in,
    buffer_size_in,
    max_pkt_size_in,
    usb_spec_in,
    usb_class_in,
    usb_subclass_in,
    usb_protocol_in, 
    device_speed_in,
    device_ver_in,
    device_sn_in,
    factory_sn_in,
    descriptor_sn_in,
    object_type_in,
    object_json_in,
    remote_addr_in,
    checkin_date_in,
    checkin_date_in
  )
  ON DUPLICATE KEY UPDATE
    host_name = host_name_in,
    -- vendor_id = vendor_id_in,
    -- product_id = product_id_in,
    -- serial_number = serial_number_in,
    vendor_name = vendor_name_in,
    product_name = product_name_in,
    product_ver = product_ver_in,
    firmware_ver = firmware_ver_in,
    software_id = software_id_in,
    port_number = port_number_in,
    bus_number = bus_number_in,
    bus_address = bus_address_in,
    buffer_size = buffer_size_in,
    max_pkt_size = max_pkt_size_in,
    usb_spec = usb_spec_in,
    usb_class = usb_class_in,
    usb_subclass = usb_subclass_in,
    usb_protocol = usb_protocol_in,
    device_speed = device_speed_in,
    device_ver = device_ver_in,
    device_sn = device_sn_in,
    factory_sn = factory_sn_in,
    descriptor_sn = descriptor_sn_in,
    object_type = object_type_in,
    object_json = object_json_in,
    remote_addr = remote_addr_in,
    last_seen = checkin_date_in,
    checkins = checkins + 1;
END//
DELIMITER ;

-- Dumping structure for procedure gocmdb.proc_usbci_insert_unserialized
DROP PROCEDURE IF EXISTS `proc_usbci_insert_unserialized`;
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `proc_usbci_insert_unserialized`(
	IN `host_name_in` VARCHAR(255),
	IN `vendor_id_in` VARCHAR(4),
	IN `product_id_in` VARCHAR(4),
	IN `serial_number_in` VARCHAR(127),
	IN `vendor_name_in` VARCHAR(127),
	IN `product_name_in` VARCHAR(127),
	IN `product_ver_in` VARCHAR(255),
	IN `firmware_ver_in` VARCHAR(255),
	IN `software_id_in` VARCHAR(255),
	IN `bus_number_in` INT(10),
	IN `bus_address_in` INT(10),
	IN `port_number_in` INT(10),
	IN `buffer_size_in` INT(10),
	IN `max_pkt_size_in` INT(10),
	IN `usb_spec_in` VARCHAR(5),
	IN `usb_class_in` VARCHAR(127),
	IN `usb_subclass_in` VARCHAR(127),
	IN `usb_protocol_in` VARCHAR(127),
	IN `device_speed_in` VARCHAR(127),
	IN `device_ver_in` VARCHAR(5),
	IN `device_sn_in` VARCHAR(127),
	IN `factory_sn_in` VARCHAR(127),
	IN `descriptor_sn_in` VARCHAR(127),
	IN `object_type_in` VARCHAR(255),
	IN `object_json_in` JSON,
	IN `remote_addr_in` VARCHAR(255),
	IN `checkin_date_in` DATETIME

)
    DETERMINISTIC
    SQL SECURITY INVOKER
BEGIN
  INSERT INTO usbci_unserialized (
    host_name,
    vendor_id,
    product_id,
    serial_number,
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
    object_json,
    remote_addr,
    first_seen,
    last_seen
  )
  VALUES (
    host_name_in,
    vendor_id_in,
    product_id_in,
    serial_number_in,
    vendor_name_in,
    product_name_in,
    product_ver_in,
    firmware_ver_in,
    software_id_in,
    port_number_in,
    bus_number_in,
    bus_address_in,
    buffer_size_in,
    max_pkt_size_in,
    usb_spec_in,
    usb_class_in,
    usb_subclass_in,
    usb_protocol_in, 
    device_speed_in,
    device_ver_in,
    device_sn_in,
    factory_sn_in,
    descriptor_sn_in,
    object_type_in,
    object_json_in,
    remote_addr_in,
    checkin_date_in,
    checkin_date_in
  )
  ON DUPLICATE KEY UPDATE
    -- host_name = host_name_in,
    -- vendor_id = vendor_id_in,
    -- product_id = product_id_in,
    -- serial_number = serial_number_in,
    vendor_name = vendor_name_in,
    product_name = product_name_in,
    product_ver = product_ver_in,
    firmware_ver = firmware_ver_in,
    software_id = software_id_in,
    -- port_number = port_number_in,
    -- bus_number = bus_number_in,
    bus_address = bus_address_in,
    buffer_size = buffer_size_in,
    max_pkt_size = max_pkt_size_in,
    usb_spec = usb_spec_in,
    usb_class = usb_class_in,
    usb_subclass = usb_subclass_in,
    usb_protocol = usb_protocol_in,
    device_speed = device_speed_in,
    device_ver = device_ver_in,
    device_sn = device_sn_in,
    factory_sn = factory_sn_in,
    descriptor_sn = descriptor_sn_in,
    object_type = object_type_in,
    object_json = object_json_in,
    remote_addr = remote_addr_in,
    last_seen = checkin_date_in,
    checkins = checkins + 1;
END//
DELIMITER ;
-- DB Generation
CREATE TABLE hum_tem_raw (
	data_id serial4 NOT NULL,
	sensor_id varchar(255) not NULL,
	hum varchar(50) NULL,
	tem varchar(50) NULL,
	updated_at timestamp NULL,
	created_at timestamp NOT null default current_date,
	CONSTRAINT dataId_pkey PRIMARY KEY (data_id)
);

CREATE TABLE device_type (
	device_type_id serial4 NOT NULL,
	type_name varchar(255) not NULL,
	updated_at timestamp NULL,
	created_at timestamp NOT null default current_date,
	CONSTRAINT deviceTypeId_pkey PRIMARY KEY (device_type_id)
);

CREATE TABLE device_list (
	device_id serial4 NOT NULL,
	device_name varchar(255) not NULL,
	device_type_id INT not NULL,
	updated_at timestamp NULL,
	created_at timestamp NOT null default current_date,
	CONSTRAINT deviceId_pkey PRIMARY KEY (device_id)
);


alter table device_list 
  add constraint fk_device_type_id
  foreign key (device_type_id) 
  references device_type (device_type_id);
 
alter table hum_tem_raw  
  add constraint fk_device_id
  foreign key (sensor_id) 
  references device_list (device_id);

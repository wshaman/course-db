package migrations

import "github.com/wshaman/migrate"

func init() {
	migrate.RegisterSQL(20210802093815, 
		"Acme Developer <my-email@acme.corp>",
		"add_locations", 
		`create table locations (
    id serial primary key,
    name text default ''
);

alter table users add column location_id int;
insert into locations (id, name) values (1, 'Normandy'), (2, 'Citadel');
`,
		`
alter table users drop column location_id;
drop table locations;
`,
	)
}

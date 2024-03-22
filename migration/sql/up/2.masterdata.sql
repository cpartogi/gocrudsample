insert into tutorial_types (id, type_name,  created_at, created_by) 
values ('1e815717-301a-4f08-bfc4-45b78e4d78c1', 'Golang',  now(), 'system');
insert into tutorial_types (id, type_name,  created_at, created_by) 
values ('93d899ef-b918-4a94-b7fb-c51df7c7e144', 'PHP',  now(), 'system');
insert into tutorial_types (id, type_name,  created_at, created_by) 
values ('5dca5b91-46cf-49c7-b827-e825f7de04ba',  'JavaScript',  now(), 'system');
insert into tutorial_types (id, type_name,  created_at, created_by) 
values ('89bc2029-1ed4-461c-8c0f-79c9489e04a2', 'Python',  now(), 'system');

insert into tutorials (id, tutorial_type_id, keywords, "sequence",  title, "description", created_at, created_by)
values ('07d790e6-4705-49c1-b61e-126f2800d586', '1e815717-301a-4f08-bfc4-45b78e4d78c1', 'golang, setup', 1, 'Mulai Menggunakan Golang', 'Download Go terbaru dari website lalu install', now(), 'system');

insert into tutorials (id, tutorial_type_id, keywords, "sequence",  title, "description", created_at, created_by)
values ('6cf17700-5a52-4b2c-a9d6-9e843efd163e', '93d899ef-b918-4a94-b7fb-c51df7c7e144', 'php, setup', 1, 'Mulai Menggunakan PHP', 'Download PHP terbaru dari website lalu install', now(), 'system');

insert into tutorials (id, tutorial_type_id, keywords, "sequence",  title, "description", created_at, created_by)
values ('a1cbb992-d507-470b-88a7-ac586c4b9a02', '93d899ef-b918-4a94-b7fb-c51df7c7e144', 'javascript, setup', 1, 'Mulai Menggunakan JavaScript', 'Download JavaScript terbaru dari website lalu install', now(), 'system');

insert into tutorials (id, tutorial_type_id, keywords, "sequence",  title, "description", created_at, created_by)
values ('209bd532-5934-4222-a3af-e3ed16d80b6b', '93d899ef-b918-4a94-b7fb-c51df7c7e144', 'python, setup', 1, 'Mulai Menggunakan Python', 'Download Python terbaru dari website lalu install', now(), 'system');


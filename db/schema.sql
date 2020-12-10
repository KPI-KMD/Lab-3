-- DDL
create table TABLETS(
    id serial primary key,
    name varchar unique not null
);

create table TELEMETRY(
    id serial primary key,
    battery int not null check (0 between 0 and 100),
    deviceTime timestamp not null,
    serverTime timestamp not null,
    currentVideo varchar,
    tablet_id int references TABLETS(id)
);

-- DML
insert into TELEMETRY (battery, deviceTime, serverTime, currentVideo, tablet_id)
values (78, '2020-12-10T14:15:16.123', '2020-12-10:15:16.321', null, 14),
       (10, '2020-12-10T14:15:16.142', '2020-12-10T14:15:16.349', 'dependency injection', 1),
       (56, '2020-12-10T14:15:17.141', '2020-12-10T14:15:17.569', null, 2),
       (46, '2020-12-10T14:15:17.142', '2020-12-10T14:15:17.657', 'git hub architecture', 3),
       (36, '2020-12-10T14:15:17.121', '2020-12-10T14:15:17.123', null, 4),
       (16, '2020-12-10T14:15:17.121', '2020-12-10T14:15:17.123', null, 5),
       (6, '2020-12-10T14:15:17.148', '2020-12-10T14:15:17.149', null, 6),
       (5, '2020-12-10T14:15:17.149', '2020-12-10T14:15:17.150', null, 7),
       (4, '2020-12-10T14:15:17.150', '2020-12-10T14:15:17.151', null, 8),
       (3, '2020-12-10T14:15:17.151', '2020-12-10T14:15:17.152', null, 9),
       (2, '2020-12-10T14:15:17.152', '2020-12-10T14:15:17.153', null, 10);
       


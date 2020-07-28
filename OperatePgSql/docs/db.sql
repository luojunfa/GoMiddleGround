create table user_prop_log
(
 id              integer                     not null default nextval('user_prop_log_id_seq'::regclass),
 game_id         integer                     not null,
 platform_id     integer                     not null default 0,
 user_id         integer                     not null,
 user_prop_id    integer                     default 0,
 update_num      integer                     not null,
 update_type     integer                     not null,
 update_user_id  integer                     not null default 0,
 log_date        date                        not null,
 log_time        timestamp without time zone not null,
 prop_id         integer                     not null default 0,
 raw             jsonb,
 tag             character varying           not null default ''::character varying
)
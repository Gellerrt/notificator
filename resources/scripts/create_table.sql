CREATE TABLE public.notification_data (
id serial PRIMARY KEY,
rquid text UNIQUE,
message text,
need_notification bool
)
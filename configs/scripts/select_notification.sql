SELECT rquid, message FROM public.notification_data nd
WHERE need_notification = true AND sent IS NOT true
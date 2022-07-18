create
or replace function fn_code_stage_modified() returns trigger as $psql$
begin
    if NEW.deleted > (now() - INTERVAL '10 minute'):: timestamp with time zone then
    perform pg_notify(
            'code_confirm',NEW.id_user::varchar
        );
return new;
end if;
return new;
end;$psql$ language plpgsql;

create trigger code_stage before
    update
    on line_owner_validation for each row
    execute procedure fn_code_stage_modified();
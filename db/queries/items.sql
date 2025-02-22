-- name: AddItem :exec
insert into items (
  title,
  description,
  collection
) values (
  @title,
  @description,
  @collection
);

-- name: GetCollections :many
select distinct collection
from items;

-- name: GetItems :many
select *
from items
where
  collection = @collection;

-- name: UpdateItem :exec
update items
set
  title = @title,
  description = @description,
  status = @status,
  collection = @collection
where
  id = @id;

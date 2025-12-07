-- name: ListCategories :many 
select * from categories;

-- name: CateogryExists :one
select Exists(select * from categories where slug = $1) ;

-- name: DeleteCategory :exec
delete from categories where id = $1;

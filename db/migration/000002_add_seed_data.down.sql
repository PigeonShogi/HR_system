DELETE FROM statuses WHERE "name" = '工時未達標準' OR "name" = '工時正常';
DELETE FROM employees WHERE "full_name" LIKE '%(seed)';
DELETE FROM identities WHERE "name" = 'HR-Admin' OR "name" = 'employee';
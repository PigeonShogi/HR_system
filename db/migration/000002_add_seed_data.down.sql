DELETE FROM employees WHERE full_name LIKE '%(seed)';
DELETE FROM identities WHERE name = 'HR-Admin' OR name = 'employee';
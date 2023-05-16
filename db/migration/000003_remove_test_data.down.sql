-- 刪除因執行測試檔而產出的資料
DELETE FROM entries
WHERE employee_id IN (
  SELECT e.id
  FROM employees AS e
  WHERE e.full_name LIKE '%(seed)'
);
DELETE FROM transfers
WHERE from_employee_id IN (
  SELECT e.id
  FROM employees AS e
  WHERE e.full_name LIKE '%(seed)'
);
-- 還原因執行測試檔而變動的資料
UPDATE employees
SET stock = 10000
WHERE full_name LIKE '%(seed)';

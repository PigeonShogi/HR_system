INSERT INTO identities (name) VALUES ('employee');
INSERT INTO identities (name) VALUES ('HR-Admin');
INSERT INTO statuses (name) VALUES ('工時正常');
INSERT INTO statuses (name) VALUES ('工時未達標準');

-- code 始於 S、姓名後加註 (seed) 以表示該記錄為種子資料
INSERT INTO employees (identity_id, code, full_name, stock) VALUES (
    (SELECT id FROM identities WHERE name = 'employee'), 'S2023050001', '王怡君(seed)', 10000);

-- 練習不用 subquery 插入新記錄
INSERT INTO employees (identity_id, code, full_name, stock)
SELECT i.id, 'S2023050002', '王小明(seed)', 10000
FROM identities AS i
WHERE i.name = 'HR-Admin';
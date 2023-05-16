INSERT INTO identities (name) VALUES ('employee');
INSERT INTO identities (name) VALUES ('HR-Admin');
-- code 始於 S、姓名後加註 (seed) 以表示該記錄為種子資料
INSERT INTO employees (identity_id, code, full_name, stock) VALUES (
    (SELECT id FROM identities WHERE name = 'employee'), 'S2023050001', '王怡君(seed)', 10000);
INSERT INTO employees (identity_id, code, full_name, stock) VALUES (
    (SELECT id FROM identities WHERE name = 'HR-Admin'), 'S2023050002', '王小明(seed)', 10000);
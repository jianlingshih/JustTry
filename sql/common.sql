# MySQL避免重复插入
INSERT INTO cms_customers (openid) SELECT "iwnxaninasx" FROM dual 
 WHERE not exists (SELECT * FROM cms_customers WHERE unionid='oR0D703X3FxtkwVYS_djaBxdTk')

 # 简易计算
 SELECT ROUND(1.2333,3)

# 正则 查询 
 select  * FROM cms_customers_child where childid REGEXP ( SELECT CONCAT('^', REPLACE (child_ids, ',', '$|^'), '$')   FROM  cms_customers_order   WHERE orderid='1805171637792028' )

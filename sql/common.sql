# MySQL避免重复插入
INSERT INTO cms_customers (openid) SELECT "iwnxaninasx" FROM dual 
 WHERE not exists (SELECT * FROM cms_customers WHERE unionid='oR0D703X3FxtkwVYS_djaBxdTk')

 # 简易计算
 SELECT ROUND(1.2333,3)

# 正则 查询 
 select  * FROM cms_customers_child where childid REGEXP ( SELECT CONCAT('^', REPLACE (child_ids, ',', '$|^'), '$')   FROM  cms_customers_order   WHERE orderid='1805171637792028' )

#日期查询  
SELECT  FROM_UNIXTIME(1526217150, '%Y-%m-%d %H:%i:%S')

# 带IF查询
SELECT *   FROM  `table` WHERE IF(  `parentID` =1,  `plan_id` <10,  `plan_id` >500 ) LIMIT 0 , 30 

# 表空间占用情况查询
SELECT TABLE_NAME,DATA_LENGTH+INDEX_LENGTH,TABLE_ROWS,concat(round((DATA_LENGTH+INDEX_LENGTH)/1024/1024,2), 'MB') as data
FROM information_schema.tables WHERE TABLE_SCHEMA='tablename' ORDER BY DATA_LENGTH+INDEX_LENGTH desc;

#更新
UPDATE cms_customers_order  co inner join (...) tmp on co.orderid=tmp.orderid set co.type=if(co.skuid like 'V%','3',tmp.type)

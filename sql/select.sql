-- 取得特定 user 的 urls
select id, url, code, created_at, total_click
from user_urls 
where user_id = '4425ff13-354f-4e45-897f-ac76476305d5';

-- 從 urls, user_urls 兩個 table 取得特定 url
SELECT u.url, total_click, user_id
FROM (SELECT url, code FROM urls union SELECT url, code FROM user_urls) AS u
LEFT OUTER JOIN user_urls uu
ON uu.code = u.code
WHERE u.code = 'abc';
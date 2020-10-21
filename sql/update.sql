-- 增加 user url 的 total_click
UPDATE user_urls
SET total_click = total_click + 1
WHERE code = 'abcde';
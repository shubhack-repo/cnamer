# cnamer
Go-lang tool to find cname of hostnames !!

hostnames.txt
---------------------------------
a1.bing.com
bing.com
origin.bn1.bing.com

cat hostnames.txt | cnamer -c 30 #concurrency Flag

Output:
----------------------------------
origin.bn1.bing.com No CNAME record
bing.com bing.com
a1.bing.com a134.lm.akamai.net

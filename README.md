## Experimental Command
 - auth
  ```bash
  pinata auth :auth_key
  ```
  create file for store key api before use other command, put your JWT Key to create file in current folder to save your key for request pinata API.
 - list
  ```bash
  pinata list :file_name :query
  ```
  save result of query to file (or create file if not exists).
 - unpin-hash
  ```bash
  pinata unpin-hash :hash
  ```
  remove/unpin file by hash id from pinata.
 - unpin-query
  ```bash
  pinata unpin-query :query
  ```
  remove/unpin file by query from pinata.

## EXAMPLE UNPIN BY QUERY

```bash
pinata unpin-query "status=pinned"
```
remove/unpin file pinata where status of file is pinned

```bash
pinata unpin-query "metadata[name]=test"
```
remove/unpin file pinata where name in metadata file is test

```bash
pinata unpin-query "metadata[keyvalues][key1]={'value':'valkey1','op':'eq'}"
```
remove/unpin file pinata where keyvalues key1 value is equal 'valkey1'

for more example query you can find in pinata documentation https://docs.pinata.cloud/pinata-api/data/query-files

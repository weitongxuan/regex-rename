# regex-rename

Regex-rename is used to rename files using regex.

## usage

### For help
```
rename -h
```
### Sample command

add prefix "aaa_" to all files end with .txt

```
rename -p ./path -in (./)\.txt -out aaa_\$1.txt
```

note that there is a `\` before `$`. 
* windows users don't need the `\`

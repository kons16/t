# t
t は scp の設定を省略してサーバーにファイルを転送することのできるCLIコマンドです。

```
$ echo "hostname=my_server" >> .direction
$ echo "path=/home/me/src" >> .direction
$ t train.py
```
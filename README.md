# t
t は scp の設定を省略してサーバーにファイルを転送することのできるCLIコマンドです。  
t を実行するディレクトリに、転送先のサーバー情報を記した `.direction` ファイルを用意してください。

```
$ echo "hostname=my_server" >> .direction
$ echo "path=/home/me/src" >> .direction
$ t train.py
```
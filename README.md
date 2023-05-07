# init

1. pwd

```terminal
Projectrootpass/
```

2. ビルド

```terminal
$ docker-compose build
```

3. 起動

```terminal
$ docker-compose up
```

3. 起動（デーモン）

```terminal
$ docker-compose up -d
```

4. ログ
```terminal
$ docker-compose logs -f
```

# もしエラーが出た時の対処方法

1. コンテナを強制終了
```
$ docker-compose kill
```

2. 停止中のコンテナ削除
```
$ docker-compose rm
```

3. ビルドするときにキャッシュを使わない

```
$ docker-compose build --no-cache
```

4. リスタートしてみる
```
$ docker-compose restart
```

5. コンテナ作成
```
$ docker-compose up
```

---

imageとコンテナを全て消してみる

```
$ docker stop $(docker ps -q) && docker rmi $(docker images -q) -f
```

## Docker上DBへWorkBenchを接続する

1. MySQL Connectionsに設定を追加
2. Connection Testでつながっていることを確認
3. 確認できたらOKボタンを押して完了
[参考資料](https://qiita.com/hihats/items/370a195209bf3bef2401)





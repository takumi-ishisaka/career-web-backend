# プロジェクトのビジョン
## 学生が望むキャリアに進むためのアクションを起こすことをサポートする。
## 就活を楽しんでもらう。自分のキャリアを楽しみながら見つけ、楽しみながら前進する。
一人の時間、何となく将来について考えて、不安になってしまった。
「就活やらなきゃ」というときに開くアプリケーションにしたい。

学生にとって就活は今後のキャリアを決める上で最も重要な要素の一つである。
しかし、就活はコトが大きすぎるが故、その重要性が理解できていないことに加えて
不安になることや、何をすれば良いか分からないことで、就活を後回しにする学生が多い。
その結果、就活にかける時間が足りずに自分のキャリアを妥協している学生が多い。

そのため、アクションに対して意味づけをおこなうこと、何をすれば良いか分からない状態を解決すること、
扱うアクションを誰でもできる小さなものにすることで、学生が、早期から就活に役立つようなアクション
を積み上げていくことを実現する。その結果、周りよりも自信を持って楽しく就活に望める状態をつくること
がこのアプリケーションのビジョンである。

# アプリケーション概要
学生生活の間で、就活をやらなきゃと感じるタイミングで、誰でもアクションが起こせるようになるアプリケーション。
基本的には就活生のタスク管理アプリケーション。このアプリの魅力は就活に必要なアクションを誰もがその場で実行できる
粒度まで落とし込んだものを提案してくれること。提案されているアクションの中からユーザがアクションを選び、登録する。
ユーザがタスクを完了したらタスクを振り返り、次のアクションへとつなげる振り返りも行える。


## 要求定義
- タスク管理機能
通知
Googleカレンダー
タスク検索
タスクの振り返り
タスクのカテゴリ分け

- タスクの内容
タイトル
レベル感
取り組み方のコツ
内容
意味づけ
重要度

- モチベーション維持
他者意識
自分のレベル表示
タスク解決のヒント
短期的に終わるような小さなアクションにする
アプリのデザインをおしゃれに

- 意思決定をする
タスクの決定
振り返りで次の行動を決定する（理由を付けて）
徐々に大きな意思決定をできるようにする

- 要望に応える機能
マッチングやチャット
足りない情報を発信する
お問い合せ

- アプリを開く習慣を付ける
〇〇のときはアプリを開くみたいなブランディング
もしくは機能によって開く仕組みを創る。

- アプリを拡散する仕組み
機能を充実
デザインをおしゃれに
UI・UX


# 扱う領域(ユースケース)
## 初期段階
ユーザがタスクを登録する
ユーザがタスクの状態を変更する
ユーザがタスクに対する振り返りを行う
ユーザがタスクを評価する
ユーザがプロフィールを変更する
ユーザが自分がこなしたタスク一覧を見る
ユーザがタスク一覧を見る
ユーザがタスクの詳細を見る
ユーザがタスクを検索する
ユーザがお問い合わせをする
ユーザが自分が行ったタスクの振り返りを見る
タスクがカテゴリー分けされる
広告が掲載される

## 今後の拡張希望
ユーザが自分に合うタスクのレコメンドを受け取る
ユーザが自分の就活偏差値（順位）を見れる
ユーザが就活偏差値ランキングを見れる
ユーザがGoogleカレンダーにタスクを登録できる
ユーザが就活に関する情報を収集できる
ユーザがフレンドと繋がれる
ユーザがマッチングで繋がれる
ユーザが通知を受け取る
ユーザが就活に関してQ&A出来る
ユーザがコミュニティツールと連携できる
ユーザがアクションに挑戦している他のユーザの数を見れる。

--------------------------
# 技術選定

* 言語
Go言語
選定理由：
シンプルな言語仕様。処理スピードの速さがあるためAPI開発においていろいろな企業でAPI開発に使われているから。
また、チューニングをしていくことを考えると並列処理が可能な言語だからこそのチューニング。
[３企業のGOの選定理由](https://geechs-magazine.com/tag/event/20160425)
* フレームワーク
Echo
選定理由：
echoの公式ドキュメントがしっかりしている。middlewareが豊富。処理速度が早い。
リソースの量に比例して全体のスループットが向上するシステム（スケーラブル）でプログラムの外乱に対する抵抗性のある（ロバスト）なRESTAPIを作れる。以上の理由から選定した。
RestAPI向けに最適化されたフレームワークであり、拡張性が高く、カスタマイズが容易である点が特徴的
[echoを選んだ理由](https://note.com/mkudo/n/n6482c47e9708)
[echoを選んだ理由](https://rightcode.co.jp/blog/become-engineer/go-flamework)
* フロントエンド
Vue.js
選定理由：
スケーラビリティ、書きやすさ
[vueの選定理由1](https://blog.ecbeing.tech/entry/2019/05/22/112828)
[vueの選定理由2](https://employment.en-japan.com/engineerhub/entry/2018/09/25/110000)
* DB
MySQL
選定理由：
  * 今まで使ってきたから
  * Twitterなどの大規模情報から少量のデータを取る際に有用
[MySQLとpostgreSQLどっちを使うべき？](https://employment.en-japan.com/engineerhub/entry/2017/09/05/110000#%E7%B5%90%E8%AB%96%E3%81%A9%E3%81%A3%E3%81%A1%E3%82%92%E3%81%A9%E3%82%93%E3%81%AA%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%81%AB%E4%BD%BF%E3%81%86%E3%81%B9%E3%81%8D)
[リレーショナルデータベースとは](https://www.atmarkit.co.jp/ait/articles/0807/16/news149.html)
* Webサーバー
Nginx
選定理由：
Apacheに比べ、メモリ使用量が少ない、並列処理が得意、小機能という
[Nginx vs Apache](https://kinsta.com/jp/blog/nginx-vs-apache/#caching)
* クラウドサービス
GCP
選定理由：
  * シンプルな規模に向いている
[GCPとAWSを運用してみてわかったこと](https://speakerdeck.com/showmurai/aws-vs-gcp-jin-karazuo-ru-naratotutikaiifalse)

# Web開発の流れ

1. 要求モデル(要望の洗い出し)
2. 要求モデル(要求の洗い出し)
3. 要求モデル(要件の洗い出し)
4. ドメインモデル
5. ユースケースモデリング(ユースケース図)
6. 画面設計(画面遷移図)
7. 画面設計(ワイヤーフレーム)
8. ユースケースモデリング(ユースケース記述)
9. ロバストネス分析
10. クラス図
11. データモデリング(ER図)
12. UI設計
13. 実装に入る…

# 具体的に何をしたか

* 要望の洗い出し
  * 
* 要求の洗い出し
  * 
* 要件の洗い出し
  * 
* ドメインモデル
  * 
* ユースケースモデリング
  * 
* 画面遷移図
  * 
* ユースケース図
  * 
* ロバウトネス分析
  * 
* クラス図
  * 
* ER図
  * 
* UI設計
  * 
* 実装
  * バックエンド
  * フロント
  * インフラ

# バックエンド(golangAPI開発)
## 採用アーキテクチャ
レイヤードアーキテクチャを採用

## API仕様
[swagger.io](https://editor.swagger.io/)
## 認証
### JWT
* 使用アルゴリズム選定
今回デフォルトのアルゴリズムであるHS256を採用している。
</br>理由：HS256は共通鍵暗号かをしようしており、TwitterOAuthなどの認証局と認可局が分かれている場合秘密鍵が認証局分流出してしまうためHS256は推奨されない。しかし、今回は認証局認可局が同一になっているため、HS256でも事足りると判断した。
* JWTの保存先
  config.iniファイルに作成

### OAuth
## キャッシュ
Redisを用いてキャッシュをマネジメント。キャッシュのアーキテクチャには2種類ある。Cache-Aside PatternとBroker Patternの2種類
今回はCache-Aside Patternを採用する。[参考文献](https://buildersbox.corp-sansan.com/entry/2019/03/25/150000)
## テスト

# インフラ（Docker・GCP環境構築）

## 設計方針
  * α版(ローカル)
    * OSに依存しない開発環境で行いたい
    * ホットリロードな開発環境
    * ミニマムな設計
  * β版(完璧ではないため、エラーがあっても怒らないでね)
    * スケールは度外視して、デプロイ＝特定のユーザーに使ってもらいUXUIについて情報収集する→最小限のインフラ構築
  * 製品版
    * リクエスト数に応じた自動スケール調整と開発環境と同じ環境にしたい
    * ロードバランサーを使った
    * →kubernates

# 具体的な手順
## α版(Docker)
* docker-compose.yaml
```yaml
version: "3"

services:
  web:
    build: 
      context: ./
      dockerfile: ./build/nginx/Dockerfile
    container_name: career-web-web
    hostname: 'web-dev'
    ports:
      - 8000:80
    depends_on:
      - ap
  ap:
    build: 
        context: ./
        dockerfile: ./Dockerfile #dockerfileの指定
    container_name: career-web-ap
    tty: true
    # volumes: 
    #   - ./:/go/src/
    ports:
      - 8081:8081
    depends_on: #DBに依存している　APIサーバがDBにアクセスする。アクセス方向にdepends_onを指定
      - db
      - redis
    links: 
      - db
    entrypoint: ./wait-for-it.sh -t 60 --strict db:3306 -- ./career-web-backend 
  db:
    image: mysql:latest
    restart: always #コンテナ起動時に自動起動する設定。自動起動させたくない場合はこの記述を削除すれば良い
    container_name: career-web-db #コンテナの名前を決める、あってもなくても良い
    ports:
      - 3306:3306 #どのポートを開放するかを設定":"の左側はホストのポート、右側はコンテナのポート
    volumes: # ./mysqlと言うローカルディレクトりをコンテナの指定ディレクトリにマウント
      - ./db/docker-init:/docker-entrypoint-initdb.d
      - ./db/conf/:/etc/mysql/conf.d
      - ./db/log/:/var/log/mysql
    environment: #環境変数を指定する場合はこのように記述する。
      MYSQL_ROOT_PASSWORD: rootpassword
      TZ: Asia/Tokyo
  redis:
    image: redis:latest
    restart: always
    container_name: career-web-redis
    ports:
      - 6379:6379
```
* dockerfile(golang)
```dokcerfile
# Use the official Golang image to create a build artifact.
    # This is based on Debian and sets the GOPATH to /go.
    # https://hub.docker.com/_/golang
    FROM golang:latest as builder

	MAINTAINER career<allofcareerapp@gmail.com>

    # Copy local code to the container image.
    COPY . /usr/local/go/src/career-web-backend 
	WORKDIR /usr/local/go/src/career-web-backend
	RUN go get -d
    # Build the binary. remove -mod=readonly
    RUN CGO_ENABLED=0 go build -o career-web-backend . 

    # Use the official Alpine image for a lean production container.
    # https://hub.docker.com/_/alpine
    # https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
	# alpineのシェルはbashではなくash
    FROM alpine:latest
    RUN apk add --no-cache ca-certificates
    RUN apk add --no-cache bash

    # Copy the binary to the production image from the builder stage.
    COPY --from=builder /usr/local/go/src/career-web-backend/career-web-backend .
    COPY --from=builder /usr/local/go/src/career-web-backend/config.ini .
    COPY --from=builder /usr/local/go/src/career-web-backend/wait-for-it.sh .
    RUN chmod 777 wait-for-it.sh

    # Run the web service on container startup.
    ENTRYPOINT [ "/bin/ash", "-c" ]
```
* 注意点
  * dockerfileのpath指定にてこずった(rootに置くことでfile not foundを回避)
  * volumeでddl.sqlなどをどのファイルにマウントするかてこずった(docker起動時に走るプログラムmysqldについて書かれたファイルをcontainerの中にbashで入って確認することで特定する)[参考資料](https://noumenon-th.net/programming/2019/04/01/docker-entrypoint-initdb01/)
  * docker-composeで作られた各containerはdocker-compose buildでipが割り振られそれぞれが接続される。ここで、ipは自動割り当てになるので再起動をしたらIPアドレスが変わってしまう可能性がある。自動割り当てされた際にその都度プログラムでIPの設定を変えるのは面倒であるため、具体的なIPの代わりにdocker-composeで指定したcontainer_nameに書き換える。[参考資料](https://qiita.com/tsukapah/items/677b1f5c89dcbe520344)
  * DBが起動する前にgoアプリケーションが起動してしまうと、DB接続ができずエラーが出てしまう。そこでdocker-composeのコンテナ起動制御をするために[wait-for-it.sh](http://docs.docker.jp/compose/startup-order.html)を導入した。
  * [multi stage build](https://qiita.com/minamijoyo/items/711704e85b45ff5d6405)を使ってコンテナの容量を少なくする
* ブラッシュアップすべき点
  * Cloud runを使ってデプロイ
  * Dockerをホットリロードに
  * JWTの秘密鍵をユーザーごとに変える[(KMS)](https://cloud.google.com/kms/docs?hl=ja)

## β版
* アーキテクチャ
 ```
User--(PublicIP)-->GCE(Vue/Nginx)--(PrivateIP)-->GAE(Go)--(PrivateIP)-->CloudSQL(MySQL)
                         |
                         |->GCS(/img)
```
※アーキテクチャは全画面にウィンドウを広げて見てください

### 仕様
#### GAE(Go)
* app.yamlにてVPC-Network-Connectorを作成する
```yaml
runtime: go113  # or go113 for Go 1.13 runtime (beta)

#this conector is the roll that GAE connecting GCE
vpc_access_connector:
  name: "projects/<project-ID>/locations/<region>/connectors/<vpc-connector>"
```
vpc-connectorは`VPCネットワーク->サーバレスVPCアクセス`にて`コネクタを作成`ボタンから作る
注意しなければならないことは`connectorの指定ネットワークはつなげたいGCEと同じネットワークでなければならない（今回はdefaultネットワークを用いた）`
* mysqlへの接続
```go
var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", user, password, instanceConnectionName, database)
```
注意点：
* instanceConnectionNameは`<projectID>:<region>:<db-name>`の形をしている
* `parseTime=true`がないとtime.Time型のデータが型が違うことを理由にデータベースに入らない
* スタンダード環境のGAEはTCPを保証していないため、GAEからCloudSQLへの接続は上記のようになる。
* CloudSQLとGAEとの接続をするには`Cloud SQL Admin API`を有効化しなければならない(APIとサービス→検索→有効化)

#### GCE(Vue/Nginx)
##### Vueの設定
* `/usr/local/src/`にリポジトリをgit cloneしビルドする。ビルドするとdistフォルダが新たに生成される。

##### Nginxの設定
* /etc/nginx/sites-available/defaultにおいて以下の設定を行う。
```
        root /usr/local/src/all_of_career/dist;

        index index.html;

        server_name allofcareer.com;

        location / {
                root /usr/local/src/all_of_career/dist;
                try_files $uri $uri/ /index.html =404;
        }
}
```
Vueプロジェクトをビルドするとdistフォルダにindex.htmlが作成される。このindex.htmlをindexに指定してあげるとvue側で指定されたpass通りに画面遷移することができる。
nginxについて参考にした記事は以下の通り
* [インストールからssl適応まで](https://www.logw.jp/cloudserver/8475.html)
* [aliasとindexについて](https://heartbeats.jp/hbblog/2012/04/nginx05.html)
* [https nginx gce](https://cloud.google.com/community/tutorials/https-load-balancing-nginx?hl=ja)
* [公式ホームページ](https://www.nginx.com/blog/free-resources-for-websites-impacted-by-covid-19/)
* [公式ホームページのブログ](https://www.nginx.com/blog/)

#### CloudSQL
##### sqlファイルのインポート
注意点：
* mysqlのバージョンがGCPではmysql5.7を用いていているため最新のバージョンで書かれたSQL文は構文が違うためエラーになる。そのため、WorkBentchのmysqlバージョン設定を5.7に変更したのちコード生成しなければならない
* 文字コードは日本語で文字化けさせないためにもcharset = utf8mb4：COLLATE = utf8mb4_binにする。
* 設定し終わったら、確認のためCloudShellからmysqlに入ってデータが入っているのか確認する。


# 全体を通した疑問
  1. GCSに入れるイメージのファイル名がバラバラだとめんどくさくない？もしファイル名を変えるとしたら、どうする？
  2. どんなことを考慮に入れてネットワーク構築を行なっているのか
  3. VPCネットワーク設定の際に気を付けることは?
  [Amazon VPC設計時に気をつけたい基本の5のこと](https://dev.classmethod.jp/articles/amazon-vpc-5-tips/)
  [ネットワーク設計](https://www.slideshare.net/serverworks/aws-121012921)



# エラーログ解析
* nginxのaccesslogとerrorlog参照
* mysqlのデータ参照
* mysqlのaccesslogとerrorlog参照
* GAEのlog参照
* `gcloud app logs read`をCloud Shellでうち、ログを見る



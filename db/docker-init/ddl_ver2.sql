-- MySQL dump 10.13  Distrib 8.0.18, for Linux (x86_64)
--
-- Host: localhost    Database: career_db
-- ------------------------------------------------------
-- Server version	8.0.18-google

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `career_db`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `career_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `career_db`;

--
-- Table structure for table `action`
--

DROP TABLE IF EXISTS `action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `action` (
  `action_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `category_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `title` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(512) COLLATE utf8mb4_bin NOT NULL,
  `standard_time` varchar(16) COLLATE utf8mb4_bin NOT NULL,
  `action_type` int(11) NOT NULL,
  `url` varchar(512) COLLATE utf8mb4_bin DEFAULT NULL,
  `after` varchar(512) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`action_id`),
  KEY `tag_id_idx` (`category_id`),
  CONSTRAINT `fk_category_action_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `action`
--

-- LOCK TABLES `action` WRITE;
-- /*!40000 ALTER TABLE `action` DISABLE KEYS */;
-- INSERT INTO `action` VALUES ('action_0000','category_00','振り返りのアクションをするモチベーションを付けよう。','ネットの記事を調べ振り返りの重要性を認識し、なぜ自分に必要なのかをまとめよう。そしてその答えは振り返りの「やったこと」欄に記述しよう」','45分',2,'https://www.katsuiku-academy.org/media/effectivereflection/','振り返りのスキルを磨こうというモチベができることを目指しましょう。'),('action_0001','category_00','メンターがフィードバックをしやすい振り返りの心構え。','メンターの気持ちになって振り返りを書いてみよう。評価するのに必要な情報が含まれているか、論点はどこかを考えて行うと非常に有効です。例としてあなたが味噌汁をつくったとして、振り返りを書きましょう。','15分',2,'','メンターへの気遣いができるようになりましょう。'),('action_0002','category_00','振り返りのフレームワークを知ろう。','振り返りをより効率よく行うためのフレームワークを参考URLの記事をみて勉強しよう。そして、好きなフレームワークを一つ選び、昨日の行動に対してQOLを向上させることを目的として振り返りをしよう。','30分',2,'https://www.kikakulabo.com/tpl-reflection/','世の中にある洗練された振り返り方法があることを知り、形だけでも使えるようになることを目指しましょう。'),('action_0003','category_00','PDCAサイクルを強化しよう。','あなたが最近提出した学校の課題に対してPDCAを回してみよう。やったこと欄に以下のP~Aまでの4項目を書きましょう。P...何が目的でどのような行動をしたのか。D...計画を実行する。C...実行して得られた結果とアクションの目的を比較し、不足している部分を明らかにする。A...不足分を補えるアクションに修正する。（そして実行する。）反省欄には今回自分が行ったPDCAの回し方の反省をしましょう。','60分',2,'','あなたのPDCAサイクルにおいて、改善ポイントを見つけることを目指しましょう。'),('action_0004','category_00','①振り返りの「やったこと」欄の質を高めよう。','「学生時代に最も力を入れたこと」を書く作業でどのようなことをやったのかをやったこと欄に書きましょう。その際に①何が目的で②具体的にどのようなアクションを③どんなことに意識しながら行ったのかを示してください。','60分',2,'','やったことを反省に繋げやすくなることを目指しましょう。'),('action_0005','category_00','②振り返りの「次のアクション」欄の質を高めよう。','「学生時代に最も力を入れたこと」を書く作業の反省でみつかった問題点を解決するために必要なアクションを次のアクション欄に書きましょう。反省で見つかった問題点を解消する方法は無数にありますが、その中からより適したアクションを選択しその理由も付けれるようにしましょう。','60分',2,'https://infinity-agent.co.jp/lab/logic-tree/','なぜそのネクストアクションを選択したのかを説明できることを目指しましょう。'),('action_0006','category_00','③振り返りの「反省」欄の質を高めよう。','「学生時代に最も力を入れたこと」を書く作業での反省点を書きましょう。ここではアクションのゴールを達成した場合でも達成していない場合でもその理由を書きましょう。','60分',2,'','次のアクションに繋がる反省ができるようになることを目指しましょう。'),('action_0100','category_01','自己分析① 自己分析を知ろう！','以下のurlの記事を読んで自己分析の意義、自己分析では何をやるのかのイメージをつけましょう。','30分',2,'https://note.com/maron0917/n/n9d554d1fd90b','自己分析が就活のためだけではなく、人生においても重要という感覚を持ち、自己分析を行うモチベーションがつくられることを目指しましょう。。'),('action_0101','category_01','自己分析② 自分の中で一番大事にしている価値観を仮定しよう！','自分の一番深い位置にある価値観（憧れられたい、自信を付けたい等）を仮定してみましょう。また、仮定した価値観がなぜ生まれたのか（原因）を示すエピソードを探してみよう。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','自分が仮定した価値観への納得度が50%以上になることを目指しましょう。'),('action_0102','category_01','自己分析③ 仮定した根源的欲求を掘り下げてみよう！','「自己分析②」で仮定した価値観よりも奥にある価値観がある可能性があります。仮定した価値観をなぜ満たしたいのか（目的）を自問自答してみましょう。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','自分が仮定した価値観への納得度が70%以上になることを目指しましょう。'),('action_0103','category_01','自己分析④ 深掘りされた9つの価値観のどれに当てはまりそうか考えてみよう！','以下のurlの記事にある9種類の根源的欲求を把握し、深掘りした価値観がどれに当てはまるのか見てみよう。当てはまりそうにない場合は、まだ深掘りの余地があるかも知れないです。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','9種類の根源的欲求に対して自分なりに優先順位を付けることを目指しましょう。'),('action_0104','category_01','自己分析⑤ 根源的欲求のルーツを探ろう！','自分の根源的欲求はなにがきっかけでうまれたのか（原因）を知りましょう。自分の家庭環境や子供のころの印象に残っているエピソードは説得力が強いです。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','仮定した根源的欲求への納得度が50%以上になることを目指しましょう。'),('action_0105','category_01','自己分析⑥ 自分の欲求に基づく一貫した行動を一言で表してみよう！','色々なエピソードを振り返り、一貫している部分を言語化してみましょう。例「安定していたい」 → 「何かに依存しないでも生きていけるように自分を磨く」等','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','自分の一貫している軸が言語化できることによって、根源的欲求の納得度を上げることを目指しましょう。'),('action_0106','category_01','自己分析⑦ 自分が定義した根源的欲求の納得度を90%以上にしよう！','過去の特徴的なエピソードにおける自分の動機が遠回しにでも根源的欲求に結び付いているかどうかを確認してみましょう。また、価値観が動機となって形成された自分の強みを探ってみましょう。例「なにかに依存しないでも生きていける能力を磨く」 → 強みは「ストイック」や「上昇志向」等。','30分',1,'https://note.com/maron0917/n/n9d554d1fd90b','自分の強みを根拠をもって話せるようになることを目指しましょう。'),('action_0107','category_01','自己分析⑧ 働くことで成し遂げたいことを見つける方法の流れを知ろう！','urlの記事の「あなたが人生をかけて成し遂げたいことは何か」を読んでアウトラインを把握しましょう。','30分',1,'https://note.com/maron0917/n/n9d554d1fd90b','働くことで成し遂げたいことを決める際になにをすればいいかわからない状態を抜け出すことを目指しましょう。'),('action_0108','category_01','自己分析⑨ 根源的欲求から働くことで成し遂げたいこと見つけてみよう！','自分の根源的欲求が満たされない状況や要因を挙げてみましょう。なるべく抽象度の高いものが良いです。例「安定していたい」→「安定を構成する要因を自分でコントロールできない状態」等。難しいので何回も壁打ちするのがいいと思います。','120分',1,'https://note.com/maron0917/n/n9d554d1fd90b','例のように自分なりの言葉で、そして抽象度の高いものを挙げられることを目指しましょう。難しければ要因を複数個挙げることを目指しましょう。'),('action_0109','category_01','自己分析⑩ 自分にとって重要な問題を解決する手段を決めよう！','「自己分析⑨」で挙げた問題を解決する手段を複数個挙げてみよう。その中から自分が一番モチベーションを維持できそうなモノを選びましょう。モチベーションに関しては参考リンク内のモチベーションの構成図を参考にしてください。','120分',1,'https://note.com/maron0917/n/n905030b895b5','自分のモチベーションを最も引き出せる目標を見つけることを目指しましょう。'),('action_0110','category_01','自己分析⑪ 目標を自分のモノにしよう！（オプション）','「自己分析⑩」で挙げた目標を自分の言葉で表現しよう。あなただからこそ出てくる言葉で表現しよう。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','働くことで成し遂げたいことを自分の言葉で表現できることを目指しましょう。'),('action_0111','category_01','自己分析⑫ ビジョンを実現する方法をイメージしよう！','具体的にどんな業界があって、それぞれがどのような課題にアプローチしているのか、それによってどのような影響を与えるかを調べてみましょう。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','ビジョン達成するための手段を仕事に落とし込むことを目指しましょう。'),('action_0112','category_01','自己分析⑬ 自分が決めたビジョンを合理化してみよう！','仮定した「働くことで成し遂げたいこと」のやる意義を3つ見つけましょう。合理化することで自信を持つことができ、モチベーション維持に大きく影響します。','60分',1,'https://note.com/maron0917/n/n9d554d1fd90b','自分が決めたビジョンに対して他人に話せるくらいに自信を持てることを目指しましょう。'),('action_0116','category_01','他人と自己分析をして、自分の強みと弱みを相対的に評価しよう！','自分の強みや弱みが他者と完全に一致することは無いので、どのように違うのかを細かく分析してみましょう。また、他人の強みや弱みを知ることで相場感というものを付けることができるとなおよいです。','30分',1,'','自分の強みの固有性を話せるようになることを目指しましょう。'),('action_0117','category_01','自分の3年後までのキャリアプランを仮定してみよう！','自分の人生の目標から逆算して3年後までにはどのようになっていなければいけないのかを400字以内でまとめてみましょう。気になる企業の社員インタビューを見るとイメージが付くかもしれません。','180分',1,'','業界や職業、役職等まで明確になっていることを目指しましょう。'),('action_0200','category_02','学生時代頑張ったことを書いてみよう！','ESや面接で聞かれる「学生時代頑張ったこと」を400文字で書いてみましょう。ここでは自分の人間性や将来のビジョンと一貫性を持たせた内容にするのが理想です。また結論ファーストを意識することと、構成はPDCAを回していることが分かるように書きましょう。振り返りの「やったこと」欄に学チカを400字以内で記入してください。','180分',2,'https://shukatsu-mirai.com/archives/88839','自分がどれくらい頑張れているのかを把握することを目指しましょう。理想は次のアクションプランが見つかることです。'),('action_0201','category_02','自分が気になる業界の課題を見つけよう！','自分が気になる業界を一つ挙げ、その業界が解決しようとしている問題やその業界自体の問題を3つ挙げよう。そして自分がその問題を解決することに魅力を感じるかどうかを直感的に考察してみよう。','60分',2,'https://gyokai-search.com/','自分が興味を持っている企業や業界を志す一つの理由が見つかることを目指しましょう。'),('action_0202','category_02','自己PRを書いてみよう！','面接やESでよく聞かれる自己PRを400文字で考えましょう。自己分析で深掘りした内容を人に伝えられる形に整理しよう。結論ファーストや、伝わりやすい構成を意識しながら作りましょう。また、一番伝えたいことをあらかじめ決めておけばコンパクトにまとまると思います。振り返りの「やったこと」欄に自己PRを400字以内で記入してください。','180分',2,'https://www.kurihaku.jp/2021/special/entrysheet/464/','自分の魅力を把握し、業界、企業選びの一つの軸ができることを目指しましょう。'),('action_0203','category_02','志望企業の面接の質問内容を知ろう。','ネット検索で志望企業でどんな質問をされているかを調べ、自分なりの回答を簡易的に作ろう。','45分',2,'https://syukatsu-kaigi.jp/','面接で聞かれうる質問を把握している状態になることを目指しましょう。理想はいろんな質問に対しても一貫した回答ができるようになることです。'),('action_0204','category_02','あなたの弱みを語ろう！','あなたの弱みを400文字で書こう。結論ファーストを意識することと、弱みを克服するためにどのようにアクションをしているか、またはする予定かもまとめましょう。振り返りの反省の部分に400文字以内で記入してください。','180分',2,'url','自分の欠点を把握し、どのように克服していくかのアクションプランが立てれることを目指しましょう。');
-- /*!40000 ALTER TABLE `action` ENABLE KEYS */;
-- UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `category_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `goal` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

-- LOCK TABLES `category` WRITE;
-- /*!40000 ALTER TABLE `category` DISABLE KEYS */;
-- INSERT INTO `category` VALUES ('category_00','振り返り力','振り返り能力は行動を改善し、物事を前進させていくために必要なスキルです。全ての人におススメですが、特に「何をやってもうまくいかない」や「成長をなかなか実感できていない」という人におすすめのカテゴリです。またアプリ内の振り返り機能をフルに使えるようになる事で、より効果的な学習が可能になります。思考力との相性がいいです。'),('category_01','自己分析','自己分析は自分の快楽と苦痛を知ることで自分が進むべき道（目標）を明らかにする行為です。就活をやっている人だけでなく、全ての人がすぐにやるべき内容です。自分のやりたいことが見つかったり、相手の視点に立って物事を考えられるようになります。'),('category_02','選考対策','就活の準備としてやっておいた方が良いことがまとめられています。本選考やインターン選考に進む予定がある人はやるべきです。大切な機会を棒に振るわないようにしましょう。'),('category_03','思考力','思考力とは問題を発見し問題解決の過程を考え、最適な選択肢を判断する力です。全ての人が優先的にやる必要があります。振り返りで見つかった問題をうまく解決できたり、物事の本質を見極めたりできるようになります。振り返り力と相性がいいです。');
-- /*!40000 ALTER TABLE `category` ENABLE KEYS */;
-- UNLOCK TABLES;

--
-- Table structure for table `feedback`
--

DROP TABLE IF EXISTS `feedback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `feedback` (
  `feedback_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `user_action_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `comment` varchar(1024) COLLATE utf8mb4_bin NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`feedback_id`),
  KEY `fk_feedback_user_action1_idx` (`user_action_id`),
  CONSTRAINT `fk_feedback_user_action1` FOREIGN KEY (`user_action_id`) REFERENCES `user_action` (`user_action_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feedback`
--

-- LOCK TABLES `feedback` WRITE;
-- /*!40000 ALTER TABLE `feedback` DISABLE KEYS */;
-- /*!40000 ALTER TABLE `feedback` ENABLE KEYS */;
-- UNLOCK TABLES;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `profile` (
  `user_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `university` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `major` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `graduation_year` int(11) NOT NULL,
  `aspiring_occupation` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `aspiring_field` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `sentence` varchar(1024) COLLATE utf8mb4_bin DEFAULT NULL,
  `image_path` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL,
  `job_hunting_status` int(11) NOT NULL DEFAULT '0',
  `deviation_value` decimal(20,10) NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `user_id_idx` (`user_id`),
  CONSTRAINT `fk_user_profile_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profile`
--

-- LOCK TABLES `profile` WRITE;
-- /*!40000 ALTER TABLE `profile` DISABLE KEYS */;
-- INSERT INTO `profile` VALUES ('1b500546-e408-47f1-87be-eeca9c758b07','松下雄哉','静岡大学','情報学部',3,'エンジニア/コンサルタント','コンサルタント','','undefined',0,0.0000000000);
-- /*!40000 ALTER TABLE `profile` ENABLE KEYS */;
-- UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `user_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `password` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `status` int(11) DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

-- LOCK TABLES `user` WRITE;
-- /*!40000 ALTER TABLE `user` DISABLE KEYS */;
-- INSERT INTO `user` VALUES ('09312f4d-e2b0-49a4-9d38-114f77b8c020','y_rainbow25@icloud.com','$2a$10$Qr0.mUeJiVLRsIAEX3soU.wan5w60smg8Q5b.g8Nofktnugk2SufG',1,NULL),('12b061a2-f10b-4c50-8be5-40b6eba2d321','torikatsu923@gmail.com','$2a$10$IiSXL.4Zyjgwod8KGGB6dOvLiHEv5fnfzfmcLSRkTUS8TLknhy6JW',1,NULL),('16af2f2c-6bde-44f5-ae5c-09aa5f7663b3','gantaidaomu@gmail.com','$2a$10$jRCaleYuIiGjXBviUfBJ2uMk1Z6IoQcMFXq4cDvkykyiM9eFZCE/e',1,NULL),('1911e565-35c9-41cf-afad-1f2eea3be752','hiroki0506.lions1999@gmail.com','$2a$10$E9T9ODle4Mh3VapTgsW9v.PeKXtKmfL8sHQEeljPEQpHbk7jCxtYC',1,NULL),('1b500546-e408-47f1-87be-eeca9c758b07','yuyachiko821@gmail.com','$2a$10$1nQbx/wox8ELHGkpn1ja0.pcdkOdDAGqHIjOzStCBsvQ9tJwRLGHu',1,NULL),('23847f19-c36a-4e32-9849-cd59e2b8e023','water.0316@icloud.com','$2a$10$jpi32ifz9sM0WYFzvZoIbOyn7Z2ezUtNcE/GupHO6/8DCEEH07Grq',1,NULL),('3dafb0b9-320e-46e6-a1fd-5c9f0ec4ba7a','test_account@gmail.com','$2a$10$tLsxqD/I2056t0vFTQz9KeOId.6A5k4iNewkFofaRlzoB7nyVsHWK',1,NULL),('3f98ffd0-b96f-48b5-9fa0-70460c63de89','tkinoken19@icloud.com','$2a$10$ozg0wXaJnWVELXJGcyFCjeAWXXz46zcCCcoR0xvnCFkQftv//wwdC',1,NULL),('485018b4-57c4-4207-aa10-31598345419c','khmck7@gmail.com','$2a$10$sIbIbSMqI5KHX0v9RFAyI.7XavmqbzPHCWXlayAY4TwHn71m2JySu',1,NULL),('51b6748b-a53a-437f-931e-6463f3b9a559','mariyui-chokora@i.softbank.jp','$2a$10$g4jaXuBuyMBK4Kq.LxEZmOQaobGM1NSXfLraNqHTEPgYsErjipm8i',1,NULL),('53f5a1a5-0774-4d4c-8e83-951cb5ff36b7','takahirodesu0812@yahoo.co.jp','$2a$10$5sdje.vls.dNbXwD/04kQOwJ3BqeC2aW7JCLiHgEL.4FOD40x9tUe',1,NULL),('570a3807-81ed-4ef1-b1f1-d384c30bacb3','ngw1p3@gmail.com','$2a$10$bxuhRPTlb/hp1k/Bw1t0vu49Sl.FWWU8ZcqaQ0FKpV/kk9TV4lROm',1,NULL),('5999aaef-66e0-4369-aa2f-457de40fb97c','nba.nba.chako@gmail.com','$2a$10$QJRS6nlMk4L16Won6VgF7OOg/4TafpIapyZeTJv0kRoX6xKH5.C0a',1,NULL),('699a143b-faf0-4016-a921-458556bf8f5d','abe.yy.0225@outlook.com','$2a$10$9.jGNh/40IIrOQH1KrLds.4guv2IfakGBXoYYAL0csMKYviDMFsEG',1,NULL),('708c6427-a5f5-4a5c-ab3a-c00c19a1e0e7','ajf23y@gmail.com','$2a$10$Yd/T2gGwzbbDKjf8RB2r5.ejwMaYjLZk5.ZotADQ1m0KhdBmtApuy',1,NULL),('709e07db-5837-451e-a2d2-265bdc8f8a73','allofcareer@gmail.com','$2a$10$mwhkDb2bg6ARnHi29BPbTuHZwEETyUUbIf/WpDbFuii9AJNxwj04K',75392,NULL),('7671b853-93e3-4ba2-b657-688070887a58','kokokikutetote@gmail.com','$2a$10$LrDE1DjpPyndN65/5Ti.HubuCH4/E8kvcd29KjH6zL.uri0rI5dO6',1,NULL),('7e0a0a60-f169-4265-801d-f36ab143e4d1','m5203781040@lost-corona.com','$2a$10$4.02oIbYikRwdp2Ts44xD.P22CyBpXO8t6YJq20AzdTro8jtyW.ai',1,NULL),('861e9b18-fed1-43ad-a01c-ff9146d9023c','k.iwakami@gmail.com','$2a$10$wIdliZc6/TawStkNHi3D8OlxjAkAKytvmCem4vc27AR19g75klxxO',1,NULL),('8a8d1539-45df-4013-9dd0-08ca873ecb34','kai.acm092@gmail.com','$2a$10$PBqR16y19GvSZXunxeUTUu6llCQt2z1vLtByj09C8rdsHmm7B2LHe',1,NULL),('9ef2629a-5b12-46e8-a924-84dc6784a24f','takatomomomo@gmail.com','$2a$10$W4Su7yuOnbweHRQP0Q348.37/Tec.GD1KauRpJw27HK7Eq8pxFgLO',1,NULL),('a0df79c6-4964-4e73-97c1-7624ad05b3e6','oda.naoya.18@shizuoka.ac.jp','$2a$10$3qmYWelYkETpw0UAQRdpZOOr9FA4W8zVtLs8zkQeLJuFzWmXgJfEK',1,NULL),('b60d10c0-3b0b-4cb0-9519-e646504bd77b','kobayashi.shuno.18@shizuoka.ac.jp','$2a$10$Z5MS2Voe4mUkUhzTKo696efKSOcyuPnVri9DRctq25rtNpIvxiOS.',1,NULL),('c2be1480-740f-4fbb-ae0f-6b7ce0503d52','toti-xm.1204@au.com','$2a$10$B4x6hEMD/cUyOaeKAs7CquwN5VxWS4clSxOZ1EVlfJH3DBW/J2JLq',1,NULL),('c69d9c4b-5551-432f-b6d7-03aa2497a868','ozuwarudo1115@icloud.com','$2a$10$YFUWnpztZ/ylZ4G/NwcQmupnnTDP.0FFLtdhi3CRQmes/ibLQtmFq',1,NULL),('c99cea82-4f8e-4ebd-b0b9-3f5b2ec15178','ku091056@icloud.com','$2a$10$GPC1tumynfuo5oTnJ9rneuDH2m2SU2kgZHKcJZ2nmnw65Lp0XHAxO',1,NULL),('cba7062f-0a68-4f86-82fb-daba3bc016d6','fcgifuyuushou@yahoo.co.jp','$2a$10$cdGsSPBbqOQhVEzDxbFN7.sDFKR2WR3LzD5Pu3FRqjiZ55KlkN30i',1,NULL),('d333fdfc-c1de-4d59-aef0-2ce1f398a73f','morioka.miki.18@shizuoka.ac.jp','$2a$10$wedqrX94iaPfy80pX6VXu.pMCT8GxMoRt/FZrdDRW9.0ULIXQhVLW',1,NULL),('d915dbe8-7990-4fe0-b12d-4a6971e3bfa7','udon.yuto@gmail.com','$2a$10$PXEA4sxOgSu1WaMdZqCaP.HW6/HIwPbesP2mMtKGuNIOlw0EooNOa',1,NULL),('dbe513f3-fc16-4aa3-afd0-d2748eec290e','qyl21060@cuoly.com','$2a$10$7XZwoR0hUmOouAaGmN9UGOM6/tOrO4ADJx72iPAGMjLFwVx4rbttC',1,NULL),('e5780f6d-2df4-43d8-80dc-a1aee9aa13a6','skimura3102@gmail.com','$2a$10$RFh5JJbQPmUENtlB.zcojOPD0vBReyuSWLZnhe8aElg0b7beisKC.',1,NULL),('f372e49d-af5e-4c72-88f0-8bf3b684c2ff','hanayabu1011@gmail.com','$2a$10$YjgEBV.MzyxUPQaYBuVDLujYquZ0azIE7ktKjUpzpaAo3vHe6YnZ.',1,NULL),('f3ebf622-15a8-4182-977d-5fa7951f3e60','yuki.yamamoto0915@gmail.com','$2a$10$bdoZiR.6uirrCDYmrVyf3eQ7sa18yK4BCpi/xqES9X7HndhyqccUu',1,NULL),('ff34e6a2-6d35-497e-aafd-6c596741b0f6','shunya358@yahoo.co.jp','$2a$10$58nGog/VFu4raptcsOQoY.vyVwWrU603cvIZpjNCxi8siWTw7nhZ2',1,NULL);
-- /*!40000 ALTER TABLE `user` ENABLE KEYS */;
-- UNLOCK TABLES;

--
-- Table structure for table `user_action`
--

DROP TABLE IF EXISTS `user_action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_action` (
  `user_action_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `action_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `status` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `do` varchar(1024) COLLATE utf8mb4_bin DEFAULT NULL,
  `reflection` varchar(1024) COLLATE utf8mb4_bin DEFAULT NULL,
  `next_action` varchar(1024) COLLATE utf8mb4_bin DEFAULT NULL,
  `evaluate_value` int(11) DEFAULT NULL,
  PRIMARY KEY (`user_action_id`),
  UNIQUE KEY `user_action_id_UNIQUE` (`user_action_id`),
  KEY `action_id_idx` (`action_id`),
  KEY `user_id_idx` (`user_id`),
  CONSTRAINT `fk_action_user_action` FOREIGN KEY (`action_id`) REFERENCES `action` (`action_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_user_action_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_action`
--

-- LOCK TABLES `user_action` WRITE;
-- /*!40000 ALTER TABLE `user_action` DISABLE KEYS */;
-- INSERT INTO `user_action` VALUES ('0256acdd-4c3c-4162-96cc-57a99df3bf36','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0000',1,'2020-11-09 09:20:31',NULL,NULL,NULL,NULL),('0683da97-4108-467b-a673-967f1e9b4fdd','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0002',1,'2020-11-09 11:12:07',NULL,NULL,NULL,NULL),('0b3ca3ca-7bc4-492f-9c27-cddbc9f2c0ed','1b500546-e408-47f1-87be-eeca9c758b07','action_0200',1,'2020-11-11 14:36:44',NULL,NULL,NULL,NULL),('1f550d70-c0f1-41e1-a2bc-ed140562fd01','1b500546-e408-47f1-87be-eeca9c758b07','action_0001',1,'2020-11-12 15:48:04',NULL,NULL,NULL,NULL),('402c7b9e-012d-4666-aaf3-ea081a468a7f','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0005',1,'2020-11-10 16:49:59',NULL,NULL,NULL,NULL),('4464f958-5a40-4270-b1b0-e04c92efcfce','8a8d1539-45df-4013-9dd0-08ca873ecb34','action_0200',1,'2020-11-02 07:00:45',NULL,NULL,NULL,NULL),('457a25ef-e3d4-4bbc-8a86-cdce2bef4f10','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0001',1,'2020-11-09 11:12:00',NULL,NULL,NULL,NULL),('58120519-b658-4801-8303-f08af7ee8e35','3dafb0b9-320e-46e6-a1fd-5c9f0ec4ba7a','action_0002',1,'2020-10-28 06:05:02',NULL,NULL,NULL,NULL),('5cc01e64-8f5e-4eeb-af4d-7f49d82fd33f','3dafb0b9-320e-46e6-a1fd-5c9f0ec4ba7a','action_0200',2,'2020-10-27 20:43:41','abc','abc','abc',3),('5e2e1b2a-d5a9-4263-bfa8-18d49e15d759','3dafb0b9-320e-46e6-a1fd-5c9f0ec4ba7a','action_0100',1,'2020-10-29 08:12:03',NULL,NULL,NULL,NULL),('7619cb72-25fc-4c13-8790-f1456dffd95f','8a8d1539-45df-4013-9dd0-08ca873ecb34','action_0201',1,'2020-11-02 07:00:52',NULL,NULL,NULL,NULL),('a6df1ddb-9804-43f2-ac01-8ae22b654c1c','1b500546-e408-47f1-87be-eeca9c758b07','action_0100',1,'2020-11-11 14:31:53',NULL,NULL,NULL,NULL),('afbabb8f-500c-457d-be3a-ccac7890fe81','8a8d1539-45df-4013-9dd0-08ca873ecb34','action_0000',1,'2020-10-30 13:30:26',NULL,NULL,NULL,NULL),('c76fb26f-20f4-4659-9131-8bbac8b9a49b','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0003',1,'2020-11-09 11:12:12',NULL,NULL,NULL,NULL),('d1846740-d634-4c53-b4af-71384cda1d91','23847f19-c36a-4e32-9849-cd59e2b8e023','action_0200',1,'2020-11-05 06:44:53',NULL,NULL,NULL,NULL),('dbe8ebeb-adc7-491f-b404-4f29f8da0ac7','1b500546-e408-47f1-87be-eeca9c758b07','action_0102',1,'2020-11-12 15:48:13',NULL,NULL,NULL,NULL),('eaf0bbd2-85f8-4991-b7a9-f677609bb57f','12b061a2-f10b-4c50-8be5-40b6eba2d321','action_0004',1,'2020-11-09 11:12:15',NULL,NULL,NULL,NULL);
-- /*!40000 ALTER TABLE `user_action` ENABLE KEYS */;
-- UNLOCK TABLES;
-- /*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-25  7:07:15

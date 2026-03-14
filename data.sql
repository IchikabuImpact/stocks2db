-- MySQL dump 10.13  Distrib 8.0.45, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: stocks2db
-- ------------------------------------------------------
-- Server version	8.0.45-0ubuntu0.24.04.1

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
-- Dumping data for table `stock_master`
--

LOCK TABLES `stock_master` WRITE;
/*!40000 ALTER TABLE `stock_master` DISABLE KEYS */;
INSERT INTO `stock_master` (`stock_code`, `company_name`) VALUES ('196A','ＭＦＳ'),('2164','地域新聞社'),('2342','トランスＧＧ'),('2351','ＡＳＪ'),('241A','ＲＯＸＸ'),('281A','インフォメティス'),('288A','ラクサス'),('3042','セキュアヴェイル'),('3137','ファンデリー'),('3187','ミラタップ'),('3237','イントランス'),('3444','菊池製作所'),('3550','スタジオアタオ'),('3624','アクセルマーク'),('3645','メディカルネット'),('3674','オークファン'),('3680','ホットリンク'),('3727','アプリックス'),('3908','コラボス'),('3917','アイリッジ'),('3929','ソーシャルワイヤー'),('3936','グローバルウェイ'),('3987','エコモット'),('4052','フィーチャ'),('4167','ココペリ'),('4170','カイゼン'),('4179','ジーネクスト'),('4240','クラスターテクノロジー'),('4260','ハイブリッド'),('4265','ＩＧＳ'),('4381','ビープラッツ'),('4424','Ａｍａｚｉａ'),('4438','Ｗｅｌｂｙ'),('4484','ランサーズ'),('4490','ビザスク'),('5817','ＪＭＡＣＳ'),('6433','ヒーハイスト'),('7726','黒田精工');
/*!40000 ALTER TABLE `stock_master` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `stock_price_daily`
--

LOCK TABLES `stock_price_daily` WRITE;
/*!40000 ALTER TABLE `stock_price_daily` DISABLE KEYS */;
/*!40000 ALTER TABLE `stock_price_daily` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-03-15  8:27:01

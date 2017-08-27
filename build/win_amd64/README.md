# 動作環境
* OS: Windows
* CPU: 64bit

# 準備

## 簡易的な準備（試しに使う場合）
1. 適当な作業フォルダにtw-toolbox.exeを配置する
2. コマンドプロンプトを起動
3. コマンド「cd パス」で作業フォルダに移動
※使用するたびにステップ2~3の手順が必要

## ちゃんとした準備（今後も使う場合）
1. コマンドプロンプトを起動
2. コマンド「%PATH%」を実行
3. 表示された場所にtw-toolbox.exeを配置
※コマンドプロンプトを起動するだけで使用できる

# 使い方
## ヘルプ
tw-toolbox help

## ツイートの収集
tw-toolbox collect キーワード 取得数
例：tw-toolbox collect 花火 100

### 使用できるオプション（flag）
--show または -s ツイートを全件表示
--export または -e  取得したツイート全件をCSV出力（UTF-8）
--path または -p  出力先を設定（デフォルトはカレントディレクトリ）
--lang または -l  ツイートの言語を指定（言語コードはISO 639-1に準拠）

### 例1：ツイートを取得して全件コマンドプロンプトに表示
tw-toolbox collect 花火 100 -s

### 例2：ツイートを取得して全件CSVに出力
tw-toolbox collect 花火 100 -e

### 例3：ツイートを取得して全件CSVに出力（出力先をdataフォルダに変更）
tw-toolbox collect 花火 100 -e -p /data

### 例4：言語を日本語に指定してCSVに出力
tw-toolbox collect 花火 100 -e -l ja

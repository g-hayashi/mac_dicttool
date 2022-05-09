# name
mac dictionary tool

## Overview

これはmacOSからWindowsへのユーザー辞書の移行を支援するツールです。
 Windowｓからユーザー辞書をエクスポートしてWindowsにインポートするとき、下記の課題があります。

### Google 日本語入力

macOSではUTF-8で出力されるため、UTF-16 (Little Endian with BOM)に変換する必要がある。

### macOS標準 日本語入力

plistで出力されるため、TSVに変換する必要がある。

このプログラムでは、この変換をコマンド一つで実現します。

## Usage

### Google 日本語入力から出力した場合

```
dicttool google <辞書ファイル名>
```

### mac標準のFEPから出力した場合

```
dicttool mac <辞書ファイル名>
```

いずれも同じフォルダに out_<辞書ファイル名>.txt で出力されます。
このファイルをMS-IMEの辞書ツールから読み込んでください。


## Author

[twitter](https://twitter.com/ginjih)

## Licence

[MIT]

go practice

#エニグマとは　　
エニグマとは第二次世界大戦でナチス・ドイツが用いた暗号機である。  
対戦中の1939年にイギリスのアラン・チューリングらによってエニグマの解読に成功した。  

#構造  
大きく分けてプラグボード、ローター、リフレクターの三つのパーツで構成されている。  
それぞれのパーツを組み合わせることで強固な暗号を実現していた。  
  
・プラグボード  
アルファベットのペアを設定する。例えばアルファベットのAとBをペアとして設定した場合、Aと入力した場合Bに暗号化される。また複合化する場合は暗号化されたBを入力するとAに複合化される。  
使用者が任意に設定できる部分であり、エニグマ暗号の鍵の1つである。
  
・ローター  
ローター1つでアルファベット26字の換字表となる。プラグボードから渡された文字を一文字づつ暗号化し、一文字暗号化するたびにローターは1メモリ進む。このローターを複数個繋げて使用することにより膨大なパターンとなる。また一文字変換するたびに回転するメモリ数は変えられる。（コードの中ではoffsetとして表現）  
この一つ一つのローターもエニグマ暗号の鍵の1つである。  
  
・リフレクター  
アルファベットA-Zの中で2文字のペアが設定されているパーツ。例えばCとMがペアになっていた時Cと入力があればMを、Mと入力されればCを返す仕組み。  
この反転性がエニグマの複合機能を実現させている。例えば ABC と入力して DBG と出力された際に逆に DBG と入力すれば ABC と出力されるのは自明である。
  
  
これらのパーツ（オブジェクト）を組み合わせてGo言語でenigmaを再現する
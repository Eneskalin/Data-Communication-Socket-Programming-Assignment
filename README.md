
# Data-Communication-Socket-Programming-Assignment

Bu proje, verilerin bir ağ üzerinden iletilmesi sırasında oluşabilecek hataları simüle etmek ve çeşitli hata tespit algoritmalarını kullanarak bu hataları yakalamak amacıyla geliştirilmiştir.

Sistem; veriyi gönderen bir Client (Sender), gürültülü bir kanalı simüle ederek veriyi bozan bir Server (Channel) ve veriyi alıp doğrulayan bir Client (Receiver) olmak üzere üç ana bileşenden oluşur.


## Desteklenen Hata Tespit Algoritmaları (Error Detection Methods)

Gönderici tarafında seçilen ve alıcı tarafında doğrulanan algoritmalar:

- Parity Check (Tek Boyutlu Eşlik Denetimi)

- 2D Parity Check (İki Boyutlu Eşlik Denetimi)

- CRC-16 (Cyclic Redundancy Check - CCITT)

- Hamming Code (Hata düzeltme ve tespit)

- Internet Checksum (Sağlama Toplamı)


  
## Hata Enjeksiyonu (Error Injection)
Server, veriyi iletirken config.json dosyasındaki ayarlara göre aşağıdaki bozulmaları rastgele uygular:

- Bit Flip: Rastgele bir bitin ters çevrilmesi.

- Character Substitution: Bir karakterin rastgele başka bir karakterle değiştirilmesi.

- Character Insertion: Rastgele bir karakter eklenmesi.

- Character Deletion: Rastgele bir karakterin silinmesi.

- Character Swapping: Yan yana iki karakterin yer değiştirmesi.

- Multiple Bit Flips: Birden fazla bitin ters çevrilmesi.

- Burst Error: Belirli bir bloktaki bitlerin toplu halde bozulması.
## Ekran Görüntüleri

![Uygulama Ekran Görüntüsü](https://i.imgur.com/b2ov0JN.jpeg)

  
## Yükleme 



```bash 
  git clone https://github.com/Eneskalin/Data-Communication-Socket-Programming-Assignment.git
```
    
```bash 
  cd Data-Communication-Socket-Programming-Assignment
```

### Alıcıyı (Receiver) Başlatın
```bash 
    cd clientReceiver
    go run main.go
```

### Sunucuyu (Server) Başlatın
```bash 
    cd server
    go run main.go
```

### Göndericiyi (Sender) Başlatın
```bash 
    cd clientSender
    go run main.go
```


    
## Yapılandırma (Config)

Hata simülasyonlarını `server/config/config.json` dosyasından açıp kapatabilirsiniz. true olan hata türleri, sunucu tarafından rastgele seçilerek uygulanır.



```bash
{
    "ports": {
        "sender_port": "8000",
        "receiver_port": "9000"
    },
    "ErrorInjection": {
        "BitFlip": true,
        "characterSubstitution": true,
        "characterDeletion": true,
        "characterInsertion": true,
        "characterSwapping": false,
        "multipleBitFlips": true,
        "burstError": true,
        "noError": false
    }
}
 ```


  

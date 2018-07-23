# Skycryptor CLI

Skycryptor CLI is a tool which enables you to work with Skycryptor Crypto API functionality. 
SkyCryptor Crypto is an open-source high-level cryptographic library that allows you to perform proxy re-encryption as well other necessary operations for secure storing and transferring data in your decentralized app (Blockchain, IoT, etc...)  

```
$ skycryptor --help

  Skycryptor CLI is a tool which enables you to work with Skycryptor Crypto API functionality.
  
  Usage:
    skycryptor [command]
  
  Available Commands:
    help        Help about any command
    keys        A brief description of your command
  
  Flags:
    -h, --help   help for skycryptor
```
  
```
$ skycryptor keys --help
  A brief description of keys command
  
  Usage:
    skycryptor keys [flags]
    skycryptor keys [command]
  
  Available Commands:
    capsulate   A brief description of key capsulate command
    generate    A brief description of key generate command
  
  Flags:
    -h, --help            help for keys
    -o, --output string   output file
    -p, --pubKey string   publicKey file path
    -s, --secKey string   secretKey file path
```


```
$ skycryptor keys generate --help
  A brief description of key generate command
  
  Usage:
    skycryptor keys generate [flags]
  
  Flags:
    -h, --help           help for generate
        --publicKey      get public key
        --reEncryptKey   generate reEncryption key
  
  Global Flags:
    -o, --output string   output file
    -p, --pubKey string   publicKey file path
    -s, --secKey string   secretKey file path
```


  Examples
  

  ### generate private key
  ```
  $ skycryptor keys generate 
  
    -----BEGIN PRIVATE KEY-----
    6db79c2689f688c3bbdf0ccdcabe6c3887429b9938fca3a0572b9efb127aaaa6
    ------END PRIVATE KEY------
```


  ### get public key from private key
  ```
  $ ./skycryptor keys generate --publicKey -s {YOUR_PRIVATE_KEY_FILE_PATH}
  
    -----BEGIN PUBLIC KEY-----
  303232433746334536463833394146373443354531433742453546333831374641394537323339354146453646343937363536373544423532323346334442423343
    ------END PUBLIC KEY------
 ```
 
  
  ### generate reEncryption key
  ```
  $ skycryptor keys generate --reEncryptKey -s {YOUR_PRIVATE_KEY_FILE_PATH} -p {YOUR_PUBLIC_KEY_FILE_PATH}
  
    -----BEGIN REENCRYPTION KEY-----
    eb897d2bdcb6123b6fdbd6676f6cf22f4a4bf80782c58f69303650665424cbf4
    ------END REENCRYPTION KEY------
 ```
  
  
```
$ skycryptor keys capsulate --help
  brief description of key capsulate command
  
  Usage:
    skycryptor keys capsulate [flags]
  
  Flags:
    -c, --capsule string    capsule path
    -d, --decapsulate       decapsulate key
    -e, --encapsulate       encapsulate key
    -h, --help              help for capsulate
    -r, --reEncKey string   reEncryption key path
        --reEncrypt         reEncrypt key
  
  Global Flags:
    -o, --output string   output file (default "pipe:1")
    -p, --pubKey string   publicKey file path
    -s, --secKey string   secretKey file path
```

  ### Encapsulate
  ```
  $ skycryptor keys capsulate -e -p {YOUR_PUBLIC_KEY_FILE_PATH}
  
    -----BEGIN CAPSULE-----    
   0000004230324531434438433945433641443532444542464246343937344538394646303733303643443639353045383846453339314242363146333541433834413131393430333030373641383145343241444532364435424330333339373641393134423641304338433233393734443435324130413431353433434244354330333032423300000020ebedf6503d9196814ec69cd321613016eeb5094618f2efba371044b46b67fa5b00
    ------END CAPSULE------
    -----BEGIN SYMMETRIC KEY-----
cdcc0fc68d2dde276fa795fa76df237104fe971df9297a9eee96260dd10544c3af9e501baad6d70d55e8a5dcd18ae6233f1f5b9f03258cf3337f6e519d84720e3bdb05674a638b1345f5d14d50e1bc1143a70395391fa31bff1a4e02f2dc7e82f94e73875c9e71c5767cce5d3566a9cac41324aa960ae9609c0a57ed740be88e
    ------END SYMMETRIC KEY------
```    
  
  ### Decapsulate
  ```
  $ ./skycryptor keys capsulate -d -s {YOUR_PRIVATE_KEY_FILE_PATH} -c {YOUR_CAPSULE_FILE_PATH}
  
    -----BEGIN SYMMETRIC KEY-----
cdcc0fc68d2dde276fa795fa76df237104fe971df9297a9eee96260dd10544c3af9e501baad6d70d55e8a5dcd18ae6233f1f5b9f03258cf3337f6e519d84720e3bdb05674a638b1345f5d14d50e1bc1143a70395391fa31bff1a4e02f2dc7e82f94e73875c9e71c5767cce5d3566a9cac41324aa960ae9609c0a57ed740be88e
    ------END SYMMETRIC KEY------
```

### ReEncrypt
```
$ skycryptor keys capsulate --re-encrypt -c {YOUR_CAPSULE_FILE_PATH} --re-enc-key {YOUR_REENCRYPTION_KEY_FILE_PATH} 

    -----BEGIN CAPSULE-----    
   0000004230324531434438433945433641443532444542464246343937344538394646303733303643443639353045383846453339314242363146333541433834413131393430333030373641383145343241444532364435424330333339373641393134423641304338433233393734443435324130413431353433434244354330333032423300000020ebedf6503d9196814ec69cd321613016eeb5094618f2efba371044b46b67fa5b00
    ------END CAPSULE------

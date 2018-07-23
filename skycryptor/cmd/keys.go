package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"skycryptor-sdk-go/skycryptor"
)

const (
	PrivateKeyLen     = 28
	PublicKeyLen      = 27
	ReEncrytionKeyLen = 33
	CapsuleLen        = 24
)

// ioWriter create and return io.Writer object or error
func ioWriter(filename string) (w io.Writer, err error) {
	if filename == "pipe:1" {
		w = os.Stdout
	} else {
		w, err = os.Create(filename)
	}
	return w, err
}

// createMessage return final output message with generated value
func createMessage(key []byte, keyType string) []byte {
	return append(append([]byte("-----BEGIN "+keyType+"-----\n"), key...), []byte("\n------END "+keyType+"------\n")...)
}

// read params
func read(fileName string) []byte {
	var data []byte
	if fileName == "" {
		data, _ = ioutil.ReadAll(os.Stdin)
	} else {
		fmt.Println(fileName)
		data, _ = ioutil.ReadFile(fileName)
	}
	return data
}

// decode represents decoding function for keys
func decode(data []byte, n int) []byte {
	c := data[n:(len(data) - n - 1)]
	d := make([]byte, hex.DecodedLen(len(c)))
	hex.Decode(d, c)
	return d
}

// encode represent encoding function for keys
func encode(data []byte) []byte {
	d := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(d, data)
	return d
}

// getSecretAndPublicKeys read and return decoded private and public keys
func getSecretAndPublicKeys() ([]byte, []byte) {
	if viper.GetString("sec-key") != "" && viper.GetString("pub-key") != "" {
		return decode(read(viper.GetString("sec-key")), PrivateKeyLen), decode(read(viper.GetString("pub-key")), PublicKeyLen)
	} else {
		data := read("")
		return decode(data[0:119], PrivateKeyLen), decode(data[119:len(data)], PublicKeyLen)
	}
}

// getSecretAndCapsule read and return decoded private key and capsule
func getSecretAndCapsule() ([]byte, []byte) {
	if viper.GetString("sec-key") != "" && viper.GetString("capsule") != "" {
		return decode(read(viper.GetString("sec-key")), PrivateKeyLen), decode(read(viper.GetString("capsule")), CapsuleLen)
	} else {
		data := read("")
		return decode(data[0:119], PrivateKeyLen), decode(data[119:len(data)], CapsuleLen)
	}
}

// getReEncKeyAndCapsule read and return decoded reencryption key and capsule
func getReEncKeyAndCapsule() ([]byte, []byte) {
	if viper.GetString("re-enc-key") != "" && viper.GetString("capsule") != "" {
		return decode(read(viper.GetString("re-enc-key")), ReEncrytionKeyLen), decode(read(viper.GetString("capsule")), CapsuleLen)
	} else {
		data := read("")
		return decode(data[0:469], ReEncrytionKeyLen), decode(data[469:len(data)], CapsuleLen)
	}
}

// command variables
var (
	// keysCmd represents the keys command
	keysCmd = &cobra.Command{
		Use:   "keys",
		Short: "A brief description of keys command",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	// generateKeysCmd represents the key generator command
	generateKeysCmd = &cobra.Command{
		Use:   "generate",
		Short: "A brief description of key generate command",
		Run:   manageGenerateKeysCmdFlags,
	}
	// capsulateKeysCmd represents key capsulation command
	capsulateKeysCmd = &cobra.Command{
		Use:   "capsulate",
		Short: "A brief description of key capsulate command",
		Run:   manageCapsulateKeysFlags,
	}
)

// init represent command initialization
func init() {
	rootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(generateKeysCmd, capsulateKeysCmd)

	// keysCmd flags
	keysCmd.PersistentFlags().StringP("sec-key", "s", "", "secretKey file path")
	keysCmd.PersistentFlags().StringP("pub-key", "p", "", "public-key file path")
	keysCmd.PersistentFlags().StringP("output", "o", "pipe:1", "output file")

	viper.BindPFlag("output", keysCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("sec-key", keysCmd.PersistentFlags().Lookup("sec-key"))
	viper.BindPFlag("pub-key", keysCmd.PersistentFlags().Lookup("pub-key"))

	// generate flags
	generateKeysCmd.Flags().BoolP("public-key", "", false, "get public key")
	generateKeysCmd.Flags().BoolP("re-encrypt-key", "", false, "generate reEncryption key")

	viper.BindPFlag("public-key", generateKeysCmd.Flags().Lookup("public-key"))
	viper.BindPFlag("re-ncryptKey", generateKeysCmd.Flags().Lookup("reEncryptKey"))

	// capsulate flags
	capsulateKeysCmd.Flags().BoolP("encapsulate", "e", false, "encapsulate key")
	capsulateKeysCmd.Flags().BoolP("decapsulate", "d", false, "decapsulate key")
	capsulateKeysCmd.Flags().BoolP("re-encrypt", "", false, "reEncrypt key")
	capsulateKeysCmd.Flags().StringP("re-enc-key", "r", "", "reEncryption key path")
	capsulateKeysCmd.Flags().StringP("capsule", "c", "", "capsule path")
	capsulateKeysCmd.Flags().StringP("symmetric-key", "", "", "symmetric key path")

	viper.BindPFlag("encapsulate", capsulateKeysCmd.Flags().Lookup("encapsulate"))
	viper.BindPFlag("decapsulate", capsulateKeysCmd.Flags().Lookup("decapsulate"))
	viper.BindPFlag("re-encrypt", capsulateKeysCmd.Flags().Lookup("re-encrypt"))
	viper.BindPFlag("re-enc-key", capsulateKeysCmd.Flags().Lookup("re-enc-key"))
	viper.BindPFlag("capsule", capsulateKeysCmd.Flags().Lookup("capsule"))
	viper.BindPFlag("symmetric-key", capsulateKeysCmd.Flags().Lookup("symmetric-key"))
}

// manageGenerateKeysCmdflags represents generateKeys command flags management
func manageGenerateKeysCmdFlags(cmd *cobra.Command, args []string) {
	sc := skycryptor.NewSkycryptor()
	defer sc.Clean()

	if viper.GetBool("re-encrypt-key") {
		w, _ := ioWriter(viper.GetString("output"))
		generateReEncryptionKey(sc, w)
		//w.Write(generateReEncryptionKey(sc))
	} else if viper.GetBool("public-key") {
		w, _ := ioWriter(viper.GetString("output"))
		getPublicKey(sc, w)
	} else {
		w, _ := ioWriter(viper.GetString("output"))
		generatePrivateKey(sc, w)
	}
}

// manageCapsulateKeysCmdflags represents capsulateKeys command flags management
func manageCapsulateKeysFlags(cmd *cobra.Command, args []string) {
	sc := skycryptor.NewSkycryptor()
	defer sc.Clean()

	if viper.GetBool("encapsulate") {
		w1, _ := ioWriter(viper.GetString("output"))
		w2, _ := ioWriter(viper.GetString("symmetric-key"))
		encapsulate(sc, w1, w2)
	} else if viper.GetBool("decapsulate") {
		w, _ := ioWriter(viper.GetString("output"))
		decapsulate(sc, w)
	} else if viper.GetBool("re-encrypt") {
		w, _ := ioWriter(viper.GetString("output"))
		reEncrypt(sc, w)
	} else {
		fmt.Println("\nMissing subcommand...\n")
	}
}

// Generate public and private keys
func generatePrivateKey(sc *skycryptor.SkyCryptor, w io.Writer) {
	privateKey, publickey := sc.Keys.Generate()
	defer privateKey.Clean()
	defer publickey.Clean()

	w.Write(createMessage(encode(privateKey.ToBytes()), "PRIVATE KEY"))
}

// Getting public key from our private key
func getPublicKey(sc *skycryptor.SkyCryptor, w io.Writer) {
	secKey := decode(read(viper.GetString("sec-key")), PrivateKeyLen)
	sk := sc.PrivateKeyFromBytes(secKey)
	defer sk.Clean()

	pk := sk.GetPublicKey()
	defer pk.Clean()

	w.Write(createMessage(encode(pk.ToBytes()), "PUBLIC KEY"))
}

// Getting ReEncryption key using current PrivateKey and given PublicKey
func generateReEncryptionKey(sc *skycryptor.SkyCryptor, w io.Writer) {
	secKey, pubKey := getSecretAndPublicKeys()
	sk := sc.PrivateKeyFromBytes(secKey)
	defer sk.Clean()
	pk := sc.PublicKeyFromBytes(pubKey)
	defer pk.Clean()

	reEncryptionKey := sk.GenerateReKey(pk)
	defer reEncryptionKey.Clean()

	w.Write(createMessage(encode(reEncryptionKey.ToBytes()), "REENCRYPTION KEY"))
}

// Making encapsulation and getting Capsule with symmetric key
func encapsulate(sc *skycryptor.SkyCryptor, w1 io.Writer, w2 io.Writer) {
	pubKey := decode(read(viper.GetString("pub-key")), PublicKeyLen)
	pk := sc.PublicKeyFromBytes(pubKey)
	defer pk.Clean()

	c, k := pk.Encapsulate()
	w1.Write(createMessage(encode(c.ToBytes()), "CAPSULE"))
	w2.Write(createMessage(encode(k), "SYMMETRIC KEY"))
}

// Decapsulating given capsule and getting back symmetric key
func decapsulate(sc *skycryptor.SkyCryptor, w io.Writer) {
	secKey, capsule := getSecretAndCapsule()
	sk := sc.PrivateKeyFromBytes(secKey)
	defer sk.Clean()

	c := sc.CapsuleFromBytes(capsule)
	defer c.Clean()

	k := sk.Decapsulate(c)
	w.Write(createMessage(encode(k), "SYMMETRIC KEY"))
}

// Running re-encryption for given capsule and returning re-encrypted capsule
func reEncrypt(sc *skycryptor.SkyCryptor, w io.Writer) {
	reEncKey, capsule := getReEncKeyAndCapsule()
	rk := sc.ReEncryptionKeyFromBytes(reEncKey)
	defer rk.Clean()

	c := sc.CapsuleFromBytes(capsule)
	defer c.Clean()

	k := rk.ReEncrypt(c)
	defer k.Clean()

	w.Write(createMessage(encode(k.ToBytes()), "CAPSULE"))
}

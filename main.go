package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type counter struct {
	dirs    int
	secrets int
	name    string
}

func (counter *counter) index(path string) {
	if strings.HasSuffix(path, "/") {
		counter.dirs++
	} else {
		counter.secrets++
	}
}

func (counter *counter) output() string {
	return fmt.Sprintf("\n%d paths, %d %s", counter.dirs, counter.secrets, counter.name)
}

func dirnamesFrom(base string, logical *vault.Logical) []string {
	if !strings.HasSuffix(base, "/") {
		return nil
	}

	secret, err := logical.List(base)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	var keys []string
	keyObjects := secret.Data["keys"].([]interface{})
	for _, v := range keyObjects {
		elem := fmt.Sprint(v)
		keys = append(keys, elem)
	}

	sort.Strings(keys)
	return keys
}

func tree(counter *counter, base string, prefix string, logical *vault.Logical) {
	names := dirnamesFrom(base, logical)

	for index, name := range names {
		subpath := base + name
		counter.index(subpath)

		if index == len(names)-1 {
			fmt.Println(prefix+"└──", strings.TrimSuffix(name, "/"))
			tree(counter, subpath, prefix+"    ", logical)
		} else {
			fmt.Println(prefix+"├──", strings.TrimSuffix(name, "/"))
			tree(counter, subpath, prefix+"│   ", logical)
		}
	}
}

func main() {

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatalln("err: VAULT_TOKEN not defined")
	}

	vaultServerEndpoint := os.Getenv("VAULT_ADDR")
	if vaultServerEndpoint == "" {
		log.Fatalln("err: VAULT_ADDR not defined")
	}

	subcommand := flag.String("subcommand", "kv", "Specify the object to query. Options are: kv, policy and kroles. kv by default")
	kubernetesCluster := flag.String("kubernetes", "kubernetes", "Specify the kubernetes integration name for kroles. kubernetes by default")
	flag.Parse()

	config := vault.DefaultConfig()
	config.Address = vaultServerEndpoint

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	var directory string
	var elementName string
	logical := client.Logical()
	switch *subcommand {
	case "policy":
		directory = "sys/policy/"
		elementName = "policies"
	case "kroles":
		directory = fmt.Sprintf("auth/%s/role/", *kubernetesCluster)
		elementName = "kroles"
	default:
		directory = "kv/metadata/"
		elementName = "secrets"
	}

	counter := &counter{
		name: elementName,
	}
	fmt.Println(strings.TrimSuffix(directory, "/"))

	tree(counter, directory, "", logical)
	fmt.Println(counter.output())
}

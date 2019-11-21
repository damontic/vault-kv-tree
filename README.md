# vault-kv-tree
Shows Hashicorp Vault kv secrets (v2) as a tree

## Install

```bash
$ go get github.com/damontic/vault-kv-tree
$ go install github.com/damontic/vault-kv-tree
```

## Usage

### List kv secrets

1. Export both VAULT_ADDR and VAULT_TOKEN

2. Execute

   ```bash
   $ vault-kv-tree
   kv/metadata
   ├── a_secret
   ├── a_secret_2
   │   ├── prod
   │   └── qa
   ├── a_secret_3
   │   ├── prod
   │   └── qa
   ├── a_secret_4
   │   └── something
   ├── qa
   │   ├── test
   │   └── something_else
   └── tools
       └── anchore
   
   6 paths, 9 secrets
   ```

### List policies

1. Export both VAULT_ADDR and VAULT_TOKEN

2. Execute

   ```bash
   $ vault-kv-tree -subcommand policy
   sys/policy
   ├── default
   └── root
   
   0 paths, 2 policies
   ```

### List kubernetes authentication roles

1. Export both VAULT_ADDR and VAULT_TOKEN

2. Execute

   ```bash
   $ vault-kv-tree -subcommand kroles -kubernetes kubernetes-prod
   auth/kubernetes-prod/role
   ├── another-k-role
   └── secrets-read-role
   
   0 paths, 2 kroles
   ```


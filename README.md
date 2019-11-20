# vault-kv-tree
Shows Hashicorp Vault kv secrets (v2) as a tree

## Usage

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


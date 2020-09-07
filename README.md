# github-app-installer

# Usage

```shell
$ export GHE_ACCESS_TOKEN=...
$ export GHE_BASE_URL=https://git.example.com/

$ github-app-installer list user
$ github-app-installer install --installation-id <id> user/repo

# Dry-Run
$ github-app-installer auto-merge --dry-run --branch-name 'renovate/configure' --title 'Configure Renovate' --retry-count 10 user/repo
$ github-app-installer auto-merge --branch-name 'renovate/configure' --title 'Configure Renovate' --retry-count 10 user/repo
```

If you do not know the installation id, it is the number at the end of the URL of the setting screen of the GitHub App you want to install.  
For example, `https://git.example.com/organizations/test/settings/installations/41` is `41` .

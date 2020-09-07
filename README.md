# github-app-installer

A tool that automates the installation of the GitHub App

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

## List

List repositories for the specified organization.

```
$ github-app-installer list <organizations>

# Show only default branch changes since 2020-03-17
$ github-app-installer list --after 2020-03-07 <organizations>
```

## Install

Install the GitHub App in the specified repository. This is useful if you grep the repository from the `list` command and pass it with xargs.  
If you do not know the installation id, it is the number at the end of the URL of the setting screen of the GitHub App you want to install.  
For example, `https://git.example.com/organizations/test/settings/installations/41` is `41` .

```
$ github-app-installer install --installation-id <id> <organization/repository>
```

## Auto Merge

There is also a GitHub App that create the files as a PR. This command is for automatically merging that PR.

For example, renovate creates PR according to the following rules:

- PR title : Configure Renovate
- Branch Name : renovate/configure

```shell
# Dry-Run
$ github-app-installer auto-merge --dry-run --branch-name 'renovate/configure' --title 'Configure Renovate' --retry-count 20 --dry-run <organization/repository>

# Merge
$ github-app-installer auto-merge --branch-name 'renovate/configure' --title 'Configure Renovate' --retry-count 20 --dry-run <organization/repository>
```

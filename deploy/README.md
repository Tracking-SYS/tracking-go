# DEPLOYMENT SETUP (TRAVIS CI)

## Installation

1. Install Travis CLI
```
gem install travis
```

2. Create API Token on Github

_Under the GitHub account settings for the user you want to use, navigate to **Settings** > **Developer settings**, and then generate a **Personal access tokens**. Make sure the token has the **repo** scope._

3. Login Travis CLI
```  
travis login --pro --com --github-token <github personal access token>
```

4. Add token to your ```.travis.yml```
```
travis encrypt CI_USER_TOKEN="<github personal access token>" --pro --com --add
```


## Fractal 
go version for fractal blockchain

master | dev
-------|----------
[![TravisCI](https://travis-ci.org/fractal-platform/fractal.svg?branch=master)](https://travis-ci.org/fractal-platform/fractal) | [![TravisCI](https://travis-ci.org/fractal-platform/fractal.svg?branch=dev)](https://travis-ci.org/fractal-platform/fractal)

## Build Steps
1. make your own go env, set $GOPATH, and add $GOPATH/bin to PATH env
    ```
    cd ~
    mkdir -p go/bin go/src
    export GOPATH=~/go
    export PATH=$PATH:$GOPATH/bin
    ```
    
2. clone source (The default branch is 0.2.x)
    ```
    cd $GOPATH/src
    git clone https://github.com/GuoxiW/fractal github.com/GuoxiW/fractal
    ```
    
3. build
    ```
    cd $GOPATH/src/github.com/GuoxiW/fractal
    sudo cp transaction/txexec/libwasmlib.dylib /usr/local/lib/
    go install -v -ldflags "-X main.gitCommit=$(git log --pretty=format:'%h' -1)" ./cmd/gftl/  
    go install -v -ldflags "-X main.gitCommit=$(git log --pretty=format:'%h' -1)" ./cmd/gtool/
    ```

## Running Fractal blockchain
### Setup Basic Environment and Start mining
1. Create a work directory
    ```
    cd ~
    mkdir test
    cd test
    ```

2. Get the private key from the wallet, import it to the 'data/keys' folder and set the password
    ```
    gtool keys --pri <private key> --keys data/keys --pass <password> import
    ```

3. Generate the mining key, the address can be copied from the wallet, the password is set in the previous step
    ```
    gtool keys --keys data/keys --pass <password> --addr <address> newminingkey
    ```

4. Register for mining
    ```
    gtool keys --keys data/keys --pass <password> --chainid 1 --rpc http://3.226.27.211:8545 regminingkey
    ```

5. Start mining
    ```
    gftl --datadir ./data --port 60006 --verbosity 3 --mine --unlock <password>
    ```

    Note: If you don’t want to mine by yourself, we can help you. First, You need to copy your mining key to us. The mining keys are stored in the 'data/keys/mining_keys' folder, you can package the 'mining_keys' folder and send it to us. At the same time, you also need to inform us of your private key password. Note that the 'data/keys/account.json' file holds your account key. Never send this file to us.

6. publish data, the address can be copied from the wallet
    ```
    gtool tx --rpc http://35.170.127.58:8545 --to <address> --chainid 1 --keys data/keys --pass <password> --data <your data> send
    ```



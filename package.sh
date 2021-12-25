# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

# cross_compiles
make -f ./Makefile

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        nsq_auth_dir_name="nsq_auth_${os}_${arch}"
        nsq_auth_path="./packages/nsq_auth_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./nsq_auth_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${nsq_auth_path}
            mv ./nsq_auth_${os}_${arch}.exe ${nsq_auth_path}/nsq_auth.exe
        else
            if [ ! -f "./nsq_auth_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${nsq_auth_path}
            mv ./nsq_auth_${os}_${arch} ${nsq_auth_path}/nsq_auth
        fi
        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${nsq_auth_dir_name}.zip ${nsq_auth_dir_name}
        else
            tar -zcf ${nsq_auth_dir_name}.tar.gz ${nsq_auth_dir_name}
        fi
        cd ..
        rm -rf ${nsq_auth_path}
    done
done

cd -

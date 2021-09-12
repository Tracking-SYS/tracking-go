 #!/bin/bash

set -ex

WORK_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

clone_repo() {
  DST=/tmp/genproto
  rm -rf "${DST}"
  git clone https://github.com/tikivn/genproto ${DST}
  cd ${DST}
  git config user.name "hienduyph"
  git config user.email "hien.pham2@tiki.vn"
  echo ${DST}
}

commit_n_push() {
  if [[ -z $(git status -s) ]]
  then
    echo "Tree is clean. Exit!"
    exit
  fi

  echo "Start Commit"
  git add .
  git commit -m "Gen artifacts"
  git push ${GIT_PUSH_ARGS} origin main
}

gen_py_setup() {
  DST=$1
  cat ${WORK_DIR}/setup.py | envsubst > ${DST}/setup.py
  cp "${WORK_DIR}/requirements.txt" "${DST}"
}

push_tag() {
  VERSION="$1"
  git tag "${VERSION}"
  git push ${GIT_PUSH_ARGS}  origin "${VERSION}"
}

main() {
  CURR_DIR="$1"
  DST="$(clone_repo)"
  cd ${DST}
  echo "WORK DIR ${CURR_DIR}"
  echo "Copy go artifacts"

  rm -rf ${DST}/go ${DST}/python

  # copy go
  cp -r ${CURR_DIR}/_go ${DST}/go
  cp -r ${CURR_DIR}/guides ${DST}
  cp ${CURR_DIR}/README.md ${DST}/README.md

  # copy python
  echo "Copy python artifacts"
  cp -r ${CURR_DIR}/python ${DST}/python
  gen_py_setup "${DST}" "${VERSION}"

  cd ${DST}

  go mod tidy

  commit_n_push

  VERSION="$2"
  if [[ ! -z $VERSION ]]; then
    echo "Creating tag name ${VERSION}"
    push_tag "${VERSION}"
  fi
}

main "$@"
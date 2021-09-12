#!/usr/bin/env bash
# this script helps to if we have connection to some services of tiki
#tested in macos + linux
#how to run: ./vpn-check.sh list_hosts.txt

timeout=1 #seconds

if [ -t 1 ]; then
  # see if it supports colors...
  ncolors=$(tput colors)
  if test -n "$ncolors" && test $ncolors -ge 8; then
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    NC='\033[0m' # No Color
  fi
fi

function process_line() {
  host=$1
  port=$2
  if [ -z $1 ]; then
    return
  else
    if [ "$(uname)" == "Darwin" ]; then
      nc -z -G "$timeout" $host $port &>/dev/null
    elif [ "$(expr substr "$(uname)" 1 5)" == "Linux" ]; then
      nc -z -w "$timeout" $host $port &>/dev/null
    else
      echo "Not in macos or linux"
      exit 1
    fi

    if [ "$?" -eq 0 ]; then
      printf "${GREEN}Online${NC}  $host:$port \n"
    else
      printf "${RED}Offline${NC} $host:$port \n"
    fi
  fi
}

function process_file() {
  input_file=$1
  cat "$input_file" | while read line; do
    # remove character after #
    line=${line%#*}
    # remove leading whitespace characters
    line="${line#"${line%%[![:space:]]*}"}"
    # remove trailing whitespace characters
    line="${line%"${line##*[![:space:]]}"}"
    line=$(echo "$line" | awk -F':' '{if (NF!=1) {print $1 " " $2} else {print $1 " " 80}}')
    if [ -z "$line" ]; then
      echo " empty"
    else
      process_line $line
    fi
  done
}

process_file $1
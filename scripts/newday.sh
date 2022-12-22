#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Set up defaults
base_dir=~/code/projects/advent/
year="2022"

function usage() {
    echo "Usage: newday [options]"
    echo " Option: "
    echo "  [-d|--day]   - the number of the new day to set up."
    echo "  [-y|--year]  - the year to set the new day in. default: "${year}
    echo "  [-h|--help]      - display this help message"
    exit 0;
}

if [[ $# -eq 0 ]]; then
    usage
fi

# override defaults with command line args
while [[ $# -gt 1 ]]
do
key="$1"

case ${key} in
    -d|--day)
    day="$2"
    if [[ ${#day} < 2 ]]; then
        day="00${day}"
        day="${day: -2}"
    fi
    shift
    ;;
    -y|--year)
    year="$2"
    shift
    ;;
    -h|--help)
    usage
    ;;
    *)
            # unknown option
    ;;
esac
shift
done

echo ${year}"/day"${day}
dir=${base_dir}${year}"/day"${day}

if [[ -d "$dir" ]]; then
  echo "day already exists!"
  exit 0;
else
  mkdir -p "$dir"
  cp -R templates/day/ "$dir"
  echo "$dir created"
  touch "$dir/input.txt"
  touch "$dir/testinput.txt"
  dayN=$(echo $day | sed 's/^0*//')
  open -a 'Google Chrome' "https://adventofcode.com/${year}/day/${dayN}/input"
  echo "input files created"
fi

cd "$dir"

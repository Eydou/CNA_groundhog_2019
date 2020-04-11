#!/usr/bin/env bash
##
## EPITECH PROJECT, 2020
## CNA_groundhog_2019
## File description:
## graphic
##

graphic() {
    len=$(wc -l < .data.csv)
    lenght=$(echo $(($len)))
    touch .dat
    for i in `seq 1 $lenght`;
    do
        echo $i >> .dat
    done
    pr -mts' ' .dat .data.csv > data.csv
    rm .dat
    rm .data.csv
    
}

editFile() {
    sed -i 's/\(.\{25\}\).*/\1/' .data.csv
    sed -i '/n/d' .data.csv
    sed -i '/n/d' .data.csv
    sed -i '$d' .data.csv
    sed -i '$d' .data.csv
    sed -i 's/w//g' .data.csv
    sed -i 's/a//g' .data.csv
    sed -i 's/g//g' .data.csv
    sed -i 's/r//g' .data.csv
    sed -i 's/s//g' .data.csv
    sed -i 's/=//g' .data.csv
    sed -i 's/i//g' .data.csv
    sed -i 's/t//g' .data.csv
}

help() {
    echo "Usage:
        ./graphic [-h] [[-g or -a] file period]
Description:
        -h      print the help
        file    put A file with your temperature values
        period  the number of days defining a period
        -g       option for the general print
        -a       option for the analyst print"
}

main() {
    if [ $# -lt 1 ]; then
        help
        exit 0
    fi
    if [ $1 = "--help" ] || [ $1 = "-h" ]; then
        help
        exit 0
    fi
    if [ "$#" -ne 3 ]; then
    echo "error need file or period or opt"
    exit 84
    fi
    if [ $# -eq 3 ]; then
        cd .. && make re &>/dev/null && cd bonus
        cat $2 | .././groundhog $3 > .data.csv
        status=$(echo $?)
        editFile
        echo -e "\e[92mOk\033[0m Edit the file"
        graphic
        echo -e "\e[92mOk\033[0m Convert in png"
        if [ $1 == "-g" ]; then
            cat drawer.gnu | gnuplot
        elif [ $1 == "-a" ]; then
            cat drawer1.gnu | gnuplot
        else
            echo "invalid opt"
        fi
        rm data.csv
        echo -e "\e[92mCompleted !\033[0m"
        exit $status
    fi
}

main "$@"
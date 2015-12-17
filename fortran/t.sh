#!/bin/sh

file='./case'

for ff in $file
do
    echo $file 
    while read LINE
    do
        echo $LINE
    done
done


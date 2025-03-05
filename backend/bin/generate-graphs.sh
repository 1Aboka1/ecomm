#!/usr/bin/env bash 

for dir in graphqls/*/
do
  dir=${dir%*/}
  echo "generating for ${dir}..."
  cd $dir
  go run github.com/99designs/gqlgen generate
  cd ../../
done
echo "done"

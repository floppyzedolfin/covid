#!/usr/bin/env bash

function comparejson {
  reference=$1
  obtained=$2

  # from https://stackoverflow.com/questions/31930041/using-jq-or-alternative-command-line-tools-to-compare-json-files
  res=$(cmp <(jq -cS . $reference) <(jq -cS . $obtained))
  if [[ ! -z "${res}" ]]; then
    echo unidentical files $reference and $obtained
    exit 1
  fi
}

function cleanup {
  kill -9 $(ps | awk '/covid.out/{print $1}')
}

# run the binary
make run &

# wait a few seconds for the binary to be up and running
while [[ true ]]; do
  pid=$(ps | awk '/covid.out/{print $1}')
  if [[ ! -z "${pid}" ]]; then
    break
  fi
done
trap cleanup EXIT

# load the database
curl localhost:8405/admin/load
if [[ "$?" != 0 ]]; then
  echo "error while loading database"
  exit 1
fi

# test api endpoints
curl localhost:8405/api/departements -o testdata/obtained_departements.json
comparejson testdata/result_departements.json testdata/obtained_departements.json

curl localhost:8405/api/departements/06/dates/2020-12-23 -o testdata/obtained_departement_date.json
comparejson testdata/result_departement_date.json testdata/obtained_departement_date.json

curl localhost:8405/api/departements/73,74/dates/2020-11-01/2020-11-04 -o testdata/obtained_departements_dates.json
comparejson testdata/result_departements_dates.json testdata/obtained_departements_dates.json

curl localhost:8405/api/rangeDates -o testdata/obtained_range_dates.json
comparejson testdata/result_range_dates.json testdata/obtained_range_dates.json

curl localhost:8405/api/classAges -o testdata/obtained_classages.json
comparejson testdata/result_classages.json testdata/obtained_classages.json


# test analytics endpoints
curl localhost:8405/analytics/national/dates/2020-07-08/2020-07-12 -o testdata/obtained_national.json
comparejson testdata/result_national.json testdata/obtained_national.json

curl localhost:8405/analytics/departements/33/brief -o testdata/obtained_brief.json
comparejson testdata/result_brief.json testdata/obtained_brief.json

curl localhost:8405/analytics/dates/2020-10-27/top5 -o testdata/obtained_top5.json
comparejson testdata/result_top5.json testdata/obtained_top5.json

exit 0

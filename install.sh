#!/bin/bash
function create_config {
  DIR=$1
  cat >./service.yaml <<EOF
product:
  source: ${DIR}/data/source/products/
  report: ${DIR}/data/result/report/
  sent: ${DIR}/data/source/processed/products/
  in_progress: ${DIR}/data/source/inprogress/
  success_result: ${DIR}/data/result/sent/
  fail_result: ${DIR}/data/result/report/
  mapping: ${DIR}/data/mapping/mapping.yaml
  ontology: ${DIR}/data/ontology/rules.csv
offer:
  source: ${DIR}/data/source/offers/
  sent: ${DIR}/data/source/processed/offers/
common:
  source: ${DIR}/data/source/
  sent: ${DIR}/data/source/processed/
  sheet:
    products: "Products"
    offers: "Offers"
    failures: "Attributes"
tradeshift_api:
  # set Tradeshift API parameters from API Access To Own Account in Tradeshift pannel
  base_url:
  consumer_key:
  consumer_secret:
  token:
  token_secret:
  tenant_id:
EOF
}

while getopts d: OPT; do
  case "$OPT" in
  d)
    DIR="$OPTARG"
    ;;
  [?])
    # got invalid option
    echo "Usage: $0 [-d work directory]" >&2
    exit 1
    ;;
  esac
done

go test ./...
go get ./...

mkdir -p $DIR
go build -o ./product-catalog-import-tool

cp ./product-catalog-import-tool $DIR/product-catalog-import-tool
create_config $DIR
cp ./service.yaml $DIR/
cp -R ./data $DIR/

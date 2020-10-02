#!/bin/sh

# Fictional secret
api_key=$(secrethub read company/demo/api_key)
echo $api_key

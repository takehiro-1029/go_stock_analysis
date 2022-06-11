#!/bin/bash
base64 task-definition.json > task-definition.txt

# 確認
# base64 -d task-definition.txt
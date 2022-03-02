#!/opt/homebrew/bin/python3

import argparse

parser = argparse.ArgumentParser(description='Find AWS endpoints info from SDK source codes and generates a csv')
parser.add_argument('--aws-sdk-go', required=True, help='path to aws-sdk-go source repo')
parser.add_argument('--aws-sdk-go-v2', required=True, help='path to aws-sdk-go-v2 source repo')
parser.add_argument('--boto3', required=True, help='path to boto3 source repo')

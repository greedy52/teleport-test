import boto3
import json
from botocore.config import Config

proxy_definitions = {
        'http': 'https://127.0.0.1:34567',
        'https': 'https://127.0.0.1:34567',
}

my_config = Config(
    region_name='ca-central-1',
    signature_version='v4',
    proxies=proxy_definitions,
    proxies_config={
        'proxy_ca_bundle': '/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-app/teleport.dev.aws.stevexin.me/awsconsole-localhost-ca.pem',
    },
)

sts = boto3.client('sts', config=my_config)
print(json.dumps(sts.get_caller_identity(), indent=2))

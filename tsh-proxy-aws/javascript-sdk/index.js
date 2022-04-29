var ProxyAgent = require('proxy-agent');
var AWS = require('aws-sdk');

if (process.env.https_proxy) {
  AWS.config.update({
    httpOptions: { agent: ProxyAgent(process.env.https_proxy) }
  });
};

var sts = new AWS.STS();
var params = {};
sts.getCallerIdentity(params, function(err, data) {
  if (err) console.log(err, err.stack); // an error occurred
  else     console.log(data);           // successful response
});

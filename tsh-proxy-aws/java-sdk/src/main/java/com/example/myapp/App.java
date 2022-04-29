package com.example.myapp;

import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.http.SdkHttpClient;
import software.amazon.awssdk.http.apache.Apache:ttpClient;
import software.amazon.awssdk.http.apache.ProxyConfiguration;
import software.amazon.awssdk.services.sts.StsClient;
import software.amazon.awssdk.services.sts.model.GetCallerIdentityResponse;
import software.amazon.awssdk.services.sts.model.StsException;

public class App {
    public static void main(String[] args) {
       final SdkHttpClient httpClient = ApacheHttpClient.builder()
            .proxyConfiguration(ProxyConfiguration.builder()
                    .useSystemPropertyValues(true)
                    .build())
            .build();
        Region region = Region.US_EAST_1;
        StsClient stsClient = StsClient.builder()
                .region(region)
                .httpClient(httpClient)
                .build();

        getCallerId(stsClient);
        stsClient.close();
    }

    public static void getCallerId(StsClient stsClient) {

        try {
            GetCallerIdentityResponse response = stsClient.getCallerIdentity();

            System.out.println("The user id is" +response.userId());
            System.out.println("The ARN value is" +response.arn());

        } catch (StsException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}

#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { CdkTodosStack } from '../lib/cdk-todos-stack';

// Initialize CDK App
const app = new cdk.App();

// Instantiate the Stack
new CdkTodosStack(app, 'CdkTodosStack', {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
});

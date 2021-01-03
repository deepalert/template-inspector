#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { InspectorStack } from '../lib/stack';

const app = new cdk.App();
new InspectorStack(app, 'TestStack', {
  taskTopicARN: 'arn:aws:sns:us-east-2:123456789012:MyTopic',
  findingQueueARN: 'arn:aws:sqs:us-east-2:444455556666:queue1',
  attributeQueueARN: 'arn:aws:sqs:us-east-2:444455556666:queue2',
});

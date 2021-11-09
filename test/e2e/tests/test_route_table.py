# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the RouteTable API.
"""

import pytest
import time
import logging

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_ec2_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources

RESOURCE_PLURAL = "routetables"

CREATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

def get_route_table(ec2_client, route_table_id: str) -> dict:
    try:
        resp = ec2_client.describe_route_tables(
            Filters=[{"Name": "route-table-id", "Values": [route_table_id]}]
        )
    except Exception as e:
        logging.debug(e)
        return None

    if len(resp["RouteTables"]) == 0:
        return None
    return resp["RouteTables"][0]


def route_table_exists(ec2_client, route_table_id: str) -> bool:
    return get_route_table(ec2_client, route_table_id) is not None

@service_marker
@pytest.mark.canary
class TestRouteTable:
    def test_create_delete(self, ec2_client):
        test_resource_values = REPLACEMENT_VALUES.copy()
        resource_name = random_suffix_name("route-table-test", 24)
        test_vpc = get_bootstrap_resources().SharedTestVPC
        vpc_id = test_vpc.vpc_id

        test_resource_values["ROUTE_TABLE_NAME"] = resource_name
        test_resource_values["VPC_ID"] = vpc_id

        # Load Route Table CR
        resource_data = load_ec2_resource(
            "route_table",
            additional_replacements=test_resource_values,
        )
        logging.debug(resource_data)

        # Create k8s resource
        ref = k8s.CustomResourceReference(
            CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
            resource_name, namespace="default",
        )
        k8s.create_custom_resource(ref, resource_data)
        cr = k8s.wait_resource_consumed_by_controller(ref)

        assert cr is not None
        assert k8s.get_resource_exists(ref)

        resource = k8s.get_resource(ref)
        resource_id = resource["status"]["routeTableID"]

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        # Check Route Table exists
        exists = route_table_exists(ec2_client, resource_id)
        assert exists

        # Delete k8s resource
        _, deleted = k8s.delete_custom_resource(ref)
        assert deleted is True

        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        # Check Route Table doesn't exist
        exists = route_table_exists(ec2_client, resource_id)
        assert not exists

    def test_terminal_condition(self):
        test_resource_values = REPLACEMENT_VALUES.copy()
        resource_name = random_suffix_name("route-table-fail", 24)
        test_resource_values["ROUTE_TABLE_NAME"] = resource_name
        test_resource_values["VPC_ID"] = "InvalidVpcId"

        # Load RouteTable CR
        resource_data = load_ec2_resource(
            "route_table",
            additional_replacements=test_resource_values,
        )
        logging.debug(resource_data)

        # Create k8s resource
        ref = k8s.CustomResourceReference(
            CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
            resource_name, namespace="default",
        )
        k8s.create_custom_resource(ref, resource_data)
        cr = k8s.wait_resource_consumed_by_controller(ref)

        assert cr is not None
        assert k8s.get_resource_exists(ref)

        expected_msg = "InvalidVpcID.NotFound: The vpc ID 'InvalidVpcId' does not exist"
        terminal_condition = k8s.get_resource_condition(ref, "ACK.Terminal")
        # Example condition message:
        #   InvalidVpcID.NotFound: The vpc ID 'InvalidVpcId' does not exist
        #   status code: 400, request id: 5801fc80-67cf-465f-8b83-5e02d517d554
        # This check only verifies the error message; the request hash is irrelevant and therefore can be ignored.
        assert expected_msg in terminal_condition['message']
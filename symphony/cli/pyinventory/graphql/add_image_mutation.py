#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .add_image_input import AddImageInput


@dataclass
class AddImageMutation(DataClassJsonMixin):
    @dataclass
    class AddImageMutationData(DataClassJsonMixin):
        @dataclass
        class File(DataClassJsonMixin):
            id: str
            fileName: str

        addImage: File

    data: AddImageMutationData

    __QUERY__: str = """
    mutation AddImageMutation($input: AddImageInput!) {
  addImage(input: $input) {
    id
    fileName
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddImageInput) -> AddImageMutationData:
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data

#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from dataclasses import dataclass
from datetime import datetime
from functools import partial
from gql.gql.datetime_utils import DATETIME_FIELD
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .property_input import PropertyInput
@dataclass
class AddLocationInput(DataClassJsonMixin):
    name: str
    type: str
    properties: List[PropertyInput]
    parent: Optional[str] = None
    latitude: Optional[Number] = None
    longitude: Optional[Number] = None
    externalID: Optional[str] = None


#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from enum import Enum

class PropertyKind(Enum):
    string = "string"
    int = "int"
    bool = "bool"
    float = "float"
    date = "date"
    enum = "enum"
    range = "range"
    email = "email"
    gps_location = "gps_location"
    equipment = "equipment"
    location = "location"
    service = "service"
    datetime_local = "datetime_local"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value: str) -> "PropertyKind":
        return cls.MISSING_ENUM

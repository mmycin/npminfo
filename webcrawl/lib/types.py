from typing_extensions import TypeAlias
from typing import Final, List, Dict
import requests

Response: TypeAlias = Final[requests.Response]
Data: TypeAlias = List[Dict[str, List[str]]]
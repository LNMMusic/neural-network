import numpy as np
from numpy.typing import NDArray
from typing import List

# PatternFinderArray2D
class PatternFinderArray2D:
    schema:NDArray[np.int64]
    threshold:int

    # constructor
    def __init__(self, schema:NDArray[np.int64], threshold:int):
        # set schema
        self.schema = schema
        # set threshold
        self.threshold = threshold

    # Find allows to find a pattern in the schema
    def Find(self, data:List[NDArray[np.int64]]) -> NDArray[np.int64]:
        # counter
        patternArray = np.zeros(self.schema.shape)
    
        # iterate over data
        counterArray = np.zeros(self.schema.shape)
        for a in data:
            # iterate over array
            for (i, j), v in np.ndenumerate(a):
                if a[i,j] == 1:
                    counterArray[i,j] += 1

        # check in counterArray if it reaches the threshold value
        for (i, j), v in np.ndenumerate(counterArray):
            if counterArray[i,j] >= self.threshold:
                patternArray[i,j] = 1

        # return patternArray
        return patternArray
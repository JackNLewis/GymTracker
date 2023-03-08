'''
This is a python scrapper that retives the list of exercises and their cooresponding muscler group from 
https://www.strengthlog.com/exercise-directory/
This is used to populate the Exercise table in the database.
'''

import requests
from bs4 import BeautifulSoup
import re
import pandas as pd
import numpy as np
import os

#Initialize bs4 and make page request
url = "https://www.strengthlog.com/exercise-directory/"
page = requests.get(url)

soup = BeautifulSoup(page.content, "html.parser")

header_names = [
    'h-chest-exercises',
    'h-shoulder-exercises',
    'h-bicep-exercises',
    'h-triceps-exercises',
    'h-leg-exercises',
    'h-back-exercises',
    'h-glute-exercises',
    'h-ab-exercises',
    'h-calves-exercises',
    'h-forearm-*', #Merge the 2 forarm exercises
    ] 


#Create empty 2d numpy array
myArr = np.empty((0,2), dtype="U40")

for h in header_names:

    #Extract muscle name from header
    muscle_re = re.compile('.*-([a-zA-Z]*)-.*')
    match = muscle_re.search(h)
    if match is None:
        continue
    muscle_name = match.group(1)

    #Retrive exercise names from lists
    exercise = chest = soup.find_all(id=re.compile(h))
    for e in exercise:
        lis = e.find_next_sibling().find_all('a')
        for li in lis:
            row = np.array([[li.getText(),muscle_name]])
            myArr = np.append(myArr,row,axis=0)
  
# print(os.getcwd())
# print(type(os.getcwd()))
pd.DataFrame(myArr, columns=['Name','Muscle']).to_csv(os.getcwd()+'/exercises.csv')
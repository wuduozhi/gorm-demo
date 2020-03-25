import requests
import random
import sys


def test_his_meter(n):
    year = 2016
    month = 1
    day = 1
    payload = {
        'meterno':'000000000004',
        'cctno':'111111111',
    }
    url = "http://127.0.0.1:10004/collect/v1/InsertHistMeterData"
    for i in range(n):
        if i % 100 == 0:    
            month += 1
            if month == 13:
                month = 1
                year += 1
        day += 1
        if day == 29:
            day = 1
        MDate = str(year) + "-" + str(month) +  "-" + str(day)
        lasttotalall = str(random.random())
        payload["MDate"] = MDate
        payload["lasttotalall"] = lasttotalall
        # payload["ID"] = i
        r = requests.post(url, data=payload)
        print(r.text)

def test_metedata():
    url = "http://127.0.0.1:8087/MeterData/addMeterData"
    payload = {
        'meterno':'000000000004',
        'cctno':'111111111',
        'ccttype':'D2',
        'Lst_Totalall':'9.0000',
        'Jsr_TotalAll':'12334',
        'MDate':'2020-2-18'
    }

    for i in range(1000):
        payload['meterno'] = str(i)
        r = requests.post(url, data=payload)
        print(r.text)

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("error")
        exit()
    
    method = sys.argv[1]
    if method == "his":
        n = int(sys.argv[2])
        test_his_meter(n)
    elif method == "data":
        test_metedata()
    else:
        print("unkonw")
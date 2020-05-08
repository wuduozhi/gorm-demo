import requests
import random
import sys
import json


def select_meter_info(n):
    payload = {
        'pageNo':'0',
        'pageSize':'50',
    }
    url = "http://127.0.0.1:8087/MeterInfo/searchDataByPara"
   
    r = requests.post(url, data=payload)
    result = json.loads(r.text)

    if n != result["totalNo"]:
        print("meter check fail\n")
    else:
        print("meter check success")
    
    meters = result["list"]
    # for m in meters:
    #     print(m["meterno"])

def select_cct_info(n):
    payload = {
        'pageNo':'0',
        'pageSize':'50',
    }
    url = "http://127.0.0.1:8087/CctInfo/selectCcInfoByPara"
   
    r = requests.post(url, data=payload)
    result = json.loads(r.text)
    if n != result["totalNo"]:
        print("cct check fail\n")
    else:
        print("cct check success")
    
    ccts = result["list"]
    # for c in ccts:
    #     print(c["cct_no"])

def select_company_info(n):
    payload = {
        'pageNo':'0',
        'pageSize':'50',
    }
    url = "http://127.0.0.1:8087/CompanyInfo/selectCompanyInfo"
   
    r = requests.post(url, data=payload)
    result = json.loads(r.text)
    if n != result["totalNo"]:
        print("company check fail\n")
    else:
        print("company check success")
    
    companys = result["list"]
    # for c in companys:
    #     print(c["name"])

def select_meter_data(n):
    payload = {
        'pageno':'0',
        'pagesize':'50',
        'usercode':'1',
    }
    url = "http://127.0.0.1:10004/collect/v1/getMeterDatalist"
    r = requests.get(url,params=payload)

    result = json.loads(r.text)
    if n != result["body"]["totalCount"]:
        print("meter-data check fail")
    else:
        print("meter-data check success")

def select_his_meter(n):
    payload = {
        'pageNo':'0',
        'pageSize':'50',
    }
    url = "http://127.0.0.1:8087/Failreadmeter/searchDataByPara"
    r = requests.post(url,data=payload)

    result = json.loads(r.text)
    if n != result["totalNo"]:
        print("his-meter-data check fail")
    else:
        print("his-meter-data check success")



if __name__ == '__main__':
    companyCount = 8
    select_company_info(companyCount)
    cctIfoCount = 40
    select_cct_info(cctIfoCount)
    meterInfoCount = 200
    select_meter_info(meterInfoCount)
    meterDataCount = 25
    select_meter_data(meterDataCount)
    hisMeterDataCount = 1600
    select_his_meter(hisMeterDataCount)
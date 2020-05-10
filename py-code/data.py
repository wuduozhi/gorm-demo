import requests
import random
import sys
import time
import datetime

class MeterData:
    def __init__(self,meterno,cctno,ccttype,Lst_Totalall,Jsr_TotalAll,MDate):
        self.meterno = meterno
        self.cctno = cctno
        self.ccttype = ccttype
        self.lasttotalall = Lst_Totalall
        self.jsrutotalall = Jsr_TotalAll
        self.MDate = MDate

    def payload(self):
        return self.__dict__

class MeterInfo:
    def __init__(self, user_id, meterno,custno,cct_no,snr_no,custName="",areano="",id=""):
        self.user_id = user_id
        self.meterno = meterno
        self.custno = custno
        self.cct_no = cct_no
        self.snr_no = snr_no
        self.custName = custName
        self.areano = areano
        self.id = id
    
    def payload(self):
        return self.__dict__

class CctInfo:
    def __init__(self, cct_no, cct_name,userid,cct_type,areano="",cct_status=""):
        self.cct_no = cct_no
        self.cct_name = cct_name
        self.userid = userid
        self.cct_type = cct_type
        self.areano = areano
        # self.cct_status = cct_status
    
    def payload(self):
        return self.__dict__

class CompanyInfo:
    def __init__(self, code, name,province,city,satus="",createTime="2020-3-1"):
        self.code = code
        self.name = name
        self.province = province
        self.city = city
        self.satus = satus
        # self.createTime = createTime

    def payload(self):
        return self.__dict__



def getCode(prefix):
    code_list = []
    for i in range(10): # 0-9数字
        code_list.append(str(i))
    for i in range(65, 91): # A-Z
        code_list.append(chr(i))
    for i in range(97, 123): # a-z
        code_list.append(chr(i))

    myslice = random.sample(code_list, 10)  # 从list中随机获取6个元素，作为一个片断返回
    verification_code = ''.join(myslice) # list to string
    return prefix+verification_code


def getProvince():
    l = []
    for i in range(130123,130134):
        l.append(i)
    index = random.randint(0,len(l)-1)
    return str(l[index])

def getCctType():
    cctTypes = ["D1","D2","G1","N1","Y0","Y1","Y2","Y3"]
    index = random.randint(0,len(cctTypes)-1)
    return str(cctTypes[index])

def getAreaNo():
    areas = []
    for i in range(10000004,10000024):
        areas.append(i)

    index = random.randint(0,len(areas)-1)
    return str(areas[index])

def getRandomInt(val):
    return str(random.randint(0,val))

def strTimeProp(start, end, prop, frmt):
    stime = time.mktime(time.strptime(start, frmt))
    etime = time.mktime(time.strptime(end, frmt))
    ptime = stime + prop * (etime - stime)
    return int(ptime)


def randomDate(start, end, frmt='%Y-%m-%d %H:%M:%S'):
    return time.strftime(frmt, time.localtime(strTimeProp(start, end, random.random(), frmt)))



def insertConpany(conpany):
    payload = company.payload()
    url = "http://127.0.0.1:8087/CompanyInfo/addCompanyInfo"
    r = requests.post(url, data=payload)
    print(r.text)

def insertCct(cct):
    payload = cct.payload()
    url = "http://127.0.0.1:8087/CctInfo/addCctInfo"
    r = requests.post(url, data=payload)
    print(r.text)

def insertMeter(cct):
    payload = cct.payload()
    url = "http://127.0.0.1:8087/MeterInfo/addMeterInfo"
    r = requests.post(url, data=payload)
    print(r.text)


def insertMeterData(meterData):
    payload = meterData.payload()
    url = "http://127.0.0.1:10004/collect/v1/InsertMeterData"
    r = requests.post(url, data=payload)
    print(r.text)

def insertHisMeterData(meterData):
    url = "http://127.0.0.1:10004/collect/v1/InsertHistMeterData"
    payload = meterData.payload()
    r = requests.post(url, data=payload)
    print(r.text)

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("error")
    

    companuCount = 8
    cctCount = 5
    meterCount = 5
    hisMeterDataCount = 8

    start = '2017-06-02 12:12:12'
    end = '2020-11-01 00:00:00'

    meterID = 1
    startTime = datetime.datetime.now()
    for i in range(companuCount):
        company_code = str(i+1) + getCode("company-code")
        company_name = getCode("company")
        company_province = getProvince()
        company_city = getProvince()
        company = CompanyInfo(company_code,company_name,company_province,company_city)
        insertConpany(company)

        for j in range(cctCount):
            cctName = getCode("cct-name")
            cctNo = getCode("cct-no")
            ccyType = getCctType()
            cctAreano = getAreaNo()
            cct = CctInfo(cctNo,cctName,company_code,ccyType,cctAreano)
            insertCct(cct)

            for k in range(meterCount):
                meterCustNo = company_code
                meterNo = getCode("meter-no")
                meterAreano = getAreaNo()
                meterCustName = getCode("meter-custom")
                meterSnr_no = getCode("meter-snr-no")
                meter = MeterInfo(company_code,meterNo,meterCustNo,cctNo,meterSnr_no,meterCustName,meterAreano,meterID)
                insertMeter(meter)
                meterID += 1

                for x in range(hisMeterDataCount):
                    Lst_Totalall = getRandomInt(10)
                    Jsr_TotalAll = getRandomInt(10)
                    MDate = randomDate(start,end)
                    meterData =  MeterData(meterNo,cctNo,ccyType,Lst_Totalall,Jsr_TotalAll,MDate)
                    insertHisMeterData(meterData)
                    if x == 0:
                        insertMeterData(meterData)


    endTime = datetime.datetime.now()
    print((endTime-startTime).seconds)
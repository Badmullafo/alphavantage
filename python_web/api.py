import requests, datetime
from datetime import date
from requests.exceptions import HTTPError


def days_since_date(last_date):
    #Compute date difference
    today = datetime.date.today()
    #print("The date today is {0}".format(today))

    delta =  today - last_date
    return(delta.days)

def myprint(response_json, day_range):

    sum = 0.0

    #Get just the time series data, not the meta data
    time_series_json = response_json['Time Series (Daily)']

    for time, elements in time_series_json.items():
        #print(time)
        #Parse the date header as time object for comparison - the date bit makes sure its a date and has no hour/min/second values
        date_time_obj = datetime.datetime.strptime(time, '%Y-%m-%d').date()

        #Return from loop if days is larger than limit
        if days_since_date(date_time_obj) > day_range:
            continue

        for subkey, value in elements.items():
            print("Key: {0}, Value: {1}".format(subkey, value))
            #Only add up the close
            if subkey ==  "4. close":
                sum += float(value)

    print("The sum is: {0}".format(sum))
      
def main():

    symbol = "MSFT"
    apikey = "demo"

    url = 'https://www.alphavantage.co/query?apikey='+apikey+'&function=TIME_SERIES_DAILY_ADJUSTED&symbol='+symbol

    try:
        #Don't have to convert the object to json later on, it is invoked as json
        response_json = requests.get(url).json()

        # If the response_json was successful, no Exception will be raised
        #response_json.raise_for_status()
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')  # Python 3.6
    except Exception as err:
        print(f'Other error occurred: {err}')  # Python 3.6
    else:
        print('Success!')
    # print(response_json)

    myprint(response_json, 5)

if __name__ == "__main__":
    main()
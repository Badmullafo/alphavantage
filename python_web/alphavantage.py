from http.server import BaseHTTPRequestHandler, HTTPServer
import socketserver, os, time, requests, datetime, time, threading
from requests.exceptions import HTTPError
from socketserver import ThreadingMixIn

#The port to listen on - can be defined via docker

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>https://pythonbasics.org</title></head>", "utf-8"))
        self.wfile.write(bytes("<p>Request: %s from host %s</p>" % (self.path, hostName), "utf-8"))
        self.wfile.write(bytes("<body>", "utf-8"))
        self.wfile.write(bytes("<p>Getting the last %s days worth of results for %s, the list is %s average is %s.</p>" % (nDays,symbol,valueList,str(avg)), "utf-8"))
        self.wfile.write(bytes("</body></html>", "utf-8"))

class ThreadingSimpleServer(ThreadingMixIn, HTTPServer):
    pass

def days_since_date(last_date):
    #Compute date difference
    today = datetime.date.today()

    delta =  today - last_date
    return(delta.days)

def get_api(url):

    try:
        #Don't have to convert the object to json later on, it is invoked as json with the .json bit

        response_json = requests.get(url).json()
        
        # If the response_json was successful, no Exception will be raised
        #response_json.raise_for_status()
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')  # Python 3.6
    except Exception as err:
        print(f'Other error occurred: {err}')  # Python 3.6
    else:
        print('Success - got a response from: '+url )
        
        return response_json

def get_values(response_json):

    valuelist = []

    #Get just the time series data, not the meta data
    time_series_json = response_json['Time Series (Daily)']

    for time, elements in time_series_json.items():
        #Parse the date header as date object for comparison - the date bit makes sure its a date and has no hour/min/second values
        date_time_obj = datetime.datetime.strptime(time, '%Y-%m-%d').date()

        #Return to start of loop if days is larger than limit (avoids adding values out of range to valuelist)
        if days_since_date(date_time_obj) > nDays:
            continue

        for subkey, value in elements.items():
            #Only add up the close
            if subkey ==  "4. close":
                valuelist.append(float(value))

    return valuelist

def average(lst):
    return round(sum(lst) / len(lst),2)

def main():

    global serverPort, symbol, apiKey, nDays, avg, hostName, valueList

    serverPort = int(os.getenv('LISTEN_PORT'))
    #The hostname of the server - in this case the docker container name
    hostName = 'localhost'
    symbol = os.getenv('SYMBOL')
    apiKey = os.getenv('APIKEY')
    nDays = int(os.getenv('NDAYS'))

    url='https://www.alphavantage.co/query?apikey='+apiKey+'&function=TIME_SERIES_DAILY_ADJUSTED&symbol='+symbol

    response_json = get_api(url)
    valueList=get_values(response_json)
    avg = average(valueList)

    webServer = ThreadingSimpleServer((hostName, serverPort), Handler)

    print("Server started http://%s:%s" % (hostName, serverPort))

    try:
        webServer.serve_forever()

    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print("Server stopped.")

if __name__ == "__main__":
    main()

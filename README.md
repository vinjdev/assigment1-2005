# Assigment 1

SEE the default page, this sends you to the main page

## Main page
Main page is an html page, which shows what possible url there are
INFO
POPULATION
STATUS

### INFO
has two parameters

val(mandatory) which is for two letter country code
limit(optional) which is for how many cities to be showned
    if defined it will only show the first countries in alphabetic order
    if not it shows all the cities
    use ?limt=<some_number>
    

### POPOULATION
has two paramter

val(mandatory) which is for two letter country code
years(optional) struct with two numbers, to restrict the population to be between the years 
    use ?limit=2010-2015
    has to be a 4 digit number
    if not provided, it chooses the range between 1960 and 2018

### STATUS
No paramters

request the two apis used, and checks if the apis are up and runninng
shows the version for the application
shows the uptime for when the service was deployed

## RENDER
main page:       https://assigment1-2005.onrender.com/countryinfo/v1/
info page:       https://assigment1-2005.onrender.com/countryinfo/v1/info/{val}
population page: https://assigment1-2005.onrender.com/countryinfo/v1/population/{val}
status page:     https://assigment1-2005.onrender.com/countryinfo/v1/status






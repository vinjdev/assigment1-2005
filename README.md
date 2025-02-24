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
    use ?limt=<some_number>

### POPOULATION
has two paramter

val(mandatory) which is for two letter country code
years(optional) struct with two numbers, to restrict the population to be between the years 
    use ?limit=2010-2015
    has to be a 4 digit number

### STATUS
No paramters

request the two apis used, and checks if the apis are up and runninng
shows the version for the application
shows the uptime for when the service was deployed


# nbd
Next business date.
------------------

This app aims to serve current or next or previous business date information via web server in order to facilitate different housekeeping script needs.

This comes up alot especially when you have various housekeeping checks or backups scripts that runs on the following day.  

To try it from command line, do:
go run prototype2.go get -t="utc" -b "2019-06-02"

-t to specify timezone.
-b to specify business date enquire. E.g. What is the next or previous business date for 2019-06-02.

The acceptable date format are 20190602 or 2019-06-02 or 2019/06/02 or 2019.06.02.

Todo:
Implement help options.
Create test script to test the application.
Create function(s) to read Bank holiday listing for UK.
Create Web server to serve post and get business day.  

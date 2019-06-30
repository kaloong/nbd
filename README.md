# nbd
Next business date.
------------------

This app aims to serve current or next or previous business date information via web server in order to facilitate different housekeeping script needs.

This comes up alot especially when you have checks or backups that is run from the next day.  

To run do:
go run prototype.go get -t="utc" -b "2019-06-02"
-t for timezone.
-b for business date enquire. E.g. What is the previous business date for 2019-06-02.

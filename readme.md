# rabbit watcher

Monitor RabbitMQ and store the statistics in InfluxDB.

Loads configuration from config.json in current directory or the
file specified in the command line.

Uses UDP to send metrics to InfluxDB

Running from cron. For example:

*/5 * * * * /opt/bin/rabbit_watcher /opt/etc/rabbit_watcher.json


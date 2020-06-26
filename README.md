# Go vs Node.js (w/Express.js) _Naive_ Benchmark

## Purpose

This benchmark aims to provide a naive gauge of performance on a typical simple production REST API (running with net/http and Express.js).
It is NOT meant to demonstrate the maximum performance achievable from Go and Node.js.

## Test Details

Receive a pair of numbers and return their sum as the result.

Request is a JSON message containing two numbers:

```json
{
  "a": 123,
  "b": 456
}
```

Response is a JSON message containing the result:

```json
{
  "result": 579
}
```

## Setup

**Testing Rig Specs**

```
OS: Ubuntu 20.04 LTS x86_64
Kernel: 5.4.0-37-generic
CPU: Intel i5-8350U (8 CPUs) @ 3.600GHz
Memory: 15871MiB
```

**Benchmark Tool**
Apache's _ab_ tool was used with the following options:

```bash
$ ab -p data.json -T application/json -c 1000 -n 50000 http://localhost:<PORT>/
```

**Go**

- net/http
- chi-go for routing

**Node.js**

- Express.js web application framework
- Disabled Etag and X-Powered-By in HTTP response headers
- PM2 for scaling with cluster mode

## Results

### Apache ab results

**Go**

```
Concurrency Level:      1000
Time taken for tests:   2.893 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      6700000 bytes
Total body sent:        8200000
HTML transferred:       850000 bytes
Requests per second:    17280.86 [#/sec] (mean)
Time per request:       57.867 [ms] (mean)
Time per request:       0.058 [ms] (mean, across all concurrent requests)
Transfer rate:          2261.36 [Kbytes/sec] received
                        2767.64 kb/s sent
                        5029.00 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   28   6.3     28      54
Processing:    10   29   6.3     29      56
Waiting:        0   18   7.4     16      45
Total:         38   57   2.9     57      84

Percentage of the requests served within a certain time (ms)
  50%     57
  66%     57
  75%     58
  80%     59
  90%     60
  95%     62
  98%     64
  99%     66
 100%     84 (longest request)
```

**Node.js**

Without PM2 (single thread):

```
Concurrency Level:      1000
Time taken for tests:   6.766 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      7900000 bytes
Total body sent:        8200000
HTML transferred:       800000 bytes
Requests per second:    7390.41 [#/sec] (mean)
Time per request:       135.311 [ms] (mean)
Time per request:       0.135 [ms] (mean, across all concurrent requests)
Transfer rate:          1140.32 [Kbytes/sec] received
                        1183.62 kb/s sent
                        2323.94 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   15  93.2      5    1023
Processing:    46  119  18.9    116     368
Waiting:       18   89  17.7     88     346
Total:         46  134  97.3    121    1391

Percentage of the requests served within a certain time (ms)
  50%    121
  66%    122
  75%    124
  80%    125
  90%    138
  95%    167
  98%    252
  99%    260
 100%   1391 (longest request)
```

With PM2 (multiple threads):

```
Concurrency Level:      1000
Time taken for tests:   6.931 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      7900000 bytes
Total body sent:        8200000
HTML transferred:       800000 bytes
Requests per second:    7213.76 [#/sec] (mean)
Time per request:       138.624 [ms] (mean)
Time per request:       0.139 [ms] (mean, across all concurrent requests)
Transfer rate:          1113.06 [Kbytes/sec] received
                        1155.33 kb/s sent
                        2268.39 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   5.3      1      37
Processing:    17  134  20.2    130     216
Waiting:       15  132  19.8    129     216
Total:         53  137  18.8    132     217

Percentage of the requests served within a certain time (ms)
  50%    132
  66%    142
  75%    147
  80%    152
  90%    163
  95%    174
  98%    188
  99%    194
 100%    217 (longest request)
```

### CPU load

![cpu load](https://github.com/joshuous/Go_vs_Node_API_Benchmark/raw/master/cpu_load.png)

## Observations

**Speed of handling requests**

For this Sum API, Go appears to handle the request around 2x faster than Node.js.

**CPU load**

Go: all CPUs running at around 20%-40% load

Node.js (single thread): one CPU running at around 90% load

Node.js (multiple threads w/PM2): all CPUs running at around 80%-90% load

**Others**

Interestingly, Node.js with a single thread handled requests at about the same speed as Node.js with PM2. Perhaps Node.js with PM2 will show a greater advantage for longer-running and more CPU intensive tasks.

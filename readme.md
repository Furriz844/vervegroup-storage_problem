# Storage API
## How to start and use app

- Run db using `docker-compose.yaml`
```
docker-compose up 
```
- Create table
```
CREATE TABLE promotions (
	id uuid NULL,
	price float4 NULL,
	expiration_date timestamptz NULL
);
```
- Compile and run app
```
make
./out/storage-api
```
- Put `promotions.csv` into `./resources/` directory
- Invoke `/admin/load` endpoint to upload data
```
curl -v http://localhost:8080/admin/load
```
- Get promotion using `/promotions/{uuid}` endpoint
```
curl -v http://localhost:8080/promotions/8aaa4862-c976-42a0-839e-43533837eb2c
```

## What should to be improvement
 - Testing - right now code is not covered by unit test. I beleive that it should be done with first priority.
 - Logging - To provide data for monitoring it's better to think more about that. Right now logging is not ideal. Things for improvements:
    - Log levels - use different log levels to understand what attention it required - INFO, ERROR, WARN. Also maybe it's better to add TRACE logs for debugging purpose.
    - Structured logs - If we use some external system (e.g. Splunk) it's better to have structured logs (json as example) to have oportunity to query them. Also it helps to build dashboards.
 - Migration sripts - automate some work
 - Configuration file - most variables (e.g. db configs) are hardcoded. To change them without building new version it's better to extract them to configuration file.
 - Storage - I beleive that relationship databases it's not good option for this app. It's seems that its better to store data in some distributed cache. But to choose the best solution for this task I need to know all scenarious how we want to use our data:
    - How we want to use this data? only return as response by uuid? 
    - Should we store data for history to create some statistic and charts? 
    - Do we have another entities that have relations with promotion entity? 
    - As an example for using SQL db we can add new field version: Users get only last version of data. If we want to upload new data we load it with new version. When uploading procedure is done, we switching all users for recieveing new version of data. It means that we should not block DB to update data, we can rollback to previos version if data corrupted and have historical data.
## Operating app in production
- Deployment
    - We can use canary deployment to limitate load on new version, make the switching of users to the new version seamless and have rollback option.
- Scaling
    - Depends on infrastructure: K8S, AWS lambda, e.t.c.
    - Think about caching to decrease load on storage. 
- Monitoring
    - Implement dashboards that uses logs as input data. Do not forget to monitor infrastructure (storage loading, free space e.t.c)

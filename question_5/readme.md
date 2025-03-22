# Question 5

You are tasked with refactoring a monolithic service that handles multiple responsibilities such as authentication, file uploads, and user data processing. The system has become slow and hard to maintain. How would you approach refactoring the service?
- What steps would you take to decompose the service into smaller, more manageable services?
- How would you ensure that the new system is backward compatible with the old one during the transition?

# Answer

1. What steps would you take to decompose the service into smaller, more manageable services?

- Before we consider migrate to Microservices, we can optimize the monolith service with the following:
  - Check Problem: We need to check where is the cause of slow, is it from database query, or third party slow response, or application problem.
  - Optimize Problem: 
    - Database Problem: We can optimize the query such as tuning the query, add index, add materialize view for faster query.
    - Application Problem: Many problem can accure from application problem, such as bottleneck, complex logic, CPU bound, logic is too long. We can optimize the problem with caching for faster response, fix the logic, we can consider async process to make it faster.
    - Third Party Problem: Sometime slow problem is not from our side, it is from third party side. We can optimize this with cache the result of third party response, follow up the third party client, and we can consider using circuit breaker so that the application will fail fast.
  - CQRS: If the problem is fixed but the performance is still quite slow, we can consider to using CQRS design to separate the read and write operation.

- If we want to migrate to Microservices, we can do the following step for migrating:
  - Identify Service Boundaries and Dependencies: Analyze and list the functionalities such as authentication, file uploads, etc, and dependencies such as databases, 3rd party services, etc. There functionalities and dependencies will form basis for new microservices architecture.
  - Design Microservices: Define the APIs and data models for each service. We need to ensure that each service has single responsibility and can scalable independently. Also we need to design the complexity communication from 1 service to other if the transaction need from service to service based on dependencies. We also need to plan what technology, frameworks and tool need to use for services.
  - Transision Plan: We need to create transition plan. The plan should include timeline, milestone, and rollback plan. This can help team incrementally migrate each service.
  - Implement Microservice: After everything is planned, its time for implementation microservices one by one start from least critical functionality to minimal risk and also for trial migrate to prepare for critical functionality.
  - Data Management: After services ready to use, we need to strategize where the database should be. Either we can use shared database, or different database each services. If we use different database, we need to migrate the data from main database to microservice database.
  - Dependencies Services Communication: We also need to strategize how to communicate each service if the functionality need other service. Either we can use RestAPI, gRPC, or message broker. Ensure that the communication is secure and realible.
  - Testing: After everything is already set up on staging environment, we need to test each service. Either its load testing, unit testing, integration testing and end-to-end testing, to ensure the functionality, the resources, and performance is good to go.
  - Monitoring and Logging: Implement monitoring and logging to ensure each service run as expected and no issue.

2. How would you ensure that the new system is backward compatible with the old one during the transition?

- To ensure the backward compability during transition, we can do this following:
  - Versioning: Implement API versioning to ensure both old and new API support all clients. This ensure that the client using old API can still use its functionality, and the client using new API can use new microservices.
  - Feature Toogle: Use feature toggle to enable or disable the new functionalities. This ensure us to test the new feature from microservices in production.
  - Persentage Rollout: Use percentage rollout, so that small percentage of user can use the new feature to see the system run as expected. Gradually increase the rollout and monitor the performance before full deployment.

# **Wallet**
We are going to have a microservice to keep all data of the user wallet. We need to have two
API to expose them to other microservices.
## **API Information:**
- **get-balance:**
  This API should return a JSON to show the current balance of a user. The parameter which is
  needed for this API is user_id and the output should be like the below sample:
>
> Input: user_id int
>
> Output: {"balance":4000}
>
- **add-money:**
  This API should add money to the wallet of a user and at the end return the transaction reference
  number. The parameter which is needed for this API is user_id and amount and the output should
  be like the below sample:
>
> Input: user_id int amount int (this parameter can be negative)
>
> Output: {"reference_id":12312312312}
>

## **Project Detail:**

Please consider the below points:

- Please Dockerize the project
- Use MySql as a database to store your data
- We need to Save all transaction logs of user
- We need an API to show balance of each user
- We need an API to add money to wallet of user
- We need to have some necessary test cases (just 6 test case to make sure you know
  about this procedure)
- We need a daily job to calculate total amount of transactions and print it on terminal
- You don’t have to develop any API or service for user, just develop the necessary
  services which are related to **wallet**


## **Run Project:**

```
cp config/config.env.example config.env
```

Edit your config file and run:
```
docker-compose up
```
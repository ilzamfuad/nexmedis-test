# Question 1

You are tasked with designing an API for an e-commerce platform. The system must support the following features:
- User registration and authentication
- Viewing and searching products
- Adding items to a shopping cart
- Completing a purchase
Design the RESTful endpoints for the above features. Describe your choice of HTTP methods (GET, POST, PUT, DELETE), URL structure, and the expected response formats. Assume that users need to authenticate before performing certain actions (e.g., adding items to the cart).

# Answer

### RESTful API Endpoints

#### User Registration and Authentication
- **POST /api/register**
  - Registers a new user.
  - Using POST because this API need to do something that require to process
  - Request body: `{ "username": "string", "password": "string", "email": "string" }`
  - Response: `201 Created` with user details or `400 Bad Request` for errors.
  - Response: 
    - `201 Created`: `{ "id": "string", "username": "string", "email": "string" }`
    - `400 Bad Request`: `{ "error": "string" }`
    - `500 Internal Server Error`: `{ "error": "string" }`

- **POST /api/login**
  - Authenticates a user.
  - Using POST because this API need to do something that require to process and validate
  - Request body: `{ "username": "string", "password": "string" }`
  - Response: `200 OK` with authentication token or `401 Unauthorized` for invalid credentials.
  - Response: 
    - `200 OK`: `{ "token": "access_token(string)", "ttl": "time_to_live_in_int(int)", "refresh_token": "refresh_token(string)" }`
    - `400 Bad Request`: `{ "error": "string" }`
    - `401 Unauthorized`: `{ "error": "string" }`
    - `500 Internal Server Error`: `{ "error": "string" }`

#### Viewing and Searching Products
- **GET /api/products?search=string&limit=number&offset=number&sort=string&order=string**
  - Retrieves a list of products.
  - Using Get because this API get resource from database
  - Query parameters: 
    - `search=string` for searching products.
    - `sort=string`: Sort products by a specific field (e.g., price, name).
    - `order=string`: Order products by ascending or descending (e.g., asc, desc).
    - `limit=number`: Number of products per page.
    - `offset=number`: Begin number from product page number for pagination.
  - Response: `200 OK` with a list of products.
  - Response: 
    - `200 OK`: `{ data: [{"id": "number", name: "string", "price": "number", "quantity": "number"}], meta: {"limit": "number", "offset": "number", "total": "number"}}`
    - `500 Internal Server Error`: `{ "error": "string" }`

- **GET /api/products/{id}**
  - Retrieves details of a specific product.
  - Using GET because this API get resource from database
  - Response: `200 OK` with product details or `404 Not Found` if the product does not exist.
  - Response: 
    - `200 OK`: `{ data: {"id": "number", name: "string", description: "string", "price": "number", "quantity": "number"} }`
    - `404 Not Found`: `{ "error": "string" }`
    - `500 Internal Server Error`: `{ "error": "string" }`

#### Adding Items to a Shopping Cart
- **POST /api/cart**
  - Adds an item to the authenticated user's shopping cart, Required Authentication.
  - Using POST because this API need to do something that require to process and validate
  - Header:
    - `Authentication: Bearer Token`
  - Request body: `{ "product_id": "string", "quantity": "number" }`
  - Response: `200 OK` with updated cart details or `400 Bad Request` for errors.
  - Response: 
    - `201 Created`: `{ data: {"id": "number", product: {"name": "string", "price": "string", "description": "string"}, "total_price": "number", "quantity": "number"}}`
    - `422 Unprocessable Entity`: `{ "error": "string" }`
    - `500 Internal Server Error`: `{ "error": "string" }`


- **GET /api/cart?limit=number&offset=number**
  - Retrieves the authenticated user's shopping cart, Required Authentication.
  - Using GET because this API get resource from database
  - Header:
    - `Authentication: Bearer Token`
  - Query parameters: 
    - `limit=number`: Number of products per page.
    - `offset=number`: Begin number from product page number for pagination.
  - Response: `200 OK` with cart details.
  - Response: 
    - `200 OK`: `{ data: [{"id": "number", product: {"name": "string", "price": "string", "description": "string"}, "total_price": "number", "quantity": "number"}], meta: {"limit": "number", "offset": "number", "total": "number"}}`
    - `500 Internal Server Error`: `{ "error": "string" }`

#### Completing a Purchase
- **POST /api/checkout**
  - Completes the purchase for the authenticated user's shopping cart, Required Authentication.
  - Using POST because this API need to do something that require to process and validate
  - Header:
    - `Authentication: Bearer Token`
  - Request body: `{ "cart": [{"id": "number", "total_price": "number", "quantity": "number"}], "total_price": "number" }`
  - Response: `200 OK` with order confirmation or `400 Bad Request` for errors.
  - Response: 
    - `201 Created`: `{ data: {"id": "number", "customer": {"id":"number", "email": "string", "name": "string"}, "cart": {"id": "number", "items": [{"id": "number", product: {"name": "string", "price": "string", "description": "string"}, "total_price": "number", "quantity": "number"}], "total_price": "number"}, "status": "string" }}`
    - `400 Bad Request`: `{ "error": "string" }`
    - `422 Unproccesable Entity`: `{ "error": "string" }`
    - `500 Internal Server Error`: `{ "error": "string" }`

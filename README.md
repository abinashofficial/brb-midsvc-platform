# BRB Mid Service Platform

This project is a GoLang microservice for managing service listings, vendor bookings, and basic user flows (admin/customer), built for Beauty Right Back (BRB).

It covers:
- Service and vendor management
- Basic booking with non-overlapping rules
- Notification simulation
- Simple user roles
- Booking summary per vendor
- Swagger documentation
- Basic retry handler and rate limiter
- Pagination for customer booking lists
- Docker-based deployment
- Unit Test Cases

---

## 🚀 Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/brb-midsvc-platform.git
cd brb-midsvc-platform


2. Run with Docker Compose

docker-compose up --build

This will spin up:
    PostgreSQL database
    GoLang API server on localhost:8080


📖 Swagger/OpenAPI Documentation
Once the service is running:

👉 Access Swagger UI here: http://localhost:8080/swagger/index.html

It contains:
    All API endpoints
    Request/response models
    Example payloads


📌 API Endpoints Overview
Base URL: http://localhost:8080/api

🧠 Health Check

Method	Endpoint	Description
GET	    /health	    Check DB health

🔧 Services (Admin Only)

Method	Endpoint	                Description
POST	/services/	                Create a service
PUT	    /services/:id		        Update a service 
PUT	    /services/assign-vendor	    Link service to vendor
PATCH	/services/toggle	        Toggle service status

🏢 Vendors (Admin Only)

Method	Endpoint	Description
POST	/vendors/	Create a vendor

📅 Bookings

Method	Endpoint	Description
POST	/bookings/	Create a booking
GET	    /bookings/	List all bookings

📊 Summary (Admin Only)

Method	Endpoint	            Description
GET	    /summary/vendor/:id	    Vendor booking summary

📚 Swagger Docs
View API documentation: http://localhost:8080/swagger/index.html



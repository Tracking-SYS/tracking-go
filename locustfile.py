from locust import HttpUser, task, between
import json
import secrets

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task(1)
    def fetch_product_pb(self):
        self.client.get("/v1/products?limit=10")

    @task(1)
    def fetch_product(self):
        self.client.get("/products?limit=10")

    @task(1)
    def get_home(self):
        self.client.get("/")

    @task(1)
    def get_product_detail(self):
        for item_id in range(1, 10):
            self.client.get(f"/v1/product/{item_id}", name="/item")

    @task(1)
    def create_product(self):   
        headers = {'Content-Type': 'application/json','Accept-Encoding':'gzip'}
        self.client.post("/v1/product", data=json.dumps(
            {"data": {"id": secrets.randbelow(200), "name": "Samsung", "price": secrets.randbelow(30000), "type": "simple", "status": secrets.randbelow(2)}}), 
            headers=headers, 
            name = "Create a new product")
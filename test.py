import requests
import json

# Base URL
BASE_URL = "http://localhost:8080/api"


# Sample data
user_data = {
    "email": "user1@example.com",
    "name": "John Doe",
    "mobile": "1234567890",
    "id": "1bc9e841-79f8-4050-88ea-76bd799326ae"
}

expense_data = {
    "description": "Dinner",
    "amount": 50.0,
    "splitMethod": "equal",
    "participants": [
        {"userId": "2bc9e841-79f8-4050-88ea-76bd799326ae", "amount": 25.0}
    ]
}

def login():
    response = requests.post("http://localhost:8080/login", json={"username": "user", "password": "user@1"})
    if response.status_code == 200:
        token = response.json().get("token")
        print(token)
        return token
    else:
        print("Failed to login:", response.status_code, response.text)
        return None

def test_create_user(token):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(f"{BASE_URL}/users", headers=headers, json=user_data)
    print("Create User Response:", response.status_code)
    print("Response Text:", response.text)
    try:
        print("Response JSON:", response.json())
    except requests.exceptions.JSONDecodeError as e:
        print("JSONDecodeError:", e)

def test_add_expense(token):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(f"{BASE_URL}/expenses", headers=headers, json=expense_data)
    print("Add Expense Response:", response.status_code, response.json())

def test_get_user_expenses(token):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(f"{BASE_URL}/expenses/{user_data['id']}", headers=headers)
    print("Get User Expenses Response:", response.status_code, response.json())

def test_fetch_all_expenses(token):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(f"{BASE_URL}/all-expenses", headers=headers)
    print("Fetch All Expenses Response:", response.status_code, response.json())

def test_download_balance_sheet(token):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(f"{BASE_URL}/balance-sheet", headers=headers)
    print("Download Balance Sheet Response:", response.status_code)
    if response.status_code == 200:
        with open('balance_sheet.json', 'wb') as file:
            file.write(response.content)
import requests

def test_protected_endpoint(token):
    headers = {
        "Authorization": f"Bearer {token}"
    }
    
    response = requests.get(f"{BASE_URL}/protected-endpoint", headers=headers)
    
    print("Status Code:", response.status_code)
    print("Response Headers:", response.headers)
    print("Raw Response Text:", response.text)

    if response.status_code == 401:
        print("Authorization failed: Unauthorized")
    elif response.status_code == 403:
        print("Authorization failed: Forbidden")
    else:
        print("Unexpected response:", response.status_code, response.text)

if __name__ == "__main__":
    token = login()
    if token:
        test_create_user(token)
        test_add_expense(token)
        test_get_user_expenses(token)
        test_fetch_all_expenses(token)
        test_download_balance_sheet(token)
        test_protected_endpoint(token)

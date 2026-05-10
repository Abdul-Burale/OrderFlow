# OrderFlow

OrderFlow is a multi-tenant backend API for order management. It is designed for businesses that need to manage users, products, orders, and payment workflows while keeping each tenant’s data isolated.

The project is currently in active development. The goal is to build a secure backend system using Go, MongoDB, Firebase Authentication, role-based access control, and Stripe payment links.

---

## Project Status

**Status:** In active development

This project is being built incrementally. The API structure, documentation, endpoint examples, authentication flow, and architecture diagrams will continue to evolve as the backend implementation is completed.

---

## Purpose

The main goal of OrderFlow is to demonstrate backend system design beyond basic CRUD operations.

The project focuses on:

- Multi-tenant data isolation
- Secure authentication
- Role-based access control
- Order and product management
- Stripe payment link generation
- Clear API boundaries
- Maintainable Go backend structure
- Developer-focused documentation

The important engineering challenge is not just creating endpoints. The key challenge is ensuring that every request is authenticated, scoped to the correct tenant, checked against user permissions, and safely handled by the API.

---

## Planned System Overview

OrderFlow is designed around a request flow similar to this:

```txt
Client / API Consumer
        ↓
Firebase Authentication
        ↓
Go REST API
        ↓
Tenant Resolution Middleware
        ↓
Role-Based Access Control Middleware
        ↓
Route Handlers
        ↓
MongoDB / Stripe
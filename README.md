# Blog Platform API

## Overview

This application serves as a backend for a blog platform, providing essential operations for managing users, blogs, and comments. The API includes features such as authentication (login, signup), user profile updates, password changes, blog creation, updating, and deletion. Users can also list all blogs, retrieve specific blogs, and filter blogs by topic. Additionally, the API supports comment management, allowing users to add, delete, and update comments.

## Technologies Used

- Golang
- Gin
- Gorm
- PostgreSQL

## ERD
```mermaid
erDiagram
users {
    integer id
    varchar(255) name
    varchar(255) email
    text password
    text image
}

blogs {
    integer id
    integer author_id
    varchar(255) title
    text content
    text image_cover
}

topics {
    integer id
    integer author_id
    varchar(255) name
    
}

blog_topics {
    integer blog_id
    integer topic_id
}

comments {
    integer id
    integer blog_id
    integer author_id
}

likes {
    integer id
    integer blog_id
    integer user_id
}
blogs ||--o{ comments:has
blogs ||--o{ likes:has
blogs }o--o{ topics:has
users ||--o{ comments:has
users ||--o{ likes:has
users ||-- o{ blogs:has
topics ||--o{ blog_topics:has
blogs ||--o{ blog_topics:has
```
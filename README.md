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
    id integer
    name varchar(255)
    email varchar(255)
    password text
    image text
}

blogs {
    id integer
    author_id integer
    title varchar(255)
    content text
    image_cover text
}

topics {
    id integer
    author_id integer
    name varchar(255)
    
}

blog_topics {
    blog_id integer
    topic_id integer
}

comments {
    id integer
    blog_id integer
    author_id integer
}

likes {
    id integer
    blog_id integer
    user_id integer
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
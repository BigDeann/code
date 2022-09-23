package com.fromzerotoexpert.Service;

import com.fromzerotoexpert.entity.User;

import java.util.List;

public interface UserService {
    // 业务接口类

    int add(User user);
    List<User> queryAll();
}

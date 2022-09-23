package com.fromzerotoexpert.UserServicelmpl;

import com.fromzerotoexpert.entity.User;
import com.fromzerotoexpert.mapper.UserMapper;
import com.fromzerotoexpert.Service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserMapper userMapper;


    @Override
    public int add(User user) {
        return userMapper.add(user);
    }

    @Override
    public List<User> queryAll() {
        return userMapper.queryAll();
    }

}

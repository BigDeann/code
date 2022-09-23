package com.fromzerotoexpert;

import com.fromzerotoexpert.entity.User;
import com.fromzerotoexpert.mapper.UserMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.List;

@SpringBootTest
class FromZerotoExpertApplicationTests {
    @Autowired
    private UserMapper userMapper;
    @Test
    void contextLoads() {
        User user = new User("123","1222");
        userMapper.add(user);
        List<User> users = userMapper.queryAll();
        users.forEach(System.out::println);

    }
}

package com.fromzerotoexpert.mapper;

import com.fromzerotoexpert.entity.User;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

import java.util.List;

@Mapper
public interface UserMapper {


    @Insert("INSERT INTO User VALUES(#{username},#{password})")
    int add(User user);
    @Select("SELECT * FROM User")
    List<User> queryAll();



/*    *//**
     * 根据用户名和密码查询用户对象
     *
     * @param username
     * @param password
     * @return
     *//*
    @Select("select * from User where username = #{username} and password = #{password}")
    static User select(@Param("username") String username, @Param("password") String password) {
        return null;
    }*/
}

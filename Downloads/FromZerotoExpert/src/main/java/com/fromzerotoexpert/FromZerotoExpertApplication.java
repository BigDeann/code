package com.fromzerotoexpert;


import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

//因为会连接到数据库，但是我又没连接则会访问不了web页面。    (exclude= DataSourceAutoConfiguration.class)
// 所以必须1.建立数据库 2.注视xml数据库连接的JDBC 3.排除 jdbc 的自动装配机制即可
@SpringBootApplication
public class FromZerotoExpertApplication {

    public static void main(String[] args) {
        SpringApplication.run(FromZerotoExpertApplication.class, args);
    }

}

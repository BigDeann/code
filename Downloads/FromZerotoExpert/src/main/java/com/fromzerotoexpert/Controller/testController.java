package com.fromzerotoexpert.Controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@RestController
public class testController {
    /**
     * 利用HttpServletRequest获取所有cookie
     * @param request
     * @param response
     * @return
     */
    @GetMapping ("/FromZerotoExpert")
    public String hello(HttpServletRequest request, HttpServletResponse response) {
        Cookie[] cookies = request.getCookies();
        //判断cookies是否为空
        if(cookies != null) {
            for (Cookie cookie : cookies) {
                if (cookie.getName().equals("userName")&&cookie.getValue().equals("Dean")) {
                    return "嗨嗨嗨，欢迎您再次来到 from zero to expert.";
                }
            }
        }
        Cookie newCookie = new Cookie("userName","Dean");
        newCookie.setMaxAge(24*60*60*7);
        response.addCookie(newCookie);
        return "嗨，欢迎您来到 from zero to expert.";
    }

}

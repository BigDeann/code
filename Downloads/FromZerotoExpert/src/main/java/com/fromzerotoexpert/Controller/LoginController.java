package com.fromzerotoexpert.Controller;

import com.baomidou.mybatisplus.extension.activerecord.Model;
import com.fromzerotoexpert.Service.UserService;
import com.fromzerotoexpert.entity.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

@Controller
public class LoginController {
    //将Service注入Web层bean注入
    @Autowired
    public UserService userService;

    @RequestMapping("/toLogin")//当你在浏览器输入/toLogin，我们会直接跳转到登陆界面
    public String toLogin(){
        return "login";
    }
    @RequestMapping("/LoginSuccess")//这个是判断是否成功登陆
    public String  LoginSuccess( User user, RedirectAttributes attributes){
        if(user.getUsername()!=null){
            attributes.addAttribute("kkk",user.getUsername());
            return "success";//我们就返回到我们的成功页面上边去
        }
        else {
            attributes.addAttribute("data","请输入你的密码");
            return "login";
        }

    }

    @RequestMapping("/toRegister")//这个是注册页面
    public String toRegister(){
        return "register";
    }

    @RequestMapping("/RegisterSuccess")//判断注册成功，然后将我们的数据放入数据库
    public String RegisterSuccess(User user){

        int add = userService.add(user);
        return "login";

    }
    @RequestMapping("/toMessage")//这个是显示我们的用户登陆信息（对应我们的显示用户信息按钮）
    public String toMessage(RedirectAttributes attributes){
        attributes.addAttribute("users",userService.queryAll());
        return "showAll";
    }


}

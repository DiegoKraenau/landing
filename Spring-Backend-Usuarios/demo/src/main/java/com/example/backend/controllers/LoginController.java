package com.example.backend.controllers;

import java.util.ArrayList;
import java.util.Base64;
import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.AuthorityUtils;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.example.backend.entities.User;
import com.example.backend.services.IUserService;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

@RestController
@RequestMapping("/api/login")
public class LoginController {

	@Autowired
	private IUserService userService;
	
	
	//JWT
	@PostMapping
	public ResponseEntity<User> login(@RequestParam("user") String username, @RequestParam("password") String pwd) throws Exception {
		/*
		String token = getJWTToken(username);
		User user = new User();
		user.setEmail(username);
		user.setToken(token);	
		*/
		List<User> listaUsuarios=new ArrayList<>();
		listaUsuarios=userService.findAll();
		User userLogged=new User();
		Boolean finded=false;
	
		for (User user : listaUsuarios) {
			byte[] byteArray = Base64.getDecoder().decode(user.getPassword().getBytes());
			String passwordDecrypt = new String(byteArray);
			if(user.getEmail().equals(username)&& passwordDecrypt.equals(pwd)) {
				finded=true;
				userLogged=user;
			}
		}

		if(finded==true) {
			
			String token=getJWTToken(username);
			userLogged.setEmail(username);
			//userLogged.setPassword(pwd);
			userLogged.setToken(token);
			
			return new ResponseEntity<User>(userLogged, HttpStatus.OK);
		}else {
			return new ResponseEntity<User>(HttpStatus.INTERNAL_SERVER_ERROR);

		}
		
		
	}

	private String getJWTToken(String username) {
		String secretKey = "mySecretKey";
		List<GrantedAuthority> grantedAuthorities = AuthorityUtils
				.commaSeparatedStringToAuthorityList("ROLE_USER");
		
		String token = Jwts
				.builder()
				.setId("softtekJWT")
				.setSubject(username)
				.claim("authorities",
						grantedAuthorities.stream()
								.map(GrantedAuthority::getAuthority)
								.collect(Collectors.toList()))
				.setIssuedAt(new Date(System.currentTimeMillis()))
				.setExpiration(new Date(System.currentTimeMillis() + 600000))
				.signWith(SignatureAlgorithm.HS512,
						secretKey.getBytes()).compact();

		return "Bearer " + token;
	}
}

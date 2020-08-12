package com.example.backend.controllers;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PruebaController {

	@RequestMapping("/helloPublic")
	public String helloPublic() {
		return "Hello Test-Publico";
	}
	
	@RequestMapping("/helloPrivate")
	public String helloPrivate() {
		return "Hello Test-Private";
	}
	
	
}

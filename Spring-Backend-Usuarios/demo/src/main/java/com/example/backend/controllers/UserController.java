package com.example.backend.controllers;

import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Base64;
import java.util.Date;
import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.AuthorityUtils;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import com.example.backend.entities.User;
import com.example.backend.services.IUserService;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

@RestController
@RequestMapping("/api/users")
public class UserController {
	

	@Autowired
	private IUserService userService;
	
	
	@GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<List<User>> findAll(){
		
		try {
			List<User> users = new ArrayList<>();
			users = userService.findAll();
			return new ResponseEntity<List<User>>(users, HttpStatus.OK);
			
		} catch (Exception e) {
			// TODO: handle exception
			return new ResponseEntity<List<User>>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	@PostMapping(consumes=MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<User> registerUser(@Validated @RequestBody User user){
		
		try {
			User userNew = new User();
			user.setPassword(Base64.getEncoder().encodeToString(user.getPassword().getBytes()));//Se encripta la contrasñea
			userNew=userService.save(user);
			
			return new ResponseEntity<User>(userNew, HttpStatus.OK);
			
		} catch (Exception e) {
			// TODO: handle exception
			return new ResponseEntity<User>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	@PostMapping(value = "/{id}/updatePhoto",produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<String> uploadPhoto(@PathVariable("id") UUID id,@RequestParam("file") MultipartFile foto) throws Exception{
		
			if (!foto.isEmpty()) {
				User userNew=new User();
				try {
					
					//Se procede a poner la foto
					Path directorioRecursos = Paths.get("src//main//resources//static//uploads");
					String rootPath = directorioRecursos.toFile().getAbsolutePath();
					
					byte[] bytes = foto.getBytes();
					Path rutaCompleta = Paths.get(rootPath + "//" + foto.getOriginalFilename());
					Files.write(rutaCompleta, bytes);
					userNew=userService.findById(id).get();
					userNew.setFoto(foto.getOriginalFilename());
					userNew=userService.save(userNew);
					return new ResponseEntity<String>(userNew.getFoto(), HttpStatus.OK);
					
				} catch (Exception e1) {
					// TODO Auto-generated catch block
					
					e1.printStackTrace();
					return new ResponseEntity<String>("No se encontre el usuario",HttpStatus.BAD_REQUEST);
				}
			
				
			}else {
				
				return new ResponseEntity<String>("Debe ingresar una foto",HttpStatus.BAD_REQUEST);
			}
			
		
	}
	
	@GetMapping(value="/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<User> getById(@PathVariable("id") UUID id){
		
		try {

			
			return new ResponseEntity<User>(userService.findById(id).get(), HttpStatus.OK);
		} catch (Exception e) {
			
			return null;
		}
		
	}

	@PutMapping(consumes=MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<User> updateUser(@Validated @RequestBody User user){
		try {
			
			return new ResponseEntity<User>(userService.save(user),HttpStatus.OK);
			
		} catch (Exception e) {
			// TODO: handle exception
			return new ResponseEntity<User>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	
	
	
}

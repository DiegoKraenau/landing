package com.example.backend.impl;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.example.backend.entities.User;
import com.example.backend.repository.UserRepository;
import com.example.backend.services.IUserService;

@Service
@Transactional(readOnly = true)
public class UserServiceImpl implements IUserService{
	
	@Autowired
	UserRepository userRepository;
	
	
	
	

	@Override
	@Transactional
	public User save(User t) throws Exception {
		// TODO Auto-generated method stub
		
		//t.setPassword(Base64.getEncoder().encodeToString(t.getPassword().getBytes()));
		return userRepository.save(t);
	}

	@Override
	public void deleteById(int id) throws Exception {
		// TODO Auto-generated method stub
		
	}

	

	@Override
	public List<User> findAll() throws Exception {
		// TODO Auto-generated method stub
		return userRepository.findAll();
	}

	@Override
	public Optional<User> findById(UUID id) throws Exception {
		// TODO Auto-generated method stub
		return userRepository.findById(id);
		
		
	}

}

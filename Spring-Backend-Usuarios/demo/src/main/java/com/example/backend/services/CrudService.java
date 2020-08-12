package com.example.backend.services;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface CrudService<T> {
	T save(T t) throws Exception;
	void deleteById(int id)throws Exception;
	Optional<T> findById(UUID id)throws Exception;
	List<T> findAll()throws Exception;

}
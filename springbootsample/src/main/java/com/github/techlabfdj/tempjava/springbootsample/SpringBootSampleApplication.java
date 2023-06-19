package com.github.techlabfdj.tempjava.springbootsample;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class SpringBootSampleApplication {
	@RequestMapping("/")
	String home() {
		return "Hello World!\n";
	}

	@RequestMapping("/gc")
	String gc() {
		System.gc();
		return "GC Done !";
	}

	public static void main(String[] args) {
		SpringApplication.run(SpringBootSampleApplication.class, args);
	}

}

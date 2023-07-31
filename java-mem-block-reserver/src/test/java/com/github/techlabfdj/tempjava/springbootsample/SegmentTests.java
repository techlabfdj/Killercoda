package com.github.techlabfdj.tempjava.springbootsample;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;

import org.slf4j.LoggerFactory;
import org.slf4j.Logger;

@SpringBootTest
public class SegmentTests {
  private Logger logger = LoggerFactory.getLogger(SegmentTests.class);
  @Autowired
  private Jackson2ObjectMapperBuilder mapperBuilder;
 
  @Test
  void jsonDeserialize() throws JsonMappingException, JsonProcessingException {
    String json = "{\"size\":16}";
    Segment s = mapperBuilder.build().readValue(json, Segment.class);
    logger.info("json serialization from {} resulted in {}", json, s);
  }

  @Test
  void jsonSerialize() throws JsonProcessingException,Exception {
    Segment s = new Segment(16);
    String serializedS = mapperBuilder.build().writeValueAsString(s);
    logger.info("initial serialized segment: {}", serializedS);
    Segment newS = mapperBuilder.build().readValue(serializedS, Segment.class);
    String serializedNewS = mapperBuilder.build().writeValueAsString(newS);
    logger.info("resulting serialized segment: {}", serializedNewS);
    if (!s.equals(newS)) {
      throw new Exception("deserialization of a serialized segment failed");
    }
	}
}

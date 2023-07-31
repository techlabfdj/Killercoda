package com.github.techlabfdj.tempjava.springbootsample;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import org.springframework.web.server.ResponseStatusException;

import com.github.techlabfdj.tempjava.springbootsample.SegmentRepository.SegmentList;

import org.springframework.http.HttpStatus;
import org.springframework.validation.annotation.Validated;
import javax.validation.constraints.Max;
import javax.validation.constraints.Min;

import org.slf4j.LoggerFactory;
import org.slf4j.Logger;

import java.util.UUID;

@RestController
@RequestMapping("/segments")
@Validated
class SegmentController {
	Logger logger = LoggerFactory.getLogger(SegmentController.class);

	private final SegmentRepository repository = new SegmentRepositoryMemoryImpl();

	@PostMapping
	@ResponseStatus(HttpStatus.CREATED)
	public Segment create(@RequestBody Segment segment) {
		logger.debug("received creation for {}", segment.toString());
		if (repository.create(segment)) {
			logger.info("created {}", segment.toString());
			return segment;
		}
		logger.error("failed to create {}", segment.toString());
		throw new ResponseStatusException(HttpStatus.CONFLICT, "uuid conflict detected");
	}

	@DeleteMapping(value = "/{uuid}")
  @ResponseStatus(HttpStatus.NO_CONTENT)
	public void delete(@PathVariable("uuid") UUID uuid) {
		logger.debug("received delete for {}", uuid);
		Segment deleted = repository.delete(uuid);
		if (deleted == null) {
			logger.warn("failed to delete segment with uuid '{}': not found", uuid.toString());
			throw new ResponseStatusException(HttpStatus.NOT_FOUND, "uuid not found");
		}
		logger.info("deleted {}", deleted.toString());
		return;
	}
	
	@GetMapping(value = "/{uuid}")
  @ResponseStatus(HttpStatus.OK)
	public Segment get(@PathVariable("uuid") UUID uuid) {
		logger.debug("received get for {}", uuid);
		Segment s = repository.get(uuid);
		if (s == null) {
			logger.warn("failed to get segment with uuid '{}': not found", uuid.toString());
			throw new ResponseStatusException(HttpStatus.NOT_FOUND, "uuid not found");
		}
		logger.info("got {}", s.toString());
		return s;
	}

	@PutMapping(value = "/{uuid}")
  @ResponseStatus(HttpStatus.OK)
	public Segment put(@PathVariable("uuid") UUID uuid, @RequestBody Segment segment) {
		logger.debug("received update for {} with {}", uuid, segment.toString());
		Segment current_segment = repository.get(uuid);
		if (current_segment != null) {
			segment.created_at = current_segment.created_at;
			segment.id = current_segment.id;
			if (repository.update(segment)) {
				logger.info("updated {} with {}", uuid, segment.toString());
				return segment;
			}
		}
		logger.warn("failed to update {} with {}: uuid not found", uuid, segment.toString());
		throw new ResponseStatusException(HttpStatus.NOT_FOUND, "uuid not found");
	}

	@GetMapping
  @ResponseStatus(HttpStatus.OK)
	public SegmentList list(@RequestParam(value = "index", defaultValue = "0") @Min(0) int index,
			@RequestParam(value = "limit", defaultValue = "10") @Min(1) @Max(20) int limit) {
		return repository.list(index, limit);
	}

}

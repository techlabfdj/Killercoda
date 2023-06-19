package com.github.techlabfdj.tempjava.springbootsample;

import java.util.ArrayList;
import java.util.Enumeration;
import java.util.concurrent.ConcurrentHashMap;

import java.util.UUID;

public class SegmentRepositoryMemoryImpl implements SegmentRepository {

  private final ConcurrentHashMap<UUID, Segment> segments = new ConcurrentHashMap<UUID, Segment>(100);
  
  @Override
  public Segment get(UUID id) {
    return segments.get(id);
  }

  @Override
  public boolean create(Segment segment) {
    return segments.putIfAbsent(segment.id, segment) == null;
  }

  @Override
  public boolean update(Segment segment) {
    if (!segments.containsKey(segment.id)) {
      return false;
    }
    segments.put(segment.id, segment);
    return true;
  }

  @Override
  public Segment delete(UUID id) {
    return segments.remove(id);
  }

  @Override
  public SegmentList list(int index, int limit) {
    SegmentList list = new SegmentList();
    list.total_count = segments.mappingCount();
    list.segments = new ArrayList<Segment>(limit);
    Enumeration<Segment> enumeration = segments.elements();
    if (index > 0) {
      int i = 0;
      while (i < index && enumeration.hasMoreElements()) {
        enumeration.nextElement();
        i++;
      }
    }
    int i = 0;
    while ((i < limit) && enumeration.hasMoreElements()) {
      list.segments.add(enumeration.nextElement());
      i++;
    }
    list.count = i;
    return list;
  }
}

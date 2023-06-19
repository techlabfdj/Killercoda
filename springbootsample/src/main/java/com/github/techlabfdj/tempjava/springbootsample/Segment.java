package com.github.techlabfdj.tempjava.springbootsample;

import java.time.Instant;
import java.util.Objects;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonCreator;

import org.slf4j.LoggerFactory;
import org.slf4j.Logger;

@JsonPropertyOrder({"id","size","unit","data"})
public class Segment {
	Logger logger = LoggerFactory.getLogger(SegmentController.class);

    public class SizeException extends RuntimeException {
        public SizeException(String message) {
            super(message);
        }
    }

    public enum Units {
        BYTES,
        KILOBYTES,
        MEGABYTES,
        GIGABYTES;
        
        public int getByteMultiplier() {
            switch (this) {
                default:
                    return 1;
                case KILOBYTES:
                    return 1024;
                case MEGABYTES:
                    return 1024 * 1024;
                case GIGABYTES:
                    return 1024 * 1024 * 1024;
            }
        }

        public int maxValue() {
            return Integer.MAX_VALUE / getByteMultiplier();
        }
    }

    public UUID id;
    public int size;
    public Units unit;
    protected byte[] data;
    public Instant created_at;
    public Instant updated_at;


    public Segment(int size) {
        this(size, Units.BYTES);
    }

    @JsonCreator
    public Segment(@JsonProperty("size") int size, @JsonProperty("unit") Units unit) {
        this.id = UUID.randomUUID();
        Instant now = Instant.now();
        this.created_at = now;
        this.updated_at = now;
        this.size = 0;
        this.data = null;
        this.unit = Units.BYTES;
        this.setSize(size, unit);
    }
    
    public void setSize(int size, Units unit) {
        unit = unit != null ? unit : Units.BYTES;
        if (size == this.size && unit == this.unit) {
            return;
        }
        if (size > unit.maxValue()) {
            throw new SizeException("size exceeds maximum allowed capacity");
        }
        if (size < 1) {
            throw new SizeException("size must be at least greater or equal to 1");
        }
        this.size = size;
        this.unit = unit;
        this.data = new byte[size * unit.getByteMultiplier()];
    }

    @Override
    public int hashCode() {
        return Objects.hash(id,size,data);
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        } else if (!(obj instanceof Segment)) {
            return false;
        } else {
            Segment other = (Segment) obj;
            return id.equals(other.id) && size == other.size && unit == other.unit;
        }
    }

    @Override
    public String toString() {
        return String.format("Segment [id=%s,size=%d,unit=%s]",id.toString(),size,unit.toString()) ;
    }
}

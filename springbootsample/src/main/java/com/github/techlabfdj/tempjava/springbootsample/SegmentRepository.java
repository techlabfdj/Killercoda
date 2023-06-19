package com.github.techlabfdj.tempjava.springbootsample;

import java.util.List;
import java.util.UUID;

public interface SegmentRepository {
    public class SegmentList {
        public int count;
        public long total_count;
        public List<Segment> segments;
    }

    /**
     * Permet de récupérer tous les enregistrements de la table Segment
     * @return Liste de Segment
     */
    public SegmentList list(int index, int limit);

    /**
     * Permet de rechercher un Segment à partir de son identifiant
     * @param id
     * @return Segment
     */
    public Segment get(UUID id);

    /**
     * Permet d'ajouter un nouvel enregistrement de la table Segment
     * @param Segment
     * @return Segment
     */
    public boolean create(Segment segment);

    /**
     * 
     * Permet de supprimer un Segment
     * @param Segment
     */
    public Segment delete(UUID id);

    /**
     * Permet de mettre à jour un Segment 
     * @param Segment
     */
    public boolean update(Segment segment);
}

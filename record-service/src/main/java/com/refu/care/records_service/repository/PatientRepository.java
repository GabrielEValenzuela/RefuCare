package com.refu.care.records_service.repository;

import org.springframework.stereotype.Repository;
import org.springframework.data.jpa.repository.JpaRepository;
import com.refu.care.records_service.model.Patient;

@Repository
public interface PatientRepository extends JpaRepository<Patient, Long> {
    // Custom query methods can be defined here if needed
}
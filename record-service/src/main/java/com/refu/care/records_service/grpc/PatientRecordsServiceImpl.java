package com.refu.care.records_service.grpc;


import io.grpc.stub.annotations.GrpcGenerated;
import io.grpc.stub.StreamObserver;


@GrpcGenerated
public class PatientRecordsServiceImpl extends PatientRecordsGrpc.PatientRecordsImplBase {
    
    public void AppendVitals(VitalsRequest request, StreamObserver<StatusResponse> responseObserver) {
        // Implement the logic to append vitals here
        StatusResponse response = StatusResponse.newBuilder()
                .setStatus("Vitals appended successfully")
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    public void GetPatientInfo(PatientId request, StreamObserver<PatientInfo> responseObserver) {
        // TODO: Fetch patient information from database

        PatientInfo info = PatientInfo.newBuilder()
                .setName("Test Patient")
                .setAge(45)
                .setHasDiabetes(false)
                .build();

        responseObserver.onNext(info);
        responseObserver.onCompleted();
    }
}

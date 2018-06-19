import demorpcservice_pb2 as pb
import demorpcservice_pb2_grpc as rpcservice
import grpc


if __name__ == "__main__":
    channel = grpc.insecure_channel('grpc-server:6000')
    client = rpcservice.DemoRPCServiceStub(channel)

    # this request will fail as 'surname' field can't be blank
    print("- TEST #1:\n")

    try:
        client.CreatePerson(
            pb.CreatePersonArgs(
                name='Gianfranco',
                age=36,
            )
        )
    except grpc.RpcError as e:
        print("received grpc.RpcError: %s" % e.details())

    # successfully create a person
    print("\n- TEST #2:\n")

    a_person = client.CreatePerson(
        pb.CreatePersonArgs(
            name='Gianfranco',
            surname='Reppucci',
            age=36,
        )
    )

    print("person created with id: %s" % a_person.id)

    # get person from id
    print("\n- TEST #3:\n")

    person_clone = client.ReadPerson(
        pb.ReadPersonArgs(
            id=a_person.id,
        )
    )

    print("got person with id: %s, has name: %s and surname: %s" %
          (person_clone.id, person_clone.name, person_clone.surname))

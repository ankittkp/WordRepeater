
package main

import (
	"context"
	pb "github.com/jinxankit/WordRepeater/proto/frequency"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewWordFrequencyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Calculate(ctx, &pb.InputRequest{Text: "Another antigen, the Rh antigen similar to one present in Rhesus monkeys\n(hence Rh), is also observed on the surface of RBCs of majority (nearly 80\nper cent) of humans. Such individuals are called Rh positive (Rh+ve)\nand those in whom this antigen is absent are called Rh negative (Rh-ve).\nAn Rh-ve person, if exposed to Rh+ve blood, will form specific antibodies\nagainst the Rh antigens. Therefore, Rh group should also be matched\nbefore transfusions. A special case of Rh incompatibility (mismatching)\nhas been observed between the Rh-ve blood of a pregnant mother with\nRh+ve blood of the foetus. Rh antigens of the foetus do not get exposed to\nthe Rh-ve blood of the mother in the first pregnancy as the two bloods are\nwell separated by the placenta. However, during the delivery of the first\nchild, there is a possibility of exposure of the maternal blood to small\namounts of the Rh+ve blood from the foetus. In such cases, the mother\nstarts preparing antibodies against Rh antigen in her blood. In case of\nher subsequent pregnancies, the Rh antibodies from the mother (Rh-ve)\ncan leak into the blood of the foetus (Rh+ve) and destroy the foetal RBCs.\nThis could be fatal to the foetus or could cause severe anaemia and\njaundice to the baby. This condition is called erythroblastosis foetalis.\nThis can be avoided by administering anti-Rh antibodies to the mother\nimmediately after the delivery of the first child.",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %v", r.Result)
}
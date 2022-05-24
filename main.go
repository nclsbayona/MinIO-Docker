package main
//El siguiente codigo fue diseñado apartir de las especificaciones de la documentacion de MinIO. En este codigo se hace uso del API de AWS S3 para listar los "buckets" o los espacios, podriamos decir que tenemos creados para guardar objetos. Es importante mencionar que este script se provo en un ambiente de pruebas, haciendo uso de un contenedor de Docker.
import (
    "context"
    "fmt"
    "io"
	"os"
	"flag"
    "time"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// `listMIOBuckets` is a function that takes in a writer, minioAccessKeyID, minioKeySecret, and
// minioserverURL and returns an error
func listMIOBuckets(w io.Writer, minioAccessKeyID string, minioKeySecret string, minioserverURL string) error {
    // minioAccessKeyID := "MinIO Access Key ID"
    // minioAccessKeySecret := "MinIO Access Key Secret"
    // minioServerURL := "MinIO Server URL"
	sess := session.Must(session.NewSession(&aws.Config{
        Region:      aws.String("auto"),
        Endpoint:    aws.String(minioserverURL),
        Credentials: credentials.NewStaticCredentials(minioAccessKeyID, minioKeySecret, ""),
    }))
    client := s3.New(sess)
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Second*10)
    defer cancel()
    result, err := client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
    if err != nil {
        return fmt.Errorf("ListBucketsWithContext: %v", err)
    }
    fmt.Fprintf(w, "Buckets:")
    for _, b := range result.Buckets {
        fmt.Fprintf(w, "\n - %s\n", aws.StringValue(b.Name))
    }
    return nil
}

var minioAccessKeyID *string = flag.String("k", "", "MinIO Access Key ID")
var minioKeySecret *string = flag.String("s", "", "MinIO Access Key Secret")
var minioserverURL *string = flag.String("u", "","MinIO Server URL, default (localhost:9000)")

// It takes a writer, a minio access key, a minio secret key, and a minio server URL, and it lists all
// the buckets on the minio server
func main() {
	flag.Parse()
	listMIOBuckets(os.Stdout, *minioAccessKeyID, *minioKeySecret, *minioserverURL)
}
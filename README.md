# Notes

[Google Cloud Platform provides the following ways of authenticating from your local machine](https://medium.com/google-cloud/local-remote-authentication-with-google-cloud-platform-afe3aa017b95):

- Creating a service account, downloading a corresponding JSON credentials file and setting the environmental variable of `JSON` file path to `GOOGLE_APPLICATION_CREDENTIALS`.
- Using gcloud auth application-default login to authenticate with a user identity (via a web flow) but using the credentials as a proxy for a service account.
- Running the gcloud auth login command to authenticate with a user identity (via web flow) which then authorizes gcloud and other SDK tools to access Google Cloud Platform. Note that this auth command is also one of the suite of commands run when gcloud init is run.
- Using GCP service specific methods such as generating Private/Public key pairs for use with Core IoT Auth when your local machine is an embedded device. These service specific examples are out of scope for this article.

### Links
- [Create projects independent of $GOPATH using Go Modules](https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)

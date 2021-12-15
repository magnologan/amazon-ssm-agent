// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package textractiface provides an interface to enable mocking the Amazon Textract service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package textractiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/textract"
)

// TextractAPI provides an interface to enable mocking the
// textract.Textract service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Textract.
//    func myFunc(svc textractiface.TextractAPI) bool {
//        // Make svc.AnalyzeDocument request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := textract.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTextractClient struct {
//        textractiface.TextractAPI
//    }
//    func (m *mockTextractClient) AnalyzeDocument(input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTextractClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type TextractAPI interface {
	AnalyzeDocument(*textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error)
	AnalyzeDocumentWithContext(aws.Context, *textract.AnalyzeDocumentInput, ...request.Option) (*textract.AnalyzeDocumentOutput, error)
	AnalyzeDocumentRequest(*textract.AnalyzeDocumentInput) (*request.Request, *textract.AnalyzeDocumentOutput)

	AnalyzeExpense(*textract.AnalyzeExpenseInput) (*textract.AnalyzeExpenseOutput, error)
	AnalyzeExpenseWithContext(aws.Context, *textract.AnalyzeExpenseInput, ...request.Option) (*textract.AnalyzeExpenseOutput, error)
	AnalyzeExpenseRequest(*textract.AnalyzeExpenseInput) (*request.Request, *textract.AnalyzeExpenseOutput)

	DetectDocumentText(*textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error)
	DetectDocumentTextWithContext(aws.Context, *textract.DetectDocumentTextInput, ...request.Option) (*textract.DetectDocumentTextOutput, error)
	DetectDocumentTextRequest(*textract.DetectDocumentTextInput) (*request.Request, *textract.DetectDocumentTextOutput)

	GetDocumentAnalysis(*textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error)
	GetDocumentAnalysisWithContext(aws.Context, *textract.GetDocumentAnalysisInput, ...request.Option) (*textract.GetDocumentAnalysisOutput, error)
	GetDocumentAnalysisRequest(*textract.GetDocumentAnalysisInput) (*request.Request, *textract.GetDocumentAnalysisOutput)

	GetDocumentTextDetection(*textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error)
	GetDocumentTextDetectionWithContext(aws.Context, *textract.GetDocumentTextDetectionInput, ...request.Option) (*textract.GetDocumentTextDetectionOutput, error)
	GetDocumentTextDetectionRequest(*textract.GetDocumentTextDetectionInput) (*request.Request, *textract.GetDocumentTextDetectionOutput)

	StartDocumentAnalysis(*textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error)
	StartDocumentAnalysisWithContext(aws.Context, *textract.StartDocumentAnalysisInput, ...request.Option) (*textract.StartDocumentAnalysisOutput, error)
	StartDocumentAnalysisRequest(*textract.StartDocumentAnalysisInput) (*request.Request, *textract.StartDocumentAnalysisOutput)

	StartDocumentTextDetection(*textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error)
	StartDocumentTextDetectionWithContext(aws.Context, *textract.StartDocumentTextDetectionInput, ...request.Option) (*textract.StartDocumentTextDetectionOutput, error)
	StartDocumentTextDetectionRequest(*textract.StartDocumentTextDetectionInput) (*request.Request, *textract.StartDocumentTextDetectionOutput)
}

var _ TextractAPI = (*textract.Textract)(nil)
package constant

// Method API Web Service
const (
	POST string = "POST"
	GET  string = "GET"
)

const (
	Required string = "required"
	Min      string = "min"
	Max      string = "max"
	Length   string = "len"
	Numeric  string = "numeric"
)

const (
	SuccessCode                  string = "0"
	RequiredField                string = "101000"
	InvalidCode                  string = "102000"
	OverScope                    string = "106000"
	UnderScope                   string = "107000"
	OverLengthCode               string = "108000"
	NoMatchTypeNumeric           string = "109000"
	DataNotFoundCode             string = "201000"
	DataNotFoundCodeInTDAASystem string = "801"
	SQLException                 string = "995"
	AuthorizeInCorrect           string = "996"
	TokenIsExpired               string = "997"
	ConditionNotFound            string = "998"
	ErrorException               string = "999"
	DataReadAuthorizeNotFound    string = "201000"
	DataWriteAuthorizeNotFound   string = "201000"
	DataCxTransDefNotFound       string = "201000"
	ApplicationError             string = "400000"
	DataProductCRMNotFound       string = "201000"
	DataInvalid                  string = "203000"
	UpSertCampSubmitDone         string = "204000"
)

const (
	Success        string = "Success"
	DataNotFoundEn string = "Data not found"
	ContactAdminEn string = "Please contact admin."
)

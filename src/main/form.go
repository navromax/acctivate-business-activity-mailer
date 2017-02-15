package main

import (
	"mime/multipart"
	"os"
	"bufio"
	"io/ioutil"
)

type HttpForm struct {
	Name         string `form:"name"`
	EMail        string `form:"email";binding:"Email"`
	Company      string `form:"company"`
	ActivityCode string `form:"ActivityCode"`
	ProdName     string `form:"prodname"` 
	SerNo	     string `form:"serno"`
	Message      string `form:"message"`
	
	File         *multipart.FileHeader `form:"file"`
}

type BusinessActivity struct {
	ActivityType   string
	ActivityCode   string
	ActivityDescription string
	AssignedTo     string
	Reference      string
	Reference2     string
	OpenedBy       string
	Contact        string
	Email          string
	AddressName    string
	Discussion     string
	Description    string
	AttachmentPath *string
  AttachmentName *string
}

var activityCodes = map[string]string{
	"COM": "Complaint",
	"IR": "Information Request",
}

func BuildBusinessActivity(form HttpForm) (*BusinessActivity, error) {
	ba := BusinessActivity{}

	// Select which business activity type to use.
	// This is the internal code found in the Business Activities section
	// of Configuration Management in ACCTivate!.
	ba.ActivityType = "WEB"

	// Use the business activity Code that the user selected on the web form
	ba.ActivityCode = form.ActivityCode
	ba.ActivityDescription = activityCodes[form.ActivityCode]

	// Leave this business activity unassigned
	ba.AssignedTo = ""

	// Fill in the contact information on the business activity
	ba.Contact = form.Name
	ba.Email = form.EMail

	// Fill in Product Name and Serial/Lot Number
	ba.Reference = form.ProdName
	ba.Reference2 = form.SerNo

	// Fill in the company name and address for this request
	ba.AddressName = form.Company

	// Specify which ACCTivate! user will be used for the "Opened By" field on the business activity.
	ba.OpenedBy = "GST"

	// Set business activity description to include the contact name
	ba.Description = "Service Request from " + ba.Contact

	// Add the company name to the description if it was supplied
	if ba.AddressName != "" {
		ba.Description = ba.Description + " (" + ba.AddressName + ")"
	}

	// Set discussion to the message entered by the user
	ba.Discussion = form.Message

	// Handle uploaded file
	if form.File != nil {
		if uploadedFile, err := form.File.Open(); err != nil {
			return nil, err
		} else {
			if f, err := ioutil.TempFile(os.TempDir(), "-upload-"); err != nil {
				return nil, err
			} else {
				w := bufio.NewWriter(f)
				if _, err := bufio.NewReader(uploadedFile).WriteTo(w); err != nil {
					return nil, err
				} else {
          attachmentPath := f.Name()
          attachmentName := form.File.Filename
          ba.AttachmentPath = &attachmentPath
          ba.AttachmentName = &attachmentName
				}
			}
		}
	}

	return &ba, nil
}

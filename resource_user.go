package triton

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/joyent/triton-go/identity"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create:   resourceUserCreate,
		Read:     resourceUserRead,
		Update:   resourceUserUpdate,
		Delete:   resourceUserDelete,
		Exists:   resourceUserExists,
		Timeouts: fastResourceTimeout,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Description: "Users E-mail",
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},
			"login": {
				Description: "User Login",
				Required:    true,
				Type:        schema.TypeString,
			},
			"password": {
				Description: "password to login",
				Required:    true,
				Type:        schema.TypeString,
			},
			"companyname": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"firstname": {
				Description: "Users First Name",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"lastname": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"address": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"postalcode": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"city": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"state": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"country": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"phone": {
				Description: "Name of Company",
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	i, err := client.Identity()
	if err != nil {
		return err
	}

	user, err := i.Users().Create(context.Background(), &identity.CreateUserInput{
		Email:       d.Get("email").(string),
		Login:       d.Get("login").(string),
		CompanyName: d.Get("companyname").(string),
		FirstName:   d.Get("firstname").(string),
		LastName:    d.Get("lastname").(string),
		Address:     d.Get("address").(string),
		PostalCode:  d.Get("postalcode").(string),
		City:        d.Get("city").(string),
		State:       d.Get("state").(string),
		Country:     d.Get("country").(string),
		Phone:       d.Get("phone").(string),
		Password:    d.Get("password").(string),
	})

	if err != nil {
		return err
	}
	d.SetId(user.ID)

	return resourceUserRead(d, meta)

}
func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	i, err := client.Identity()
	if err != nil {
		return err
	}

	user, err := i.Users().Get(context.Background(), &identity.GetUserInput{
		UserID: d.Id(),
	})
	if err != nil {
		return nil
	}

	d.Set("login", user.Login)
	d.Set("email", user.EmailAddress)
	d.Set("companyname", user.CompanyName)
	d.Set("firstname", user.FirstName)
	d.Set("lastname", user.LastName)
	d.Set("address", user.Address)
	d.Set("postalcode", user.PostalCode)
	d.Set("city", user.City)
	d.Set("state", user.State)
	d.Set("country", user.Country)
	d.Set("phone", user.Phone)
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	i, err := client.Identity()
	if err != nil {
		return err
	}
	_, err = i.Users().Update(context.Background(), &identity.UpdateUserInput{
		UserID:      d.Id(),
		Email:       d.Get("email").(string),
		Login:       d.Get("login").(string),
		CompanyName: d.Get("companyname").(string),
		FirstName:   d.Get("firstname").(string),
		LastName:    d.Get("lastname").(string),
		Address:     d.Get("address").(string),
		PostalCode:  d.Get("postalcode").(string),
		City:        d.Get("city").(string),
		State:       d.Get("state").(string),
		Country:     d.Get("country").(string),
		Phone:       d.Get("phone").(string),
	})
	if err != nil {
		return err
	}

	return nil
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	i, err := client.Identity()
	if err != nil {
		return err
	}
	return i.Users().Delete(context.Background(), &identity.DeleteUserInput{
		UserID: d.Id(),
	})
}
func resourceUserExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*Client)
	i, err := client.Identity()
	if err != nil {
		return false, err
	}

	return resourceExists(i.Users().Get(context.Background(), &identity.GetUserInput{
		UserID: d.Id(),
	}))
}

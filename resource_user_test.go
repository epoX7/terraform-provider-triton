package triton

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/joyent/triton-go/identity"
	"github.com/joyent/triton-go/errors"
)

func TestAccTritonUser_basic *testing.T) {
	config := testAccTritonUser_basic(acctest.RandIntRange(3, 2048))

	resource.Test( t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: testCheckTritonUserDelete,
		Steps: []resource.TestStep{
			{

				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckTritonUserExists("triton_user.test")
				),
			},
		},
		
	})
}
func TestAccTritonUser_update(t. *testing.T) {
	userNumber := acctest.RandIntRange(3, 2048)
	preConfig := TestAccTritonUser_basic(userNumber)
	postConfig := TestAccTritonUser_basic(userNumber)


	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		CheckDestroy: testCheckTritonUserDelete,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckTritonUserExists("triton_user.test"),
					resource.TestCheckResourceAttr("triton_user.test", "email", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "login" , "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "password" , "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "companyname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "firstname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "lastname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "address", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "postalcode", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "city", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "state", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "country", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "phone", "test-user")
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckTritonUserExists("triton_user.test")
					resource.TestCheckResourceAttr("triton_user.test", "email", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "login" , "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "password" , "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "companyname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "firstname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "lastname", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "address", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "postalcode", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "city", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "state", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "country", "test-user")
					resource.TestCheckResourceAttr("triton_user.test", "phone", "test-user")
					
				),
			},
		},
	})

}
//pretty sure I fucked this up but talk to brucie
func testCheckTritonUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		conn := testAccProvider.Meta().(*Client)
		n, err := conn.Identity()
		if err != nil {
			return err
		}
		id, err = resourceUserExists(i.User().Get(context.Background(), &identity.GetUserInput{
			UserID: d.Id(),
		})
		if err != nil {
			return err
		}
		return nil
		})
	}

}


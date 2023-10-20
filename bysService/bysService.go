package byservice 

import "context"

type Byservice struct {
}
// DTO (DATA TRANSFER OBJECT)
type FooReq struct {
	Username string
	Email    string 

}

type SignupResp struct{
	ID string 
}

func (d*bysService) SignUp(ctx context.Context, username , email string ) (*SignupResp, error){
	userName, err := user.NewUsername(username)
	id err != nil{
	  return nil, err
  }
	userEmail, err := user.NewEmail(email)
	 if err != nil {
		return nil, err
  }

  u := user.New(UserEmail,userName)

  // pass to repos

  return &SignupResp{
	ID: u.ID.String(),
  }, nil
}
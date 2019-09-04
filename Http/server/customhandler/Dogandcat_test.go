package customhandler_test

import (
	"github.com/golang/mock/gomock"
	"testing"
	"webservice/Http/server/mocks"
	. "webservice/Http/server/customhandler"
)

func TestPetName(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAnimal := mocks.NewMockAnimal(mockCtrl)

	Human := &Human{Pet:mockAnimal}


	mockAnimal.EXPECT().Myname().Return("xiaohuanhuan").Times(1)
	petname := Human.Mypetname()
	if petname != "xiaohuanhuan"{
		t.Fail()
	}

}
package core

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"isp.accounts.api/core/config"
	"isp.accounts.api/core/models"
	"os"
)

var accounts = "accounts"

func (receiver *Repository) CreateAccount(request *models.CreateAccountModel) (bool, error) {

	var account models.Account
	account.ID = GenerateCode()
	account.Created = Now()
	account.Modified = Now()
	account.CustomerID = request.CustomerId
	account.Status = models.NewStatus(config.ACTIVE)
	results, err := receiver.Database.Collection(accounts).InsertOne(ctx, account)
	if err != nil { return false, err }
	log.Infof("[core.CreateAccount] Account Created: %s", results.InsertedID)

	return true, nil
}

func (receiver *Repository) FetchAccounts() ([]*models.Account, error) {

	collection := receiver.Database.Collection(accounts)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil { log.Error(err) }
	var list []*models.Account
	if err = cursor.All(ctx, &list); err != nil { log.Error(err); return nil, err }

	return list, nil
}

func (receiver *Repository) FetchAccount(id string) (*models.Account, error) {

	filter := bson.M{"id": id}
	var acc models.Account
	collection := receiver.Database.Collection(accounts)

	err := collection.FindOne(ctx, filter).Decode(&acc)

	if err != nil { log.Error(err); return nil, err }

	return &acc, nil
}

func (receiver *Repository) UpdateAccountStatus(id string, name string) bool {

	status := models.NewStatus(name)
	status.Created = Now()
	log.Infof("[core.UpdateAccountStatus]...Going to update an account")
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": status}}
	updatedResult, err := receiver.Database.Collection(accounts).UpdateOne(ctx, filter, update)
	if err != nil { log.Error(err); return false }
	log.Infof("[core.UpdateAccountStatus] %v", updatedResult.ModifiedCount)

	return true
}

/////////////// ACCOUNT MANAGEMENT SECTION ///////////////////

func (receiver *Repository) FundAccount(id string, amount float64) bool {

	account, err := receiver.FetchAccount(id)
	if err != nil {
		log.Errorf("Error Fetch an account: [%v]", err)
		return false
	}
	filter := bson.M{"id": id}
	newAmount := 0.0
	// TODO: PUT A SWITCH
	balance := account.Balance
	var newBalance models.Amount
	newBalance.Currency = os.Getenv("CURRENCY")
	if balance == nil {
		newBalance.Value = models.AddAmount(newAmount, amount)
		newBalance.DisplayName = models.NewDisplayName(amount)
	} else {
		newAmount = models.AddAmount(amount, account.Balance.Value)
		newBalance.Value = newAmount
		newBalance.DisplayName = models.NewDisplayName(newAmount)
	}
	update := bson.M{"$set":bson.M{"balance": newBalance}}
	fundResults, err := receiver.Database.Collection(accounts).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Errorf("[core.FundAccount] Error fund an account [%v]", err)
		return false
	}
	log.Infof("[core.FundAccount] Account Successfully Funded: [%v]", fundResults.ModifiedCount)
	return true
}

func (receiver *Repository) SuspendAccount(id string) bool {

	filter := bson.M{"id": id}
	suspended := models.NewStatus(config.SUSPENDED)
	update := bson.M{"$set": bson.M{"status": suspended}}
	suspendResults, err := receiver.Database.Collection(accounts).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Errorf("[core.SuspendAccount] Error suspending an account [%v]", err)
		return false
	}
	log.Infof("[core.SuspendAccount] Account Successfully Suspended: [%v]", suspendResults.ModifiedCount)
	return true
}

func (receiver *Repository) ActivateAccount(id string) bool {

	filter := bson.M{"id": id}
	activated := models.NewStatus(config.ACTIVE)
	update := bson.M{"$set": bson.M{"status": activated}}
	activationResults, err := receiver.Database.Collection(accounts).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Errorf("[core.ActivateAccount] Error activating an account [%v]", err)
		return false
	}
	log.Infof("[core.ActivateAccount] Account Successfully Activated: [%v]", activationResults.ModifiedCount)
	return true
}

func (receiver *Repository) BlockAccount(id string) bool {

	filter := bson.M{"id": id}
	blocked := models.NewStatus(config.BLOCKED)
	update := bson.M{"$set": bson.M{"status": blocked}}
	blockedResults, err := receiver.Database.Collection(accounts).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Errorf("[core.BlockAccount] Error Blocking an account [%v]", err)
		return false
	}
	log.Infof("[core.BlockAccount] Account Successfully Blocked: [%v]", blockedResults.ModifiedCount)
	return true
}
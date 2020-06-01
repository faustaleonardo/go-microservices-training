package main

type Item struct {
	Id    uint
	Title string
}

var Database []Item

type API int

func (a *API) addItem(item Item, reply *Item) error {
	Database = append(Database, item)
	reply = &item

	return nil
}

func (a *API) editItem(itemEdited Item, reply *Item) error {
	for index, item := range Database {
		if item.Id == itemEdited.Id {
			Database[index] = itemEdited
			reply = &itemEdited
		}
	}

	return nil
}

func (a *API) getItemByID(id uint, reply *Item) error {
	for _, item := range Database {
		if item.Id == id {
			reply = &item
			break
		}
	}

	return nil
}

func main() {}

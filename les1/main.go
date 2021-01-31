package main

import (
	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/pkg/sendmail"
	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/pkg/tgbot"
	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/repository"
	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"awesomeProject/GoArchitecture/gb-go-architecture-master/lesson-2/shop/pkg/gorilla/mux"
)

func main() {
	sendmail.SendMail()
	tg, err := tgbot.NewTelegramAPI("1446832842:AAE55Z-MRFhNmyNdtE1Gnw-vRDppr", 1325926)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	db := repository.NewMapDB()
	service := service.NewService(tg, db)
	handler := &shopHandler{
		service: service,
		db:      db,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

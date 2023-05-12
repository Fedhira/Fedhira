package url

import (
	"github.com/Fedhira/Tugas_1214028/controller"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	page.Get("/presen", controller.GetPresensi)
	page.Get("/nasabah", controller.GetNasabah)
	page.Get("/penagih", controller.GetPenagih)
	page.Get("/tagihan", controller.GetTagihan)
	page.Get("/bank", controller.GetBank)
	page.Get("/test", controller.GetAll)
	page.Post("/insnasabah", controller.InsertNasabah)
	page.Post("/inspenagih", controller.InsertPenagih)
	page.Post("/instagihan", controller.InsertTagihan)
	page.Post("/insbank", controller.InsertBank)
	page.Post("/ins", controller.InsertData)

	page.Get("/presensi", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id

	page.Get("/all", controller.GetAllBank) //menampilkan seluruh data presensi
}

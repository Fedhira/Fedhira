package controller

import(
	inimodel "github.com/Fedhira/be_tagihan/model"
	inimodul "github.com/Fedhira/be_tagihan/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"github.com/Fedhira/Tugas_1214028/config"
	"net/http"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func GetAllNew(c *fiber.Ctx) error {
// 	ps := inimodul.GetAllTagihan(config.Ulbimongoconn, "bank")
// 	return c.JSON(ps)
// }

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/Fedhira/Fedhira",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetNasabah(c *fiber.Ctx) error {
	nl := inimodul.GetNasabahFromNama("Auliyah Safana","nasabah", config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetPenagih(c *fiber.Ctx) error {
	nl := inimodul.GetPenagihFromNama("Marlina", "penagih", config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetTagihan(c *fiber.Ctx) error {
	nl := inimodul.GetTagihanFromNama_nasabah("Auliyah Safana", "tagihan", config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetBank(c *fiber.Ctx) error {
	nl := inimodul.GetBankFromNama_bank("bank abc", "bank", config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllTagihanFromNama_nasabah("Auliyah Safana", config.Ulbimongoconn, "tagihan")
	return c.JSON(nl)
}

func InsertNasabah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var nasabah inimodel.Nasabah
	if err := c.BodyParser(&nasabah); err != nil {
		return err
	}
	insertedID := inimodul.InsertNasabah(db, "nasabah",
		nasabah.Nama_nasabah,
		nasabah.Email,
		nasabah.Phone_number,
		nasabah.Alamat)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertPenagih(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var penagih inimodel.Penagih
	if err := c.BodyParser(&penagih); err != nil {
		return err
	}
	insertedID := inimodul.InsertPenagih(db, "penagih",
	penagih.Nama_penagih,
	penagih.Email,
	penagih.Phone_number,
	penagih.Total_Tagihan)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertTagihan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var tagihan inimodel.Tagihan
	if err := c.BodyParser(&tagihan); err != nil {
		return err
	}
	insertedID := inimodul.InsertTagihan(db, "tagihan",
	tagihan.Total_Tagihan,
	tagihan.Deskripsi,
	tagihan.Status,
	tagihan.Tanggal_jatuhtempo,
	tagihan.Biodata,
	tagihan.Location,
	tagihan.Longitude,
	tagihan.Latitude)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertBank(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var bank inimodel.Bank
	if err := c.BodyParser(&bank); err != nil {
		return err
	}
	insertedID := inimodul.InsertBank(db, "bank",
	bank.Nama_bank,
	bank.Lokasi,
	bank.Daftar)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// Photoshop Jr. is the code used to convert images to 500x500px images, one that retains the original color and the other turning the image into monochrome.
// It also can take a valid PNG file and automatically convert it to the correct name and format.
package photoshopjr

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/response"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"

	"github.com/disintegration/gift"
	"github.com/harrydb/go/img/grayscale"
)

// separateFilenameFromExtension() strips the last 4 characters from a string
func separateFilenameFromExtension(filename string) string {
	strippedfilename := filename[0 : len(filename)-4]
	return strippedfilename
}

// ProcessImage() is the function which is intially called whenever an image is POSTed to a valid endpoint
// It's job is to create the intial file, check to see if the filetype is a valid PNG or JPEG
// Once it confirms it is a valid JPEG or PNG, it calls furtherProcessing() asyncronously (after sending a success response JSON object back to the front-end) because if the image processing gets this far, we can assume the image was valid.
func ProcessImage(w http.ResponseWriter, r *http.Request, color_filename, mono_filename string) {
	file, header, err := r.FormFile("file")
	if err != nil {
		json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Couldn't decode http request"})
		return
	}
	defer file.Close()

	out, err1 := os.Create(utils.FilesDir() + header.Filename)
	if err1 != nil {
		json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Couldn't create file" + header.Filename})
		return
	}
	defer out.Close()
	_, err2 := io.Copy(out, file)
	if err2 != nil {
		json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Couldn't copy image binary as the file " + header.Filename})
		return
	}

	test, errt := os.Open(utils.FilesDir() + header.Filename)
	if errt != nil {
		fmt.Println("Can't open test file")
	}
	_, jpgerr := jpeg.Decode(test)
	_, pngerr := png.Decode(test)
	if jpgerr != nil && pngerr != nil {
		json.NewEncoder(w).Encode(response.Response{Status: false, Message: header.Filename + " is not a valid jpeg/png file. Please upload valid image formats"})
		rmerr := os.Remove(utils.FilesDir() + header.Filename)
		if rmerr != nil {
			fmt.Println("Error deleting file: " + utils.FilesDir() + header.Filename)
		}
		return
	} else {
		json.NewEncoder(w).Encode(response.Response{Status: true, Message: "File uploaded successfully"})
		go furtherProcessing(color_filename, mono_filename, header.Filename)
	}

}

// furtherProcessing() standardizes the photos in the following steps.
//	1) If image is a PNG, convert to JPEG
//	2) Rename image to fit the color_filename argument
//	3) If image is < 500x500px, upsize the image so that the smallest dimention will be exactly 500px, and crop the image to be 500x500. If the image is larger than 500x500px, crop the image to be 500x500px.
func furtherProcessing(color_filename, mono_filename, current_filename string) {

	if current_filename[len(current_filename)-3:] == "jpg" || current_filename[len(current_filename)-3:] == "JPG" || current_filename[len(current_filename)-4:] == "JPEG" || current_filename[len(current_filename)-4:] == "jpeg" {

		err := os.Rename(utils.FilesDir()+current_filename, utils.FilesDir()+color_filename)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {
		replacePNGwithJPEG(color_filename, current_filename)
	}

	cropImage(color_filename)
	MakeMonoChromeImage(color_filename, mono_filename)

}

func replacePNGwithJPEG(color_filename, current_filename string) {
	// fmt.Println("replacePNGwithJPEG")
	file, nexterr := os.Open(utils.FilesDir() + current_filename)
	if nexterr != nil {
		log.Fatal(nexterr)
	}
	defer file.Close()
	img, pngerr := png.Decode(file)
	if pngerr != nil {
		log.Fatal(pngerr)
	}

	outFile, createErr := os.Create(utils.FilesDir() + color_filename)
	if createErr != nil {
		log.Fatal(createErr)
	}
	defer outFile.Close()

	jpegerr := jpeg.Encode(outFile, img, nil)
	if jpegerr != nil {
		log.Fatal(jpegerr)

	}
	rmerr := os.Remove(utils.FilesDir() + current_filename)
	if rmerr != nil {
		log.Fatal(rmerr)
	}
}

func cropImage(filename string) {
	// fmt.Println("cropImage")

	file, err := os.Open(utils.FilesDir() + filename)
	if err != nil {
		fmt.Println("can't open " + filename)
	}
	defer file.Close()

	// fmt.Println("trying to delete")
	// os.Remove(utils.FilesDir() + filename)

	img, err2 := jpeg.Decode(file)
	if err2 != nil {
		fmt.Println("can't decode " + filename + " as a jpeg")
		log.Fatal(err2)
	}

	bounds := img.Bounds()
	width := float32(bounds.Max.X)
	height := float32(bounds.Max.Y)

	if width < 500 || height < 500 {
		// fmt.Println("if width < 500 || height < 500 {")
		// fmt.Println("width", width)
		// fmt.Println("height", height)
		if width < height {
			scale := 500 / width
			// fmt.Println("scale", scale)
			width *= scale
			height *= scale
		} else {
			scale := 500 / height
			// fmt.Println("scale", scale)
			width *= scale
			height *= scale
		}
		// fmt.Println("width", width)
		// fmt.Println("height", height)
		g1 := gift.New(
			gift.Resize(int(width), int(height), gift.LinearResampling),
		)
		dst1 := image.NewRGBA(g1.Bounds(img.Bounds()))
		g1.Draw(dst1, img)
		g2 := gift.New(
			gift.CropToSize(500, 500, gift.CenterAnchor),
		)
		dst2 := image.NewRGBA(g2.Bounds(dst1.Bounds()))
		g2.Draw(dst2, dst1)

		outFile, createErr := os.Create(utils.FilesDir() + filename)
		if createErr != nil {
			fmt.Println("cant save " + filename)
			log.Fatal(createErr)
		}
		defer outFile.Close()

		jpegerr := jpeg.Encode(outFile, dst2, nil)
		if jpegerr != nil {
			fmt.Println("Cant decode jpeg for file " + filename)
			log.Fatal(jpegerr)

		}

	} else {
		g := gift.New(
			gift.CropToSize(500, 500, gift.CenterAnchor),
		)
		dst := image.NewRGBA(g.Bounds(img.Bounds()))
		g.Draw(dst, img)
		outFile, createErr := os.Create(utils.FilesDir() + filename)
		if createErr != nil {
			fmt.Println("cant save " + filename)
			log.Fatal(createErr)
		}
		defer outFile.Close()

		jpegerr := jpeg.Encode(outFile, dst, nil)
		if jpegerr != nil {
			fmt.Println("Cant decode jpeg for file " + filename)
			log.Fatal(jpegerr)

		}
	}

}

func MakeMonoChromeImage(color_filename, mono_filename string) {
	// fmt.Println("MakeMonoChromeImage")
	file, err := os.Open(utils.FilesDir() + color_filename)
	// file, err := os.Open("../../../assets/files/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", utils.FilesDir()+color_filename, err)
	}

	grayImg := grayscale.Convert(img, grayscale.ToGrayLuminance)

	// outFile, err := os.Create("../../../assets/files/" + separateFilenameFromExtension(filename) + ".mono.jpg")
	outFile, err := os.Create(utils.FilesDir() + mono_filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, grayImg, nil)

}

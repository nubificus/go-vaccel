package vaccel

/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func ImageClassificationFromFile(sess *Session, imagePath string) (string, int) {

	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	cImageBytes := (*C.uchar)(&imageBytes[0])
	c_img_buf := unsafe.Pointer(cImageBytes)
	c_img_len := C.ulong(len(imageBytes))

	cText := (*C.uchar)(C.malloc(C.size_t(256)))
	cOutImageName := (*C.uchar)(C.malloc(C.size_t(256)))

	/* Free the memory when done */
	defer C.free(unsafe.Pointer(cText))
	defer C.free(unsafe.Pointer(cOutImageName))

	csess := sess.cSess

	c_ret := C.vaccel_image_classification(
		&csess, c_img_buf, cText, cOutImageName,
		c_img_len, C.ulong(256), C.ulong(256))

	var golangOut string

	if int(c_ret) == 0 {

		ptr := unsafe.Pointer(cText)
		typeCast := (*C.char)(ptr)
		golangOut = C.GoString(typeCast)

	} else {

		golangOut =
			"A problem occurred while running the Operation"
	}

	return golangOut, int(c_ret)
}

func ImageClassification(sess *Session, image []byte) (string, int) {

	cImageBytes := (*C.uchar)(&image[0])
	c_img_buf := unsafe.Pointer(cImageBytes)
	c_img_len := C.ulong(len(image))

	cText := (*C.uchar)(C.malloc(C.size_t(256)))
	cOutImageName := (*C.uchar)(C.malloc(C.size_t(256)))

	/* Free the memory when done */
	defer C.free(unsafe.Pointer(cText))
	defer C.free(unsafe.Pointer(cOutImageName))

	csess := sess.cSess

	c_ret := C.vaccel_image_classification(
		&csess, c_img_buf, cText, cOutImageName,
		c_img_len, C.ulong(256), C.ulong(256))

	var golangOut string

	if int(c_ret) == 0 {

		ptr := unsafe.Pointer(cText)
		typeCast := (*C.char)(ptr)
		golangOut = C.GoString(typeCast)

	} else {

		golangOut =
			"A problem occurred while running the Operation"
	}

	return golangOut, int(c_ret)
}

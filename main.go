package main

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// #import <IOKit/pwr_mgt/IOPMLib.h>
//
// uint32_t take() {
//  CFStringRef reasonForActivity= CFSTR("Describe Activity Type");
//
//  IOPMAssertionID assertionID;
//  IOReturn success = IOPMAssertionCreateWithName(kIOPMAssertionTypeNoDisplaySleep,
//                                     kIOPMAssertionLevelOn, reasonForActivity, &assertionID);
//
//  return success == kIOReturnSuccess ? assertionID : 0;
// }
//
// void release(uint32_t id) {
//  IOPMAssertionRelease(id);
// }
//
import "C"
import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
    id := C.take()

    go func() {
        ch := make(chan os.Signal)
        signal.Notify(ch, os.Interrupt)

         <- ch

        fmt.Println("releasing")
        C.release(id)

        os.Exit(0)
    }()

    select {}
}

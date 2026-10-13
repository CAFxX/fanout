package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int = 32767
    var v2 uint32 = 2147483648
    v3 := []uint32{}
    v4 := []string{"a", "\n\t", "a"}
    v5 := make(map[string]bool)
    var v6 uint64 = 2
    switch "" {
    case "\n\t":
        var v7 int32 = int32(v2)
        _ = v7
    case "hello":
        switch "a" {
        case "a":
            for fi8 := 0; fi8 < 6; fi8++ {
                if false {
                    if false {
                        var v9 int8 = -1
                        _ = v9
                    }
                }
                _ = fi8
            }
        case "\n\t":
            var v10 uint8 = 2
            _ = v10
        default:
            var v11 bool = ((128 > 2) && (("hello" > ((("hello" + "\n\t") + "hello") + "\n\t")) && (0 < (v1 + 677))))
            _ = v11
        }
    case "foo bar":
        v1 -= v1
    }
    var v12 int8 = 127
    var v13 uint64 = 42
    for fi14 := 0; fi14 < 19; fi14++ {
        if ((v2 * (v2 * 4294967295)) != (v2 + (v2 % (((v2 - v2)) | 1)))) {
            var v15 bool = true
            _ = v15
        } else {
            var v16 bool = (true && (!true))
            _ = v16
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        switch (v1 / ((-1) | 1)) {
        case -473:
            _ = len(v3)
        case 42:
            var v18 uint16 = 255
            _ = v18
        case -2147483648:
            var v19 bool = ((!false) || false)
            _ = v19
        }
        _ = fi17
    }
    var v20 bool = (!("" >= "hello"))
    var v21 int = v1
    delete(v5, "\n\t")
    v22 := GenS1{F0: 707, F1: 388}
    _ = len(v5)
    v1 -= (v1 / ((v21) | 1))
    _ = len(v3)
    var v23 int16 = 1
    var v24 string = ("" + "hello")
    v1 += (v21 + 127)
    var v25 uint32 = (v2 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

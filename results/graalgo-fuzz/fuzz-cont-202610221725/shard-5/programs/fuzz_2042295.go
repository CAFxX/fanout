package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0() uint {
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 14; fi5++ {
        if (false || ((!true) && false)) {
            if (-424 > fi5) {
                var v6 uint8 = 128
                fmt.Println(fi5)
                fmt.Println(v4)
                _ = v6
            }
        }
        if false {
            var v7 float32 = 3.14159
            fmt.Println(fi5, v7)
            var v8 uint8 = 128
            _ = v7
            _ = v8
        } else {
            if false {
                v9 := GenS1{F0: 594}
                v10 := []int{0, (fi5 >> ((255) & 63)), fi5, 32768}
                _ = v10
                _ = v9
            }
        }
        v11 := []int{}
        _ = fi5
        _ = v11
    }
    var v12 int16 = 1
    _ = v12
    _ = v4
    return 0
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 2147483647
    var v3 float64 = -2.25
    switch -9223372036854775808 {
    case 704:
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: 42}
            _ = fi13
            _ = v14
        }
    default:
        v15 := GenS2{F0: 82}
        _ = v15
    }
    v16 := []int{}
    v1 -= v1
    v1 = v1
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint64 = (v1 ^ v1)
        fmt.Println(fi17, v18)
        _ = fi17
        _ = v18
    }
    var v19 uint32 = v2
    if ("\n\t" != "a") {
        if false {
            v20 := GenS3{F0: 654, F1: 722}
            _ = v20
        }
    }
    v1 = v1
    var v21 float64 = (float64(0.1) / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    v3 := []uint32{0, 954, 4294967295, 2147483647}
    var v4 bool = false
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint64 = 747
        _ = fi5
        _ = v6
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        v4 = (false && (2147483647 < 1))
    case "a":
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint = 24
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 591, F1: 2, F2: 18446744073709551615}
    if v4 {
        switch 1 {
        case -1:
            v10 := GenS2{F0: 157}
            _ = v10
        case 42:
            _ = len(v2)
        case 2147483647:
            fmt.Println(v4)
        default:
            v11 := GenS1{F0: 984, F1: 255, F2: 777}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS3{F0: 0}
            v4 = (!(v4 || true))
            v4 = (!((432 <= (fi12 % (((fi12 * fi12)) | 1))) && (!v4)))
            _ = fi12
            _ = v13
        }
    }
    var v14 int16 = 1
    var v15 int32 = -999
    if false {
        v14 -= v14
        v16 := GenS3{F0: 128}
        fmt.Println(v4)
        _ = v16
    }
    _ = len(v2)
    for fi17 := 0; fi17 < 14; fi17++ {
        v15 += (v15 & v15)
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    var v1 bool = (!(575 == 88))
    v2 := []uint32{}
    var v3 int16 = 1
    var v4 float64 = (-2.25 - (1e10 + (1.5 * float64((3.14159 * float32((-2.25 * (-2.25 * (((0.1 / 1e10) / (1e10 + ((1.5 * (0.0 / 1.0)) * 3.14159))) + 1.5)))))))))
    v3 += (v3 & v3)
    switch 1 {
    case -2147483648:
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := GenS1{F0: 42, F1: 711, F2: 110}
            v2 = append(v2, 2147483647)
            var v7 int64 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
    case 2147483647:
        v8 := GenS2{F0: 2147483647, F1: 0, F2: 0}
        _ = v8
    }
    v9 := GenS3{F0: -128}
    v3 -= v3
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS4{F0: 818, F1: 610, F2: 0}
        var v12 int64 = -362
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 9223372036854775807 {
    case 127:
        switch ("foo bar" + "hello") {
        case "":
            var v13 int32 = 0
            _ = v13
        }
    }
    if (1 >= 255) {
        v3 = v3
    } else {
        if (v3 >= (v3 - 42)) {
            v14 := GenS4{F0: 540, F1: 42, F2: 4294967295}
            _ = v14
        } else {
            v1 = (v4 != v4)
        }
        var v15 uint8 = 255
        _ = v15
    }
    var v16 float32 = 3.14159
    for fi17 := 0; fi17 < 11; fi17++ {
        v16 -= (v16 * 0.1)
        v16 = (v16 - (((v16 * v16) - 0.0) * v16))
        _ = fi17
    }
    v18 := GenS5{F0: 18446744073709551615, F1: -1, F2: -2147483648}
    v1 = true
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

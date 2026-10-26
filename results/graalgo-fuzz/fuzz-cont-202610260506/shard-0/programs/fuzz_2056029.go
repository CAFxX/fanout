package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0() string {
    if (!("foo bar" == "\n\t")) {
        if (32768 >= 4294967295) {
            var v5 int8 = -9
            _ = v5
        }
    } else {
        switch "a" {
        case "hello":
            v6 := []int{4294967295}
            _ = v6
        case "":
            v7 := GenS1{F0: 0, F1: 32767, F2: -930}
            _ = v7
        case "a":
            var v8 float64 = 3.14159
            _ = v8
        }
        v9 := []int{628}
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int32 = 1
        _ = fi10
        _ = v11
    }
    var v12 uint = 808
    _ = v12
    return "a"
}

func main() {
    var v1 uint8 = 128
    v2 := []uint32{414, 1, 42, 1}
    var v3 string = "foo bar"
    v4 := []int{}
    v1 = (v1 | v1)
    v1 = v1
    var v13 int16 = 737
    fmt.Println(v3, v3)
    if ((0 > -40) || false) {
        fmt.Println(v1, v13, v13)
    }
    v14 := GenS1{F0: v13, F1: (v13 | (^(v13))), F2: 9223372036854775807}
    if (!true) {
        v13 -= v13
        v15 := GenS2{F0: 18446744073709551615, F1: 255}
        _ = v15
    } else {
        v16 := GenS3{F0: 1, F1: 255}
        if len(v4) > 0 {
            v4[0] = 32767
        }
        _ = v16
    }
    var v17 float64 = float64((3.14159 / float32((float64(0.1) + 0.0))))
    var v18 float32 = (float32(v17) * (0.0 * float32(0.1)))
    if (13 > 972) {
        fmt.Println(v17)
        var v19 float64 = 0.0
        v17 = 3.14159
        _ = v19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!false) {
        if true {
            var v20 uint16 = 32767
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v21 int8 = -83
            _ = v20
            _ = v21
        }
        if (-1 != 274) {
            var v22 uint8 = (v1 % ((255) | 1))
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

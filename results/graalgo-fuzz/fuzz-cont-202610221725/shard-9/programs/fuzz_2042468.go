package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int16) bool {
    var v7 int8 = -74
    var v8 int8 = v7
    var v9 bool = (true && (9223372036854775807 <= -9223372036854775808))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    v1 := []int{32767, -961}
    var v2 uint16 = 255
    var v3 int = -2147483648
    var v4 int = v3
    var v5 int32 = 415
    v6 := make(map[uint32]string)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint32 = 2147483647
        v12 := GenS1{F0: 42, F1: "hello", F2: v2}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: -2, F1: int16(v2)}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 string = ""
        var v16 uint8 = 2
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v5, v2, v5)
    fmt.Println(v2, v5, v4)
    if (!(!false)) {
        if len(v1) > 0 {
            v1[0] = (v4 >> ((2) & 63))
        }
        switch v4 {
        case -951:
            v3 = (v3 * 786)
        case -897:
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int8 = -128
                _ = fi17
                _ = v18
            }
        }
    }
    switch (v3 * (v4 << ((4294967295) & 63))) {
    case 2147483647:
        switch (v4 / ((int(v5)) | 1)) {
        case -464:
            v2 -= 255
        case -629:
            v19 := GenS3{F0: (v3 % ((v3) | 1))}
            _ = v19
        }
    case 32768:
        switch "a" {
        case "\n\t":
            var v20 uint8 = 42
            _ = v20
        default:
            var v21 int32 = -2147483648
            _ = v21
        }
    case 32767:
        var v22 string = ""
        _ = v22
    }
    var v23 int8 = 100
    for fi24 := 0; fi24 < 18; fi24++ {
        v1 = append(v1, (fi24 >> ((4294967295) & 63)))
        var v25 uint16 = (v2 | (v2 | (v2 + (v2 - (v2 + v2)))))
        var v26 bool = (!true)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

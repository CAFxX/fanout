package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 int16, p2 string) uint8 {
    v7 := GenS1{F0: 42, F1: -128, F2: 1}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{127, fi8}
        _ = fi8
        _ = v9
    }
    v10 := []int{478}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -115
    v3 := []uint32{508, 4294967295}
    var v4 uint16 = 2
    var v5 float64 = 0.1
    var v6 bool = false
    if (-1 > 127) {
        v6 = v6
        if false {
            v11 := GenS1{F0: 4294967295, F1: 1, F2: 127}
            switch ("a" + ("" + "hello")) {
            case "\n\t":
                var v12 uint32 = 4294967295
                _ = v12
            case "hello":
                v4 += v4
            default:
                v13 := GenS2{F0: 1, F1: 187}
                _ = v13
            }
            _ = v11
        } else {
            _ = len(v1)
        }
        v2 += v2
    }
    var v14 bool = v6
    v15 := GenS2{F0: 42, F1: 4294967295}
    v4 = (v4 + (v4 << ((4294967295) & 63)))
    var v16 int16 = 32767
    v17 := GenS3{F0: 0}
    fmt.Println(v2, v14)
    for fi18 := 0; fi18 < 8; fi18++ {
        v5 += float64((3.14159 - 3.14159))
        var v19 int32 = -949
        _ = fi18
        _ = v19
    }
    if (v14 || (float32(v5) > 3.14159)) {
        v1[""] = 266
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int16 = (v16 % ((v16) | 1))
        v14 = (2 >= 42)
        _ = fi20
        _ = v21
    }
    fmt.Println(v4, v4)
    v6 = v14
    var v22 float32 = ((3.14159 * ((((1e10 - float32(v5)) * (float32(-2.25) * -2.25)) * float32((v5 / 0.1))) / float32(v5))) + -2.25)
    v23 := GenS4{F0: 255, F1: v16}
    v24 := v1[("hello" + "hello")]
    var v25 string = "\n\t"
    var v26 int16 = int16(v2)
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
func genFn0() uint64 {
    var v5 int32 = 172
    var v6 uint16 = 2
    var v7 int = 32767
    switch ("foo bar" + "") {
    case "foo bar":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 uint16, p2 int32) bool {
    var v9 float64 = float64(((1e10 - float32((float64(1.5) + 3.14159))) - 1.5))
    v10 := []int{9223372036854775807}
    p2 = p0
    var v11 string = (("foo bar" + "") + (("\n\t" + "hello") + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (!true)
}

func main() {
    var v1 int32 = 767
    v2 := []uint32{4294967295, 1}
    var v3 int16 = int16(v1)
    v4 := []int{-743, 32767, 128}
    v12 := GenS2{F0: "a", F1: (false || true), F2: genFn0()}
    fmt.Println(v3, v1)
    switch 4294967295 {
    case 582:
        v13 := GenS1{F0: (v3 << ((820) & 63))}
        _ = v13
    }
    var v14 bool = (-128 >= 1)
    if (-453 != 0) {
        _ = len(v4)
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint32 = 989
            _ = fi15
            _ = v16
        }
        _ = len(v4)
    }
    if v14 {
        fmt.Println(v3, v3, v1)
        var v17 int8 = -1
        _ = v17
    } else {
        _ = len(v2)
        v4 = append(v4, 430)
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int = fi18
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v1)
    var v20 int16 = (v3 / (((v3 * v3)) | 1))
    v1 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 0 {
    case -721:
        var v21 uint64 = 600
        _ = v21
    case -901:
        var v22 uint32 = 2
        _ = v22
    case 127:
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                switch "" {
                case "":
                    v14 = v14
                    v1 = v1
                }
                switch "a" {
                case "":
                    v1 += 146
                case "foo bar":
                    var v25 bool = (2 <= 166)
                    _ = v25
                }
                v4 = append(v4, (fi23 >> ((4294967295) & 63)))
                _ = fi24
            }
            var v26 bool = (!v14)
            _ = fi23
            _ = v26
        }
    default:
        v3 += -1
    }
    v4 = append(v4, -373)
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 int16 = 42
        var v29 int8 = 100
        v20 -= (v20 >> ((551) & 63))
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

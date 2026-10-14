package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 uint }

func main() {
    v1 := []int{32768, 127, -264, -2147483648}
    v2 := make(map[string]uint32)
    var v3 int32 = -415
    var v4 uint64 = 0
    if false {
        v4 = v4
        var v5 string = "\n\t"
        _ = v5
    }
    var v6 int32 = v3
    var v7 uint8 = 76
    if false {
        v6 += v3
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(fi8, v3, fi9)
                delete(v2, ("foo bar" + "a"))
                _ = fi9
            }
            var v10 int = fi8
            _ = fi8
            _ = v10
        }
    } else {
        v4 = (v4 | (v4 + v4))
        var v11 uint64 = v4
        _ = v11
    }
    var v12 int8 = 1
    fmt.Println(v7, v6)
    var v13 int = 109
    switch (v13 % ((-9223372036854775808) | 1)) {
    case 4294967295:
        v3 = v3
    case -179:
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 bool = true
                _ = fi15
                _ = v16
            }
            v1 = append(v1, (v13 % ((v13) | 1)))
            for fi17 := 0; fi17 < 4; fi17++ {
                if false {
                    var v18 uint32 = 2147483648
                    v19 := GenS1{F0: 721, F1: 576, F2: 2147483647}
                    _ = v18
                    _ = v19
                } else {
                    v4 += (v4 >> ((4294967295) & 63))
                    switch (("\n\t" + ("foo bar" + "")) + (("a" + "") + "")) {
                    case "foo bar":
                        var v20 int8 = (v12 >> ((0) & 63))
                        _ = v20
                    }
                }
                _ = len(v2)
                _ = len(v2)
                _ = fi17
            }
            _ = fi14
        }
    default:
        var v21 float32 = 3.14159
        _ = v21
    }
    fmt.Println(v7)
    if (v7 != v7) {
        delete(v2, (("foo bar" + "") + ("\n\t" + ("" + ""))))
        v22 := GenS2{F0: true}
        v7 = v7
        _ = v22
    } else {
        v12 -= (^(v12))
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v6, v12)
            v4 += (v4 - (^(v4)))
            _ = fi24
        }
        if true {
            var v25 int32 = (v3 * (v6 + int32(v7)))
            v12 = (v12 << ((4294967295) & 63))
            fmt.Println(v12)
            _ = v25
        } else {
            v26 := GenS3{F0: 1000, F1: 851, F2: 983}
            v13 -= -152
            for fi27 := 0; fi27 < 15; fi27++ {
                if true {
                    v28 := GenS1{F0: (v6 - (-(v3))), F1: 223, F2: -740}
                    fmt.Println(v12)
                    if true {
                        var v29 int32 = v6
                        var v30 int = 1
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
                _ = fi27
            }
            _ = v26
        }
        v6 = v3
        _ = fi23
    }
    var v31 float64 = (float64(float32(float64(0.1))) * 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v6
    _ = v7
}

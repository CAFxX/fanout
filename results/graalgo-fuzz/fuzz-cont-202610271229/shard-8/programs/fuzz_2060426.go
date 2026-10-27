package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{1, -488}
    v3 := make(map[uint32]string)
    v4 := []uint32{1, 65535, 4294967295}
    v5 := make(map[string]bool)
    var v6 int16 = -32768
    if (182 == -442) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v7 := GenS1{F0: 486, F1: 30, F2: 140}
        var v8 int64 = 9223372036854775807
        _ = v7
        _ = v8
    } else {
        _ = len(v5)
        fmt.Println(v6)
        if (false && false) {
            _ = len(v3)
        } else {
            v9 := GenS1{F0: 32768, F1: 100, F2: 119}
            v10 := GenS1{F0: 218, F1: 100, F2: 193}
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int8 = -1
        _ = fi11
        _ = v12
    }
    var v13 float64 = (3.14159 + 1e10)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 string = (("foo bar" + "") + ("a" + "\n\t"))
        fmt.Println(v6, v13)
        v16 := GenS1{F0: (fi14 & 2147483647), F1: 0, F2: 127}
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = float64((((float32(v13) - -2.25) + (float32((v13 - v13)) + float32((v13 / -2.25)))) - float32(v13)))
    if true {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(v6)
                fmt.Println(v13)
                var v19 int16 = (v6 << ((2) & 63))
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        var v20 float32 = float32(float64((1e10 * -2.25)))
        switch ("" + (("foo bar" + "hello") + "")) {
        case "":
            var v21 int = 2147483647
            _ = v21
        case "hello":
            v22 := GenS2{F0: -732}
            _ = v22
        default:
            v6 = v6
        }
        _ = v20
    } else {
        if (0.1 >= v13) {
            fmt.Println(v6, v6)
            v13 -= v13
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 9; fi24++ {
                    delete(v5, "\n\t")
                    v25 := GenS3{F0: 1}
                    if false {
                        v6 = (^(v6))
                        v13 += 3.14159
                        v26 := v1[2147483648]
                        _ = v26
                    }
                    v6 = 592
                    if ("foo bar" != "foo bar") {
                        var v27 int32 = 42
                        v5[("a" + "a")] = true
                        if len(v2) > 0 {
                            v2[0] = fi23
                        }
                        _ = v27
                    } else {
                        var v28 int8 = 1
                        v13 = 0.1
                        _ = v28
                    }
                    v29 := GenS3{F0: 65535}
                    v30 := GenS2{F0: 135}
                    v13 = v13
                    _ = fi24
                    _ = v25
                    _ = v29
                    _ = v30
                }
                _ = fi23
            }
        }
    }
    var v31 float32 = -2.25
    v6 += v6
    var v32 string = "\n\t"
    if len(v2) > 0 {
        v2[0] = 859
    }
    v13 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

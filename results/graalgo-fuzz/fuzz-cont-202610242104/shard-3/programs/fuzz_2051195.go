package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64, p1 uint64, p2 int32) uint32 {
    v5 := []int{42}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint8 = 128
        if len(v5) > 0 {
            v5[0] = fi6
        }
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = -128
            if len(v5) > 0 {
                v5[0] = fi8
            }
            v10 = v10
            _ = fi9
            _ = v10
        }
        v5 = append(v5, (fi8 | (^(fi8))))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 289
}
func genFn1(p0 uint32) uint8 {
    p0 -= (p0 & genFn0(36, 1, -169))
    fmt.Println(p0)
    v11 := []int{-713, 128, -639, 32767}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 uint16 = 32767
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := []string{"hello", (("\n\t" + "foo bar") + "\n\t"), "hello", "hello"}
    v2 := make(map[string]uint32)
    var v3 float64 = 0.1
    var v4 uint16 = 1
    v13 := GenS1{F0: 1}
    v4 = v4
    delete(v2, "")
    fmt.Println(v4)
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint16 = v4
            v17 := GenS1{F0: 1}
            _ = fi15
            _ = v16
            _ = v17
        }
        v3 = v3
        _ = fi14
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint = 65535
        _ = fi18
        _ = v19
    }
    delete(v2, "\n\t")
    v1 = append(v1, "")
    switch 459 {
    case -9223372036854775808:
        var v20 int16 = 1
        _ = v20
    case -976:
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 += 1
            v4 -= v4
            _ = fi21
        }
    default:
        var v22 string = "\n\t"
        _ = v22
    }
    var v23 uint64 = 370
    var v24 int16 = -32768
    var v25 int32 = 42
    if false {
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v4)
            v27 := GenS1{F0: v24}
            if ((0.0 <= ((v3 * float64(float32(v3))) - (0.1 - (-2.25 - v3)))) && (1.5 == float32(((1e10 * ((v3 - (((v3 / ((1.5 / v3) - v3)) + v3) - v3)) + ((1.5 / v3) + 1e10))) * -2.25)))) {
                var v28 float64 = (float64(0.1) / 1e10)
                var v29 uint16 = v4
                fmt.Println(fi26, v23)
                _ = v28
                _ = v29
            } else {
                fmt.Println(v25)
                var v30 int = (fi26 - (^(fi26)))
                var v31 uint16 = v4
                _ = v30
                _ = v31
            }
            _ = fi26
            _ = v27
        }
    }
    var v32 int = -114
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}

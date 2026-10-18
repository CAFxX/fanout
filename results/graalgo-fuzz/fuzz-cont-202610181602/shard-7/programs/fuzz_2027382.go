package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int, p1 uint64) uint32 {
    if (128 >= 255) {
        var v4 uint16 = 255
        fmt.Println(v4, p0, v4)
        _ = v4
    } else {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5, p0, p1)
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 float64 = (((0.0 + (0.1 - (-2.25 / float64(0.1)))) + 0.0) * (float64(0.1) + float64((3.14159 + (float32(3.14159) / (0.0 - 0.1))))))
                _ = fi6
                _ = v7
            }
            if true {
                p0 += (p0 - -850)
            } else {
                var v8 uint64 = p1
                v9 := []int{}
                switch (p0 | p0) {
                case 0:
                    v9 = append(v9, (p0 << ((0) & 63)))
                }
                _ = v8
                _ = v9
            }
            _ = fi5
        }
    }
    var v10 uint8 = 0
    _ = p0
    _ = p1
    _ = v10
    return 2
}
func genFn1(p0 uint32, p1 uint8) uint16 {
    p0 -= p0
    p1 = (^(p1))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint8 = 1
        fmt.Println(p1, p1)
        _ = fi11
        _ = v12
    }
    if (9223372036854775807 == -1) {
        var v13 uint8 = p1
        _ = v13
    } else {
        var v14 uint = 992
        var v15 string = "\n\t"
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 uint32 = 2147483647
    var v3 uint32 = (v2 << ((42) & 63))
    var v16 float32 = ((0.1 - v1) - (float32(1e10) - float32(float64(0.0))))
    v17 := []int{4294967295, -165}
    v3 -= v3
    for fi18 := 0; fi18 < 18; fi18++ {
        if len(v17) > 0 {
            v17[0] = (fi18 * fi18)
        }
        _ = fi18
    }
    if true {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v19 int = -1
        v17 = append(v17, v19)
        _ = v19
    }
    switch "" {
    case "foo bar":
        v20 := GenS1{F0: 4294967295}
        _ = v20
    default:
        var v21 float32 = v1
        _ = v21
    }
    fmt.Println(v1, v2)
    v22 := GenS1{F0: 2}
    var v23 string = "foo bar"
    v23 = ""
    v2 += 4294967295
    var v24 uint8 = 176
    v23 = (((v23 + "foo bar") + v23) + v23)
    v3 += (v2 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}

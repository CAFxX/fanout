package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32) bool {
    v6 := GenS1{F0: (p0 & (p0 - 201))}
    p0 = -2147483648
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(p0)
        p0 -= (p0 % (((p0 - (-(p0)))) | 1))
        var v8 int64 = -9223372036854775808
        _ = fi7
        _ = v8
    }
    v9 := []int{10, 4294967295, -1, 534}
    switch -720 {
    case 128:
        var v10 int8 = 100
        _ = v10
    case 155:
        if ((-2.25 / 1e10) >= (1e10 - (1e10 * ((float32(((1.5 + float64(1.5)) - 0.0)) / 1e10) + (float32((float64(1.5) + 0.1)) / 0.1))))) {
            if true {
                _ = len(v9)
            }
            var v11 float64 = 0.0
            if len(v9) > 0 {
                v9[0] = 127
            }
            _ = v11
        } else {
            p0 -= p0
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    _ = v9
    return (!true)
}
func genFn1(p0 uint32) int8 {
    if false {
        var v12 uint32 = 42
        _ = v12
    }
    if (2 >= 988) {
        v13 := GenS1{F0: -838}
        _ = v13
    }
    p0 -= (p0 + (p0 << ((2) & 63)))
    _ = p0
    return 113
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    v4 := []int{-700}
    var v5 float64 = 1.5
    var v14 string = "\n\t"
    v1 -= v1
    var v15 float64 = ((3.14159 + v5) + 0.0)
    switch "" {
    case "a":
        v1 = v1
    case "":
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = true
            _ = fi16
            _ = v17
        }
    default:
        v1 -= v1
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            var v19 string = "a"
            v15 = ((float64(0.0) * v5) / v15)
            if true {
                var v20 int = (fi18 / (((fi18 - 703)) | 1))
                fmt.Println(v1, v14)
                _ = v20
            }
            _ = v19
        } else {
            fmt.Println(fi18, v5)
        }
        var v21 uint16 = v1
        v15 = 1.5
        _ = fi18
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        switch (fi22 % (((fi22 >> ((1) & 63))) | 1)) {
        case 2147483647:
            var v23 int32 = 147
            _ = v23
        case -509:
            v2["\n\t"] = 86
        case 663:
            v3 = (^(v3))
        }
        var v24 uint32 = 1
        _ = fi22
        _ = v24
    }
    v3 = (v3 >> ((18446744073709551615) & 63))
    v25 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 32767
    v3 := []int{584, 127}
    var v4 int = -674
    var v5 int64 = 1
    var v6 float64 = (3.14159 - 0.1)
    v5 += (v5 & v5)
    var v7 float64 = v6
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int16 = 42
            fmt.Println(v4, v2)
            _ = fi9
            _ = v10
        }
        var v11 int64 = 0
        _ = fi8
        _ = v11
    }
    v1 = v1
    if (1 <= 42) {
        var v12 int64 = (v5 & v5)
        var v13 int16 = 1000
        _ = v12
        _ = v13
    } else {
        v2 = v2
        v5 -= v5
    }
    v3 = append(v3, (v4 << ((v1) & 63)))
    var v14 uint = v1
    switch (^(v4)) {
    case 42:
        var v15 string = "foo bar"
        _ = v15
    case -1:
        var v16 uint16 = 706
        _ = v16
    case 392:
        var v17 uint32 = 2147483648
        _ = v17
    }
    var v18 float64 = v7
    if ((float32(((((3.14159 / v18) + float64((1.5 - 0.1))) * 1e10) + -2.25)) == 1.5) && (!true)) {
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                v4 = (v4 ^ (fi19 - (fi19 | v4)))
                fmt.Println(v18)
                _ = fi20
            }
            v21 := GenS1{F0: -128, F1: 0}
            _ = fi19
            _ = v21
        }
        fmt.Println(v18)
        var v22 uint16 = 2
        _ = v22
    } else {
        switch "a" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            if (v5 != v5) {
                if false {
                    v23 := GenS1{F0: 127, F1: -882}
                    v24 := GenS1{F0: int8(v14), F1: 534}
                    v3 = append(v3, v4)
                    _ = v23
                    _ = v24
                } else {
                    if (-128 >= -1) {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            _ = len(v3)
                            _ = fi25
                        }
                        fmt.Println(v14, v2, v18)
                        _ = len(v3)
                    }
                    var v26 uint64 = 4294967295
                    v4 += v4
                    _ = v26
                }
                if len(v3) > 0 {
                    v3[0] = v4
                }
                v27 := GenS1{F0: -4, F1: v5}
                _ = v27
            }
        case "foo bar":
            fmt.Println(v2, v5, v4)
        }
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        switch -9 {
        case 128:
            fmt.Println(v4, v6, v4)
        case 32767:
            _ = len(v3)
        default:
            v29 := GenS2{F0: 4294967295, F1: 127, F2: 121}
            _ = v29
        }
        _ = fi28
    }
    var v30 int8 = 0
    if false {
        v2 = (-(v2))
        v1 = (v1 * (v14 / (((^(v14))) | 1)))
        var v31 string = "a"
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint64 {
    v6 := GenS1{F0: ""}
    var v7 float32 = 3.14159
    v7 = v7
    _ = v6
    _ = v7
    return 933
}
func genFn1(p0 uint, p1 bool) uint8 {
    var v8 uint16 = 940
    v9 := []int{-2147483648, 127, 42, 33}
    if p1 {
        var v10 uint8 = 127
        _ = len(v9)
        fmt.Println(v10, p0, v8)
        _ = v10
    } else {
        switch 0 {
        case 32767:
            p0 += 489
        case 426:
            v9 = append(v9, int(v8))
        case 9223372036854775807:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    var v3 float32 = float32(3.14159)
    v4 := make(map[int]bool)
    var v5 int64 = 110
    v5 = (v5 | v5)
    fmt.Println(v3, v2, v5)
    v11 := GenS1{F0: ""}
    var v12 uint32 = 2
    var v13 uint64 = 4294967295
    v14 := []int{128, 614, -21, 0}
    if (((0.1 / 1e10) - 0.0) != -2.25) {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint = 650
            _ = fi15
            _ = v16
        }
        if (genFn0() == (v13 & v13)) {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v14 = append(v14, 42)
            var v17 int64 = v5
            _ = v17
        }
        if (-1 > -603) {
            v1 += v3
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                delete(v4, fi18)
                _ = fi18
            }
            var v19 uint = 18446744073709551615
            _ = v19
        }
    } else {
        var v20 float32 = v1
        var v21 int = 128
        if (false && (-1 > 0)) {
            var v22 float64 = 0.1
            for fi23 := 0; fi23 < 9; fi23++ {
                v21 -= 1
                if (v12 >= 498) {
                    v20 -= -2.25
                }
                switch fi23 {
                case 972:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v5 -= v5
                        v25 := GenS1{F0: "a"}
                        var v26 int8 = -128
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                default:
                    v2 = ((v22 > 0.0) || (false && (v22 > v22)))
                }
                _ = fi23
            }
            fmt.Println(v3)
            _ = v22
        }
        _ = v20
        _ = v21
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2)
    }
    if (!(true && (32767 == 681))) {
        v2 = true
        v12 = v12
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

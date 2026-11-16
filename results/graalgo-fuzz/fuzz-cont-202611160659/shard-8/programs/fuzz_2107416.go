package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int8) int16 {
    p0 -= p0
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 float32 = 1.5
        var v7 uint32 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint64, p1 string, p2 uint32) int {
    var v8 int16 = -1
    var v9 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 827
}

func main() {
    var v1 uint16 = 1
    v2 := []int{-1, 127}
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]uint32)
    switch -9223372036854775808 {
    case 9223372036854775807:
        v1 -= v1
    case -792:
        v1 = (v1 + (v1 | 2))
    }
    var v10 int64 = (v3 * v3)
    fmt.Println(v3, v3)
    if ((0.0 + (float32((((float64(((1e10 - 1e10) * 0.1)) / 1.5) / float64(-2.25)) + 1e10)) + 0.1)) >= 3.14159) {
        v11 := GenS1{F0: 0}
        v10 = v10
        fmt.Println(v10)
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v1, v10, v10)
        var v13 int32 = 2147483647
        fmt.Println(v10)
        _ = fi12
        _ = v13
    }
    var v14 int = -9223372036854775808
    v1 -= v1
    fmt.Println(v14, v14, v3)
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            v14 = 128
            v17 := GenS1{F0: 63}
            _ = fi16
            _ = v17
        }
        switch "foo bar" {
        case "foo bar":
            if false {
                if true {
                    if (((^(v1)) != (^(v1))) || (false || (3.14159 >= float32(-2.25)))) {
                        v14 += (v14 / ((v14) | 1))
                        var v18 float32 = 1e10
                        _ = v18
                    }
                }
            } else {
                v1 -= v1
                delete(v4, 0)
            }
        }
        switch fi15 {
        case 76:
            v4[42] = 2147483648
        }
        _ = fi15
    }
    v19 := GenS1{F0: 18446744073709551615}
    if (3 != 100) {
        var v20 uint64 = 2
        _ = v20
    }
    var v21 uint32 = 2147483647
    var v22 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16) uint32 {
    var v7 float64 = float64(float32(1.5))
    p1 -= (^(p1))
    switch -615 {
    case 0:
        var v8 int16 = 42
        _ = v8
    case 127:
        var v9 uint8 = 2
        _ = v9
    case 32768:
        v10 := []int{-253, 2147483647, 4294967295, 9223372036854775807}
        _ = v10
    default:
        var v11 float64 = v7
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := []int{-1, (-(fi12)), fi12}
        if len(v13) > 0 {
            v13[0] = (^(fi12))
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]uint32)
    v3 := []int{128, 128}
    v4 := []string{"foo bar", "", "", "hello"}
    var v5 uint = 396
    var v6 uint16 = 785
    fmt.Println(v5, v5)
    v14 := GenS1{F0: -128}
    _ = len(v4)
    switch -949 {
    case -451:
        for fi15 := 0; fi15 < 5; fi15++ {
            v5 -= (v5 / (((v5 / ((18446744073709551615) | 1))) | 1))
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint64 = 42
                if false {
                    v5 -= v5
                    var v18 uint = v5
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    default:
        switch 32768 {
        case 1:
            switch "" {
            case "\n\t":
                v1 += v1
            }
        }
    }
    var v19 int = 1
    var v20 int16 = 1000
    v6 = v6
    var v21 uint = v5
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int = (v19 / ((v19) | 1))
        fmt.Println(v21)
        if (v21 >= v5) {
            _ = len(v2)
            var v24 uint32 = 1
            v21 += v21
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

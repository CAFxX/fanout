package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 0}
    v10 := GenS1{F0: 2147483648, F1: 2}
    var v11 uint8 = 32
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]string)
    var v3 string = ("hello" + "\n\t")
    v4 := make(map[int]string)
    var v5 uint64 = 687
    var v6 float64 = -2.25
    v6 += ((((1.5 / -2.25) / (v6 - ((v6 / v6) * v6))) + v6) + v6)
    fmt.Println(v1, v6, v3)
    v12 := GenS1{F0: (v1 / ((2) | 1)), F1: 42}
    var v13 int64 = -1
    v3 = "a"
    v6 = (v6 + -2.25)
    v5 -= v5
    var v14 int64 = v13
    switch v3 {
    case "\n\t":
        fmt.Println(v14, v13)
    }
    if (!(1 == 42)) {
        switch -9223372036854775808 {
        case 32768:
            switch 128 {
            case 32767:
                v15 := []int{}
                _ = v15
            case 4294967295:
                for fi16 := 0; fi16 < 6; fi16++ {
                    fmt.Println(fi16, v13)
                    _ = fi16
                }
            default:
                var v17 int32 = 2147483647
                _ = v17
            }
        case -534:
            v18 := GenS2{F0: -225, F1: 1000}
            _ = v18
        case 9223372036854775807:
            v5 += 785
        default:
            var v19 string = "\n\t"
            _ = v19
        }
        v3 = ((v3 + v3) + ("foo bar" + "hello"))
    } else {
        var v20 float32 = -2.25
        v13 = (v14 * (-(v13)))
        v13 += (-(v13))
        _ = v20
    }
    switch -924 {
    case 169:
        for fi21 := 0; fi21 < 7; fi21++ {
            v14 = (v14 << ((255) & 63))
            var v22 uint = 812
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint = v22
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    case -935:
        if (71 != uint8(v1)) {
            v1 = 2
        }
    case 76:
        v25 := []int{}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

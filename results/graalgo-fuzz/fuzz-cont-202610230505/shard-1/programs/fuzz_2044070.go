package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0() bool {
    var v6 int = 2147483647
    v6 += -2147483648
    v6 = v6
    var v7 uint64 = 1
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int, p1 int16, p2 uint8) int {
    v8 := []int{}
    p0 = p0
    fmt.Println(p0, p2)
    if (true || ((p1 << ((339) & 63)) != p1)) {
        var v9 int8 = -1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 float64 = (-2.25 * 0.0)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    var v5 bool = true
    v10 := []int{}
    if false {
        for fi11 := 0; fi11 < 1; fi11++ {
            switch fi11 {
            case 32767:
                v12 := GenS1{F0: 4294967295, F1: (v5 || (false && v5)), F2: -128}
                _ = v12
            case 9223372036854775807:
                v13 := GenS1{F0: 4294967295, F1: true, F2: -1}
                _ = v13
            }
            _ = fi11
        }
        fmt.Println(v4, v4)
        var v14 uint = 2
        _ = v14
    }
    v15 := GenS1{F0: 2147483647, F1: (100 == -128), F2: 103}
    v16 := GenS2{F0: 979, F1: 1, F2: -1}
    v4 = (v4 + ("" + v4))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v17 uint64 = 42
    var v18 uint16 = 32767
    if len(v10) > 0 {
        v10[0] = genFn1(-1, 32767, 127)
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v4 = v4
        if ((1 < 2147483648) || true) {
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int = genFn1(0, 32767, 92)
                v4 = "a"
                var v22 float64 = (0.0 * (3.14159 - v1))
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = len(v10)
        } else {
            v18 -= (v18 >> ((4294967295) & 63))
        }
        if ((v1 <= float64(1.5)) || (100 > 0)) {
            switch (v4 + "\n\t") {
            case "a":
                fmt.Println(fi19)
            default:
                v4 = v4
            }
        } else {
            fmt.Println(v5)
            if len(v10) > 0 {
                v10[0] = -482
            }
            for fi23 := 0; fi23 < 12; fi23++ {
                switch (-(fi19)) {
                case -786:
                    v24 := v3[-2147483648]
                    _ = v24
                }
                v25 := GenS2{F0: 1, F1: -754, F2: 0}
                v18 += v18
                v18 = (v18 % ((v18) | 1))
                v18 = v18
                _ = fi23
                _ = v25
            }
        }
        _ = fi19
    }
    v26 := GenS3{F0: 128, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

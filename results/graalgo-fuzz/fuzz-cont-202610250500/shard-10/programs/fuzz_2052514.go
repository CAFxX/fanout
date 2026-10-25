package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 bool, p2 int64) uint {
    if true {
        var v6 uint16 = 0
        _ = v6
    } else {
        if (true && (0 <= 42)) {
            switch 32767 {
            case -1:
                if p1 {
                    p1 = true
                    var v7 int32 = -844
                    _ = v7
                }
            case 1:
                switch 2147483647 {
                case 9223372036854775807:
                    var v8 int = 32767
                    _ = v8
                default:
                    v9 := GenS1{F0: (^(p2)), F1: -9223372036854775808, F2: 0}
                    _ = v9
                }
            case -263:
                p0 = p0
            }
        }
        var v10 int16 = 972
        p2 = (-(p2))
        _ = v10
    }
    p1 = (-44 != 1)
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 uint64) int32 {
    fmt.Println(p0)
    var v11 int16 = 222
    p0 += p0
    var v12 uint16 = 32767
    _ = p0
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    var v1 int32 = 725
    var v2 int = -600
    var v3 int = -2147483648
    v4 := []int{32768, 248}
    var v5 int = -9223372036854775808
    fmt.Println(v1, v1)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int16 = -1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v5 = -1
        _ = fi15
    }
    var v16 uint = 605
    switch "a" {
    case "a":
        v17 := GenS1{F0: 996, F1: 56, F2: (^(v1))}
        _ = v17
    }
    var v18 uint16 = 32767
    v3 += v5
    if (750 > 4294967295) {
        fmt.Println(v16, v18, v2)
        fmt.Println(v5)
        switch (v5 + v5) {
        case -2147483648:
            if (!true) {
                var v19 int32 = 2147483647
                if (!(v16 > v16)) {
                    v20 := GenS1{F0: -611, F1: -559, F2: (v19 & v19)}
                    _ = v20
                }
                var v21 bool = (92 >= 1)
                _ = v19
                _ = v21
            }
        case 106:
            v5 = (v3 << ((genFn0("", (32767 != -136), 0)) & 63))
        case -468:
            var v22 bool = (!(2147483648 > 0))
            _ = v22
        default:
            v1 -= v1
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v16 += (v16 * (^(v16)))
        var v24 float64 = (0.0 + (-2.25 * (1e10 / 1.0)))
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, v5, v5)
        case "":
            var v25 uint = 603
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    v3 -= -1
    v4 = append(v4, v5)
    if (!false) {
        var v26 int8 = -1
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS2{F0: (^(v16)), F1: 4294967295}
        for fi29 := 0; fi29 < 2; fi29++ {
            if ((!false) || (v18 == v18)) {
                if len(v4) > 0 {
                    v4[0] = (v5 ^ fi27)
                }
                if true {
                    if (1 <= -26) {
                        v5 = -132
                    }
                } else {
                    v30 := GenS1{F0: 144, F1: v3, F2: 1}
                    switch -1 {
                    case 128:
                        v31 := GenS3{F0: ("foo bar" + "")}
                        _ = v31
                    case 865:
                        var v32 float32 = (((1e10 - 0.0) + (0.1 * 0.1)) / (1.5 / ((-2.25 + ((3.14159 * (0.0 + 0.0)) - (1.5 * 0.0))) * -2.25)))
                        _ = v32
                    }
                    if true {
                        v33 := GenS1{F0: 1, F1: v5, F2: v1}
                        _ = v33
                    } else {
                        if false {
                            v1 = (v1 * genFn1(786))
                        }
                        var v34 int = v3
                        _ = v34
                    }
                    _ = v30
                }
            }
            _ = fi29
        }
        v18 = v18
        _ = fi27
        _ = v28
    }
    var v35 int32 = v1
    switch v5 {
    case 4294967295:
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 float64 = (float64(float32(0.1)) * -2.25)
            _ = fi36
            _ = v37
        }
    case -9223372036854775808:
        fmt.Println(v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}

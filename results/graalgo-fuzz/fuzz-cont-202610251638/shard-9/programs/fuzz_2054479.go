package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 int32, p2 int8) uint64 {
    var v4 int8 = (p2 / ((p0) | 1))
    v5 := GenS1{F0: (p2 != v4)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 562
}

func main() {
    var v1 uint8 = 127
    var v2 float64 = ((-2.25 / 1.0) * -2.25)
    var v3 uint8 = (v1 + 128)
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(v3)
        fmt.Println(fi6)
        v3 += (v3 % (((v3 ^ (^(v1)))) | 1))
        _ = fi6
    }
    switch -9223372036854775808 {
    case 127:
        v7 := []int{2147483647, 128}
        _ = v7
    default:
        v2 = (((3.14159 + 3.14159) + ((v2 + 3.14159) - v2)) / 1e10)
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        if (true && false) {
            v2 = (v2 - (0.0 / (((0.1 - v2) - 0.0) + -2.25)))
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint64 = 2
            _ = fi9
            _ = v10
        }
        v1 = 255
        _ = fi8
    }
    switch "\n\t" {
    case "foo bar":
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: false, F1: 0, F2: 0}
            _ = fi11
            _ = v12
        }
    case "\n\t":
        if true {
            v13 := []int{127, 127, 2147483647}
            fmt.Println(v3, v3)
            var v14 int = 42
            _ = v13
            _ = v14
        }
    }
    var v15 uint32 = 2147483648
    if false {
        var v16 int8 = 42
        _ = v16
    } else {
        var v17 uint8 = v3
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := []int{}
        var v20 uint64 = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v1 += v3
    case 32768:
        v1 = 183
    }
    v2 += v2
    v21 := []int{485}
    v22 := GenS2{F0: true, F1: -1, F2: 1}
    var v23 int64 = -671
    if true {
        if true {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 uint16 = 283
                switch 32768 {
                case 32768:
                    v15 = v15
                case 2147483647:
                    v21 = append(v21, 1)
                case -779:
                    if true {
                        var v26 uint = 4294967295
                        fmt.Println(v15, v26, v26)
                        if len(v21) > 0 {
                            v21[0] = 303
                        }
                        _ = v26
                    }
                default:
                    var v27 float64 = 3.14159
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        } else {
            v28 := GenS1{F0: (-2147483648 == 42)}
            var v29 int16 = -32768
            _ = v28
            _ = v29
        }
    } else {
        if len(v21) > 0 {
            v21[0] = 42
        }
    }
    var v30 uint8 = (v1 * v1)
    var v31 string = "\n\t"
    v32 := GenS3{F0: -1, F1: ("" + v31), F2: 1}
    if ((v23 != v23) && false) {
        switch v31 {
        case "":
            if ((("foo bar" + v31) >= v31) || true) {
                v3 = 128
                if false {
                    var v33 uint = 145
                    if len(v21) > 0 {
                        v21[0] = -2147483648
                    }
                    _ = v33
                }
            }
        }
        _ = len(v21)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}

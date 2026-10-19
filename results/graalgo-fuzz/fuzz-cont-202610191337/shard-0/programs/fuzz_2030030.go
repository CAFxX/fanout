package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 int16, p1 uint16, p2 int) int32 {
    switch p2 {
    case -1:
        v4 := GenS1{F0: 511}
        _ = v4
    case -2147483648:
        p2 -= -9223372036854775808
    default:
        v5 := []int{}
        _ = v5
    }
    p0 += 280
    var v6 int16 = 42
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    if false {
        v1[(("hello" + ("hello" + ("\n\t" + "foo bar"))) + ("\n\t" + "a"))] = 0
        var v7 float32 = (1e10 * 0.0)
        _ = v7
    }
    if (!(false || true)) {
        v2[2147483647] = 42
        var v8 float32 = 3.14159
        _ = v8
    }
    if true {
        var v9 float64 = -2.25
        v9 += 3.14159
        if false {
            var v10 uint64 = 612
            if ((v9 - v9) != 3.14159) {
                fmt.Println(v10)
            } else {
                v11 := GenS1{F0: 1}
                if true {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        var v12 uint8 = 0
                        _ = v12
                    }
                    switch "\n\t" {
                    case "":
                        switch 32768 {
                        case -98:
                            v10 = (v10 % ((18446744073709551615) | 1))
                        case -2147483648:
                            var v13 bool = true
                            _ = v13
                        default:
                            var v14 bool = true
                            _ = v14
                        }
                    }
                }
                _ = v11
            }
            _ = v10
        } else {
            if (true || (!(-128 <= 127))) {
                fmt.Println(v9)
            }
        }
        _ = v9
    }
    var v15 int8 = 127
    if (!false) {
        v16 := GenS1{F0: 1}
        switch 1 {
        case 0:
            var v17 int32 = 2147483647
            _ = v17
        default:
            v18 := v3[970]
            _ = v18
        }
        if false {
            v19 := GenS2{F0: 616, F1: 472}
            if (false && false) {
                fmt.Println(v15)
            }
            v15 -= (v15 | 65)
            _ = v19
        }
        _ = v16
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint = 65535
        switch fi20 {
        case 128:
            if true {
                fmt.Println(fi20, v15)
            } else {
                v15 = v15
                if true {
                    _ = len(v1)
                    switch "hello" {
                    case "":
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v21 -= v21
                            _ = fi22
                        }
                    case "\n\t":
                        v15 = (v15 * (v15 - v15))
                    }
                    v21 += (v21 | (v21 | (v21 - v21)))
                }
                v15 = (v15 & (v15 | (v15 ^ (v15 + (v15 & v15)))))
            }
        case 127:
            fmt.Println(v15)
        case 2147483647:
            var v23 int = fi20
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch ("" + "foo bar") {
    case "foo bar":
        fmt.Println(v15)
    case "a":
        v24 := GenS2{F0: 1, F1: 32768}
        _ = v24
    case "":
        v15 -= v15
    default:
        v15 = (-(v15))
    }
    var v25 float64 = 1e10
    fmt.Println(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
}

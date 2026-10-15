package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int32) int64 {
    p0 -= p0
    p0 = p0
    _ = p0
    return -1
}
func genFn1(p0 int64, p1 string) uint16 {
    switch -206 {
    case -341:
        p1 = p1
    case -2147483648:
        switch p1 {
        case "hello":
            p1 = "hello"
        case "\n\t":
            var v6 uint64 = 288
            _ = v6
        }
    }
    switch int(p0) {
    case -883:
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                p1 = ""
                v9 := GenS1{F0: 4294967295, F1: (p0 - p0)}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 42, F1: p0}
            _ = fi7
            _ = v10
        }
    case -1:
        v11 := []int{}
        _ = v11
    case -234:
        v12 := GenS1{F0: 42, F1: 42}
        _ = v12
    default:
        p1 = ("hello" + (p1 + p1))
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 2
    var v3 float64 = 3.14159
    var v4 uint32 = 1
    var v5 uint = v2
    v13 := []int{-9223372036854775808, -2147483648}
    fmt.Println(v3)
    var v14 uint64 = 4294967295
    if (!(-2.25 < (v3 + v3))) {
        v5 -= uint(v14)
        switch 1 {
        case 765:
            var v15 float64 = 0.0
            _ = v15
        default:
            v13 = append(v13, -2147483648)
        }
    }
    var v16 uint8 = 127
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            v3 = v3
            _ = len(v13)
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi18
        }
        var v19 int = (fi17 ^ fi17)
        var v20 bool = false
        _ = fi17
        _ = v19
        _ = v20
    }
    var v21 uint32 = 42
    if true {
        v3 = (((v3 * v3) - -2.25) - (v3 * v3))
        var v22 int64 = 42
        var v23 uint64 = 343
        _ = v22
        _ = v23
    }
    var v24 int16 = v1
    var v25 float64 = 3.14159
    v13 = append(v13, 9223372036854775807)
    var v26 uint32 = 65535
    if ((v2 < v2) || (!false)) {
        v16 = v16
        var v27 float64 = v3
        _ = v27
    }
    if len(v13) > 0 {
        v13[0] = 700
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

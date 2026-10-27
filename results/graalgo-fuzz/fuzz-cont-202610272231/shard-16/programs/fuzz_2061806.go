package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 int8, p1 uint16, p2 bool) int8 {
    v6 := GenS1{F0: 32767, F1: 18446744073709551615}
    var v7 int8 = 127
    if (629 == 1) {
        fmt.Println(p1, p2, p1)
        var v8 uint = 65535
        p0 = (p0 - (p0 >> ((v8) & 63)))
        _ = v8
    } else {
        p1 = 812
    }
    v9 := []int{}
    var v10 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v9
    return v7
}
func genFn1() uint {
    var v11 uint16 = 694
    v11 += (^(v11))
    _ = v11
    return 65535
}

func main() {
    var v1 uint = 18
    var v2 int64 = -1
    v3 := []string{("\n\t" + "a"), ("hello" + "foo bar"), "a"}
    var v4 int8 = 1
    var v5 uint16 = 710
    fmt.Println(v1, v2, v1)
    v5 += (v5 * (v5 / ((v5) | 1)))
    var v12 float64 = 0.0
    v13 := GenS1{F0: -2147483648, F1: 65535}
    if (!true) {
        v12 = 3.14159
    }
    if true {
        if true {
            var v14 int8 = (v4 * v4)
            if true {
                var v15 int = 202
                var v16 uint8 = 127
                _ = v15
                _ = v16
            } else {
                fmt.Println(v1)
            }
            var v17 bool = (!(!false))
            _ = v14
            _ = v17
        }
        v2 = int64(v1)
    }
    fmt.Println(v1, v12, v4)
    for fi18 := 0; fi18 < 17; fi18++ {
        if false {
            var v19 bool = false
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch ("a" + ("hello" + "")) {
            case "":
                var v20 uint64 = 4294967295
                _ = v20
            case "\n\t":
                var v21 uint64 = 0
                _ = v21
            default:
                v22 := GenS1{F0: fi18, F1: (v1 >> ((v1) & 63))}
                _ = v22
            }
            _ = v19
        }
        var v23 int64 = v2
        _ = fi18
        _ = v23
    }
    var v24 float32 = 0.1
    v5 += (v5 ^ v5)
    fmt.Println(v2, v5, v2)
    var v25 uint32 = uint32(v5)
    fmt.Println(v24, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

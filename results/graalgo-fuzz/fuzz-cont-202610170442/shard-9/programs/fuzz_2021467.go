package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int64, p1 uint32, p2 uint32) int32 {
    if (!(-2147483648 <= 9223372036854775807)) {
        p1 = (p2 | p1)
    }
    var v6 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    var v1 bool = true
    var v2 uint8 = 1
    var v3 float32 = (0.0 / (1e10 + (3.14159 / 1e10)))
    var v4 int16 = 567
    var v5 string = ("a" + "hello")
    var v7 int8 = 1
    if (!v1) {
        v2 = (v2 / ((v2) | 1))
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v1)
            _ = fi8
        }
    } else {
        v7 = (v7 ^ v7)
    }
    var v9 float32 = (v3 - v3)
    var v10 float32 = 0.0
    var v11 uint8 = v2
    switch ("" + v5) {
    case "foo bar":
        var v12 uint16 = uint16(v11)
        _ = v12
    case "a":
        v4 -= v4
    case "":
        v13 := GenS1{F0: 64, F1: 388}
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        if true {
            v2 -= v11
        }
        _ = fi14
    }
    if v1 {
        v5 = v5
        if (v4 < -32768) {
            fmt.Println(v3, v11, v11)
            v5 = ("" + v5)
            if (false && (!(-2.25 <= 0.0))) {
                v11 -= (v11 ^ (v2 >> ((18446744073709551615) & 63)))
            } else {
                v9 += (v3 + v10)
                v1 = ((32767 >= 0) && (v7 < (v7 & v7)))
                v7 -= 101
            }
        } else {
            var v15 int64 = int64(v7)
            v16 := GenS1{F0: 1, F1: 0}
            _ = v15
            _ = v16
        }
        var v17 uint32 = 65535
        _ = v17
    }
    var v18 bool = (1.5 > v10)
    if (-9223372036854775808 > -324) {
        v19 := []int{int(v4)}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}

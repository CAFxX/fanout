package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 75
    var v2 uint16 = 2
    var v3 int32 = -1
    var v4 bool = (1.5 <= 0.0)
    var v5 float64 = 0.0
    var v6 string = "foo bar"
    switch 578 {
    case 668:
        v7 := []int{42, 127, 128}
        _ = v7
    case 566:
        fmt.Println(v3, v1, v3)
    default:
        fmt.Println(v5, v4, v1)
    }
    var v8 float32 = (1e10 * 0.0)
    if false {
        var v9 int = -1
        var v10 uint64 = v1
        _ = v10
        _ = v9
    } else {
        if ((!(v8 == (v8 * 0.1))) || (((!v4) || v4) || (413 > 922))) {
            v11 := []int{0, 2147483647}
            _ = v11
        } else {
            v4 = true
            fmt.Println(v6, v3)
            if (9223372036854775807 <= -1) {
                v12 := GenS1{F0: 141}
                fmt.Println(v2)
                if v4 {
                    v5 = v5
                }
                _ = v12
            }
        }
    }
    fmt.Println(v2, v5)
    v13 := []int{}
    v14 := GenS1{F0: 127}
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v1, v8)
    if (false || (4294967295 == 65535)) {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v4)
            if v4 {
                var v19 uint64 = v1
                v20 := GenS1{F0: (v15 % ((2) | 1))}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        var v21 uint8 = 88
        if v4 {
            var v22 uint = 663
            v2 += uint16(v21)
            var v23 int = 2147483647
            _ = v22
            _ = v23
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

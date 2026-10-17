package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint8) uint {
    var v7 bool = true
    v8 := []int{980}
    if false {
        v7 = (v7 && v7)
        var v9 uint32 = 81
        v7 = (-1 >= 0)
        _ = v9
    } else {
        var v10 string = "\n\t"
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, fi11, v7)
        fmt.Println(v7, fi11, v7)
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: -2147483648, F1: 549}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint16, p1 int16, p2 bool) uint32 {
    fmt.Println(p1, p2)
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := []int{fi15, -2147483648}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    if true {
        fmt.Println(p1, p2)
    } else {
        var v17 uint64 = 4294967295
        _ = v17
    }
    if p2 {
        p1 = (p1 | (^(p1)))
        if ((1 >= genFn0(179)) || (412 < 2147483648)) {
            fmt.Println(p0, p0, p2)
            var v18 int = 1
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 32767
    var v3 uint16 = 65535
    var v4 float64 = float64(((0.1 / (0.1 / 1.5)) - 1e10))
    var v5 int16 = (v2 + (^(v2)))
    v6 := make(map[int]bool)
    v1 -= (v1 >> ((4294967295) & 63))
    switch 128 {
    case 127:
        var v19 int64 = -9223372036854775808
        _ = v19
    case -329:
        v4 = v4
    }
    switch "" {
    case "a":
        if ((1 > 100) || false) {
            v3 = (v3 >> ((2) & 63))
        } else {
            switch "" {
            case "\n\t":
                switch 2147483647 {
                case 42:
                    v2 = (v2 + v5)
                }
            default:
                fmt.Println(v4)
            }
            v20 := v6[127]
            v1 -= 210
            _ = v20
        }
    }
    v5 -= v2
    switch int(v3) {
    case 32767:
        v2 -= (v2 | (v2 + (v5 * (v2 * (v2 / ((1000) | 1))))))
    case -166:
        fmt.Println(v4, v1, v2)
    default:
        var v21 uint16 = (v3 % ((0) | 1))
        _ = v21
    }
    var v22 int = 9223372036854775807
    var v23 uint8 = 128
    fmt.Println(v2, v3, v23)
    switch v22 {
    case 4294967295:
        var v24 string = ""
        _ = v24
    }
    v5 = 992
    var v25 int8 = -17
    var v26 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

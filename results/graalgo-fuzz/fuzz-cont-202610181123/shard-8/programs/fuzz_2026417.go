package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 uint16, p1 bool) uint32 {
    v7 := []int{-180, 2147483647}
    fmt.Println(p0, p1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (-32768 > 651) {
        p1 = p1
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}
func genFn1(p0 bool) uint {
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    var v10 int64 = -1
    _ = p0
    _ = v10
    return 608
}

func main() {
    var v1 bool = (false || (!(3.14159 == (0.0 * 1.5))))
    v2 := make(map[int]uint32)
    var v3 float32 = float32((1.5 - -2.25))
    var v4 uint16 = 695
    var v5 float64 = (-2.25 - 1.5)
    var v6 int64 = -1
    if v1 {
        v3 = 1e10
        v11 := GenS1{F0: (!(v6 < v6)), F1: 42}
        v5 = v5
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := []int{4294967295, (fi12 << ((genFn1(v1)) & 63)), fi12}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        switch "foo bar" {
        case "hello":
            v15 := []int{-9223372036854775808, 42}
            _ = v15
        case "foo bar":
            v6 += (v6 >> ((42) & 63))
        case "":
            var v16 float32 = 1e10
            _ = v16
        default:
            var v17 uint64 = 4294967295
            _ = v17
        }
        fmt.Println(v1)
        _ = fi14
    }
    if (688 >= -2147483648) {
        var v18 float32 = v3
        var v19 int64 = (v6 << ((255) & 63))
        var v20 float32 = v18
        _ = v18
        _ = v19
        _ = v20
    }
    if (42 == 46) {
        v3 -= 0.0
    }
    if v1 {
        v21 := GenS1{F0: (!((-333 <= 1000) && false)), F1: -445}
        _ = v21
    }
    var v22 uint8 = 255
    var v23 int = 32768
    for fi24 := 0; fi24 < 14; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            switch 127 {
            case -9223372036854775808:
                v26 := []int{(fi24 >> ((255) & 63)), (fi24 - (v23 * 450))}
                _ = v26
            default:
                fmt.Println(v22)
            }
            fmt.Println(fi25, v23)
            _ = fi25
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS1{F0: false, F1: 0}
            _ = fi28
            _ = v29
        }
        var v30 uint64 = 4294967295
        _ = fi27
        _ = v30
    }
    v22 += 5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

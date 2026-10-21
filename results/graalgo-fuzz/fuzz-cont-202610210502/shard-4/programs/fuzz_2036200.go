package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 int64, p2 int32) uint8 {
    fmt.Println(p1, p0, p0)
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 2; fi8++ {
        p0 -= (p0 >> ((508) & 63))
        v9 := []int{}
        var v10 string = "foo bar"
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int64 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 0
}

func main() {
    var v1 bool = (!(2147483647 == 1))
    var v2 int8 = 127
    var v3 int64 = 9223372036854775807
    var v4 int = 70
    var v5 int8 = (v2 | 127)
    var v6 float32 = (3.14159 + 0.1)
    v4 = v4
    if (!(!v1)) {
        var v12 int8 = 1
        fmt.Println(v3, v12, v5)
        _ = v12
    } else {
        var v13 int8 = (v5 ^ (v2 * (v5 & -128)))
        fmt.Println(v6)
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v2 -= -128
        _ = fi14
    }
    fmt.Println(v2, v5)
    switch (v4 << ((449) & 63)) {
    case 0:
        v2 = (v2 >> ((0) & 63))
    }
    var v15 uint16 = 429
    v1 = v1
    if ((!(v1 && v1)) || true) {
        var v16 uint32 = 783
        v17 := []int{}
        var v18 int8 = 42
        _ = v16
        _ = v17
        _ = v18
    } else {
        v4 -= v4
        var v19 bool = (!(-628 >= -174))
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: 599}
            switch ("\n\t" + "hello") {
            case "a":
                fmt.Println(fi21, v2)
            case "\n\t":
                if (!((1e10 * 1e10) < 1.5)) {
                    v23 := []int{}
                    _ = v23
                } else {
                    v1 = v1
                    v2 -= v2
                    v24 := GenS3{F0: (fi21 >> ((691) & 63)), F1: 0}
                    _ = v24
                }
            default:
                var v25 float64 = 0.0
                _ = v25
            }
            v26 := GenS4{F0: int16(v5)}
            _ = fi21
            _ = v22
            _ = v26
        }
        var v27 string = ("" + "hello")
        _ = fi20
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

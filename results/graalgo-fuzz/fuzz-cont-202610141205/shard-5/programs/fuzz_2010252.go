package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    var v1 bool = (false && false)
    var v2 uint64 = 1
    var v3 uint32 = 2147483648
    var v4 bool = false
    if (v1 || true) {
        v4 = (0 <= 0)
        if (!(v4 || v4)) {
            var v5 uint = 4294967295
            v1 = true
            v6 := GenS1{F0: (v5 % ((2) | 1)), F1: (v1 || true)}
            _ = v5
            _ = v6
        }
    }
    if (((("" + "a") + "foo bar") + ("hello" + "foo bar")) < "a") {
        v1 = (!(!((-2.25 * ((1.5 * -2.25) / (1.5 - (float64(1.5) * (3.14159 - (0.0 * -2.25)))))) >= (3.14159 / float64(0.1)))))
        var v7 int16 = 42
        fmt.Println(v2, v7)
        _ = v7
    } else {
        var v8 uint8 = 75
        var v9 float32 = (1e10 / 1.0)
        v1 = v1
        _ = v8
        _ = v9
    }
    var v10 uint64 = 367
    var v11 uint16 = 42
    var v12 int16 = 1000
    fmt.Println(v12)
    v4 = false
    v13 := []int{1, -2147483648}
    for fi14 := 0; fi14 < 8; fi14++ {
        v11 += uint16(v3)
        _ = fi14
    }
    var v15 float64 = 0.1
    fmt.Println(v4, v12, v10)
    if v4 {
        var v16 uint16 = (v11 | 636)
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v10 = (v2 | v2)
        _ = fi17
    }
    fmt.Println(v3)
    if true {
        fmt.Println(v4)
        v13 = append(v13, -863)
        if (-1 == -1) {
            if (100 == -90) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 uint32 = uint32(v11)
                        v1 = (!v4)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
            var v21 int64 = -919
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v21
        } else {
            v10 -= v2
            var v22 float64 = v15
            _ = v22
        }
    }
    var v23 int32 = -99
    switch -2147483648 {
    case 32767:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    case 0:
        if ((v11 ^ 65535) >= (v11 << ((65535) & 63))) {
            v24 := GenS1{F0: 966, F1: v4}
            fmt.Println(v23, v12, v23)
            _ = v24
        }
    default:
        fmt.Println(v15, v15, v1)
    }
    switch -2147483648 {
    case -406:
        var v25 float64 = v15
        _ = v25
    case -2147483648:
        var v26 uint64 = v2
        _ = v26
    case 2147483647:
        v2 += v2
    default:
        var v27 string = "\n\t"
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}

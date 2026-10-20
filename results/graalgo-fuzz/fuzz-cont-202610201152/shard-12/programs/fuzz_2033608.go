package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
func genFn0(p0 int16, p1 int16, p2 int8) int8 {
    p0 += p1
    var v5 float64 = 1e10
    var v6 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 100
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 2
    var v3 bool = (127 != 128)
    var v4 int = 128
    switch (-(v4)) {
    case 1:
        v7 := []int{-745, (v4 | 1), (v4 + (v4 / (((v4 + (v4 % (((v4 - (v4 % ((32768) | 1)))) | 1)))) | 1)))}
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float32 = 0.1
            var v10 float32 = -2.25
            v10 = v10
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    fmt.Println(v3, v1, v2)
    fmt.Println(v1)
    var v11 int16 = -836
    v2 -= v2
    var v12 int32 = 2147483647
    v13 := GenS1{F0: (v12 % ((1) | 1)), F1: (v12 - (^(v12))), F2: v4}
    if v3 {
        for fi14 := 0; fi14 < 14; fi14++ {
            switch "" {
            case "a":
                v2 = 65535
            }
            v3 = true
            var v15 uint16 = 255
            _ = fi14
            _ = v15
        }
        var v16 uint64 = 1
        _ = v16
    } else {
        fmt.Println(v3, v1)
        var v17 string = ("foo bar" + "hello")
        switch v17 {
        case "a":
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 bool = ((-2.25 == 3.14159) || (336 < 42))
                _ = fi18
                _ = v19
            }
        case "hello":
            if (!v3) {
                var v20 int32 = v12
                _ = v20
            } else {
                fmt.Println(v4)
            }
        }
        _ = v17
    }
    var v21 int = -1
    if ((!v3) || true) {
        fmt.Println(v3, v11, v21)
    }
    v22 := []int{}
    v22 = append(v22, v4)
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}

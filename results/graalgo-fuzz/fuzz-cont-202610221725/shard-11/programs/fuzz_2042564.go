package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 int32, p1 int8, p2 bool) int32 {
    v4 := []int{4294967295, -1, -874}
    var v5 float64 = 1e10
    fmt.Println(v5)
    v6 := GenS1{F0: p1, F1: -32768, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 457
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 int32 = 42
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = v2
        for fi8 := 0; fi8 < 5; fi8++ {
            v3 = -1
            v9 := GenS1{F0: 42, F1: 1000, F2: 1}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int32 = (v3 >> ((987) & 63))
            v12 := []int{fi10, fi7, 42, fi7}
            v3 -= 42
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v2 += v2
    for fi13 := 0; fi13 < 5; fi13++ {
        if (!true) {
            var v14 uint8 = 192
            _ = v14
        }
        switch 1 {
        case 42:
            v3 = (-(v3))
        }
        switch "a" {
        case "foo bar":
            v3 += (v3 & v3)
        case "a":
            var v15 bool = (!(-716 != 42))
            _ = v15
        default:
            var v16 float64 = (0.1 / 1.0)
            _ = v16
        }
        _ = fi13
    }
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 11; fi18++ {
        v1["\n\t"] = true
        v17 += (v17 >> ((0) & 63))
        _ = fi18
    }
    var v19 int8 = 39
    switch 2147483647 {
    case 1:
        if (0.0 < (v2 / 3.14159)) {
            v2 = float64(-2.25)
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                fmt.Println(v3, v17)
                var v21 string = "hello"
                _ = fi20
                _ = v21
            }
        }
    case -9223372036854775808:
        v2 = v2
    case 2147483647:
        v17 = (v17 << ((65535) & 63))
    default:
        if false {
            var v22 float64 = float64(3.14159)
            v19 = (v19 * -128)
            v19 = v19
            _ = v22
        }
    }
    v19 = 1
    var v23 uint8 = 0
    fmt.Println(v3, v17)
    var v24 int32 = (v3 ^ v3)
    for fi25 := 0; fi25 < 20; fi25++ {
        _ = len(v1)
        switch -9223372036854775808 {
        case 927:
            var v26 int64 = -9223372036854775808
            _ = v26
        case 4294967295:
            fmt.Println(v19, v19)
        default:
            if ((!((v3 / ((v24) | 1)) > v3)) || (1e10 < (float32(1e10) - 0.0))) {
                v27 := []int{(fi25 % ((fi25) | 1)), (fi25 + (-(fi25)))}
                _ = v27
            }
        }
        _ = fi25
    }
    v23 -= (v23 & (v23 ^ v23))
    var v28 int = 32767
    var v29 int64 = -34
    var v30 float64 = v2
    var v31 uint16 = 517
    for fi32 := 0; fi32 < 14; fi32++ {
        v23 = (v23 - 0)
        fmt.Println(fi32)
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
func genFn0(p0 string, p1 bool, p2 uint16) uint64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        switch ("" + ((p0 + p0) + (p0 + p0))) {
        case "hello":
            v6 := []int{(^(fi5))}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(fi5, fi7)
                _ = fi7
            }
        case "foo bar":
            p0 = (p0 + "foo bar")
        default:
            switch 128 {
            case 2147483647:
                p0 = p0
            }
        }
        if (!p1) {
            p1 = (!(32768 <= (fi5 % (((fi5 | (-(fi5)))) | 1))))
            v8 := []int{fi5, (fi5 >> ((182) & 63)), 878, fi5}
            _ = v8
        }
        v9 := []int{42, 32767}
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if p1 {
            if p1 {
                var v11 bool = ((fi10 * 2147483647) >= -1)
                _ = v11
            } else {
                v12 := GenS1{F0: -935, F1: 1, F2: 42}
                _ = v12
            }
        } else {
            v13 := GenS1{F0: -1, F1: 9223372036854775807, F2: 515}
            _ = v13
        }
        _ = fi10
    }
    fmt.Println(p0)
    var v14 bool = (0.0 == 0.0)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 32767, F1: -455, F2: -1}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 411
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    var v4 float64 = v3
    var v17 float64 = 3.14159
    delete(v2, -311)
    v4 = float64(0.0)
    switch 636 {
    case 4294967295:
        v17 = v4
    }
    fmt.Println(v3)
    var v18 int16 = 0
    if (false && false) {
        v4 = v3
        v19 := []int{4294967295, 127, 0}
        var v20 uint8 = 14
        _ = v19
        _ = v20
    } else {
        _ = len(v2)
        fmt.Println(v17)
    }
    fmt.Println(v1, v3, v4)
    var v21 float32 = 0.1
    var v22 uint16 = 42
    v17 -= float64(((v21 - (-2.25 * (1e10 + (0.1 * (float32(v4) * (((1.5 - v21) / (v21 - 1e10)) * (1e10 + v21))))))) - v21))
    var v23 float32 = v21
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            switch 0 {
            case 591:
                v3 = 3.14159
            default:
                _ = len(v2)
            }
            var v25 int8 = -59
            _ = v25
        }
        v26 := GenS1{F0: v18, F1: -1, F2: v18}
        _ = fi24
        _ = v26
    }
    v27 := GenS1{F0: v18, F1: -9223372036854775808, F2: 430}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}

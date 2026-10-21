package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
func genFn0(p0 uint32, p1 uint64, p2 string) string {
    v6 := GenS1{F0: 1, F1: 128, F2: 42}
    if ((448 == 1) || true) {
        var v7 float32 = ((0.1 / 0.1) / 1e10)
        _ = v7
    } else {
        var v8 string = p2
        v9 := []int{}
        var v10 bool = false
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{}
    switch 128 {
    case 135:
        var v12 int64 = 1
        _ = v12
    default:
        switch -1 {
        case -388:
            v13 := GenS1{F0: -1, F1: 39, F2: -32768}
            _ = v13
        case 2147483647:
            p2 = "a"
        default:
            p2 = (("" + p2) + "a")
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return ("\n\t" + p2)
}
func genFn1() int8 {
    v14 := GenS1{F0: 18, F1: 42, F2: 32767}
    var v15 uint32 = 750
    if (-1 != 2147483647) {
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := GenS2{F0: -662, F1: "foo bar", F2: 2147483647}
            for fi18 := 0; fi18 < 13; fi18++ {
                v15 -= v15
                var v19 int8 = -54
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v19 = v19
                default:
                    v20 := GenS1{F0: (v19 - v19), F1: 0, F2: 200}
                    if false {
                        v21 := []int{(fi16 / ((-167) | 1)), fi16, -495, (fi18 | (fi18 * fi16))}
                        v19 = v19
                        _ = v21
                    } else {
                        v19 = (v19 | (v19 * (v19 * v19)))
                        var v22 int16 = -32768
                        v22 = v22
                        v22 = 0
                        fmt.Println(v15)
                        _ = v22
                    }
                    _ = v20
                }
                _ = fi18
                _ = v19
            }
            var v23 uint64 = 18446744073709551615
            _ = fi16
            _ = v17
            _ = v23
        }
        v15 += (^(v15))
        var v24 uint32 = v15
        _ = v24
    } else {
        if ((^(v15)) <= (v15 % ((v15) | 1))) {
            v15 = (v15 * 0)
            var v25 int = 854
            _ = v25
        } else {
            v15 = (v15 * 0)
            fmt.Println(v15)
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v15 += 2147483647
        _ = fi26
    }
    _ = v14
    _ = v15
    return int8(v15)
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 4294967295
    var v3 int16 = 1000
    v4 := make(map[uint32]bool)
    var v5 uint = 587
    var v27 bool = ((float64(1.5) < ((float64(float32(-2.25)) + (0.0 * 1e10)) + 0.1)) && false)
    v27 = true
    switch -1 {
    case -2147483648:
        v1 -= (v1 >> ((106) & 63))
    case 127:
        v28 := []int{-9223372036854775808, -40, -2147483648}
        _ = v28
    }
    var v29 bool = (65535 == 2147483648)
    v30 := GenS1{F0: int8(v3), F1: 127, F2: (^(v3))}
    v31 := GenS1{F0: 42, F1: 128, F2: v3}
    v32 := v4[2147483648]
    v2 -= 67
    switch 42 {
    case 621:
        fmt.Println(v1)
    case 32767:
        var v33 int16 = v3
        _ = v33
    default:
        var v34 int16 = v3
        _ = v34
    }
    v1 = (v1 & v1)
    fmt.Println(v5, v3, v3)
    var v35 string = (("" + "foo bar") + ("hello" + ""))
    var v36 float32 = 0.1
    for fi37 := 0; fi37 < 20; fi37++ {
        switch 32767 {
        case 425:
            v5 = (^(v2))
        case 37:
            var v38 uint16 = v1
            _ = v38
        default:
            var v39 int8 = 1
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}

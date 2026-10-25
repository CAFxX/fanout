package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() int64 {
    v4 := GenS1{F0: 2, F1: 65535, F2: -2147483648}
    v5 := GenS1{F0: 0, F1: 65535, F2: 0}
    if false {
        if (true && true) {
            if (true || true) {
                var v6 float64 = (1.5 * ((1e10 / 1.5) * (-2.25 * 1e10)))
                v7 := []int{127, 32768, 32768, -321}
                switch "a" {
                case "a":
                    v7 = append(v7, 0)
                case "foo bar":
                    var v8 int = 32767
                    _ = v8
                }
                _ = v6
                _ = v7
            }
        } else {
            v9 := []int{}
            var v10 uint16 = 32767
            _ = v10
            _ = v9
        }
        v11 := []int{661, 960}
        _ = v11
    } else {
        v12 := []int{128, 32767}
        switch -77 {
        case -2147483648:
            var v13 uint16 = 0
            _ = v13
        }
        var v14 uint8 = 2
        _ = v12
        _ = v14
    }
    v15 := []int{754, 9223372036854775807, 9223372036854775807}
    var v16 float64 = 1.5
    _ = v15
    _ = v16
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 uint16, p2 int64) int16 {
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int8 = 42
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 1, F1: (^(p1)), F2: 0}
    p1 = (p1 ^ (p1 * p1))
    p2 = (p2 | p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 32767
}

func main() {
    var v1 int8 = 0
    var v2 uint = 65535
    var v3 float32 = 1.5
    v3 += (-2.25 * float32(-2.25))
    var v20 int64 = -550
    v21 := GenS1{F0: 127, F1: 381, F2: -1}
    fmt.Println(v1)
    v1 += (-(v1))
    v22 := GenS2{F0: (v1 + 0), F1: 2}
    if (360 == 42) {
        var v23 int8 = (v1 + v1)
        _ = v23
    }
    var v24 int8 = (v1 >> ((v2) & 63))
    var v25 uint16 = 694
    for fi26 := 0; fi26 < 3; fi26++ {
        v1 += (v24 - -1)
        _ = fi26
    }
    v2 = v2
    v3 -= v3
    fmt.Println(v2, v20)
    var v27 int64 = (v20 % ((v20) | 1))
    var v28 int64 = genFn0()
    var v29 uint16 = (v25 + (^(v25)))
    v30 := GenS1{F0: 128, F1: (v25 % ((v29) | 1)), F2: -2147483648}
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint = 0
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}

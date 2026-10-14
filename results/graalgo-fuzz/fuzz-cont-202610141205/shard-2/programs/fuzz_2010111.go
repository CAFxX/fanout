package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint; F1 int32; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 int32, p2 int) uint16 {
    fmt.Println(p1, p0)
    p2 += (p2 | (p2 - 32768))
    switch p2 {
    case -9223372036854775808:
        if true {
            p2 -= -2147483648
        } else {
            var v5 int = p2
            var v6 uint16 = 1
            _ = v5
            _ = v6
        }
    case 32768:
        switch (((("hello" + "hello") + (((("\n\t" + "\n\t") + "a") + "foo bar") + "\n\t")) + "hello") + "a") {
        case "foo bar":
            var v7 uint16 = 42
            _ = v7
        case "":
            v8 := []int{p2, 847, p2, p2}
            _ = v8
        default:
            v9 := GenS1{F0: 127}
            _ = v9
        }
    }
    if true {
        p2 += p2
    } else {
        v10 := GenS2{F0: -965}
        _ = v10
    }
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(p0, p0)
    case 754:
        p1 = (p1 << ((184) & 63))
    case 0:
        p0 = (^(p0))
    default:
        switch p2 {
        case 9223372036854775807:
            var v11 int32 = -1
            _ = v11
        case -1:
            if false {
                v12 := GenS1{F0: 1}
                v13 := []int{803}
                for fi14 := 0; fi14 < 14; fi14++ {
                    fmt.Println(p0, p2)
                    v15 := GenS3{F0: 847, F1: (-2.25 != 1.5)}
                    _ = fi14
                    _ = v15
                }
                _ = v12
                _ = v13
            } else {
                switch (p2 - p2) {
                case 2147483647:
                    for fi16 := 0; fi16 < 8; fi16++ {
                        fmt.Println(p1, p0)
                        p2 += 9223372036854775807
                        _ = fi16
                    }
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 94
    var v3 uint32 = 4294967295
    v4 := make(map[int]bool)
    v1 = (v1 + (v1 + v1))
    var v17 float32 = (((0.1 + 1.5) * 0.1) + 1e10)
    if ((42 < 1) || false) {
        v18 := GenS4{F0: -556, F1: 42}
        v2 = v2
        var v19 int = -147
        _ = v18
        _ = v19
    } else {
        v20 := v4[640]
        fmt.Println(v17, v1)
        _ = v20
    }
    v21 := GenS5{F0: 42, F1: v2, F2: 42}
    var v22 uint16 = 2
    v23 := v4[-9223372036854775808]
    var v24 int = -1
    v25 := GenS6{F0: 4294967295, F1: 516, F2: 0}
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            v22 += (^(v22))
            _ = fi27
        }
        _ = fi26
    }
    switch v24 {
    case 9223372036854775807:
        var v28 bool = false
        _ = v28
    }
    fmt.Println(v24, v17)
    switch (v24 / (((v24 * (v24 * v24))) | 1)) {
    case 32768:
        var v29 string = (("\n\t" + "") + v1)
        _ = v29
    case -800:
        v30 := GenS6{F0: 18446744073709551615, F1: 42, F2: 318}
        _ = v30
    default:
        for fi31 := 0; fi31 < 17; fi31++ {
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := GenS7{F0: -1, F1: v22, F2: -865}
                _ = fi32
                _ = v33
            }
            var v34 uint8 = (v2 / ((v2) | 1))
            _ = fi31
            _ = v34
        }
    }
    v23 = (0.0 <= (1.5 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

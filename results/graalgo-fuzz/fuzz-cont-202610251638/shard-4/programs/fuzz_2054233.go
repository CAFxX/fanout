package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64, p1 int, p2 int64) uint32 {
    var v5 uint32 = 748
    v6 := []int{(p1 - (p1 + (p1 % ((p1) | 1)))), (p1 << ((0) & 63)), 15, (p1 | 127)}
    if true {
        var v7 uint16 = 643
        _ = v7
    } else {
        v8 := GenS1{F0: 1, F1: 325}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v5 + v5)
}
func genFn1(p0 int8, p1 int32) uint16 {
    switch 568 {
    case 122:
        var v9 int64 = 0
        _ = v9
    case 881:
        p0 -= (p0 ^ p0)
    case 9223372036854775807:
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int32 = -35
            if false {
                var v12 int16 = 32767
                var v13 uint64 = 42
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 9; fi14++ {
                    fmt.Println(v11, p0)
                    v11 += (v11 * (p1 << ((2) & 63)))
                    _ = fi14
                }
            }
            _ = fi10
            _ = v11
        }
    }
    p0 = (p0 >> ((42) & 63))
    _ = p0
    _ = p1
    return 268
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = ((0.1 - (0.1 + (-2.25 - (((0.1 / float64(0.1)) + 1.5) * ((float64((0.0 - 3.14159)) * 1.5) / 1.5))))) - 1e10)
    var v3 uint8 = v1
    var v4 bool = true
    if ("\n\t" > ("hello" + "a")) {
        v15 := GenS2{F0: 42, F1: 2147483648, F2: 65535}
        v16 := GenS3{F0: 18446744073709551615, F1: ""}
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: false, F1: 42}
    v18 := GenS1{F0: -1, F1: -508}
    if false {
        var v19 uint8 = (v1 ^ (v1 - v1))
        fmt.Println(v3, v2)
        v4 = (float64((1.5 / ((((3.14159 - 1e10) - 0.0) * 1.5) + float32((0.0 - v2))))) <= (0.0 * v2))
        _ = v19
    }
    v1 -= 255
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int64 = 42
        var v22 float32 = ((1e10 - 0.1) / float32(((v2 / (((v2 / v2) - (-2.25 / 3.14159)) - 1.5)) / 1.0)))
        v23 := GenS5{F0: (v21 * v21)}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    switch ((("\n\t" + "hello") + "foo bar") + "\n\t") {
    case "foo bar":
        fmt.Println(v2, v3)
    case "hello":
        v1 -= 0
    }
    v3 = v3
    v4 = ((v1 & v3) <= (v1 ^ (v1 + v1)))
    var v24 uint8 = (v3 / ((127) | 1))
    v3 -= (^(v3))
    if v4 {
        v25 := GenS2{F0: 289, F1: 4294967295, F2: 479}
        v26 := []int{}
        var v27 string = "foo bar"
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}

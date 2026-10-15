package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 uint) uint {
    var v4 int16 = 42
    v5 := []int{-2147483648, 4294967295, 32767}
    v6 := GenS1{F0: "\n\t", F1: 143, F2: uint16(p1)}
    p1 = 255
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (p1 & p0)
}

func main() {
    var v1 int32 = -2147483648
    var v2 bool = (18446744073709551615 > 295)
    v3 := []uint32{}
    if true {
        v3 = append(v3, 2147483648)
        v1 = (^(v1))
        switch "" {
        case "":
            v7 := GenS2{F0: 475, F1: 559, F2: 2}
            _ = v7
        case "foo bar":
            switch 128 {
            case 32767:
                var v8 int = 266
                _ = v8
            case 9223372036854775807:
                var v9 uint8 = 128
                _ = v9
            }
        }
    } else {
        var v10 uint16 = 65535
        _ = len(v3)
        _ = v10
    }
    if (128 > 128) {
        v11 := GenS2{F0: 42, F1: 105, F2: 72}
        for fi12 := 0; fi12 < 13; fi12++ {
            switch (("hello" + "a") + "foo bar") {
            case "foo bar":
                for fi13 := 0; fi13 < 9; fi13++ {
                    if v2 {
                        var v14 uint16 = 42
                        _ = v14
                    }
                    var v15 uint8 = 108
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi13
                    _ = v15
                }
            case "":
                if true {
                    v3 = append(v3, 2)
                    var v16 string = "a"
                    _ = v16
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v2 = (49 == 164)
                    }
                }
            case "hello":
                var v17 uint8 = 128
                _ = v17
            default:
                var v18 int32 = v1
                _ = v18
            }
            _ = fi12
        }
        v1 = (v1 % ((v1) | 1))
        _ = v11
    }
    v2 = v2
    var v19 bool = false
    v19 = v19
    v3 = append(v3, 4294967295)
    for fi20 := 0; fi20 < 16; fi20++ {
        v1 = (-(v1))
        v2 = ((((3.14159 * (0.1 * (3.14159 / 1.0))) / -2.25) * 1.5) <= 0.1)
        _ = fi20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v22 float32 = (-2.25 - 0.1)
        fmt.Println(v22, v2, v2)
        _ = fi21
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (-1 != 0)
    fmt.Println(v1, v2, v1)
    var v23 int = -964
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}

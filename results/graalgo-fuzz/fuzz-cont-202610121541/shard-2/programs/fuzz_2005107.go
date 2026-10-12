package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint64 = 2
    var v2 bool = ((v1 / ((v1) | 1)) > v1)
    var v3 float32 = 0.0
    var v4 bool = v2
    var v5 int8 = -128
    var v6 int16 = -344
    v2 = v4
    if (!(32767 >= (v6 - (v6 >> ((1) & 63))))) {
        v2 = (!false)
        v7 := []int{-2147483648, 42, 128}
        if v4 {
            switch ("" + "a") {
            case "\n\t":
                v2 = (!(!v4))
            default:
                var v8 uint8 = 255
                _ = v8
            }
            v2 = v4
        } else {
            v6 = (v6 * v6)
            for fi9 := 0; fi9 < 12; fi9++ {
                v4 = false
                _ = fi9
            }
            v10 := GenS1{F0: 18446744073709551615}
            _ = v10
        }
        _ = v7
    }
    fmt.Println(v2)
    var v11 uint64 = v1
    if (v1 < 2) {
        v5 -= (v5 * (v5 + v5))
        switch "foo bar" {
        case "foo bar":
            v12 := GenS2{F0: 42, F1: -1}
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 9; fi13++ {
                if v2 {
                    v4 = (!v2)
                    v14 := []int{(fi13 % ((523) | 1)), 4294967295, 150}
                    _ = v14
                }
                switch "a" {
                case "hello":
                    v15 := []int{(fi13 & fi13)}
                    _ = v15
                case "":
                    v16 := GenS2{F0: 1, F1: 317}
                    _ = v16
                case "foo bar":
                    var v17 float64 = ((0.0 * 1.5) + (0.1 / 1.0))
                    _ = v17
                default:
                    var v18 int = fi13
                    _ = v18
                }
                switch ("hello" + "a") {
                case "\n\t":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v4 = (!v4)
                        v20 := GenS2{F0: 0, F1: -471}
                        _ = fi19
                        _ = v20
                    }
                }
                _ = fi13
            }
        }
    }
    v21 := GenS3{F0: 9223372036854775807, F1: v2, F2: 401}
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint8 = 42
        v24 := []int{}
        _ = fi22
        _ = v23
        _ = v24
    }
    v4 = (((((((0.1 - 1.5) - (0.1 * (1e10 / -2.25))) * 1.5) >= 0.0) && (-9223372036854775808 > -9223372036854775808)) || (v6 == (v6 ^ (-(v6))))) || true)
    if (!(!v2)) {
        fmt.Println(v11, v5, v11)
    }
    fmt.Println(v2, v1, v2)
    fmt.Println(v6, v5)
    v3 = (1.5 / (v3 + (float32(1e10) + 0.0)))
    v25 := []int{130, 271, -946}
    if true {
        v25 = append(v25, 4294967295)
        if ("" > "") {
            var v26 uint8 = 218
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = v26
        }
    }
    if v4 {
        for fi27 := 0; fi27 < 11; fi27++ {
            if v2 {
                v3 += 0.0
                v1 -= 0
            } else {
                v25 = append(v25, fi27)
            }
            v28 := GenS4{F0: (^(fi27))}
            _ = fi27
            _ = v28
        }
        if (!v4) {
            v4 = (v2 && v2)
        } else {
            var v29 int32 = 0
            _ = v29
        }
        var v30 int32 = int32(v11)
        _ = v30
    }
    var v31 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

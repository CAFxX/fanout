package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }

func main() {
    var v1 int = 42
    var v2 uint32 = 52
    v3 := []uint32{}
    var v4 uint32 = (v2 & (v2 >> ((587) & 63)))
    switch "foo bar" {
    case "hello":
        var v5 int16 = 168
        _ = v5
    case "foo bar":
        switch "a" {
        case "hello":
            var v6 int = (^(v1))
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: 2, F1: 18446744073709551615}
            _ = v7
        default:
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v4 -= v2
                    var v10 bool = false
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: 255, F1: 217}
                if (77 != (^(v2))) {
                    var v12 string = "hello"
                    _ = v12
                }
                _ = fi8
                _ = v11
            }
        }
    }
    var v13 int8 = -1
    v14 := GenS2{F0: 242, F1: 2147483648, F2: 32768}
    switch ("hello" + "") {
    case "":
        var v15 uint16 = 1
        _ = v15
    }
    switch (v1 / ((v1) | 1)) {
    case 4294967295:
        v2 = v2
    case 9223372036854775807:
        if (("\n\t" + "a") == "hello") {
            v4 = (^(v4))
        } else {
            var v16 float64 = 0.1
            var v17 uint8 = 0
            _ = v16
            _ = v17
        }
    case 1:
        v3 = append(v3, (v2 % ((1) | 1)))
    default:
        v18 := GenS1{F0: 681, F1: 859}
        _ = v18
    }
    var v19 int8 = -1
    v4 += (v4 * (v4 << ((2) & 63)))
    var v20 float64 = 1.5
    var v21 uint64 = 42
    for fi22 := 0; fi22 < 20; fi22++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v23 float32 = 0.0
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}

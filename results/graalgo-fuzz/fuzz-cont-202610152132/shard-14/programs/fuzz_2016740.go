package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 1
    var v3 uint32 = 1
    v4 := make(map[uint32]bool)
    v5 := []int{}
    v6 := GenS1{F0: ("" + "a"), F1: 0}
    var v7 uint16 = 65535
    var v8 float64 = (((0.0 - float64((3.14159 + (3.14159 / 1.5)))) * 1.5) / 0.1)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: "hello", F1: 871}
        if ("hello" < "hello") {
            switch fi9 {
            case -338:
                v11 := GenS2{F0: -9223372036854775808, F1: (v7 * (v7 ^ v7))}
                _ = v11
            default:
                fmt.Println(v3)
            }
            v4[(v3 - 262)] = ((0 >= 282) || (-128 >= -1))
            v7 = v7
        } else {
            var v12 uint = 2
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v4[(v3 + v3)] = false
    v8 = float64(1e10)
    v13 := GenS3{F0: -9223372036854775808, F1: (v3 ^ v3), F2: "\n\t"}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint8 = (^(v2))
        v7 += v7
        if false {
            switch "hello" {
            case "a":
                v7 = (v7 - (v7 - v7))
            case "hello":
                var v16 uint = 4294967295
                _ = v16
            default:
                if (false && (-1 == -87)) {
                    v8 = -2.25
                    fmt.Println(fi14)
                }
            }
        } else {
            switch (("foo bar" + "\n\t") + "a") {
            case "":
                fmt.Println(v2, v2, fi14)
            case "\n\t":
                v15 = (v2 & (v15 + 255))
            }
            var v17 float32 = (1e10 / -2.25)
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    switch 32767 {
    case 600:
        v18 := GenS2{F0: 99, F1: 65535}
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v3 -= (v3 / ((0) | 1))
            var v21 string = ("foo bar" + ("foo bar" + ""))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v22 int64 = 782
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 float64 = v8
        _ = fi23
        _ = v24
    }
    v7 = 255
    var v25 uint32 = 2147483648
    v26 := GenS1{F0: "foo bar", F1: -433}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

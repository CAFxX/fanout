package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 1
    var v3 bool = ((0.0 / -2.25) > ((-2.25 / float64((float32((1.5 * 0.1)) / float32(0.1)))) + float64((0.0 + 1.5))))
    var v4 bool = true
    v4 = v4
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int64 = -9223372036854775808
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -1}
    v3 = ((v2 | v2) != -28)
    var v8 float64 = 0.1
    var v9 uint = 155
    v10 := []int{128, -26, 4294967295, 4294967295}
    v11 := GenS2{F0: 9223372036854775807, F1: v4, F2: -558}
    v12 := GenS2{F0: -9223372036854775808, F1: v4, F2: 42}
    v3 = (!false)
    if false {
        var v13 uint64 = 33
        _ = v13
    } else {
        v14 := GenS3{F0: 40}
        var v15 int32 = 0
        _ = v14
        _ = v15
    }
    var v16 uint32 = 65535
    for fi17 := 0; fi17 < 17; fi17++ {
        switch fi17 {
        case -842:
            if true {
                v18 := GenS3{F0: 741}
                fmt.Println(v4)
                var v19 uint64 = 573
                _ = v18
                _ = v19
            }
        case 32767:
            if ("a" <= "\n\t") {
                if (731 != 18446744073709551615) {
                    v2 = (v2 | 0)
                } else {
                    var v20 uint16 = 2
                    _ = v20
                }
            }
        default:
            if ((v2 >> ((v9) & 63)) < (v2 * (^(v2)))) {
                v1["hello"] = -210
            } else {
                switch "foo bar" {
                case "a":
                    switch "hello" {
                    case "a":
                        v3 = v3
                    }
                case "foo bar":
                    v21 := GenS4{F0: (v2 >> (((v9 | v9)) & 63))}
                    _ = v21
                }
                v22 := GenS4{F0: v2}
                _ = v22
            }
        }
        var v23 string = "\n\t"
        _ = fi17
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

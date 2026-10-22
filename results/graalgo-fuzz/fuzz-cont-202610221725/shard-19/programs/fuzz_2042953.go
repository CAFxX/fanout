package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 string; F1 string }

func main() {
    v1 := []int{-205}
    var v2 int32 = 0
    var v3 uint = 785
    v4 := []int{32768, 32768, 42, -855}
    v5 := []int{}
    var v6 int8 = 0
    switch "a" {
    case "hello":
        _ = len(v5)
    default:
        v7 := GenS1{F0: 2147483647}
        _ = v7
    }
    v6 = (-(v6))
    if true {
        switch 32768 {
        case 1:
            v8 := GenS2{F0: (v3 / (((^(v3))) | 1)), F1: 1, F2: 9223372036854775807}
            _ = v8
        case 67:
            if len(v4) > 0 {
                v4[0] = -664
            }
        case 128:
            fmt.Println(v2, v3)
        default:
            var v9 int16 = -1
            _ = v9
        }
        v6 = v6
        var v10 uint16 = 255
        _ = v10
    }
    v11 := GenS3{F0: ("" + "foo bar"), F1: "a"}
    v12 := GenS1{F0: -583}
    v6 -= -128
    if true {
        v3 = v3
        v3 = 0
        var v13 int32 = (-(v2))
        _ = v13
    } else {
        fmt.Println(v3, v6, v2)
        fmt.Println(v2, v3)
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int8 = -125
            var v16 string = ("\n\t" + "a")
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS2{F0: 1, F1: (fi17 % ((127) | 1)), F2: (fi17 - -2147483648)}
            _ = fi18
            _ = v19
        }
        _ = len(v4)
        _ = fi17
    }
    var v20 int32 = (v2 | 458)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

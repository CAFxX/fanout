package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 2
    var v3 uint = 42
    v4 := []int{-9223372036854775808, 32767, -80, 1}
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 3; fi6++ {
        switch ("\n\t" + "") {
        case "\n\t":
            v7 := GenS1{F0: "foo bar"}
            _ = v7
        case "hello":
            v5 -= v5
        default:
            var v8 uint = v3
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS2{F0: v3, F1: 159, F2: 873}
    for fi10 := 0; fi10 < 12; fi10++ {
        v2 = uint64(v5)
        var v11 int64 = -1
        var v12 bool = (!false)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = 1
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint = 671
        _ = fi14
        _ = v15
    }
    switch 32767 {
    case 811:
        fmt.Println(v5, v2, v1)
    default:
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS3{F0: 42, F1: 782}
            _ = fi16
            _ = v17
        }
    }
    v13 += v13
    var v18 float64 = (-2.25 - 1e10)
    v19 := GenS4{F0: 0}
    _ = len(v4)
    v20 := GenS1{F0: "foo bar"}
    v21 := GenS5{F0: 637}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

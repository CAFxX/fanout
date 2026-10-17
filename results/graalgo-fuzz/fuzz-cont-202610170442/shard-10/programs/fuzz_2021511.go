package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int64 = 577
    var v2 float64 = (-2.25 + 0.0)
    v3 := []uint32{2147483648}
    var v4 uint8 = 42
    v5 := GenS1{F0: true}
    fmt.Println(v4)
    v4 += (^(v4))
    if (3.14159 <= (1e10 + 0.0)) {
        v4 -= v4
    } else {
        fmt.Println(v1, v1)
    }
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 := GenS1{F0: false}
    if false {
        var v7 uint8 = (^(v4))
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS2{F0: 0, F1: 42, F2: -589}
            v10 := GenS3{F0: -145}
            if (v2 > v2) {
                fmt.Println(v4)
                if false {
                    v7 = 128
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        switch "a" {
        case "":
            v2 = 0.0
        case "a":
            v3 = append(v3, 0)
        }
        _ = v7
    }
    v3 = append(v3, 1)
    v11 := GenS4{F0: "", F1: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 int32 = 42
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint = 1
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int = fi15
            _ = fi15
            _ = v16
        }
        v4 -= v4
        _ = fi13
        _ = v14
    }
    _ = len(v3)
    v1 = v1
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint32 = 2147483647
            _ = fi17
            _ = v18
        }
        var v19 float64 = 1e10
        v20 := GenS5{F0: (^(v1))}
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

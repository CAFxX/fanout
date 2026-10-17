package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]bool)
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: (("" + ("hello" + "a")) + "a"), F1: 990}
    var v6 uint16 = 65535
    v1 = v1
    var v7 bool = true
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 977
        _ = fi8
        _ = v9
    }
    if v7 {
        if v7 {
            v10 := GenS2{F0: 100}
            v6 -= (v6 ^ (^(v6)))
            _ = v10
        } else {
            var v11 uint = 42
            var v12 uint8 = 255
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int = 0
                for fi15 := 0; fi15 < 8; fi15++ {
                    v16 := GenS2{F0: 100}
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint16 = v6
        v19 := GenS2{F0: -105}
        _ = len(v2)
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 float64 = 1.5
    v7 = false
    if (42 <= -57) {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 uint16 = v6
            _ = fi21
            _ = v22
        }
    } else {
        v7 = v7
        fmt.Println(v1, v6)
    }
    v6 = 253
    v23 := v4[279]
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS3{F0: 85}
        var v26 int64 = -940
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

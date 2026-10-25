package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []uint32{4294967295, 515, 4294967295}
    var v2 uint16 = 32767
    var v3 float32 = 1.5
    var v4 int32 = 0
    var v5 int32 = v4
    var v6 int64 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 1, F1: 0}
        v9 := GenS2{F0: (v2 ^ (v2 / (((v2 / ((v2) | 1))) | 1)))}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v5
    v2 = v2
    var v10 int8 = 42
    for fi11 := 0; fi11 < 14; fi11++ {
        if (v2 != 65535) {
            v2 = v2
            v12 := GenS3{F0: 0}
            _ = v12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi13
        }
        _ = fi11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

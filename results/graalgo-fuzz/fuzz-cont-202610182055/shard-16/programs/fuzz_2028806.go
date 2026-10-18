package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 uint64, p2 uint16) bool {
    p0 = 128
    v5 := GenS1{F0: 0, F1: 855, F2: -626}
    var v6 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (1.5 <= (3.14159 + 3.14159))
}

func main() {
    v1 := []uint32{}
    v2 := []int{392}
    v3 := []string{}
    var v4 int8 = 42
    fmt.Println(v4)
    var v7 uint8 = 42
    var v8 int = -782
    v9 := GenS2{F0: 42, F1: 592, F2: -1}
    switch 710 {
    case -1000:
        var v10 bool = (1 >= 9223372036854775807)
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 65535
            _ = len(v2)
            _ = fi11
            _ = v12
        }
    case -408:
        var v13 float32 = ((3.14159 + 0.0) / ((3.14159 - -2.25) - 1.5))
        _ = v13
    }
    if false {
        var v14 uint8 = 131
        _ = v14
    } else {
        switch (v8 - v8) {
        case 128:
            switch "" {
            case "\n\t":
                fmt.Println(v7, v8, v7)
            case "hello":
                v15 := GenS3{F0: 42, F1: int8(v7)}
                _ = v15
            default:
                v7 = 199
            }
        case 1:
            var v16 float32 = 1e10
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v18 := GenS1{F0: 1, F1: 4294967295, F2: -293}
        _ = fi17
        _ = v18
    }
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

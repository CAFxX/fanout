package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 32
    v3 := make(map[string]uint32)
    v4 := []uint32{2147483647, 428, 2147483648}
    v5 := make(map[string]bool)
    var v6 string = "\n\t"
    delete(v3, v6)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(fi7)
        _ = len(v4)
        v6 = ("hello" + v6)
        _ = fi7
    }
    v2 = 341
    v8 := v3["\n\t"]
    v2 += v2
    fmt.Println(v6)
    if ((9223372036854775807 >= -1) || (0.0 != 0.1)) {
        var v9 int64 = -1
        _ = v9
    } else {
        v10 := GenS1{F0: v2, F1: 0, F2: (v8 & v8)}
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v10
    }
    var v11 uint8 = 1
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS2{F0: 42, F1: 42, F2: 1}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS3{F0: (^(v11))}
        _ = fi14
        _ = v15
    }
    v6 = v6
    v16 := GenS1{F0: uint(v11), F1: 100, F2: (v8 << (((v2 & v2)) & 63))}
    var v17 uint64 = 726
    var v18 uint8 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

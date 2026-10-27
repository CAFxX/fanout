package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (82 > 2147483647)
    var v3 uint8 = 255
    var v4 int32 = 2147483647
    var v5 uint64 = 666
    v6 := GenS1{F0: -418}
    if (350 == int(v3)) {
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        fmt.Println(v3, v3, v3)
    }
    v5 += v5
    var v8 uint32 = 663
    if v2 {
        v4 = 2147483647
    } else {
        v5 = v5
        v5 = (^(v5))
        v4 += 1
    }
    var v9 int64 = 731
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int = (fi10 << ((235) & 63))
        v8 = (v8 & v8)
        _ = fi10
        _ = v11
    }
    switch 32767 {
    case 4294967295:
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v12 float32 = 0.0
            _ = v12
        }
    default:
        if ((v3 << ((745) & 63)) < v3) {
            if true {
                delete(v1, -275)
                var v13 uint16 = 837
                _ = v13
            } else {
                var v14 int32 = (v4 << ((18446744073709551615) & 63))
                v15 := GenS1{F0: 1}
                if v2 {
                    v4 -= v4
                } else {
                    v8 = v8
                    v5 -= v5
                    var v16 int8 = 100
                    _ = v16
                }
                _ = v14
                _ = v15
            }
            var v17 float64 = -2.25
            var v18 uint16 = 0
            _ = v17
            _ = v18
        }
    }
    v3 = 0
    var v19 bool = true
    var v20 uint8 = (^(v3))
    v9 -= int64(v8)
    fmt.Println(v8, v3, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

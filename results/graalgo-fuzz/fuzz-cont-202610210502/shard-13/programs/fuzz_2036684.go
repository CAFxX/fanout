package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    v3 := []int{}
    var v4 uint16 = 255
    var v5 uint16 = (v4 % ((0) | 1))
    var v6 uint = 4294967295
    if true {
        if (-591 < 882) {
            v7 := GenS1{F0: -1, F1: 939}
            switch -9223372036854775808 {
            case -2147483648:
                var v8 int16 = 32767
                _ = v8
            }
            v2 += -2147483648
            _ = v7
        }
        var v9 int64 = -9223372036854775808
        v5 = 483
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int16 = 1000
        v3 = append(v3, (fi10 << ((42) & 63)))
        if false {
            v12 := GenS1{F0: 9223372036854775807, F1: 65535}
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := GenS2{F0: (v6 + v6), F1: (v5 >> (((v6 | v6)) & 63))}
    v4 = v4
    v14 := GenS3{F0: 42, F1: int64(v6), F2: 80}
    var v15 int16 = 498
    v16 := GenS4{F0: 0, F1: 42}
    var v17 uint = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

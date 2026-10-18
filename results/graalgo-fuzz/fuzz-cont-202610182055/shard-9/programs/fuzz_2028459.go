package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 string; F2 string }

func main() {
    var v1 uint64 = 299
    var v2 uint8 = 128
    v3 := make(map[uint32]bool)
    var v4 int16 = 42
    var v5 int8 = -128
    var v6 int64 = 821
    if ((v5 * v5) <= 1) {
        v7 := []int{}
        var v8 bool = (false && false)
        _ = v7
        _ = v8
    } else {
        v6 -= v6
        var v9 int = 0
        _ = v9
    }
    v1 = v1
    if (!true) {
        fmt.Println(v1, v4, v5)
        v5 = int8(v2)
    }
    v2 += 127
    v5 -= v5
    if false {
        v10 := GenS1{F0: 4294967295, F1: 32767, F2: 0}
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS2{F0: 65535, F1: 42}
            _ = fi11
            _ = v12
        }
        var v13 int32 = 0
        _ = v10
        _ = v13
    }
    v5 = 0
    switch "" {
    case "hello":
        switch -1 {
        case -1:
            v14 := GenS3{F0: 2}
            _ = v14
        case -9223372036854775808:
            fmt.Println(v2)
        }
    case "foo bar":
        var v15 int = 2147483647
        _ = v15
    default:
        v5 += 51
    }
    v16 := GenS4{F0: v6}
    v17 := v3[708]
    switch 0 {
    case -2147483648:
        fmt.Println(v5)
    }
    var v18 uint32 = 2
    var v19 int16 = (v4 & (v4 << ((4294967295) & 63)))
    v20 := GenS5{F0: (^(v18)), F1: ("\n\t" + "hello"), F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 string }
func genFn0() int {
    var v7 uint64 = 4294967295
    var v8 bool = false
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = 1.5
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 int16 = 32767
    v6 := make(map[string]string)
    var v9 int = 4294967295
    v10 := GenS1{F0: genFn0(), F1: (v5 - (v5 * v5))}
    var v11 uint32 = 2147483648
    var v12 uint64 = 828
    var v13 float32 = v2
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            if (!(!(v1 != v1))) {
                fmt.Println(v9, fi14)
                var v16 bool = true
                _ = v16
            } else {
                var v17 uint32 = (v11 << ((42) & 63))
                v18 := GenS1{F0: genFn0(), F1: (v5 + (v5 - -1))}
                _ = v17
                _ = v18
            }
            v19 := GenS2{F0: -1, F1: v5, F2: ("a" + "foo bar")}
            _ = fi15
            _ = v19
        }
        v9 = fi14
        _ = fi14
    }
    var v20 int8 = 42
    v21 := GenS1{F0: int(v12), F1: v5}
    switch genFn0() {
    case 127:
        fmt.Println(v12, v20, v11)
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 string = ""
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

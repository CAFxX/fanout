package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }

func main() {
    var v1 int = 128
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 int8 = -119
    var v5 int32 = -2147483648
    var v6 float32 = 1e10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 string = "foo bar"
    _ = len(v2)
    fmt.Println(v5)
    if ((!true) && false) {
        v7 = "\n\t"
    }
    v1 += v1
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8)
        v9 := GenS1{F0: 9223372036854775807, F1: (-(v1)), F2: -63}
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(v4, v7)
            var v11 int64 = 1
            if (true && true) {
                switch (fi10 << ((255) & 63)) {
                case 1:
                    v4 += int8(v5)
                case 32768:
                    v7 = "hello"
                }
            } else {
                v12 := GenS1{F0: int64(v4), F1: 42, F2: 1}
                v1 = (v1 << ((1) & 63))
                switch v7 {
                case "hello":
                    v13 := GenS1{F0: -9223372036854775808, F1: fi10, F2: v4}
                    _ = v13
                default:
                    v14 := GenS1{F0: v11, F1: 688, F2: 113}
                    _ = v14
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v6)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

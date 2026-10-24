package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []int{128, 873, -2147483648, -9223372036854775808}
    _ = len(v3)
    v4 := v2[1]
    v4 = ("hello" + "hello")
    var v5 uint = 65535
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint16 = uint16(v5)
        for fi8 := 0; fi8 < 18; fi8++ {
            v5 += (^(v5))
            v9 := v1[65535]
            _ = fi8
            _ = v9
        }
        var v10 int = fi6
        _ = fi6
        _ = v10
        _ = v7
    }
    v3 = append(v3, -920)
    var v11 float32 = float32((0.0 * 0.1))
    v4 = (v4 + (v4 + v4))
    if (-1 <= 128) {
        fmt.Println(v4, v4)
    }
    switch -9223372036854775808 {
    case 42:
        switch (("hello" + v4) + (v4 + "\n\t")) {
        case "hello":
            if (-1 == -2147483648) {
                var v12 uint64 = 503
                fmt.Println(v11, v4, v12)
                _ = v12
            }
        default:
            v13 := v1[4294967295]
            _ = v13
        }
    default:
        var v14 uint32 = 42
        _ = v14
    }
    v15 := GenS1{F0: -675}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

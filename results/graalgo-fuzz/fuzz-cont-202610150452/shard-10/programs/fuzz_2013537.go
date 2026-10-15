package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
func genFn0() int8 {
    if false {
        switch "" {
        case "":
            var v5 int64 = 42
            _ = v5
        }
    }
    var v6 uint = 953
    v6 += 65535
    v6 = 480
    _ = v6
    return 83
}

func main() {
    var v1 uint8 = 2
    var v2 int = -9223372036854775808
    var v3 uint32 = 2147483647
    v4 := make(map[string]uint32)
    switch v2 {
    case 32768:
        v1 = v1
    case 2147483647:
        var v7 uint32 = 2147483648
        _ = v7
    case 0:
        v1 = v1
    default:
        if (-1 > 0) {
            var v8 int = v2
            _ = v8
        } else {
            var v9 uint = 18446744073709551615
            v2 = -202
            _ = v9
        }
    }
    fmt.Println(v1)
    var v10 bool = false
    v10 = v10
    var v11 uint32 = (v3 & uint32(v1))
    v12 := GenS1{F0: "\n\t", F1: ("" + "\n\t"), F2: (v2 ^ v2)}
    if v10 {
        switch v2 {
        case 32767:
            var v13 int64 = 0
            _ = v13
        default:
            v14 := []int{v2, v2, (v2 ^ -2147483648), v2}
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 bool = (1 == 42)
            _ = fi15
            _ = v16
        }
        fmt.Println(v3, v10)
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v10 = (v10 || (v10 || v10))
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}

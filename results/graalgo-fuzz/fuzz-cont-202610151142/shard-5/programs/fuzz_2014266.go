package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 85
    var v3 int8 = 0
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    var v6 string = "\n\t"
    v6 = ""
    if false {
        var v7 int = 770
        var v8 float64 = -2.25
        if true {
            v3 += (v3 << ((216) & 63))
            fmt.Println(v2, v6, v2)
        } else {
            var v9 int8 = (v3 % ((v3) | 1))
            _ = v9
        }
        _ = v7
        _ = v8
    } else {
        v6 = "hello"
    }
    v10 := v4[42]
    var v11 int64 = 245
    v12 := GenS1{F0: -587, F1: 111, F2: 1}
    var v13 uint16 = 255
    v3 = 127
    for fi14 := 0; fi14 < 17; fi14++ {
        v6 = v10
        if true {
            var v15 int32 = int32(v3)
            switch int(v15) {
            case 9223372036854775807:
                v13 = v13
            case 128:
                var v16 uint32 = 0
                _ = v16
            default:
                fmt.Println(v13, fi14)
            }
            var v17 uint64 = 42
            _ = v15
            _ = v17
        }
        fmt.Println(v11, v6, fi14)
        _ = fi14
    }
    v13 = v13
    switch v10 {
    case "hello":
        var v18 uint = 2
        _ = v18
    }
    v6 = v10
    v4[2147483647] = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

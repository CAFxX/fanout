package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int32 = -288
    v2 := make(map[int]int)
    var v3 uint8 = 178
    v4 := make(map[int]int)
    var v5 string = "hello"
    var v6 uint16 = 65535
    var v7 int64 = 0
    var v8 uint32 = 2147483648
    var v9 int8 = -101
    if true {
        v6 = uint16(v3)
        v10 := GenS1{F0: 32767, F1: v8}
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                v12 := GenS1{F0: (v6 * uint16(v8)), F1: (v8 & (^(v8)))}
                _ = v12
            } else {
                v8 = 996
                if (true || true) {
                    v13 := GenS1{F0: v6, F1: (v8 ^ v8)}
                    _ = v13
                }
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v14 uint32 = (^(v8))
        if true {
            var v15 int32 = v1
            _ = v15
        } else {
            fmt.Println(v8, v7, v1)
            v9 = (v9 >> ((474) & 63))
        }
        _ = v14
    }
    var v16 uint8 = v3
    switch int(v9) {
    case 2147483647:
        var v17 string = "\n\t"
        _ = v17
    case 4294967295:
        fmt.Println(v16)
    default:
        var v18 int32 = -1
        _ = v18
    }
    var v19 uint = 111
    v7 = (^(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

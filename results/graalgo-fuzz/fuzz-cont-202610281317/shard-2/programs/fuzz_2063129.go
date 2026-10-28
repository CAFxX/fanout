package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 string = ("" + ("" + "\n\t"))
    var v2 uint64 = 42
    var v3 uint = 255
    var v4 int64 = 9223372036854775807
    var v5 uint8 = 1
    var v6 int16 = 42
    var v7 int16 = (v6 * v6)
    v6 = v6
    var v8 int64 = v4
    var v9 uint16 = 268
    v10 := GenS1{F0: (v4 % ((9223372036854775807) | 1))}
    if (!((0 > -1) && false)) {
        switch 32767 {
        case -410:
            v7 += v7
        case 1:
            v8 = (v8 | 9223372036854775807)
        default:
            var v11 uint8 = 127
            _ = v11
        }
        var v12 string = (v1 + v1)
        _ = v12
    }
    v13 := GenS1{F0: v8}
    v14 := GenS1{F0: (v4 & v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

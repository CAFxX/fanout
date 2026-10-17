package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    v1 := []int{9223372036854775807, 531, 32768}
    var v2 int = 127
    var v3 int16 = -787
    var v4 int8 = -1
    var v5 int64 = 1
    var v6 int32 = 42
    var v7 uint = 42
    fmt.Println(v2)
    fmt.Println(v5, v6)
    _ = len(v1)
    switch "" {
    case "hello":
        v8 := GenS1{F0: (-(v6)), F1: v7, F2: 146}
        _ = v8
    default:
        if (("hello" >= ("foo bar" + ("hello" + ("a" + ("\n\t" + "a"))))) || (!(v4 > (v4 * -103)))) {
            v7 = v7
            v5 += (v5 * (v5 / ((v5) | 1)))
            v9 := GenS1{F0: v6, F1: (v7 << ((v7) & 63)), F2: 2}
            _ = v9
        }
    }
    var v10 int32 = (v6 - (^(v6)))
    var v11 uint64 = 4294967295
    var v12 uint64 = v11
    if false {
        if ((!true) && false) {
            v13 := GenS1{F0: (v6 + (v6 ^ v6)), F1: v7, F2: 175}
            var v14 uint8 = 128
            _ = v13
            _ = v14
        } else {
            v6 -= -1
            var v15 uint32 = 4294967295
            _ = v15
        }
    } else {
        fmt.Println(v3, v7, v11)
        var v16 int8 = -1
        _ = v16
    }
    var v17 uint16 = 2
    v18 := GenS1{F0: int32(v12), F1: v7, F2: 185}
    if true {
        fmt.Println(v5, v6, v4)
    } else {
        var v19 string = ("foo bar" + "")
        _ = v19
    }
    var v20 string = ""
    v12 = uint64(v7)
    v10 += v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

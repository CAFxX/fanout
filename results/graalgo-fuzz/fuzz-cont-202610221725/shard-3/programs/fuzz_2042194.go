package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int8 }

func main() {
    var v1 uint64 = 92
    v2 := make(map[uint32]int)
    var v3 int64 = 42
    var v4 int32 = 42
    var v5 uint8 = 255
    var v6 float32 = 3.14159
    var v7 bool = (!(false && false))
    delete(v2, 423)
    var v8 string = "foo bar"
    v8 = ("hello" + (v8 + ((v8 + ("foo bar" + v8)) + v8)))
    if (186 != v5) {
        fmt.Println(v1, v7, v4)
        var v9 int16 = 497
        _ = v9
    } else {
        v10 := []int{128}
        _ = v10
    }
    if (!true) {
        v1 = v1
        fmt.Println(v6, v7)
        v11 := []int{0, 1}
        _ = v11
    }
    v12 := GenS1{F0: (v1 << ((uint(v5)) & 63)), F1: 1, F2: 127}
    switch v8 {
    case "":
        var v13 uint8 = (^(v5))
        _ = v13
    case "foo bar":
        if false {
            switch "a" {
            case "foo bar":
                v14 := GenS1{F0: v1, F1: 823, F2: -128}
                _ = v14
            case "":
                v1 += (v1 | 18446744073709551615)
            case "hello":
                v15 := GenS1{F0: 978, F1: 299, F2: 0}
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(fi16, v5, v8)
                v17 := GenS1{F0: v1, F1: 65535, F2: -128}
                var v18 int8 = 95
                _ = fi16
                _ = v17
                _ = v18
            }
            switch 127 {
            case 2147483647:
                v6 -= v6
            case 128:
                fmt.Println(v8)
            }
            var v19 uint8 = v5
            _ = v19
        }
    case "hello":
        v20 := GenS1{F0: (v1 - (^(v1))), F1: 31, F2: -1}
        _ = v20
    default:
        v4 = v4
    }
    v4 = (v4 << ((718) & 63))
    switch -1 {
    case 1:
        if (uint8(v1) < 0) {
            v8 = "foo bar"
            v3 += v3
            v21 := GenS1{F0: 0, F1: 255, F2: 1}
            _ = v21
        }
    default:
        var v22 uint16 = 42
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

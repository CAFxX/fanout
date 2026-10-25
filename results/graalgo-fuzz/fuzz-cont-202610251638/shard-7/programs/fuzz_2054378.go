package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }

func main() {
    var v1 uint64 = 0
    var v2 int32 = 0
    var v3 int64 = 0
    v4 := []string{"hello", ""}
    var v5 uint = 42
    var v6 uint32 = 2147483648
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 int32 = v2
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    switch 127 {
    case -1:
        v5 = 65535
    case 2147483647:
        if (true && (!(false && ((^(v7)) < v7)))) {
            fmt.Println(v6)
            if (42 >= 100) {
                var v8 uint64 = (v1 / (((v1 - v1)) | 1))
                _ = v8
            }
            v9 := GenS1{F0: (v1 / ((1) | 1)), F1: v1, F2: -9223372036854775808}
            _ = v9
        } else {
            v5 = (v5 % (((v5 * (v5 + (^(v5))))) | 1))
            fmt.Println(v5, v3, v6)
        }
    case 0:
        var v10 uint64 = v1
        _ = v10
    default:
        var v11 int8 = -128
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: v1, F1: 664, F2: int(v7)}
        v4 = append(v4, "hello")
        _ = fi12
        _ = v13
    }
    switch 769 {
    case 32768:
        var v14 int32 = (v7 + v7)
        _ = v14
    case 839:
        fmt.Println(v5)
    case 635:
        v7 = (v2 + v2)
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 = 1
        var v16 bool = (1000 != 32767)
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
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

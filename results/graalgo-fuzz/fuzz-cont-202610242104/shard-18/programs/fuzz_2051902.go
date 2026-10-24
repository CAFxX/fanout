package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int32 }
func genFn0() uint16 {
    if false {
        var v4 uint = 18446744073709551615
        for fi5 := 0; fi5 < 13; fi5++ {
            if false {
                fmt.Println(fi5, fi5)
                fmt.Println(v4, fi5)
                v4 = (v4 + 18446744073709551615)
            }
            if ((v4 / ((18446744073709551615) | 1)) <= 1) {
                var v6 uint8 = 128
                fmt.Println(fi5, v6)
                _ = v6
            }
            _ = fi5
        }
        _ = v4
    }
    var v7 uint32 = 426
    var v8 uint64 = 1
    var v9 int32 = 239
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 int8, p1 uint32) string {
    var v10 uint8 = 42
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            fmt.Println(fi11, p0, p1)
        } else {
            v12 := []int{-351}
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return "hello"
}

func main() {
    var v1 int16 = 387
    v2 := make(map[int]string)
    var v3 string = ("foo bar" + ("" + "a"))
    var v13 int64 = -9223372036854775808
    fmt.Println(v13, v13, v13)
    v14 := v2[-1]
    var v15 uint64 = 4294967295
    v13 = (v13 << ((236) & 63))
    fmt.Println(v15, v1)
    v16 := []int{1, -2147483648, 127}
    v17 := GenS1{F0: v1, F1: 255, F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }

func main() {
    var v1 float64 = 3.14159
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]bool)
    v4 := []string{(("a" + "hello") + "foo bar"), "foo bar", ("foo bar" + "")}
    v5 := []int{-429}
    v6 := make(map[int]int)
    if (true || (!true)) {
        if (!(-9223372036854775808 <= -741)) {
            v2 -= v2
            var v7 uint32 = (v2 / (((v2 ^ v2)) | 1))
            _ = v7
        }
    } else {
        v1 = v1
    }
    v2 += v2
    var v8 uint16 = 42
    if true {
        var v9 uint8 = 128
        _ = v9
    }
    fmt.Println(v1, v1, v8)
    v10 := GenS1{F0: -1, F1: v2, F2: -1}
    v2 = 2147483647
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = (2 == 0)
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
    }
    var v13 int8 = 127
    var v14 int32 = 42
    v15 := GenS1{F0: -1, F1: 42, F2: -2147483648}
    var v16 int32 = (-(v14))
    var v17 uint32 = (v2 ^ 0)
    v18 := GenS1{F0: v14, F1: (^(v17)), F2: 32767}
    if false {
        fmt.Println(v8, v1)
    } else {
        v16 = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

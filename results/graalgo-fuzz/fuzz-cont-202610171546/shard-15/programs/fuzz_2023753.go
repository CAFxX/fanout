package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
func genFn0(p0 bool, p1 int32) uint {
    p1 += p1
    if (!p0) {
        var v7 int64 = 838
        _ = v7
    } else {
        v8 := GenS1{F0: -1, F1: 18446744073709551615, F2: (364 >= 42)}
        _ = v8
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 42
    var v3 uint = 65535
    var v4 int16 = -32768
    v5 := []uint32{83, 327}
    var v6 uint16 = 32767
    _ = len(v1)
    v9 := GenS1{F0: 0, F1: (v3 >> (((v3 >> (((v3 + v3)) & 63))) & 63)), F2: ((false && false) || true)}
    var v10 bool = false
    if v10 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1["a"] = "foo bar"
    } else {
        v11 := v1[("hello" + "hello")]
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v4)
        v2 = v2
        _ = len(v1)
        _ = fi12
    }
    var v13 uint = v3
    v14 := GenS1{F0: 9223372036854775807, F1: (v13 % ((genFn0(v10, v2)) | 1)), F2: v10}
    fmt.Println(v2, v10, v2)
    v6 += 1
    var v15 float64 = ((0.0 - 1e10) + ((3.14159 / (1.5 - (-2.25 + 0.1))) - 1.5))
    var v16 int = 9223372036854775807
    v2 = (-(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

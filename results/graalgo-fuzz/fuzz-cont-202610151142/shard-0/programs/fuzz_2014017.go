package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 string = ("foo bar" + "hello")
    v4 := []uint32{1, 42, 1, 748}
    v5 := make(map[int]int)
    v6 := make(map[string]string)
    v7 := GenS1{F0: 898, F1: 681, F2: 42}
    v8 := v1[v3]
    var v9 int8 = 68
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint16 = 65535
        if len(v4) > 0 {
            v4[0] = 0
        }
        v11 = v11
        _ = fi10
        _ = v11
    }
    if (42 <= v2) {
        v9 = 42
    } else {
        if false {
            var v12 int16 = 32767
            _ = v12
        } else {
            _ = len(v4)
        }
        v3 = "hello"
        var v13 int64 = v2
        _ = v13
    }
    v14 := GenS1{F0: 18446744073709551615, F1: 1, F2: 65535}
    v8 = v3
    var v15 int16 = 364
    v9 = (-(v9))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

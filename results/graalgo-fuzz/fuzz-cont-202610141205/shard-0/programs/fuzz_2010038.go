package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint {
    var v7 int32 = 193
    v7 -= v7
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    v3 := make(map[int]string)
    var v4 int64 = 42
    var v5 uint8 = 2
    var v6 int = -9223372036854775808
    var v8 int32 = 363
    var v9 int = 1
    v5 += (v5 * (v5 + 240))
    fmt.Println(v4)
    v1[1] = ("" + (("a" + "foo bar") + ("foo bar" + "")))
    _ = len(v2)
    v10 := GenS1{F0: (false && true)}
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float64 = 1.5
        if len(v2) > 0 {
            v2[0] = fi11
        }
        switch (v6 % ((v9) | 1)) {
        case -288:
            v13 := v3[(-(v9))]
            _ = v13
        case -9223372036854775808:
            var v14 int8 = 1
            _ = v14
        case 127:
            var v15 float64 = (3.14159 - (3.14159 + v12))
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint64 = 42
        var v18 bool = true
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v9, v9, v5)
    fmt.Println(v6, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

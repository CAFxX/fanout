package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool, p1 int16) int16 {
    v6 := GenS1{F0: "a"}
    switch 4294967295 {
    case -1:
        p0 = true
    case 42:
        v7 := GenS1{F0: "hello"}
        _ = v7
    case 128:
        var v8 uint8 = 42
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return p1
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    var v3 int8 = 127
    v4 := []string{"a"}
    v5 := []int{302}
    var v10 uint = 0
    v3 += (v3 ^ (v3 & (v3 | v3)))
    var v11 uint = (v10 ^ 881)
    fmt.Println(v1, v11, v1)
    var v12 int32 = v1
    v13 := GenS1{F0: "a"}
    fmt.Println(v10, v10, v11)
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 = (v3 | v3)
        _ = fi14
    }
    v1 -= (v1 % ((v1) | 1))
    var v15 uint = (v10 * (v11 ^ 65535))
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int8 = -39
        v17 = v17
        v3 = v17
        _ = fi16
        _ = v17
    }
    switch -1 {
    case 4294967295:
        fmt.Println(v1)
    case 128:
        var v18 int8 = -4
        _ = v18
    case 32768:
        var v19 float64 = 0.0
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

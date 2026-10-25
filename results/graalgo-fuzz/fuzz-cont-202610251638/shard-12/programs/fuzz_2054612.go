package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint32 {
    v7 := []int{46, 0}
    var v8 uint = 139
    v8 += v8
    v8 = v8
    v8 += v8
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16) int8 {
    v9 := GenS1{F0: 42}
    p0 = p0
    _ = p0
    _ = v9
    return 100
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]uint32)
    var v3 uint = 532
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]bool)
    fmt.Println(v3)
    var v10 float64 = 1e10
    v11 := []int{4294967295}
    var v12 uint32 = 947
    var v13 string = ""
    if false {
        var v14 float32 = (0.1 - (1.5 + (-2.25 * (1e10 + 1e10))))
        if (((1e10 + -2.25) <= (1e10 / v10)) && false) {
            v3 = v3
            var v15 uint = (^(v3))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(v12, v13, v14)
                fmt.Println(v14)
                var v17 uint32 = (v12 ^ (v12 >> ((v3) & 63)))
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS1{F0: v12}
        _ = v14
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = (^(v4))
        var v21 uint64 = 224
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint16 = uint16(v12)
    _ = len(v11)
    fmt.Println(v12, v4)
    var v23 uint64 = 121
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

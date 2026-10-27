package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 int64, p1 int, p2 int32) bool {
    v5 := []int{}
    p2 = (-(p2))
    v6 := GenS1{F0: 2147483648, F1: -1, F2: p0}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 65535
        if true {
            v8 -= 65535
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 string) uint32 {
    var v9 float32 = (1.5 * ((1e10 + 3.14159) * 3.14159))
    v9 = 0.0
    fmt.Println(v9, v9)
    v9 = 3.14159
    _ = p0
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := []string{""}
    var v4 uint64 = 1
    _ = len(v3)
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (fi10 - 82) {
        case 128:
            if true {
                var v11 int64 = -1
                var v12 uint64 = 18446744073709551615
                _ = v11
                _ = v12
            }
        case -873:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi10
    }
    v1["\n\t"] = -204
    fmt.Println(v4)
    _ = len(v3)
    switch ("" + "a") {
    case "\n\t":
        fmt.Println(v4)
    case "a":
        var v13 uint32 = 904
        _ = v13
    default:
        _ = len(v2)
    }
    if (0.0 >= 0.1) {
        v14 := GenS1{F0: 879, F1: 1, F2: 9223372036854775807}
        var v15 uint32 = 375
        var v16 float32 = (1.5 + 1.5)
        _ = v14
        _ = v15
        _ = v16
    }
    v4 = (^(v4))
    var v17 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}

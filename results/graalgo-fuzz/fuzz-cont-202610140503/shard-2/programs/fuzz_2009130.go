package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8) int32 {
    var v7 float32 = ((0.0 / 1.0) / 1e10)
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint32 = 4294967295
        _ = fi8
        _ = v9
    }
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := []int{-1, 42}
    var v2 uint = 42
    var v3 int32 = 0
    var v4 bool = false
    var v5 int16 = -529
    v6 := make(map[uint32]uint32)
    v5 = v5
    if (!v4) {
        var v10 bool = (int16(v2) == 531)
        var v11 uint64 = 0
        _ = v10
        _ = v11
    } else {
        fmt.Println(v3, v4)
        switch "hello" {
        case "a":
            var v12 float32 = 0.1
            _ = v12
        case "foo bar":
            if true {
                v3 += v3
            } else {
                delete(v6, 2147483647)
            }
        case "":
            v13 := GenS1{F0: (-2.25 >= 1.5), F1: 21}
            _ = v13
        }
        var v14 uint = (^(v2))
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if len(v1) > 0 {
            v1[0] = (fi15 - 128)
        }
        v4 = false
        var v16 uint64 = 18446744073709551615
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 835
    v4 = (v4 || (-2.25 > (-2.25 * 0.1)))
    var v18 uint64 = 42
    switch (("" + "foo bar") + ("\n\t" + "")) {
    case "foo bar":
        v6[v17] = uint32(v18)
    case "a":
        if (231 > -2147483648) {
            var v19 bool = (42 >= 65535)
            var v20 uint16 = 403
            var v21 uint16 = 920
            _ = v19
            _ = v20
            _ = v21
        }
    }
    v17 = (v17 & (^(v17)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

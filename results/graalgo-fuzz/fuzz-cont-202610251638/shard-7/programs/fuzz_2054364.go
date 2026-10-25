package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 uint64) uint64 {
    if true {
        v5 := []int{4294967295, 128, 0}
        var v6 string = "a"
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 2, F1: 128}
        _ = v7
    }
    var v8 bool = true
    var v9 uint8 = 255
    if (735 != 4294967295) {
        for fi10 := 0; fi10 < 9; fi10++ {
            p2 -= (^(p2))
            fmt.Println(p0, p0)
            switch "" {
            case "a":
                var v11 uint16 = 0
                _ = v11
            }
            _ = fi10
        }
        var v12 bool = v8
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 int16) int16 {
    v13 := GenS2{F0: 2147483647, F1: 4294967295}
    v14 := GenS1{F0: 0, F1: 0}
    p0 -= (p0 + 1)
    fmt.Println(p0)
    _ = p0
    _ = v13
    _ = v14
    return (p0 >> ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]bool)
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS2{F0: -856, F1: v3}
        _ = fi15
        _ = v16
    }
    if (128 <= 6) {
        v4["a"] = (true || ((v3 / (((v3 * v3)) | 1)) < v3))
        v17 := []int{128, 703}
        switch v2 {
        case "hello":
            v1[4294967295] = "\n\t"
        default:
            v3 += 42
        }
        _ = v17
    }
    switch (v2 + v2) {
    case "\n\t":
        v18 := GenS1{F0: 65535, F1: 128}
        _ = v18
    default:
        var v19 uint = 65535
        _ = v19
    }
    v2 = "a"
    var v20 uint8 = 192
    switch -2147483648 {
    case 2147483647:
        v21 := []int{-615, 9223372036854775807, -477}
        _ = v21
    case 4294967295:
        for fi22 := 0; fi22 < 15; fi22++ {
            v20 = v20
            _ = fi22
        }
    case 706:
        if false {
            v23 := GenS1{F0: 82, F1: (v20 & v20)}
            var v24 uint = 42
            var v25 int64 = 1
            _ = v23
            _ = v24
            _ = v25
        }
    default:
        v26 := v4["hello"]
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            if (!(((((1e10 / -2.25) - 0.0) * (0.1 / (float64(1.5) - 0.0))) + 0.0) >= 1e10)) {
                v29 := v1[fi28]
                v3 -= v3
                _ = v29
            }
            _ = fi28
        }
        _ = fi27
    }
    fmt.Println(v20)
    var v30 float64 = (0.0 / 1e10)
    if false {
        v20 -= v20
        v31 := []int{}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
func genFn0(p0 int64, p1 string, p2 uint32) int8 {
    v7 := GenS1{F0: 42, F1: 65535, F2: (((p1 + p1) + p1) + "a")}
    p0 = p0
    p2 = p2
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint32, p1 uint64, p2 int64) int64 {
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: 1000, F1: 42, F2: "hello"}
        v11 := []int{(fi9 % ((-581) | 1)), (fi9 << ((4294967295) & 63)), fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    p2 += (p2 << ((18446744073709551615) & 63))
    var v12 uint8 = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return p2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := make(map[uint32]bool)
    var v4 int64 = 0
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    v13 := []int{-1, 9223372036854775807}
    fmt.Println(v4)
    if ((v2 & (v2 ^ v2)) >= (v2 & (v2 << ((4294967295) & 63)))) {
        fmt.Println(v2, v4)
        var v14 float64 = float64((0.1 + (-2.25 / (0.0 + 1e10))))
        v4 -= -1
        _ = v14
    } else {
        var v15 uint8 = 12
        fmt.Println(v15, v2)
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: -363, F1: 42, F2: "foo bar"}
        _ = fi16
        _ = v17
    }
    v13 = append(v13, -368)
    for fi18 := 0; fi18 < 13; fi18++ {
        if (32768 == fi18) {
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS1{F0: 32767, F1: 303, F2: "hello"}
                var v21 uint8 = 42
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        var v22 int8 = -1
        _ = fi18
        _ = v22
    }
    var v23 bool = false
    v24 := GenS1{F0: 1000, F1: 1, F2: ("\n\t" + "foo bar")}
    v2 = v2
    var v25 uint32 = 778
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

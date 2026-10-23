package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    v4 := GenS1{F0: 2147483648, F1: 2, F2: (p0 | p0)}
    v5 := GenS2{F0: 1, F1: 127}
    var v6 bool = false
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 89
}
func genFn1(p0 int8, p1 int8) string {
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        switch (fi7 ^ fi7) {
        case 9223372036854775807:
            var v8 int8 = p1
            _ = v8
        }
        _ = fi7
    }
    p1 = 100
    var v9 uint64 = 2
    _ = p0
    _ = p1
    _ = v9
    return ("a" + ("a" + "foo bar"))
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]string)
    v3 := []string{}
    v10 := GenS2{F0: 4294967295, F1: 72}
    var v11 uint32 = 2
    v12 := GenS2{F0: 763, F1: (v1 / (((-(v1))) | 1))}
    v1 = -127
    v1 -= (v1 ^ 1)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v3 = append(v3, "hello")
        _ = fi13
    }
    var v14 int64 = -1
    _ = len(v3)
    var v15 uint32 = (^(v11))
    v1 = -105
    if (uint64(v11) < 42) {
        if (127 < 128) {
            v16 := GenS3{F0: 9223372036854775807, F1: (((v14 | v14) > (-(v14))) && true), F2: "foo bar"}
            _ = v16
        }
        var v17 int32 = 2147483647
        delete(v2, v11)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}

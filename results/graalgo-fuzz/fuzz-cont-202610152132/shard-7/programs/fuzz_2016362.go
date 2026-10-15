package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint32, p1 uint32, p2 int16) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 -= p0
        _ = fi4
    }
    p2 = p2
    v5 := []int{42, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 & 42)
}
func genFn1(p0 int64) string {
    p0 = p0
    v6 := GenS1{F0: 529, F1: 607, F2: ("foo bar" + ("foo bar" + "a"))}
    _ = p0
    _ = v6
    return ""
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    v7 := []int{}
    fmt.Println(v3, v3)
    v8 := GenS2{F0: (^(v3)), F1: 453}
    for fi9 := 0; fi9 < 20; fi9++ {
        v1 = v1
        _ = fi9
    }
    v10 := GenS1{F0: 472, F1: 18446744073709551615, F2: "hello"}
    v11 := GenS1{F0: 509, F1: 18446744073709551615, F2: "foo bar"}
    if len(v7) > 0 {
        v7[0] = 32768
    }
    var v12 uint32 = 0
    var v13 bool = (128 <= 127)
    var v14 int8 = 95
    var v15 int8 = (v14 | -70)
    var v16 uint16 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}

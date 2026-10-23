package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 bool }
func genFn0(p0 string) uint16 {
    v7 := GenS1{F0: (p0 + p0), F1: 788}
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS2{F0: 42}
        _ = fi8
        _ = v9
    }
    v10 := []int{-423, -2147483648}
    _ = p0
    _ = v10
    _ = v7
    return 683
}
func genFn1(p0 string, p1 int64) uint16 {
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: p0, F1: -1}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: 1}
        _ = fi13
        _ = v14
    }
    switch (("hello" + (p0 + p0)) + p0) {
    case "a":
        v15 := []int{32768, 32767}
        _ = v15
    }
    var v16 int32 = 0
    var v17 uint64 = 2
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := []string{}
    var v2 uint64 = 1
    var v3 uint8 = 127
    v4 := []uint32{1, 2, 65535, uint32(v3)}
    v5 := make(map[int]string)
    var v6 uint8 = 127
    var v18 int64 = 0
    var v19 uint64 = (v2 << ((753) & 63))
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float32 = 3.14159
        _ = fi20
        _ = v21
    }
    _ = len(v5)
    v22 := GenS3{F0: 1, F1: 42, F2: (!true)}
    var v23 string = "foo bar"
    fmt.Println(v23, v23, v2)
    var v24 float64 = -2.25
    var v25 int = 146
    v24 = -2.25
    var v26 int8 = 42
    fmt.Println(v26, v3, v2)
    var v27 bool = true
    v27 = (!v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

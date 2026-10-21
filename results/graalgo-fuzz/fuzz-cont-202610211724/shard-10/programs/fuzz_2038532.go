package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0(p0 string, p1 uint32) string {
    v5 := GenS1{F0: -157, F1: (!true), F2: 1}
    var v6 float32 = float32(1.5)
    v7 := []int{-406, 96, 0, 0}
    p0 = "\n\t"
    if (((p0 + p0) + p0) != p0) {
        p0 = (p0 + "hello")
        var v8 uint = 65535
        p0 = (p0 + ((p0 + (p0 + p0)) + "hello"))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 bool = false
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1.5
    var v9 int16 = 0
    _ = len(v3)
    var v10 uint32 = 4294967295
    var v11 uint64 = 0
    _ = len(v3)
    var v12 int = -2147483648
    switch (-(v12)) {
    case -2147483648:
        var v13 int8 = 100
        _ = v13
    case 127:
        var v14 bool = false
        _ = v14
    }
    v10 = 235
    var v15 int16 = 1000
    for fi16 := 0; fi16 < 10; fi16++ {
        v15 -= (v15 / (((v15 >> ((4294967295) & 63))) | 1))
        var v17 uint8 = 1
        v18 := GenS2{F0: 9223372036854775807, F1: v10}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v9, v11)
    if (!v1) {
        switch v2 {
        case "hello":
            v19 := GenS3{F0: 621}
            _ = v19
        case "":
            var v20 uint32 = v10
            _ = v20
        case "a":
            var v21 int = v12
            _ = v21
        }
    }
    if v1 {
        v22 := GenS4{F0: 2, F1: (v11 / ((184) | 1))}
        v9 = v9
        _ = v22
    }
    v23 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}

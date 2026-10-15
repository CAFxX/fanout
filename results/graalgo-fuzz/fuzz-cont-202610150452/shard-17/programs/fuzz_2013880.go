package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 int; F1 string; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0() int {
    var v5 bool = (!true)
    v6 := []int{}
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 uint32) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 bool = false
        var v9 uint16 = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 = p1
    v10 := GenS1{F0: 1, F1: (("\n\t" + "hello") + ("" + "a")), F2: (0 != 65535)}
    if true {
        var v11 int64 = p0
        var v12 int32 = 2147483647
        v13 := GenS2{F0: ("hello" + "hello"), F1: 15, F2: -332}
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 float64 = 1e10
        p0 += (p0 / ((p0) | 1))
        var v15 int32 = 42
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 0
    var v3 float64 = (-2.25 + -2.25)
    var v4 int64 = -1
    v16 := []int{32767, 127}
    v17 := GenS2{F0: ("\n\t" + "hello"), F1: 42, F2: 32767}
    v16 = append(v16, -509)
    var v18 uint = 2
    v4 += (v4 << ((v18) & 63))
    _ = len(v16)
    v19 := GenS3{F0: genFn0(), F1: ("hello" + "a"), F2: v2}
    var v20 bool = false
    v4 = (v4 ^ (v4 >> (((v18 + (v18 | (v18 / (((v18 << ((245) & 63))) | 1))))) & 63)))
    if true {
        fmt.Println(v20, v20, v20)
        v18 = 65535
        v21 := GenS4{F0: (v4 - (^(v4))), F1: 65535}
        _ = v21
    } else {
        v20 = (2 <= 2147483648)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

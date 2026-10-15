package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() uint16 {
    var v4 int = 128
    var v5 int16 = -1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint, p2 uint8) int {
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            p2 = 53
            var v8 bool = (!(42 <= -570))
            v9 := GenS1{F0: 1}
            var v10 uint = 65535
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        fmt.Println(p1, p1)
        _ = fi6
    }
    v11 := GenS1{F0: 4294967295}
    p2 += p2
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return int(p0)
}

func main() {
    var v1 int8 = -1
    var v2 float32 = 1.5
    var v3 uint64 = 18446744073709551615
    v13 := GenS2{F0: 2147483647}
    switch "\n\t" {
    case "":
        fmt.Println(v1, v3, v2)
    }
    var v14 float32 = v2
    var v15 uint64 = v3
    v16 := GenS3{F0: 241, F1: 323}
    var v17 string = ("" + "foo bar")
    v18 := []int{9223372036854775807, 0, 42}
    var v19 float32 = v14
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int32 = -424
        v18 = append(v18, 1)
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int32 = 42
            fmt.Println(fi20, v19, v14)
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v15, v15, v17)
    var v24 bool = false
    v25 := GenS2{F0: 65535}
    v17 = v17
    var v26 string = v17
    if ((v24 && v24) || (!true)) {
        if len(v18) > 0 {
            v18[0] = 32768
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}

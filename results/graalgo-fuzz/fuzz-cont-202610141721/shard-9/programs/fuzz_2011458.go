package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 uint, p1 uint32) bool {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p0, fi6)
        _ = fi6
    }
    _ = p0
    _ = p1
    return false
}
func genFn1() string {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    var v8 int16 = 32767
    var v9 int8 = 29
    v9 -= 0
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 int32 = -906
    var v2 bool = false
    v3 := []string{"\n\t", "", "\n\t"}
    var v4 uint64 = 42
    v5 := make(map[uint32]int)
    v4 = 42
    var v10 uint8 = 0
    var v11 int16 = -1
    v11 = (v11 + v11)
    if (42 < 42) {
        v12 := GenS1{F0: 59, F1: 1}
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS2{F0: (v10 / ((v10) | 1)), F1: 936}
            delete(v5, 42)
            v15 := GenS3{F0: v1, F1: 32767}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(fi16, v11)
            _ = fi16
        }
        if false {
            v17 := GenS3{F0: 42, F1: (v11 - v11)}
            _ = v17
        }
        fmt.Println(v2)
    }
    if v2 {
        v2 = (v2 || (v4 > v4))
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "":
        var v18 int32 = (v1 % ((v1) | 1))
        _ = v18
    }
    var v19 float32 = float32((0.0 - -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32, p1 int32) bool {
    var v5 uint64 = 1
    v6 := []int{9223372036854775807, -1}
    v7 := GenS1{F0: -983, F1: (v5 | v5)}
    if true {
        p1 = 2147483647
        if false {
            v8 := GenS2{F0: 130, F1: -27}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: 165, F1: -14}
            var v11 uint16 = 42
            v6 = append(v6, fi9)
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 += (^(p0))
    }
    if false {
        p0 = (p0 / ((0) | 1))
        p1 -= -1
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (-112 > 74)
}

func main() {
    var v1 bool = (910 < 42)
    var v2 bool = ("hello" < ("a" + ("\n\t" + (("" + "a") + "hello"))))
    v3 := make(map[string]bool)
    v4 := []int{2147483647, 436}
    var v12 int8 = 42
    if ((-79 > 42) || v2) {
        for fi13 := 0; fi13 < 15; fi13++ {
            _ = len(v3)
            fmt.Println(v2)
            fmt.Println(v12)
            _ = fi13
        }
    } else {
        v14 := GenS2{F0: 0, F1: (v12 - (v12 >> ((4294967295) & 63)))}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v14
    }
    _ = len(v4)
    v1 = (1e10 >= 3.14159)
    var v15 int = -897
    v2 = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v16 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

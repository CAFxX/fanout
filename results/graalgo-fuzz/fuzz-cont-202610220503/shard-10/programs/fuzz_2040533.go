package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 int32 }
type GenS4 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 int32, p1 int32, p2 int8) bool {
    p2 += -1
    v7 := GenS1{F0: 1}
    if true {
        fmt.Println(p2)
        v8 := []int{-139, -9223372036854775808, -94}
        _ = v8
    }
    v9 := GenS2{F0: 2147483647, F1: 815, F2: 1}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("\n\t" + "")}
    var v3 int32 = 2147483647
    var v4 int8 = 127
    var v5 uint32 = 2
    var v6 int64 = 600
    v4 = v4
    v10 := GenS3{F0: v3, F1: (v6 << ((2) & 63)), F2: v3}
    var v11 uint = 0
    var v12 bool = ((4294967295 < 1) && true)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint16 = 39
        v12 = false
        switch fi13 {
        case -2147483648:
            var v15 int16 = 0
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 uint64 = 42
    v6 += v6
    v17 := GenS4{F0: (v11 | uint(v5)), F1: (v11 * uint(v16)), F2: v12}
    var v18 string = (("" + "hello") + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

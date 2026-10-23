package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 int64 }
func genFn0() int64 {
    var v4 uint8 = 106
    var v5 uint16 = 0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    var v3 int32 = 2147483647
    v1 -= -2.25
    var v6 uint64 = 2
    fmt.Println(v3)
    switch -88 {
    case -531:
        v3 -= (v3 | (-(v3)))
    }
    var v7 uint = 65535
    v8 := GenS1{F0: (("hello" + ("\n\t" + ("hello" + ""))) + ("" + "\n\t"))}
    if false {
        if true {
            v7 -= (v7 * (^(v7)))
        } else {
            var v9 int32 = (^(v3))
            v10 := []int{32767, -751}
            v11 := GenS2{F0: ("hello" + "hello"), F1: 912, F2: 0}
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 bool = (v1 != v1)
        _ = v12
    }
    var v13 uint64 = v6
    v14 := GenS3{F0: 42}
    v15 := []int{2147483647, 32767, 621, 42}
    var v16 int8 = 1
    v3 = -77
    fmt.Println(v6, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}

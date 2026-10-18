package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }

func main() {
    var v1 string = (("a" + (("" + "foo bar") + (("hello" + "hello") + ""))) + "a")
    v2 := []int{}
    v3 := []string{v1, (v1 + v1), ""}
    v4 := []string{(v1 + v1)}
    v5 := make(map[int]int)
    var v6 int8 = 0
    var v7 uint16 = 42
    if true {
        if (!(false || (-618 > 0))) {
            v1 = v1
            v8 := GenS1{F0: v1, F1: 456}
            fmt.Println(v1, v7, v7)
            _ = v8
        }
        var v9 uint64 = 42
        v10 := GenS2{F0: false, F1: 42, F2: 42}
        _ = v10
        _ = v9
    } else {
        fmt.Println(v7)
        v6 = v6
    }
    var v11 uint16 = 207
    if false {
        if true {
            v12 := GenS2{F0: (2 == uint64(v7)), F1: uint8(v7), F2: -52}
            _ = v12
        }
        v13 := GenS2{F0: (true || false), F1: 0, F2: 63}
        v1 = v1
        _ = v13
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) > 0 {
            v4[0] = v1
        }
    } else {
        v6 += v6
        var v14 bool = (65535 == 4294967295)
        _ = v14
    }
    v15 := v5[0]
    fmt.Println(v6)
    var v16 int16 = 1
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = 42
        v19 := GenS2{F0: false, F1: 2, F2: (v16 / (((v16 | v16)) | 1))}
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

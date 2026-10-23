package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    var v4 int8 = 111
    var v5 uint32 = 2147483647
    fmt.Println(v5)
    if false {
        var v6 uint32 = (v5 >> ((0) & 63))
        v1 -= v1
        _ = v6
    }
    fmt.Println(v2, v5, v5)
    var v7 bool = ((v4 % ((v4) | 1)) == v4)
    v8 := GenS1{F0: 42, F1: 0, F2: 193}
    if true {
        v9 := []int{1, -985}
        v10 := GenS2{F0: "a", F1: 25}
        delete(v3, 2147483647)
        _ = v10
        _ = v9
    } else {
        v1 = v1
    }
    fmt.Println(v5)
    if v7 {
        switch -2147483648 {
        case 4294967295:
            fmt.Println(v2)
        case 32767:
            v7 = v7
        }
    } else {
        v11 := GenS1{F0: 127, F1: -32768, F2: 1000}
        v12 := GenS1{F0: 193, F1: 1000, F2: 0}
        _ = v11
        _ = v12
    }
    _ = len(v3)
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: "hello", F1: 2}
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 1
    if ((!false) || false) {
        var v16 float32 = 1.5
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := []int{fi17, fi17, (fi17 & 42)}
            _ = fi17
            _ = v18
        }
        var v19 uint8 = (^(v15))
        _ = v16
        _ = v19
    } else {
        var v20 uint32 = (v5 / ((v5) | 1))
        v20 = (v20 >> ((18446744073709551615) & 63))
        var v21 uint16 = 674
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }

func main() {
    var v1 uint = 512
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    if false {
        v4 := []int{}
        var v5 uint8 = 42
        _ = v4
        _ = v5
    }
    switch 42 {
    case 128:
        switch -2147483648 {
        case 0:
            for fi6 := 0; fi6 < 14; fi6++ {
                _ = len(v3)
                v1 -= 1
                v3[(^(fi6))] = "hello"
                _ = fi6
            }
        case 127:
            v7 := GenS1{F0: -316}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 816}
        var v10 uint = v1
        v11 := GenS2{F0: -128, F1: 762, F2: 0}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{203}
    v13 := GenS3{F0: 273, F1: -1, F2: 2}
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    v2 = 2
    v3[32768] = "hello"
    var v14 uint = 6
    for fi15 := 0; fi15 < 12; fi15++ {
        if len(v12) > 0 {
            v12[0] = fi15
        }
        _ = fi15
    }
    if false {
        delete(v3, -567)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    } else {
        v16 := GenS3{F0: 492, F1: -1, F2: 65535}
        _ = v16
    }
    var v17 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}

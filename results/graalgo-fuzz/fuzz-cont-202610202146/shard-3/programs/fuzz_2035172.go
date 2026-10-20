package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint32 = 341
    var v2 int32 = 0
    var v3 string = ("a" + "")
    for fi4 := 0; fi4 < 9; fi4++ {
        v2 = 42
        fmt.Println(v1, v3)
        _ = fi4
    }
    if (2 == 42) {
        var v5 int8 = 100
        _ = v5
    }
    fmt.Println(v1, v2)
    var v6 uint16 = 168
    v7 := []int{-9223372036854775808, 4294967295, 0, 2147483647}
    var v8 int32 = (v2 & -1)
    if (-9223372036854775808 == -9223372036854775808) {
        v2 = (^(v2))
        if false {
            v7 = append(v7, 127)
            var v9 float32 = 0.0
            _ = v9
        }
    } else {
        var v10 uint32 = v1
        v11 := GenS1{F0: -1, F1: 42, F2: 171}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v14 uint64 = 18446744073709551615
            if true {
                v3 = "\n\t"
                if false {
                    v15 := GenS2{F0: (42 == 42), F1: v14}
                    v6 -= 303
                    _ = v15
                } else {
                    v3 = ("a" + "")
                }
            }
            if len(v7) > 0 {
                v7[0] = (fi12 >> ((315) & 63))
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v2 = 42
        var v17 uint32 = (v1 % (((v1 / ((v1) | 1))) | 1))
        v6 -= (v6 / ((v6) | 1))
        _ = fi16
        _ = v17
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (!(!false)) {
        var v18 int64 = -1
        _ = v18
    }
    _ = len(v7)
    v2 = v2
    var v19 string = (v3 + v3)
    var v20 int8 = -1
    var v21 int16 = 32767
    var v22 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}

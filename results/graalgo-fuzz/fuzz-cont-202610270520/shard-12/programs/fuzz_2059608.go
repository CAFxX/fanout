package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0() string {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint64 = 242
        var v6 float32 = (1.5 + float32(3.14159))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint = 1
    v8 := []int{-269, 32767, -164, -2147483648}
    _ = v7
    _ = v8
    return (("\n\t" + "hello") + "foo bar")
}

func main() {
    v1 := []int{127, 349, -2147483648}
    var v2 uint16 = 0
    v3 := make(map[uint32]uint32)
    v9 := GenS1{F0: 42, F1: 0}
    v2 += (v2 ^ v2)
    var v10 int16 = 32767
    if (false && false) {
        v11 := GenS1{F0: -68, F1: 2147483647}
        fmt.Println(v2, v2)
        _ = v11
    } else {
        v1 = append(v1, 0)
    }
    if (42 != 0) {
        v3[42] = 2
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS2{F0: "a"}
                fmt.Println(fi12, fi13, fi13)
                if (-112 > -9223372036854775808) {
                    var v15 int8 = 81
                    _ = v15
                } else {
                    var v16 bool = (292 > 65535)
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    } else {
        var v17 uint = 1
        if (!(-100 > 0)) {
            var v18 int = 128
            _ = v18
        } else {
            var v19 int8 = 42
            _ = v19
        }
        _ = v17
    }
    _ = len(v1)
    v20 := GenS3{F0: 1, F1: 42, F2: 2147483647}
    var v21 int32 = 0
    fmt.Println(v10)
    var v22 int = 32768
    v21 = v21
    var v23 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}

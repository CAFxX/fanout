package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 uint64 = 1
    v2 := []string{"hello", "foo bar"}
    var v3 float32 = float32(3.14159)
    v4 := make(map[int]string)
    var v5 bool = false
    v1 += 107
    var v6 float32 = (((v3 / v3) * v3) / 1.5)
    fmt.Println(v5, v5, v3)
    fmt.Println(v5, v1)
    _ = len(v2)
    if false {
        if (689 == 42) {
            var v7 uint16 = 963
            _ = v7
        }
    } else {
        v8 := GenS1{F0: 42}
        _ = v8
    }
    if ((true || v5) || v5) {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                v1 += 1
                _ = len(v2)
                v11 := GenS2{F0: 127}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                v1 = 0
                v3 = v6
                _ = fi12
            }
            _ = fi9
        }
        if true {
            v13 := GenS3{F0: 32767, F1: 0, F2: 128}
            _ = v13
        }
        var v14 int16 = -762
        _ = v14
    } else {
        v15 := GenS4{F0: 9223372036854775807, F1: -451, F2: -32768}
        var v16 int64 = -191
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

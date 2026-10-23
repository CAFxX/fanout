package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 58
    var v3 bool = false
    var v4 string = "a"
    v5 := []uint32{2147483647}
    v6 := make(map[string]int)
    fmt.Println(v3, v4, v4)
    if v3 {
        switch (v4 + ("\n\t" + v4)) {
        case "a":
            v7 := GenS1{F0: -600, F1: 9223372036854775807, F2: 65535}
            _ = v7
        }
        switch v4 {
        case "foo bar":
            var v8 uint8 = (v1 / ((v1) | 1))
            _ = v8
        default:
            v9 := GenS2{F0: "a", F1: -90, F2: 7}
            _ = v9
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        v2 -= (v2 + 196)
        v2 += v2
    }
    var v10 bool = (142 <= 32767)
    if v3 {
        var v11 int32 = -1
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint64 = v2
            if false {
                var v14 uint64 = (v13 % ((387) | 1))
                _ = v14
            } else {
                v4 = v4
            }
            if len(v5) > 0 {
                v5[0] = 2
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        var v15 int8 = 117
        v16 := GenS3{F0: v4, F1: (^(v15))}
        var v17 bool = v10
        _ = v15
        _ = v16
        _ = v17
    }
    _ = len(v6)
    v18 := GenS3{F0: (v4 + v4), F1: 100}
    var v19 string = v4
    if false {
        var v20 float32 = (((float32(0.0) * float32(1e10)) / (1e10 / float32((float64(0.1) * 0.1)))) / 0.1)
        _ = len(v6)
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 int32; F2 int32 }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 int8; F2 uint32 }
type GenS6 struct { F0 int }

func main() {
    var v1 uint32 = 762
    var v2 int16 = 313
    v3 := make(map[uint32]string)
    var v4 float64 = 0.0
    var v5 float32 = 1e10
    fmt.Println(v4, v5, v4)
    v4 = v4
    v6 := []int{32768, -1}
    v7 := GenS1{F0: (("a" + "a") + "foo bar"), F1: "\n\t"}
    if (0 >= 128) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v3[(v1 ^ v1)] = "\n\t"
        v8 := GenS2{F0: 189, F1: -1, F2: 42}
        _ = v8
    }
    v9 := GenS2{F0: 1, F1: -2147483648, F2: -2147483648}
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS3{F0: 2147483647, F1: (0 == 343), F2: 32767}
        fmt.Println(fi10, v4, v5)
        _ = fi10
        _ = v11
    }
    if (true || ("foo bar" != "")) {
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v1, v5)
            v3[v1] = "a"
            if false {
                switch (("a" + "") + "") {
                case "hello":
                    v1 += (^(v1))
                default:
                    v13 := GenS2{F0: 210, F1: 42, F2: -278}
                    _ = v13
                }
            }
            _ = fi12
        }
    } else {
        var v14 uint = 42
        _ = v14
    }
    var v15 uint32 = (v1 / (((v1 / ((v1) | 1))) | 1))
    v16 := GenS4{F0: 579, F1: 796}
    if false {
        v1 = (v15 / ((813) | 1))
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        var v17 uint64 = 0
        _ = v17
    }
    v18 := GenS5{F0: 65535, F1: -1, F2: v15}
    v19 := GenS4{F0: -457, F1: 42}
    fmt.Println(v5)
    var v20 uint32 = 2
    v21 := GenS6{F0: 680}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

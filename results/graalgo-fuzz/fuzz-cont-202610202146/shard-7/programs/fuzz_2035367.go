package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }

func main() {
    v1 := []string{(((("\n\t" + "foo bar") + "hello") + ("foo bar" + ("" + ("" + "\n\t")))) + "hello"), ("foo bar" + (("\n\t" + "hello") + "foo bar")), "foo bar"}
    v2 := make(map[int]string)
    var v3 string = "\n\t"
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int = (^(fi4))
            v3 = "a"
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v2[32768] = ((v3 + v3) + v3)
    var v7 int8 = -1
    v7 = (v7 << ((667) & 63))
    fmt.Println(v3, v3)
    fmt.Println(v3, v7)
    var v8 uint16 = 0
    if true {
        var v9 uint32 = 1
        v8 += v8
        _ = v9
    } else {
        var v10 int32 = 0
        var v11 float32 = (1e10 / 3.14159)
        _ = v10
        _ = v11
    }
    if (65535 == 944) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v12 := GenS1{F0: 128, F1: 402}
        _ = v12
    } else {
        var v13 uint = 2
        v14 := GenS2{F0: v7, F1: -417, F2: "foo bar"}
        v15 := GenS3{F0: 650, F1: (v7 ^ (^(v7))), F2: (v13 & (v13 | (v13 + v13)))}
        _ = v13
        _ = v14
        _ = v15
    }
    v7 += (v7 % ((v7) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}

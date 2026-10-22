package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    v3 := make(map[int]uint32)
    var v4 uint16 = 2
    v4 = (v4 ^ (v4 | v4))
    switch 127 {
    case 127:
        var v5 uint8 = 60
        _ = v5
    case -20:
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{fi6, (fi6 / (((fi6 | (fi6 >> ((226) & 63)))) | 1)), (-(fi6))}
            v4 -= (^(v4))
            _ = fi6
            _ = v7
        }
    }
    v8 := []int{-2147483648, 134, -416}
    v1[""] = "foo bar"
    fmt.Println(v2, v2)
    for fi9 := 0; fi9 < 13; fi9++ {
        v4 = v4
        v10 := GenS1{F0: -2147483648, F1: -500}
        var v11 uint16 = 255
        _ = fi9
        _ = v10
        _ = v11
    }
    delete(v1, ("foo bar" + ("hello" + ("a" + "\n\t"))))
    v12 := GenS1{F0: 0, F1: -878}
    for fi13 := 0; fi13 < 20; fi13++ {
        v2 = (v2 % ((v2) | 1))
        v14 := GenS1{F0: 971, F1: -693}
        v4 = 961
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 52, F1: (464 < 4294967295), F2: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

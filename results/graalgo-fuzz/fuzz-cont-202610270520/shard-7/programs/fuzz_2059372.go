package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    var v3 int = 1
    v4 := make(map[uint32]bool)
    v5 := []string{"foo bar", ("hello" + ""), ("foo bar" + "hello"), ("hello" + "a")}
    _ = len(v5)
    v6 := GenS1{F0: 42, F1: true, F2: 676}
    v7 := GenS2{F0: v3}
    var v8 uint64 = 153
    v8 = v8
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS3{F0: (("" + (("foo bar" + ("a" + ("" + ""))) + "")) + "a")}
            v8 += 2
            _ = fi10
            _ = v11
        }
        v2 = 752
        _ = fi9
    }
    v8 += (v8 >> ((255) & 63))
    switch (v3 << ((13) & 63)) {
    case 895:
        var v12 uint = 1
        _ = v12
    case 761:
        v3 -= (^(v3))
    case -1:
        var v13 int = v3
        _ = v13
    default:
        _ = len(v5)
    }
    var v14 float32 = float32(0.0)
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS3{F0: "\n\t"}
        _ = fi15
        _ = v16
    }
    v2 = 0
    v17 := GenS3{F0: "hello"}
    v18 := GenS3{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

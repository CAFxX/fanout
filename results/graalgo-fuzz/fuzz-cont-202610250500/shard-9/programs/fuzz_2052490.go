package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 0
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    v6 := GenS1{F0: (!((1 == 42) && (true || (!true)))), F1: 966, F2: -2147483648}
    v2 += (v2 << ((2) & 63))
    v7 := []int{}
    v8 := GenS2{F0: 4294967295, F1: (!("hello" >= "a"))}
    _ = len(v7)
    v1 = ("a" + v1)
    var v9 float64 = (3.14159 + 1.5)
    v10 := GenS3{F0: 65535, F1: 9223372036854775807, F2: 2147483648}
    v11 := GenS2{F0: 32767, F1: (false && false)}
    switch "hello" {
    case "hello":
        v12 := GenS4{F0: v2, F1: v2}
        _ = v12
    }
    if (("" + v1) == "foo bar") {
        switch v1 {
        case "\n\t":
            v2 = 0
        case "":
            var v13 uint64 = v2
            _ = v13
        }
        v9 += v9
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS4{F0: v2, F1: v2}
            _ = fi14
            _ = v15
        }
    }
    v4[589] = 4294967295
    v1 = v1
    if (!true) {
        v16 := v4[77]
        fmt.Println(v16)
        fmt.Println(v1)
        _ = v16
    }
    _ = len(v7)
    delete(v5, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

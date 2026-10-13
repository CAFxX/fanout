package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() int64 {
    switch (("" + "hello") + ((("a" + "hello") + ("hello" + "hello")) + "hello")) {
    case "":
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: 252, F1: int8(fi4)}
            _ = fi4
            _ = v5
        }
    case "hello":
        v6 := GenS2{F0: 111, F1: 65535}
        _ = v6
    case "foo bar":
        var v7 uint8 = 255
        _ = v7
    default:
        v8 := []int{-9223372036854775808, 127}
        _ = v8
    }
    v9 := GenS2{F0: 0, F1: 2}
    var v10 uint64 = 434
    _ = v10
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    v3 := []int{78, 127}
    var v11 int64 = 1
    for fi12 := 0; fi12 < 16; fi12++ {
        switch (fi12 ^ -937) {
        case -364:
            v11 -= (v11 * v11)
        }
        v13 := v2[("hello" + "a")]
        var v14 bool = v13
        _ = fi12
        _ = v13
        _ = v14
    }
    v1[42] = 128
    if false {
        v15 := GenS3{F0: 1, F1: 42, F2: 32767}
        fmt.Println(v11)
        v11 = genFn0()
        _ = v15
    }
    var v16 int8 = 127
    var v17 float32 = 1.5
    v18 := GenS4{F0: "foo bar", F1: (v16 + 90)}
    _ = len(v2)
    var v19 uint16 = 747
    fmt.Println(v16, v16)
    var v20 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}

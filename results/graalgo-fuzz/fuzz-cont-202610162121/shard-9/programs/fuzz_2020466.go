package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 uint8) uint8 {
    switch "\n\t" {
    case "foo bar":
        p1 = (p1 + p1)
    case "a":
        v6 := GenS1{F0: 467}
        _ = v6
    case "hello":
        var v7 int32 = -726
        _ = v7
    default:
        for fi8 := 0; fi8 < 14; fi8++ {
            p1 += (p1 * p1)
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        switch (("" + "\n\t") + "a") {
        case "a":
            var v10 int8 = 36
            _ = v10
        case "foo bar":
            v11 := GenS2{F0: 128, F1: 843}
            _ = v11
        case "\n\t":
            v12 := GenS3{F0: 30}
            _ = v12
        }
        v13 := []int{(fi9 | (-(fi9))), fi9}
        v14 := GenS3{F0: 120}
        _ = fi9
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 2
    v3 := make(map[int]bool)
    var v4 uint16 = 2
    var v5 uint = 65535
    var v15 uint64 = 352
    fmt.Println(v4, v15)
    fmt.Println(v5, v15)
    v3[-1] = false
    var v16 uint = 4294967295
    v17 := []int{127, int(v5)}
    v18 := GenS4{F0: (0 > (v4 - (v4 % ((425) | 1)))), F1: 0}
    v15 = (v15 / ((uint64(v5)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

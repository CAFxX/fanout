package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 int8) int8 {
    var v4 int64 = -1
    var v5 bool = false
    v6 := []int{-504, -548, 2147483647}
    v7 := GenS1{F0: (^(p0))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (p1 * p1)
}

func main() {
    var v1 uint8 = 121
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v1 = v1
    switch "hello" {
    case "":
        v1 = v1
    case "a":
        switch ("hello" + "") {
        case "hello":
            var v8 int32 = 786
            _ = v8
        case "\n\t":
            var v9 string = "a"
            _ = v9
        default:
            v1 = v1
        }
    case "foo bar":
        v10 := GenS1{F0: 65535}
        _ = v10
    }
    var v11 uint64 = 48
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 string = "\n\t"
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v11)
    v14 := GenS1{F0: 42}
    if (!true) {
        v1 = 21
        v15 := GenS2{F0: "\n\t", F1: 42}
        v11 += (v11 ^ 2)
        _ = v15
    }
    v1 = v1
    v16 := GenS2{F0: "foo bar", F1: 0}
    v1 = v1
    switch "hello" {
    case "a":
        if true {
            v17 := GenS3{F0: 65535, F1: 859, F2: 4294967295}
            for fi18 := 0; fi18 < 3; fi18++ {
                fmt.Println(v11)
                var v19 int8 = genFn0(502, 127)
                _ = fi18
                _ = v19
            }
            _ = v17
        } else {
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(v1, fi20, fi20)
                _ = fi20
            }
        }
    case "foo bar":
        v2 = append(v2, 2147483647)
    case "\n\t":
        var v21 int = -968
        _ = v21
    }
    v22 := GenS3{F0: 255, F1: 1, F2: 2}
    v1 = (v1 + (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}

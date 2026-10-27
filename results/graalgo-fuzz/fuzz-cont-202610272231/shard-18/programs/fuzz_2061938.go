package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 int8, p1 uint16, p2 int64) uint64 {
    p1 = p1
    v4 := GenS1{F0: 65535}
    fmt.Println(p0, p0, p0)
    switch int(p1) {
    case -351:
        var v5 uint8 = 128
        _ = v5
    }
    var v6 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 1
}

func main() {
    v1 := []int{}
    var v2 uint16 = 42
    var v3 int = 32767
    fmt.Println(v3, v3)
    v2 += (^(v2))
    var v7 string = ("hello" + ((("foo bar" + "") + "") + "foo bar"))
    v3 = -1
    v2 -= v2
    switch v3 {
    case 32767:
        v8 := GenS2{F0: -1, F1: 127}
        _ = v8
    case -1:
        v3 += (v3 % ((v3) | 1))
    }
    v1 = append(v1, 789)
    v9 := GenS3{F0: 948, F1: (("a" + (("foo bar" + v7) + (("" + ("hello" + "hello")) + "a"))) + v7)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}

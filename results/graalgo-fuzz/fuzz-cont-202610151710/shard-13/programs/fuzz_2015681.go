package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 bool, p2 int32) uint8 {
    p0 = p0
    v7 := []int{735, 1, 2147483647}
    if len(v7) > 0 {
        v7[0] = 42
    }
    p2 += (p2 >> ((977) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint, p2 bool) int8 {
    v8 := []int{(p0 % ((p0) | 1)), 42, p0, (^(p0))}
    p0 = p0
    switch "\n\t" {
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: true}
            p1 = p1
            _ = fi9
            _ = v10
        }
    default:
        fmt.Println(p2)
    }
    if p2 {
        v11 := GenS1{F0: false}
        _ = v11
    }
    switch ("" + "hello") {
    case "\n\t":
        v12 := GenS1{F0: (773 >= 287)}
        _ = v12
    case "a":
        var v13 int = p0
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{1, 756, 428}
    v3 := []int{}
    v4 := make(map[string]uint32)
    var v5 uint16 = 0
    v6 := []string{((("foo bar" + "foo bar") + "\n\t") + ("hello" + "foo bar")), ("\n\t" + ("foo bar" + "a")), "a"}
    v14 := GenS2{F0: genFn1(-270, 4294967295, true), F1: 65535}
    v15 := GenS1{F0: (true && true)}
    var v16 int = -549
    v17 := GenS1{F0: (true || false)}
    for fi18 := 0; fi18 < 19; fi18++ {
        v16 = fi18
        switch (fi18 << ((42) & 63)) {
        case -9223372036854775808:
            v16 = v16
        }
        _ = fi18
    }
    fmt.Println(v16, v16)
    var v19 uint32 = 42
    v20 := GenS2{F0: genFn1(-930, 610, true), F1: 65535}
    delete(v4, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 uint) int32 {
    v6 := GenS1{F0: 4294967295, F1: (1 > -128), F2: false}
    v7 := GenS1{F0: 782, F1: false, F2: (!(false && (false || false)))}
    switch "foo bar" {
    case "\n\t":
        switch -938 {
        case 956:
            if (32767 >= 0) {
                p0 = p0
            }
        case -1:
            v8 := []int{-2147483648}
            _ = v8
        case 128:
            v9 := []int{-355, 4294967295, 42, -9223372036854775808}
            _ = v9
        default:
            p0 -= p0
        }
    case "hello":
        fmt.Println(p0)
    case "":
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 bool, p2 bool) string {
    v10 := GenS1{F0: 538, F1: p2, F2: p1}
    var v11 int64 = 1
    var v12 int64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{1, 4294967295}
    v3 := make(map[int]int)
    var v4 string = ("foo bar" + "")
    var v5 int = 563
    fmt.Println(v4, v4)
    _ = len(v3)
    var v13 int = (-(v5))
    v4 = v4
    switch ("\n\t" + (("a" + "hello") + v4)) {
    case "\n\t":
        var v14 uint16 = 518
        _ = v14
    case "hello":
        _ = len(v3)
    }
    fmt.Println(v4, v13)
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint32 = 65535
        _ = fi16
        _ = v17
    }
    fmt.Println(v5, v5, v15)
    var v18 int8 = 0
    if (!true) {
        fmt.Println(v4, v18, v5)
        v1[2147483647] = (v5 % ((v13) | 1))
        switch v4 {
        case "\n\t":
            _ = len(v2)
        }
    }
    if (2 >= 447) {
        v18 = 42
        v19 := GenS2{F0: 211}
        _ = v19
    }
    fmt.Println(v13, v13, v15)
    v20 := GenS1{F0: (v5 * v5), F1: true, F2: (!(1 == v18))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

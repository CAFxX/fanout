package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 uint }
func genFn0() string {
    var v7 uint64 = 42
    switch "\n\t" {
    case "a":
        v8 := []int{32768, 32768, -1, 859}
        _ = v8
    case "hello":
        fmt.Println(v7)
    default:
        var v9 uint64 = v7
        _ = v9
    }
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64, p1 int16, p2 bool) bool {
    if p2 {
        var v10 int8 = 42
        _ = v10
    } else {
        var v11 uint8 = 255
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(p0, p2)
        p2 = p2
        switch ((("a" + ("foo bar" + "foo bar")) + "") + (((("hello" + "a") + "") + "") + "a")) {
        case "":
            v13 := GenS1{F0: -709, F1: 618, F2: true}
            _ = v13
        case "\n\t":
            p2 = (true && (((p1 - p1) < 276) || (735 > 515)))
        case "hello":
            var v14 string = ("a" + "")
            _ = v14
        }
        _ = fi12
    }
    var v15 int8 = 50
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 0
    v3 := []int{-9223372036854775808, 465, -2147483648}
    var v4 uint = 0
    var v5 int64 = 1
    v6 := make(map[uint32]uint32)
    v16 := GenS1{F0: 32767, F1: 2147483647, F2: (-833 == -2147483648)}
    v2 += v2
    if false {
        switch ("hello" + ("" + "a")) {
        case "a":
            if false {
                v4 += (v4 - (v4 | v4))
            } else {
                v3 = append(v3, 4294967295)
            }
        case "":
            v17 := GenS2{F0: -14, F1: true}
            _ = v17
        case "\n\t":
            var v18 int16 = 1
            _ = v18
        }
        v19 := GenS1{F0: 1000, F1: -2147483648, F2: true}
        v4 = (v4 ^ (v4 - v4))
        _ = v19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 uint = (v4 / ((v4) | 1))
    fmt.Println(v4, v4, v20)
    var v21 bool = true
    if (((("hello" + "hello") + "\n\t") + "\n\t") == ("\n\t" + ("a" + "a"))) {
        v22 := GenS2{F0: 42, F1: true}
        if v21 {
            var v23 bool = ((944 <= 65535) && v21)
            _ = v23
        }
        _ = v22
    }
    v24 := GenS3{F0: 494, F1: 2}
    v21 = true
    var v25 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

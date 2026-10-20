package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32) int64 {
    fmt.Println(p0)
    var v4 int32 = p0
    v4 = (v4 - v4)
    _ = p0
    _ = v4
    return -1
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]int)
    var v3 int64 = 1
    switch ((((("hello" + "hello") + "") + ("\n\t" + "")) + "hello") + "hello") {
    case "hello":
        v5 := []int{127, 9223372036854775807, 326, -9223372036854775808}
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 32767, F1: v3}
        _ = v6
    case "\n\t":
        var v7 float32 = -2.25
        _ = v7
    default:
        v3 += (v3 % ((v3) | 1))
    }
    fmt.Println(v1, v3)
    var v8 int32 = -1
    _ = len(v2)
    if ((-(v8)) == v8) {
        fmt.Println(v1, v3)
        v9 := GenS2{F0: 2}
        _ = v9
    } else {
        var v10 int = 954
        _ = v10
    }
    if true {
        var v11 int32 = v8
        if false {
            _ = len(v2)
            var v12 bool = (v3 >= v3)
            var v13 uint = 593
            _ = v12
            _ = v13
        } else {
            v2[42] = 127
            v3 += (v3 * v3)
            v3 = (v3 & v3)
        }
        _ = v11
    }
    switch ("foo bar" + ("a" + ((("hello" + "") + "a") + ""))) {
    case "":
        fmt.Println(v3, v3, v1)
    case "hello":
        switch 128 {
        case -192:
            var v14 int8 = 0
            _ = v14
        case -2147483648:
            fmt.Println(v3)
        default:
            v15 := []int{956, 1}
            _ = v15
        }
    case "foo bar":
        var v16 bool = false
        _ = v16
    default:
        var v17 uint16 = 2
        _ = v17
    }
    fmt.Println(v3, v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}

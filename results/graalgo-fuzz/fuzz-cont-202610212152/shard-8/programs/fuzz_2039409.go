package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0() bool {
    switch 1 {
    case 128:
        v5 := []int{9223372036854775807, -776}
        _ = v5
    case 9223372036854775807:
        var v6 bool = true
        _ = v6
    case -549:
        var v7 float64 = (1.5 - (0.0 / (0.1 - float64((3.14159 * 0.1)))))
        _ = v7
    default:
        var v8 int8 = -128
        _ = v8
    }
    switch "foo bar" {
    case "":
        var v9 int32 = -1
        _ = v9
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "foo bar":
        v11 := []int{32768, 593}
        _ = v11
    }
    var v12 int64 = -399
    var v13 int16 = 42
    v14 := []int{9223372036854775807, 2147483647, -1}
    _ = v12
    _ = v13
    _ = v14
    return true
}

func main() {
    var v1 uint = 840
    var v2 bool = true
    var v3 uint64 = 0
    v4 := make(map[string]bool)
    v1 = 34
    var v15 int8 = -11
    var v16 int = -262
    switch ("foo bar" + (("a" + "") + "hello")) {
    case "\n\t":
        v1 += 65535
    }
    v1 -= v1
    v17 := GenS1{F0: 42, F1: 885}
    if v2 {
        var v18 string = "foo bar"
        _ = v18
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 string = "hello"
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := []int{(-(fi21)), -543, int(v15), (fi21 % ((127) | 1))}
        fmt.Println(v2, v16)
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: 0, F1: (^(v16))}
    var v24 int64 = 1
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

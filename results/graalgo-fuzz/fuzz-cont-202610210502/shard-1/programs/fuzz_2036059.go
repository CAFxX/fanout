package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint) int32 {
    v7 := GenS1{F0: -32768, F1: (p1 % (((^(p1))) | 1))}
    var v8 float32 = ((1.5 + -2.25) - (((3.14159 - 0.0) / (0.0 + float32(3.14159))) * -2.25))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16) int {
    switch "hello" {
    case "\n\t":
        p0 += p0
    case "":
        v9 := []int{128}
        _ = v9
    case "foo bar":
        p0 = p0
    default:
        p0 = (p0 + (p0 + (^(p0))))
    }
    p0 = (p0 << ((2) & 63))
    v10 := []int{-543, 860, -9223372036854775808, 956}
    fmt.Println(p0)
    v11 := GenS1{F0: 0, F1: 27}
    _ = p0
    _ = v10
    _ = v11
    return -171
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 int16 = -456
    v4 := make(map[string]bool)
    var v5 bool = false
    v6 := []string{("foo bar" + "hello"), (("a" + "\n\t") + "hello")}
    fmt.Println(v5)
    switch "\n\t" {
    case "\n\t":
        var v12 string = "foo bar"
        _ = v12
    case "foo bar":
        switch ("" + "\n\t") {
        case "hello":
            for fi13 := 0; fi13 < 1; fi13++ {
                v3 = -32768
                _ = fi13
            }
        case "\n\t":
            var v14 string = "foo bar"
            _ = v14
        default:
            var v15 int16 = (v3 >> ((538) & 63))
            _ = v15
        }
    default:
        v6 = append(v6, (("a" + "") + "foo bar"))
    }
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int8 = 100
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: (v3 - v3), F1: 865}
    var v19 int64 = 708
    fmt.Println(v5)
    fmt.Println(v19)
    v20 := GenS2{F0: 255}
    fmt.Println(v5, v3)
    v5 = false
    switch 79 {
    case 9223372036854775807:
        v5 = (!v5)
    default:
        v21 := v4["a"]
        _ = v21
    }
    var v22 string = ("hello" + "hello")
    v19 += (v19 ^ v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

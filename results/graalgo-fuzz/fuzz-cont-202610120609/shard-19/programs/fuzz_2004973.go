package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16) bool {
    var v7 int16 = 42
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 15; fi9++ {
        p0 += p0
        switch "hello" {
        case "hello":
            v7 += 32767
        default:
            p0 = (p0 - p0)
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint32 = 979
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 int = 529
    v5 := make(map[int]string)
    var v6 string = ((("a" + ("\n\t" + ((("\n\t" + "\n\t") + "a") + ("foo bar" + (("hello" + "hello") + "a"))))) + "a") + ("" + ""))
    switch v6 {
    case "hello":
        var v12 uint8 = 2
        _ = v12
    default:
        v6 = v6
    }
    fmt.Println(v6)
    var v13 string = v6
    v13 = v6
    v6 = "a"
    v13 = (("\n\t" + v6) + v13)
    var v14 string = (v6 + "foo bar")
    v15 := GenS1{F0: -2147483648, F1: 1}
    v16 := GenS1{F0: -2147483648, F1: 0}
    if (100 < -128) {
        v17 := []int{(-(v4)), (v4 / ((724) | 1)), v4}
        _ = v17
    } else {
        v6 = ("\n\t" + v14)
        v18 := GenS2{F0: v6}
        v14 = v6
        _ = v18
    }
    fmt.Println(v6, v4)
    v4 = (v4 >> ((0) & 63))
    v19 := []int{v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

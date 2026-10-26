package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 int32 = 1
    if true {
        var v6 float64 = 0.1
        v7 := GenS1{F0: -128, F1: 42}
        switch "\n\t" {
        case "\n\t":
            v8 := GenS1{F0: 42, F1: 872}
            _ = v8
        case "":
            var v9 int32 = v5
            _ = v9
        case "a":
            var v10 bool = (!false)
            _ = v10
        }
        _ = v6
        _ = v7
    } else {
        v11 := GenS2{F0: 4294967295}
        switch "foo bar" {
        case "a":
            delete(v4, "")
        default:
            for fi12 := 0; fi12 < 7; fi12++ {
                v5 = (v5 ^ v5)
                _ = fi12
            }
        }
        _ = v11
    }
    if false {
        fmt.Println(v5)
    } else {
        v13 := []int{}
        var v14 uint16 = 269
        var v15 uint8 = 1
        _ = v13
        _ = v14
        _ = v15
    }
    v5 = -325
    v16 := []int{32768, 1, -573}
    fmt.Println(v5)
    var v17 string = ((("\n\t" + (((("foo bar" + "hello") + "hello") + "") + "")) + "foo bar") + "a")
    var v18 float64 = -2.25
    v17 = (("\n\t" + ((v17 + v17) + "hello")) + v17)
    v18 -= v18
    var v19 float32 = (((-2.25 - 3.14159) - ((-2.25 - float32(v18)) - -2.25)) * 1e10)
    var v20 bool = ((v19 != v19) && true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

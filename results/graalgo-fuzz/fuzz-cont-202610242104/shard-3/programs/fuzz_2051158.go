package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 bool, p1 uint64, p2 bool) int8 {
    p2 = p0
    var v4 bool = false
    for fi5 := 0; fi5 < 9; fi5++ {
        switch "hello" {
        case "hello":
            v6 := []int{}
            _ = v6
        default:
            v7 := GenS1{F0: "hello", F1: 128}
            _ = v7
        }
        var v8 int64 = 9223372036854775807
        _ = fi5
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    for fi9 := 0; fi9 < 16; fi9++ {
        switch ("foo bar" + "a") {
        case "":
            v10 := []int{}
            _ = v10
        case "a":
            var v11 uint32 = v1
            _ = v11
        default:
            v3 = v3
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint16 = 255
            _ = fi12
            _ = v13
        }
        _ = fi9
    }
    v3 = (v3 * (v3 & v3))
    var v14 uint32 = 841
    var v15 float64 = 1.5
    fmt.Println(v1, v1, v1)
    fmt.Println(v14)
    for fi16 := 0; fi16 < 11; fi16++ {
        v15 = 1.5
        _ = fi16
    }
    fmt.Println(v14, v15)
    v17 := GenS1{F0: (("\n\t" + "") + ("hello" + (((("hello" + "hello") + "") + "") + "foo bar"))), F1: -9223372036854775808}
    fmt.Println(v3, v3)
    fmt.Println(v1, v1)
    v15 = ((v15 - float64(float32((v15 / v15)))) * (3.14159 * v15))
    v18 := []int{127, 42}
    var v19 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}

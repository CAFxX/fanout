package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 int16) uint16 {
    p0 = p0
    v6 := GenS1{F0: 682, F1: 114}
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint32, p1 int64) int {
    if true {
        p0 = (p0 + p0)
        var v7 uint8 = 127
        _ = v7
    }
    fmt.Println(p0, p0)
    p1 = (p1 - 1)
    p1 = p1
    _ = p0
    _ = p1
    return 338
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 float32 = 0.0
    v4 := make(map[uint32]uint32)
    var v5 int32 = 1
    if true {
        switch ("foo bar" + ("" + "\n\t")) {
        case "foo bar":
            var v8 int16 = int16(v5)
            _ = v8
        }
    } else {
        delete(v1, 4294967295)
        fmt.Println(v5, v3)
        var v9 uint16 = 32767
        _ = v9
    }
    v10 := []int{}
    v3 = v3
    v5 += 1
    switch "foo bar" {
    case "foo bar":
        v11 := GenS2{F0: (42 <= 4294967295), F1: false, F2: 2147483647}
        _ = v11
    }
    var v12 int = -945
    if len(v10) > 0 {
        v10[0] = (v12 ^ v12)
    }
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

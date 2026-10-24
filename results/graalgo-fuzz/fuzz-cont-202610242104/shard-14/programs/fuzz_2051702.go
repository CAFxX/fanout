package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0(p0 int32, p1 uint64, p2 uint64) int16 {
    if true {
        fmt.Println(p1, p2)
        v4 := GenS1{F0: p0, F1: 230, F2: 1}
        p1 -= 1
        _ = v4
    }
    v5 := []int{32767, -736}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    v1 := []int{42}
    v2 := make(map[uint32]bool)
    var v3 int8 = 91
    fmt.Println(v3)
    switch (((("a" + (("hello" + "foo bar") + "")) + "a") + "foo bar") + "foo bar") {
    case "\n\t":
        v6 := GenS1{F0: 825, F1: 1, F2: 42}
        _ = v6
    }
    var v7 int = 42
    var v8 float32 = -2.25
    v1 = append(v1, v7)
    switch "\n\t" {
    case "hello":
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint8 = 2
            var v11 uint16 = uint16(v10)
            var v12 int64 = -193
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v7 = v7
    v13 := GenS1{F0: 1, F1: 42, F2: -393}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int = (v7 * v7)
        v16 := GenS1{F0: 994, F1: 1, F2: -9223372036854775808}
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v17 := GenS1{F0: -1, F1: 164, F2: -633}
            _ = v17
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v18 := v2[2]
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}

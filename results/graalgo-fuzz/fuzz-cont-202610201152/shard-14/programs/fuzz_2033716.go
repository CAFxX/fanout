package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8) uint8 {
    v5 := GenS1{F0: ""}
    v6 := []int{42, 1, -2147483648}
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 uint, p1 string) string {
    fmt.Println(p0)
    switch (p1 + "\n\t") {
    case "foo bar":
        v7 := GenS2{F0: false, F1: 2147483647}
        _ = v7
    case "hello":
        p1 = "\n\t"
    }
    _ = p0
    _ = p1
    return ((p1 + p1) + p1)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483648, 2}
    var v3 bool = (false && false)
    v4 := make(map[uint32]uint32)
    for fi8 := 0; fi8 < 16; fi8++ {
        v3 = (!v3)
        _ = fi8
    }
    v3 = v3
    delete(v4, 589)
    var v9 int64 = -37
    var v10 string = "a"
    switch (((("foo bar" + "\n\t") + v10) + (v10 + ((("a" + "foo bar") + "foo bar") + v10))) + "") {
    case "hello":
        if false {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint64 = 382
                _ = fi11
                _ = v12
            }
        }
    default:
        var v13 int64 = 1
        _ = v13
    }
    var v14 uint16 = 255
    var v15 uint32 = 4294967295
    if (v9 < 1) {
        fmt.Println(v3)
        var v16 int8 = 0
        var v17 uint8 = genFn0(5)
        _ = v16
        _ = v17
    } else {
        var v18 uint8 = 1
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 uint) uint32 {
    p0 = p0
    var v4 int64 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = 283
        v7 := []int{-710, (fi5 >> ((p1) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (("a" + "hello") + "\n\t")}
    if (-128 >= 42) {
        var v9 uint16 = 32767
        if true {
            p1 -= p1
            fmt.Println(v9, v9)
        }
        v9 = v9
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 633
}
func genFn1(p0 int, p1 int8) int8 {
    fmt.Println(p1, p0)
    switch "foo bar" {
    case "foo bar":
        var v10 uint = 0
        _ = v10
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    v3 := make(map[string]int)
    _ = len(v1)
    v2 -= v2
    v2 -= (-(v2))
    v11 := GenS1{F0: ("a" + ("hello" + "\n\t"))}
    fmt.Println(v2)
    var v12 int64 = -966
    var v13 float64 = (1e10 / ((1.5 - (1.5 * 1.5)) + (-2.25 - float64(3.14159))))
    _ = len(v1)
    switch "a" {
    case "":
        var v14 uint16 = 65535
        _ = v14
    case "a":
        if (!(0 == 0)) {
            var v15 int32 = 0
            _ = v15
        }
    case "\n\t":
        fmt.Println(v2)
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: "\n\t"}
        _ = fi16
        _ = v17
    }
    var v18 int64 = (v12 * (v12 / ((v12) | 1)))
    v18 -= v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
}

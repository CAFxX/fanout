package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint, p1 int8, p2 int16) int64 {
    fmt.Println(p0, p1)
    if (!((p0 * 42) > (p0 & 18446744073709551615))) {
        fmt.Println(p0, p0, p1)
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16, p1 uint32) bool {
    switch "" {
    case "a":
        v7 := []int{}
        _ = v7
    case "":
        v8 := []int{}
        _ = v8
    }
    fmt.Println(p1, p1)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return (false && false)
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1000
    var v4 float64 = 0.1
    var v5 int32 = -1
    var v6 uint8 = 255
    fmt.Println(v6, v3)
    v9 := GenS1{F0: ("hello" + "\n\t")}
    fmt.Println(v6, v4)
    var v10 uint64 = 4294967295
    if false {
        _ = len(v1)
        v11 := GenS1{F0: ""}
        _ = v11
    }
    v5 -= -899
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v3, v5)
        _ = fi12
    }
    switch "foo bar" {
    case "foo bar":
        switch 0 {
        case 32767:
            var v13 int32 = (^(v5))
            _ = v13
        }
    case "":
        fmt.Println(v4, v10)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

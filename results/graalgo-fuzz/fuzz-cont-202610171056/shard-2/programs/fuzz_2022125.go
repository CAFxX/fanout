package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int, p1 int, p2 int64) uint64 {
    var v6 uint16 = 42
    v7 := []int{p1, p1, (p0 + (p0 ^ p0)), 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []int{-1}
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 int64 = -845
    v5 = v5
    if true {
        v5 += (v5 | v5)
    }
    var v8 float64 = 3.14159
    for fi9 := 0; fi9 < 3; fi9++ {
        delete(v3, 119)
        for fi10 := 0; fi10 < 16; fi10++ {
            v4[0] = (!true)
            v1 = append(v1, fi10)
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS1{F0: 0}
    var v12 int = 0
    if len(v1) > 0 {
        v1[0] = -217
    }
    switch "foo bar" {
    case "a":
        switch "a" {
        case "":
            v13 := GenS1{F0: 100}
            _ = v13
        case "\n\t":
            switch ("" + "hello") {
            case "foo bar":
                v4[42] = false
            }
        }
    }
    if false {
        var v14 int32 = 42
        switch ("\n\t" + ("a" + ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + (("a" + "") + ("foo bar" + "hello"))))) {
        case "":
            delete(v2, v12)
        case "foo bar":
            v15 := GenS1{F0: 1}
            _ = v15
        case "a":
            var v16 float64 = 0.0
            _ = v16
        default:
            switch "\n\t" {
            case "hello":
                switch "" {
                case "a":
                    v8 -= v8
                case "":
                    var v17 int16 = -32768
                    _ = v17
                default:
                    fmt.Println(v14, v12, v12)
                }
            }
        }
        switch (v12 | v12) {
        case -1:
            v14 -= -32
        case -169:
            fmt.Println(v8, v5)
        }
        _ = v14
    } else {
        var v18 int64 = v5
        v5 = v18
        _ = v18
    }
    var v19 uint64 = 263
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

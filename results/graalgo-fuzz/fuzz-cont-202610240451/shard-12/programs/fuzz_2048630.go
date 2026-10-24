package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int {
    v4 := GenS1{F0: 855}
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(fi5 + (fi5 | (fi5 * fi5))), fi5}
        var v7 int32 = -1
        if false {
            switch ("a" + ("hello" + "a")) {
            case "hello":
                for fi8 := 0; fi8 < 15; fi8++ {
                    fmt.Println(v7, v7)
                    _ = fi8
                }
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if (!true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            switch ("foo bar" + "\n\t") {
            case "a":
                var v10 float32 = 0.1
                _ = v10
            case "":
                var v11 uint = 18446744073709551615
                _ = v11
            }
            var v12 uint16 = 0
            _ = fi9
            _ = v12
        }
        var v13 float64 = 1.5
        _ = v13
    }
    var v14 int = 1
    var v15 float32 = 1e10
    _ = v14
    _ = v15
    _ = v4
    return v14
}
func genFn1(p0 int16) uint32 {
    var v16 int32 = 275
    v16 = v16
    _ = p0
    _ = v16
    return 388
}

func main() {
    v1 := []int{2147483647, 241, -189, -36}
    var v2 uint = 18446744073709551615
    var v3 uint32 = 1
    v17 := GenS1{F0: 0}
    fmt.Println(v2, v3)
    switch "" {
    case "\n\t":
        v3 = genFn1(-483)
    case "foo bar":
        fmt.Println(v2, v3)
    default:
        var v18 uint = 42
        _ = v18
    }
    v3 = (v3 * v3)
    var v19 uint8 = 42
    var v20 int64 = -1
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v20)
        _ = fi21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint32 = v3
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 4294967295}
    switch "" {
    case "foo bar":
        v20 = (v20 * (v20 * v20))
    case "a":
        switch "hello" {
        case "hello":
            fmt.Println(v20)
        default:
            v19 += v19
        }
    case "hello":
        var v25 uint = (v2 & v2)
        _ = v25
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 uint16) int16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 int64 = 0
        var v6 int8 = -73
        _ = fi4
        _ = v5
        _ = v6
    }
    switch "hello" {
    case "\n\t":
        switch -9223372036854775808 {
        case 32768:
            v7 := GenS1{F0: 127, F1: 1, F2: 388}
            _ = v7
        case -2147483648:
            p0 = (p0 / ((1) | 1))
        case -285:
            fmt.Println(p0)
        }
    default:
        var v8 int16 = 42
        _ = v8
    }
    var v9 uint16 = 255
    _ = p0
    _ = v9
    return -160
}

func main() {
    v1 := []string{"hello"}
    var v2 int32 = -1
    v3 := []uint32{2147483648}
    fmt.Println(v2)
    fmt.Println(v2)
    var v10 int64 = -9223372036854775808
    v11 := GenS1{F0: 32768, F1: 1, F2: 42}
    v1 = append(v1, "\n\t")
    for fi12 := 0; fi12 < 15; fi12++ {
        v3 = append(v3, 763)
        v13 := GenS1{F0: (fi12 - fi12), F1: -103, F2: 4294967295}
        var v14 int8 = 100
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v3)
    if (!false) {
        var v15 int64 = (-(v10))
        _ = v15
    } else {
        var v16 int16 = genFn0(0)
        v16 = v16
        var v17 string = "a"
        _ = v16
        _ = v17
    }
    v10 = (^(v10))
    v2 += (v2 * (-(v2)))
    fmt.Println(v2, v2)
    v10 -= (-(v10))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 9223372036854775807 {
    case 914:
        if (((("hello" + ("\n\t" + ("hello" + ""))) + "hello") + "hello") != "a") {
            for fi18 := 0; fi18 < 11; fi18++ {
                switch fi18 {
                case 1:
                    var v19 string = "hello"
                    _ = v19
                case 4294967295:
                    fmt.Println(v2, v10)
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                _ = fi18
            }
        }
    }
    if (!(!true)) {
        var v20 uint8 = 0
        var v21 uint64 = 18446744073709551615
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint32 = 32
        if true {
            v24 := GenS2{F0: fi22, F1: 42, F2: 100}
            _ = len(v1)
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    v25 := GenS3{F0: ((!false) || ((!false) && (2147483647 < 1))), F1: 2, F2: 476}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
}

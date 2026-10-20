package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 uint = 255
    v4 := []int{32768, -1, 32767, 127}
    v5 := []int{1, 127, 2147483647, 127}
    v1["\n\t"] = false
    switch "\n\t" {
    case "hello":
        if (!true) {
            v6 := GenS1{F0: -32768, F1: 0, F2: 1}
            v7 := GenS2{F0: 1, F1: ""}
            _ = v6
            _ = v7
        }
    case "a":
        v8 := v1["foo bar"]
        _ = v8
    default:
        v9 := GenS1{F0: 539, F1: 0, F2: 42}
        _ = v9
    }
    var v10 uint32 = 4294967295
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint = 656
        v4 = append(v4, fi11)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: -1, F1: uint64(v3), F2: 0}
    _ = len(v1)
    v3 = v3
    v10 = (v10 % (((v10 << ((v3) & 63))) | 1))
    _ = len(v1)
    v1[("hello" + "foo bar")] = true
    v2 -= v2
    v14 := GenS1{F0: 0, F1: 4294967295, F2: 703}
    fmt.Println(v3, v2)
    switch "foo bar" {
    case "a":
        v10 = uint32(v3)
    case "hello":
        if (42 != 0) {
            fmt.Println(v3, v3)
            fmt.Println(v10, v3)
            v10 += (v10 % ((v10) | 1))
        } else {
            var v15 bool = (100 < 1)
            var v16 bool = (1 != 42)
            var v17 int = -463
            _ = v15
            _ = v16
            _ = v17
        }
    case "\n\t":
        v18 := GenS3{F0: "foo bar"}
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint = (v3 * v3)
            if true {
                v10 = v10
            } else {
                switch 127 {
                case -1:
                    fmt.Println(fi20)
                default:
                    v22 := GenS2{F0: 125, F1: (("foo bar" + "") + "\n\t")}
                    _ = v22
                }
            }
            _ = fi20
            _ = v21
        }
        v23 := v1["hello"]
        _ = len(v4)
        _ = fi19
        _ = v23
    }
    if false {
        v10 -= (^(v10))
    } else {
        fmt.Println(v10, v2, v3)
        v24 := GenS4{F0: v3, F1: 1}
        fmt.Println(v3)
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

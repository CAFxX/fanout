package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    v7 := GenS1{F0: 1, F1: -1, F2: 421}
    switch "foo bar" {
    case "hello":
        v8 := []int{}
        _ = v8
    case "foo bar":
        p0 += p0
    case "\n\t":
        var v9 int64 = -1
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1() uint {
    var v10 uint64 = 0
    var v11 int16 = 662
    v12 := GenS1{F0: (v10 / ((v10) | 1)), F1: 4294967295, F2: 42}
    v13 := GenS1{F0: (v10 % ((genFn0(-1)) | 1)), F1: 127, F2: 0}
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: (v10 / (((v10 % ((132) | 1))) | 1)), F1: (fi14 + 9223372036854775807), F2: 2147483648}
        _ = fi14
        _ = v15
    }
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 65535
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 42
    var v3 uint = 18446744073709551615
    var v4 int64 = 0
    v5 := make(map[int]int)
    v6 := []int{32768, 32767, 9223372036854775807}
    v1 += float32((1e10 * 1.5))
    if (true || false) {
        switch ("a" + "foo bar") {
        case "hello":
            var v16 bool = false
            _ = v16
        }
    }
    v6 = append(v6, 127)
    v17 := GenS1{F0: 454, F1: 1, F2: 974}
    fmt.Println(v1, v3, v1)
    v18 := GenS2{F0: v2}
    var v19 bool = (false && (true || (false && true)))
    v20 := GenS1{F0: 394, F1: 4294967295, F2: 1}
    v4 = v4
    v21 := GenS1{F0: 42, F1: 0, F2: 2}
    for fi22 := 0; fi22 < 20; fi22++ {
        switch (("a" + "a") + "") {
        case "a":
            _ = len(v5)
        case "\n\t":
            var v23 string = ("hello" + "a")
            _ = v23
        case "foo bar":
            v3 = v3
        default:
            if ((false || v19) && true) {
                if (false || v19) {
                    fmt.Println(fi22)
                }
                fmt.Println(v4, v3)
                if (genFn1() == uint(v2)) {
                    v24 := GenS3{F0: 690}
                    var v25 int8 = 42
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v1 = ((0.1 - (v1 - 0.1)) / (((v1 * v1) * float32((3.14159 * (float64(((v1 - ((0.0 - v1) + 1e10)) - 0.1)) * (1.5 - 0.1))))) - (float32(3.14159) - 1e10)))
                        _ = fi26
                    }
                    _ = v24
                    _ = v25
                } else {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v28 := GenS2{F0: (v2 - 42)}
                        _ = fi27
                        _ = v28
                    }
                    _ = len(v5)
                    var v29 bool = v19
                    _ = v29
                }
            }
        }
        fmt.Println(v19, v3)
        _ = fi22
    }
    if v19 {
        v30 := GenS1{F0: genFn0(1), F1: 684, F2: 2}
        fmt.Println(v19)
        v4 += (v4 % ((v4) | 1))
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 int16) uint {
    var v7 bool = false
    var v8 int8 = 63
    v8 = v8
    if ((42 == 58) || v7) {
        v9 := []int{42, 2147483647, 0, int(v8)}
        var v10 uint = 0
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                fmt.Println(fi11, v8, v8)
                _ = fi12
            }
            var v13 float32 = (1e10 - 3.14159)
            v14 := GenS1{F0: fi11, F1: -2147483648, F2: 629}
            _ = fi11
            _ = v13
            _ = v14
        }
        var v15 int = -71
        _ = v15
    }
    var v16 uint16 = 32767
    _ = p0
    _ = v16
    _ = v7
    _ = v8
    return 512
}
func genFn1(p0 int64, p1 uint32, p2 uint16) uint {
    var v17 float64 = -2.25
    v18 := GenS2{F0: genFn0(0), F1: 18446744073709551615}
    fmt.Println(p2, p1)
    v19 := GenS2{F0: 739, F1: 0}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return 822
}

func main() {
    var v1 bool = (false || false)
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    var v5 int32 = 42
    v6 := []int{}
    fmt.Println(v1, v1)
    v6 = append(v6, 1)
    var v20 uint16 = 0
    var v21 float64 = 0.1
    var v22 uint16 = v20
    var v23 uint64 = 151
    switch "\n\t" {
    case "hello":
        switch 127 {
        case -120:
            fmt.Println(v5, v1)
        case 128:
            v5 = v5
        case 1:
            for fi24 := 0; fi24 < 5; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    switch "foo bar" {
                    case "":
                        if len(v6) > 0 {
                            v6[0] = (fi25 / ((4294967295) | 1))
                        }
                    case "\n\t":
                        var v26 int8 = -38
                        _ = v26
                    }
                    v27 := GenS3{F0: 1, F1: 255, F2: -40}
                    _ = fi25
                    _ = v27
                }
                var v28 string = "foo bar"
                _ = fi24
                _ = v28
            }
        }
    case "foo bar":
        var v29 float32 = -2.25
        _ = v29
    default:
        v30 := v4[2147483647]
        _ = v30
    }
    var v31 float32 = (1.5 - 1.5)
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 int16 = 42
        _ = fi32
        _ = v33
    }
    v34 := v3["hello"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32, p1 uint32, p2 uint32) bool {
    if true {
        var v4 float64 = (((float64(0.0) / 0.1) + float64(1e10)) + 1e10)
        if true {
            v5 := GenS1{F0: "hello", F1: "foo bar", F2: 68}
            _ = v5
        } else {
            v6 := GenS1{F0: "hello", F1: "a", F2: 1}
            v7 := GenS2{F0: 113, F1: -1}
            p0 = (p0 << ((984) & 63))
            _ = v6
            _ = v7
        }
        v8 := GenS3{F0: 929}
        _ = v4
        _ = v8
    } else {
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                p2 -= p1
                v10 := GenS3{F0: 226}
                _ = fi9
                _ = v10
            }
        } else {
            v11 := []int{-794}
            fmt.Println(p0, p1, p0)
            _ = v11
        }
        if true {
            p2 = p1
            var v12 string = "a"
            _ = v12
        }
    }
    v13 := GenS1{F0: "hello", F1: "hello", F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -249
    v3 := []int{32767, (v2 * (v2 & v2))}
    fmt.Println(v2)
    v2 -= (^(v2))
    v1[(-(v2))] = 1
    if true {
        switch 9223372036854775807 {
        case 504:
            v14 := GenS3{F0: 18446744073709551615}
            _ = v14
        }
        v15 := GenS4{F0: 42, F1: 772}
        fmt.Println(v2)
        _ = v15
    } else {
        v16 := v1[v2]
        v17 := v1[(v2 % ((v2) | 1))]
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    switch -1 {
    case 303:
        v19 := GenS4{F0: 9223372036854775807, F1: 2}
        _ = v19
    case 1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v2 = (-(v2))
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}

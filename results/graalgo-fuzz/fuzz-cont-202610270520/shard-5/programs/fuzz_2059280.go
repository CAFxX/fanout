package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 uint64) bool {
    for fi4 := 0; fi4 < 15; fi4++ {
        switch 42 {
        case 4294967295:
            v5 := GenS1{F0: p0, F1: 1}
            _ = v5
        case 32767:
            p1 = 0
        default:
            v6 := []int{127, (fi4 >> ((987) & 63)), (fi4 + fi4)}
            _ = v6
        }
        p0 = (p0 | p0)
        v7 := GenS2{F0: fi4, F1: -1}
        _ = fi4
        _ = v7
    }
    v8 := GenS2{F0: 2147483647, F1: (^(p0))}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v8
    return false
}
func genFn1(p0 string) uint8 {
    v9 := []int{-302, 127, 602, 32768}
    p0 = "hello"
    p0 = "foo bar"
    _ = p0
    _ = v9
    return 255
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 int = 42
    if v1 {
        v10 := []int{v3, v3, v3}
        _ = v10
    }
    if (-412 > -405) {
        var v11 bool = ((9223372036854775807 == 808) || ((!(4294967295 > 4294967295)) || v1))
        var v12 string = ""
        _ = v11
        _ = v12
    }
    if v1 {
        if (true || false) {
            var v13 int = 128
            _ = v13
        }
    }
    v1 = true
    switch (v3 << ((18446744073709551615) & 63)) {
    case -1:
        v14 := GenS2{F0: v3, F1: 2147483647}
        _ = v14
    default:
        fmt.Println(v1, v3)
    }
    v15 := []int{(^(v3)), (v3 - v3), v3, v3}
    if v1 {
        v3 = -9223372036854775808
    }
    v16 := GenS1{F0: 1, F1: 0}
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 = true
        _ = fi17
    }
    _ = len(v2)
    v3 -= v3
    if v1 {
        var v18 uint64 = 2
        _ = v18
    } else {
        v19 := GenS3{F0: -1}
        _ = v19
    }
    var v20 bool = (v1 || ("foo bar" != "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
}

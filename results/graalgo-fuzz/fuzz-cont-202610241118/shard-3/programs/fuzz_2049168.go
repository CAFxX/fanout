package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
func genFn0(p0 int8) uint {
    p0 = (^(p0))
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = 1
        _ = fi6
    }
    _ = p0
    return 2
}
func genFn1(p0 int32, p1 bool) uint {
    p1 = (p1 && false)
    fmt.Println(p0, p0)
    if (p1 || ((!(p1 && (356 != 666))) && (!("hello" > (("" + "") + (("hello" + "a") + "\n\t")))))) {
        v7 := []int{57, 127}
        _ = v7
    } else {
        v8 := GenS1{F0: (1e10 != ((float32(0.0) - 0.0) / 3.14159)), F1: "\n\t", F2: -128}
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                if (!p1) {
                    p0 += -1
                }
                _ = fi10
            }
            p1 = (p1 && false)
            fmt.Println(p1, fi9)
            _ = fi9
        }
        fmt.Println(p1)
        _ = v8
    }
    v11 := GenS1{F0: (!true), F1: "foo bar", F2: 0}
    for fi12 := 0; fi12 < 18; fi12++ {
        if (0 >= 255) {
            v13 := GenS2{F0: 4294967295, F1: 9223372036854775807}
            p1 = (9223372036854775807 == 0)
            fmt.Println(p0, p0)
            _ = v13
        } else {
            switch "\n\t" {
            case "a":
                var v14 uint = 0
                _ = v14
            default:
                p1 = p1
            }
            v15 := GenS1{F0: (-2.25 != 1.5), F1: "", F2: 100}
            v16 := GenS3{F0: p1}
            _ = v15
            _ = v16
        }
        var v17 float64 = 0.1
        _ = fi12
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v11
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 int = -622
    var v4 int = v3
    var v5 float32 = -2.25
    for fi18 := 0; fi18 < 2; fi18++ {
        switch (v4 * (fi18 % (((v4 >> ((1) & 63))) | 1))) {
        case -2147483648:
            var v19 int32 = 1
            _ = v19
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            v3 = (-(fi20))
            var v21 int32 = -1
            _ = fi20
            _ = v21
        }
        v3 += v3
        _ = fi18
    }
    fmt.Println(v3)
    fmt.Println(v3)
    var v22 int8 = 1
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(fi23, v3)
        fmt.Println(v3, fi23, v22)
        _ = fi23
    }
    v24 := GenS4{F0: 32767}
    v25 := v2[4294967295]
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int64 = -788
        var v28 float32 = -2.25
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := []int{v25, -108}
    var v30 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

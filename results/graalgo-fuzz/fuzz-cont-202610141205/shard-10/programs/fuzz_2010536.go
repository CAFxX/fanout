package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
func genFn0(p0 string, p1 int) int64 {
    p1 += 32768
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch 9223372036854775807 {
            case -794:
                p0 = (("a" + (p0 + ("hello" + p0))) + p0)
            case 127:
                switch (-(fi6)) {
                case -1:
                    p1 = (fi6 ^ 1)
                case 32768:
                    var v8 uint64 = 2
                    _ = v8
                }
            default:
                var v9 uint8 = 2
                _ = v9
            }
            p0 = p0
            _ = fi7
        }
        fmt.Println(p1, p1, p0)
        _ = fi6
    }
    v10 := []int{-501}
    for fi11 := 0; fi11 < 1; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        p1 += fi11
        if false {
            var v12 int8 = 42
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint8 = 157
                _ = fi13
                _ = v14
            }
            var v15 uint = 1
            _ = v15
        }
        _ = fi11
    }
    var v16 uint64 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v16
    return 722
}
func genFn1(p0 string) int8 {
    var v17 uint32 = 286
    p0 = (("hello" + "hello") + "")
    _ = p0
    _ = v17
    return -1
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 int64 = -1
    v5 := make(map[int]uint32)
    var v18 int32 = 2147483647
    if false {
        var v19 float32 = float32(1.5)
        fmt.Println(v19, v1, v4)
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 string = v1
        _ = fi20
        _ = v21
    }
    var v22 float64 = 3.14159
    var v23 int32 = (v18 & (-(v18)))
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint64 = 391
        switch fi24 {
        case 1:
            var v26 int64 = (v4 - 0)
            _ = v26
        case 887:
            var v27 bool = true
            _ = v27
        case 4294967295:
            v28 := GenS1{F0: v4, F1: 138}
            _ = v28
        default:
            var v29 int32 = 1
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    if (!true) {
        delete(v3, -1)
    } else {
        var v30 int16 = int16(v18)
        fmt.Println(v4, v1, v1)
        v31 := v2["a"]
        _ = v30
        _ = v31
    }
    v32 := GenS1{F0: 694, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}

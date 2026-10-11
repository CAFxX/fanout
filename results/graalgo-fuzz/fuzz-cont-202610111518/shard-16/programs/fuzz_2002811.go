package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0(p0 uint8, p1 string, p2 uint64) uint32 {
    p2 = (p2 | p2)
    var v6 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int) uint64 {
    p0 = p0
    p0 -= p0
    p0 -= p0
    _ = p0
    return 4294967295
}

func main() {
    var v1 int32 = 882
    v2 := []string{"a"}
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 float64 = (((float64(1.5) / ((1e10 - 1e10) - (0.1 * 1e10))) + 3.14159) / 1e10)
    v1 = v1
    for fi7 := 0; fi7 < 5; fi7++ {
        v1 += 318
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    v5 = v5
    if false {
        v8 := GenS1{F0: 42, F1: 0, F2: genFn1(127)}
        var v9 int64 = 42
        _ = v8
        _ = v9
    } else {
        if false {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint32 = 1
                var v12 int16 = -1
                v13 := GenS2{F0: 2, F1: 927}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
    }
    switch -392 {
    case -921:
        v5 = v5
    case 128:
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v1, v5)
            _ = fi14
        }
    default:
        if true {
            var v15 int8 = 50
            var v16 uint = 162
            _ = v15
            _ = v16
        }
    }
    if (65535 != 0) {
        if true {
            var v17 float64 = v5
            v18 := GenS1{F0: 4294967295, F1: 1000, F2: 1}
            _ = v17
            _ = v18
        }
        var v19 uint16 = 945
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = 98
            _ = fi20
            _ = v21
        }
        v2 = append(v2, "\n\t")
        for fi22 := 0; fi22 < 8; fi22++ {
            v5 -= v5
            var v23 float32 = (0.1 / 1.0)
            v24 := GenS3{F0: 373, F1: "hello"}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v5 -= -2.25
    v25 := GenS4{F0: 2147483647, F1: "\n\t", F2: 1}
    v3[-22] = (("hello" + "\n\t") + ("\n\t" + "hello"))
    switch 769 {
    case 85:
        v1 = v1
    case 127:
        var v26 uint16 = 221
        _ = v26
    default:
        v4["hello"] = 32768
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

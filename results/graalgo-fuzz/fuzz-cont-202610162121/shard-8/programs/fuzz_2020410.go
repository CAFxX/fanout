package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 int8, p1 uint) uint8 {
    if false {
        var v4 float32 = -2.25
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: (p0 << ((46) & 63)), F1: 1, F2: -253}
            var v7 float32 = v4
            var v8 float64 = 3.14159
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := []int{}
        _ = v4
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            p0 = (p0 + 100)
            _ = fi10
        }
        v11 := []int{14, 128, 42, 127}
        if (!((2147483647 == uint32(p1)) || (1 > uint32(p1)))) {
            var v12 uint16 = 659
            _ = v12
        }
        _ = v11
    }
    v13 := GenS1{F0: 0, F1: 42, F2: 0}
    p1 = (p1 * 18446744073709551615)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v13
    return 41
}
func genFn1(p0 bool, p1 uint16, p2 int) int64 {
    v14 := GenS2{F0: 0, F1: p2, F2: 1}
    if (((true || (42 == 42)) && (!(p0 && (!(p0 && (!p0)))))) && (false || p0)) {
        v15 := []int{}
        _ = v15
    } else {
        var v16 uint = 918
        _ = v16
    }
    var v17 uint32 = 2147483647
    var v18 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 int8 = 1
    var v2 bool = (false && false)
    v3 := make(map[int]int)
    for fi19 := 0; fi19 < 3; fi19++ {
        delete(v3, fi19)
        switch (fi19 * (fi19 + 1)) {
        case 1:
            v1 += (v1 / ((v1) | 1))
        case -1:
            var v20 float32 = 1e10
            _ = v20
        default:
            for fi21 := 0; fi21 < 14; fi21++ {
                v2 = (true && v2)
                var v22 int64 = 0
                var v23 bool = true
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = fi19
    }
    var v24 int = 32768
    var v25 uint64 = 18446744073709551615
    v26 := GenS2{F0: 1, F1: (v24 / ((v24) | 1)), F2: -59}
    var v27 uint = uint(v25)
    delete(v3, -9223372036854775808)
    _ = len(v3)
    v2 = v2
    v28 := GenS2{F0: 42, F1: v24, F2: (-(v1))}
    var v29 int64 = 42
    switch "" {
    case "foo bar":
        v2 = (v25 > (v25 / ((v25) | 1)))
    case "":
        v30 := v3[(v24 / ((v24) | 1))]
        _ = v30
    default:
        var v31 bool = false
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int8, p1 int8, p2 uint64) int32 {
    var v6 int8 = (p0 ^ (p1 & (p1 - p1)))
    if true {
        var v7 int8 = p0
        v8 := GenS1{F0: 1, F1: p2}
        p0 -= v6
        _ = v7
        _ = v8
    } else {
        switch ("" + (("\n\t" + "") + "foo bar")) {
        case "hello":
            fmt.Println(p1, p0)
        default:
            var v9 uint16 = 32767
            _ = v9
        }
        p0 = (p0 - (p1 - p0))
        var v10 float32 = 1e10
        _ = v10
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS2{F0: (^(v6)), F1: p1, F2: 42}
        var v13 uint8 = 1
        var v14 uint = 255
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -315
}
func genFn1(p0 string, p1 uint8, p2 int) uint8 {
    p0 = "a"
    v15 := GenS2{F0: 1, F1: 42, F2: 24}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p1 * 127)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -109
    v3 := []string{"foo bar", "hello"}
    var v4 float32 = (0.1 / float32(0.1))
    v5 := make(map[string]string)
    if (true && (-32768 != -778)) {
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                if true {
                    var v17 uint = 42
                    _ = v17
                }
                v2 -= (v2 - 13)
                _ = fi16
            }
        } else {
            var v18 uint = 18446744073709551615
            var v19 bool = false
            switch -1 {
            case -1:
                _ = len(v3)
            case 545:
                var v20 uint64 = 18446744073709551615
                _ = v20
            case 365:
                v19 = (2147483648 >= 875)
            }
            _ = v18
            _ = v19
        }
    } else {
        v21 := GenS1{F0: genFn1("\n\t", 2, 323), F1: 71}
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 += 100
            _ = fi23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := GenS3{F0: ((("foo bar" + (("\n\t" + "hello") + ((("hello" + "foo bar") + "") + "foo bar"))) + "") + ""), F1: false}
            _ = fi24
            _ = v25
        }
        v4 = v4
        _ = fi22
    }
    fmt.Println(v4)
    v26 := GenS4{F0: 2, F1: -95}
    if false {
        var v27 int64 = 1
        fmt.Println(v2, v27)
        _ = v27
    }
    switch -2147483648 {
    case 127:
        v4 += v4
    case 32768:
        var v28 int8 = (^(v2))
        _ = v28
    }
    v4 += 0.0
    v29 := GenS5{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

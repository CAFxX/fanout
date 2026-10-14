package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 uint, p1 int16, p2 int8) int64 {
    switch 4294967295 {
    case 4294967295:
        v5 := GenS1{F0: 0, F1: ("\n\t" + (("" + ("\n\t" + "\n\t")) + (("foo bar" + "hello") + "\n\t"))), F2: 1}
        _ = v5
    case 0:
        v6 := GenS1{F0: 63, F1: "a", F2: 2}
        _ = v6
    case 593:
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                v8 := GenS2{F0: 42, F1: p0}
                _ = v8
            }
            v9 := GenS1{F0: 619, F1: ((("hello" + "hello") + ("a" + "a")) + ""), F2: 85}
            var v10 int = 1
            _ = fi7
            _ = v10
            _ = v9
        }
    default:
        switch 1 {
        case 999:
            v11 := GenS2{F0: 1, F1: (p0 / ((704) | 1))}
            _ = v11
        case 200:
            v12 := GenS3{F0: 42, F1: ((-(p1)) == int16(p2))}
            _ = v12
        default:
            p1 = 42
        }
    }
    var v13 float32 = ((1e10 * 1e10) - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -1
}
func genFn1(p0 int, p1 uint64) int32 {
    fmt.Println(p1, p0)
    v14 := GenS4{F0: 113, F1: (p1 >> ((4294967295) & 63)), F2: ("hello" + "hello")}
    var v15 string = "hello"
    var v16 float32 = (1.5 * 1.5)
    v17 := []int{(p0 | (p0 >> ((2) & 63))), p0, p0, (p0 & p0)}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -2147483648
    var v3 uint64 = v1
    var v4 bool = true
    v2 += (v2 / ((v2) | 1))
    var v18 int64 = 0
    v19 := GenS2{F0: 128, F1: 2}
    v20 := GenS4{F0: 100, F1: (v1 ^ (v1 + 2)), F2: ("" + ("hello" + "foo bar"))}
    for fi21 := 0; fi21 < 14; fi21++ {
        switch 127 {
        case 9223372036854775807:
            v1 -= v3
        default:
            fmt.Println(v18, v4)
        }
        var v22 bool = v4
        var v23 uint8 = 128
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v1, v2, v18)
    if v4 {
        fmt.Println(v4, v2, v3)
    } else {
        v24 := []int{}
        v2 = (v2 | (v2 | (v2 % ((v2) | 1))))
        _ = v24
    }
    v25 := []int{1}
    fmt.Println(v1)
    v3 = v3
    _ = len(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}

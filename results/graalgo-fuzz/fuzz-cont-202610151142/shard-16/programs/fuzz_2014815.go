package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    if false {
        v6 := GenS1{F0: true, F1: 42}
        p0 += 18446744073709551615
        _ = v6
    }
    if false {
        if true {
            var v7 float64 = 3.14159
            _ = v7
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    if (false || true) {
                        var v10 int8 = 62
                        var v11 int8 = (v10 ^ (v10 >> ((2) & 63)))
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS2{F0: 450, F1: 1}
                        p0 -= p0
                        _ = v12
                    }
                    _ = fi9
                }
                _ = fi8
            }
            v13 := GenS1{F0: (3.14159 > 0.0), F1: 1}
            v14 := GenS3{F0: 753, F1: p0, F2: -32768}
            _ = v13
            _ = v14
        }
        switch -9223372036854775808 {
        case 31:
            if true {
                var v15 int64 = 42
                var v16 uint = 1
                _ = v15
                _ = v16
            }
        default:
            v17 := []int{32768, 42, -2147483648, 1}
            _ = v17
        }
    } else {
        p0 = 42
        p0 += p0
        fmt.Println(p0)
    }
    var v18 uint32 = 42
    var v19 uint32 = 4294967295
    _ = p0
    _ = v18
    _ = v19
    return 146
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    var v20 int8 = -128
    for fi21 := 0; fi21 < 10; fi21++ {
        v20 -= v20
        for fi22 := 0; fi22 < 13; fi22++ {
            v20 += v20
            if false {
                var v23 uint64 = 928
                delete(v3, (fi22 / (((fi22 & 32767)) | 1)))
                _ = v23
            }
            _ = fi22
        }
        _ = fi21
    }
    v24 := GenS2{F0: 42, F1: 1}
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v25 uint64 = 18446744073709551615
    switch ("foo bar" + (("foo bar" + "\n\t") + (((("" + "") + "hello") + ("a" + "foo bar")) + ""))) {
    case "a":
        v26 := GenS1{F0: false, F1: 42}
        _ = v26
    case "\n\t":
        _ = len(v4)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v27 int8 = v20
    for fi28 := 0; fi28 < 8; fi28++ {
        v20 -= v20
        var v29 int16 = 607
        fmt.Println(v25)
        _ = fi28
        _ = v29
    }
    v25 -= v25
    var v30 float64 = 0.0
    v4[2147483647] = ("" + ((("hello" + ("" + ("\n\t" + "\n\t"))) + (((("foo bar" + ("" + ("" + "\n\t"))) + "a") + "hello") + "foo bar")) + "a"))
    v27 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

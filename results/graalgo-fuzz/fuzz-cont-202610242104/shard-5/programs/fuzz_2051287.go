package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint8) bool {
    if false {
        v7 := GenS1{F0: false}
        var v8 int16 = 717
        _ = v7
        _ = v8
    } else {
        if (632 != 255) {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint8 = (p0 / (((p0 / (((p0 / (((p0 * (^(p0)))) | 1))) | 1))) | 1))
                _ = fi9
                _ = v10
            }
            switch "\n\t" {
            case "\n\t":
                var v11 float32 = (-2.25 + (3.14159 * -2.25))
                _ = v11
            case "foo bar":
                v12 := GenS2{F0: 4294967295}
                _ = v12
            case "hello":
                if false {
                    var v13 string = "a"
                    _ = v13
                } else {
                    fmt.Println(p0)
                }
            default:
                var v14 int8 = 127
                _ = v14
            }
            if (!true) {
                var v15 int32 = 0
                var v16 uint64 = 0
                v16 += 0
                _ = v15
                _ = v16
            }
        } else {
            v17 := GenS3{F0: 634, F1: "hello", F2: "foo bar"}
            var v18 uint16 = 32767
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(p0)
                _ = fi19
            }
            _ = v17
            _ = v18
        }
        fmt.Println(p0)
        v20 := []int{-2147483648, 1, 32767, -414}
        _ = v20
    }
    fmt.Println(p0)
    if false {
        v21 := []int{32767}
        p0 += (p0 - p0)
        _ = v21
    }
    _ = p0
    return false
}
func genFn1(p0 int16, p1 uint16, p2 int) int8 {
    if false {
        v22 := GenS4{F0: 0, F1: (false && (-126 == 90))}
        _ = v22
    } else {
        var v23 uint = 259
        v23 += 18446744073709551615
        for fi24 := 0; fi24 < 14; fi24++ {
            switch (("foo bar" + "hello") + "a") {
            case "hello":
                p2 = (fi24 | (p2 << (((^(v23))) & 63)))
            case "foo bar":
                v25 := GenS1{F0: (true && false)}
                _ = v25
            default:
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 uint32 = uint32(p1)
                    _ = fi26
                    _ = v27
                }
            }
            v28 := GenS1{F0: (!(!false))}
            var v29 int32 = -1
            _ = fi24
            _ = v28
            _ = v29
        }
        _ = v23
    }
    var v30 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 4294967295
    v3 := []string{"hello", "foo bar", ("" + "foo bar")}
    v4 := []int{}
    var v5 string = (("hello" + ("hello" + "\n\t")) + ("a" + "\n\t"))
    var v6 int16 = -899
    v6 = 766
    fmt.Println(v6)
    var v31 int = -1
    var v32 int64 = int64(v2)
    v2 -= v2
    fmt.Println(v1)
    _ = len(v4)
    switch v5 {
    case "hello":
        v31 = v2
    }
    fmt.Println(v1, v1)
    v31 += (v31 % ((v2) | 1))
    var v33 float32 = 1.5
    for fi34 := 0; fi34 < 19; fi34++ {
        v35 := GenS5{F0: 42, F1: v1}
        _ = fi34
        _ = v35
    }
    if len(v4) > 0 {
        v4[0] = v31
    }
    v31 = (v2 & v2)
    var v36 uint8 = 246
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

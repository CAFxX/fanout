package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint8) string {
    for fi4 := 0; fi4 < 6; fi4++ {
        p0 -= p0
        switch fi4 {
        case -1:
            v5 := []int{fi4, fi4, (fi4 - fi4), fi4}
            _ = v5
        default:
            var v6 uint = 18446744073709551615
            _ = v6
        }
        _ = fi4
    }
    p0 = 0
    switch 128 {
    case 32768:
        if true {
            var v7 int8 = -1
            _ = v7
        } else {
            var v8 uint64 = 987
            v8 += (v8 & v8)
            _ = v8
        }
    case 9223372036854775807:
        p0 -= p0
    default:
        var v9 uint = 4294967295
        _ = v9
    }
    switch ((("\n\t" + "") + "hello") + "") {
    case "\n\t":
        v10 := []int{0, 834, 127, -1}
        _ = v10
    case "hello":
        var v11 uint32 = 962
        _ = v11
    case "a":
        fmt.Println(p0)
    }
    _ = p0
    return (((("a" + "\n\t") + "\n\t") + "foo bar") + "a")
}

func main() {
    var v1 bool = false
    v2 := []string{}
    v3 := []string{}
    v12 := GenS1{F0: -865, F1: 42}
    v1 = (!v1)
    var v13 int32 = 539
    v1 = (!(-9223372036854775808 == 0))
    switch (("\n\t" + "hello") + ("a" + "hello")) {
    case "":
        var v14 int8 = -128
        _ = v14
    }
    switch "hello" {
    case "foo bar":
        var v15 int = 32767
        _ = v15
    case "\n\t":
        if (-2.25 < (3.14159 - 1e10)) {
            var v16 float64 = (float64(0.0) / 3.14159)
            _ = v16
        }
    case "":
        v13 += v13
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v17 := GenS1{F0: 865, F1: -914}
    v1 = false
    v18 := GenS2{F0: 128, F1: false, F2: -483}
    fmt.Println(v1)
    v19 := GenS1{F0: v13, F1: 1}
    if false {
        switch (("a" + "\n\t") + (("foo bar" + "foo bar") + (((((("foo bar" + "a") + (("foo bar" + (("a" + "foo bar") + "foo bar")) + "")) + ("foo bar" + ("a" + "foo bar"))) + "\n\t") + "a") + "a"))) {
        case "\n\t":
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = 0
                fmt.Println(v13)
                _ = fi20
                _ = v21
            }
        case "foo bar":
            v3 = append(v3, "")
        case "hello":
            v22 := GenS2{F0: -1, F1: (false && v1), F2: -1}
            _ = v22
        default:
            v23 := GenS1{F0: 882, F1: 423}
            _ = v23
        }
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "hello")
    }
    v13 -= 42
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS3{F0: 113, F1: 829}
        _ = fi24
        _ = v25
    }
    var v26 uint = 570
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
}

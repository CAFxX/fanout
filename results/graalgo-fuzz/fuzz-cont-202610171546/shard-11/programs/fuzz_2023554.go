package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
func genFn0(p0 int8, p1 string) int32 {
    p0 = p0
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{fi5, -2147483648, (fi5 * (fi5 * 0))}
        switch p1 {
        case "a":
            fmt.Println(p1, fi5, p1)
        case "foo bar":
            for fi7 := 0; fi7 < 6; fi7++ {
                _ = len(v6)
                var v8 uint = 65535
                _ = fi7
                _ = v8
            }
        case "\n\t":
            fmt.Println(p1)
        default:
            p0 = -128
        }
        _ = fi5
        _ = v6
    }
    switch 9223372036854775807 {
    case 602:
        var v9 int = 128
        _ = v9
    case 32767:
        var v10 uint = 4294967295
        _ = v10
    case 42:
        var v11 string = p1
        _ = v11
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() int {
    if true {
        var v12 int8 = 100
        _ = v12
    }
    switch "hello" {
    case "a":
        var v13 int64 = 75
        _ = v13
    case "":
        var v14 float32 = (-2.25 + float32((0.0 / (1.5 - ((1e10 + 1e10) * 1.5)))))
        _ = v14
    default:
        if false {
            switch "foo bar" {
            case "foo bar":
                v15 := GenS1{F0: 42}
                _ = v15
            case "":
                var v16 float64 = (0.1 / 1e10)
                _ = v16
            case "hello":
                var v17 int8 = 100
                _ = v17
            default:
                var v18 float32 = float32(1e10)
                _ = v18
            }
        }
    }
    v19 := []int{-1, 9223372036854775807}
    v20 := GenS2{F0: -190, F1: 903, F2: 92}
    _ = v19
    _ = v20
    return 128
}

func main() {
    var v1 uint = 785
    v2 := []int{442, 0, 254}
    var v3 int64 = -1
    v4 := make(map[string]uint32)
    if false {
        _ = len(v2)
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS3{F0: 42, F1: (("" + ("foo bar" + "a")) + ((("" + "") + "a") + "\n\t")), F2: 100}
            var v23 uint32 = 86
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            v4[""] = 2
            _ = fi24
        }
    } else {
        fmt.Println(v1)
        if ((-978 > v3) && false) {
            var v25 bool = false
            _ = v25
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            v1 = 18446744073709551615
            _ = fi26
        }
    }
    var v27 float32 = (0.0 * (1e10 - (float32(1e10) + 3.14159)))
    v3 -= -9223372036854775808
    v4["\n\t"] = 2
    switch 1 {
    case 9223372036854775807:
        switch "hello" {
        case "a":
            var v28 string = "foo bar"
            _ = v28
        case "":
            v1 += v1
        }
    case 32767:
        var v29 uint64 = 18446744073709551615
        _ = v29
    case -2147483648:
        switch -779 {
        case 210:
            var v30 uint = v1
            _ = v30
        default:
            var v31 int8 = -128
            _ = v31
        }
    default:
        var v32 int64 = v3
        _ = v32
    }
    if false {
        v1 = 120
        var v33 uint8 = 144
        _ = v33
    } else {
        v4["a"] = 2
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v4[("\n\t" + "\n\t")] = 2147483647
        v1 = (^(v1))
        _ = fi34
    }
    v2 = append(v2, -9223372036854775808)
    v3 -= v3
    switch genFn1() {
    case 466:
        var v35 int32 = genFn0(49, ("foo bar" + ""))
        _ = v35
    default:
        _ = len(v4)
    }
    v4[("a" + "a")] = 2147483647
    var v36 int32 = 42
    fmt.Println(v3, v1, v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
}

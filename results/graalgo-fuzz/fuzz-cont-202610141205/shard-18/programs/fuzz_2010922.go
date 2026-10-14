package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 bool, p1 uint64, p2 uint64) uint {
    v4 := []int{507, -278, 9223372036854775807, 42}
    p1 = 1
    p0 = (!p0)
    var v5 int = -596
    for fi6 := 0; fi6 < 1; fi6++ {
        if p0 {
            v4 = append(v4, 4294967295)
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int32, p1 uint64, p2 uint) uint64 {
    if (("a" + "\n\t") > ((("\n\t" + ("a" + ("" + (("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar")))) + "foo bar") + "foo bar")) {
        if (!(32767 > 42)) {
            v7 := GenS1{F0: 0, F1: 127}
            _ = v7
        } else {
            p2 = (p2 / (((p2 / ((620) | 1))) | 1))
            p2 += p2
        }
    }
    var v8 uint32 = 2
    switch -735 {
    case 32767:
        var v9 int16 = 970
        _ = v9
    case 251:
        for fi10 := 0; fi10 < 20; fi10++ {
            v8 += 0
            p1 = (p1 % (((^(p1))) | 1))
            v11 := []int{(fi10 | (-(fi10))), (-(fi10)), (fi10 - 1), fi10}
            _ = fi10
            _ = v11
        }
    case 2147483647:
        switch ("a" + "\n\t") {
        case "\n\t":
            var v12 uint = genFn0(true, (^(p1)), (p1 + p1))
            _ = v12
        case "foo bar":
            var v13 float32 = (3.14159 * 0.0)
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := GenS2{F0: fi14}
                var v16 uint32 = (^(v8))
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := []int{-9223372036854775808, int(v8), (^(fi17)), (fi17 % ((int(p1)) | 1))}
                _ = fi17
                _ = v18
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p1 ^ 0)
}

func main() {
    var v1 int8 = 127
    v2 := []int{42}
    v3 := make(map[string]bool)
    var v19 int64 = 0
    switch "foo bar" {
    case "a":
        v19 = v19
    }
    switch ("a" + "hello") {
    case "":
        v20 := GenS1{F0: 172, F1: -69}
        _ = v20
    case "\n\t":
        _ = len(v3)
    default:
        fmt.Println(v1, v1)
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int32 = 1
            _ = fi22
            _ = v23
        }
        if (1.5 >= 0.0) {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint = 255
                v1 = 114
                var v26 uint64 = 740
                _ = fi24
                _ = v25
                _ = v26
            }
            v27 := GenS3{F0: ("\n\t" + "a"), F1: 42, F2: -32768}
            _ = v27
        } else {
            var v28 int = (fi21 | (fi21 ^ fi21))
            _ = v28
        }
        _ = fi21
    }
    fmt.Println(v1)
    var v29 string = "\n\t"
    var v30 int64 = 9223372036854775807
    var v31 float64 = -2.25
    v19 = (v30 / ((v19) | 1))
    var v32 uint8 = 42
    var v33 int32 = 43
    var v34 int64 = 1
    var v35 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}

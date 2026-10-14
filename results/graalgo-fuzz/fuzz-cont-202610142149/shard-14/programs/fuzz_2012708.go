package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -1
    v3 := []int{1, 9223372036854775807, -1}
    var v4 float32 = (1e10 + float32(1.5))
    v5 := make(map[string]uint32)
    if ((!false) && true) {
        switch 807 {
        case 4294967295:
            var v6 uint64 = 913
            _ = v6
        case 127:
            v2 = v2
        default:
            v2 += (v2 * (v2 ^ (v2 % (((v2 >> ((0) & 63))) | 1))))
        }
        _ = len(v3)
    }
    if (((((("foo bar" + "a") + ("\n\t" + "a")) + "a") + ("" + "hello")) + "hello") >= (("a" + (("\n\t" + "foo bar") + "")) + (("" + "") + ("foo bar" + (("a" + "foo bar") + ("" + "")))))) {
        v2 = v2
    } else {
        v4 += ((3.14159 / 1.0) + -2.25)
        var v7 uint64 = 4294967295
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        delete(v1, 144)
        _ = fi8
    }
    var v9 int = -9223372036854775808
    v10 := GenS1{F0: 42, F1: 2}
    if (263 == 0) {
        fmt.Println(v2, v2)
        switch 9223372036854775807 {
        case -298:
            var v11 float64 = ((3.14159 / 0.1) * 1.5)
            _ = v11
        case 589:
            var v12 uint32 = 1
            _ = v12
        case 4294967295:
            if (!false) {
                fmt.Println(v4)
            }
        default:
            v13 := GenS2{F0: 98, F1: 126, F2: ("hello" + "\n\t")}
            _ = v13
        }
        v2 = v2
    }
    var v14 uint32 = 433
    _ = len(v1)
    switch "hello" {
    case "":
        if (!(42 == 472)) {
            var v15 int64 = int64(v9)
            _ = v15
        }
    case "foo bar":
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 string = ("foo bar" + "")
            _ = fi16
            _ = v17
        }
    }
    switch v9 {
    case 668:
        v18 := GenS3{F0: v14, F1: ((402 != 1) && true)}
        _ = v18
    case 2147483647:
        for fi19 := 0; fi19 < 3; fi19++ {
            delete(v5, "a")
            v14 -= (v14 & 65535)
            if (true || true) {
                var v20 uint64 = 1
                var v21 int8 = 1
                _ = v20
                _ = v21
            }
            _ = fi19
        }
    case 0:
        switch "foo bar" {
        case "hello":
            for fi22 := 0; fi22 < 4; fi22++ {
                v4 = (v4 - float32((-2.25 + float64((0.1 * (v4 * (v4 / v4)))))))
                _ = fi22
            }
        case "\n\t":
            if false {
                var v23 uint32 = (v14 + 14)
                for fi24 := 0; fi24 < 18; fi24++ {
                    v14 -= v23
                    var v25 uint = 42
                    _ = fi24
                    _ = v25
                }
                delete(v1, (^(v14)))
                _ = v23
            }
        default:
            v26 := GenS2{F0: 2, F1: (v14 + v14), F2: ((("" + ("\n\t" + (("hello" + "foo bar") + ("foo bar" + "\n\t")))) + "") + "\n\t")}
            _ = v26
        }
    }
    var v27 int32 = 42
    v2 -= v2
    v28 := GenS3{F0: 2, F1: false}
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

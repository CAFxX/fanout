package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool) int32 {
    var v5 int32 = 928
    var v6 int32 = v5
    v7 := []int{0, 9223372036854775807}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (v5 + 1)
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 42
    v4 := []int{-2147483648, 0, -1, 4294967295}
    if (!(!true)) {
        v3 -= (^(v3))
        var v8 string = "\n\t"
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            switch "foo bar" {
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v3 = (^(v3))
                        _ = fi10
                    }
                default:
                    v11 := GenS1{F0: 42}
                    _ = v11
                }
            }
            var v12 string = ("hello" + ("\n\t" + (("" + (("\n\t" + ("foo bar" + "foo bar")) + (("\n\t" + (("foo bar" + ("\n\t" + "")) + "hello")) + ("hello" + "hello")))) + "")))
            _ = fi9
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                switch ("a" + "hello") {
                case "":
                    var v15 float64 = (1e10 + 0.0)
                    _ = v15
                default:
                    v3 = 4294967295
                }
                v16 := GenS1{F0: 733}
                v3 = (v3 ^ 4294967295)
                _ = fi14
                _ = v16
            }
            _ = fi13
        }
    }
    v17 := GenS2{F0: true, F1: ("" + ("foo bar" + "foo bar")), F2: -2147483648}
    v3 = 0
    if (0 == 791) {
        var v18 uint8 = 1
        _ = v18
    }
    var v19 string = ("\n\t" + "hello")
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS3{F0: 0}
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 65535
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v3, fi23)
        v22 += v22
        var v24 float32 = 1.5
        _ = fi23
        _ = v24
    }
    switch "foo bar" {
    case "\n\t":
        var v25 uint = 2
        _ = v25
    case "a":
        v19 = v19
    default:
        v22 += v22
    }
    var v26 string = (v19 + v19)
    if (-1 <= 32767) {
        fmt.Println(v22)
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := GenS3{F0: 598}
            _ = fi27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v19 = ((v19 + (("\n\t" + ("a" + (v19 + ((v26 + ("" + v26)) + v19)))) + v19)) + "hello")
        _ = fi29
    }
    v30 := GenS3{F0: 9223372036854775807}
    var v31 uint64 = 42
    v22 -= v22
    if true {
        v32 := GenS3{F0: 42}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

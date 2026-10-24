package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0() bool {
    var v7 uint64 = 0
    v8 := []int{128, 128, -130}
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            if len(v8) > 0 {
                v8[0] = 533
            }
            _ = fi10
        }
        _ = fi9
    }
    var v11 uint = 255
    switch ((("hello" + "a") + ("\n\t" + ("foo bar" + ""))) + (("a" + "foo bar") + "")) {
    case "\n\t":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if (1e10 != 3.14159) {
            var v12 float32 = (1e10 + (1e10 + (-2.25 + 0.1)))
            _ = len(v8)
            _ = v12
        }
    }
    _ = v11
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 float64 = -2.25
    v4 := []uint32{(v1 >> ((440) & 63)), v1, v1}
    var v5 uint64 = 462
    var v6 int8 = -128
    fmt.Println(v3, v1, v5)
    v5 -= (v5 % (((^(v5))) | 1))
    switch "a" {
    case "a":
        v6 -= v6
    case "\n\t":
        v13 := GenS1{F0: 2147483648, F1: "", F2: "a"}
        _ = v13
    case "foo bar":
        v14 := GenS1{F0: (v1 ^ (v1 / ((759) | 1))), F1: "a", F2: ""}
        _ = v14
    default:
        v3 -= -2.25
    }
    v6 = -128
    v3 = (v3 * v3)
    v15 := GenS1{F0: v1, F1: "\n\t", F2: ""}
    _ = len(v4)
    v16 := GenS2{F0: -1, F1: 1, F2: 42}
    if ("\n\t" != "hello") {
        if false {
            v2[v1] = (false && false)
            var v17 uint = 319
            v1 += (v1 ^ (v1 - 1))
            _ = v17
        }
        if true {
            v2[(v1 * (^(v1)))] = (true || false)
            v18 := GenS1{F0: 2, F1: "", F2: "\n\t"}
            v5 = v5
            _ = v18
        } else {
            switch "foo bar" {
            case "\n\t":
                v6 -= 120
            case "foo bar":
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int8 = (v6 * v6)
                    _ = fi19
                    _ = v20
                }
            case "":
                v6 += v6
            }
            var v21 uint64 = (^(v5))
            var v22 uint64 = (v5 - (v21 % (((v5 * 42)) | 1)))
            _ = v21
            _ = v22
        }
    } else {
        var v23 int64 = 42
        _ = v23
    }
    if false {
        var v24 string = (("foo bar" + "foo bar") + "")
        _ = v24
    }
    var v25 int = 9223372036854775807
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int = (fi26 | 32767)
        _ = fi26
        _ = v27
    }
    v3 -= v3
    for fi28 := 0; fi28 < 3; fi28++ {
        if false {
            v5 -= v5
            if false {
                v29 := GenS3{F0: 1, F1: 925}
                v30 := GenS2{F0: 889, F1: 255, F2: int(v1)}
                _ = v29
                _ = v30
            } else {
                delete(v2, v1)
            }
        } else {
            var v31 int = (v25 * (fi28 * fi28))
            _ = v31
        }
        v32 := GenS1{F0: 508, F1: "", F2: ""}
        _ = fi28
        _ = v32
    }
    fmt.Println(v25)
    var v33 uint = 832
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

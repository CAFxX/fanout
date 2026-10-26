package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
func genFn0(p0 uint16) bool {
    var v7 int8 = 0
    var v8 string = ("a" + (("hello" + ("foo bar" + "foo bar")) + ("a" + "foo bar")))
    _ = p0
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint64) int {
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{-1, fi9, 1}
        _ = fi9
        _ = v10
    }
    if true {
        var v11 int = 0
        _ = v11
    }
    p0 += (p0 & (p0 * (p0 | 42)))
    p0 -= p0
    _ = p0
    return 128
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = -358
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    v5 := []string{"\n\t", ("hello" + ""), ""}
    var v6 int32 = -1
    v1 += (v1 / (((v1 * v1)) | 1))
    var v12 int64 = (-(v2))
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if ((!true) || ((!true) || false)) {
                var v15 string = "a"
                v1 = 744
                _ = v15
            } else {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                v3 = v3
                v12 = 9223372036854775807
            }
            v12 = v12
            _ = fi14
        }
        fmt.Println(v1, fi13)
        _ = fi13
    }
    v3 = (v3 - v3)
    switch -1 {
    case 32767:
        var v16 int = 428
        _ = v16
    case 127:
        switch 128 {
        case 330:
            switch "hello" {
            case "hello":
                var v17 uint16 = 255
                _ = v17
            case "foo bar":
                v1 += v1
            case "a":
                v2 = (v2 & (v12 % (((v2 - (v2 << ((255) & 63)))) | 1)))
            }
        case -1:
            if false {
                v12 += (v12 | (v2 + v2))
                v12 = (-(v2))
                v2 += (-(v2))
            }
        case 128:
            if true {
                var v18 int16 = 42
                v19 := GenS1{F0: 2147483648}
                var v20 float32 = 0.0
                _ = v18
                _ = v19
                _ = v20
            } else {
                var v21 uint = 0
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(v6, v21, v6)
                    _ = fi22
                }
                _ = v21
            }
        }
    default:
        v23 := GenS1{F0: 4294967295}
        _ = v23
    }
    var v24 int = 32768
    v1 = 2
    v1 = 496
    v25 := GenS1{F0: 1}
    delete(v4, "hello")
    var v26 uint64 = v1
    v27 := GenS1{F0: 2147483648}
    var v28 float64 = 1e10
    v29 := GenS1{F0: 2147483647}
    v30 := GenS2{F0: ("foo bar" + ""), F1: false, F2: 794}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

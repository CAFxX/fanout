package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32; F1 string; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 uint64, p1 uint64, p2 int8) uint {
    if false {
        if false {
            fmt.Println(p1, p0)
        }
        switch ("hello" + "\n\t") {
        case "foo bar":
            p2 = (p2 + p2)
        case "\n\t":
            var v4 uint32 = 2147483648
            _ = v4
        case "a":
            var v5 float32 = 0.0
            _ = v5
        }
    } else {
        p2 -= (p2 & (p2 - (p2 / ((p2) | 1))))
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 bool = false
        var v8 int32 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 345 {
    case 288:
        v9 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807}
        _ = v9
    case 127:
        if ((18 >= -9223372036854775808) || true) {
            v10 := []int{}
            v11 := GenS2{F0: 127, F1: ""}
            _ = v10
            _ = v11
        }
    case 42:
        v12 := []int{-430, 826, -2147483648}
        _ = v12
    default:
        var v13 uint64 = (p0 / ((p0) | 1))
        _ = v13
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := []int{414, 127, 501, 4294967295}
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    v2 -= (v2 | v2)
    _ = len(v3)
    switch "" {
    case "foo bar":
        for fi14 := 0; fi14 < 5; fi14++ {
            switch ("a" + (("hello" + "") + ("foo bar" + ""))) {
            case "\n\t":
                var v15 int = fi14
                _ = v15
            case "a":
                var v16 uint32 = 922
                _ = v16
            default:
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int = fi17
                    v18 = v18
                    v19 := v3["foo bar"]
                    v1 = append(v1, fi14)
                    var v20 uint32 = 2147483647
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
            v2 -= 255
            _ = fi14
        }
    }
    switch -383 {
    case 4294967295:
        for fi21 := 0; fi21 < 19; fi21++ {
            v2 -= 242
            v2 += 8
            _ = fi21
        }
    case -9223372036854775808:
        var v22 uint16 = 32767
        _ = v22
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v2)
    switch "" {
    case "\n\t":
        for fi23 := 0; fi23 < 5; fi23++ {
            v2 = v2
            _ = fi23
        }
    case "hello":
        v24 := GenS3{F0: 42, F1: "", F2: 958}
        _ = v24
    case "":
        if true {
            fmt.Println(v2)
            v1 = append(v1, 42)
            v25 := GenS4{F0: 319}
            _ = v25
        } else {
            if false {
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint16 = (^(v2))
                    v2 = (^(v2))
                    _ = fi26
                    _ = v27
                }
                var v28 int32 = 42
                var v29 float32 = 1.5
                _ = v28
                _ = v29
            } else {
                var v30 int = 128
                v31 := GenS4{F0: 4294967295}
                _ = v30
                _ = v31
            }
            v2 += 32767
        }
    }
    v2 = (^(v2))
    v2 = (v2 / (((^(v2))) | 1))
    if ((1e10 / (0.0 + ((1e10 + (-2.25 - 0.0)) / (1e10 - -2.25)))) > (0.1 + 1e10)) {
        fmt.Println(v2)
        v32 := GenS5{F0: 2, F1: ""}
        _ = v32
    }
    v2 = (^(v2))
    var v33 int32 = 6
    var v34 int8 = 100
    for fi35 := 0; fi35 < 19; fi35++ {
        for fi36 := 0; fi36 < 13; fi36++ {
            fmt.Println(v34, v33, v34)
            var v37 bool = (733 == -202)
            _ = fi36
            _ = v37
        }
        v38 := GenS4{F0: 1}
        _ = fi35
        _ = v38
    }
    v34 = v34
    var v39 uint64 = 929
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: ""}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        switch "\n\t" {
        case "hello":
            if (645 < 2) {
                var v10 uint = 42
                fmt.Println(v10)
                _ = v10
            } else {
                switch "a" {
                case "hello":
                    var v11 uint8 = 1
                    _ = v11
                default:
                    v12 := GenS2{F0: -93, F1: ((fi9 / (((fi9 >> ((42) & 63))) | 1)) < -2147483648), F2: 314}
                    _ = v12
                }
            }
        default:
            var v13 int = fi9
            _ = v13
        }
        if (1 > 0) {
            var v14 float32 = float32((0.0 + 0.1))
            v14 -= float32(0.0)
            var v15 int8 = -61
            _ = v14
            _ = v15
        } else {
            var v16 uint64 = 2
            var v17 uint8 = 0
            _ = v16
            _ = v17
        }
        if (!false) {
            var v18 uint8 = 128
            var v19 float64 = 3.14159
            _ = v18
            _ = v19
        }
        _ = fi9
    }
    return 285
}
func genFn1(p0 uint64, p1 int8, p2 string) int16 {
    var v20 int = 128
    v21 := []int{}
    var v22 int = 0
    var v23 int = (v20 | (v20 + (^(v20))))
    var v24 uint64 = (p0 ^ p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 32767
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 91
    v3 := []string{(("hello" + "") + "foo bar")}
    v4 := make(map[string]bool)
    var v5 int8 = -128
    v6 := make(map[string]uint32)
    var v25 int8 = 0
    if true {
        v2 += v2
        var v26 float32 = (0.0 / 3.14159)
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        if true {
            fmt.Println(v1, fi27, v5)
            var v28 int32 = 1
            _ = v28
        } else {
            _ = len(v3)
            if true {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS2{F0: int16(v1), F1: ((("\n\t" + "") + "foo bar") < ("" + (((("a" + ("" + "\n\t")) + (("a" + "hello") + "a")) + ((("hello" + "hello") + "") + "foo bar")) + ""))), F2: 2147483648}
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := v6[("a" + "a")]
                    _ = fi31
                    _ = v32
                }
            }
        }
        if false {
            v2 = 1
            if false {
                v1 -= v1
            } else {
                var v33 uint8 = uint8(v2)
                var v34 int8 = v25
                var v35 bool = false
                _ = v33
                _ = v34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            v2 = (v2 % ((v2) | 1))
            var v37 int32 = -601
            _ = len(v3)
            _ = fi36
            _ = v37
        }
        _ = fi27
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        v2 += (v2 - 2)
        v5 = (v25 - v5)
        _ = fi38
    }
    v39 := GenS1{F0: ("foo bar" + "foo bar")}
    var v40 uint8 = 42
    var v41 bool = true
    v1 = (v1 / ((uint16(v40)) | 1))
    for fi42 := 0; fi42 < 9; fi42++ {
        fmt.Println(v2, v1, v41)
        v43 := v6["hello"]
        _ = fi42
        _ = v43
    }
    fmt.Println(v1, v25, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}

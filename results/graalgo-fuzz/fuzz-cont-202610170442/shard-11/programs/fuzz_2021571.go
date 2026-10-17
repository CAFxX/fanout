package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 int32, p1 bool) uint64 {
    p1 = ("" != "hello")
    v5 := []int{}
    p1 = false
    p0 += p0
    _ = p0
    _ = p1
    _ = v5
    return 343
}
func genFn1() uint16 {
    var v6 int8 = 100
    switch 9223372036854775807 {
    case 210:
        var v7 string = ""
        _ = v7
    }
    v8 := GenS1{F0: v6, F1: 33, F2: v6}
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{32768}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    fmt.Println(v6)
    _ = v6
    _ = v8
    return 947
}

func main() {
    var v1 string = (("hello" + ("foo bar" + "foo bar")) + "\n\t")
    var v2 uint8 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v12 float32 = -2.25
    switch -2147483648 {
    case 9223372036854775807:
        if false {
            v2 = 255
        } else {
            switch v1 {
            case "":
                if true {
                    var v13 int32 = -788
                    if (1 >= 1) {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            var v15 int16 = 36
                            v16 := GenS1{F0: -42, F1: 4294967295, F2: 127}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        fmt.Println(v2, v2)
                    } else {
                        v12 -= v12
                    }
                    if true {
                        v13 -= (v13 % ((v13) | 1))
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v1 = ""
                            v18 := GenS1{F0: -121, F1: 1, F2: -27}
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = v13
                }
            case "foo bar":
                if true {
                    v19 := []int{}
                    _ = v19
                } else {
                    var v20 uint32 = 0
                    _ = v20
                }
            default:
                v21 := GenS2{F0: 280, F1: (!true)}
                _ = v21
            }
            var v22 int64 = 945
            v1 = v1
            _ = v22
        }
    case 4294967295:
        var v23 string = "\n\t"
        _ = v23
    }
    v1 = ((((("a" + "\n\t") + ("a" + v1)) + "a") + v1) + v1)
    v1 = "foo bar"
    if false {
        var v24 int64 = -9223372036854775808
        var v25 uint64 = 4294967295
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            v28 := []int{(fi26 + fi27)}
            var v29 int8 = -1
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    var v30 uint8 = (v2 & v2)
    v31 := []int{1}
    if (9223372036854775807 <= 1) {
        var v32 uint32 = 0
        _ = v32
    } else {
        if false {
            fmt.Println(v1)
            for fi33 := 0; fi33 < 3; fi33++ {
                if (65535 <= 255) {
                    fmt.Println(v12, v1)
                    switch (fi33 ^ -1) {
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "":
                            fmt.Println(v1)
                        case "foo bar":
                            var v34 float32 = (float32(((0.1 + (((0.0 - -2.25) / ((1.5 + 1e10) - float64(v12))) * 0.0)) - -2.25)) + v12)
                            _ = v34
                        }
                    case -595:
                        var v35 bool = (1 < 1)
                        _ = v35
                    default:
                        if (false || (!true)) {
                            v2 += (v30 / (((v2 / (((v2 << ((2) & 63))) | 1))) | 1))
                        }
                    }
                }
                _ = fi33
            }
            v30 += (v30 / ((v2) | 1))
        }
        switch 128 {
        case 128:
            if len(v31) > 0 {
                v31[0] = 42
            }
        case -204:
            v4["a"] = ((2 != 584) || true)
        case 222:
            var v36 int64 = 42
            _ = v36
        default:
            if true {
                v37 := v3[893]
                _ = v37
            }
        }
        var v38 uint64 = 737
        _ = v38
    }
    fmt.Println(v30, v12)
    v12 = v12
    v2 += 2
    v39 := GenS1{F0: 48, F1: 42, F2: 100}
    var v40 uint64 = 266
    var v41 int16 = 1
    v2 += 127
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}

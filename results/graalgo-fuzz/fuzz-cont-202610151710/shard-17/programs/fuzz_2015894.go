package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int64) int {
    p0 -= -1
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    p0 -= -1
    _ = p0
    return 0
}
func genFn1(p0 int32) uint16 {
    fmt.Println(p0)
    if (!(!(false && (-1 <= 32768)))) {
        var v5 int8 = 1
        _ = v5
    } else {
        v6 := GenS1{F0: 42, F1: 119, F2: (p0 & (p0 | p0))}
        if (107 == 1) {
            v7 := GenS1{F0: 42, F1: 1, F2: p0}
            var v8 uint8 = 137
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: 0}
        _ = v6
        _ = v9
    }
    v10 := []int{genFn0(-910), -314}
    _ = p0
    _ = v10
    return 286
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 177
    var v3 uint64 = 42
    v11 := GenS1{F0: 604, F1: 1, F2: -2147483648}
    var v12 float64 = ((((-2.25 + 0.0) + -2.25) + 1e10) / 1.5)
    switch 4294967295 {
    case 32767:
        for fi13 := 0; fi13 < 18; fi13++ {
            switch ("" + ("hello" + "")) {
            case "a":
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int16 = 1
                    _ = fi14
                    _ = v15
                }
            case "hello":
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (true || (1 > -788)) {
                        switch fi16 {
                        case 127:
                            var v17 int = (fi16 | (fi16 ^ genFn0(-9223372036854775808)))
                            _ = v17
                        case 32768:
                            var v18 int32 = 1
                            _ = v18
                        case 153:
                            var v19 int64 = -1
                            _ = v19
                        default:
                            fmt.Println(fi16, v3)
                        }
                        if false {
                            v20 := GenS3{F0: -1}
                            v12 += v12
                            _ = v20
                        } else {
                            fmt.Println(v2)
                            v12 = 1.5
                            v21 := GenS4{F0: 768, F1: (v3 & (v3 & 42)), F2: 121}
                            _ = v21
                        }
                        fmt.Println(fi16)
                    }
                    _ = fi16
                }
            case "\n\t":
                var v22 int32 = 1
                _ = v22
            default:
                v2 = v2
            }
            v2 += v2
            _ = fi13
        }
    case 897:
        var v23 string = "a"
        _ = v23
    case 614:
        var v24 uint16 = 42
        _ = v24
    default:
        var v25 uint = v2
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 bool = (fi26 >= fi26)
        _ = fi26
        _ = v27
    }
    fmt.Println(v12)
    for fi28 := 0; fi28 < 10; fi28++ {
        v12 = (3.14159 - 3.14159)
        for fi29 := 0; fi29 < 11; fi29++ {
            switch fi28 {
            case -32:
                var v30 uint = (^(v2))
                v2 -= v2
                _ = v30
            }
            v31 := GenS5{F0: 689, F1: "hello", F2: (("a" + "foo bar") + "")}
            if (!false) {
                var v32 uint = v2
                if false {
                    v32 = (v2 - v2)
                    var v33 uint16 = genFn1(-105)
                    v34 := GenS6{F0: 242}
                    _ = v33
                    _ = v34
                }
                fmt.Println(fi29)
                _ = v32
            }
            _ = fi29
            _ = v31
        }
        _ = fi28
    }
    _ = len(v1)
    for fi35 := 0; fi35 < 6; fi35++ {
        if (!(-106 > 1)) {
            v36 := []int{-598}
            fmt.Println(fi35, v12)
            var v37 int64 = -1
            _ = v36
            _ = v37
        }
        var v38 uint16 = 944
        _ = fi35
        _ = v38
    }
    fmt.Println(v2, v2, v12)
    v12 -= v12
    var v39 uint64 = (v3 - v3)
    var v40 int32 = 42
    _ = len(v1)
    var v41 string = ("\n\t" + "")
    if ((0.0 - -2.25) < ((float32(1e10) / 3.14159) - (1e10 * 1.5))) {
        for fi42 := 0; fi42 < 17; fi42++ {
            v43 := []int{(fi42 ^ fi42), (fi42 >> ((v2) & 63)), fi42, (fi42 / (((-(fi42))) | 1))}
            var v44 int8 = 91
            _ = fi42
            _ = v43
            _ = v44
        }
        delete(v1, -806)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint32; F2 string }
func genFn0(p0 uint64) uint16 {
    p0 = (p0 - 0)
    v7 := GenS1{F0: -9223372036854775808, F1: p0}
    switch (("a" + ("foo bar" + "hello")) + "a") {
    case "":
        var v8 bool = true
        _ = v8
    case "\n\t":
        if true {
            var v9 float64 = float64(-2.25)
            if (("a" + "hello") == "\n\t") {
                var v10 uint8 = 255
                _ = v10
            } else {
                var v11 uint16 = 255
                v11 = v11
                if true {
                    v12 := GenS2{F0: 0, F1: 32767}
                    v13 := GenS1{F0: -686, F1: p0}
                    _ = v12
                    _ = v13
                } else {
                    v14 := []int{883, 993, 2147483647}
                    v9 = (float64(1.5) * 0.1)
                    _ = v14
                }
                _ = v11
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v9 += v9
                _ = fi15
            }
            _ = v9
        }
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        p0 += p0
        _ = fi16
    }
    _ = p0
    _ = v7
    return 255
}
func genFn1() int64 {
    switch 32767 {
    case 820:
        if true {
            var v17 uint = 0
            _ = v17
        }
    default:
        var v18 int32 = 0
        _ = v18
    }
    var v19 string = ((("a" + "hello") + ("foo bar" + "")) + "")
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := []int{(^(fi20)), 127, fi20}
        _ = fi20
        _ = v21
    }
    if true {
        v19 = ("foo bar" + v19)
        var v22 uint = 1
        _ = v22
    }
    v19 = ""
    _ = v19
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (float32(-2.25) / (0.0 + ((0.1 * 1e10) - -2.25)))
    v3 := make(map[int]string)
    var v4 uint32 = 42
    var v5 uint8 = 1
    var v6 int64 = 0
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS3{F0: -128}
        v1[(fi23 / (((fi23 * (fi23 << ((1) & 63)))) | 1))] = ("hello" + "hello")
        v2 += v2
        _ = fi23
        _ = v24
    }
    v5 -= 87
    v6 -= genFn1()
    v25 := GenS4{F0: "foo bar", F1: ("foo bar" + ("a" + "a"))}
    if (false || false) {
        v2 += (v2 / v2)
        v26 := GenS3{F0: 100}
        var v27 uint16 = 255
        _ = v26
        _ = v27
    }
    switch 850 {
    case 2147483647:
        if (true && (525 == 721)) {
            v28 := GenS5{F0: 65}
            _ = v28
        }
    default:
        if false {
            if (!false) {
                var v29 int = int(v4)
                var v30 int32 = 0
                for fi31 := 0; fi31 < 9; fi31++ {
                    v5 = (v5 - v5)
                    for fi32 := 0; fi32 < 3; fi32++ {
                        switch fi31 {
                        case -527:
                            v29 = (fi31 ^ (^(fi31)))
                        case 127:
                            var v33 float64 = (((0.0 / ((1.5 * 1.5) + ((((3.14159 / -2.25) + 3.14159) - 0.0) + 1e10))) - 0.1) - 0.0)
                            _ = v33
                        }
                        var v34 int = 4294967295
                        _ = fi32
                        _ = v34
                    }
                    _ = fi31
                }
                _ = v29
                _ = v30
            }
            v35 := GenS3{F0: 42}
            _ = v35
        } else {
            v36 := GenS6{F0: ((416 != 678) && ((false && ((v2 == v2) || true)) && true)), F1: v4, F2: ("\n\t" + "\n\t")}
            v2 = v2
            v4 += (v4 ^ (^(v4)))
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 float32 = float32((((float64(v2) + 1.5) - 3.14159) - (0.1 + (3.14159 + ((1.5 + 0.1) - 1.5)))))
            fmt.Println(v2, v6)
            var v40 float32 = float32(1.5)
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi37
    }
    var v41 int = 57
    var v42 uint = 1
    if (1e10 < (1e10 / -2.25)) {
        if (false || true) {
            v43 := []int{}
            _ = v43
        } else {
            for fi44 := 0; fi44 < 11; fi44++ {
                v42 -= 65535
                if false {
                    v45 := GenS1{F0: (fi44 - fi44), F1: 1}
                    v6 -= genFn1()
                    _ = v45
                } else {
                    if false {
                        switch "a" {
                        case "":
                            var v46 int16 = 1
                            _ = v46
                        case "foo bar":
                            var v47 bool = (-1 != -119)
                            _ = v47
                        }
                    }
                }
                for fi48 := 0; fi48 < 12; fi48++ {
                    switch (fi44 - fi48) {
                    case -9223372036854775808:
                        var v49 int = (-(fi48))
                        _ = v49
                    default:
                        v4 += (v4 ^ 892)
                    }
                    _ = fi48
                }
                _ = fi44
            }
        }
        var v50 int = v41
        var v51 uint16 = 2
        _ = v50
        _ = v51
    }
    var v52 int16 = -32768
    delete(v3, (v41 - v41))
    v53 := GenS4{F0: "a", F1: ""}
    switch v41 {
    case -1:
        var v54 string = ("hello" + ("\n\t" + "\n\t"))
        _ = v54
    case -991:
        v55 := []int{}
        _ = v55
    default:
        v56 := GenS5{F0: 1}
        _ = v56
    }
    if true {
        if (!true) {
            v41 += (-(v41))
            if true {
                v57 := GenS4{F0: "\n\t", F1: ""}
                _ = v57
            } else {
                v2 = float32(-2.25)
                var v58 int8 = -128
                _ = v58
            }
        }
        v59 := []int{v41}
        _ = v59
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        v41 = 0
        if (v4 == v4) {
            v5 -= v5
            delete(v3, 2147483647)
            if false {
                delete(v3, (fi60 << ((42) & 63)))
            } else {
                v42 -= v42
                v61 := GenS4{F0: "a", F1: ("a" + "")}
                _ = v61
            }
        } else {
            v52 = v52
            var v62 int8 = 42
            _ = v62
        }
        fmt.Println(v4)
        _ = fi60
    }
    v63 := []int{v41, 283, v41, v41}
    var v64 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v63
    _ = v64
}

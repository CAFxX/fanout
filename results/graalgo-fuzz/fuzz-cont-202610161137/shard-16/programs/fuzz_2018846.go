package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0(p0 int16) uint64 {
    if (((("\n\t" + "a") + "hello") + "") > "\n\t") {
        var v5 int8 = -128
        _ = v5
    }
    var v6 int32 = 456
    fmt.Println(p0)
    var v7 int64 = -385
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 string, p1 uint16) uint8 {
    v8 := []int{-627, 127, 32767, -9223372036854775808}
    for fi9 := 0; fi9 < 4; fi9++ {
        p1 -= 65535
        fmt.Println(p0, p0)
        fmt.Println(p1, p0, p0)
        _ = fi9
    }
    if true {
        fmt.Println(p1, p1)
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float64 = 1e10
            if false {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v11 -= v11
                    v8 = append(v8, fi12)
                    _ = fi12
                }
                fmt.Println(fi10, fi10)
                if false {
                    v13 := GenS1{F0: (p0 + p0)}
                    if ((false && ((("a" + p0) >= (p0 + (((("foo bar" + p0) + ("a" + ("" + p0))) + p0) + p0))) || (176 <= 0))) || (true || (1 <= 1))) {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := GenS1{F0: p0}
                            var v16 uint64 = 42
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint16 = (p1 % ((843) | 1))
                        _ = v17
                    }
                    _ = v13
                } else {
                    var v18 uint8 = 187
                    _ = v18
                }
            } else {
                p0 = "foo bar"
            }
            switch fi10 {
            case 4294967295:
                switch p0 {
                case "":
                    var v19 int = (-(fi10))
                    _ = v19
                case "a":
                    v11 -= v11
                }
            case 128:
                for fi20 := 0; fi20 < 13; fi20++ {
                    p0 = (p0 + (p0 + "\n\t"))
                    p0 = "a"
                    v21 := GenS2{F0: genFn0(-32768), F1: p0}
                    v11 = float64(0.1)
                    _ = fi20
                    _ = v21
                }
            default:
                var v22 float64 = (v11 - v11)
                _ = v22
            }
            _ = fi10
            _ = v11
        }
        v23 := GenS3{F0: 255}
        _ = v23
    } else {
        switch 32767 {
        case 231:
            v24 := GenS4{F0: 46, F1: p1}
            _ = v24
        default:
            p0 = (p0 + "")
        }
        var v25 int32 = 881
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int8 = 78
            v28 := GenS4{F0: 4294967295, F1: (p1 + (p1 | p1))}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 952
    v4 := make(map[uint32]uint32)
    delete(v2, 586)
    switch "" {
    case "a":
        var v29 uint64 = 2
        _ = v29
    case "":
        for fi30 := 0; fi30 < 16; fi30++ {
            switch ("hello" + "") {
            case "":
                v3 -= v3
            case "\n\t":
                v1[fi30] = (-32768 == -32768)
            default:
                var v31 bool = (!((v3 / ((v3) | 1)) >= v3))
                _ = v31
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                v33 := GenS2{F0: genFn0(1000), F1: (("a" + (("hello" + "hello") + "\n\t")) + ("a" + ""))}
                if false {
                    switch "\n\t" {
                    case "":
                        var v34 int16 = -1
                        _ = v34
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            var v35 uint32 = 2147483647
                            _ = v35
                        case "foo bar":
                            var v36 bool = (1 <= uint32(v3))
                            if (65535 < 66) {
                                var v37 uint16 = 255
                                var v38 uint64 = 2
                                v36 = (127 > 0)
                                _ = v37
                                _ = v38
                            }
                            fmt.Println(v3)
                            _ = v36
                        case "a":
                            var v39 int64 = int64(v3)
                            _ = v39
                        }
                    }
                    switch fi32 {
                    case 9223372036854775807:
                        var v40 int32 = 42
                        _ = v40
                    default:
                        var v41 float32 = (float32(3.14159) * (((float32(-2.25) + 1e10) * (float32(((((1.5 * (((1e10 / (3.14159 / (1.5 / (-2.25 / 1.0)))) * (0.0 * 0.1)) + 0.0)) / 1.5) * 1.5) / 1.0)) - 0.0)) + 1.5))
                        v3 = (v3 | 2)
                        var v42 int32 = 0
                        _ = v41
                        _ = v42
                    }
                } else {
                    v3 += (v3 / (((v3 / ((699) | 1))) | 1))
                    if (false && (true && ((v3 / ((v3) | 1)) == (^(v3))))) {
                        v3 += (^(v3))
                    }
                }
                _ = fi32
                _ = v33
            }
            _ = fi30
        }
    default:
        var v43 uint32 = 2
        _ = v43
    }
    if false {
        for fi44 := 0; fi44 < 4; fi44++ {
            v3 -= (v3 >> ((4294967295) & 63))
            var v45 int32 = -919
            var v46 uint16 = 32767
            _ = fi44
            _ = v45
            _ = v46
        }
        if true {
            fmt.Println(v3)
            var v47 uint16 = v3
            _ = v47
        }
    }
    if true {
        v48 := v2[42]
        if (2 > 1) {
            v49 := GenS1{F0: v48}
            v50 := GenS5{F0: 0, F1: 42}
            _ = v49
            _ = v50
        }
        _ = v48
    }
    v3 = v3
    var v51 uint8 = 2
    var v52 int16 = 0
    v51 -= (v51 | genFn1(("hello" + ""), 0))
    var v53 bool = (("\n\t" < ("foo bar" + "")) && ((float64(0.1) / 1.5) != (1e10 * 1e10)))
    var v54 uint8 = 255
    v55 := GenS4{F0: 65535, F1: (v3 % ((42) | 1))}
    v56 := []int{}
    for fi57 := 0; fi57 < 7; fi57++ {
        _ = len(v2)
        var v58 int8 = 1
        v53 = v53
        _ = fi57
        _ = v58
    }
    var v59 uint64 = 2
    v60 := GenS5{F0: 143, F1: v59}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v59
    _ = v60
}

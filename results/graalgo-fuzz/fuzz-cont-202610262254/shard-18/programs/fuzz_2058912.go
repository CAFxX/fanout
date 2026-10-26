package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 bool) int64 {
    if p0 {
        fmt.Println(p0)
        var v4 uint = 117
        p0 = (p0 && (2 != 510))
        _ = v4
    }
    p0 = (p0 || (!((!p0) || (171 > 4294967295))))
    _ = p0
    return 9223372036854775807
}
func genFn1(p0 bool) int {
    var v5 int16 = 32767
    switch 4294967295 {
    case -580:
        var v6 uint32 = 995
        _ = v6
    default:
        var v7 float32 = (((0.0 / ((3.14159 - 3.14159) - 3.14159)) - 1.5) / 1.5)
        _ = v7
    }
    fmt.Println(p0)
    switch "" {
    case "foo bar":
        var v8 uint = 0
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 18; fi9++ {
            if p0 {
                switch "a" {
                case "a":
                    if (fi9 > -1) {
                        if p0 {
                            var v10 uint64 = 989
                            _ = v10
                        }
                    } else {
                        switch 128 {
                        case 0:
                            v11 := []int{(^(fi9)), fi9, 184, -2147483648}
                            _ = v11
                        case 290:
                            switch 32768 {
                            case 0:
                                var v12 uint64 = 42
                                _ = v12
                            case 935:
                                p0 = p0
                            default:
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    v14 := GenS1{F0: v5, F1: (("hello" + ("foo bar" + "a")) + ""), F2: 2}
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        default:
                            var v15 int8 = 42
                            _ = v15
                        }
                        v16 := []int{}
                        _ = v16
                    }
                case "\n\t":
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v18 := []int{fi9, -416, fi17, (fi17 << ((241) & 63))}
                        v18 = append(v18, (fi9 << ((748) & 63)))
                        _ = fi17
                        _ = v18
                    }
                default:
                    p0 = false
                }
                v19 := []int{32767, (fi9 * fi9), fi9, fi9}
                _ = v19
            } else {
                p0 = p0
                var v20 int32 = 42
                _ = v20
            }
            _ = fi9
        }
    case "hello":
        switch "" {
        case "foo bar":
            fmt.Println(v5)
        case "hello":
            fmt.Println(v5, v5)
        case "a":
            var v21 float32 = 0.1
            _ = v21
        default:
            var v22 float64 = 3.14159
            _ = v22
        }
    default:
        if false {
            var v23 int8 = int8(v5)
            v5 -= 42
            if false {
                v24 := GenS1{F0: (v5 + -173), F1: "a", F2: 127}
                _ = v24
            } else {
                var v25 int32 = -2147483648
                var v26 uint64 = 0
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            v27 := []int{672, 4294967295, 128}
            _ = v27
        }
    }
    var v28 uint8 = 130
    _ = p0
    _ = v28
    _ = v5
    return 0
}

func main() {
    v1 := []string{"hello", "a"}
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v29 uint16 = 334
    v30 := GenS2{F0: 42, F1: 1, F2: v29}
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS3{F0: 1000}
        v29 = (v29 | v29)
        delete(v3, "a")
        _ = fi31
        _ = v32
    }
    v33 := GenS4{F0: 0, F1: 42}
    v29 += (^(v29))
    fmt.Println(v29)
    v34 := GenS5{F0: "", F1: "hello"}
    switch "a" {
    case "hello":
        var v35 int8 = 109
        _ = v35
    case "a":
        v29 = (^(v29))
    case "":
        if true {
            if true {
                v36 := GenS1{F0: 1, F1: "", F2: 128}
                v29 = v29
                if (false && ((!(("foo bar" + (("\n\t" + ("\n\t" + ("hello" + "foo bar"))) + ("foo bar" + "\n\t"))) != "foo bar")) && false)) {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 float32 = 0.0
                        _ = fi37
                        _ = v38
                    }
                }
                _ = v36
            }
            v29 = v29
        } else {
            v29 -= 0
            v29 += v29
            switch "hello" {
            case "":
                if ((("" > "") || true) || true) {
                    v29 = v29
                } else {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v29)
                    case "a":
                        var v39 float64 = 1.5
                        _ = v39
                    }
                }
            default:
                v40 := GenS2{F0: 1, F1: 9223372036854775807, F2: (v29 & v29)}
                _ = v40
            }
        }
    }
    v29 += (v29 * v29)
    v41 := GenS5{F0: ((("" + "hello") + ("\n\t" + "a")) + "\n\t"), F1: ("" + ("\n\t" + "hello"))}
    var v42 int16 = 471
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v41
    _ = v42
}

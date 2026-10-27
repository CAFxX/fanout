package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint16; F2 int8 }
type GenS6 struct { F0 int64; F1 uint }
func genFn0(p0 int64, p1 int8) uint32 {
    p0 = (p0 + (p0 - (p0 / ((0) | 1))))
    v6 := []int{42, -883, 127}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!(("" + ("foo bar" + ("hello" + "\n\t"))) >= "\n\t")) {
            delete(v5, (fi7 / (((-(fi7))) | 1)))
            var v8 uint8 = 2
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint64 = 263
                v11 := GenS1{F0: v10, F1: ((-128 == -86) && true)}
                v2[(fi9 << ((42) & 63))] = ("foo bar" + "")
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        } else {
            var v12 uint16 = 1000
            if true {
                if false {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(fi7, fi7)
                        _ = fi13
                    }
                    v12 += 32767
                }
            }
            fmt.Println(v12, fi7)
            _ = v12
        }
        v14 := GenS2{F0: "foo bar", F1: 1}
        _ = fi7
        _ = v14
    }
    if true {
        var v15 int = -2147483648
        var v16 bool = (-1 > 1)
        _ = v15
        _ = v16
    } else {
        var v17 string = "hello"
        _ = v17
    }
    var v18 int8 = -10
    v19 := []int{}
    for fi20 := 0; fi20 < 14; fi20++ {
        _ = len(v2)
        v18 -= (v18 + v18)
        _ = fi20
    }
    var v21 uint = 2
    v22 := GenS2{F0: "a", F1: 751}
    switch "foo bar" {
    case "hello":
        var v23 bool = false
        _ = v23
    case "\n\t":
        v24 := GenS2{F0: (("a" + (("a" + ((("hello" + ("\n\t" + "hello")) + "hello") + (("foo bar" + (("hello" + "\n\t") + "hello")) + ("foo bar" + "foo bar")))) + "foo bar")) + "hello"), F1: 0}
        _ = v24
    default:
        v25 := GenS3{F0: (v18 | v18)}
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            v4["\n\t"] = fi27
            _ = fi27
        }
        _ = fi26
    }
    var v28 float32 = ((0.1 * (0.0 * (float32(1e10) + -2.25))) - 0.0)
    switch 1 {
    case -909:
        var v29 string = ""
        _ = v29
    case 459:
        var v30 uint = 2
        _ = v30
    default:
        switch -1 {
        case 392:
            for fi31 := 0; fi31 < 11; fi31++ {
                v28 = v28
                v32 := GenS4{F0: 1, F1: 127}
                if true {
                    var v33 uint32 = 437
                    _ = v33
                } else {
                    fmt.Println(v18, v18)
                    v21 = 18446744073709551615
                    var v34 bool = ((true && false) || ("" < "hello"))
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
        case -622:
            fmt.Println(v28)
        case 127:
            for fi35 := 0; fi35 < 4; fi35++ {
                v1[127] = ((v28 <= (v28 + 0.1)) || false)
                v36 := GenS3{F0: -128}
                _ = fi35
                _ = v36
            }
        default:
            v28 = (v28 * v28)
        }
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        if ((((("foo bar" + "foo bar") + "foo bar") + "\n\t") + "foo bar") < "hello") {
            v38 := GenS5{F0: 837, F1: 0, F2: 42}
            fmt.Println(v28, fi37)
            switch fi37 {
            case -2147483648:
                for fi39 := 0; fi39 < 20; fi39++ {
                    fmt.Println(v21, v28)
                    _ = fi39
                }
            default:
                v28 += v28
            }
            _ = v38
        } else {
            switch ("" + "hello") {
            case "a":
                v40 := GenS6{F0: 552, F1: (^(v21))}
                _ = v40
            default:
                v21 -= 255
            }
            v41 := GenS4{F0: 856, F1: 2}
            _ = v41
        }
        _ = fi37
    }
    var v42 uint64 = 42
    switch 174 {
    case -958:
        for fi43 := 0; fi43 < 19; fi43++ {
            v21 += v21
            _ = fi43
        }
    case -34:
        v28 += v28
    }
    var v44 int32 = -1
    var v45 float64 = -2.25
    if len(v19) > 0 {
        v19[0] = int(v44)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v5
}

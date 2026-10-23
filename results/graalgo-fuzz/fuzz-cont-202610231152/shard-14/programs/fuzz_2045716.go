package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint8, p1 int16, p2 uint) int64 {
    fmt.Println(p2, p0, p0)
    switch "\n\t" {
    case "\n\t":
        switch "a" {
        case "hello":
            var v6 string = ("" + ("hello" + "foo bar"))
            _ = v6
        }
    case "hello":
        var v7 float32 = -2.25
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint64 {
    var v8 uint16 = 1
    if (("hello" + "hello") <= "") {
        var v9 uint32 = 276
        v10 := []int{-124, -70, -465}
        _ = v10
        _ = v9
    }
    var v11 uint8 = 0
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint64 = 0
        switch fi12 {
        case 416:
            v14 := []int{(fi12 << ((65535) & 63)), 2147483647, fi12, fi12}
            _ = v14
        case -493:
            var v15 bool = (uint64(v11) < (v13 + v13))
            _ = v15
        case 38:
            var v16 uint16 = (v8 % (((^(v8))) | 1))
            _ = v16
        default:
            var v17 int64 = genFn0(152, 1, 65535)
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    v18 := GenS1{F0: (v8 % ((255) | 1)), F1: 198}
    _ = v11
    _ = v18
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    var v3 uint32 = 4294967295
    v4 := make(map[string]bool)
    var v5 int16 = 568
    v3 = v3
    if false {
        if false {
            if (((943 <= 42) && (v5 < 2)) || (!(((!false) || true) && false))) {
                fmt.Println(v3, v5)
            }
            v19 := GenS1{F0: 313, F1: (v2 | (v2 | 820))}
            _ = v19
        } else {
            v3 = v3
        }
    } else {
        v2 -= (^(v2))
        if ((0 != 32767) || false) {
            if ("hello" > ("\n\t" + ("foo bar" + (("\n\t" + ((("\n\t" + "a") + ((("hello" + "foo bar") + ("a" + ("foo bar" + (((("a" + "") + "hello") + "hello") + "")))) + "\n\t")) + "")) + "")))) {
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint = (v2 | (v2 ^ 0))
                    switch "foo bar" {
                    case "\n\t":
                        v2 -= v21
                    case "a":
                        var v22 float32 = 0.1
                        _ = v22
                    case "hello":
                        if false {
                            fmt.Println(v21)
                            var v23 float64 = float64(1.5)
                            v3 = 4294967295
                            _ = v23
                        } else {
                            v24 := GenS2{F0: (v5 | v5)}
                            switch "a" {
                            case "foo bar":
                                switch fi20 {
                                case -67:
                                    v25 := GenS3{F0: genFn0(245, -1, 875)}
                                    _ = v25
                                }
                            case "\n\t":
                                var v26 uint16 = 324
                                _ = v26
                            case "hello":
                                var v27 float64 = -2.25
                                _ = v27
                            default:
                                var v28 bool = true
                                _ = v28
                            }
                            _ = v24
                        }
                    default:
                        v29 := []int{127, fi20}
                        _ = v29
                    }
                    _ = fi20
                    _ = v21
                }
                v2 += (^(v2))
            } else {
                var v30 float32 = (0.1 + float32((((0.1 + -2.25) + 3.14159) - 0.0)))
                _ = v30
            }
        } else {
            var v31 int64 = 9223372036854775807
            _ = v31
        }
        fmt.Println(v3, v3, v3)
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int = 32767
        if ((!((!false) && (!(true && true)))) && true) {
            var v34 uint32 = v3
            v3 += 2
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    var v35 uint16 = 1
    v36 := []int{int(v35), 79, int(v2), 844}
    for fi37 := 0; fi37 < 8; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 uint32 = (^(v3))
            _ = len(v36)
            v3 = v39
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 19; fi40++ {
            v35 = v35
            v2 = (^(v2))
            fmt.Println(v2, v35)
            _ = fi40
        }
        _ = len(v1)
        _ = fi37
    }
    var v41 float32 = 0.0
    var v42 int32 = 2147483647
    var v43 int32 = -1
    var v44 int8 = -1
    var v45 float64 = 0.1
    switch ("" + "a") {
    case "hello":
        switch -1 {
        case 32768:
            v46 := GenS1{F0: (v35 & 0), F1: (^(v2))}
            _ = v46
        case 42:
            for fi47 := 0; fi47 < 7; fi47++ {
                v2 = 2
                _ = fi47
            }
        default:
            v43 = 2147483647
        }
    }
    var v48 float64 = float64(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v5
}

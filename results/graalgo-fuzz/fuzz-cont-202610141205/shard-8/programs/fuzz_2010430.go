package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 int16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 uint32, p1 int8, p2 int) int16 {
    fmt.Println(p2)
    v6 := GenS1{F0: 236}
    v7 := []int{p2, p2, int(p1), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -462
}
func genFn1(p0 int, p1 uint8) int {
    var v8 float32 = 0.0
    switch ("hello" + ("hello" + "")) {
    case "":
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: 901}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case "foo bar":
        v12 := []int{p0}
        _ = v12
    default:
        v13 := []int{}
        _ = v13
    }
    v8 = v8
    switch (p0 << ((4294967295) & 63)) {
    case -2147483648:
        p0 -= (p0 + -398)
    case 826:
        if false {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint32 = 42
                _ = fi14
                _ = v15
            }
            v8 = 0.1
        }
    case -47:
        if false {
            v16 := GenS1{F0: 0}
            if (true && (false || (true && ((-(p0)) != (p0 << ((249) & 63)))))) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    if true {
                        fmt.Println(p0, p0, v8)
                        if (4294967295 == 4294967295) {
                            switch (fi17 >> ((415) & 63)) {
                            case 4294967295:
                                v18 := GenS2{F0: 32767, F1: 65535}
                                _ = v18
                            default:
                                p1 -= 2
                            }
                            var v19 int16 = 1
                            v20 := []int{}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint64 = 4294967295
                            v22 := []int{983, (p0 >> ((2) & 63)), fi17, (^(fi17))}
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v8 += 0.1
                                v24 := GenS3{F0: 2147483647, F1: 437, F2: genFn0(471, -1, (fi23 % (((-(fi23))) | 1)))}
                                v25 := GenS1{F0: -1}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                    if (1000 <= 42) {
                        v26 := []int{(fi17 >> ((475) & 63)), fi17}
                        p0 += 42
                        _ = v26
                    }
                    if true {
                        v27 := []int{p0, fi17}
                        p1 -= (p1 | 255)
                        var v28 string = "foo bar"
                        _ = v27
                        _ = v28
                    }
                    _ = fi17
                }
                p1 = p1
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint8 = (p1 * 0)
                _ = fi29
                _ = v30
            }
            _ = v16
        }
    default:
        var v31 uint32 = 1
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v8
    return (p0 % ((p0) | 1))
}

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 float64 = 1.5
    var v4 float64 = (0.0 - (v3 - (v3 + v3)))
    var v5 int8 = -128
    var v32 float32 = (0.0 * 0.1)
    v33 := GenS1{F0: 0}
    if (1 == -243) {
        v3 += v4
        for fi34 := 0; fi34 < 19; fi34++ {
            v32 = v32
            fmt.Println(v32)
            v35 := GenS2{F0: 361, F1: 521}
            _ = fi34
            _ = v35
        }
        var v36 float64 = v3
        _ = v36
    }
    v37 := GenS3{F0: 0, F1: 0, F2: 0}
    v5 -= (v5 * (v5 * (v5 ^ (v5 >> ((708) & 63)))))
    switch -176 {
    case 2147483647:
        v38 := GenS1{F0: 307}
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        if true {
            v1 = append(v1, fi39)
            switch "a" {
            case "\n\t":
                v40 := GenS4{F0: fi39, F1: 793}
                _ = v40
            }
            var v41 uint32 = 603
            _ = v41
        } else {
            v42 := GenS3{F0: 2147483647, F1: 0, F2: 571}
            var v43 float64 = v3
            var v44 float32 = v32
            _ = v42
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 11; fi45++ {
            v46 := GenS5{F0: 4294967295, F1: ((v4 <= v4) || (v5 != (v5 % ((v5) | 1)))), F2: "a"}
            v32 -= float32((0.0 * ((v3 / v3) / v4)))
            _ = fi45
            _ = v46
        }
        _ = fi39
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v4, v5)
    default:
        fmt.Println(v4, v4)
    }
    var v47 float32 = v32
    if (float64(v32) != (v3 * v3)) {
        v32 += v32
    }
    if ((("" + "\n\t") + "foo bar") != ("a" + "")) {
        delete(v2, "")
        for fi48 := 0; fi48 < 5; fi48++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(v4)
            case "hello":
                var v49 int16 = 32767
                _ = v49
            }
            for fi50 := 0; fi50 < 20; fi50++ {
                v51 := GenS4{F0: fi48, F1: 1}
                v5 += (v5 ^ (v5 | v5))
                v32 += float32(float64(float32((((v4 - float64(0.1)) * ((0.1 / (v3 / float64((v32 * (float32(1e10) - 0.0))))) / ((v4 * (v3 / (v4 + v3))) - (1.5 - -2.25)))) / (3.14159 / (0.1 * float64(-2.25)))))))
                _ = fi50
                _ = v51
            }
            _ = fi48
        }
    }
    var v52 int8 = (v5 & v5)
    if len(v1) > 0 {
        v1[0] = 681
    }
    v5 = -82
    v3 += ((v3 * v3) + ((v3 / ((v4 - v3) / (0.0 * v3))) / (v4 + v4)))
    var v53 float32 = v47
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v47
    _ = v5
    _ = v52
    _ = v53
}

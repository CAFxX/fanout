package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 string; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint, p1 int8, p2 uint32) uint16 {
    var v7 float64 = (float64(-2.25) + 1.5)
    v8 := GenS1{F0: 1, F1: p0, F2: "foo bar"}
    if false {
        var v9 int16 = 42
        fmt.Println(v9, p1)
        fmt.Println(v9, p2, p0)
        _ = v9
    } else {
        var v10 uint64 = 42
        var v11 uint32 = p2
        for fi12 := 0; fi12 < 19; fi12++ {
            if (!(false && ((-2.25 < ((3.14159 / 1.5) - 0.0)) && true))) {
                var v13 uint = p0
                v10 = v10
                if ((false && true) && true) {
                    switch fi12 {
                    case 1:
                        v14 := []int{fi12, -833, fi12}
                        _ = v14
                    case 575:
                        switch "" {
                        case "\n\t":
                            v15 := []int{fi12, -2147483648}
                            _ = v15
                        default:
                            var v16 uint8 = 1
                            _ = v16
                        }
                    case 638:
                        for fi17 := 0; fi17 < 4; fi17++ {
                            switch "hello" {
                            case "":
                                v13 -= v13
                            case "foo bar":
                                v18 := []int{fi17, (fi17 + 128), (fi12 + fi12)}
                                _ = v18
                            case "\n\t":
                                p2 = p2
                                v19 := []int{(fi17 % ((fi12) | 1))}
                                _ = v19
                            }
                            v20 := GenS1{F0: (v10 << (((v13 >> ((18446744073709551615) & 63))) & 63)), F1: (p0 * (v13 - p0)), F2: (("\n\t" + "a") + ((("\n\t" + "hello") + "") + ("" + "\n\t")))}
                            v21 := []int{-2147483648, (fi17 >> ((589) & 63)), (-(fi12))}
                            var v22 uint16 = 2
                            _ = fi17
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    default:
                        v23 := []int{}
                        _ = v23
                    }
                } else {
                    var v24 int16 = int16(p1)
                    if (true || true) {
                        var v25 uint16 = 2
                        if ((p2 ^ v11) >= p2) {
                            switch (fi12 * fi12) {
                            case -1:
                                v26 := GenS2{F0: v25}
                                _ = v26
                            case 1:
                                if (true && true) {
                                    fmt.Println(fi12)
                                    var v27 uint32 = (v11 & (^(p2)))
                                    _ = v27
                                }
                            }
                            switch "" {
                            case "":
                                v10 = 0
                            default:
                                v28 := GenS2{F0: 741}
                                _ = v28
                            }
                            v7 += 0.0
                        }
                        fmt.Println(v7, v13, v25)
                        _ = v25
                    }
                    v10 = (v10 % (((^(v10))) | 1))
                    _ = v24
                }
                _ = v13
            } else {
                v29 := []int{-9223372036854775808}
                v30 := GenS1{F0: 0, F1: (p0 ^ 42), F2: "a"}
                _ = v29
                _ = v30
            }
            v31 := GenS3{F0: 2, F1: "", F2: 545}
            p2 = (v11 & p2)
            _ = fi12
            _ = v31
        }
        _ = v10
        _ = v11
    }
    var v32 float64 = (v7 + float64(-2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v7
    _ = v8
    return 65535
}
func genFn1() uint16 {
    var v33 int = -590
    v34 := []int{735, v33, (v33 & (-(v33))), (-(v33))}
    _ = v33
    _ = v34
    return genFn0(0, 100, 1)
}

func main() {
    var v1 int8 = 101
    var v2 uint = 537
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 uint8 = 109
    var v6 float64 = ((-2.25 * 0.0) / 1.0)
    v2 -= 18446744073709551615
    v35 := GenS2{F0: 1}
    v1 += (v1 - -128)
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 float64 = -2.25
        for fi38 := 0; fi38 < 16; fi38++ {
            v6 = float64(-2.25)
            var v39 int = 0
            v2 += (v2 | (v2 << ((v2) & 63)))
            _ = fi38
            _ = v39
        }
        v6 = v6
        _ = fi36
        _ = v37
    }
    if (("\n\t" + "a") >= "hello") {
        var v40 float32 = 3.14159
        v2 = (v2 | v2)
        _ = v40
    }
    var v41 int64 = 9223372036854775807
    v42 := GenS4{F0: 2147483647, F1: "", F2: (v5 % ((v5) | 1))}
    v43 := GenS4{F0: 1, F1: "foo bar", F2: v5}
    v44 := []int{}
    var v45 float64 = v6
    v46 := GenS5{F0: 2147483647, F1: 2}
    var v47 uint16 = 797
    for fi48 := 0; fi48 < 1; fi48++ {
        v49 := GenS6{F0: (-(v1))}
        _ = len(v44)
        var v50 int = (^(fi48))
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println(v47)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

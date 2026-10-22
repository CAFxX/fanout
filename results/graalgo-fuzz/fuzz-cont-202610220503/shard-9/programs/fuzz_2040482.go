package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint32 }
func genFn0(p0 int8) uint64 {
    p0 -= p0
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 2147483647}
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint32 = 632
            v8 := []int{127, fi4, (fi6 & fi6)}
            p0 -= (p0 - (p0 << ((4294967295) & 63)))
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    _ = p0
    return 4294967295
}
func genFn1() uint32 {
    var v9 int = 2147483647
    var v10 uint = 18446744073709551615
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v9, fi11)
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := []int{42, (fi12 % ((fi11) | 1)), -271}
            if true {
                switch 128 {
                case 1:
                    v13 = append(v13, 127)
                }
                v14 := GenS2{F0: -1}
                var v15 uint16 = 89
                v13 = append(v13, (v9 * fi11))
                _ = v14
                _ = v15
            } else {
                v16 := GenS3{F0: -760, F1: (v9 | (fi12 % (((fi12 << (((v10 & v10)) & 63))) | 1))), F2: v9}
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        v17 := []int{239, (-(fi11))}
        _ = fi11
        _ = v17
    }
    v9 -= v9
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []int{127, 32768, 1, 513}
    var v2 string = ("hello" + "\n\t")
    var v3 int = -2147483648
    fmt.Println(v2, v3)
    var v18 uint16 = 529
    var v19 int16 = 401
    if false {
        var v20 uint = 1
        _ = v20
    } else {
        v3 = v3
        var v21 int16 = v19
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := GenS4{F0: 65535}
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 uint64 = 4294967295
    v2 = (v2 + ((((((v2 + v2) + "foo bar") + ("\n\t" + v2)) + v2) + v2) + "a"))
    if (1 <= -1) {
        var v25 float32 = -2.25
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = v25
    }
    var v26 float64 = float64((-2.25 - 0.1))
    switch v2 {
    case "\n\t":
        v27 := GenS5{F0: (120 >= 58)}
        _ = v27
    }
    switch v2 {
    case "\n\t":
        switch v2 {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            switch v2 {
            case "\n\t":
                var v28 int8 = 13
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 13; fi29++ {
                fmt.Println(v26, fi29, v26)
                switch fi29 {
                case 278:
                    v1 = append(v1, 4294967295)
                case 9223372036854775807:
                    fmt.Println(v26, v2, v24)
                case -460:
                    var v30 int32 = 2147483647
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v26 -= (v26 - (float64((0.0 - float32(-2.25))) - (3.14159 / (v26 + v26))))
                        v32 := GenS2{F0: 32767}
                        if ((false && (v18 > v18)) || false) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if ((v24 | genFn0(-56)) < v24) {
                                v3 = fi29
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                switch "a" {
                case "a":
                    v33 := GenS6{F0: 1}
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v35 := GenS5{F0: true}
                        var v36 int8 = -128
                        if true {
                            switch (fi29 + v3) {
                            case 878:
                                var v37 bool = ((0 != 0) && false)
                                _ = v37
                            default:
                                var v38 string = ""
                                _ = v38
                            }
                        } else {
                            if false {
                                var v39 float64 = v26
                                if (-1 <= 78) {
                                    v26 -= float64(1e10)
                                    fmt.Println(v36, v3, v26)
                                    v40 := GenS7{F0: fi29, F1: genFn1()}
                                    _ = v40
                                } else {
                                    fmt.Println(v39, v26)
                                    fmt.Println(v24, v36)
                                }
                                _ = v39
                            }
                            v41 := GenS1{F0: 217}
                            if false {
                                v3 = (^(fi29))
                                v42 := GenS6{F0: 9223372036854775807}
                                _ = v42
                            }
                            _ = v41
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi29
            }
        }
    case "hello":
        var v43 int16 = v19
        _ = v43
    default:
        var v44 int32 = 2147483647
        _ = v44
    }
    v24 = (v24 | genFn0(100))
    v45 := GenS6{F0: 135}
    v1 = append(v1, 128)
    v18 += 65535
    fmt.Println(v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v45
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 bool; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 bool; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32) uint32 {
    for fi5 := 0; fi5 < 13; fi5++ {
        p0 += 693
        p0 += (p0 | (p0 % ((p0) | 1)))
        _ = fi5
    }
    var v6 string = "foo bar"
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 int, p1 bool, p2 uint32) uint16 {
    p0 -= -1
    var v7 uint = 4294967295
    v7 -= (v7 | (v7 - (^(v7))))
    switch "foo bar" {
    case "":
        v8 := []int{-808, (p0 + p0), (-(p0))}
        _ = v8
    case "a":
        v9 := GenS1{F0: -32768}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            p2 -= p2
            v11 := []int{}
            p1 = p1
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS2{F0: p2}
        _ = v12
    }
    v13 := []int{int(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 387
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    v4 := []uint32{0, 0, (v3 - 0), v3}
    _ = len(v4)
    if (false && true) {
        v14 := GenS3{F0: 94, F1: 18446744073709551615, F2: -9223372036854775808}
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            delete(v2, fi15)
            var v17 uint32 = genFn0(42)
            _ = fi16
            _ = v17
        }
        v3 += v3
        _ = fi15
    }
    fmt.Println(v3, v3)
    v4 = append(v4, (v3 | genFn0(-184)))
    v18 := GenS4{F0: false, F1: -155, F2: -1}
    var v19 float32 = 0.0
    v19 += v19
    var v20 string = ""
    for fi21 := 0; fi21 < 1; fi21++ {
        v4 = append(v4, v3)
        switch fi21 {
        case -1:
            if ((0 == 42) || (127 != 255)) {
                fmt.Println(fi21, v20, fi21)
                if false {
                    v22 := GenS5{F0: 731, F1: true, F2: (v3 - genFn0(-530))}
                    var v23 string = "a"
                    var v24 float64 = (0.1 + 1.5)
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint = 255
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 int32 = 42
                        fmt.Println(v27)
                        switch v20 {
                        case "":
                            v3 = (v3 * 2)
                        case "hello":
                            if false {
                                var v28 string = ("a" + "\n\t")
                                var v29 uint8 = 128
                                var v30 float32 = (float32((1e10 * 0.0)) / (float32((3.14159 + float64(float32(-2.25)))) * 0.1))
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                if ((^(fi26)) != fi26) {
                                    fmt.Println(fi26)
                                    var v31 uint = (v25 << ((v25) & 63))
                                    v32 := GenS6{F0: (v1 + v1), F1: 2147483648}
                                    v25 = (v31 - 18446744073709551615)
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 int16 = 0
                                    v1 -= v1
                                    var v34 float32 = v19
                                    _ = v33
                                    _ = v34
                                }
                                var v35 int = fi26
                                _ = len(v2)
                                _ = v35
                            }
                        case "a":
                            fmt.Println(v25, v20)
                        }
                        _ = fi26
                        _ = v27
                    }
                    v19 += v19
                    _ = v25
                }
            }
        case 0:
            var v36 uint32 = v3
            _ = v36
        case 2147483647:
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
        _ = fi21
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        fmt.Println(fi37)
        var v38 int64 = 193
        for fi39 := 0; fi39 < 2; fi39++ {
            var v40 uint16 = 42
            switch -344 {
            case -2147483648:
                if (912 <= (fi39 / (((fi37 >> ((421) & 63))) | 1))) {
                    v41 := GenS5{F0: -862, F1: (!true), F2: genFn0(-390)}
                    var v42 int32 = 2147483647
                    _ = v41
                    _ = v42
                } else {
                    if (false || true) {
                        var v43 string = v20
                        v1 = (v1 << ((42) & 63))
                        v44 := GenS1{F0: -897}
                        if (((v40 * v40) >= v40) || false) {
                            v40 += (v40 | 0)
                        } else {
                            if ((true || false) || true) {
                                fmt.Println(v20, v1)
                                var v45 uint8 = 2
                                _ = v45
                            }
                            var v46 int = int(v40)
                            v47 := GenS6{F0: 42, F1: (v3 * v3)}
                            var v48 float64 = 1.5
                            v49 := GenS1{F0: -32768}
                            var v50 uint16 = (v40 + (v40 ^ v40))
                            _ = v46
                            _ = v47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        _ = v43
                        _ = v44
                    }
                    v51 := GenS6{F0: (v1 + v1), F1: genFn0(1)}
                    v52 := GenS6{F0: (v1 / ((48) | 1)), F1: v3}
                    _ = v51
                    _ = v52
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    _ = len(v2)
    var v53 uint8 = 128
    for fi54 := 0; fi54 < 14; fi54++ {
        switch fi54 {
        case 32767:
            v19 = v19
        }
        _ = fi54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v53
}

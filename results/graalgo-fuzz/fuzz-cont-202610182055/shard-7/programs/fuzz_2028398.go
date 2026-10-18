package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8, p1 int8) uint32 {
    fmt.Println(p0, p0)
    v4 := []int{}
    _ = p0
    _ = p1
    _ = v4
    return 2
}
func genFn1(p0 int32, p1 bool, p2 int8) int16 {
    for fi5 := 0; fi5 < 19; fi5++ {
        if p1 {
            fmt.Println(p1)
        } else {
            var v6 uint64 = 2
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 = p0
        v8 := GenS1{F0: fi7, F1: p2, F2: p0}
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = (v1 | v1)
    var v3 bool = (65535 == 0)
    for fi9 := 0; fi9 < 4; fi9++ {
        v2 = v2
        _ = fi9
    }
    fmt.Println(v1)
    v1 = (-(v1))
    fmt.Println(v2, v1, v1)
    var v10 uint16 = 1
    v1 += (v1 - v1)
    v1 = (v2 << ((1) & 63))
    switch -2147483648 {
    case 911:
        var v11 string = "a"
        _ = v11
    case -794:
        var v12 int = -1
        _ = v12
    default:
        v13 := GenS1{F0: -519, F1: -1, F2: (v1 >> ((4294967295) & 63))}
        _ = v13
    }
    v14 := GenS2{F0: -810, F1: (!((("a" >= ((("\n\t" + "a") + ("" + "hello")) + ("a" + (("\n\t" + ("hello" + ("a" + "a"))) + "")))) || v3) && (!false))), F2: 2147483648}
    for fi15 := 0; fi15 < 19; fi15++ {
        if false {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int16 = 32767
                _ = fi16
                _ = v17
            }
        } else {
            if (-519 == -739) {
                v18 := GenS3{F0: 513, F1: 42, F2: 477}
                v19 := []int{(^(fi15)), fi15}
                _ = v18
                _ = v19
            } else {
                v20 := GenS4{F0: 0, F1: true}
                var v21 float32 = 3.14159
                _ = v20
                _ = v21
            }
            v10 -= v10
            var v22 uint16 = v10
            _ = v22
        }
        v2 = 398
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v23 := []int{(fi15 ^ (fi15 * fi15)), (fi15 & fi15)}
            _ = v23
        case "":
            if false {
                if true {
                    if true {
                        var v24 uint16 = 42
                        switch int(v10) {
                        case -284:
                            var v25 uint16 = (^(v24))
                            _ = v25
                        case -9223372036854775808:
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v1 = v1
                                _ = fi26
                            }
                        case 4294967295:
                            if true {
                                v2 -= v1
                            }
                        default:
                            v24 -= v10
                        }
                        _ = v24
                    } else {
                        v27 := GenS3{F0: 710, F1: 1, F2: v10}
                        _ = v27
                    }
                } else {
                    var v28 uint32 = 4294967295
                    var v29 int32 = v2
                    var v30 int32 = (v1 >> ((0) & 63))
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if v3 {
                    v31 := GenS5{F0: 991, F1: (v10 + (v10 >> ((42) & 63)))}
                    _ = v31
                }
                v32 := GenS3{F0: -9223372036854775808, F1: 1, F2: v10}
                _ = v32
            } else {
                var v33 bool = true
                _ = v33
            }
        }
        _ = fi15
    }
    var v34 uint64 = 42
    var v35 uint16 = v10
    v36 := []int{770, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}

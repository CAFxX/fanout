package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 int64, p1 bool, p2 uint16) string {
    v6 := GenS1{F0: 2, F1: 42, F2: 158}
    v7 := []int{-666, 1}
    p1 = p1
    var v8 int = -615
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 += fi9
        var v10 int64 = 347
        if (-97 >= -2147483648) {
            v11 := GenS1{F0: 255, F1: 4294967295, F2: uint8(p2)}
            if p1 {
                v12 := GenS1{F0: 42, F1: 18446744073709551615, F2: uint8(p2)}
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 bool = (p1 || (!false))
                    var v15 uint32 = 1
                    p1 = (859 != 524)
                    v15 = v15
                    var v16 uint32 = (v15 | (v15 & v15))
                    var v17 uint32 = 995
                    var v18 string = (("" + "foo bar") + "")
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                v19 := GenS1{F0: 128, F1: 1, F2: 2}
                _ = v12
                _ = v19
            }
            if p1 {
                v20 := GenS1{F0: 0, F1: 0, F2: 2}
                _ = v20
            }
            _ = v11
        } else {
            p1 = p1
            var v21 int64 = 971
            fmt.Println(p0, p1)
            _ = v21
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (("a" + "a") + "\n\t")
}
func genFn1(p0 int32, p1 bool) uint64 {
    if (("" + "hello") < "\n\t") {
        v22 := []int{9223372036854775807, 0, -659}
        var v23 int = 0
        v23 = v23
        _ = v22
        _ = v23
    }
    var v24 uint8 = 2
    var v25 float32 = (3.14159 / -2.25)
    v26 := []int{9223372036854775807, 32767, -260}
    v25 = 1e10
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 114
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = -608
    v3 := make(map[uint32]bool)
    var v4 float64 = (-2.25 + 1e10)
    v5 := make(map[string]uint32)
    if false {
        var v27 float64 = v4
        _ = v27
    } else {
        v28 := GenS2{F0: 32767}
        v29 := []int{32767, 4294967295, -940, -1}
        switch 128 {
        case -203:
            var v30 int8 = 127
            _ = v30
        default:
            for fi31 := 0; fi31 < 16; fi31++ {
                fmt.Println(v1)
                if false {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        var v33 bool = true
                        _ = fi32
                        _ = v33
                    }
                } else {
                    var v34 string = "a"
                    _ = v34
                }
                var v35 int8 = -68
                _ = fi31
                _ = v35
            }
        }
        _ = v28
        _ = v29
    }
    if (v4 == v4) {
        v4 = (v4 * ((-2.25 / v4) + v4))
        if (0 <= 42) {
            var v36 string = "\n\t"
            v37 := []int{4294967295, 371, 638, 32768}
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 int64 = v2
                v40 := GenS2{F0: 862}
                if true {
                    var v41 uint32 = 42
                    for fi42 := 0; fi42 < 17; fi42++ {
                        var v43 uint = 475
                        v44 := GenS3{F0: v39, F1: 255, F2: 100}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    v41 = (v41 * v41)
                    _ = v41
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = v36
            _ = v37
        }
        v3[655] = ((!true) && false)
    }
    var v45 uint32 = 442
    v1 -= (v1 * (^(v1)))
    v46 := []int{4294967295, 32768}
    delete(v5, ("" + "hello"))
    var v47 int8 = 42
    for fi48 := 0; fi48 < 16; fi48++ {
        for fi49 := 0; fi49 < 1; fi49++ {
            var v50 uint16 = 2
            v1 += 190
            _ = fi49
            _ = v50
        }
        var v51 float64 = v4
        _ = fi48
        _ = v51
    }
    v52 := GenS2{F0: 2}
    var v53 uint64 = (v1 - v1)
    if false {
        var v54 uint64 = v1
        v4 -= v4
        _ = v54
    } else {
        v3[v45] = false
        var v55 int64 = v2
        v2 = v55
        _ = v55
    }
    var v56 float32 = (1e10 - -2.25)
    v53 = (v1 - (v53 / ((genFn1(-1, (!(v56 > v56)))) | 1)))
    v45 = v45
    var v57 int = 732
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 16; fi59++ {
            var v60 uint8 = 53
            _ = fi59
            _ = v60
        }
        for fi61 := 0; fi61 < 9; fi61++ {
            fmt.Println(v53)
            _ = fi61
        }
        _ = fi58
    }
    fmt.Println(v1, v57)
    if (((v57 * (-(v57))) != (v57 ^ 127)) && ((896 != 376) && true)) {
        v4 -= (v4 / v4)
        var v62 int16 = 119
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v57
}

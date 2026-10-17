package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0() int16 {
    v6 := []int{767}
    var v7 uint = 18446744073709551615
    _ = v6
    _ = v7
    return 1000
}
func genFn1(p0 uint64, p1 uint) uint8 {
    if true {
        if false {
            if false {
                v8 := GenS1{F0: true, F1: -128, F2: (("" + "hello") + "foo bar")}
                _ = v8
            }
            p1 = 1
        }
        switch -9223372036854775808 {
        case 128:
            for fi9 := 0; fi9 < 7; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    p1 = p1
                    p1 = p1
                    _ = fi10
                }
                _ = fi9
            }
        case 0:
            v11 := []int{32767, 128, 0, 127}
            _ = v11
        case -1:
            var v12 uint8 = 42
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS2{F0: 1}
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                p1 = p1
                fmt.Println(fi15, p0)
                p1 = (p1 - (p1 & (p1 << (((p1 % ((65535) | 1))) & 63))))
                v17 := GenS3{F0: 1}
                var v18 uint64 = (p0 * (p0 - 42))
                var v19 uint8 = 198
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            p0 -= (p0 * uint64(p1))
            var v20 int = fi15
            _ = fi15
            _ = v20
        }
    }
    var v21 float64 = 1.5
    _ = p0
    _ = p1
    _ = v21
    return 0
}

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 string = "\n\t"
    v3 := []int{32767, 2147483647}
    v4 := make(map[int]bool)
    v5 := []int{}
    if (2 >= 223) {
        switch 1 {
        case 127:
            if true {
                var v22 uint32 = 903
                v2 = v2
                _ = v22
            }
        case 32767:
            fmt.Println(v2)
        }
    }
    var v23 uint64 = 0
    if (!true) {
        v23 = 905
        if len(v1) > 0 {
            v1[0] = (v2 + ("hello" + "hello"))
        }
    }
    var v24 int8 = -8
    var v25 int16 = 902
    for fi26 := 0; fi26 < 9; fi26++ {
        if true {
            v25 += (-(v25))
        } else {
            fmt.Println(fi26, v23)
            v27 := GenS3{F0: -1}
            if (1 == 556) {
                var v28 uint16 = 712
                if ((v24 & 1) > v24) {
                    v2 = v2
                }
                _ = v28
            } else {
                v2 = v2
                for fi29 := 0; fi29 < 2; fi29++ {
                    v1 = append(v1, v2)
                    v30 := GenS4{F0: v25, F1: -361}
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 uint64 = (v23 >> ((255) & 63))
                    if false {
                        if false {
                            if (-2.25 >= (1.5 + 0.1)) {
                                fmt.Println(v32, fi26, v2)
                                fmt.Println(fi31)
                            }
                            v23 += 434
                            fmt.Println(v2, fi31, fi26)
                        } else {
                            v24 += 100
                            switch fi31 {
                            case -9223372036854775808:
                                v33 := v4[-9223372036854775808]
                                _ = v33
                            case 0:
                                fmt.Println(fi31)
                                var v34 int32 = 42
                                var v35 uint32 = 829
                                _ = v34
                                _ = v35
                            }
                            if ((v2 + (("" + "foo bar") + v2)) <= v2) {
                                _ = len(v3)
                            } else {
                                v5 = append(v5, fi31)
                            }
                            if true {
                                fmt.Println(v25, v25, fi31)
                                v23 += (v23 - v32)
                            }
                        }
                        var v36 uint32 = 1
                        _ = v36
                    }
                    var v37 uint8 = 1
                    _ = fi31
                    _ = v32
                    _ = v37
                }
            }
            _ = v27
        }
        _ = fi26
    }
    fmt.Println(v24, v2, v23)
    v23 = 42
    delete(v4, 32767)
    if (false || true) {
        delete(v4, 42)
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v25, v23)
        case "":
            var v38 bool = (!true)
            _ = v38
        case "\n\t":
            fmt.Println(v23, v2, v24)
        default:
            _ = len(v4)
        }
    }
    var v39 string = (v2 + "\n\t")
    v39 = v2
    fmt.Println(v39, v2)
    v40 := GenS5{F0: 42, F1: 915}
    var v41 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}

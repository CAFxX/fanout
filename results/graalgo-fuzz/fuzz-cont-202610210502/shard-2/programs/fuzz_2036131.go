package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 int, p1 uint16) uint {
    if (p1 < 1) {
        v6 := []int{}
        var v7 int = p0
        v6 = append(v6, (p0 >> ((810) & 63)))
        _ = v6
        _ = v7
    } else {
        var v8 int32 = 1
        var v9 string = "\n\t"
        fmt.Println(v8)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        p1 += p1
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := []int{p0, fi11, 128}
            v13 := GenS1{F0: 42, F1: 2}
            p0 = (fi10 | fi11)
            fmt.Println(fi11, p1, p0)
            var v14 int32 = 0
            p1 = (p1 >> ((42) & 63))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        p1 += 344
        _ = fi10
    }
    if (211 <= 0) {
        if (0 > 65535) {
            v15 := GenS1{F0: 127, F1: 2}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS1{F0: 42, F1: 1}
                var v18 float32 = (float32((1.5 * (0.0 + 1e10))) + (-2.25 / 3.14159))
                var v19 int = (p0 & p0)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 int = 1
            p0 += (p0 ^ (p0 / ((v20) | 1)))
            _ = v20
        }
        if false {
            switch p0 {
            case 569:
                v21 := []int{}
                _ = v21
            case -37:
                var v22 uint = 0
                _ = v22
            default:
                fmt.Println(p1, p0)
            }
            fmt.Println(p1, p0)
        }
    } else {
        if (("hello" + "") == (("" + ("foo bar" + "\n\t")) + (("a" + "a") + "\n\t"))) {
            for fi23 := 0; fi23 < 20; fi23++ {
                if ((1 < -1) && (!((fi23 >= fi23) && ("\n\t" >= "hello")))) {
                    v24 := []int{-563}
                    if false {
                        if true {
                            _ = len(v24)
                        }
                        var v25 uint64 = 321
                        for fi26 := 0; fi26 < 1; fi26++ {
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            v27 := GenS1{F0: 128, F1: 1}
                            v25 = (v25 - 2)
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    } else {
                        var v28 float64 = 1.5
                        if false {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                v28 += ((0.1 + v28) * v28)
                                _ = fi29
                            }
                        }
                        _ = v28
                    }
                    _ = v24
                } else {
                    if (1e10 >= 1.5) {
                        switch p0 {
                        case -1:
                            p1 = p1
                        }
                        p1 = p1
                    }
                    switch 658 {
                    case 9223372036854775807:
                        switch (p0 >> ((531) & 63)) {
                        case 975:
                            v30 := GenS1{F0: 42, F1: 292}
                            _ = v30
                        case -789:
                            v31 := []int{(fi23 - p0), (-(p0)), (^(fi23))}
                            _ = v31
                        case 376:
                            p0 = p0
                        default:
                            v32 := []int{(fi23 | p0), (p0 ^ (fi23 / ((80) | 1)))}
                            _ = v32
                        }
                    case 151:
                        v33 := []int{(^(p0)), (fi23 << ((42) & 63))}
                        _ = v33
                    }
                    if (279 > 2147483647) {
                        v34 := GenS2{F0: 2}
                        v35 := GenS2{F0: 65535}
                        p0 -= fi23
                        _ = v34
                        _ = v35
                    } else {
                        p1 += 0
                    }
                }
                _ = fi23
            }
        }
        switch 377 {
        case 1:
            v36 := GenS1{F0: 128, F1: 2}
            _ = v36
        case 4294967295:
            p0 += 128
        case 42:
            p1 -= (p1 + p1)
        default:
            p0 = 9223372036854775807
        }
    }
    v37 := []int{}
    v38 := GenS1{F0: 229, F1: 42}
    _ = p0
    _ = p1
    _ = v37
    _ = v38
    return 0
}
func genFn1(p0 string, p1 int, p2 int32) int {
    p0 = p0
    if (p2 > -1) {
        var v39 bool = false
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 int16 = 184
            v42 := GenS3{F0: 166, F1: 460}
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = v39
    }
    var v43 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return (-(p1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    v4 := []string{("a" + "\n\t"), "foo bar", "foo bar"}
    var v5 uint = 935
    fmt.Println(v5, v5)
    v3 = 4294967295
    if false {
        v44 := GenS4{F0: 33, F1: false}
        var v45 float32 = ((3.14159 + 1e10) / 1.0)
        v45 = v45
        _ = v44
        _ = v45
    }
    if false {
        var v46 uint = (^(v5))
        v3 = v3
        _ = v46
    }
    var v47 int8 = 0
    var v48 float32 = (3.14159 - (-2.25 / 0.1))
    v48 += v48
    var v49 int = -326
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}

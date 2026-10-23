package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint32) int64 {
    fmt.Println(p0)
    var v7 int = 32767
    var v8 uint64 = 42
    if (127 > 100) {
        if true {
            var v9 int64 = 290
            v8 = v8
            _ = v9
        } else {
            switch (v7 / ((v7) | 1)) {
            case 151:
                v10 := []int{(v7 | (^(v7))), v7}
                _ = v10
            case 2147483647:
                v11 := GenS1{F0: "a"}
                _ = v11
            case 32768:
                v12 := []int{193, -9223372036854775808, (v7 * v7), (v7 >> ((680) & 63))}
                _ = v12
            }
            v13 := GenS2{F0: 2, F1: 1, F2: (v8 + (v8 % (((^(v8))) | 1)))}
            _ = v13
        }
        fmt.Println(v8)
        fmt.Println(v7)
    } else {
        fmt.Println(v8, p0, p0)
        v14 := []int{v7, v7, v7}
        var v15 int64 = 1
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int, p1 int, p2 uint) int16 {
    var v16 uint16 = 191
    fmt.Println(p0, p2)
    var v17 uint = 517
    if (!(100 != 127)) {
        v17 += 0
        p2 -= (p2 * (v17 - (v17 * p2)))
        switch (-(p1)) {
        case 128:
            p1 -= (p1 - p0)
        case -2147483648:
            v18 := []int{p0, (p1 & p0)}
            _ = v18
        default:
            if true {
                p1 = 9223372036854775807
            } else {
                v19 := []int{42}
                _ = v19
            }
        }
    } else {
        v17 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 63
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    v5 := []string{(("hello" + ("" + v2)) + v2)}
    v6 := []string{v2, v2, ""}
    v3 = 1
    v5 = append(v5, (v2 + v2))
    var v20 int8 = 100
    switch (((("\n\t" + (v2 + v2)) + "foo bar") + "\n\t") + ("" + (v2 + v2))) {
    case "hello":
        var v21 int8 = (v20 & v20)
        _ = v21
    case "":
        if (0 >= 0) {
            var v22 int64 = 102
            switch "hello" {
            case "\n\t":
                fmt.Println(v2, v20)
            case "":
                switch 32768 {
                case 42:
                    fmt.Println(v20)
                case -786:
                    _ = len(v6)
                default:
                    if len(v5) > 0 {
                        v5[0] = v2
                    }
                }
            default:
                v2 = ("hello" + (("\n\t" + "\n\t") + v2))
            }
            for fi23 := 0; fi23 < 16; fi23++ {
                v6 = append(v6, v2)
                var v24 uint8 = 255
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    default:
        var v25 uint8 = 128
        _ = v25
    }
    if (!((-(v3)) >= v3)) {
        for fi26 := 0; fi26 < 3; fi26++ {
            if ((167 == 127) && (true || ((v20 << ((266) & 63)) <= -58))) {
                v1[2147483648] = ((true || true) && true)
                if false {
                    var v27 int = fi26
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v20 = v20
                        if (false || (21 <= 2)) {
                            if true {
                                if true {
                                    v29 := GenS2{F0: 2, F1: 1, F2: 2}
                                    v30 := GenS2{F0: 817, F1: genFn0(2), F2: 1}
                                    v31 := GenS3{F0: 1}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                v27 -= (fi26 << ((546) & 63))
                            }
                        } else {
                            v20 -= v20
                            _ = len(v1)
                            var v32 float32 = 1.5
                            _ = v32
                        }
                        var v33 int64 = 527
                        _ = fi28
                        _ = v33
                    }
                    _ = v27
                }
                v34 := GenS4{F0: 0}
                _ = v34
            }
            v35 := GenS3{F0: 284}
            _ = fi26
            _ = v35
        }
        if (false && false) {
            var v36 int64 = (v3 - (^(v3)))
            _ = v36
        }
    }
    if ((!(0 != 4294967295)) && (!false)) {
        fmt.Println(v3, v20)
    }
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v37 int8 = (-(v20))
        var v38 uint32 = 2147483647
        _ = v37
        _ = v38
    }
    fmt.Println(v2, v2, v3)
    switch -1 {
    case 32767:
        v39 := v4[4294967295]
        _ = v39
    }
    if (true || false) {
        v20 = (v20 * (v20 % ((v20) | 1)))
        _ = len(v1)
    } else {
        for fi40 := 0; fi40 < 7; fi40++ {
            v3 = v3
            v3 += genFn0(2)
            v3 = (v3 * v3)
            _ = fi40
        }
        var v41 int8 = 42
        _ = v41
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        v43 := GenS4{F0: 2}
        _ = fi42
        _ = v43
    }
    var v44 int8 = v20
    var v45 int16 = 0
    v46 := GenS5{F0: false, F1: 2}
    v20 = v44
    var v47 uint32 = 738
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

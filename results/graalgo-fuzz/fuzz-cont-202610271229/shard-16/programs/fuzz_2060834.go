package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 int, p1 int, p2 int32) uint16 {
    p0 += (-(p0))
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 += (p1 + (-(fi5)))
        _ = fi5
    }
    v6 := GenS1{F0: true, F1: 18446744073709551615, F2: 255}
    p1 = (p0 * 32768)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint64, p1 uint, p2 int8) uint {
    p2 = p2
    var v7 uint32 = 0
    v8 := GenS2{F0: 32767, F1: 87, F2: (v7 & (^(v7)))}
    v7 = uint32(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 + uint(v7))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -1
    var v3 uint = 0
    v4 := make(map[uint32]bool)
    var v9 string = ((("foo bar" + ("hello" + "hello")) + "\n\t") + "hello")
    if (!false) {
        switch -2147483648 {
        case 141:
            var v10 uint16 = 981
            _ = v10
        case -2147483648:
            v3 += v3
        case 2147483647:
            var v11 uint64 = 363
            _ = v11
        }
        _ = len(v1)
        var v12 uint16 = 0
        _ = v12
    }
    var v13 uint64 = 113
    v14 := GenS2{F0: 1, F1: 42, F2: 637}
    switch "a" {
    case "\n\t":
        v2 += (v2 & v2)
    case "foo bar":
        switch (v9 + v9) {
        case "foo bar":
            v15 := GenS3{F0: -1, F1: (v2 % (((-(v2))) | 1))}
            _ = v15
        case "hello":
            fmt.Println(v2, v3)
        default:
            v2 += -2147483648
        }
    default:
        fmt.Println(v2, v13)
    }
    var v16 uint = genFn1((v13 >> ((194) & 63)), v3, 42)
    if false {
        v17 := []int{}
        for fi18 := 0; fi18 < 20; fi18++ {
            if len(v17) > 0 {
                v17[0] = (fi18 ^ 772)
            }
            _ = fi18
        }
        var v19 int16 = -1
        _ = v17
        _ = v19
    }
    v20 := GenS2{F0: 1000, F1: (v2 >> ((v16) & 63)), F2: 1}
    v3 = 42
    if (32768 < -2147483648) {
        var v21 int = 632
        var v22 uint8 = 30
        var v23 string = (v9 + v9)
        _ = v21
        _ = v22
        _ = v23
    } else {
        switch v9 {
        case "hello":
            fmt.Println(v2, v16, v16)
        case "a":
            switch "foo bar" {
            case "hello":
                v24 := v1[("a" + ((("a" + v9) + "foo bar") + "foo bar"))]
                _ = v24
            case "foo bar":
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 uint8 = 248
                    fmt.Println(v16, v3)
                    v27 := []int{(fi25 << ((v3) & 63))}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        case "foo bar":
            var v28 int32 = (v2 & (v2 & (^(v2))))
            _ = v28
        default:
            switch "\n\t" {
            case "a":
                if ((1 > 42) || true) {
                    switch int(v13) {
                    case 32767:
                        switch 127 {
                        case 2147483647:
                            var v29 uint = (v3 - 18446744073709551615)
                            _ = v29
                        case 726:
                            var v30 uint32 = 2147483647
                            _ = v30
                        }
                    case 94:
                        switch -1 {
                        case 42:
                            var v31 uint16 = 1
                            _ = v31
                        case 1:
                            fmt.Println(v3, v2)
                        default:
                            var v32 string = v9
                            _ = v32
                        }
                    case -2147483648:
                        v13 -= (v13 * (v13 & 398))
                    }
                    v33 := []int{29, 402, -971, 1}
                    var v34 uint8 = 2
                    _ = v33
                    _ = v34
                } else {
                    v9 = "a"
                    for fi35 := 0; fi35 < 11; fi35++ {
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 uint32 = 36
                            v38 := GenS4{F0: v37, F1: ((^(fi36)) < fi36)}
                            v39 := []int{fi35}
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 20; fi40++ {
                            v41 := GenS1{F0: true, F1: 4294967295, F2: 93}
                            _ = fi40
                            _ = v41
                        }
                        var v42 uint32 = 2
                        _ = fi35
                        _ = v42
                    }
                    v43 := GenS3{F0: 107, F1: v2}
                    _ = v43
                }
            case "":
                for fi44 := 0; fi44 < 18; fi44++ {
                    for fi45 := 0; fi45 < 17; fi45++ {
                        var v46 uint8 = 128
                        v47 := []int{}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    _ = fi44
                }
            default:
                v48 := []int{-267}
                _ = v48
            }
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            if (!(false && false)) {
                for fi50 := 0; fi50 < 16; fi50++ {
                    var v51 float32 = 0.0
                    v9 = v9
                    v51 += (1.5 + v51)
                    _ = fi50
                    _ = v51
                }
            } else {
                fmt.Println(v2, v13)
                var v52 uint16 = 1
                v3 = (v3 | (^(v3)))
                _ = v52
            }
            v53 := GenS3{F0: -1, F1: v2}
            v9 = v9
            _ = fi49
            _ = v53
        }
        fmt.Println(v3, v9, v2)
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        if (v16 != v3) {
            switch (fi54 * 42) {
            case 0:
                if false {
                    v9 = v9
                    var v55 int8 = -128
                    if true {
                        v56 := GenS4{F0: 623, F1: (!true)}
                        fmt.Println(v3)
                        _ = v56
                    }
                    _ = v55
                }
            }
        }
        v57 := GenS5{F0: 32767}
        _ = fi54
        _ = v57
    }
    v58 := []int{}
    if false {
        v16 = genFn1(357, 128, -128)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v58
    _ = v9
}

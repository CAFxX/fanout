package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 string; F1 uint8; F2 int64 }
func genFn0(p0 int64) uint8 {
    var v6 int16 = -1
    if false {
        switch "foo bar" {
        case "foo bar":
            v7 := []int{-1, 56, 104}
            _ = v7
        default:
            for fi8 := 0; fi8 < 15; fi8++ {
                switch fi8 {
                case 1:
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 uint64 = 523
                        if (((!true) && false) || false) {
                            var v11 bool = ((fi8 * fi8) == (^(fi9)))
                            if false {
                                var v12 int32 = -1
                                v13 := []int{}
                                v13 = append(v13, (fi9 + fi9))
                                _ = v12
                                _ = v13
                            }
                            var v14 int = (fi9 % ((-336) | 1))
                            _ = v11
                            _ = v14
                        }
                        _ = fi9
                        _ = v10
                    }
                case 343:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := []int{-544, fi8}
                        _ = fi15
                        _ = v16
                    }
                case 127:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 uint32 = 42
                        v18 = (v18 & v18)
                        _ = fi17
                        _ = v18
                    }
                default:
                    v19 := []int{(^(fi8)), (-(fi8))}
                    _ = v19
                }
                v20 := []int{fi8, (fi8 >> ((616) & 63)), fi8, fi8}
                _ = fi8
                _ = v20
            }
        }
        v21 := []int{}
        _ = v21
    } else {
        if false {
            p0 += p0
            switch (("" + "hello") + "a") {
            case "a":
                var v22 int = 1
                _ = v22
            default:
                v6 = v6
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 uint16 = 0
                v25 := []int{(fi23 / ((int(p0)) | 1)), (fi23 | (fi23 >> ((42) & 63))), -9223372036854775808}
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            if (-2.25 == 3.14159) {
                for fi26 := 0; fi26 < 20; fi26++ {
                    v6 += -224
                    var v27 uint64 = 1
                    _ = fi26
                    _ = v27
                }
                var v28 int16 = (v6 + (v6 << ((18446744073709551615) & 63)))
                v29 := GenS1{F0: -2147483648, F1: -1}
                _ = v28
                _ = v29
            } else {
                v6 -= v6
                for fi30 := 0; fi30 < 19; fi30++ {
                    switch ("foo bar" + ("" + "a")) {
                    case "foo bar":
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v32 := []int{(fi30 - fi31), fi31, 880}
                            v33 := GenS2{F0: 1, F1: 222}
                            v34 := GenS3{F0: 1, F1: 1}
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    case "a":
                        p0 -= (p0 | p0)
                    }
                    var v35 uint16 = 32767
                    _ = fi30
                    _ = v35
                }
            }
            v36 := GenS4{F0: 42}
            if true {
                var v37 int = -431
                var v38 uint64 = 214
                var v39 uint64 = (^(v38))
                _ = v37
                _ = v38
                _ = v39
            }
            _ = v36
        }
        p0 += p0
    }
    v6 = (v6 >> ((65535) & 63))
    _ = p0
    _ = v6
    return 128
}
func genFn1(p0 uint, p1 int) uint64 {
    fmt.Println(p1, p0)
    v40 := GenS5{F0: p1, F1: -2147483648, F2: 42}
    _ = p0
    _ = p1
    _ = v40
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 int16 = 32767
    var v4 int32 = -1
    var v5 int64 = -9223372036854775808
    fmt.Println(v4, v4)
    v4 += (v4 / (((v4 | (v4 << ((815) & 63)))) | 1))
    var v41 int64 = v5
    v42 := GenS6{F0: 2147483647, F1: 2147483648, F2: 42}
    v3 = (v3 * v3)
    v43 := GenS7{F0: 42, F1: false}
    for fi44 := 0; fi44 < 9; fi44++ {
        switch 157 {
        case 128:
            v45 := GenS8{F0: ("a" + ("\n\t" + "a")), F1: 255, F2: (v5 | v41)}
            _ = v45
        case -126:
            if (!true) {
                if ((1 > -1) || true) {
                    v5 += (v41 - v41)
                    delete(v1, 2147483647)
                    var v46 string = "\n\t"
                    _ = v46
                } else {
                    var v47 float64 = float64(3.14159)
                    v48 := GenS4{F0: 127}
                    _ = v47
                    _ = v48
                }
            }
        case -860:
            v3 -= int16(v4)
        default:
            fmt.Println(fi44)
        }
        if true {
            for fi49 := 0; fi49 < 2; fi49++ {
                var v50 uint16 = 552
                switch "hello" {
                case "foo bar":
                    v50 = v50
                }
                _ = fi49
                _ = v50
            }
            var v51 int16 = v3
            v3 -= v51
            _ = v51
        }
        if ("foo bar" <= (("a" + (("foo bar" + "") + "\n\t")) + ("a" + ("" + "\n\t")))) {
            if ((v4 * 1) != (-(v4))) {
                var v52 uint16 = 379
                _ = v52
            }
        }
        _ = fi44
    }
    var v53 int64 = v5
    if ((0.1 + (3.14159 * 0.1)) <= 0.0) {
        if ("a" >= "hello") {
            _ = len(v2)
            var v54 float64 = -2.25
            v4 = (v4 << ((1) & 63))
            _ = v54
        } else {
            fmt.Println(v3, v5)
            var v55 int16 = v3
            _ = v55
        }
    }
    v5 = -483
    var v56 uint32 = 906
    v57 := v2["a"]
    v53 = v41
    fmt.Println(v53, v41, v56)
    if true {
        for fi58 := 0; fi58 < 15; fi58++ {
            var v59 string = ("foo bar" + "hello")
            v53 += (v53 >> ((42) & 63))
            _ = fi58
            _ = v59
        }
        v56 += (^(v56))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v53
    _ = v56
    _ = v57
}

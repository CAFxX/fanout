package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 int32 }
type GenS5 struct { F0 int32; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 uint, p1 string) uint8 {
    var v5 int32 = 2147483647
    var v6 int8 = int8(v5)
    v7 := []int{490}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1() int8 {
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int32 = -2147483648
        fmt.Println(v9, fi8)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint64 = 810
            if (!false) {
                v13 := []int{128, 32768, (^(fi11))}
                if true {
                    v12 = v12
                }
                _ = v13
            } else {
                v14 := []int{(fi10 - fi11)}
                var v15 bool = true
                if (!v15) {
                    var v16 int = (^(fi11))
                    v17 := GenS1{F0: 1}
                    v18 := GenS2{F0: 475, F1: 0, F2: "a"}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS2{F0: 32767, F1: (fi11 ^ fi11), F2: ""}
                    v15 = v15
                    _ = v19
                }
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if (("hello" + "hello") > ((("" + ((("hello" + ("\n\t" + (((("\n\t" + ("a" + "")) + "foo bar") + "") + ("" + ("foo bar" + "foo bar"))))) + "foo bar") + "\n\t")) + "\n\t") + "foo bar")) {
            if ((("\n\t" + ("" + "foo bar")) + ("hello" + "hello")) >= "\n\t") {
                for fi21 := 0; fi21 < 15; fi21++ {
                    if true {
                        var v22 uint = 491
                        _ = v22
                    } else {
                        if (-1 == 0) {
                            var v23 float32 = -2.25
                            var v24 string = "foo bar"
                            if false {
                                if ("hello" > ("foo bar" + (("" + v24) + v24))) {
                                    fmt.Println(v23)
                                }
                                v23 = (-2.25 / v23)
                                var v25 int16 = -869
                                var v26 int16 = 1000
                                fmt.Println(fi21)
                                v24 = (v24 + "foo bar")
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                            _ = v24
                        }
                        v27 := GenS2{F0: 897, F1: (^(fi20)), F2: "hello"}
                        _ = v27
                    }
                    _ = fi21
                }
                var v28 float64 = (-2.25 / 1.0)
                _ = v28
            }
            var v29 uint16 = 65535
            _ = v29
        } else {
            var v30 uint8 = 255
            _ = v30
        }
        v31 := []int{}
        _ = fi20
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS1{F0: 0}
        var v34 int = fi32
        v35 := []int{24, (fi32 ^ fi32), 128, 128}
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := []int{32767}
    _ = v36
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 523
    var v3 float32 = (3.14159 * 1e10)
    var v4 uint32 = 4294967295
    if (1 != 666) {
        switch 9223372036854775807 {
        case 463:
            for fi37 := 0; fi37 < 13; fi37++ {
                v38 := v1[(fi37 | fi37)]
                fmt.Println(fi37)
                _ = fi37
                _ = v38
            }
        case 4294967295:
            v39 := GenS3{F0: 4294967295}
            _ = v39
        default:
            fmt.Println(v4, v4)
        }
        var v40 uint64 = 18446744073709551615
        _ = v40
    } else {
        v4 -= v4
        switch "hello" {
        case "foo bar":
            if false {
                v2 += 2
            }
        default:
            for fi41 := 0; fi41 < 18; fi41++ {
                v42 := GenS4{F0: 2, F1: 593, F2: 1}
                v43 := v1[(fi41 ^ fi41)]
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    }
    var v44 uint64 = 4294967295
    v3 += 3.14159
    fmt.Println(v44, v2, v3)
    v4 = v4
    v2 = (v2 / (((^(v2))) | 1))
    switch "foo bar" {
    case "\n\t":
        for fi45 := 0; fi45 < 11; fi45++ {
            fmt.Println(fi45, v44, v4)
            _ = fi45
        }
    case "a":
        v44 += (v44 ^ v44)
    case "foo bar":
        for fi46 := 0; fi46 < 1; fi46++ {
            if (true && true) {
                v47 := GenS1{F0: -1}
                if (("foo bar" + "a") < "hello") {
                    var v48 float32 = v3
                    _ = v48
                }
                v2 = (^(v2))
                _ = v47
            } else {
                var v49 int16 = 0
                switch "\n\t" {
                case "foo bar":
                    v2 = (^(v2))
                case "a":
                    if true {
                        if (128 == 128) {
                            v50 := GenS5{F0: 1, F1: v4, F2: 0}
                            var v51 uint = 255
                            _ = v50
                            _ = v51
                        } else {
                            v52 := GenS6{F0: v2, F1: "foo bar"}
                            _ = v52
                        }
                        if false {
                            if true {
                                var v53 uint32 = v4
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    v55 := GenS7{F0: v49}
                                    var v56 int64 = 1
                                    v57 := GenS8{F0: fi46, F1: (-(v56))}
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                                _ = v53
                            } else {
                                v58 := GenS4{F0: 42, F1: 0, F2: 0}
                                if (true && true) {
                                    _ = len(v1)
                                    v49 = (v49 | (v49 - (v49 << ((0) & 63))))
                                } else {
                                    v59 := v1[fi46]
                                    var v60 uint64 = v44
                                    _ = v59
                                    _ = v60
                                }
                                var v61 string = "foo bar"
                                _ = v58
                                _ = v61
                            }
                            fmt.Println(v44)
                        }
                        for fi62 := 0; fi62 < 11; fi62++ {
                            _ = len(v1)
                            _ = fi62
                        }
                    } else {
                        var v63 bool = (!true)
                        _ = len(v1)
                        v64 := v1[fi46]
                        _ = v63
                        _ = v64
                    }
                default:
                    var v65 bool = (2147483647 < uint32(v44))
                    _ = v65
                }
                _ = v49
            }
            switch "hello" {
            case "a":
                v66 := []int{(^(fi46)), fi46, fi46, (fi46 + fi46)}
                _ = v66
            }
            v1[127] = (^(fi46))
            _ = fi46
        }
    }
    var v67 uint16 = (v2 * v2)
    if (false || (((0.1 - 0.1) < (float64(v3) - 0.1)) || true)) {
        v68 := GenS4{F0: 128, F1: 42, F2: 0}
        switch 1 {
        case -426:
            v3 += (float32(float64(v3)) * (v3 + 3.14159))
        case 1:
            for fi69 := 0; fi69 < 2; fi69++ {
                v44 = v44
                v70 := GenS7{F0: -32768}
                _ = fi69
                _ = v70
            }
        default:
            if false {
                v1[42] = 128
            }
        }
        _ = v68
    } else {
        var v71 string = "\n\t"
        v1[661] = -1
        _ = v71
    }
    fmt.Println(v3)
    v72 := GenS9{F0: (v44 - (v44 >> ((65535) & 63))), F1: 818, F2: -283}
    v73 := GenS3{F0: (v44 ^ v44)}
    var v74 int8 = 127
    v75 := GenS2{F0: v67, F1: -1, F2: "foo bar"}
    for fi76 := 0; fi76 < 14; fi76++ {
        switch (fi76 >> ((42) & 63)) {
        case -733:
            var v77 int8 = v74
            _ = v77
        default:
            v44 = 2
        }
        var v78 uint16 = v2
        var v79 float32 = v3
        _ = fi76
        _ = v78
        _ = v79
    }
    v44 = 739
    v44 = (v44 / (((v44 - v44)) | 1))
    fmt.Println(v67)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v67:", v67)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v67
    _ = v72
    _ = v73
    _ = v74
    _ = v75
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int64 }
type GenS3 struct { F0 string }
func genFn0() int16 {
    var v6 int32 = 797
    v7 := GenS1{F0: (("\n\t" >= (("foo bar" + (("hello" + "a") + (("" + (("hello" + "") + "")) + "foo bar"))) + ("a" + ""))) && (v6 < (v6 * 641))), F1: 59, F2: 42}
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 int, p2 int) int8 {
    switch (p1 % ((-9223372036854775808) | 1)) {
    case 127:
        switch "\n\t" {
        case "a":
            v8 := GenS1{F0: false, F1: 1, F2: 4294967295}
            _ = v8
        case "hello":
            if true {
                if (!(9223372036854775807 <= 0)) {
                    var v9 uint64 = 1
                    if ((v9 % (((v9 ^ 4294967295)) | 1)) < (v9 * v9)) {
                        var v10 int = p1
                        fmt.Println(p2, v9)
                        var v11 uint32 = 1
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                } else {
                    p0 = 116
                }
                p0 += 126
            } else {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{p2, (p1 & -84), p2}
                    for fi14 := 0; fi14 < 1; fi14++ {
                        p1 = fi14
                        var v15 int = (fi12 << ((352) & 63))
                        _ = fi14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            }
        case "foo bar":
            p2 = p2
        default:
            p0 += (-(p0))
        }
    case 0:
        var v16 uint16 = 65535
        _ = v16
    case -1:
        var v17 float64 = (3.14159 / 1.0)
        _ = v17
    default:
        v18 := GenS1{F0: true, F1: 32767, F2: 2}
        _ = v18
    }
    var v19 uint64 = 898
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 373
    var v3 int8 = 1
    v4 := make(map[string]uint32)
    var v5 string = "hello"
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint = 255
            var v23 bool = (42 > -207)
            switch (v5 + v5) {
            case "foo bar":
                var v24 int8 = int8(v2)
                _ = v24
            case "\n\t":
                switch (fi21 % ((int(v2)) | 1)) {
                case 42:
                    v22 = 18446744073709551615
                case -2147483648:
                    if v23 {
                        if v23 {
                            v22 -= v22
                            var v25 int32 = 0
                            _ = v25
                        }
                    }
                default:
                    v2 = (^(v2))
                }
            case "":
                switch ((v5 + "\n\t") + v5) {
                case "":
                    if true {
                        var v26 uint32 = 340
                        _ = v26
                    } else {
                        v27 := GenS2{F0: 2, F1: v22, F2: 9223372036854775807}
                        v28 := GenS3{F0: v5}
                        switch fi20 {
                        case 1:
                            v2 -= v2
                        default:
                            var v29 float64 = 3.14159
                            _ = v29
                        }
                        v2 += 2147483647
                        _ = v27
                        _ = v28
                    }
                default:
                    v5 = "\n\t"
                }
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        var v30 float32 = 1e10
        if true {
            var v31 string = ""
            v30 += -2.25
            if true {
                var v32 float32 = (v30 * 0.1)
                if ((v3 / (((v3 * v3)) | 1)) != (v3 + v3)) {
                    var v33 bool = true
                    var v34 int16 = 42
                    _ = v33
                    _ = v34
                } else {
                    var v35 int32 = 1
                    switch fi20 {
                    case 4294967295:
                        switch fi20 {
                        case 0:
                            if ((1e10 + v32) != (1.5 * float32(-2.25))) {
                                fmt.Println(v3, v3, v30)
                                var v36 string = v31
                                _ = v36
                            }
                        case 2147483647:
                            var v37 bool = (v35 != (v35 / (((v35 + 0)) | 1)))
                            _ = v37
                        case 9223372036854775807:
                            for fi38 := 0; fi38 < 11; fi38++ {
                                v35 = 399
                                _ = fi38
                            }
                        default:
                            v3 += (v3 * (-(v3)))
                        }
                    case -924:
                        var v39 float64 = (0.1 + 3.14159)
                        _ = v39
                    case -9223372036854775808:
                        v31 = (("\n\t" + (v5 + "hello")) + v5)
                    }
                    _ = v35
                }
                for fi40 := 0; fi40 < 11; fi40++ {
                    v30 = v32
                    v41 := GenS1{F0: false, F1: 398, F2: 915}
                    delete(v1, "")
                    _ = fi40
                    _ = v41
                }
                _ = v32
            }
            _ = v31
        }
        _ = fi20
        _ = v30
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        var v43 uint32 = v2
        fmt.Println(fi42)
        if ((^(v2)) <= (v2 / ((v2) | 1))) {
            v43 = (v2 - v2)
            var v44 bool = true
            v3 = v3
            _ = v44
        }
        _ = fi42
        _ = v43
    }
    if false {
        for fi45 := 0; fi45 < 9; fi45++ {
            var v46 float64 = 3.14159
            _ = fi45
            _ = v46
        }
    } else {
        v5 = (v5 + v5)
        var v47 int = 1
        v48 := GenS2{F0: 2, F1: 94, F2: -9223372036854775808}
        _ = v47
        _ = v48
    }
    var v49 float32 = (3.14159 * ((float32(3.14159) - 0.0) - float32(1.5)))
    delete(v4, (((v5 + "") + (("" + "a") + v5)) + "foo bar"))
    delete(v1, "\n\t")
    v49 += (0.0 + v49)
    fmt.Println(v5, v3)
    v50 := []int{0, 32767}
    for fi51 := 0; fi51 < 13; fi51++ {
        v2 -= (v2 + (v2 >> ((42) & 63)))
        _ = len(v50)
        for fi52 := 0; fi52 < 12; fi52++ {
            v49 = (1.5 / float32(-2.25))
            _ = fi52
        }
        _ = fi51
    }
    var v53 uint = 901
    v54 := GenS2{F0: 606, F1: v53, F2: 0}
    var v55 int8 = v3
    v3 -= genFn1((v55 << ((v53) & 63)), int(v2), 2147483647)
    for fi56 := 0; fi56 < 4; fi56++ {
        fmt.Println(v53)
        if false {
            v2 += v2
        } else {
            v4[v5] = (v2 / (((v2 | (^(v2)))) | 1))
            var v57 uint64 = 42
            _ = v57
        }
        for fi58 := 0; fi58 < 20; fi58++ {
            switch "hello" {
            case "foo bar":
                var v59 bool = (false || true)
                _ = v59
            case "a":
                if false {
                    v53 = v53
                    if ((((v53 / ((v53) | 1)) != (v53 / ((1) | 1))) || (-1 < 158)) && true) {
                        if len(v50) >= 2 {
                            v50 = v50[1:len(v50)]
                        }
                        var v60 string = (v5 + "a")
                        _ = len(v50)
                        _ = v60
                    } else {
                        fmt.Println(v53, fi58)
                        v2 += 446
                        v55 = (v3 | v55)
                    }
                    v61 := v4[v5]
                    _ = v61
                }
            default:
                v62 := GenS1{F0: (genFn0() != -1), F1: -857, F2: 255}
                _ = v62
            }
            _ = fi58
        }
        _ = fi56
    }
    v55 = genFn1(-1, int(v2), -1)
    if true {
        for fi63 := 0; fi63 < 12; fi63++ {
            var v64 uint16 = 717
            _ = fi63
            _ = v64
        }
        var v65 int8 = v3
        _ = v65
    } else {
        fmt.Println(v2, v49)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v55
}

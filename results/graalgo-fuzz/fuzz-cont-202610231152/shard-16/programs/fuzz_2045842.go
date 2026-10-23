package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
type GenS3 struct { F0 int32; F1 int16; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint) int {
    p0 += p0
    var v4 int16 = -32768
    _ = p0
    _ = p1
    _ = v4
    return -40
}
func genFn1(p0 bool) uint {
    p0 = (-1 != 2147483647)
    fmt.Println(p0)
    _ = p0
    return 4294967295
}

func main() {
    var v1 int16 = 42
    var v2 string = (("foo bar" + (((("a" + "a") + ("\n\t" + "a")) + ("\n\t" + ("foo bar" + ("hello" + "")))) + (("a" + "") + (("a" + "") + "hello")))) + "")
    v3 := make(map[int]int)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch v2 {
        case "foo bar":
            if ((!(true && (2 != 947))) && true) {
                v6 := GenS1{F0: 42, F1: 998}
                var v7 int8 = -1
                _ = v6
                _ = v7
            } else {
                if false {
                    v2 = "\n\t"
                    if true {
                        _ = len(v3)
                        switch (fi5 & fi5) {
                        case 127:
                            var v8 int8 = -27
                            _ = v8
                        case 9223372036854775807:
                            if (((0.0 + float32(0.1)) >= (1e10 + (0.1 - -2.25))) && true) {
                                if (true && false) {
                                    v9 := []int{32768, (^(fi5))}
                                    _ = v9
                                } else {
                                    v10 := GenS1{F0: 255, F1: 885}
                                    if (0 >= 42) {
                                        var v11 float32 = 1.5
                                        for fi12 := 0; fi12 < 16; fi12++ {
                                            var v13 int16 = v1
                                            var v14 uint32 = 2147483647
                                            var v15 uint8 = 253
                                            if (5 < 1) {
                                                var v16 float64 = 0.0
                                                _ = v16
                                            } else {
                                                _ = len(v3)
                                                var v17 uint64 = 37
                                                var v18 int16 = 922
                                                _ = v17
                                                _ = v18
                                            }
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = v11
                                    }
                                    _ = v10
                                }
                            } else {
                                v19 := v3[(fi5 / ((fi5) | 1))]
                                var v20 int32 = -65
                                _ = v19
                                _ = v20
                            }
                        case 2147483647:
                            for fi21 := 0; fi21 < 20; fi21++ {
                                fmt.Println(fi5, fi21)
                                if false {
                                    fmt.Println(fi21, fi5)
                                    v22 := GenS1{F0: 163, F1: 4294967295}
                                    _ = v22
                                } else {
                                    fmt.Println(v1, fi5, v2)
                                    v23 := GenS1{F0: 0, F1: 2}
                                    switch ("hello" + v2) {
                                    case "a":
                                        var v24 int8 = -45
                                        _ = v24
                                    case "foo bar":
                                        var v25 float64 = 1e10
                                        _ = v25
                                    case "":
                                        switch (v2 + v2) {
                                        case "foo bar":
                                            if (!(true || (2 < 174))) {
                                                var v26 uint64 = 405
                                                _ = v26
                                            } else {
                                                v27 := v3[(fi5 % ((630) | 1))]
                                                if (!((float32(float64((0.0 + 3.14159))) > 3.14159) && (false || true))) {
                                                    v1 = (^(v1))
                                                } else {
                                                    v28 := GenS1{F0: 2, F1: 42}
                                                    var v29 uint32 = 2147483648
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v2 = "a"
                                                _ = v27
                                            }
                                        case "\n\t":
                                            v2 = v2
                                        case "hello":
                                            v30 := GenS2{F0: 2, F1: 975, F2: 65535}
                                            _ = v30
                                        default:
                                            var v31 int16 = (v1 % ((32767) | 1))
                                            _ = v31
                                        }
                                    default:
                                        if (2 >= 661) {
                                            fmt.Println(fi5, fi21)
                                            switch fi5 {
                                            case 410:
                                                switch ("foo bar" + v2) {
                                                case "a":
                                                    var v32 bool = true
                                                    _ = v32
                                                }
                                            }
                                            v33 := GenS3{F0: 2147483647, F1: 0, F2: 255}
                                            _ = v33
                                        } else {
                                            var v34 int = (^(fi5))
                                            _ = v34
                                        }
                                    }
                                    _ = v23
                                }
                                var v35 uint32 = 2147483647
                                var v36 int16 = v1
                                v37 := []int{(fi21 + genFn0(4294967295, genFn1((!((!false) || (true && ((1e10 / (float64((0.0 - ((-2.25 + 3.14159) * 1e10))) - 3.14159)) > (-2.25 - float64(float32(float64((float32(-2.25) - (3.14159 + (3.14159 / (((1.5 / (0.0 - 1e10)) - float32(float64(0.1))) + 3.14159))))))))))))))), fi21}
                                _ = fi21
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        }
                        var v38 bool = false
                        _ = v38
                    }
                }
                var v39 int = genFn0(945, 212)
                if (3.14159 < (3.14159 - ((-2.25 / 1e10) + (((1e10 - 1e10) + (((1.5 - float32(float64(((float32(float64(0.0)) + 3.14159) / 0.1)))) - 0.1) - 0.0)) * ((0.1 + (1.5 / ((3.14159 / (3.14159 + (0.0 - (0.1 + 0.1)))) / 0.1))) * (float32(-2.25) * (0.0 - -2.25))))))) {
                    v1 = 32767
                    v40 := []int{(v39 + (fi5 - genFn0(42, 720))), v39, 42, fi5}
                    v1 += 97
                    _ = v40
                }
                _ = v39
            }
        case "a":
            if true {
                v1 = v1
            } else {
                v1 -= (v1 | -1)
                var v41 uint16 = 1
                for fi42 := 0; fi42 < 18; fi42++ {
                    var v43 uint64 = 708
                    v2 = v2
                    v44 := []int{}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = v41
            }
        default:
            if false {
                v1 = (-(v1))
                var v45 int8 = -128
                _ = v45
            }
        }
        if false {
            if true {
                v46 := GenS4{F0: v1, F1: 255, F2: 558}
                _ = v46
            } else {
                v2 = (v2 + ((v2 + "a") + (v2 + (v2 + (v2 + v2)))))
                for fi47 := 0; fi47 < 19; fi47++ {
                    v48 := v3[(-(fi47))]
                    v2 = "hello"
                    _ = fi47
                    _ = v48
                }
                _ = len(v3)
            }
        }
        v2 = "foo bar"
        _ = fi5
    }
    if ((((v2 + (v2 + "\n\t")) + (v2 + v2)) == "a") && true) {
        if true {
            switch v2 {
            case "\n\t":
                v1 -= v1
            case "foo bar":
                var v49 float32 = 1e10
                _ = v49
            case "":
                v50 := GenS3{F0: 42, F1: v1, F2: 1}
                _ = v50
            default:
                v1 += 32767
            }
            var v51 uint16 = 2
            _ = v51
        } else {
            fmt.Println(v2, v1)
        }
    }
    var v52 bool = true
    switch -579 {
    case -2147483648:
        v1 = (v1 >> ((886) & 63))
    case 128:
        if false {
            v1 -= v1
            if v52 {
                for fi53 := 0; fi53 < 13; fi53++ {
                    for fi54 := 0; fi54 < 12; fi54++ {
                        var v55 string = v2
                        v1 += 42
                        var v56 int16 = v1
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    _ = fi53
                }
            } else {
                fmt.Println(v1)
            }
            v57 := []int{120, 4294967295, genFn0(346, 1), 2147483647}
            _ = v57
        } else {
            for fi58 := 0; fi58 < 2; fi58++ {
                for fi59 := 0; fi59 < 4; fi59++ {
                    if v52 {
                        delete(v3, 1)
                    }
                    if true {
                        delete(v3, genFn0(18446744073709551615, 0))
                        var v60 float64 = (((float64(1e10) - 1.5) - float64((0.0 * 0.1))) * 0.0)
                        fmt.Println(v1)
                        fmt.Println(v52, fi59)
                        _ = v60
                    }
                    v61 := []int{(fi59 % ((fi58) | 1)), fi59, fi58, (fi59 ^ fi58)}
                    _ = fi59
                    _ = v61
                }
                _ = fi58
            }
        }
    case 1:
        var v62 string = v2
        _ = v62
    }
    for fi63 := 0; fi63 < 16; fi63++ {
        v2 = ("foo bar" + "\n\t")
        _ = fi63
    }
    _ = len(v3)
    for fi64 := 0; fi64 < 13; fi64++ {
        v1 = -32768
        switch fi64 {
        case 9223372036854775807:
            v52 = ((42 < 0) || (fi64 <= fi64))
        default:
            var v65 uint32 = 42
            _ = v65
        }
        switch (fi64 & genFn0(4294967295, 255)) {
        case -1:
            var v66 float64 = 1.5
            _ = v66
        }
        _ = fi64
    }
    v3[-362] = 128
    v67 := GenS5{F0: 1}
    var v68 bool = false
    v68 = v52
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v67
    _ = v68
}

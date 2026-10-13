package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0(p0 uint32, p1 uint64, p2 int) uint16 {
    p0 = p0
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint32 {
    switch 228 {
    case 997:
        switch 128 {
        case -231:
            var v5 uint = 42
            _ = v5
        }
    case 875:
        if (true && (2 == 1)) {
            v6 := GenS1{F0: 892, F1: true}
            _ = v6
        }
    default:
        v7 := []int{0, -9223372036854775808, 127}
        _ = v7
    }
    if (1 >= 1) {
        var v8 uint16 = genFn0(786, 4294967295, 434)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 bool = (32767 <= -398)
            v11 := GenS2{F0: "foo bar", F1: "hello"}
            var v12 uint = 285
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if true {
        switch "foo bar" {
        case "foo bar":
            var v13 uint8 = 1
            _ = v13
        default:
            v14 := []int{}
            _ = v14
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            if false {
                switch (^(fi15)) {
                case 42:
                    if false {
                        var v16 int32 = 42
                        v17 := GenS1{F0: 42, F1: (v16 == v16)}
                        _ = v16
                        _ = v17
                    }
                case 2147483647:
                    var v18 string = "\n\t"
                    _ = v18
                case 9223372036854775807:
                    v19 := GenS3{F0: 1, F1: 0}
                    _ = v19
                }
                fmt.Println(fi15)
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
            var v21 uint = 2
            for fi22 := 0; fi22 < 20; fi22++ {
                v21 = 255
                v23 := []int{fi22, 2147483647, 859, fi15}
                _ = fi22
                _ = v23
            }
            _ = fi15
            _ = v21
        }
        var v24 int8 = 1
        _ = v24
    }
    v25 := []int{32768, 127, -551}
    switch -358 {
    case 4294967295:
        var v26 int32 = -2147483648
        _ = v26
    case 42:
        if (!(!true)) {
            var v27 int32 = -1
            var v28 int8 = 1
            v28 = (v28 % (((-(v28))) | 1))
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                if false {
                    var v30 int64 = 0
                    if (float32(1e10) > 0.1) {
                        switch fi29 {
                        case -321:
                            if (148 >= 18446744073709551615) {
                                fmt.Println(v30)
                                var v31 uint16 = 2
                                _ = v31
                            }
                        default:
                            v32 := GenS4{F0: "", F1: 2, F2: 42}
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 15; fi33++ {
                            v30 = v30
                            v34 := GenS5{F0: ("\n\t" + "hello"), F1: 2, F2: 42}
                            if true {
                                v35 := GenS6{F0: "hello", F1: 13}
                                _ = v35
                            } else {
                                v36 := GenS7{F0: 2147483648}
                                switch fi33 {
                                case -652:
                                    switch fi29 {
                                    case -1:
                                        var v37 uint = 4294967295
                                        _ = v37
                                    default:
                                        var v38 bool = false
                                        _ = v38
                                    }
                                case -9223372036854775808:
                                    v39 := GenS5{F0: (("hello" + "a") + ("\n\t" + "\n\t")), F1: 0, F2: int32(v30)}
                                    _ = v39
                                case 42:
                                    v40 := GenS3{F0: -1, F1: (-(v30))}
                                    _ = v40
                                }
                                if (!(!(("a" + (("foo bar" + "\n\t") + (("\n\t" + ("hello" + "")) + "hello"))) == ("hello" + ("hello" + "hello"))))) {
                                    var v41 float32 = 0.1
                                    var v42 bool = false
                                    v43 := GenS8{F0: 32767, F1: 102, F2: 42}
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    if (42 < 32) {
                                        v25 = append(v25, (fi33 % ((fi33) | 1)))
                                    } else {
                                        _ = len(v25)
                                        if len(v25) > 0 {
                                            v25[0] = (fi33 ^ 128)
                                        }
                                        fmt.Println(v30, fi33, v30)
                                    }
                                    v30 = (-(v30))
                                    var v44 float32 = -2.25
                                    _ = v44
                                }
                                _ = v36
                            }
                            var v45 int32 = 42
                            fmt.Println(v30)
                            _ = fi33
                            _ = v34
                            _ = v45
                        }
                        v46 := GenS4{F0: "a", F1: 2, F2: 185}
                        _ = v46
                    } else {
                        _ = len(v25)
                        _ = len(v25)
                    }
                    var v47 uint64 = 42
                    _ = v30
                    _ = v47
                }
                _ = fi29
            }
        }
    }
    _ = v25
    return 65535
}

func main() {
    var v1 uint64 = 938
    var v2 int16 = -32768
    v3 := []string{}
    v4 := []string{"a"}
    if (127 != 1) {
        v48 := GenS4{F0: "\n\t", F1: 42, F2: 1}
        _ = v48
    }
    fmt.Println(v1)
    v4 = append(v4, "a")
    v49 := GenS8{F0: 85, F1: 4294967295, F2: 2}
    if (2 != 177) {
        v1 = v1
    }
    var v50 int64 = -84
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v2, v2)
    default:
        var v51 uint8 = 2
        _ = v51
    }
    v1 = (v1 / ((0) | 1))
    var v52 uint16 = 733
    if false {
        v2 = -1
    }
    if (405 != 94) {
        v2 += 880
    }
    var v53 int32 = -1
    v54 := GenS5{F0: "", F1: 715, F2: int32(v1)}
    switch 2147483647 {
    case -2147483648:
        v4 = append(v4, ("" + ("" + "")))
    default:
        if true {
            v1 += v1
        } else {
            v53 = v53
            var v55 bool = (0 < -1)
            fmt.Println(v55, v2)
            _ = v55
        }
    }
    if (!((127 < -27) && (!true))) {
        for fi56 := 0; fi56 < 6; fi56++ {
            for fi57 := 0; fi57 < 16; fi57++ {
                v50 = 1
                switch fi56 {
                case 1:
                    if (true && true) {
                        fmt.Println(fi57, v1, v2)
                    }
                case 4294967295:
                    v52 = uint16(v1)
                case 42:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi57
            }
            _ = len(v3)
            v4 = append(v4, "")
            _ = fi56
        }
        if false {
            var v58 uint16 = genFn0(535, 4294967295, -693)
            _ = v58
        }
        switch 9223372036854775807 {
        case 1:
            v4 = append(v4, "")
        case -389:
            if true {
                var v59 int64 = 42
                if true {
                    var v60 float64 = -2.25
                    _ = len(v4)
                    _ = v60
                }
                switch 32767 {
                case 643:
                    v61 := GenS1{F0: 18446744073709551615, F1: false}
                    _ = v61
                default:
                    var v62 uint64 = 2
                    _ = v62
                }
                _ = v59
            } else {
                if true {
                    v1 += uint64(v52)
                    fmt.Println(v1, v52, v52)
                } else {
                    v2 = (^(v2))
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
        }
    }
    if (((float64(0.0) + 3.14159) * ((-2.25 * 1e10) * ((-2.25 / 1.5) + 3.14159))) < -2.25) {
        v52 += 970
        var v63 float64 = (1e10 * (0.0 / 1.5))
        for fi64 := 0; fi64 < 13; fi64++ {
            var v65 float32 = ((1e10 * 3.14159) - ((float32(v63) + 1.5) + -2.25))
            _ = fi64
            _ = v65
        }
        _ = v63
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v52
    _ = v53
    _ = v54
}

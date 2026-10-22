package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int32, p1 uint32, p2 bool) int {
    if (!(((1.5 / (1.5 * 3.14159)) * (1.5 / 3.14159)) == ((-2.25 + -2.25) + (-2.25 + 1e10)))) {
        v7 := []int{127, -98, 89}
        var v8 bool = p2
        _ = v7
        _ = v8
    }
    if (p1 == (^(p1))) {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: (^(p1))}
            _ = fi9
            _ = v10
        }
    }
    v11 := []int{-1}
    if (!false) {
        var v12 uint32 = p1
        _ = v12
    }
    v13 := GenS1{F0: (p1 & p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    return 1
}
func genFn1(p0 uint16) uint64 {
    v14 := []int{genFn0(-831, 65535, (766 < 0)), 0, 2147483647, 689}
    var v15 int32 = 1
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 7; fi17++ {
        p0 += 255
        if false {
            if true {
                p0 -= 0
                var v18 float32 = -2.25
                v15 = (v15 >> ((65535) & 63))
                _ = v18
            }
            if len(v14) > 0 {
                v14[0] = fi17
            }
        } else {
            if true {
                v14 = append(v14, fi17)
                var v19 int64 = -293
                fmt.Println(v16, v16, v19)
                _ = v19
            } else {
                v15 = (v15 % ((int32(fi17)) | 1))
                var v20 float64 = v16
                v20 = v20
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                v16 = 0.1
                v16 = v16
                _ = fi21
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        switch "a" {
        case "foo bar":
            switch 168 {
            case -715:
                _ = len(v14)
            case 42:
                var v22 string = "a"
                _ = v22
            default:
                _ = len(v14)
            }
        case "hello":
            if len(v14) > 0 {
                v14[0] = (fi17 ^ 32768)
            }
        }
        _ = fi17
    }
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 211
}

func main() {
    var v1 uint = 506
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 uint64 = 903
    v5 := make(map[uint32]bool)
    var v6 int8 = 127
    v23 := GenS2{F0: 0}
    var v24 float32 = ((((0.0 + 1e10) / -2.25) - 0.1) * 1.5)
    if true {
        _ = len(v3)
    } else {
        if true {
            fmt.Println(v6, v6)
            if false {
                if (false || (true || true)) {
                    if false {
                        v4 -= 380
                    }
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 string = (("foo bar" + "hello") + "foo bar")
                        for fi27 := 0; fi27 < 9; fi27++ {
                            switch "\n\t" {
                            case "\n\t":
                                delete(v5, 2147483648)
                            case "foo bar":
                                if (v26 != v26) {
                                    v28 := GenS1{F0: uint32(v4)}
                                    v6 += v6
                                    _ = v28
                                }
                            case "":
                                v2[v26] = 143
                            default:
                                var v29 float64 = 3.14159
                                _ = v29
                            }
                            _ = fi27
                        }
                        if false {
                            var v30 uint8 = 0
                            _ = v30
                        }
                        _ = fi25
                        _ = v26
                    }
                    for fi31 := 0; fi31 < 18; fi31++ {
                        var v32 float64 = -2.25
                        v33 := v2[(("hello" + "foo bar") + "a")]
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                } else {
                    for fi34 := 0; fi34 < 9; fi34++ {
                        var v35 uint64 = uint64(v1)
                        v6 = v6
                        for fi36 := 0; fi36 < 19; fi36++ {
                            v1 -= 387
                            switch -1 {
                            case 2147483647:
                                _ = len(v3)
                            case -2147483648:
                                var v37 int8 = v6
                                _ = v37
                            }
                            var v38 float64 = 1e10
                            _ = fi36
                            _ = v38
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                var v39 uint = v1
                var v40 float32 = 0.0
                _ = v39
                _ = v40
            } else {
                if false {
                    if true {
                        v1 = v1
                        v24 = v24
                        var v41 int16 = -32768
                        _ = v41
                    }
                    v4 -= 0
                } else {
                    if false {
                        var v42 float32 = ((v24 - (v24 / -2.25)) - v24)
                        var v43 uint8 = 42
                        v1 = v1
                        _ = v42
                        _ = v43
                    }
                    v44 := v2["a"]
                    _ = v44
                }
                if (v6 <= v6) {
                    var v45 int32 = 647
                    v4 = v4
                    if false {
                        v46 := v2[(("\n\t" + ("hello" + "foo bar")) + "")]
                        delete(v5, (v46 >> ((v1) & 63)))
                        _ = v46
                    } else {
                        v47 := v5[4294967295]
                        _ = v47
                    }
                    _ = v45
                } else {
                    v4 -= v4
                    v6 = 1
                    var v48 int32 = 42
                    _ = v48
                }
            }
            var v49 string = "foo bar"
            _ = v49
        } else {
            v4 = v4
        }
        v50 := GenS1{F0: 4294967295}
        v1 = (v1 - (v1 / ((v1) | 1)))
        _ = v50
    }
    fmt.Println(v1)
    delete(v2, "\n\t")
    _ = len(v2)
    var v51 uint64 = genFn1(255)
    var v52 uint16 = 857
    switch 464 {
    case 127:
        var v53 int = 1
        _ = v53
    case -704:
        v54 := GenS3{F0: (("foo bar" + ("" + "a")) + "foo bar"), F1: 819, F2: false}
        _ = v54
    case -2147483648:
        var v55 int64 = -919
        _ = v55
    }
    switch "" {
    case "":
        fmt.Println(v24, v6)
    case "a":
        v56 := GenS3{F0: "\n\t", F1: int(v1), F2: (!(!(false || (2147483648 == 561))))}
        _ = v56
    default:
        v6 = v6
    }
    for fi57 := 0; fi57 < 18; fi57++ {
        fmt.Println(fi57, v1)
        for fi58 := 0; fi58 < 12; fi58++ {
            v1 += (v1 ^ uint(v51))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi58
        }
        var v59 uint64 = genFn1((v52 & (v52 | (v52 | v52))))
        _ = fi57
        _ = v59
    }
    var v60 int8 = v6
    v52 = 1
    var v61 bool = (!false)
    v62 := GenS4{F0: 42, F1: 2147483647, F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}

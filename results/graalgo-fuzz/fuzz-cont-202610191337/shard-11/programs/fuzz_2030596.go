package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int64; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 string }
func genFn0(p0 string) int {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch p0 {
        case "":
            var v5 int8 = -1
            _ = v5
        case "a":
            var v6 int = (^(fi4))
            _ = v6
        case "hello":
            var v7 uint = 1
            _ = v7
        }
        if false {
            for fi8 := 0; fi8 < 14; fi8++ {
                if (!false) {
                    switch (fi4 * (fi8 * fi4)) {
                    case 4294967295:
                        var v9 uint = 0
                        p0 = (p0 + "hello")
                        _ = v9
                    case 32767:
                        var v10 int32 = 843
                        _ = v10
                    default:
                        if false {
                            v11 := GenS1{F0: 127, F1: 205, F2: -1}
                            switch fi8 {
                            case -422:
                                if true {
                                    v12 := GenS2{F0: (fi4 / ((fi4) | 1))}
                                    p0 = "foo bar"
                                    var v13 int = (fi8 * fi8)
                                    _ = v12
                                    _ = v13
                                }
                            default:
                                fmt.Println(fi8, fi8)
                            }
                            _ = v11
                        }
                    }
                    var v14 float32 = (-2.25 - 1e10)
                    _ = v14
                }
                _ = fi8
            }
            v15 := []int{fi4}
            if false {
                switch "" {
                case "foo bar":
                    fmt.Println(fi4, fi4)
                case "\n\t":
                    var v16 uint = 255
                    _ = v16
                }
                p0 = ""
            }
            _ = v15
        }
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    v17 := GenS3{F0: 9223372036854775807}
    p0 = (("a" + p0) + (p0 + p0))
    _ = p0
    _ = v17
    return -9223372036854775808
}
func genFn1(p0 bool) uint {
    p0 = (false && (p0 || (p0 && true)))
    switch 581 {
    case 128:
        v18 := GenS4{F0: 1}
        _ = v18
    case -1:
        p0 = false
    default:
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: 0}
            v21 := GenS5{F0: -977}
            v22 := GenS3{F0: (fi19 % ((127) | 1))}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v23 := GenS6{F0: 1, F1: 128, F2: 42}
    _ = p0
    _ = v23
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    var v3 uint8 = 128
    v24 := GenS7{F0: 644, F1: 42}
    v25 := GenS4{F0: 809}
    v2 -= v2
    var v26 uint = 2
    v26 -= (v26 & (^(v26)))
    if (1e10 <= 3.14159) {
        v3 = (v3 % ((170) | 1))
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint8 = (v3 / ((v3) | 1))
            var v29 int16 = -32768
            _ = fi27
            _ = v28
            _ = v29
        }
        if true {
            v3 = uint8(v26)
            var v30 int64 = 0
            v31 := v1["a"]
            _ = v30
            _ = v31
        } else {
            if false {
                switch ("a" + ("foo bar" + "\n\t")) {
                case "\n\t":
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v33 := []int{32767, (^(fi32)), fi32, -1}
                        v34 := GenS5{F0: -672}
                        for fi35 := 0; fi35 < 3; fi35++ {
                            switch "a" {
                            case "hello":
                                v36 := GenS3{F0: genFn0("foo bar")}
                                _ = v36
                            case "a":
                                if true {
                                    v37 := GenS8{F0: 1, F1: -32768, F2: 219}
                                    var v38 string = "hello"
                                    _ = v37
                                    _ = v38
                                }
                            case "":
                                switch -1 {
                                case 32768:
                                    if (-1 >= -815) {
                                        if (-2.25 >= 0.1) {
                                            v39 := GenS9{F0: (v3 >> ((v26) & 63)), F1: -285}
                                            _ = v39
                                        } else {
                                            delete(v1, "foo bar")
                                            v2 += 67
                                        }
                                        switch -2147483648 {
                                        case 4294967295:
                                            v40 := GenS3{F0: 0}
                                            _ = v40
                                        case 9223372036854775807:
                                            v41 := GenS5{F0: 1}
                                            _ = v41
                                        }
                                    }
                                case 2147483647:
                                    v2 = v2
                                case 32767:
                                    if (genFn1((17 <= 18446744073709551615)) == v26) {
                                        v42 := GenS4{F0: (v26 % (((v26 ^ uint(v3))) | 1))}
                                        var v43 string = ("a" + "\n\t")
                                        v1[v43] = (-(fi32))
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            }
                            v3 = v3
                            v3 -= 199
                            if len(v33) > 0 {
                                v33[0] = fi35
                            }
                            v44 := GenS10{F0: "hello"}
                            _ = fi35
                            _ = v44
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                default:
                    fmt.Println(v2)
                }
            }
        }
    } else {
        v2 = (v2 / ((0) | 1))
        var v45 int16 = 1
        v46 := []int{}
        _ = v45
        _ = v46
    }
    v26 = (^(v26))
    var v47 int = 32768
    v48 := []int{v47}
    var v49 int64 = -1
    var v50 int = (-(v47))
    if ((^(v26)) >= genFn1((!true))) {
        delete(v1, "")
        v51 := GenS1{F0: 247, F1: (^(v3)), F2: 1}
        var v52 int64 = (v49 | 9223372036854775807)
        _ = v51
        _ = v52
    } else {
        var v53 uint32 = uint32(v3)
        switch ("hello" + (("" + "\n\t") + ("\n\t" + "hello"))) {
        case "":
            if (v53 < (v53 >> (((v26 ^ genFn1(false))) & 63))) {
                if false {
                    for fi54 := 0; fi54 < 6; fi54++ {
                        v26 = (v26 << ((v26) & 63))
                        v48 = append(v48, genFn0("\n\t"))
                        _ = fi54
                    }
                } else {
                    v53 = v53
                    if false {
                        var v55 int8 = (v2 + 76)
                        v56 := GenS8{F0: v49, F1: -473, F2: 1}
                        for fi57 := 0; fi57 < 17; fi57++ {
                            if ((fi57 >> ((42) & 63)) <= genFn0("")) {
                                fmt.Println(v47)
                            } else {
                                v26 -= v26
                            }
                            v48 = append(v48, fi57)
                            _ = fi57
                        }
                        _ = v55
                        _ = v56
                    } else {
                        var v58 float32 = float32((-2.25 - 0.1))
                        _ = v58
                    }
                    if false {
                        var v59 int64 = int64(v2)
                        if (!((65535 > 572) || ((!true) && ("foo bar" > "\n\t")))) {
                            fmt.Println(v3, v26)
                            for fi60 := 0; fi60 < 4; fi60++ {
                                v53 = 1
                                if (0 >= 0) {
                                    var v61 float64 = float64(0.0)
                                    var v62 int8 = (v2 & (v2 * v2))
                                    v2 += v2
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS3{F0: (-(v50))}
                                    _ = v63
                                }
                                _ = fi60
                            }
                        }
                        _ = v59
                    }
                }
                var v64 int64 = v49
                _ = v64
            }
        case "a":
            var v65 uint64 = 4294967295
            _ = v65
        default:
            v2 += (v2 / (((v2 | 42)) | 1))
        }
        _ = v53
    }
    delete(v1, (("a" + "foo bar") + "\n\t"))
    v1["a"] = -1
    for fi66 := 0; fi66 < 13; fi66++ {
        v1[("" + (("" + "") + ("hello" + "hello")))] = (v47 << ((65535) & 63))
        v49 = (v49 + (-(v49)))
        _ = fi66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}

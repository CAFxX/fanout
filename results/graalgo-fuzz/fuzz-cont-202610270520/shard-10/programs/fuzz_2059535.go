package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 string) int {
    v4 := []int{-383}
    var v5 int8 = 1
    if (-1 < 79) {
        _ = len(v4)
    }
    _ = p0
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 bool) int64 {
    if p0 {
        var v6 int16 = 42
        v6 += v6
        v7 := GenS1{F0: 229}
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "":
        p0 = p0
    case "foo bar":
        v8 := GenS1{F0: 0}
        _ = v8
    default:
        var v9 float32 = 3.14159
        _ = v9
    }
    v10 := []int{42, genFn0("foo bar")}
    for fi11 := 0; fi11 < 10; fi11++ {
        switch genFn0("foo bar") {
        case 127:
            p0 = p0
        case -172:
            switch (^(fi11)) {
            case 127:
                fmt.Println(fi11)
            case -402:
                if p0 {
                    if false {
                        var v12 uint8 = 93
                        _ = v12
                    }
                } else {
                    var v13 uint = 42
                    _ = v13
                }
            }
        default:
            var v14 int8 = 5
            _ = v14
        }
        switch (fi11 - (fi11 * fi11)) {
        case 128:
            switch "hello" {
            case "\n\t":
                if (0 > 2) {
                    switch ("\n\t" + "a") {
                    case "a":
                        v15 := GenS2{F0: 2, F1: (!(0 >= 42)), F2: 569}
                        _ = v15
                    case "\n\t":
                        fmt.Println(p0, p0)
                    default:
                        for fi16 := 0; fi16 < 9; fi16++ {
                            if p0 {
                                var v17 uint64 = 42
                                var v18 float64 = ((0.1 + (3.14159 - 1e10)) + (3.14159 - (((0.1 - (((0.1 / 1.0) - float64(1.5)) * 1e10)) / 3.14159) * 1.5)))
                                _ = v17
                                _ = v18
                            }
                            _ = fi16
                        }
                    }
                    p0 = (18446744073709551615 <= 42)
                }
            case "":
                for fi19 := 0; fi19 < 5; fi19++ {
                    if true {
                        var v20 int64 = 439
                        v21 := GenS1{F0: 4294967295}
                        var v22 float32 = 1.5
                        v10 = append(v10, -9223372036854775808)
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                }
            case "foo bar":
                p0 = ("hello" < ("" + "foo bar"))
            default:
                for fi23 := 0; fi23 < 14; fi23++ {
                    v10 = append(v10, fi23)
                    _ = fi23
                }
            }
        case 42:
            fmt.Println(fi11, fi11)
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            p0 = false
            _ = fi24
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int8 = 127
    var v2 string = ("a" + "foo bar")
    var v3 int = 296
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := []int{1}
            v1 = v1
            v2 = (v2 + v2)
            _ = fi26
            _ = v27
        }
        v3 -= -565
        _ = fi25
    }
    v28 := GenS2{F0: 18446744073709551615, F1: false, F2: 1}
    v1 = v1
    fmt.Println(v3, v3)
    fmt.Println(v3, v2)
    v3 = v3
    switch "" {
    case "\n\t":
        for fi29 := 0; fi29 < 5; fi29++ {
            if ((v1 + (-(v1))) < v1) {
                var v30 bool = ("hello" != ((v2 + "hello") + v2))
                switch v2 {
                case "foo bar":
                    v31 := []int{fi29}
                    _ = v31
                case "a":
                    for fi32 := 0; fi32 < 16; fi32++ {
                        v33 := GenS1{F0: 2}
                        _ = fi32
                        _ = v33
                    }
                default:
                    v3 += fi29
                }
                v34 := []int{v3, (fi29 | (fi29 ^ (fi29 | (fi29 << ((65535) & 63))))), (^(v3)), (v3 / ((genFn0(v2)) | 1))}
                _ = v30
                _ = v34
            } else {
                for fi35 := 0; fi35 < 12; fi35++ {
                    switch v2 {
                    case "hello":
                        var v36 int32 = 42
                        _ = v36
                    case "\n\t":
                        v1 += v1
                    }
                    var v37 float32 = float32(1.5)
                    _ = fi35
                    _ = v37
                }
                var v38 uint16 = 1
                switch (((("\n\t" + v2) + v2) + v2) + v2) {
                case "\n\t":
                    v39 := GenS1{F0: 42}
                    _ = v39
                case "":
                    for fi40 := 0; fi40 < 15; fi40++ {
                        var v41 int32 = 1
                        switch v2 {
                        case "a":
                            v42 := []int{(fi29 + 32768), v3, (v3 >> ((1) & 63))}
                            if (-2.25 == -2.25) {
                                v41 += 42
                            } else {
                                v41 += (v41 / ((v41) | 1))
                                var v43 uint16 = 252
                                v44 := GenS3{F0: genFn1(false)}
                                _ = v43
                                _ = v44
                            }
                            v1 += 1
                            _ = v42
                        case "foo bar":
                            if (0 > 0) {
                                v2 = (v2 + v2)
                                v45 := []int{-9223372036854775808}
                                var v46 int16 = 36
                                _ = v45
                                _ = v46
                            } else {
                                v47 := []int{}
                                if false {
                                    v47 = append(v47, genFn0(v2))
                                    fmt.Println(v38, v2, v3)
                                    if (!((v2 > v2) && false)) {
                                        v1 += v1
                                        v38 = 42
                                    } else {
                                        v41 = v41
                                        var v48 int16 = 32767
                                        var v49 int16 = 788
                                        v3 += (fi40 % ((fi29) | 1))
                                        _ = v48
                                        _ = v49
                                    }
                                    v50 := GenS4{F0: ("" + v2), F1: 42, F2: 18446744073709551615}
                                    if (!(((((0.0 * 3.14159) * float64(((((1e10 - 3.14159) - (0.1 - -2.25)) - float32((0.0 * 3.14159))) - -2.25))) * (float64(1e10) * float64(1.5))) * 0.1) != 0.0)) {
                                        switch (((("\n\t" + "foo bar") + v2) + (v2 + v2)) + "a") {
                                        case "\n\t":
                                            var v51 int32 = v41
                                            _ = v51
                                        case "":
                                            v1 -= v1
                                        default:
                                            v1 = v1
                                        }
                                        var v52 float32 = 1e10
                                        _ = v52
                                    }
                                    if (42 <= -261) {
                                        var v53 int64 = 1
                                        _ = len(v47)
                                        var v54 uint16 = (v38 >> ((0) & 63))
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = v50
                                }
                                v55 := GenS1{F0: 0}
                                _ = v47
                                _ = v55
                            }
                        }
                        v38 = (v38 << ((1) & 63))
                        _ = fi40
                        _ = v41
                    }
                case "a":
                    var v56 uint = 667
                    _ = v56
                default:
                    v1 = v1
                }
                _ = v38
            }
            v2 = v2
            _ = fi29
        }
    default:
        var v57 float64 = (-2.25 * (float64(-2.25) - 1e10))
        _ = v57
    }
    var v58 int64 = 0
    v3 -= 128
    fmt.Println(v3)
    if false {
        v3 -= -2147483648
    } else {
        v59 := GenS5{F0: 252}
        _ = v59
    }
    v2 = ("\n\t" + "\n\t")
    fmt.Println(v1, v1)
    v58 = v58
    var v60 int64 = -880
    var v61 int32 = 2147483647
    v62 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v58
    _ = v60
    _ = v61
    _ = v62
}

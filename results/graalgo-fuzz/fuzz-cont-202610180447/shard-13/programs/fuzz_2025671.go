package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 int16) uint32 {
    var v5 float64 = (((((1e10 / 1e10) + (0.0 - float64(float32(float64((3.14159 * 0.1)))))) - (-2.25 / -2.25)) * ((0.0 + (((0.1 - ((1.5 + -2.25) * 1e10)) * 1.5) + (float64(float32(1e10)) * 1.5))) + float64(float32(1.5)))) + 1e10)
    switch -2147483648 {
    case 1:
        fmt.Println(v5, p0)
    }
    v5 = 1e10
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = 118
        for fi8 := 0; fi8 < 20; fi8++ {
            v7 = (^(v7))
            v9 := GenS1{F0: "foo bar", F1: 65535, F2: 127}
            if (!((1.5 - -2.25) >= 0.1)) {
                v5 = (v5 - 0.1)
                fmt.Println(v5, fi6)
            }
            p0 = 103
            _ = fi8
            _ = v9
        }
        var v10 int64 = -414
        _ = fi6
        _ = v10
        _ = v7
    }
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 uint16) int32 {
    if (-1 != 128) {
        v11 := []int{}
        p0 -= p0
        _ = v11
    } else {
        if (0.0 != 1.5) {
            var v12 int = -41
            _ = v12
        } else {
            if false {
                if true {
                    p0 = p0
                    p0 = 65535
                    var v13 uint = 650
                    _ = v13
                } else {
                    p0 = 65535
                    var v14 int16 = 32767
                    v15 := GenS1{F0: ("a" + "\n\t"), F1: 18446744073709551615, F2: 0}
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := []int{(fi16 & (-(fi16))), fi16, -162}
                    for fi18 := 0; fi18 < 12; fi18++ {
                        var v19 string = ("a" + ("foo bar" + (("foo bar" + "a") + ("\n\t" + "\n\t"))))
                        v20 := GenS1{F0: v19, F1: 255, F2: -69}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    p0 = p0
                    _ = fi16
                    _ = v17
                }
            }
            var v21 int8 = 80
            switch 4294967295 {
            case 853:
                fmt.Println(p0, p0)
            case 128:
                for fi22 := 0; fi22 < 17; fi22++ {
                    if true {
                        switch ((("hello" + "a") + "foo bar") + ("hello" + "foo bar")) {
                        case "hello":
                            fmt.Println(p0, fi22, p0)
                        case "a":
                            p0 += (^(p0))
                        default:
                            if (-2.25 < (((1.5 / 1.5) * 1.5) / (1.5 / (1.5 / 1e10)))) {
                                var v23 int32 = 42
                                v23 = v23
                                _ = v23
                            } else {
                                v24 := GenS1{F0: "", F1: 18446744073709551615, F2: (v21 & v21)}
                                var v25 string = ("foo bar" + ("" + ("foo bar" + ("hello" + "foo bar"))))
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    if true {
                                        v27 := []int{-589, (fi22 - 32767)}
                                        if (false || (0 >= 1)) {
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                        }
                                        switch fi22 {
                                        case -191:
                                            fmt.Println(v21, fi26, v25)
                                        default:
                                            var v28 uint16 = p0
                                            _ = v28
                                        }
                                        var v29 int8 = (v21 * 100)
                                        v25 = v25
                                        _ = v27
                                        _ = v29
                                    }
                                    _ = fi26
                                }
                                _ = v24
                                _ = v25
                            }
                        }
                        var v30 uint64 = 0
                        _ = v30
                    }
                    if false {
                        var v31 uint16 = p0
                        _ = v31
                    }
                    _ = fi22
                }
            case 127:
                v21 = -94
            default:
                for fi32 := 0; fi32 < 18; fi32++ {
                    v33 := GenS1{F0: "", F1: 18446744073709551615, F2: v21}
                    _ = fi32
                    _ = v33
                }
            }
            _ = v21
        }
    }
    var v34 int32 = 42
    _ = p0
    _ = v34
    return (v34 | v34)
}

func main() {
    v1 := []string{"a", "hello", "", "foo bar"}
    var v2 int16 = 42
    var v3 uint32 = 1
    var v4 uint64 = 1
    switch ("hello" + "foo bar") {
    case "\n\t":
        v35 := GenS2{F0: genFn0((v2 & (v2 | (v2 + 0)))), F1: 127, F2: int8(v4)}
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 int16 = int16(v3)
        v3 = (v3 | v3)
        _ = fi36
        _ = v37
    }
    v38 := GenS1{F0: "foo bar", F1: uint(v3), F2: -63}
    switch ((("" + "") + "foo bar") + "foo bar") {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v39 bool = true
        _ = v39
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if (false && (32767 < (v2 * -715))) {
        for fi40 := 0; fi40 < 15; fi40++ {
            v41 := GenS2{F0: v3, F1: 99, F2: 42}
            _ = fi40
            _ = v41
        }
        fmt.Println(v4, v4)
    } else {
        if false {
            v42 := GenS2{F0: (v3 >> ((2) & 63)), F1: 127, F2: -1}
            v4 += (v4 + (v4 >> ((18446744073709551615) & 63)))
            switch "hello" {
            case "hello":
                v43 := GenS2{F0: v3, F1: 42, F2: int8(v2)}
                _ = v43
            default:
                v1 = append(v1, "hello")
            }
            _ = v42
        } else {
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 int32 = -25
                _ = fi44
                _ = v45
            }
        }
    }
    var v46 int32 = -2147483648
    if (("foo bar" + "\n\t") >= "a") {
        var v47 int16 = (v2 >> ((0) & 63))
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 float64 = 0.0
            if (v46 >= (v46 ^ v46)) {
                var v50 uint64 = v4
                var v51 int = (fi48 >> ((4294967295) & 63))
                _ = v50
                _ = v51
            } else {
                if (238 > 0) {
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v4 = v4
                        _ = fi52
                    }
                }
                for fi53 := 0; fi53 < 6; fi53++ {
                    switch "" {
                    case "":
                        var v54 int = fi53
                        _ = v54
                    case "\n\t":
                        v1 = append(v1, ("hello" + ("" + ("foo bar" + ("\n\t" + ("\n\t" + ("" + "foo bar")))))))
                    case "foo bar":
                        if false {
                            v55 := GenS3{F0: 127, F1: "a", F2: 9223372036854775807}
                            v47 += (v2 << ((972) & 63))
                            _ = v55
                        } else {
                            v56 := GenS4{F0: -459, F1: v4, F2: 1}
                            _ = v56
                        }
                    default:
                        switch (^(fi48)) {
                        case -414:
                            v57 := GenS3{F0: -128, F1: ("a" + "a"), F2: -856}
                            _ = v57
                        case -1:
                            fmt.Println(fi48)
                        }
                    }
                    v1 = append(v1, "\n\t")
                    v2 = v47
                    v58 := GenS5{F0: uint32(v4), F1: fi53, F2: 2}
                    var v59 int16 = 32767
                    _ = fi53
                    _ = v58
                    _ = v59
                }
                fmt.Println(v46)
            }
            _ = fi48
            _ = v49
        }
        if false {
            var v60 int64 = -40
            var v61 int16 = (v47 / ((int16(v4)) | 1))
            _ = v60
            _ = v61
        }
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v46
}

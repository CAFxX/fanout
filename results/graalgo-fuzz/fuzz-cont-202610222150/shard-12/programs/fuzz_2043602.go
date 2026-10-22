package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
func genFn0() uint {
    if (0 == 100) {
        var v6 int8 = 1
        v6 -= 100
        _ = v6
    }
    var v7 uint32 = 835
    var v8 float64 = 1.5
    fmt.Println(v7, v7)
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 bool) int64 {
    if (!p0) {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = "hello"
            p0 = (((v10 + v10) + v10) <= v10)
            _ = fi9
            _ = v10
        }
    } else {
        p0 = (false && (((1e10 - 0.1) * (0.0 / 3.14159)) > (1e10 - 1e10)))
        switch 2147483647 {
        case 32767:
            switch "a" {
            case "foo bar":
                fmt.Println(p0)
            default:
                p0 = (p0 || (900 >= 0))
            }
        case -2147483648:
            var v11 uint = 65535
            _ = v11
        }
        switch ("a" + "foo bar") {
        case "foo bar":
            v12 := GenS1{F0: 1, F1: 42}
            _ = v12
        case "\n\t":
            switch (("foo bar" + ((("\n\t" + "a") + ("\n\t" + "")) + ("hello" + ((("foo bar" + "a") + "hello") + "hello")))) + "") {
            case "a":
                v13 := []int{575}
                _ = v13
            }
        default:
            v14 := []int{}
            _ = v14
        }
    }
    if (342 == 32767) {
        if (1 <= 18446744073709551615) {
            v15 := []int{-9223372036854775808, 127, 0, 513}
            if (p0 && (!p0)) {
                var v16 int32 = 84
                _ = v16
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 uint32 = 2147483648
                    _ = fi17
                    _ = v18
                }
                var v19 int = 9223372036854775807
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := GenS1{F0: 42, F1: 2}
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v19 -= fi22
                        v23 := GenS1{F0: -11, F1: 42}
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            _ = v15
        }
        p0 = false
        v24 := []int{-9223372036854775808, -994, 4294967295, -648}
        _ = v24
    }
    v25 := []int{790, 42, 2147483647, 817}
    for fi26 := 0; fi26 < 13; fi26++ {
        switch "" {
        case "foo bar":
            p0 = (933 <= 2)
        case "":
            fmt.Println(fi26, fi26)
        }
        _ = fi26
    }
    p0 = false
    _ = p0
    _ = v25
    return 893
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int = 42
    v4 := []uint32{136, 2, 89, 42}
    v5 := make(map[int]bool)
    v3 = (^(v3))
    fmt.Println(v3)
    var v27 uint64 = 256
    if (-61 >= 127) {
        _ = len(v4)
        v28 := GenS2{F0: 641, F1: (^(v27))}
        if ((100 > 100) && (!false)) {
            if (!false) {
                fmt.Println(v27)
            } else {
                for fi29 := 0; fi29 < 12; fi29++ {
                    fmt.Println(fi29, v27, v3)
                    _ = fi29
                }
                var v30 uint = 65535
                if false {
                    if true {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            for fi32 := 0; fi32 < 10; fi32++ {
                                var v33 int8 = -128
                                _ = fi32
                                _ = v33
                            }
                            _ = fi31
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                } else {
                    var v34 int64 = 1
                    _ = v34
                }
                _ = v30
            }
            v3 += v3
            if false {
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 uint32 = 65535
                    if (-9223372036854775808 < 0) {
                        v4 = append(v4, (v36 | v36))
                        var v37 uint32 = 0
                        _ = v37
                    } else {
                        v5[842] = true
                    }
                    _ = fi35
                    _ = v36
                }
                v3 = v3
            } else {
                _ = len(v4)
                if true {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        fmt.Println(v3, v3, v27)
                        if len(v4) > 0 {
                            v4[0] = 2
                        }
                        _ = fi38
                    }
                    if (0.1 != 1.5) {
                        v39 := GenS3{F0: 42, F1: (v3 * v3), F2: v27}
                        _ = v39
                    }
                }
                var v40 string = "\n\t"
                _ = v40
            }
        }
        _ = v28
    } else {
        v41 := v1["a"]
        _ = v41
    }
    var v42 int = (v3 | 32767)
    var v43 int64 = 0
    for fi44 := 0; fi44 < 16; fi44++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v45 := GenS3{F0: 42, F1: (v42 ^ fi44), F2: (^(v27))}
        var v46 int32 = -211
        _ = fi44
        _ = v45
        _ = v46
    }
    _ = len(v4)
    var v47 bool = (1.5 >= 1e10)
    v42 -= (^(v42))
    var v48 uint = 2
    for fi49 := 0; fi49 < 17; fi49++ {
        v47 = ((4294967295 != 1) || v47)
        v43 = v43
        v48 = v48
        _ = fi49
    }
    switch v42 {
    case -2147483648:
        v50 := GenS3{F0: 1, F1: v42, F2: 245}
        _ = v50
    case 4294967295:
        var v51 uint16 = 737
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v5
}

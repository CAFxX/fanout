package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int16 }
type GenS4 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 uint64, p1 int32, p2 string) int64 {
    switch p2 {
    case "foo bar":
        switch p2 {
        case "a":
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int16 = 42
                for fi9 := 0; fi9 < 3; fi9++ {
                    p0 -= 1
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
        }
    case "":
        fmt.Println(p1, p1)
    case "a":
        v10 := []int{2147483647, -9223372036854775808, -9223372036854775808, 32768}
        _ = v10
    }
    var v11 int64 = 9223372036854775807
    v12 := []int{32768, 32767, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (v11 - v11)
}
func genFn1() string {
    v13 := GenS1{F0: 128, F1: 4294967295, F2: 2147483647}
    var v14 int64 = -9223372036854775808
    v15 := []int{}
    var v16 bool = false
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS2{F0: 235, F1: 255, F2: -1}
        _ = fi17
        _ = v18
    }
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return ("a" + "hello")
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 255
    v3 := make(map[int]string)
    var v4 int8 = -128
    var v5 uint8 = v2
    v6 := []int{2147483647, 1}
    for fi19 := 0; fi19 < 11; fi19++ {
        v5 -= 127
        var v20 int32 = 0
        if true {
            var v21 uint = 788
            v6 = append(v6, fi19)
            v22 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: v20}
            _ = v21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v1[("" + "a")] = 42
    v2 -= v5
    if (296 == 255) {
        v1["hello"] = 4294967295
        switch ("" + "foo bar") {
        case "foo bar":
            for fi23 := 0; fi23 < 18; fi23++ {
                _ = len(v3)
                _ = fi23
            }
        case "\n\t":
            v2 -= 217
        default:
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 float32 = -2.25
                _ = fi24
                _ = v25
            }
        }
        var v26 float32 = 3.14159
        _ = v26
    } else {
        v5 += (v5 | (^(v2)))
    }
    if true {
        for fi27 := 0; fi27 < 15; fi27++ {
            v3[fi27] = ("foo bar" + "a")
            _ = fi27
        }
        v28 := GenS2{F0: 2147483647, F1: v5, F2: (v4 / (((v4 + (v4 + -128))) | 1))}
        var v29 uint = 0
        _ = v28
        _ = v29
    } else {
        v4 = 100
        v2 += v2
    }
    v30 := GenS1{F0: 537, F1: 1, F2: 1}
    for fi31 := 0; fi31 < 6; fi31++ {
        switch fi31 {
        case 1:
            var v32 uint8 = 42
            _ = v32
        case 24:
            fmt.Println(v4, v2, v5)
        case 128:
            v2 = 0
        default:
            fmt.Println(fi31, fi31)
        }
        _ = fi31
    }
    v33 := GenS2{F0: 0, F1: v2, F2: v4}
    if (!true) {
        v34 := GenS1{F0: 32768, F1: 1, F2: -2147483648}
        _ = v34
    }
    var v35 int = 9223372036854775807
    switch v35 {
    case 971:
        v36 := GenS1{F0: 128, F1: 1, F2: 42}
        _ = v36
    case -2147483648:
        var v37 float64 = 3.14159
        _ = v37
    case 919:
        for fi38 := 0; fi38 < 2; fi38++ {
            v1["a"] = (v35 / ((779) | 1))
            _ = fi38
        }
    }
    switch "\n\t" {
    case "a":
        var v39 string = "\n\t"
        _ = v39
    case "foo bar":
        for fi40 := 0; fi40 < 14; fi40++ {
            switch ("hello" + ((((("" + "a") + "hello") + "") + "\n\t") + ("foo bar" + "foo bar"))) {
            case "a":
                v6 = append(v6, (v35 * 4294967295))
            case "\n\t":
                var v41 int16 = 1
                _ = v41
            }
            switch "hello" {
            case "foo bar":
                v42 := GenS2{F0: 42, F1: (v2 | v2), F2: (v4 - (v4 % ((42) | 1)))}
                _ = v42
            case "\n\t":
                var v43 uint = 65535
                _ = v43
            case "hello":
                if (false && false) {
                    v35 = (-(fi40))
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            if ((42 < 0) || (0.1 != (0.1 * (3.14159 - ((((0.1 - 3.14159) + -2.25) * 1e10) / 3.14159))))) {
                                fmt.Println(v4, v5, v35)
                                v35 -= v35
                                fmt.Println(v4, fi40, v5)
                            }
                        }
                        v44 := GenS3{F0: 4294967295, F1: -77, F2: -32768}
                        v45 := GenS4{F0: (fi40 | 4294967295), F1: 2147483647, F2: 18446744073709551615}
                        _ = v44
                        _ = v45
                    }
                    delete(v3, 0)
                } else {
                    var v46 uint32 = 0
                    var v47 float32 = 0.1
                    _ = v46
                    _ = v47
                }
            }
            fmt.Println(v2)
            _ = fi40
        }
    case "hello":
        v6 = append(v6, (v35 & -9223372036854775808))
    }
    if true {
        for fi48 := 0; fi48 < 17; fi48++ {
            delete(v1, "hello")
            _ = fi48
        }
        for fi49 := 0; fi49 < 18; fi49++ {
            if false {
                v2 = 133
                switch (((("\n\t" + ("hello" + (("a" + ("" + "")) + ("\n\t" + ("" + "foo bar"))))) + "") + "") + ("a" + "foo bar")) {
                case "a":
                    delete(v3, (v35 >> ((18446744073709551615) & 63)))
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "\n\t":
                    if (32767 != 42) {
                        for fi50 := 0; fi50 < 19; fi50++ {
                            var v51 int = fi50
                            _ = fi50
                            _ = v51
                        }
                    } else {
                        fmt.Println(v35)
                        var v52 int8 = (v4 / ((int8(v35)) | 1))
                        _ = v52
                    }
                }
                if false {
                    var v53 uint = 739
                    switch v35 {
                    case -9223372036854775808:
                        v5 = (v5 | uint8(v53))
                    }
                    var v54 uint = (^(v53))
                    _ = v53
                    _ = v54
                }
            }
            if ((!false) || (663 > 2)) {
                for fi55 := 0; fi55 < 8; fi55++ {
                    switch 127 {
                    case 32768:
                        var v56 string = (("foo bar" + ("hello" + "foo bar")) + "hello")
                        _ = v56
                    case -9223372036854775808:
                        v5 = v2
                    case -629:
                        delete(v1, "\n\t")
                    }
                    _ = fi55
                }
                v4 -= (v4 % ((1) | 1))
                v57 := GenS2{F0: 0, F1: (v5 << ((255) & 63)), F2: v4}
                _ = v57
            } else {
                for fi58 := 0; fi58 < 7; fi58++ {
                    fmt.Println(v5, fi58)
                    switch (fi58 * (fi49 << ((18446744073709551615) & 63))) {
                    case -817:
                        fmt.Println(v35)
                    }
                    fmt.Println(v4)
                    _ = fi58
                }
                switch "a" {
                case "hello":
                    var v59 float32 = ((float32((0.0 - 0.0)) - ((-2.25 / -2.25) - (3.14159 / 0.1))) - 0.0)
                    _ = v59
                }
            }
            var v60 int = (fi49 << ((2) & 63))
            _ = fi49
            _ = v60
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if len(v6) > 0 {
            v6[0] = v35
        }
        var v61 int64 = -1
        _ = v61
    }
    fmt.Println(v4)
    var v62 float64 = ((0.1 * 0.1) / 0.1)
    v5 = v5
    var v63 int32 = -657
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}

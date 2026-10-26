package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint16; F2 int }
type GenS5 struct { F0 uint64; F1 int8; F2 string }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int32; F1 uint8; F2 int16 }
type GenS9 struct { F0 int; F1 int32; F2 int64 }
func genFn0() int64 {
    var v5 uint64 = 18446744073709551615
    if (1 != 42) {
        fmt.Println(v5)
        v5 = 753
        fmt.Println(v5)
    } else {
        fmt.Println(v5)
        switch 9223372036854775807 {
        case 2147483647:
            var v6 uint32 = 2147483647
            _ = v6
        case 128:
            fmt.Println(v5)
        case 86:
            v7 := []int{}
            _ = v7
        default:
            v8 := GenS1{F0: 255, F1: 65535, F2: 42}
            _ = v8
        }
        v9 := []int{0, 1, 2147483647, -328}
        _ = v9
    }
    switch 4294967295 {
    case 42:
        if (0.0 > (0.0 * 3.14159)) {
            var v10 bool = false
            switch 4294967295 {
            case -306:
                v11 := GenS2{F0: 32767, F1: 2147483647}
                _ = v11
            case 1:
                fmt.Println(v10, v5)
            case 70:
                v5 = 638
            }
            if v10 {
                var v12 uint = 18446744073709551615
                switch (("foo bar" + ("\n\t" + "")) + "a") {
                case "\n\t":
                    if v10 {
                        v13 := GenS2{F0: 477, F1: 2147483647}
                        v14 := GenS2{F0: 634, F1: -1}
                        _ = v13
                        _ = v14
                    } else {
                        v10 = ((("" + ("hello" + "hello")) + "") <= "\n\t")
                    }
                case "hello":
                    v12 = 0
                case "":
                    v12 = 2
                }
                var v15 int64 = -642
                _ = v12
                _ = v15
            } else {
                v16 := GenS1{F0: 1, F1: 42, F2: 580}
                _ = v16
            }
            _ = v10
        } else {
            if ((42 == 2147483647) && false) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := GenS3{F0: 0}
                    fmt.Println(fi17, fi17)
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 float32 = (1.5 / 3.14159)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                var v21 uint8 = 1
                for fi22 := 0; fi22 < 13; fi22++ {
                    if ("\n\t" != "a") {
                        v23 := GenS4{F0: 975, F1: 918, F2: (fi22 + -563)}
                        _ = v23
                    }
                    v24 := []int{fi22, int(v5), fi22, (fi22 & 290)}
                    _ = fi22
                    _ = v24
                }
                _ = v21
            } else {
                v25 := GenS5{F0: v5, F1: 42, F2: "foo bar"}
                if (1 == -2147483648) {
                    v26 := GenS6{F0: 813, F1: -9223372036854775808}
                    var v27 bool = false
                    v28 := []int{-2147483648, 128, 1, 42}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v29 string = ("" + "")
                        _ = v29
                    }
                }
                v30 := GenS7{F0: 42, F1: 42}
                _ = v25
                _ = v30
            }
            if true {
                v31 := []int{128, 1}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v5 = (^(v5))
                    if true {
                        v33 := []int{}
                        _ = v33
                    }
                    _ = fi32
                }
            }
        }
    }
    v34 := []int{32767}
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    _ = v34
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = -128
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    if true {
        switch 841 {
        case 952:
            if (!(v1 > v1)) {
                var v35 bool = (!true)
                v2 += v2
                _ = v35
            }
        }
    }
    v2 = (v2 ^ 127)
    var v36 int = -674
    v37 := []int{int(v1)}
    if false {
        for fi38 := 0; fi38 < 4; fi38++ {
            delete(v3, ("" + "\n\t"))
            if (-2.25 > 3.14159) {
                if len(v37) >= 2 {
                    v37 = v37[1:len(v37)]
                }
            }
            _ = fi38
        }
        v1 = v1
    }
    if (false && true) {
        switch "hello" {
        case "":
            v2 = v2
        case "a":
            v39 := GenS3{F0: 1}
            _ = v39
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            if (true || true) {
                for fi41 := 0; fi41 < 10; fi41++ {
                    var v42 int8 = -128
                    var v43 uint32 = 2
                    var v44 uint = 18446744073709551615
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                if true {
                    var v45 float32 = (1e10 - float32(1.5))
                    var v46 string = "hello"
                    _ = v45
                    _ = v46
                }
                var v47 int64 = 0
                _ = v47
            }
            var v48 int8 = (v2 - -1)
            for fi49 := 0; fi49 < 14; fi49++ {
                var v50 uint = 0
                v2 = -128
                _ = fi49
                _ = v50
            }
            _ = fi40
            _ = v48
        }
    } else {
        for fi51 := 0; fi51 < 11; fi51++ {
            for fi52 := 0; fi52 < 7; fi52++ {
                if true {
                    if (0.0 > (0.1 - -2.25)) {
                        v1 -= 32
                        if true {
                            v53 := GenS8{F0: -930, F1: 42, F2: 1}
                            delete(v4, -582)
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                            _ = v53
                        } else {
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        }
                        v1 = v1
                        if true {
                            var v54 float64 = ((3.14159 * 0.0) - (1e10 - 3.14159))
                            v4[fi52] = -2147483648
                            var v55 uint64 = 2
                            _ = v54
                            _ = v55
                        } else {
                            v36 = -861
                        }
                    } else {
                        v56 := GenS3{F0: 668}
                        _ = v56
                    }
                    v57 := GenS3{F0: 44}
                    switch (("\n\t" + "hello") + "\n\t") {
                    case "hello":
                        if len(v37) > 0 {
                            v37[0] = 1
                        }
                    case "\n\t":
                        v2 += v2
                    case "":
                        v36 -= (fi52 | (fi52 >> ((916) & 63)))
                    default:
                        fmt.Println(v2, v36)
                    }
                    _ = v57
                }
                var v58 uint8 = 255
                v1 += 937
                _ = fi52
                _ = v58
            }
            fmt.Println(fi51)
            _ = fi51
        }
        for fi59 := 0; fi59 < 6; fi59++ {
            v60 := GenS9{F0: (fi59 >> ((163) & 63)), F1: -1, F2: -90}
            _ = fi59
            _ = v60
        }
        v1 = (v1 & (v1 << ((65535) & 63)))
    }
    for fi61 := 0; fi61 < 7; fi61++ {
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
        _ = fi61
    }
    var v62 int8 = 42
    switch "" {
    case "\n\t":
        switch ("foo bar" + ("" + "hello")) {
        case "foo bar":
            fmt.Println(v2, v1)
        case "hello":
            v36 = (v36 ^ v36)
        case "a":
            fmt.Println(v36)
        default:
            var v63 float32 = 1.5
            _ = v63
        }
    default:
        v64 := GenS4{F0: -2147483648, F1: (v1 * (v1 - (v1 ^ (v1 << ((2) & 63))))), F2: (v36 & 724)}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v62
}

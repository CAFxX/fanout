package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 string }
func genFn0(p0 int32, p1 int8) uint16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "\n\t" {
        case "hello":
            if (!false) {
                var v7 bool = false
                v8 := GenS1{F0: -32768}
                _ = v7
                _ = v8
            } else {
                fmt.Println(fi6, p1)
                p0 = -1
            }
        case "\n\t":
            var v9 float64 = 0.0
            _ = v9
        }
        v10 := GenS2{F0: 338, F1: 42, F2: 9223372036854775807}
        v11 := GenS3{F0: p0}
        _ = fi6
        _ = v10
        _ = v11
    }
    v12 := []int{127, 32768, 127}
    switch "a" {
    case "a":
        p1 += 42
    }
    if false {
        var v13 int64 = 473
        _ = v13
    } else {
        var v14 float64 = ((3.14159 - (1.5 + float64((1e10 + float32(-2.25))))) + (0.1 * ((1e10 + (-2.25 / 1.0)) * (0.1 / -2.25))))
        _ = v14
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 float64 = 1.5
    v3 := []string{"a", ""}
    v4 := []string{"foo bar", (("a" + (("hello" + ("hello" + "\n\t")) + "foo bar")) + "foo bar"), (("a" + "hello") + "a")}
    var v5 uint8 = 1
    v15 := GenS1{F0: 1}
    v16 := GenS3{F0: 1}
    v2 = -2.25
    for fi17 := 0; fi17 < 17; fi17++ {
        if ((!(!true)) && ((3.14159 * float32(v2)) > float32((float64(1e10) + v2)))) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                if true {
                    if true {
                        v5 = (v5 | v5)
                    }
                } else {
                    if (1e10 != 0.1) {
                        var v18 int64 = -553
                        var v19 float32 = 3.14159
                        v1 -= 42
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(v2, v1)
                        v2 = v2
                        var v20 int16 = 32767
                        _ = v20
                    }
                }
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                v2 = v2
                var v22 uint16 = 255
                v23 := GenS4{F0: 2, F1: "foo bar"}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        fmt.Println(v2)
        _ = fi17
    }
    if true {
        v24 := GenS4{F0: 18446744073709551615, F1: "a"}
        var v25 uint = 255
        _ = v24
        _ = v25
    }
    v26 := GenS5{F0: 416, F1: 42}
    if (!(598 > 682)) {
        v5 = v5
        var v27 int8 = (v1 * 42)
        if false {
            v5 += v5
            var v28 int32 = 2147483647
            _ = v28
        } else {
            if (false || false) {
                var v29 uint = 2
                v3 = append(v3, "hello")
                v30 := GenS2{F0: genFn0(-1, 100), F1: 1, F2: 1}
                _ = v29
                _ = v30
            } else {
                v31 := GenS2{F0: 2, F1: 0, F2: 753}
                _ = v31
            }
            var v32 uint = 18446744073709551615
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v32
        }
        _ = v27
    }
    v2 = float64(-2.25)
    v33 := GenS5{F0: 1, F1: -861}
    v5 = (^(v5))
    fmt.Println(v2, v5, v1)
    var v34 int8 = 42
    if (!false) {
        v35 := GenS6{F0: 1, F1: 18446744073709551615}
        switch "foo bar" {
        case "":
            var v36 uint64 = 1
            _ = v36
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = "hello"
            }
        default:
            if false {
                if (!(-2147483648 != 42)) {
                    var v37 bool = (2 <= 42)
                    switch "\n\t" {
                    case "a":
                        for fi38 := 0; fi38 < 10; fi38++ {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                v40 := GenS2{F0: genFn0(2147483647, v1), F1: genFn0(0, (v1 / ((89) | 1))), F2: 649}
                                _ = fi39
                                _ = v40
                            }
                            _ = fi38
                        }
                    case "\n\t":
                        _ = len(v3)
                    }
                    _ = v37
                }
                if ("a" > "hello") {
                    for fi41 := 0; fi41 < 13; fi41++ {
                        v4 = append(v4, (("hello" + (("foo bar" + "hello") + "foo bar")) + ""))
                        v2 = float64(1.5)
                        _ = fi41
                    }
                    switch -9223372036854775808 {
                    case 1:
                        v5 = (v5 / ((v5) | 1))
                    default:
                        if (-350 == 1) {
                            for fi42 := 0; fi42 < 2; fi42++ {
                                fmt.Println(v1, v5)
                                if len(v4) > 0 {
                                    v4[0] = (("a" + ("a" + (("" + ("a" + "\n\t")) + "a"))) + "hello")
                                }
                                var v43 string = "a"
                                _ = fi42
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 10; fi44++ {
                                for fi45 := 0; fi45 < 4; fi45++ {
                                    v46 := GenS6{F0: 170, F1: 43}
                                    switch ("a" + ("\n\t" + "a")) {
                                    case "":
                                        v34 = v1
                                    case "\n\t":
                                        v4 = append(v4, ("a" + ""))
                                    default:
                                        var v47 int = fi45
                                        _ = v47
                                    }
                                    _ = len(v4)
                                    _ = fi45
                                    _ = v46
                                }
                                v48 := GenS5{F0: 4294967295, F1: 472}
                                _ = fi44
                                _ = v48
                            }
                            switch (("hello" + "hello") + "") {
                            case "foo bar":
                                v5 += v5
                            case "a":
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    v4 = append(v4, "")
                                    _ = fi49
                                }
                            case "\n\t":
                                var v50 float32 = 3.14159
                                _ = v50
                            default:
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    var v52 int = fi51
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        } else {
                            var v53 uint64 = 526
                            v54 := GenS5{F0: 999, F1: 9223372036854775807}
                            _ = v53
                            _ = v54
                        }
                    }
                }
                switch 0 {
                case 9223372036854775807:
                    fmt.Println(v5)
                }
            } else {
                fmt.Println(v2, v2)
                v55 := GenS7{F0: 42, F1: "\n\t"}
                _ = v55
            }
        }
        _ = v35
    } else {
        var v56 string = "\n\t"
        v5 = (v5 | (v5 & 255))
        _ = v56
    }
    v57 := GenS7{F0: -1, F1: "hello"}
    v34 = (^(v1))
    if (18446744073709551615 < 1) {
        switch 778 {
        case 0:
            v58 := GenS8{F0: ""}
            _ = v58
        case 2147483647:
            v34 = (v34 / (((v1 * 0)) | 1))
        case -470:
            var v59 uint16 = 0
            _ = v59
        default:
            for fi60 := 0; fi60 < 14; fi60++ {
                fmt.Println(v2)
                _ = fi60
            }
        }
    } else {
        if true {
            for fi61 := 0; fi61 < 1; fi61++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                _ = fi61
            }
            var v62 int8 = v34
            v5 -= v5
            _ = v62
        } else {
            v5 = 133
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v57
}

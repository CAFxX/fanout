package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 int16; F2 uint }
type GenS9 struct { F0 int }
func genFn0() int {
    v6 := []int{128, -1, -9223372036854775808}
    var v7 float64 = (1e10 / 0.1)
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            fmt.Println(fi8, v7)
        } else {
            var v9 uint32 = 871
            switch fi8 {
            case 86:
                v10 := GenS1{F0: ""}
                _ = v10
            case -395:
                v7 += 1e10
            case 2147483647:
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch ((("" + "\n\t") + ("" + "\n\t")) + (("hello" + ("hello" + "")) + "\n\t")) {
                    case "foo bar":
                        var v12 bool = (!true)
                        _ = v12
                    case "a":
                        v13 := GenS1{F0: "foo bar"}
                        _ = v13
                    case "":
                        fmt.Println(fi11, fi11, fi11)
                    default:
                        v9 += v9
                    }
                    v14 := GenS1{F0: ""}
                    switch ("hello" + "\n\t") {
                    case "a":
                        if true {
                            v6 = append(v6, 0)
                        }
                    case "":
                        v15 := GenS1{F0: "a"}
                        _ = v15
                    case "hello":
                        if (v7 > 1.5) {
                            var v16 uint = 65535
                            _ = v16
                        }
                        var v17 uint8 = 59
                        _ = len(v6)
                        _ = v17
                    }
                    _ = fi11
                    _ = v14
                }
            }
            _ = v9
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 bool) int32 {
    v18 := []int{702, genFn0(), 128, 2147483647}
    var v19 string = ""
    v20 := GenS2{F0: 200, F1: 9223372036854775807}
    switch 127 {
    case 2147483647:
        if true {
            var v21 int8 = -66
            _ = v21
        }
    case 32767:
        v22 := GenS1{F0: v19}
        _ = v22
    case -9223372036854775808:
        var v23 uint8 = 255
        _ = v23
    default:
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: "foo bar"}
            if p1 {
                if (p1 || p1) {
                    v18 = append(v18, 128)
                    switch (fi24 / ((fi24) | 1)) {
                    case 0:
                        switch 1 {
                        case 564:
                            p1 = p0
                        case 0:
                            switch (fi24 / (((fi24 ^ genFn0())) | 1)) {
                            case 32767:
                                v26 := GenS2{F0: 612, F1: -417}
                                _ = v26
                            case 42:
                                if p1 {
                                    v19 = v19
                                }
                            case -2147483648:
                                v27 := GenS3{F0: 849}
                                _ = v27
                            }
                        case 488:
                            switch -633 {
                            case -393:
                                fmt.Println(v19, p1, fi24)
                            default:
                                v28 := GenS1{F0: v19}
                                _ = v28
                            }
                        default:
                            for fi29 := 0; fi29 < 19; fi29++ {
                                if (667 > 4294967295) {
                                    v19 = v19
                                    var v30 uint64 = 2
                                    v19 = v19
                                    _ = v30
                                }
                                _ = fi29
                            }
                        }
                    case -273:
                        p1 = ((!p0) || p0)
                    case 396:
                        var v31 uint16 = 32767
                        _ = v31
                    default:
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v19 = v19
                            _ = fi32
                        }
                    }
                }
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            } else {
                fmt.Println(p1, v19, v19)
            }
            _ = fi24
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return 2147483647
}

func main() {
    v1 := []string{("foo bar" + ("hello" + "\n\t")), ("foo bar" + "")}
    var v2 string = ""
    var v3 int16 = 0
    v4 := make(map[int]bool)
    var v5 int32 = -2147483648
    var v33 float32 = (-2.25 / (0.0 + 1.5))
    v34 := GenS1{F0: v2}
    var v35 int8 = -128
    var v36 uint = 65535
    v5 -= v5
    v36 += 1
    for fi37 := 0; fi37 < 8; fi37++ {
        switch v2 {
        case "":
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 uint8 = 127
                _ = fi38
                _ = v39
            }
        case "hello":
            for fi40 := 0; fi40 < 11; fi40++ {
                v2 = ""
                fmt.Println(v5)
                switch genFn0() {
                case 2147483647:
                    v3 -= v3
                }
                var v41 int8 = 1
                v42 := GenS1{F0: ("" + (v2 + ""))}
                var v43 uint32 = 2
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        case "\n\t":
            for fi44 := 0; fi44 < 4; fi44++ {
                v45 := GenS4{F0: 32767}
                v46 := GenS4{F0: 709}
                _ = fi44
                _ = v45
                _ = v46
            }
        default:
            switch (^(fi37)) {
            case 32767:
                var v47 float64 = 0.1
                _ = v47
            case -168:
                for fi48 := 0; fi48 < 9; fi48++ {
                    v49 := v4[(-(fi37))]
                    delete(v4, genFn0())
                    v35 -= (v35 >> ((65535) & 63))
                    v50 := GenS5{F0: (v5 - 816), F1: false, F2: v49}
                    v51 := GenS2{F0: 0, F1: fi48}
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
            }
        }
        v52 := GenS6{F0: -89}
        if (true && (!false)) {
            for fi53 := 0; fi53 < 4; fi53++ {
                switch (fi37 | -1) {
                case 32767:
                    v54 := GenS7{F0: 1000}
                    _ = v54
                case 2147483647:
                    fmt.Println(v3)
                case 691:
                    v33 += v33
                }
                var v55 int16 = (v3 | v3)
                _ = fi53
                _ = v55
            }
            switch int(v5) {
            case 32767:
                v56 := GenS2{F0: 65535, F1: (fi37 << (((^(v36))) & 63))}
                _ = v56
            default:
                var v57 int = fi37
                _ = v57
            }
        } else {
            var v58 uint8 = 127
            for fi59 := 0; fi59 < 6; fi59++ {
                v60 := GenS3{F0: 4294967295}
                v61 := GenS8{F0: (^(v36)), F1: (v3 / (((v3 ^ 0)) | 1)), F2: 691}
                _ = fi59
                _ = v60
                _ = v61
            }
            _ = v58
        }
        _ = fi37
        _ = v52
    }
    v4[32767] = (!((1e10 <= (-2.25 + 0.1)) || ((1.5 * 1.5) == (1.5 - (1e10 * -2.25)))))
    v62 := GenS9{F0: -2147483648}
    delete(v4, 328)
    v33 = v33
    v63 := GenS7{F0: v3}
    var v64 uint64 = 42
    for fi65 := 0; fi65 < 2; fi65++ {
        delete(v4, fi65)
        v33 = (((3.14159 * float32(0.0)) - v33) * 0.1)
        v2 = ""
        _ = fi65
    }
    var v66 float32 = float32(((-2.25 / 3.14159) - -2.25))
    v2 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v66
}

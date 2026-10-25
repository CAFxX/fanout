package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int8, p1 int64, p2 int32) int64 {
    switch "hello" {
    case "hello":
        if false {
            switch 443 {
            case 2147483647:
                switch -829 {
                case -1:
                    switch -913 {
                    case -13:
                        var v7 bool = (true || false)
                        _ = v7
                    case 4294967295:
                        p2 += 2147483647
                    case 0:
                        if (-208 <= -32768) {
                            var v8 uint64 = 240
                            v9 := GenS1{F0: 518, F1: ("a" + ("" + "")), F2: "a"}
                            _ = v8
                            _ = v9
                        } else {
                            v10 := []int{}
                            var v11 int32 = p2
                            _ = v10
                            _ = v11
                        }
                    default:
                        var v12 int64 = -60
                        _ = v12
                    }
                case -9223372036854775808:
                    var v13 string = ""
                    _ = v13
                case 368:
                    fmt.Println(p0, p0)
                default:
                    var v14 uint64 = 22
                    _ = v14
                }
            case 0:
                v15 := []int{-2147483648, 128, 4294967295}
                _ = v15
            case 855:
                v16 := GenS2{F0: 255, F1: 124}
                _ = v16
            }
            v17 := GenS2{F0: 1, F1: 0}
            _ = v17
        }
    case "":
        var v18 string = "hello"
        _ = v18
    }
    var v19 int = -673
    if (1 < 671) {
        v20 := GenS1{F0: (p1 - (p1 >> ((0) & 63))), F1: "\n\t", F2: ("\n\t" + "")}
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS1{F0: p1, F1: "foo bar", F2: "a"}
            v23 := GenS2{F0: 2, F1: 107}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    } else {
        v24 := []int{(v19 << ((834) & 63)), (v19 >> ((42) & 63)), (v19 << ((65535) & 63)), int(p1)}
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS3{F0: 2147483648}
            v19 = (v19 % (((fi25 % ((v19) | 1))) | 1))
            switch fi25 {
            case 127:
                v27 := GenS2{F0: 1, F1: 255}
                _ = v27
            case -313:
                var v28 uint8 = 247
                _ = v28
            case 7:
                switch "" {
                case "hello":
                    var v29 uint32 = 1
                    _ = v29
                case "\n\t":
                    var v30 float64 = 0.0
                    _ = v30
                case "":
                    v19 = -9223372036854775808
                }
            default:
                for fi31 := 0; fi31 < 8; fi31++ {
                    v19 += 128
                    v32 := GenS3{F0: 4294967295}
                    if false {
                        p0 = (p0 >> ((660) & 63))
                        p2 += (p2 & p2)
                    }
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi25
            _ = v26
        }
        if true {
            v24 = append(v24, v19)
            var v33 float64 = 1e10
            var v34 float32 = 0.1
            _ = v33
            _ = v34
        }
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p1
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 42
    var v3 int64 = -9223372036854775808
    var v4 uint = 2
    var v5 int64 = 0
    var v6 int16 = -1
    v35 := GenS4{F0: (((v1 + (v1 + v1)) + "a") + v1), F1: ((v4 << (((v4 - v4)) & 63)) < (v4 * (^(v4))))}
    var v36 int8 = -40
    if (1.5 == (1e10 + 1e10)) {
        switch -35 {
        case 128:
            var v37 float64 = (3.14159 - 0.0)
            _ = v37
        case 0:
            var v38 string = (v1 + v1)
            _ = v38
        case 32768:
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 uint16 = 4
                _ = fi39
                _ = v40
            }
        default:
            if false {
                if (!(((((1.5 / -2.25) + (((((0.0 / 1.0) - 1.5) / 1.0) / 3.14159) + (0.0 + (0.0 - 0.0)))) / 0.1) > ((-2.25 / 3.14159) * 1e10)) && (65535 > 101))) {
                    v41 := []int{0, -6}
                    for fi42 := 0; fi42 < 18; fi42++ {
                        v43 := GenS4{F0: "a", F1: (136 < 128)}
                        v5 -= v3
                        _ = fi42
                        _ = v43
                    }
                    _ = v41
                }
            }
        }
        v3 = -9223372036854775808
        var v44 float64 = 1e10
        _ = v44
    }
    v45 := GenS4{F0: v1, F1: false}
    var v46 uint32 = 623
    v1 = ((v1 + "hello") + v1)
    v46 -= uint32(v4)
    var v47 int = 32767
    var v48 float32 = 0.1
    v49 := []int{}
    var v50 int64 = (^(v3))
    if len(v49) > 0 {
        v49[0] = (v47 | v47)
    }
    v49 = append(v49, v47)
    var v51 bool = ((255 < 128) && (true && true))
    v52 := GenS1{F0: v50, F1: "", F2: v1}
    v53 := GenS3{F0: (v46 & v46)}
    var v54 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}

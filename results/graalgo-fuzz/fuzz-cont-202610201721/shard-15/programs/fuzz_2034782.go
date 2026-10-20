package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int; F2 int64 }
func genFn0() uint8 {
    var v4 int8 = 1
    v5 := []int{-261, 127, 32768, 1}
    v4 = -128
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 int32, p1 string) int8 {
    v6 := GenS1{F0: 869}
    var v7 uint32 = 65535
    var v8 int = 307
    v9 := []int{v8}
    if false {
        if false {
            for fi10 := 0; fi10 < 18; fi10++ {
                p0 = 1
                var v11 uint16 = 978
                var v12 int8 = int8(v8)
                _ = fi10
                _ = v11
                _ = v12
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    } else {
        if len(v9) > 0 {
            v9[0] = v8
        }
        var v13 int64 = -9223372036854775808
        v14 := GenS2{F0: 42}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = float32(3.14159)
    v3 := make(map[uint32]int)
    switch 161 {
    case -9223372036854775808:
        v15 := GenS2{F0: 2147483647}
        _ = v15
    }
    switch "a" {
    case "":
        if true {
            v16 := GenS1{F0: 42}
            v17 := GenS2{F0: 578}
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 float64 = (1.5 * (1e10 / 1e10))
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        }
    default:
        v20 := GenS3{F0: "foo bar", F1: 407}
        _ = v20
    }
    var v21 int32 = 0
    switch 971 {
    case 127:
        v2 -= float32(-2.25)
    default:
        v2 += (float32(((0.0 + (3.14159 * 0.0)) + (float64(((v2 - ((1.5 / v2) + v2)) * (0.0 * (v2 - v2)))) - float64(0.1)))) - (v2 / (float32(0.1) * 1.5)))
    }
    if true {
        if (!true) {
            v21 -= v21
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 int = fi22
                _ = fi22
                _ = v23
            }
            var v24 float32 = v2
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                v2 += 3.14159
                fmt.Println(v21, fi25, fi25)
                v21 = v21
                _ = fi25
            }
            v26 := GenS2{F0: (v21 + -2147483648)}
            _ = v26
        }
        v27 := GenS1{F0: 42}
        fmt.Println(v21, v2)
        _ = v27
    } else {
        v28 := GenS4{F0: -32768}
        var v29 uint16 = 2
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: 2147483647}
    if false {
        v31 := GenS4{F0: 42}
        if (128 > 0) {
            switch 128 {
            case 2147483647:
                if false {
                    v21 = (v21 | v21)
                }
            case 42:
                var v32 bool = (false || (1 <= 42))
                _ = v32
            case 902:
                if false {
                    v2 = float32(1.5)
                } else {
                    switch "" {
                    case "\n\t":
                        var v33 int64 = -417
                        _ = v33
                    case "foo bar":
                        v21 = -820
                    default:
                        var v34 uint = 422
                        _ = v34
                    }
                    switch 14 {
                    case 0:
                        var v35 int64 = -348
                        _ = v35
                    default:
                        var v36 string = ("\n\t" + ("\n\t" + "foo bar"))
                        _ = v36
                    }
                    fmt.Println(v2)
                }
            }
            v37 := GenS2{F0: 0}
            v38 := []int{9223372036854775807}
            _ = v37
            _ = v38
        }
        v39 := []int{9223372036854775807}
        _ = v31
        _ = v39
    } else {
        v21 -= v21
        if true {
            v2 = (0.0 - ((-2.25 - 0.0) + v2))
        } else {
            var v40 int16 = 32767
            for fi41 := 0; fi41 < 13; fi41++ {
                v2 = float32(0.0)
                _ = fi41
            }
            var v42 int8 = 1
            _ = v40
            _ = v42
        }
    }
    v21 += -1
    var v43 int32 = (v21 / (((v21 % (((v21 / ((v21) | 1))) | 1))) | 1))
    var v44 uint16 = 65535
    var v45 uint8 = genFn0()
    var v46 uint16 = (v44 % (((^(v44))) | 1))
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 int = (^(fi47))
        v49 := GenS3{F0: "", F1: -1}
        if true {
            v48 = 0
            fmt.Println(fi47, v44)
            if true {
                var v50 uint64 = 664
                _ = v50
            }
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    v46 = (v46 << ((0) & 63))
    v51 := []int{}
    var v52 uint64 = 787
    for fi53 := 0; fi53 < 8; fi53++ {
        for fi54 := 0; fi54 < 17; fi54++ {
            if (3.14159 < float64(3.14159)) {
                switch "\n\t" {
                case "foo bar":
                    v46 = (v46 % ((v46) | 1))
                }
            } else {
                v44 = v46
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
                v45 -= 127
            }
            _ = fi54
        }
        switch 0 {
        case 32767:
            for fi55 := 0; fi55 < 19; fi55++ {
                var v56 uint = uint(v52)
                var v57 uint = uint(v45)
                fmt.Println(fi53, fi55, v21)
                _ = fi55
                _ = v56
                _ = v57
            }
        case 101:
            if true {
                var v58 float64 = float64(v2)
                v51 = append(v51, (-(fi53)))
                _ = v58
            }
        default:
            if ((-9223372036854775808 > 9223372036854775807) || true) {
                for fi59 := 0; fi59 < 7; fi59++ {
                    switch ("" + "a") {
                    case "foo bar":
                        var v60 uint16 = v46
                        if true {
                            switch 32767 {
                            case 32767:
                                var v61 uint8 = v45
                                _ = v61
                            }
                        }
                        v62 := GenS3{F0: "a", F1: (v21 / (((v43 << ((1) & 63))) | 1))}
                        _ = v60
                        _ = v62
                    case "":
                        if (v2 >= (v2 * float32(float64((0.0 + -2.25))))) {
                            v44 -= (^(v46))
                            var v63 int64 = -9223372036854775808
                            var v64 float64 = (0.1 + ((-2.25 * 1e10) - (0.1 - float64(v2))))
                            switch "foo bar" {
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    var v65 int64 = v63
                                    _ = v65
                                case "a":
                                    var v66 int32 = (v43 & (v21 % ((v43) | 1)))
                                    _ = v66
                                }
                            case "foo bar":
                                v64 -= ((float64(-2.25) / float64(v2)) - 3.14159)
                            default:
                                var v67 float32 = 0.1
                                _ = v67
                            }
                            _ = v63
                            _ = v64
                        } else {
                            v52 = 83
                            var v68 uint64 = 42
                            v69 := GenS5{F0: (fi53 & (fi53 << ((42) & 63)))}
                            _ = v68
                            _ = v69
                        }
                        v70 := GenS6{F0: int16(v43), F1: fi59, F2: 274}
                        _ = v70
                    case "hello":
                        if ("" <= "\n\t") {
                            v46 -= v46
                        } else {
                            v43 -= -911
                            v2 -= 1.5
                        }
                    }
                    _ = fi59
                }
                fmt.Println(v2)
                v71 := GenS6{F0: 1, F1: fi53, F2: 0}
                _ = v71
            } else {
                switch (("" + "a") + ("" + "foo bar")) {
                case "a":
                    for fi72 := 0; fi72 < 11; fi72++ {
                        if len(v51) > 0 {
                            v51[0] = (fi72 / (((fi72 - (fi72 - 32768))) | 1))
                        }
                        _ = fi72
                    }
                case "foo bar":
                    v52 = v52
                case "":
                    var v73 uint = 986
                    _ = v73
                default:
                    _ = len(v3)
                }
            }
        }
        var v74 float64 = 0.0
        _ = fi53
        _ = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v51
    _ = v52
}

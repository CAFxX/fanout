package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS7 struct { F0 int64 }
func genFn0() int8 {
    v5 := GenS1{F0: 255, F1: -32768}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = "\n\t"
        for fi8 := 0; fi8 < 9; fi8++ {
            v7 = v7
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if false {
            var v10 uint64 = 0
            _ = v10
        }
        var v11 float64 = 0.1
        var v12 int64 = 1
        _ = fi9
        _ = v11
        _ = v12
    }
    var v13 int64 = 42
    v14 := GenS1{F0: 1, F1: 123}
    _ = v13
    _ = v14
    _ = v5
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int8 = 1
    v4 := make(map[string]bool)
    fmt.Println(v3)
    if (!(!false)) {
        for fi15 := 0; fi15 < 5; fi15++ {
            v3 = v3
            v3 = v3
            _ = fi15
        }
    } else {
        var v16 uint32 = 2
        var v17 bool = true
        if v17 {
            for fi18 := 0; fi18 < 5; fi18++ {
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int16 = 42
                    _ = fi19
                    _ = v20
                }
                v21 := GenS2{F0: 294}
                v22 := GenS1{F0: 32767, F1: -1}
                _ = fi18
                _ = v21
                _ = v22
            }
        }
        _ = v16
        _ = v17
    }
    v23 := v2[-50]
    if (2 != 567) {
        fmt.Println(v23, v23)
        v24 := GenS3{F0: 18, F1: 0, F2: false}
        switch "a" {
        case "":
            if true {
                for fi25 := 0; fi25 < 7; fi25++ {
                    switch ("hello" + "") {
                    case "":
                        var v26 int = (fi25 % (((fi25 % ((fi25) | 1))) | 1))
                        _ = v26
                    case "\n\t":
                        fmt.Println(v23, fi25)
                    }
                    switch fi25 {
                    case 2147483647:
                        v3 += 1
                    case 128:
                        var v27 uint = 220
                        _ = v27
                    case -826:
                        if false {
                            var v28 int = fi25
                            if (!true) {
                                fmt.Println(fi25, v3, v3)
                                v23 = (^(v23))
                                if false {
                                    _ = len(v1)
                                    v28 += (fi25 & fi25)
                                    if false {
                                        v29 := []int{v28}
                                        var v30 int8 = (v3 >> ((255) & 63))
                                        var v31 uint = 673
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 uint8 = 1
                                        _ = v32
                                    }
                                }
                            }
                            for fi33 := 0; fi33 < 9; fi33++ {
                                switch "foo bar" {
                                case "hello":
                                    v28 = fi25
                                default:
                                    v28 = (fi33 ^ (-(fi25)))
                                }
                                var v34 int64 = int64(fi25)
                                v3 += v3
                                _ = fi33
                                _ = v34
                            }
                            _ = v28
                        } else {
                            var v35 uint64 = 412
                            _ = v35
                        }
                    default:
                        var v36 int8 = 1
                        _ = v36
                    }
                    _ = fi25
                }
                v37 := []int{}
                if len(v37) > 0 {
                    v37[0] = 43
                }
                _ = v37
            }
        case "\n\t":
            var v38 uint = 18446744073709551615
            _ = v38
        case "hello":
            v39 := GenS4{F0: "foo bar", F1: 42}
            _ = v39
        default:
            v40 := GenS5{F0: 1}
            _ = v40
        }
        _ = v24
    } else {
        switch (("" + "\n\t") + "foo bar") {
        case "\n\t":
            for fi41 := 0; fi41 < 11; fi41++ {
                var v42 uint64 = 1
                _ = fi41
                _ = v42
            }
        }
        v23 = 1
        v3 -= (^(v3))
    }
    switch ("hello" + "a") {
    case "hello":
        if (128 > 821) {
            var v43 int8 = (v3 & (v3 % (((v3 * (^(v3)))) | 1)))
            switch int(v23) {
            case 1:
                switch 1 {
                case 2147483647:
                    v23 = 2
                }
            }
            _ = v43
        } else {
            v23 = v23
            v3 += 0
            if (742 > 255) {
                v44 := GenS6{F0: false, F1: 127, F2: 337}
                var v45 int8 = 100
                _ = v44
                _ = v45
            } else {
                v46 := GenS2{F0: 42}
                for fi47 := 0; fi47 < 9; fi47++ {
                    var v48 uint8 = 128
                    v49 := GenS5{F0: 1}
                    v23 = v23
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                for fi50 := 0; fi50 < 17; fi50++ {
                    switch int(v3) {
                    case 367:
                        v3 = (v3 ^ v3)
                    case -224:
                        v23 = v23
                    case -2147483648:
                        v51 := GenS7{F0: 850}
                        _ = v51
                    default:
                        switch "" {
                        case "hello":
                            fmt.Println(v23)
                        case "a":
                            v23 -= v23
                        default:
                            v52 := []int{(fi50 & (fi50 ^ fi50)), (fi50 & 1)}
                            _ = v52
                        }
                    }
                    if (false && (false && false)) {
                        switch "\n\t" {
                        case "a":
                            v53 := GenS1{F0: uint16(v23), F1: -32768}
                            _ = v53
                        case "":
                            for fi54 := 0; fi54 < 14; fi54++ {
                                v3 = (v3 % ((52) | 1))
                                v55 := v2[(fi50 << ((65535) & 63))]
                                _ = fi54
                                _ = v55
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                _ = len(v1)
                            case "a":
                                fmt.Println(v3, v3, v3)
                            case "":
                                fmt.Println(fi50)
                            }
                        default:
                            v3 = (v3 & (v3 | 127))
                        }
                        for fi56 := 0; fi56 < 19; fi56++ {
                            var v57 string = "\n\t"
                            _ = fi56
                            _ = v57
                        }
                        switch "hello" {
                        case "\n\t":
                            for fi58 := 0; fi58 < 18; fi58++ {
                                var v59 string = ("foo bar" + "a")
                                _ = fi58
                                _ = v59
                            }
                        case "foo bar":
                            fmt.Println(v23, fi50)
                        case "a":
                            for fi60 := 0; fi60 < 19; fi60++ {
                                var v61 uint32 = (v23 - (v23 ^ v23))
                                _ = fi60
                                _ = v61
                            }
                        }
                    } else {
                        v23 = (^(v23))
                        fmt.Println(v3)
                    }
                    _ = fi50
                }
                _ = v46
            }
        }
    case "foo bar":
        v62 := GenS3{F0: 458, F1: 722, F2: (false && ((-9223372036854775808 < -310) || false))}
        _ = v62
    }
    var v63 float32 = ((1e10 - 1e10) / 3.14159)
    for fi64 := 0; fi64 < 20; fi64++ {
        for fi65 := 0; fi65 < 8; fi65++ {
            var v66 uint = 720
            v3 -= (v3 / ((127) | 1))
            var v67 uint32 = v23
            _ = fi65
            _ = v66
            _ = v67
        }
        if (("\n\t" + ("foo bar" + "a")) > "\n\t") {
            switch "foo bar" {
            case "hello":
                v68 := v2[0]
                _ = v68
            case "\n\t":
                _ = len(v1)
            case "a":
                switch (^(fi64)) {
                case 2147483647:
                    var v69 bool = true
                    _ = v69
                case -9223372036854775808:
                    v70 := GenS3{F0: 449, F1: 396, F2: (!true)}
                    _ = v70
                default:
                    v71 := []int{}
                    _ = v71
                }
            }
            switch (-(fi64)) {
            case 790:
                var v72 uint8 = 0
                _ = v72
            case -1:
                for fi73 := 0; fi73 < 1; fi73++ {
                    fmt.Println(v23, fi73, v63)
                    v74 := GenS7{F0: 9223372036854775807}
                    _ = fi73
                    _ = v74
                }
            }
        }
        var v75 float64 = -2.25
        _ = fi64
        _ = v75
    }
    if true {
        if true {
            v76 := GenS3{F0: 1, F1: 4294967295, F2: (2147483647 == 1)}
            _ = v76
        } else {
            var v77 float64 = -2.25
            var v78 float64 = 0.1
            _ = len(v1)
            _ = v77
            _ = v78
        }
    }
    v1[("hello" + ("hello" + "foo bar"))] = 0
    var v79 uint = 848
    if (42 < 906) {
        if (0 == 242) {
            var v80 uint8 = 42
            fmt.Println(v23, v80, v80)
            _ = v80
        }
    }
    var v81 uint = v79
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v63:", v63)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v63
    _ = v79
    _ = v81
}

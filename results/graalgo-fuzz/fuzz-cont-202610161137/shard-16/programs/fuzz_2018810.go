package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 int8; F1 int16; F2 int32 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 int32) int8 {
    v7 := GenS1{F0: 0}
    p0 = 2147483647
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int) bool {
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 2
        v9 = 4294967295
        if true {
            v10 := GenS1{F0: int16(p0)}
            switch "foo bar" {
            case "hello":
                switch (p0 & (^(p0))) {
                case 32767:
                    v9 = v9
                case 119:
                    v11 := []int{-503}
                    _ = v11
                case 9223372036854775807:
                    var v12 int = (p0 * 128)
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        var v14 uint64 = 536
                        _ = v14
                    } else {
                        if false {
                            v15 := []int{(p0 % (((fi8 % ((fi13) | 1))) | 1)), 42}
                            v16 := GenS2{F0: 0, F1: true}
                            p0 -= (p0 | 2147483647)
                            _ = v15
                            _ = v16
                        } else {
                            p0 -= 215
                        }
                        switch p0 {
                        case 176:
                            if false {
                                v17 := []int{(fi13 - (-(fi8)))}
                                var v18 float32 = 3.14159
                                if len(v17) > 0 {
                                    v17[0] = fi13
                                }
                                _ = v17
                                _ = v18
                            }
                        case -9223372036854775808:
                            if true {
                                p0 = fi13
                                v9 = (v9 - v9)
                            } else {
                                v19 := GenS3{F0: 1000, F1: ("hello" + "foo bar")}
                                var v20 float32 = 0.0
                                _ = v19
                                _ = v20
                            }
                        case -170:
                            fmt.Println(fi13)
                        }
                        v9 -= 42
                    }
                    _ = fi13
                }
            }
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    if (0.0 <= 3.14159) {
        v21 := GenS2{F0: 279, F1: false}
        if true {
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS2{F0: -32768, F1: (128 != fi22)}
                p0 += fi22
                _ = fi22
                _ = v23
            }
            if true {
                var v24 uint = 42
                _ = v24
            } else {
                fmt.Println(p0)
                v25 := GenS2{F0: 1000, F1: (true || (1 == 42))}
                if (!false) {
                    var v26 float64 = (1.5 - -2.25)
                    v27 := []int{p0, (p0 - p0), (p0 + 42), (p0 % ((p0) | 1))}
                    p0 = p0
                    _ = v26
                    _ = v27
                }
                _ = v25
            }
            v28 := GenS2{F0: 1000, F1: (!(!(!true)))}
            _ = v28
        }
        if (174 > 0) {
            p0 = p0
            for fi29 := 0; fi29 < 16; fi29++ {
                for fi30 := 0; fi30 < 13; fi30++ {
                    if true {
                        var v31 int = (-(p0))
                        _ = v31
                    } else {
                        if false {
                            var v32 uint8 = 42
                            _ = v32
                        } else {
                            var v33 uint16 = 65535
                            switch p0 {
                            case 1:
                                fmt.Println(fi29, v33, p0)
                            default:
                                p0 += fi30
                            }
                            _ = v33
                        }
                    }
                    _ = fi30
                }
                switch ("\n\t" + "a") {
                case "a":
                    for fi34 := 0; fi34 < 11; fi34++ {
                        p0 -= fi34
                        p0 = 527
                        switch fi29 {
                        case -25:
                            v35 := GenS3{F0: 42, F1: "hello"}
                            _ = v35
                        }
                        _ = fi34
                    }
                case "":
                    v36 := []int{(fi29 * (-(p0))), 42}
                    _ = v36
                }
                _ = fi29
            }
            if (!(("\n\t" < "") || (100 > -67))) {
                if true {
                    switch p0 {
                    case -1:
                        p0 = p0
                    case 4294967295:
                        p0 = p0
                    }
                }
                fmt.Println(p0)
            }
        } else {
            var v37 int = (p0 ^ (p0 ^ p0))
            _ = v37
        }
        _ = v21
    }
    var v38 uint32 = 2147483648
    for fi39 := 0; fi39 < 12; fi39++ {
        v40 := []int{(fi39 % ((128) | 1)), p0, 9223372036854775807}
        _ = fi39
        _ = v40
    }
    _ = p0
    _ = v38
    return false
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    var v3 uint16 = 1
    var v4 uint = 760
    var v5 uint8 = 42
    v6 := make(map[uint32]int)
    v3 += (v3 + 845)
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(fi41, v4, v3)
        if false {
            v3 = v3
            var v42 string = "foo bar"
            v4 -= v4
            _ = v42
        }
        fmt.Println(v2, fi41)
        _ = fi41
    }
    var v43 uint32 = 2147483647
    for fi44 := 0; fi44 < 14; fi44++ {
        for fi45 := 0; fi45 < 15; fi45++ {
            v46 := GenS2{F0: -191, F1: v2}
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    var v47 float32 = 1e10
    v48 := GenS4{F0: (v3 * v3), F1: 42, F2: 9223372036854775807}
    var v49 string = ""
    for fi50 := 0; fi50 < 1; fi50++ {
        if (!(1 <= 165)) {
            var v51 uint16 = 32767
            var v52 uint8 = (^(v5))
            _ = v51
            _ = v52
        }
        v53 := GenS5{F0: 0, F1: -32768, F2: 0}
        _ = fi50
        _ = v53
    }
    if true {
        if v2 {
            v54 := []int{}
            _ = v54
        } else {
            if (v49 != "") {
                if (!(v5 > uint8(v4))) {
                    switch 9223372036854775807 {
                    case 677:
                        v55 := GenS6{F0: (v4 << ((v4) & 63)), F1: 1}
                        _ = v55
                    case 42:
                        v56 := GenS6{F0: (v4 << ((4294967295) & 63)), F1: 1}
                        _ = v56
                    }
                    v57 := GenS7{F0: -2147483648}
                    v3 = (v3 >> ((v4) & 63))
                    _ = v57
                }
            } else {
                switch 4294967295 {
                case 128:
                    v43 -= v43
                }
                var v58 bool = v2
                v6[(v43 % (((v43 - v43)) | 1))] = 833
                _ = v58
            }
        }
        if (float32(((float64(v47) * 1.5) / ((3.14159 * (3.14159 + (-2.25 + (0.0 + 3.14159)))) / -2.25))) < v47) {
            switch "hello" {
            case "a":
                for fi59 := 0; fi59 < 15; fi59++ {
                    if (v5 <= (v5 / (((v5 - (^(v5)))) | 1))) {
                        v60 := GenS5{F0: -15, F1: 1000, F2: 2147483647}
                        fmt.Println(v5)
                        _ = v60
                    } else {
                        fmt.Println(v5, v5)
                        var v61 string = v49
                        v49 = v61
                        _ = v61
                    }
                    _ = fi59
                }
            default:
                for fi62 := 0; fi62 < 1; fi62++ {
                    v49 = "hello"
                    _ = fi62
                }
            }
        } else {
            v3 = v3
            v49 = v49
            for fi63 := 0; fi63 < 9; fi63++ {
                var v64 uint64 = 669
                v64 += (v64 + (^(v64)))
                for fi65 := 0; fi65 < 20; fi65++ {
                    v64 = v64
                    var v66 int = (fi65 / (((^(fi63))) | 1))
                    fmt.Println(fi65, v47)
                    v67 := GenS8{F0: (^(v5)), F1: 42}
                    var v68 uint32 = v43
                    _ = fi65
                    _ = v66
                    _ = v67
                    _ = v68
                }
                _ = fi63
                _ = v64
            }
        }
        v69 := []int{}
        _ = v69
    } else {
        switch v49 {
        case "\n\t":
            v4 = v4
        case "hello":
            var v70 uint32 = v43
            _ = v70
        case "a":
            v71 := []int{1, -793, 2147483647}
            _ = v71
        default:
            v72 := GenS4{F0: 2, F1: 9223372036854775807, F2: 1}
            _ = v72
        }
        v73 := v1[2]
        switch 2147483647 {
        case -9223372036854775808:
            if (18446744073709551615 == 4294967295) {
                for fi74 := 0; fi74 < 14; fi74++ {
                    v75 := GenS3{F0: 1, F1: v49}
                    v76 := GenS4{F0: 32767, F1: 0, F2: 42}
                    var v77 uint64 = 42
                    _ = fi74
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
        case 42:
            v3 = (^(v3))
        case 128:
            v3 = (v3 + (v3 / ((v3) | 1)))
        }
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}

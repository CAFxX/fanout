package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int16; F1 int64; F2 uint64 }
type GenS8 struct { F0 int }
func genFn0() bool {
    var v5 int32 = 42
    var v6 uint32 = 2
    v6 = 2147483648
    _ = v5
    _ = v6
    return (1 >= 232)
}
func genFn1(p0 int, p1 uint16) int64 {
    if ((-2.25 + 1.5) <= ((1e10 * -2.25) + 1.5)) {
        var v7 uint16 = (^(p1))
        v7 += (p1 - p1)
        fmt.Println(p0)
        _ = v7
    } else {
        v8 := GenS1{F0: -119, F1: 255}
        fmt.Println(p1, p0)
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                var v11 int32 = 0
                _ = v11
            } else {
                v12 := GenS2{F0: 42, F1: 880, F2: (-289 < 42)}
                _ = v12
            }
            var v13 int64 = 9223372036854775807
            _ = fi10
            _ = v13
        }
        if false {
            p1 = (p1 | (p1 & p1))
        } else {
            v14 := GenS3{F0: (fi9 - (fi9 & (fi9 ^ 1)))}
            v15 := GenS3{F0: p0}
            p1 += p1
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: -128, F1: 127}
        _ = fi9
        _ = v16
    }
    v17 := GenS1{F0: -7, F1: 128}
    _ = p0
    _ = p1
    _ = v17
    return -9
}

func main() {
    var v1 uint32 = 163
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 int32 = v2
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            if false {
                if true {
                    switch -1 {
                    case 4294967295:
                        var v20 int32 = v2
                        _ = v20
                    case 346:
                        var v21 uint16 = 2
                        _ = v21
                    case -1:
                        var v22 float32 = (0.0 + (0.1 * 3.14159))
                        _ = v22
                    default:
                        v23 := GenS4{F0: (^(fi19)), F1: 0}
                        _ = v23
                    }
                    var v24 int16 = -1
                    _ = v24
                }
            } else {
                v25 := v3[(-(fi19))]
                var v26 int16 = -1
                v27 := GenS1{F0: 100, F1: 42}
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi19
        }
        _ = fi18
    }
    fmt.Println(v1)
    var v28 int32 = (v4 + -1)
    if (true && (18 < 100)) {
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := v3[42]
            _ = fi29
            _ = v30
        }
        var v31 float64 = 3.14159
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := []int{(^(fi32))}
            v28 -= v4
            switch (-(fi32)) {
            case -4:
                v34 := GenS5{F0: 0, F1: 709}
                _ = v34
            case 9223372036854775807:
                v35 := GenS6{F0: v28, F1: false}
                _ = v35
            }
            _ = fi32
            _ = v33
        }
        _ = v31
    } else {
        fmt.Println(v4)
        var v36 uint16 = 42
        _ = v36
    }
    if false {
        var v37 int8 = 0
        fmt.Println(v2)
        if (!false) {
            var v38 uint32 = (v1 << ((699) & 63))
            switch -1 {
            case 42:
                v39 := GenS7{F0: 0, F1: 1, F2: 42}
                _ = v39
            }
            _ = v38
        }
        _ = v37
    } else {
        v40 := []int{}
        var v41 float32 = 0.0
        _ = v40
        _ = v41
    }
    var v42 string = "foo bar"
    for fi43 := 0; fi43 < 8; fi43++ {
        switch v42 {
        case "":
            switch fi43 {
            case 0:
                var v44 float32 = float32((0.1 * 3.14159))
                _ = v44
            case 128:
                if false {
                    if (4294967295 != 69) {
                        v28 -= -1
                        if ((true && (0 < 255)) && (432 > v1)) {
                            if (-21 >= int8(v1)) {
                                if ((850 <= 127) || (fi43 <= fi43)) {
                                    v1 -= (^(v1))
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        _ = len(v3)
                                        var v46 int32 = (v4 | v4)
                                        _ = fi45
                                        _ = v46
                                    }
                                    v47 := GenS8{F0: 1}
                                    _ = v47
                                }
                            } else {
                                fmt.Println(v1)
                                var v48 int8 = 1
                                var v49 int32 = 0
                                _ = v48
                                _ = v49
                            }
                            v1 -= v1
                        } else {
                            fmt.Println(v1)
                            if true {
                                if true {
                                    _ = len(v3)
                                    if (false || (((true || (42 < 1)) && ((fi43 >> ((551) & 63)) >= fi43)) && true)) {
                                        var v50 int = fi43
                                        _ = v50
                                    }
                                }
                                delete(v3, fi43)
                                if false {
                                    for fi51 := 0; fi51 < 5; fi51++ {
                                        switch 2147483647 {
                                        case 0:
                                            fmt.Println(v28, v28)
                                        default:
                                            v28 = -1
                                        }
                                        var v52 uint64 = 19
                                        var v53 uint8 = 42
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            } else {
                                for fi54 := 0; fi54 < 1; fi54++ {
                                    var v55 int8 = 42
                                    switch "a" {
                                    case "foo bar":
                                        v56 := GenS1{F0: v55, F1: 128}
                                        _ = v56
                                    case "":
                                        v2 -= v4
                                        if false {
                                            v57 := GenS2{F0: v2, F1: 4294967295, F2: (0 < -32768)}
                                            _ = v57
                                        }
                                        v58 := GenS2{F0: -2147483648, F1: 2, F2: (true || (361 >= 399))}
                                        _ = v58
                                    }
                                    v59 := GenS5{F0: uint16(v1), F1: 42}
                                    _ = fi54
                                    _ = v55
                                    _ = v59
                                }
                                if (42 == 42) {
                                    v1 = v1
                                } else {
                                    var v60 uint = 1
                                    for fi61 := 0; fi61 < 17; fi61++ {
                                        var v62 uint32 = v1
                                        switch 128 {
                                        case 4294967295:
                                            v42 = v42
                                            if false {
                                                var v63 int = fi43
                                                var v64 bool = true
                                                fmt.Println(v62, v60)
                                                var v65 float64 = 1e10
                                                v66 := []int{v63, -9223372036854775808, 480, fi43}
                                                if len(v66) > 0 {
                                                    v66[0] = -131
                                                }
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                fmt.Println(v28, v60, v1)
                                            }
                                        }
                                        switch v42 {
                                        case "foo bar":
                                            fmt.Println(v4)
                                        }
                                        _ = fi61
                                        _ = v62
                                    }
                                    _ = v60
                                }
                                switch v42 {
                                case "":
                                    var v67 int = (fi43 & (fi43 << ((547) & 63)))
                                    _ = v67
                                default:
                                    fmt.Println(fi43, v2, v42)
                                }
                            }
                            var v68 int64 = -1
                            _ = v68
                        }
                        fmt.Println(fi43, v1, fi43)
                    }
                    var v69 int32 = v4
                    _ = v69
                }
            case 39:
                if (127 == 42) {
                    fmt.Println(v1)
                } else {
                    fmt.Println(v42)
                    v70 := v3[32767]
                    var v71 uint64 = 1
                    _ = v70
                    _ = v71
                }
            }
        }
        if (false || true) {
            var v72 int64 = -9223372036854775808
            switch 9223372036854775807 {
            case 50:
                fmt.Println(v28)
            case 42:
                var v73 int32 = (v4 >> ((42) & 63))
                _ = v73
            case 128:
                for fi74 := 0; fi74 < 20; fi74++ {
                    v75 := []int{(fi43 >> ((uint(v1)) & 63)), (^(fi74)), (fi74 % (((fi43 % ((539) | 1))) | 1)), (fi43 & fi43)}
                    v76 := GenS5{F0: 32767, F1: 0}
                    _ = fi74
                    _ = v75
                    _ = v76
                }
            default:
                v77 := GenS8{F0: fi43}
                _ = v77
            }
            var v78 int64 = (v72 - (v72 >> ((822) & 63)))
            _ = v72
            _ = v78
        } else {
            var v79 uint8 = 42
            v80 := v3[127]
            switch v42 {
            case "":
                v28 = v28
            default:
                v81 := []int{fi43, (v80 + fi43), (^(fi43))}
                _ = v81
            }
            _ = v79
            _ = v80
        }
        _ = fi43
    }
    v82 := GenS3{F0: -345}
    var v83 bool = false
    v2 += 1
    v28 = (^(v28))
    if v83 {
        v4 -= v4
        var v84 int16 = 1
        _ = v84
    } else {
        var v85 uint16 = 1
        var v86 bool = ((v85 / ((65535) | 1)) < v85)
        _ = v85
        _ = v86
    }
    switch 42 {
    case 42:
        switch 9223372036854775807 {
        case -1:
            fmt.Println(v83, v4)
        }
    case 0:
        for fi87 := 0; fi87 < 6; fi87++ {
            fmt.Println(v1, fi87)
            v3[fi87] = 346
            var v88 bool = v83
            _ = fi87
            _ = v88
        }
    }
    v83 = (1 == 42)
    if ((v42 + v42) < "") {
        v42 = (v42 + v42)
        v1 = (^(v1))
        v3[-1] = -1
    }
    var v89 int8 = -87
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v83:", v83)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v82
    _ = v83
    _ = v89
}

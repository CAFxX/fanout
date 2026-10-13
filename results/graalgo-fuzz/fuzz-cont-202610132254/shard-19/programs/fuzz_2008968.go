package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
func genFn0() string {
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v7 int8 = -1
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "\n\t":
        var v8 uint8 = 255
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(fi9)
            fmt.Println(fi9)
            _ = fi9
        }
    case "hello":
        var v10 string = "a"
        _ = v10
    }
    switch 128 {
    case 32768:
        var v11 int16 = 32767
        _ = v11
    case 127:
        var v12 float32 = 3.14159
        _ = v12
    case 1:
        if len(v5) > 0 {
            v5[0] = -337
        }
    default:
        _ = len(v5)
    }
    _ = v5
    return ("hello" + "a")
}
func genFn1(p0 uint) uint16 {
    var v13 int8 = 0
    p0 += 18446744073709551615
    _ = p0
    _ = v13
    return 65535
}

func main() {
    var v1 uint32 = 731
    var v2 int32 = 42
    var v3 uint = 434
    var v4 uint32 = 0
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint8 = 1
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 31
    v3 = v3
    var v17 float32 = (3.14159 / float32(0.1))
    for fi18 := 0; fi18 < 15; fi18++ {
        if (1 < 538) {
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(v17, v4, v4)
                _ = fi19
            }
        } else {
            v20 := GenS1{F0: 32767, F1: ((fi18 / (((fi18 * 32767)) | 1)) != (fi18 / ((fi18) | 1))), F2: 4294967295}
            if ((v2 / (((v2 & int32(v16))) | 1)) != v2) {
                switch fi18 {
                case 128:
                    v2 += (v2 | 42)
                case -690:
                    var v21 int16 = 42
                    _ = v21
                case -9223372036854775808:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := []int{(fi22 * fi18), fi18, fi18, fi18}
                        _ = fi22
                        _ = v23
                    }
                default:
                    v24 := GenS1{F0: (fi18 ^ -1), F1: (-1 == -2147483648), F2: (^(v4))}
                    _ = v24
                }
                v25 := []int{(fi18 + fi18), (fi18 >> ((v3) & 63))}
                var v26 uint32 = v4
                _ = v25
                _ = v26
            }
            if true {
                v27 := GenS1{F0: (fi18 - (-(fi18))), F1: (v16 == (v16 % ((uint8(v1)) | 1))), F2: v1}
                v28 := GenS1{F0: fi18, F1: false, F2: (v4 + v1)}
                _ = v27
                _ = v28
            } else {
                var v29 uint = (v3 + (v3 >> ((0) & 63)))
                if (float32(1.5) >= 1e10) {
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                } else {
                    v1 = (v4 * 2147483647)
                    v17 += -2.25
                    v4 = uint32(v29)
                }
                var v31 string = "\n\t"
                _ = v29
                _ = v31
            }
            _ = v20
        }
        if false {
            v32 := []int{fi18, 42, (fi18 & fi18)}
            var v33 int32 = v2
            for fi34 := 0; fi34 < 1; fi34++ {
                if true {
                    if false {
                        if ((0 == -1) && false) {
                            if false {
                                if false {
                                    fmt.Println(v3, v1, fi34)
                                    var v35 uint16 = genFn1((^(v3)))
                                    switch fi34 {
                                    case 32767:
                                        v36 := GenS2{F0: "hello", F1: (^(v35))}
                                        _ = v36
                                    case -712:
                                        if len(v32) > 0 {
                                            v32[0] = 4294967295
                                        }
                                    }
                                    fmt.Println(v4, v3)
                                    _ = v35
                                } else {
                                    var v37 int = (fi34 / ((741) | 1))
                                    var v38 int16 = 1
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                var v39 float64 = 1e10
                                _ = v39
                            }
                            switch fi18 {
                            case -851:
                                v1 += v4
                            case -946:
                                var v40 bool = false
                                _ = v40
                            case 42:
                                var v41 uint8 = 127
                                _ = v41
                            default:
                                v33 += v33
                            }
                        }
                    }
                    v42 := GenS1{F0: (fi18 % ((fi18) | 1)), F1: true, F2: (v4 % ((2147483647) | 1))}
                    switch fi18 {
                    case 2147483647:
                        var v43 float32 = (0.1 - (v17 - v17))
                        _ = v43
                    case 32767:
                        var v44 int16 = 222
                        _ = v44
                    case -1:
                        v1 = (v4 - (^(v4)))
                    }
                    _ = v42
                } else {
                    if len(v32) > 0 {
                        v32[0] = (fi34 / ((fi18) | 1))
                    }
                    fmt.Println(v17)
                }
                if (174 <= (v16 - (v16 ^ 52))) {
                    if (1e10 < 1e10) {
                        var v45 uint = 65535
                        if (!((-1 > int8(v4)) && (false && (1 <= -1)))) {
                            v17 -= v17
                        } else {
                            if (0 > genFn1(0)) {
                                v46 := GenS3{F0: 1, F1: -32768, F2: (^(v16))}
                                _ = len(v32)
                                _ = v46
                            }
                            var v47 uint = (^(v3))
                            _ = v47
                        }
                        v48 := GenS1{F0: (fi18 - (fi18 >> ((0) & 63))), F1: true, F2: (v4 & v1)}
                        _ = v45
                        _ = v48
                    }
                } else {
                    v49 := GenS2{F0: "\n\t", F1: 2}
                    var v50 int64 = 573
                    v4 -= (v4 - (^(v4)))
                    _ = v49
                    _ = v50
                }
                var v51 uint8 = (^(v16))
                v52 := GenS4{F0: 1, F1: -950, F2: 32767}
                _ = fi34
                _ = v51
                _ = v52
            }
            _ = v32
            _ = v33
        } else {
            var v53 int32 = (v2 >> (((v3 << ((v3) & 63))) & 63))
            _ = v53
        }
        _ = fi18
    }
    v54 := []int{32767, 42}
    v55 := GenS1{F0: 979, F1: false, F2: v4}
    for fi56 := 0; fi56 < 6; fi56++ {
        _ = len(v54)
        var v57 int32 = v2
        v57 += 142
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 5; fi58++ {
        var v59 float32 = float32(-2.25)
        for fi60 := 0; fi60 < 8; fi60++ {
            v61 := GenS1{F0: 72, F1: true, F2: v4}
            _ = fi60
            _ = v61
        }
        _ = fi58
        _ = v59
    }
    fmt.Println(v4)
    v4 = v4
    fmt.Println(v3, v17, v17)
    for fi62 := 0; fi62 < 10; fi62++ {
        var v63 uint64 = 4294967295
        var v64 uint = v3
        switch -106 {
        case -782:
            v3 = (^(v3))
        case -1:
            switch fi62 {
            case 32768:
                fmt.Println(v64, v16, v64)
            case 737:
                v1 = v4
            }
        case 9223372036854775807:
            var v65 int32 = (v2 - (v2 << (((^(v64))) & 63)))
            _ = v65
        }
        _ = fi62
        _ = v63
        _ = v64
    }
    fmt.Println(v3)
    var v66 int16 = 792
    var v67 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v66
    _ = v67
}

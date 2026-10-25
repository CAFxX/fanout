package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 bool) int8 {
    if p2 {
        p2 = (true || (p2 && (77 > 603)))
    } else {
        if (p2 || (65535 > 781)) {
            p0 = (^(p0))
            var v5 uint32 = p0
            p0 = (^(p0))
            _ = v5
        }
        fmt.Println(p2, p2, p1)
    }
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := []int{-9223372036854775808, -1, 1}
    var v2 string = "a"
    v3 := []string{"", ("\n\t" + v2), (v2 + v2)}
    var v4 bool = (true && true)
    v4 = (!(v4 && v4))
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        var v7 uint = 2
        _ = v7
    case 669:
        var v8 uint8 = 42
        _ = v8
    case -1:
        for fi9 := 0; fi9 < 3; fi9++ {
            switch v2 {
            case "a":
                v4 = (255 == 255)
            case "\n\t":
                var v10 int64 = 42
                _ = v10
            case "":
                switch fi9 {
                case 127:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32768:
                    switch v2 {
                    case "a":
                        fmt.Println(fi9, v4)
                    case "\n\t":
                        if ((42 <= 62) && false) {
                            fmt.Println(v4, v2)
                        } else {
                            fmt.Println(fi9, fi9, v4)
                            switch v2 {
                            case "foo bar":
                                if v4 {
                                    for fi11 := 0; fi11 < 4; fi11++ {
                                        fmt.Println(v4)
                                        v12 := GenS1{F0: -1, F1: 42, F2: 168}
                                        switch fi9 {
                                        case -1:
                                            if v4 {
                                                v4 = true
                                            } else {
                                                fmt.Println(v4, fi9)
                                                var v13 uint16 = 2
                                                _ = v13
                                            }
                                        case 484:
                                            v2 = (v2 + v2)
                                        }
                                        _ = fi11
                                        _ = v12
                                    }
                                }
                            case "\n\t":
                                for fi14 := 0; fi14 < 4; fi14++ {
                                    var v15 int64 = -9223372036854775808
                                    _ = fi14
                                    _ = v15
                                }
                            case "a":
                                v16 := GenS2{F0: 65535}
                                _ = v16
                            default:
                                v4 = v4
                            }
                        }
                    case "foo bar":
                        v2 = v2
                    default:
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v18 := GenS1{F0: 1, F1: 18446744073709551615, F2: 168}
                            v4 = (164 <= 42)
                            _ = fi17
                            _ = v18
                        }
                    }
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS3{F0: (v2 + v2)}
                var v21 uint32 = 42
                var v22 int64 = int64(fi9)
                if (v2 <= v2) {
                    v2 = ""
                }
                if v4 {
                    if v4 {
                        if true {
                            var v23 uint16 = 65535
                            v24 := GenS2{F0: (v21 + v21)}
                            _ = v23
                            _ = v24
                        } else {
                            switch v2 {
                            case "\n\t":
                                if ((((0.1 < 0.1) || v4) || v4) && (v4 || v4)) {
                                    v4 = (-9223372036854775808 == v22)
                                } else {
                                    var v25 int64 = (v22 - (v22 - (v22 >> ((65535) & 63))))
                                    _ = v25
                                }
                            case "foo bar":
                                v26 := GenS3{F0: v2}
                                _ = v26
                            case "hello":
                                v2 = v2
                                if (-2.25 <= 0.1) {
                                    v2 = v2
                                }
                                var v27 uint8 = 72
                                _ = v27
                            }
                        }
                        var v28 int64 = 0
                        v28 += (^(v28))
                        _ = v28
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            var v29 string = "a"
            _ = fi9
            _ = v29
        }
    }
    v30 := GenS4{F0: v4, F1: 4294967295, F2: 4294967295}
    v31 := GenS3{F0: v2}
    if true {
        var v32 uint32 = 4294967295
        var v33 uint16 = 32767
        _ = v32
        _ = v33
    }
    switch "a" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if false {
            fmt.Println(v2, v4)
            switch ("" + v2) {
            case "a":
                v34 := GenS3{F0: "a"}
                _ = v34
            case "\n\t":
                for fi35 := 0; fi35 < 9; fi35++ {
                    v36 := GenS5{F0: 2, F1: 951}
                    _ = fi35
                    _ = v36
                }
            case "":
                var v37 uint = 1
                _ = v37
            }
        } else {
            v4 = v4
        }
    }
    var v38 uint32 = 1
    var v39 int8 = 100
    v38 = 0
    var v40 int32 = -1
    var v41 bool = v4
    switch v2 {
    case "\n\t":
        v4 = v4
    default:
        var v42 uint16 = uint16(v38)
        _ = v42
    }
    switch "" {
    case "":
        if (659 < 475) {
            v43 := GenS4{F0: v4, F1: 4294967295, F2: 32767}
            _ = v43
        }
    default:
        switch "foo bar" {
        case "":
            v44 := GenS5{F0: 4294967295, F1: 2}
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        v46 := GenS6{F0: 9223372036854775807, F1: v38}
        var v47 int64 = -9223372036854775808
        v48 := GenS7{F0: -247}
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    if ((v39 <= -123) || v41) {
        if (!(!(genFn0(v38, 127, v4) != (v39 << ((18446744073709551615) & 63))))) {
            v49 := GenS8{F0: v38}
            _ = v49
        }
        switch ("" + v2) {
        case "":
            var v50 float64 = ((0.0 * 0.0) * (float64(3.14159) * 1.5))
            _ = v50
        }
    } else {
        if (v41 || (0 > 32767)) {
            if true {
                var v51 int16 = 32767
                _ = v51
            } else {
                var v52 uint16 = 0
                v53 := GenS4{F0: (559 > 1), F1: 4294967295, F2: -9223372036854775808}
                _ = v52
                _ = v53
            }
            fmt.Println(v2, v38, v4)
            for fi54 := 0; fi54 < 1; fi54++ {
                for fi55 := 0; fi55 < 10; fi55++ {
                    if v41 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = (-(fi55))
                        }
                        var v56 int8 = v39
                        _ = v56
                    } else {
                        v40 -= 0
                    }
                    _ = fi55
                }
                var v57 int8 = (v39 >> ((939) & 63))
                var v58 int64 = -559
                _ = fi54
                _ = v57
                _ = v58
            }
        }
        var v59 float64 = 0.1
        _ = v59
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        if true {
            var v61 int64 = -752
            _ = v61
        } else {
            for fi62 := 0; fi62 < 9; fi62++ {
                _ = len(v3)
                if len(v1) > 0 {
                    v1[0] = (fi62 & 128)
                }
                _ = fi62
            }
        }
        v1 = append(v1, (fi60 / (((^(fi60))) | 1)))
        for fi63 := 0; fi63 < 16; fi63++ {
            var v64 uint32 = 1
            v1 = append(v1, 127)
            var v65 uint64 = 42
            v41 = false
            v39 = -1
            _ = fi63
            _ = v64
            _ = v65
        }
        _ = fi60
    }
    v2 = (("\n\t" + v2) + (v2 + ((v2 + v2) + "hello")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}

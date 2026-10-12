package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 bool; F2 int32 }
type GenS6 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    var v1 bool = (42 != 1)
    v2 := []uint32{339, 1, 539, 590}
    v3 := make(map[int]int)
    v1 = (!v1)
    v4 := GenS1{F0: "hello"}
    var v5 uint16 = 1
    if (!v1) {
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 int8 = 40
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 int = 128
                _ = len(v2)
                fmt.Println(v10)
                v1 = false
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: "foo bar"}
            _ = fi6
            _ = v11
        }
        if true {
            var v12 bool = false
            v13 := GenS1{F0: ("\n\t" + "a")}
            _ = v12
            _ = v13
        } else {
            if (!(!false)) {
                if v1 {
                    v5 -= 255
                }
            }
            v1 = v1
        }
        v1 = v1
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            if (398 != 18446744073709551615) {
                v5 = (v5 * 42)
                v15 := v3[42]
                var v16 int8 = -105
                _ = v15
                _ = v16
            } else {
                if (float64((0.0 + 1e10)) == 0.0) {
                    v17 := GenS2{F0: 1}
                    switch fi14 {
                    case 4294967295:
                        var v18 uint64 = 52
                        _ = v18
                    }
                    if v1 {
                        for fi19 := 0; fi19 < 8; fi19++ {
                            switch (fi19 ^ (fi14 + (fi14 >> ((0) & 63)))) {
                            case 32767:
                                switch "foo bar" {
                                case "":
                                    switch (fi19 / ((fi19) | 1)) {
                                    case -683:
                                        if v1 {
                                            if (684 <= 18446744073709551615) {
                                                _ = len(v2)
                                                v20 := GenS3{F0: 2, F1: 2}
                                                var v21 int16 = 32767
                                                v1 = (v1 || v1)
                                                fmt.Println(v1, fi19, fi14)
                                                _ = v20
                                                _ = v21
                                            }
                                            v2 = append(v2, 2147483647)
                                            v22 := GenS4{F0: 4294967295}
                                            _ = v22
                                        }
                                    }
                                case "\n\t":
                                    if true {
                                        var v23 uint32 = 167
                                        _ = v23
                                    } else {
                                        v5 = v5
                                        var v24 uint8 = 248
                                        _ = v24
                                    }
                                    v5 = (^(v5))
                                case "foo bar":
                                    if v1 {
                                        switch (fi14 + fi14) {
                                        case 1:
                                            switch "hello" {
                                            case "a":
                                                v1 = (845 != -32768)
                                            default:
                                                v5 += 871
                                            }
                                        case 127:
                                            switch ("\n\t" + "hello") {
                                            case "":
                                                if v1 {
                                                    fmt.Println(v1)
                                                    var v25 int32 = 42
                                                    _ = v25
                                                } else {
                                                    var v26 uint16 = (v5 % (((v5 ^ (^(v5)))) | 1))
                                                    var v27 uint64 = 554
                                                    v28 := GenS4{F0: (fi19 | -9223372036854775808)}
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                }
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                            }
                                            var v29 float64 = 3.14159
                                            _ = v29
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    } else {
                                        var v30 string = "foo bar"
                                        switch v30 {
                                        case "":
                                            v31 := v3[42]
                                            _ = v31
                                        case "\n\t":
                                            fmt.Println(v5, fi19)
                                        case "foo bar":
                                            var v32 int = fi19
                                            _ = v32
                                        default:
                                            v5 = (v5 % ((v5) | 1))
                                        }
                                        fmt.Println(v5, fi19, v30)
                                        _ = v30
                                    }
                                default:
                                    var v33 int64 = -9223372036854775808
                                    _ = v33
                                }
                            case 1:
                                v2 = append(v2, 0)
                            case 351:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v34 uint32 = 4294967295
                                _ = v34
                            }
                            _ = fi19
                        }
                        v35 := GenS4{F0: fi14}
                        _ = v35
                    } else {
                        var v36 float32 = 0.0
                        v37 := GenS5{F0: ("hello" + (("a" + "foo bar") + ("\n\t" + ""))), F1: false, F2: 0}
                        var v38 uint = 18446744073709551615
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = v17
                } else {
                    for fi39 := 0; fi39 < 20; fi39++ {
                        if true {
                            v2 = append(v2, 748)
                            var v40 uint16 = v5
                            var v41 bool = v1
                            _ = v40
                            _ = v41
                        }
                        if v1 {
                            delete(v3, fi39)
                            v5 -= v5
                        } else {
                            v2 = append(v2, 65535)
                            var v42 int16 = -1
                            _ = v42
                        }
                        v2 = append(v2, 1)
                        _ = fi39
                    }
                    var v43 int8 = 42
                    _ = v43
                }
            }
            v44 := GenS6{F0: "hello", F1: int8(v5), F2: v5}
            _ = len(v3)
            _ = fi14
            _ = v44
        }
    }
    var v45 uint16 = 65535
    if (9223372036854775807 != 0) {
        v5 = v45
        for fi46 := 0; fi46 < 6; fi46++ {
            v5 = 255
            if ((fi46 ^ (fi46 >> ((42) & 63))) != (^(fi46))) {
                var v47 uint32 = 42
                for fi48 := 0; fi48 < 16; fi48++ {
                    var v49 uint8 = 127
                    _ = fi48
                    _ = v49
                }
                _ = v47
            }
            _ = fi46
        }
        var v50 uint = 350
        _ = v50
    } else {
        var v51 int = -9223372036854775808
        var v52 bool = (true || (65535 < 871))
        v52 = v52
        _ = v51
        _ = v52
    }
    _ = len(v3)
    v53 := GenS6{F0: "", F1: -105, F2: (v5 >> ((uint(v45)) & 63))}
    for fi54 := 0; fi54 < 13; fi54++ {
        v1 = ("hello" > "\n\t")
        for fi55 := 0; fi55 < 17; fi55++ {
            var v56 uint32 = uint32(v5)
            v45 = (v45 * v5)
            v56 += v56
            v56 += 39
            fmt.Println(fi54)
            var v57 bool = v1
            v3[42] = (fi54 | -2147483648)
            _ = fi55
            _ = v56
            _ = v57
        }
        var v58 int = 9223372036854775807
        _ = fi54
        _ = v58
    }
    var v59 int8 = -62
    fmt.Println(v1, v1)
    for fi60 := 0; fi60 < 5; fi60++ {
        v5 -= (v45 | (v5 % (((v5 / (((v5 ^ 0)) | 1))) | 1)))
        v5 = v45
        switch 628 {
        case 0:
            var v61 int32 = 2147483647
            _ = v61
        case 32768:
            for fi62 := 0; fi62 < 18; fi62++ {
                var v63 uint16 = (v45 % (((v45 << ((0) & 63))) | 1))
                _ = fi62
                _ = v63
            }
        case -795:
            v45 = v5
        default:
            var v64 float32 = 1.5
            _ = v64
        }
        _ = fi60
    }
    if (!v1) {
        if (false || true) {
            var v65 bool = (-2.25 < (0.1 / 1.0))
            v59 = (v59 / ((-77) | 1))
            if (!v65) {
                var v66 uint = 0
                _ = v66
            } else {
                var v67 uint32 = 863
                _ = v67
            }
            _ = v65
        }
        var v68 uint64 = 320
        var v69 uint8 = 42
        _ = v68
        _ = v69
    } else {
        var v70 int64 = -1
        v45 += v5
        _ = v70
    }
    if (88 == 199) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = v5
        switch "a" {
        case "foo bar":
            var v71 uint64 = uint64(v45)
            _ = v71
        default:
            v3[32767] = int(v5)
        }
    }
    var v72 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v53
    _ = v59
    _ = v72
}

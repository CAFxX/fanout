package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int; F1 int; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8; F1 int }
func genFn0(p0 int64) string {
    v6 := []int{127, 42}
    v7 := GenS1{F0: 645}
    _ = p0
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{}
    var v3 string = ("" + ("\n\t" + "foo bar"))
    v4 := make(map[int]int)
    var v5 int8 = 93
    v5 -= v5
    for fi8 := 0; fi8 < 4; fi8++ {
        if (true && (149 >= 127)) {
            var v9 int8 = (v5 / (((^(v5))) | 1))
            v10 := GenS1{F0: 32767}
            _ = v10
            _ = v9
        }
        _ = len(v2)
        switch v3 {
        case "":
            var v11 uint = 4294967295
            _ = v11
        case "foo bar":
            if true {
                v5 -= (v5 ^ (-(v5)))
                v12 := GenS2{F0: true}
                for fi13 := 0; fi13 < 11; fi13++ {
                    if true {
                        var v14 uint16 = 1
                        var v15 float64 = float64(v1)
                        if false {
                            switch fi13 {
                            case -1:
                                v3 = v3
                                v16 := GenS3{F0: 18446744073709551615}
                                _ = v16
                            case 32767:
                                if false {
                                    v17 := GenS3{F0: 4294967295}
                                    _ = v17
                                } else {
                                    var v18 bool = false
                                    _ = v18
                                }
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                switch "foo bar" {
                case "hello":
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v20 := GenS1{F0: 255}
                        _ = fi19
                        _ = v20
                    }
                }
                var v21 float32 = v1
                _ = v21
            }
        case "a":
            for fi22 := 0; fi22 < 8; fi22++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = v3
                fmt.Println(fi8, v5)
                switch fi22 {
                case -943:
                    var v23 uint8 = 144
                    v3 = v3
                    v24 := GenS4{F0: 1}
                    _ = v23
                    _ = v24
                }
                switch fi22 {
                case 306:
                    var v25 int = (fi8 - 32768)
                    _ = v25
                case 1:
                    var v26 uint16 = 1
                    _ = v26
                case 9223372036854775807:
                    var v27 int32 = 42
                    _ = v27
                }
                if (true && false) {
                    v28 := GenS1{F0: 156}
                    if true {
                        v5 = (-(v5))
                    }
                    _ = v28
                } else {
                    fmt.Println(v1, v1)
                }
                if (((true || (false && true)) || true) || (!false)) {
                    switch ((v3 + "foo bar") + v3) {
                    case "\n\t":
                        if false {
                            _ = len(v4)
                        }
                    default:
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(fi22, fi8)
                        case "\n\t":
                            if false {
                                _ = len(v4)
                                var v29 float32 = v1
                                var v30 uint32 = 660
                                _ = v29
                                _ = v30
                            } else {
                                v31 := GenS2{F0: (0 == -32768)}
                                if ((3.14159 != (v1 + v1)) && true) {
                                    v32 := GenS4{F0: 0}
                                    var v33 int16 = -523
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 string = ""
                                    _ = len(v4)
                                    v35 := GenS4{F0: -9223372036854775808}
                                    _ = v34
                                    _ = v35
                                }
                                fmt.Println(v3, v1)
                                _ = v31
                            }
                        case "foo bar":
                            var v36 uint8 = 128
                            _ = v36
                        default:
                            if (2147483648 > 398) {
                                v5 -= v5
                                var v37 int8 = (v5 | (v5 << ((42) & 63)))
                                _ = v37
                            } else {
                                v38 := GenS4{F0: -1}
                                v3 = v3
                                _ = v38
                            }
                        }
                    }
                    if true {
                        var v39 uint16 = 409
                        v2 = append(v2, 42)
                        v5 += 127
                        _ = v39
                    } else {
                        var v40 uint = 4294967295
                        _ = v40
                    }
                    var v41 uint64 = 18446744073709551615
                    _ = v41
                } else {
                    v1 -= v1
                    if false {
                        fmt.Println(v3, v1, v3)
                        fmt.Println(v1, fi22)
                        v4[fi22] = fi8
                    }
                }
                v42 := GenS1{F0: 42}
                _ = fi22
                _ = v42
            }
        }
        _ = fi8
    }
    var v43 float64 = 3.14159
    fmt.Println(v5, v3, v1)
    for fi44 := 0; fi44 < 15; fi44++ {
        switch "hello" {
        case "a":
            v3 = v3
        case "\n\t":
            v45 := GenS5{F0: -1, F1: -9223372036854775808}
            _ = v45
        case "":
            v46 := GenS6{F0: fi44, F1: -9223372036854775808, F2: 403}
            _ = v46
        default:
            switch fi44 {
            case 608:
                v47 := GenS6{F0: fi44, F1: 1, F2: 4294967295}
                _ = v47
            default:
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            }
        }
        _ = fi44
    }
    fmt.Println(v3, v3)
    switch -724 {
    case 42:
        fmt.Println(v43)
    }
    var v48 float64 = float64(v1)
    fmt.Println(v43, v43)
    if false {
        if (!(!true)) {
            v49 := GenS7{F0: true}
            for fi50 := 0; fi50 < 8; fi50++ {
                var v51 uint = 65535
                v52 := GenS8{F0: 9223372036854775807}
                for fi53 := 0; fi53 < 6; fi53++ {
                    var v54 uint64 = 468
                    var v55 uint32 = 763
                    v3 = v3
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = v49
        }
        var v56 bool = (!(!((933 < 572) || true)))
        switch 2147483647 {
        case -905:
            var v57 float64 = (float64(v1) + 0.1)
            _ = v57
        }
        _ = v56
    } else {
        for fi58 := 0; fi58 < 5; fi58++ {
            switch 4294967295 {
            case 42:
                v2 = append(v2, 65535)
            default:
                v43 += float64(1.5)
            }
            for fi59 := 0; fi59 < 17; fi59++ {
                if (v3 < v3) {
                    var v60 int8 = v5
                    v3 = "a"
                    v5 = (v60 >> ((18446744073709551615) & 63))
                    if (541 < 0) {
                        v61 := GenS3{F0: 27}
                        v62 := GenS9{F0: 58, F1: 127}
                        switch (^(fi58)) {
                        case -353:
                            switch (("\n\t" + v3) + "a") {
                            case "hello":
                                v48 -= v43
                            case "":
                                var v63 uint64 = 0
                                _ = v63
                            default:
                                var v64 string = v3
                                _ = v64
                            }
                        case 4294967295:
                            v4[(fi59 >> ((571) & 63))] = (fi58 >> ((42) & 63))
                        }
                        var v65 uint8 = 0
                        _ = v61
                        _ = v62
                        _ = v65
                    } else {
                        var v66 uint = 0
                        v5 = -108
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        v43 = v48
                        if (2 < 884) {
                            v43 = (((1.5 + (0.1 / 0.1)) - (-2.25 + ((v48 + ((v48 * float64(v1)) / ((float64(v1) + 0.1) - v48))) * -2.25))) - v43)
                            var v67 float32 = 1.5
                            _ = v67
                        }
                        delete(v4, fi58)
                        var v68 float64 = 3.14159
                        _ = v66
                        _ = v68
                    }
                    var v69 int64 = -9223372036854775808
                    _ = v60
                    _ = v69
                } else {
                    fmt.Println(v48)
                    var v70 uint = 255
                    var v71 string = v3
                    _ = v70
                    _ = v71
                }
                var v72 uint32 = 238
                v2 = append(v2, (v72 / (((^(v72))) | 1)))
                _ = fi59
                _ = v72
            }
            var v73 int64 = 479
            _ = fi58
            _ = v73
        }
    }
    var v74 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v5
    _ = v74
}

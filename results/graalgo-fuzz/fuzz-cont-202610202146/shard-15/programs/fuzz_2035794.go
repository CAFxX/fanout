package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0() uint {
    var v6 int = -1
    v6 += (v6 | v6)
    if (42 >= 2147483647) {
        v6 = v6
    }
    _ = v6
    return 772
}
func genFn1(p0 uint32, p1 uint8, p2 string) int32 {
    fmt.Println(p0, p1)
    var v7 bool = true
    switch p2 {
    case "a":
        var v8 int = 1
        _ = v8
    }
    var v9 uint8 = (p1 ^ (^(p1)))
    p1 = 203
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 122
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 902
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    v1 = ("hello" + ("foo bar" + "foo bar"))
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{(fi11 & fi11)}
            fmt.Println(v2)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := GenS1{F0: 995}
    if v3 {
        var v14 uint64 = 18446744073709551615
        var v15 int16 = 32767
        v16 := GenS1{F0: 210}
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint8 = 161
            v20 := GenS2{F0: 0}
            switch fi17 {
            case 32768:
                if (v3 && (69 != -1)) {
                    v21 := GenS2{F0: 42}
                    _ = v21
                }
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v1 = (("hello" + v1) + v1)
    v22 := []int{221}
    v1 = "foo bar"
    for fi23 := 0; fi23 < 13; fi23++ {
        if (-2.25 == -2.25) {
            switch -281 {
            case 32768:
                for fi24 := 0; fi24 < 10; fi24++ {
                    if (int64(fi24) < -276) {
                        var v25 int8 = -1
                        v26 := v5[v2]
                        fmt.Println(v25, v26)
                        _ = v25
                        _ = v26
                    }
                    v3 = true
                    var v27 int8 = 0
                    _ = fi24
                    _ = v27
                }
            }
        } else {
            var v28 int64 = 9223372036854775807
            var v29 int32 = genFn1((v2 & v2), 72, "hello")
            _ = v28
            _ = v29
        }
        v30 := GenS2{F0: -1}
        fmt.Println(fi23, v3, v1)
        _ = fi23
        _ = v30
    }
    var v31 uint = 931
    for fi32 := 0; fi32 < 19; fi32++ {
        v2 = (v2 - v2)
        _ = fi32
    }
    v3 = false
    fmt.Println(v3, v2)
    switch (v1 + v1) {
    case "\n\t":
        switch 4294967295 {
        case 9223372036854775807:
            v33 := GenS3{F0: 105, F1: 42}
            _ = v33
        case 638:
            var v34 uint32 = v2
            _ = v34
        case 42:
            var v35 float64 = ((1e10 + ((-2.25 - -2.25) / 0.1)) * ((1.5 * -2.25) / 3.14159))
            _ = v35
        default:
            switch 453 {
            case 128:
                v4[v2] = (v2 ^ v2)
            case 42:
                var v36 int32 = genFn1(128, 127, v1)
                _ = v36
            }
        }
    case "foo bar":
        v2 -= (v2 >> (((v31 | v31)) & 63))
    }
    v37 := GenS4{F0: 127, F1: v3}
    _ = len(v4)
    if true {
        v3 = false
        for fi38 := 0; fi38 < 7; fi38++ {
            if v3 {
                v39 := GenS5{F0: uint64(v31), F1: 451}
                if ((v2 >= v2) && (true || false)) {
                    v31 = (v31 ^ (v31 * v31))
                    for fi40 := 0; fi40 < 14; fi40++ {
                        fmt.Println(v31, fi38)
                        _ = fi40
                    }
                    v31 = (v31 | (v31 ^ (v31 + (^(v31)))))
                } else {
                    var v41 uint64 = 979
                    for fi42 := 0; fi42 < 16; fi42++ {
                        _ = len(v5)
                        _ = fi42
                    }
                    delete(v5, (v2 - v2))
                    _ = v41
                }
                _ = v39
            }
            _ = fi38
        }
        var v43 string = v1
        _ = v43
    } else {
        if v3 {
            v1 = (v1 + "hello")
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 int32 = 1
                if (fi44 >= int(v45)) {
                    var v46 uint = 42
                    for fi47 := 0; fi47 < 9; fi47++ {
                        _ = len(v22)
                        switch (("hello" + v1) + "a") {
                        case "hello":
                            v45 = v45
                        case "a":
                            var v48 uint16 = 1
                            _ = v48
                        case "foo bar":
                            fmt.Println(fi44, fi47, v2)
                        default:
                            v1 = ((("foo bar" + v1) + v1) + (("" + "\n\t") + v1))
                        }
                        if (true && v3) {
                            var v49 float32 = 1.5
                            _ = v49
                        }
                        v3 = (!(0 != 9223372036854775807))
                        var v50 int32 = (v45 * v45)
                        switch fi47 {
                        case 32768:
                            v51 := GenS2{F0: 100}
                            _ = v51
                        default:
                            if (!false) {
                                if false {
                                    fmt.Println(fi44)
                                    v52 := GenS4{F0: int8(v46), F1: v3}
                                    _ = len(v22)
                                    _ = v52
                                }
                                fmt.Println(v1)
                                fmt.Println(v3, v1, fi44)
                            }
                        }
                        _ = fi47
                        _ = v50
                    }
                    v45 = (v45 / (((^(v45))) | 1))
                    _ = v46
                }
                v3 = (-2.25 == (1.5 * 0.1))
                _ = fi44
                _ = v45
            }
            for fi53 := 0; fi53 < 1; fi53++ {
                var v54 int8 = -128
                _ = fi53
                _ = v54
            }
        } else {
            v55 := GenS2{F0: -24}
            v56 := GenS3{F0: -1, F1: 2147483647}
            _ = v55
            _ = v56
        }
        if ((float64(((-2.25 * 1e10) / 3.14159)) != 1e10) && v3) {
            for fi57 := 0; fi57 < 18; fi57++ {
                if (v31 >= v31) {
                    for fi58 := 0; fi58 < 16; fi58++ {
                        switch ("foo bar" + v1) {
                        case "\n\t":
                            var v59 uint = (^(v31))
                            _ = v59
                        case "foo bar":
                            fmt.Println(v31, v1)
                        }
                        _ = fi58
                    }
                    if true {
                        v1 = v1
                        var v60 float64 = (-2.25 / -2.25)
                        _ = v60
                    }
                } else {
                    v22 = append(v22, int(v31))
                    var v61 float32 = 3.14159
                    if v3 {
                        switch "hello" {
                        case "foo bar":
                            switch 1 {
                            case -9223372036854775808:
                                v62 := GenS3{F0: 0, F1: genFn1(v2, 127, "a")}
                                _ = v62
                            default:
                                for fi63 := 0; fi63 < 6; fi63++ {
                                    v2 = (v2 * v2)
                                    _ = fi63
                                }
                            }
                        default:
                            var v64 int = fi57
                            _ = v64
                        }
                        v65 := GenS6{F0: 100}
                        if true {
                            v66 := GenS4{F0: 1, F1: false}
                            _ = len(v22)
                            _ = v66
                        }
                        _ = v65
                    } else {
                        v2 = (v2 | 0)
                    }
                    _ = v61
                }
                _ = fi57
            }
        } else {
            switch -489 {
            case 2147483647:
                v67 := GenS3{F0: 1, F1: 0}
                _ = v67
            case -951:
                for fi68 := 0; fi68 < 15; fi68++ {
                    v69 := GenS4{F0: -128, F1: (((v2 * v2) < (^(v2))) || v3)}
                    v70 := GenS7{F0: "", F1: 0}
                    var v71 float32 = 0.1
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
            }
            v2 -= v2
            if (!(!(!(55 < -1)))) {
                fmt.Println(v31)
            } else {
                var v72 string = v1
                for fi73 := 0; fi73 < 9; fi73++ {
                    var v74 int32 = 2147483647
                    _ = fi73
                    _ = v74
                }
                v75 := GenS4{F0: -128, F1: false}
                _ = v72
                _ = v75
            }
        }
    }
    var v76 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v76
}

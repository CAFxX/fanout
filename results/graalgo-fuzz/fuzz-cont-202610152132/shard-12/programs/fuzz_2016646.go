package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int; F1 uint64; F2 string }
type GenS5 struct { F0 uint8; F1 string; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int32; F2 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 uint8 }
func genFn0() int {
    switch -2147483648 {
    case 122:
        if true {
            var v5 uint32 = 0
            v5 += 42
            var v6 string = ""
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                if (!(0 == 480)) {
                    fmt.Println(fi7)
                    v8 := GenS1{F0: (fi7 + 0)}
                    var v9 uint32 = 941
                    _ = v8
                    _ = v9
                }
                if (false && (0 != 1)) {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(fi7)
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := []int{32767}
                        var v13 uint8 = 42
                        fmt.Println(fi7, v13, fi7)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v15 := []int{fi14}
                        var v16 uint8 = 242
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    switch 32768 {
                    case 449:
                        var v17 bool = false
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 uint64 = 643
                            v20 := []int{fi18}
                            v19 = 280
                            fmt.Println(fi7, v19)
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                    var v21 string = ("foo bar" + "foo bar")
                    _ = v21
                }
                _ = fi7
            }
            var v22 uint64 = 339
            v23 := GenS1{F0: 32768}
            _ = v22
            _ = v23
        }
    case 128:
        var v24 uint32 = 0
        _ = v24
    case 32767:
        var v25 uint16 = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 uint = 19
            v28 = (v28 ^ v28)
            v29 := []int{fi27, (fi26 - fi27), fi26, 426}
            var v30 uint32 = 1
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi26
    }
    var v31 uint = 65535
    _ = v31
    return 440
}
func genFn1(p0 int64) string {
    if true {
        p0 -= 0
        var v32 uint = 42
        _ = v32
    } else {
        var v33 uint8 = 255
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 bool = false
            p0 = (^(p0))
            v36 := []int{(fi34 + fi34), (fi34 & 128)}
            _ = fi34
            _ = v35
            _ = v36
        }
        v33 = (v33 | 128)
        _ = v33
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        p0 += 42
        for fi38 := 0; fi38 < 2; fi38++ {
            if true {
                var v39 int32 = 0
                fmt.Println(fi38, fi37)
                v39 = (v39 / ((v39) | 1))
                _ = v39
            } else {
                p0 -= (p0 - (p0 >> ((42) & 63)))
                var v40 int32 = 1
                p0 -= -561
                _ = v40
            }
            fmt.Println(fi38)
            p0 -= p0
            _ = fi38
        }
        v41 := []int{genFn0()}
        _ = fi37
        _ = v41
    }
    _ = p0
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"a", "a"}
    var v3 float32 = 0.0
    var v4 uint16 = 94
    var v42 uint8 = 127
    v43 := GenS2{F0: -23, F1: 0, F2: (^(v42))}
    switch genFn0() {
    case 128:
        v4 -= (v4 << ((350) & 63))
    case 32768:
        v44 := GenS3{F0: v42, F1: 1, F2: (^(v42))}
        _ = v44
    case 32767:
        var v45 uint32 = 1
        _ = v45
    default:
        var v46 uint32 = 65535
        _ = v46
    }
    var v47 int8 = 127
    fmt.Println(v42)
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 bool = (!((fi48 & (fi48 - fi48)) <= 2147483647))
        v47 -= v47
        if true {
            switch (("\n\t" + "") + ("a" + ("foo bar" + "foo bar"))) {
            case "":
                v4 = 0
            case "\n\t":
                if v49 {
                    fmt.Println(v49, v47)
                    for fi50 := 0; fi50 < 8; fi50++ {
                        switch fi48 {
                        case -803:
                            v49 = v49
                        default:
                            var v51 int64 = -436
                            _ = v51
                        }
                        if v49 {
                            if (false && (false || false)) {
                                if (384 <= 42) {
                                    switch ((("foo bar" + "foo bar") + "a") + ("a" + "hello")) {
                                    case "foo bar":
                                        v4 -= (v4 ^ (v4 ^ v4))
                                    }
                                }
                                if false {
                                    switch -284 {
                                    case 42:
                                        fmt.Println(v49)
                                    case -412:
                                        v52 := GenS1{F0: (fi50 | (fi48 | fi50))}
                                        _ = v52
                                    case -9223372036854775808:
                                        v53 := GenS4{F0: (-(fi48)), F1: 42, F2: "\n\t"}
                                        _ = v53
                                    default:
                                        switch "a" {
                                        case "":
                                            v42 = 0
                                        case "a":
                                            v4 -= (v4 << ((2) & 63))
                                        }
                                    }
                                    v42 = 255
                                    v2 = append(v2, "")
                                } else {
                                    switch genFn0() {
                                    case -9223372036854775808:
                                        var v54 bool = (-498 == fi50)
                                        var v55 int32 = 1
                                        _ = v54
                                        _ = v55
                                    case 128:
                                        v49 = (!(-2147483648 <= 0))
                                    default:
                                        var v56 int32 = 0
                                        _ = v56
                                    }
                                    switch fi48 {
                                    case 446:
                                        v57 := GenS5{F0: (v42 << ((458) & 63)), F1: "foo bar", F2: 2}
                                        _ = v57
                                    case 2147483647:
                                        switch genFn0() {
                                        case 4294967295:
                                            if v49 {
                                                var v58 int64 = -9223372036854775808
                                                _ = v58
                                            } else {
                                                v4 = 42
                                            }
                                        case -9223372036854775808:
                                            var v59 int32 = 1
                                            _ = v59
                                        default:
                                            if (!(!true)) {
                                                fmt.Println(fi50)
                                            } else {
                                                v49 = (42 > 1)
                                                v42 = 255
                                                if false {
                                                    if (!v49) {
                                                        v60 := GenS6{F0: fi48}
                                                        v3 = float32(0.1)
                                                        var v61 int = fi50
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    v47 -= (v47 >> ((65535) & 63))
                                                    v62 := GenS3{F0: (^(v42)), F1: -1, F2: v42}
                                                    _ = v62
                                                }
                                                v42 += v42
                                                v63 := GenS6{F0: (fi48 ^ -174)}
                                                _ = v63
                                            }
                                        }
                                    }
                                    v64 := GenS7{F0: v47, F1: -2147483648, F2: true}
                                    _ = v64
                                }
                            }
                            var v65 string = "a"
                            v66 := GenS8{F0: 4294967295}
                            _ = v65
                            _ = v66
                        }
                        _ = fi50
                    }
                } else {
                    switch "hello" {
                    case "a":
                        if false {
                            v67 := GenS3{F0: v42, F1: -1, F2: (v42 * v42)}
                            _ = v67
                        }
                    case "foo bar":
                        _ = len(v2)
                    case "hello":
                        for fi68 := 0; fi68 < 9; fi68++ {
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    delete(v1, 2147483647)
                                } else {
                                    switch (fi68 / (((^(fi68))) | 1)) {
                                    case -648:
                                        fmt.Println(v3, v42, v3)
                                    case -396:
                                        v3 += v3
                                    default:
                                        if (!v49) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                    v3 = (0.1 - 0.1)
                                }
                            case "\n\t":
                                if (!false) {
                                    var v69 string = "hello"
                                    _ = v69
                                }
                            case "hello":
                                if (false && false) {
                                    fmt.Println(v47, v47)
                                    var v70 uint16 = 754
                                    _ = v70
                                } else {
                                    if true {
                                        v71 := GenS7{F0: v47, F1: 1, F2: true}
                                        fmt.Println(v4, v49)
                                        var v72 uint16 = 0
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v49 = v49
                                        var v73 uint = 411
                                        fmt.Println(v3, v73, v47)
                                        _ = v73
                                    }
                                }
                            }
                            if false {
                                var v74 uint16 = (v4 | 833)
                                if v49 {
                                    var v75 uint32 = 306
                                    _ = v75
                                } else {
                                    fmt.Println(fi48, fi48, v47)
                                    fmt.Println(fi68, v4, v4)
                                    v3 -= 1.5
                                    v3 -= v3
                                    var v76 uint64 = 0
                                    _ = v76
                                }
                                var v77 uint8 = 0
                                if v49 {
                                    v42 = 1
                                    fmt.Println(v74, v4, v49)
                                }
                                _ = len(v2)
                                _ = v74
                                _ = v77
                            }
                            v78 := GenS9{F0: 939, F1: v42}
                            var v79 int64 = 9223372036854775807
                            v80 := GenS3{F0: v42, F1: 0, F2: (v42 << ((2) & 63))}
                            _ = fi68
                            _ = v78
                            _ = v79
                            _ = v80
                        }
                    }
                    v42 = 57
                }
            }
        }
        _ = fi48
        _ = v49
    }
    if false {
        if len(v2) > 0 {
            v2[0] = ("a" + "hello")
        }
    } else {
        fmt.Println(v42, v3)
        fmt.Println(v3, v47, v3)
        v47 = v47
    }
    v81 := GenS7{F0: v47, F1: -1, F2: true}
    var v82 float64 = (float64(v3) - 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v81
    _ = v82
}

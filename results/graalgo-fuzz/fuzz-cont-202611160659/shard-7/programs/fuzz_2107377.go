package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 int }
type GenS8 struct { F0 int32; F1 int32 }
func genFn0(p0 uint16, p1 int) int32 {
    v5 := []int{p1, 0, p1}
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint64 = 299
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: 42, F1: int16(fi8)}
                _ = fi8
                _ = v9
            }
        } else {
            var v10 int16 = -557
            fmt.Println(p0, p0)
            v5 = append(v5, p1)
            _ = v10
        }
        switch ("" + "foo bar") {
        case "foo bar":
            v11 := GenS2{F0: (fi6 + fi6), F1: (("foo bar" + "hello") + "hello")}
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float32 = -2.25
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 4294967295
    p0 -= (p0 - (p0 & (p0 >> ((1) & 63))))
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return -289
}
func genFn1(p0 int8) uint8 {
    switch 32767 {
    case 1:
        p0 = p0
    case 565:
        v15 := GenS2{F0: -278, F1: "a"}
        _ = v15
    default:
        if (false && (!(4294967295 < 18446744073709551615))) {
            v16 := []int{32767, -624, 127, 1}
            var v17 uint16 = 1
            fmt.Println(v17)
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 string = "hello"
                _ = fi18
                _ = v19
            }
            v20 := []int{}
            _ = v20
        }
    }
    var v21 int64 = 1
    _ = p0
    _ = v21
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 uint32 = 65535
    var v4 bool = false
    switch 42 {
    case -9223372036854775808:
        var v22 float32 = 0.0
        _ = v22
    case 42:
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint32 = 207
                if (1 > 2) {
                    v25 = v3
                    var v26 uint8 = 118
                    var v27 int32 = 2147483647
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v3, v4, v2)
                    v25 -= v3
                    var v28 int8 = 106
                    v29 := v1["\n\t"]
                    _ = v28
                    _ = v29
                }
                v25 = (v25 % (((^(v3))) | 1))
                _ = fi24
                _ = v25
            }
            if v4 {
                if v4 {
                    var v30 uint = 538
                    _ = v30
                } else {
                    v31 := GenS3{F0: (^(fi23)), F1: -106}
                    v32 := []int{(^(fi23)), (fi23 >> ((2) & 63)), fi23, (fi23 | fi23)}
                    _ = v31
                    _ = v32
                }
                delete(v1, "")
                var v33 uint = 250
                _ = v33
            }
            _ = fi23
        }
    default:
        if true {
            var v34 int32 = genFn0(32767, -9223372036854775808)
            _ = v34
        } else {
            var v35 uint8 = 127
            v36 := []int{}
            v4 = v4
            _ = v35
            _ = v36
        }
    }
    v37 := v1[""]
    if (-1 > -1) {
        v2 = v2
        var v38 int8 = 1
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        switch fi39 {
        case -1:
            switch "foo bar" {
            case "hello":
                var v40 uint = 14
                _ = v40
            case "foo bar":
                fmt.Println(fi39, v2, v4)
            }
        default:
            v4 = true
        }
        v41 := GenS4{F0: "a"}
        _ = fi39
        _ = v41
    }
    fmt.Println(v4, v3, v37)
    switch -364 {
    case 128:
        v37 = (v37 && (-921 > 42))
    case -1:
        v42 := GenS5{F0: 42}
        _ = v42
    case 9223372036854775807:
        v2 -= v2
    default:
        var v43 int32 = 1
        _ = v43
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        for fi45 := 0; fi45 < 19; fi45++ {
            v46 := GenS1{F0: (fi44 + fi44), F1: -32768}
            var v47 float32 = (-2.25 - 0.0)
            var v48 uint8 = 255
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        v2 += 0
        v49 := v1["a"]
        _ = fi44
        _ = v49
    }
    if true {
        if (v37 && (-32768 >= 32767)) {
            if true {
                v3 -= v3
            } else {
                v4 = v37
            }
        } else {
            v2 -= -9223372036854775808
            v50 := GenS6{F0: -128, F1: 42}
            _ = v50
        }
    }
    v2 = (v2 << ((1) & 63))
    v51 := GenS7{F0: 128, F1: 1, F2: 358}
    var v52 uint64 = 0
    var v53 int16 = -32768
    var v54 float64 = 0.1
    if false {
        if (42 != -778) {
            switch -2147483648 {
            case -436:
                v4 = (!true)
            default:
                var v55 uint = 1
                _ = v55
            }
            var v56 int16 = (v53 << ((2) & 63))
            _ = v56
        }
        v52 -= 691
    } else {
        if ((v37 && (127 == -117)) || v37) {
            if (1 == -1) {
                v2 = v2
            } else {
                var v57 int8 = -1
                switch -9223372036854775808 {
                case -524:
                    if (-544 >= genFn0(32767, 128)) {
                        v58 := GenS8{F0: 1, F1: -1}
                        var v59 int8 = (-(v57))
                        _ = v58
                        _ = v59
                    }
                case -587:
                    switch "a" {
                    case "hello":
                        fmt.Println(v53)
                    }
                case 42:
                    v3 = v3
                default:
                    v3 -= 42
                }
                for fi60 := 0; fi60 < 13; fi60++ {
                    fmt.Println(v53)
                    _ = fi60
                }
                _ = v57
            }
        } else {
            v61 := GenS4{F0: ""}
            var v62 string = "hello"
            _ = v61
            _ = v62
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}

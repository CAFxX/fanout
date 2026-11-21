package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int8; F2 int }
func genFn0() int64 {
    var v4 string = ("foo bar" + ("\n\t" + "foo bar"))
    v5 := []int{-191}
    v4 = ("\n\t" + (((v4 + "") + v4) + (v4 + "\n\t")))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (427 <= 42) {
        if ((true || true) && false) {
            var v6 int32 = -2147483648
            for fi7 := 0; fi7 < 8; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := GenS1{F0: 42, F1: 4294967295, F2: 255}
                    var v10 int8 = 42
                    if ("foo bar" < "\n\t") {
                        var v11 int64 = 42
                        _ = v11
                    }
                    v6 = (v6 + v6)
                    fmt.Println(v6)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v6 += v6
                v6 += (v6 - (v6 | 318))
                _ = fi7
            }
            switch v4 {
            case "a":
                v4 = ("hello" + (v4 + ("hello" + (("" + ((((v4 + (v4 + ((v4 + (v4 + v4)) + "foo bar"))) + ((((("hello" + "foo bar") + "a") + v4) + v4) + ("hello" + v4))) + v4) + "")) + "\n\t"))))
            }
            _ = v6
        }
        _ = len(v5)
    }
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 uint16, p2 int32) int8 {
    if (p2 > (p2 >> ((324) & 63))) {
        var v12 uint64 = (p0 - (p0 / (((p0 + p0)) | 1)))
        var v13 uint32 = 2
        switch int(p2) {
        case -895:
            fmt.Println(p0)
        case 128:
            fmt.Println(v13, p2)
        }
        _ = v12
        _ = v13
    } else {
        var v14 uint8 = 33
        _ = v14
    }
    p2 -= (p2 >> ((0) & 63))
    v15 := GenS1{F0: (p1 + (p1 - p1)), F1: 973, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -1
}

func main() {
    var v1 uint8 = 255
    v2 := []int{-9223372036854775808, 128}
    v3 := []string{}
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            v1 += v1
            if false {
                v18 := GenS2{F0: 440, F1: v1}
                var v19 uint8 = (v1 + 42)
                _ = v18
                _ = v19
            }
            v2 = append(v2, -283)
            _ = fi17
        }
        if false {
            switch "" {
            case "hello":
                fmt.Println(fi16, v1)
            case "a":
                fmt.Println(fi16, v1)
            default:
                var v20 string = "a"
                _ = v20
            }
            var v21 int32 = 0
            _ = v21
        } else {
            fmt.Println(v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = (v1 | v1)
        }
        fmt.Println(v1, v1)
        _ = fi16
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS3{F0: -612}
        v1 -= v1
        if true {
            v1 = (v1 & v1)
        } else {
            var v24 int64 = genFn0()
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS4{F0: -1}
        v27 := GenS2{F0: 1, F1: v1}
        _ = fi25
        _ = v26
        _ = v27
    }
    if false {
        fmt.Println(v1)
        var v28 uint8 = (v1 & (v1 ^ v1))
        _ = v28
    }
    v1 = (^(v1))
    if false {
        var v29 int64 = 658
        var v30 bool = ((v1 >> ((606) & 63)) >= 255)
        if v30 {
            if (1 < 0) {
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS3{F0: -1}
                    _ = fi31
                    _ = v32
                }
            } else {
                var v33 int16 = -1
                var v34 float64 = (3.14159 - -2.25)
                _ = v33
                _ = v34
            }
            v29 = (v29 % ((genFn0()) | 1))
            fmt.Println(v30, v30, v29)
        }
        _ = v29
        _ = v30
    } else {
        switch (((("\n\t" + "") + "") + "a") + "foo bar") {
        case "\n\t":
            _ = len(v2)
        case "hello":
            v35 := GenS4{F0: -2147483648}
            _ = v35
        default:
            var v36 int16 = -1
            _ = v36
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            fmt.Println(v1, v1)
            v38 := GenS5{F0: "hello", F1: 2147483648, F2: -1}
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 7; fi39++ {
            var v40 float32 = 1.5
            switch 326 {
            case 265:
                var v41 uint8 = 0
                _ = v41
            case 32767:
                var v42 uint32 = 42
                _ = v42
            case 1:
                v1 += (^(v1))
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if (false && (!true)) {
                v43 := GenS6{F0: 733, F1: uint32(v1), F2: 1}
                _ = v43
            } else {
                switch fi39 {
                case 127:
                    for fi44 := 0; fi44 < 3; fi44++ {
                        v1 -= (v1 - v1)
                        _ = len(v2)
                        _ = fi44
                    }
                case -9223372036854775808:
                    fmt.Println(v40)
                default:
                    for fi45 := 0; fi45 < 9; fi45++ {
                        fmt.Println(fi45, v1, v40)
                        var v46 uint32 = 42
                        _ = fi45
                        _ = v46
                    }
                }
            }
            _ = fi39
            _ = v40
        }
    }
    v1 = (v1 | v1)
    v47 := GenS7{F0: 1, F1: 100, F2: 4294967295}
    v48 := GenS1{F0: 279, F1: 371, F2: 874}
    var v49 int64 = 1
    v50 := GenS3{F0: 841}
    if (826 <= 0) {
        for fi51 := 0; fi51 < 9; fi51++ {
            fmt.Println(v49, v1)
            _ = fi51
        }
        if (!(1 > 42)) {
            var v52 float32 = -2.25
            _ = v52
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 int64 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
type GenS7 struct { F0 string }
func genFn0(p0 int16, p1 string) uint32 {
    switch 131 {
    case -391:
        if (-2147483648 != -970) {
            var v5 int = -9223372036854775808
            switch (v5 % (((v5 ^ 128)) | 1)) {
            case 4294967295:
                p0 = 32767
            case -1:
                v6 := []int{(v5 % (((v5 + v5)) | 1)), (-(v5)), -442}
                _ = v6
            }
            _ = v5
        }
    case 4294967295:
        v7 := []int{-1}
        _ = v7
    default:
        fmt.Println(p0)
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 uint64, p1 uint32, p2 uint8) int {
    p0 = p0
    p2 += p2
    var v9 float32 = (float32(1e10) + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 int8 = -121
    for fi10 := 0; fi10 < 8; fi10++ {
        if false {
            v4 += -128
            switch "" {
            case "hello":
                var v11 int = fi10
                _ = v11
            case "foo bar":
                v4 += (v4 * (^(v4)))
            case "\n\t":
                fmt.Println(v4)
            default:
                v4 = v4
            }
        } else {
            var v12 uint8 = 22
            if false {
                v13 := GenS1{F0: (v12 ^ (v12 | v12)), F1: 100, F2: 65535}
                fmt.Println(fi10, v4, fi10)
                v14 := []int{}
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: (^(v12)), F1: (v4 + v4), F2: 4294967295}
            _ = v12
            _ = v15
        }
        _ = len(v2)
        fmt.Println(fi10, v4)
        _ = fi10
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 += (^(v4))
        switch "a" {
        case "hello":
            v4 = -1
        }
        _ = fi16
    }
    if true {
        v17 := []int{}
        switch (("\n\t" + "\n\t") + (("hello" + ("\n\t" + "hello")) + "a")) {
        case "foo bar":
            switch "hello" {
            case "\n\t":
                var v18 float32 = 0.1
                _ = v18
            }
        case "":
            v19 := GenS1{F0: 65, F1: 42, F2: 1}
            _ = v19
        default:
            var v20 uint32 = 42
            _ = v20
        }
        if false {
            var v21 float32 = 1.5
            fmt.Println(v21)
            _ = v21
        } else {
            var v22 int8 = v4
            fmt.Println(v22)
            switch "\n\t" {
            case "a":
                for fi23 := 0; fi23 < 14; fi23++ {
                    fmt.Println(fi23)
                    _ = fi23
                }
            case "\n\t":
                var v24 uint8 = 42
                _ = v24
            default:
                _ = len(v2)
            }
            _ = v22
        }
        _ = v17
    } else {
        for fi25 := 0; fi25 < 5; fi25++ {
            v4 = v4
            v26 := []int{(fi25 & -876)}
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := []int{genFn1(4294967295, genFn0(-530, "hello"), 255), 9223372036854775807}
            v4 = v4
            _ = fi27
            _ = v28
        }
    }
    v29 := []int{1}
    var v30 uint32 = 796
    if true {
        var v31 int8 = (v4 - v4)
        for fi32 := 0; fi32 < 12; fi32++ {
            if false {
                fmt.Println(fi32, v4)
            } else {
                v33 := GenS1{F0: 26, F1: int8(fi32), F2: v30}
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 int = genFn1(18446744073709551615, (v30 << ((979) & 63)), 2)
                    _ = fi34
                    _ = v35
                }
                var v36 int64 = 921
                _ = v33
                _ = v36
            }
            for fi37 := 0; fi37 < 12; fi37++ {
                _ = len(v29)
                var v38 int16 = 1
                fmt.Println(fi32, v38)
                v39 := GenS2{F0: 769}
                v4 += (v31 + v31)
                _ = fi37
                _ = v38
                _ = v39
            }
            var v40 uint64 = 18446744073709551615
            _ = fi32
            _ = v40
        }
        _ = v31
    }
    if len(v29) > 0 {
        v29[0] = 127
    }
    if (!false) {
        fmt.Println(v30, v30)
        v4 = v4
        v30 = v30
    }
    v41 := GenS3{F0: 18446744073709551615, F1: -1, F2: 9223372036854775807}
    fmt.Println(v4, v4)
    if false {
        if false {
            var v42 int64 = 42
            if (!(!(2147483647 >= -1))) {
                for fi43 := 0; fi43 < 15; fi43++ {
                    _ = len(v29)
                    v30 -= (v30 / ((4294967295) | 1))
                    v30 = (v30 & (v30 * 42))
                    _ = fi43
                }
                var v44 uint16 = 0
                v45 := GenS3{F0: uint64(v30), F1: 0, F2: 1}
                _ = v44
                _ = v45
            } else {
                var v46 bool = false
                _ = v46
            }
            _ = v42
        } else {
            if len(v29) > 0 {
                v29[0] = -1
            }
        }
        v47 := GenS3{F0: 2, F1: -2147483648, F2: 0}
        v30 = genFn0(32767, "foo bar")
        _ = v47
    }
    if (42 <= 32768) {
        v4 = 0
        for fi48 := 0; fi48 < 18; fi48++ {
            fmt.Println(fi48, fi48)
            if false {
                v30 += genFn0(694, "hello")
                var v49 int16 = 566
                _ = v49
            }
            v50 := GenS4{F0: 2147483647, F1: 0, F2: 65535}
            _ = fi48
            _ = v50
        }
    } else {
        if (true && false) {
            for fi51 := 0; fi51 < 17; fi51++ {
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                fmt.Println(fi51)
                var v52 string = (("hello" + (("hello" + "\n\t") + "a")) + ("a" + ("hello" + ("\n\t" + ("" + "")))))
                _ = fi51
                _ = v52
            }
            fmt.Println(v4)
        } else {
            if (false && true) {
                if len(v29) > 0 {
                    v29[0] = 9223372036854775807
                }
                v29 = append(v29, 4294967295)
                switch (("hello" + ("hello" + ("foo bar" + ("a" + "")))) + ("\n\t" + "")) {
                case "\n\t":
                    v30 = 0
                }
            } else {
                fmt.Println(v30, v30)
            }
            switch -926 {
            case 468:
                var v53 float64 = 0.0
                _ = v53
            case 0:
                v30 = genFn0(42, "a")
            default:
                var v54 uint16 = 255
                _ = v54
            }
        }
    }
    v55 := GenS5{F0: 2147483647}
    var v56 int32 = -1
    if (v30 != (v30 % ((v30) | 1))) {
        if (true || false) {
            var v57 uint = 4294967295
            for fi58 := 0; fi58 < 6; fi58++ {
                switch "foo bar" {
                case "foo bar":
                    v56 += (v56 - v56)
                default:
                    var v59 float64 = (1e10 / -2.25)
                    _ = v59
                }
                v60 := GenS3{F0: 0, F1: v56, F2: 42}
                _ = fi58
                _ = v60
            }
            var v61 int = 4294967295
            _ = v57
            _ = v61
        } else {
            v4 += (^(v4))
            if (false || (!(1 != -1))) {
                var v62 bool = false
                _ = v62
            } else {
                var v63 uint16 = 109
                _ = v63
            }
        }
        for fi64 := 0; fi64 < 15; fi64++ {
            v65 := GenS6{F0: true, F1: 217, F2: ""}
            if false {
                if false {
                    v30 = v30
                    var v66 int = genFn1(2, 168, 238)
                    v67 := GenS5{F0: v56}
                    _ = v66
                    _ = v67
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v29)
                    case "a":
                        v68 := GenS7{F0: "hello"}
                        _ = v68
                    default:
                        var v69 uint = 65535
                        _ = v69
                    }
                }
            }
            _ = fi64
            _ = v65
        }
        if true {
            var v70 bool = (3.14159 <= float64(3.14159))
            if ((!v70) && (0 >= 0)) {
                for fi71 := 0; fi71 < 20; fi71++ {
                    switch fi71 {
                    case 817:
                        delete(v2, fi71)
                    }
                    _ = fi71
                }
                for fi72 := 0; fi72 < 12; fi72++ {
                    fmt.Println(v30, v30, v70)
                    _ = fi72
                }
            } else {
                fmt.Println(v4, v30, v4)
            }
            if v70 {
                switch (((("" + "foo bar") + "foo bar") + ("hello" + "hello")) + "hello") {
                case "\n\t":
                    var v73 int8 = (v4 / ((int8(v30)) | 1))
                    _ = v73
                case "a":
                    var v74 uint16 = 65535
                    _ = v74
                default:
                    for fi75 := 0; fi75 < 9; fi75++ {
                        v4 -= v4
                        if len(v29) > 0 {
                            v29[0] = 4294967295
                        }
                        _ = fi75
                    }
                }
                fmt.Println(v4, v30, v56)
            }
            _ = v70
        }
    } else {
        for fi76 := 0; fi76 < 8; fi76++ {
            v56 = (v56 & v56)
            fmt.Println(v56, v30)
            for fi77 := 0; fi77 < 11; fi77++ {
                v78 := GenS7{F0: (("" + "") + ("\n\t" + ("foo bar" + "\n\t")))}
                _ = fi77
                _ = v78
            }
            _ = fi76
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v55
    _ = v56
}

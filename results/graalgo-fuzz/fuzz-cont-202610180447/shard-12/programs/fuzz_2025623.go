package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 bool) string {
    p0 = (true || false)
    p0 = (2147483648 == 557)
    v7 := []int{1, 127}
    p0 = p0
    p0 = p0
    _ = p0
    _ = v7
    return "a"
}
func genFn1(p0 bool) uint {
    for fi8 := 0; fi8 < 3; fi8++ {
        switch (-(fi8)) {
        case 2147483647:
            v9 := GenS1{F0: ((0.0 * 1.5) < 0.0), F1: 4294967295, F2: 2}
            _ = v9
        case 1:
            var v10 string = ""
            _ = v10
        case -876:
            switch (fi8 >> ((4294967295) & 63)) {
            case 9223372036854775807:
                if false {
                    p0 = p0
                } else {
                    fmt.Println(p0, fi8)
                    var v11 uint = 980
                    _ = v11
                }
            default:
                p0 = (!false)
            }
        }
        _ = fi8
    }
    if ((1e10 != ((((1.5 + (3.14159 * float64(float32((((float64((float32(-2.25) * 0.0)) - (-2.25 * (3.14159 + ((float64(1e10) * 0.0) - ((1e10 - 0.1) + 3.14159))))) - ((-2.25 / 1.0) / (-2.25 * 3.14159))) + 0.0))))) + (((0.1 * float64(3.14159)) * -2.25) / 1e10)) - (1e10 / (float64(-2.25) * 1e10))) + (((((0.1 / float64(float32(-2.25))) * 0.1) + (((3.14159 + float64(1.5)) * (float64(3.14159) * float64(3.14159))) / 3.14159)) / 3.14159) / 0.1))) && (1.5 == (0.1 * 0.0))) {
        p0 = false
        v12 := []int{}
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 float64 = 3.14159
            p0 = p0
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v15 := GenS1{F0: true, F1: 2, F2: 2}
    var v16 uint32 = 172
    v17 := GenS1{F0: p0, F1: 537, F2: 0}
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (false && (((100 <= 127) && true) && (false || false)))
    v3 := make(map[uint32]int)
    var v4 int = 32768
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    if (!(499 > 4294967295)) {
        v2 = (!(1 == 819))
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS2{F0: 2, F1: (!(!((0 == 42) || (v2 || v2))))}
            _ = fi18
            _ = v19
        }
    }
    switch "\n\t" {
    case "hello":
        var v20 float32 = float32(1e10)
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 1; fi21++ {
            delete(v6, 142)
            _ = fi21
        }
    case "foo bar":
        v22 := GenS1{F0: v2, F1: 1, F2: 2}
        _ = v22
    }
    var v23 bool = true
    v4 -= 127
    if true {
        v24 := []int{}
        v25 := v5[376]
        for fi26 := 0; fi26 < 13; fi26++ {
            v25 -= (v25 % ((v4) | 1))
            if false {
                v27 := GenS3{F0: 42}
                var v28 float64 = -2.25
                switch 70 {
                case -2147483648:
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v28 = 0.1
                        v30 := v1[2147483647]
                        v2 = true
                        _ = fi29
                        _ = v30
                    }
                case 527:
                    if v2 {
                        fmt.Println(v28, v4)
                        if (4294967295 > 1) {
                            for fi31 := 0; fi31 < 19; fi31++ {
                                v4 += (v25 >> ((22) & 63))
                                v32 := GenS2{F0: 128, F1: true}
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            v25 = v4
                            delete(v3, 2)
                            var v33 string = ((("" + "") + ("" + "foo bar")) + "")
                            _ = v33
                        }
                    }
                }
                _ = v27
                _ = v28
            }
            if (42 > 4294967295) {
                var v34 int32 = -551
                _ = v34
            } else {
                v35 := GenS4{F0: "a"}
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 float32 = (((3.14159 - 0.0) / 1.0) + -2.25)
                    v3[2147483648] = (fi26 % (((v4 - v25)) | 1))
                    v5[0] = (fi26 + (v25 * 127))
                    _ = fi36
                    _ = v37
                }
                v38 := GenS3{F0: 2}
                _ = v35
                _ = v38
            }
            _ = fi26
        }
        _ = v24
        _ = v25
    } else {
        switch (("hello" + "") + "\n\t") {
        case "hello":
            fmt.Println(v4)
        case "a":
            if true {
                v39 := GenS4{F0: "foo bar"}
                if (!v23) {
                    if v2 {
                        if true {
                            fmt.Println(v23, v23)
                        }
                        var v40 uint64 = 1
                        _ = v40
                    }
                }
                _ = v39
            } else {
                var v41 bool = false
                var v42 uint32 = 31
                if v41 {
                    v43 := GenS4{F0: "hello"}
                    switch "foo bar" {
                    case "":
                        var v44 uint = 368
                        _ = v44
                    case "hello":
                        if true {
                            var v45 int = (v4 | (v4 << ((984) & 63)))
                            _ = v45
                        }
                    }
                    _ = v43
                }
                _ = v41
                _ = v42
            }
        case "foo bar":
            fmt.Println(v2)
        }
        var v46 int16 = 0
        switch (-(v4)) {
        case 128:
            for fi47 := 0; fi47 < 18; fi47++ {
                v48 := GenS5{F0: (v4 - v4), F1: 4294967295}
                delete(v5, 4294967295)
                for fi49 := 0; fi49 < 3; fi49++ {
                    var v50 int64 = -1
                    delete(v3, 0)
                    var v51 uint32 = 1
                    if (fi49 < (fi49 & v4)) {
                        if (100 == -97) {
                            v52 := GenS5{F0: (fi49 / ((v4) | 1)), F1: 2}
                            var v53 bool = false
                            fmt.Println(fi49, v23, fi47)
                            _ = v52
                            _ = v53
                        }
                        fmt.Println(fi47, v4, v51)
                    }
                    var v54 uint32 = 228
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v54
                }
                _ = fi47
                _ = v48
            }
        }
        _ = v46
    }
    v55 := GenS6{F0: 640, F1: ""}
    var v56 bool = (0.1 >= 3.14159)
    v57 := []int{(v4 >> ((786) & 63))}
    _ = len(v6)
    switch v4 {
    case 843:
        fmt.Println(v23)
    case -558:
        _ = len(v57)
    default:
        v57 = append(v57, -9223372036854775808)
    }
    v23 = ((132 <= 255) && (false || (26 >= 1)))
    switch v4 {
    case 372:
        fmt.Println(v56, v23)
    case 209:
        v4 = (v4 + (-(v4)))
    }
    if v56 {
        fmt.Println(v23)
        if true {
            for fi58 := 0; fi58 < 10; fi58++ {
                fmt.Println(v4, v56, fi58)
                _ = fi58
            }
        } else {
            if false {
                var v59 uint8 = 214
                var v60 uint16 = 255
                for fi61 := 0; fi61 < 10; fi61++ {
                    v57 = append(v57, (-(fi61)))
                    if v2 {
                        v59 = (v59 % ((v59) | 1))
                    } else {
                        fmt.Println(v56, v4)
                        v57 = append(v57, 32767)
                    }
                    v59 = v59
                    _ = fi61
                }
                _ = v59
                _ = v60
            } else {
                if (100 >= 91) {
                    var v62 string = "\n\t"
                    _ = v62
                }
                switch v4 {
                case 9223372036854775807:
                    v63 := GenS7{F0: 2, F1: 42, F2: 4294967295}
                    _ = v63
                default:
                    v2 = (((false || (1000 > -1)) || v23) || ((v4 >> ((0) & 63)) != v4))
                }
            }
            for fi64 := 0; fi64 < 3; fi64++ {
                fmt.Println(v4)
                _ = fi64
            }
            var v65 string = ""
            _ = v65
        }
        var v66 float32 = (float32(3.14159) - float32((float64(1e10) + 1.5)))
        _ = v66
    } else {
        v67 := GenS2{F0: 244, F1: (933 == 9223372036854775807)}
        if len(v57) >= 2 {
            v57 = v57[1:len(v57)]
        }
        _ = v67
    }
    v23 = v23
    v2 = (!v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}

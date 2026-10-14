package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint; F1 uint8; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int; F2 int }
type GenS11 struct { F0 uint; F1 int32; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 int8) uint {
    p2 += (p2 >> ((4294967295) & 63))
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1() int {
    for fi4 := 0; fi4 < 11; fi4++ {
        switch fi4 {
        case 128:
            v5 := GenS1{F0: 42}
            _ = v5
        default:
            var v6 float32 = (1.5 - ((float32(0.0) - (-2.25 * (1.5 - (1.5 / (0.1 + (0.0 - (0.0 / (float32(3.14159) + 3.14159)))))))) + float32(1.5)))
            _ = v6
        }
        if (255 > 65535) {
            var v7 uint16 = 255
            var v8 float64 = (1.5 + (0.0 + 0.0))
            _ = v7
            _ = v8
        } else {
            if true {
                v9 := []int{fi4, fi4}
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch -27 {
                    case 841:
                        var v11 int = -2147483648
                        _ = v11
                    case 1:
                        if (42 < 9223372036854775807) {
                            var v12 float32 = ((float32(-2.25) / 3.14159) + float32(1e10))
                            switch ((("" + "hello") + (("\n\t" + ("" + "\n\t")) + "a")) + ("hello" + ("hello" + "foo bar"))) {
                            case "":
                                var v13 uint8 = 2
                                _ = v13
                            case "hello":
                                if (false && (v12 > (v12 * v12))) {
                                    var v14 int64 = -1
                                    var v15 int = -9223372036854775808
                                    var v16 float64 = 1e10
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    switch 32768 {
                                    case -9223372036854775808:
                                        v12 = v12
                                    case 4294967295:
                                        _ = len(v9)
                                    case 2147483647:
                                        v17 := GenS1{F0: -619}
                                        _ = v17
                                    }
                                    v9 = append(v9, (fi10 >> ((65535) & 63)))
                                }
                            default:
                                var v18 float32 = v12
                                _ = v18
                            }
                            _ = v12
                        }
                    }
                    switch (fi10 & fi10) {
                    case 378:
                        switch (fi10 << ((255) & 63)) {
                        case 42:
                            var v19 bool = (((3.14159 / 1e10) / float64(-2.25)) != -2.25)
                            _ = v19
                        default:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                    }
                    var v20 bool = true
                    v21 := GenS2{F0: -229, F1: -504, F2: 65535}
                    _ = fi10
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi4)
                _ = v9
            } else {
                var v22 uint32 = 65535
                for fi23 := 0; fi23 < 9; fi23++ {
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        var v24 string = "\n\t"
                        _ = v24
                    case "a":
                        switch (fi23 & fi23) {
                        case -119:
                            v22 = 42
                        }
                    case "foo bar":
                        v22 -= v22
                    }
                    _ = fi23
                }
                fmt.Println(fi4)
                _ = v22
            }
            switch (fi4 / ((fi4) | 1)) {
            case 32768:
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 int = fi25
                    switch ((("\n\t" + ("\n\t" + "foo bar")) + ("hello" + "\n\t")) + (("foo bar" + "hello") + "hello")) {
                    case "\n\t":
                        v26 = fi25
                    case "foo bar":
                        var v27 bool = false
                        _ = v27
                    case "":
                        var v28 int8 = 100
                        _ = v28
                    default:
                        v29 := GenS1{F0: 1}
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                }
            case 32767:
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 string = "a"
                    var v32 uint8 = 87
                    v31 = v31
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            default:
                var v33 uint = 2
                _ = v33
            }
        }
        _ = fi4
    }
    var v34 int64 = -9223372036854775808
    _ = v34
    return -176
}

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 0
    for fi35 := 0; fi35 < 18; fi35++ {
        if true {
            if true {
                v36 := GenS3{F0: "", F1: (42 <= 4294967295)}
                if false {
                    fmt.Println(v3, fi35)
                } else {
                    if true {
                        for fi37 := 0; fi37 < 20; fi37++ {
                            fmt.Println(v3, fi37, fi37)
                            v38 := GenS4{F0: ((v3 >> ((65535) & 63)) == v3), F1: "\n\t"}
                            var v39 int = (-(fi35))
                            var v40 uint32 = 1
                            _ = fi37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        if (((fi35 > 2147483647) && true) || true) {
                            v3 += (v3 | (^(v3)))
                        }
                        for fi41 := 0; fi41 < 10; fi41++ {
                            if false {
                                v3 = (v3 << ((375) & 63))
                            }
                            v3 = (v3 | 128)
                            _ = fi41
                        }
                    } else {
                        var v42 string = "\n\t"
                        _ = v42
                    }
                    v43 := GenS5{F0: 582, F1: v3, F2: (!(fi35 <= fi35))}
                    switch "hello" {
                    case "hello":
                        v2[1] = 2
                    case "a":
                        switch (fi35 | fi35) {
                        case -2147483648:
                            v3 = v3
                        case -1:
                            v3 -= (v3 + v3)
                        case 9223372036854775807:
                            switch 201 {
                            case 0:
                                var v44 int = genFn1()
                                _ = v44
                            case -767:
                                for fi45 := 0; fi45 < 11; fi45++ {
                                    fmt.Println(fi45, fi45, v3)
                                    v3 = v3
                                    _ = fi45
                                }
                            case 1:
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    if true {
                                        delete(v2, 819)
                                        fmt.Println(fi46)
                                        v3 = (v3 & (v3 & (v3 * v3)))
                                    } else {
                                        v3 = 42
                                    }
                                    var v47 int = fi46
                                    v48 := GenS3{F0: ("\n\t" + "foo bar"), F1: (0 <= 4294967295)}
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            default:
                                switch "foo bar" {
                                case "":
                                    for fi49 := 0; fi49 < 19; fi49++ {
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v50 int8 = 127
                                            _ = v50
                                        case "a":
                                            if false {
                                                var v51 uint16 = 65535
                                                _ = v51
                                            } else {
                                                var v52 float64 = (3.14159 + float64((0.0 + -2.25)))
                                                var v53 uint16 = 225
                                                var v54 uint16 = v53
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        default:
                                            var v55 float32 = 0.1
                                            _ = v55
                                        }
                                        fmt.Println(fi35)
                                        var v56 float64 = -2.25
                                        _ = fi49
                                        _ = v56
                                    }
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = 819
                                    }
                                }
                            }
                        default:
                            v3 = (^(v3))
                        }
                    case "":
                        var v57 uint64 = 0
                        _ = v57
                    }
                    _ = v43
                }
                _ = v36
            } else {
                switch genFn1() {
                case 42:
                    v58 := GenS6{F0: -1}
                    _ = v58
                default:
                    for fi59 := 0; fi59 < 8; fi59++ {
                        v3 -= (v3 * v3)
                        v60 := GenS7{F0: 55}
                        _ = fi59
                        _ = v60
                    }
                }
                fmt.Println(v3, v3)
                if false {
                    var v61 int = (fi35 / ((32767) | 1))
                    _ = v61
                } else {
                    var v62 int16 = 42
                    switch (fi35 << ((uint(v3)) & 63)) {
                    case 42:
                        v2[4294967295] = 65535
                    case -589:
                        switch (("" + "") + "hello") {
                        case "":
                            v62 += v62
                        case "\n\t":
                            v62 = (^(v62))
                        }
                    default:
                        v2[4294967295] = 611
                    }
                    _ = v62
                }
            }
        } else {
            if true {
                v63 := GenS8{F0: 42, F1: 0}
                switch (^(fi35)) {
                case -857:
                    for fi64 := 0; fi64 < 12; fi64++ {
                        switch (("hello" + (("foo bar" + (("hello" + "foo bar") + ("" + "\n\t"))) + "\n\t")) + "") {
                        case "foo bar":
                            v3 -= (v3 & (v3 ^ v3))
                        case "hello":
                            fmt.Println(fi35, v3, v3)
                        }
                        var v65 int16 = 1000
                        _ = fi64
                        _ = v65
                    }
                }
                _ = v63
            } else {
                if (0 < -1) {
                    var v66 uint = 403
                    fmt.Println(v66, fi35, v3)
                    fmt.Println(v3, v3, v3)
                    _ = v66
                }
                for fi67 := 0; fi67 < 9; fi67++ {
                    v68 := GenS9{F0: 0, F1: 386}
                    if (4294967295 <= 2) {
                        v3 = v3
                        fmt.Println(v3)
                    }
                    _ = fi67
                    _ = v68
                }
            }
            switch "hello" {
            case "\n\t":
                if false {
                    v3 += (v3 + 127)
                    fmt.Println(v3, fi35)
                }
            case "hello":
                var v69 bool = ((!(v3 >= 1)) || false)
                _ = v69
            case "":
                if ((0.1 + -2.25) == 0.0) {
                    v70 := GenS10{F0: 255, F1: fi35, F2: (fi35 & -416)}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v70
                }
            default:
                _ = len(v2)
            }
        }
        if (0 > 644) {
            v3 += 145
            switch fi35 {
            case 334:
                v3 -= (v3 % ((v3) | 1))
            case 0:
                switch "foo bar" {
                case "":
                    var v71 int32 = 42
                    _ = v71
                case "foo bar":
                    _ = len(v2)
                case "\n\t":
                    switch fi35 {
                    case 9223372036854775807:
                        for fi72 := 0; fi72 < 2; fi72++ {
                            fmt.Println(v3, v3, v3)
                            v3 += (v3 % ((v3) | 1))
                            fmt.Println(fi72, fi35)
                            _ = fi72
                        }
                    }
                }
            }
            var v73 int = (fi35 << ((uint(v3)) & 63))
            _ = v73
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v74 uint8 = v3
            _ = v74
        }
        _ = fi35
    }
    if true {
        v75 := v2[937]
        if (v75 == v75) {
            if (false && (v3 == 115)) {
                switch -982 {
                case 4294967295:
                    fmt.Println(v3, v75)
                case 630:
                    v2[v75] = (v75 * 1)
                }
            } else {
                fmt.Println(v3, v75)
                v3 = v3
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v76 string = "foo bar"
            var v77 uint8 = 42
            _ = v76
            _ = v77
        } else {
            v75 -= v75
            if (0.0 != -2.25) {
                var v78 int8 = 0
                switch "foo bar" {
                case "a":
                    var v79 uint16 = 32767
                    _ = v79
                case "":
                    if len(v1) > 0 {
                        v1[0] = (v75 << ((uint(v75)) & 63))
                    }
                case "foo bar":
                    v75 = v75
                }
                _ = v78
            }
        }
        _ = v75
    }
    v3 += v3
    v80 := GenS8{F0: -2147483648, F1: 232}
    v3 -= v3
    v3 -= v3
    var v81 uint64 = 0
    v82 := GenS11{F0: 18446744073709551615, F1: 42, F2: (^(v3))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v80
    _ = v81
    _ = v82
}

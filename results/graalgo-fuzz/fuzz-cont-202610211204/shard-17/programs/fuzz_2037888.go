package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 string; F1 string; F2 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 42
    v3 := make(map[string]string)
    var v4 string = "foo bar"
    var v5 string = "a"
    v6 := make(map[uint32]string)
    v7 := []int{127}
    _ = len(v1)
    v2 += (v2 / (((v2 >> ((393) & 63))) | 1))
    var v8 int8 = 42
    if len(v7) > 0 {
        v7[0] = 128
    }
    fmt.Println(v2, v5, v2)
    if ((42 >= 65535) || true) {
        if false {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v9 int = -138
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint64 = 1
                v5 = "\n\t"
                _ = fi10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                switch (v5 + v5) {
                case "":
                    switch (fi12 % ((-500) | 1)) {
                    case 9223372036854775807:
                        switch fi12 {
                        case 127:
                            var v14 int = 1
                            _ = v14
                        case 9223372036854775807:
                            v5 = ((v5 + v5) + v5)
                            var v15 int = 9223372036854775807
                            v15 = (fi13 - (v15 - -2147483648))
                            _ = v15
                        case 4294967295:
                            var v16 float32 = -2.25
                            _ = v16
                        default:
                            delete(v3, v5)
                            switch v5 {
                            case "\n\t":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            case "":
                                v4 = ((v5 + ((((("foo bar" + ((v4 + (v4 + ((((v4 + ("" + ("hello" + (((("\n\t" + v4) + v4) + v4) + ("a" + v5))))) + "a") + "hello") + ("" + ((v5 + (((v5 + (v4 + "foo bar")) + "") + v5)) + v5))))) + "")) + "\n\t") + ((v5 + v4) + "\n\t")) + v4) + v4)) + v4)
                            }
                            _ = len(v7)
                            fmt.Println(v5)
                        }
                    default:
                        var v17 bool = true
                        _ = v17
                    }
                }
                var v18 uint32 = 1
                var v19 string = ((v4 + (v5 + (v5 + ("foo bar" + v5)))) + ((v5 + "") + ("hello" + "foo bar")))
                switch fi12 {
                case 0:
                    var v20 uint8 = 0
                    _ = v20
                default:
                    fmt.Println(v5, v4)
                    fmt.Println(v8, v18)
                    if (!true) {
                        v21 := GenS1{F0: (v2 - -831), F1: 225, F2: 42}
                        v19 = v4
                        _ = v21
                    }
                }
                v2 = (v2 | (v2 * int32(v8)))
                _ = fi13
                _ = v18
                _ = v19
            }
            if (!true) {
                if true {
                    switch (fi12 | (fi12 * int(v2))) {
                    case 42:
                        for fi22 := 0; fi22 < 14; fi22++ {
                            var v23 int8 = (v8 * 42)
                            v24 := GenS1{F0: (v2 % (((v2 | (v2 * (^(v2))))) | 1)), F1: 19, F2: 975}
                            var v25 int16 = 32767
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    default:
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 uint16 = 939
                            if false {
                                v5 = v5
                                switch (v5 + (((v5 + (v5 + v4)) + "foo bar") + v4)) {
                                case "\n\t":
                                    v27 = (v27 + v27)
                                case "a":
                                    var v28 int64 = -953
                                    _ = v28
                                }
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                } else {
                    if (565 <= 2) {
                        if true {
                            v29 := GenS1{F0: -2147483648, F1: 255, F2: 814}
                            _ = v29
                        } else {
                            var v30 int = (fi12 << ((189) & 63))
                            _ = v30
                        }
                        var v31 int16 = int16(v2)
                        var v32 float32 = float32(1.5)
                        _ = v31
                        _ = v32
                    }
                }
                v33 := GenS1{F0: (v2 | (v2 | (v2 + (v2 | 1)))), F1: 1, F2: 42}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v33
            }
            if ((v4 > ("hello" + ((v4 + ((v4 + v5) + v5)) + v4))) && (true || (((!(0.1 != (1.5 * (1e10 * (float32((1.5 + ((3.14159 - 0.1) + float64(((-2.25 - 0.1) * -2.25))))) - 1.5))))) && (!false)) || (float32(0.0) == ((((float32((float64(1e10) - -2.25)) - ((float32(0.1) * -2.25) + (1e10 / 1.5))) / 1.0) / 0.1) * 0.1))))) {
                fmt.Println(v2, v8)
            } else {
                for fi34 := 0; fi34 < 7; fi34++ {
                    switch (v4 + v5) {
                    case "a":
                        fmt.Println(v5, v4)
                    case "foo bar":
                        switch fi34 {
                        case 1:
                            v8 += (v8 + v8)
                        case 127:
                            v35 := v3[(v5 + v4)]
                            _ = v35
                        default:
                            var v36 int16 = 1000
                            _ = v36
                        }
                    }
                    switch v5 {
                    case "":
                        if true {
                            v37 := GenS1{F0: v2, F1: 4294967295, F2: 1}
                            _ = v37
                        } else {
                            v4 = (v5 + "\n\t")
                            if true {
                                v38 := GenS2{F0: 2}
                                if false {
                                    if true {
                                        if (!true) {
                                            fmt.Println(v4, fi12)
                                            if true {
                                                if false {
                                                    v8 = 51
                                                    if len(v7) > 0 {
                                                        v7[0] = fi12
                                                    }
                                                    fmt.Println(v2, v4, v4)
                                                }
                                                if (352 == 2) {
                                                    var v39 string = v5
                                                    delete(v6, 2)
                                                    switch (v4 + v5) {
                                                    case "a":
                                                        if ((fi12 << ((4294967295) & 63)) <= -662) {
                                                            v40 := GenS3{F0: 0}
                                                            delete(v1, (fi34 + fi34))
                                                            if false {
                                                                v41 := GenS2{F0: 255}
                                                                v2 -= 42
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi12 >> ((65535) & 63))
                                                                }
                                                                _ = v41
                                                            }
                                                            switch (fi34 << ((2) & 63)) {
                                                            case 1:
                                                                if true {
                                                                    v4 = (v39 + "")
                                                                    var v42 float64 = 3.14159
                                                                    _ = len(v7)
                                                                    var v43 string = (v4 + v5)
                                                                    _ = v42
                                                                    _ = v43
                                                                } else {
                                                                    switch ((v5 + v5) + "\n\t") {
                                                                    case "hello":
                                                                        v44 := GenS4{F0: ("" + (("foo bar" + v5) + (v39 + v39)))}
                                                                        _ = v44
                                                                    case "foo bar":
                                                                        v2 = (v2 % ((v2) | 1))
                                                                    default:
                                                                        _ = len(v6)
                                                                    }
                                                                }
                                                            case 754:
                                                                if true {
                                                                    if (false && (255 < 239)) {
                                                                        if false {
                                                                            var v45 uint8 = 0
                                                                            var v46 int = (fi12 & fi12)
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                        _ = len(v3)
                                                                        if ("\n\t" > v5) {
                                                                            v8 = v8
                                                                            v2 = v2
                                                                        }
                                                                        v47 := GenS3{F0: (-(v2))}
                                                                        _ = v47
                                                                    } else {
                                                                        var v48 int8 = (v8 % (((v8 >> ((42) & 63))) | 1))
                                                                        v7 = append(v7, (fi12 >> ((255) & 63)))
                                                                        v5 = "\n\t"
                                                                        v49 := GenS5{F0: 1, F1: 239, F2: 318}
                                                                        v50 := GenS3{F0: v2}
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                    var v51 bool = false
                                                                    _ = v51
                                                                }
                                                                fmt.Println(fi34, fi12, v39)
                                                            }
                                                            _ = v40
                                                        }
                                                    case "hello":
                                                        delete(v3, ((((v4 + (v4 + ("a" + "a"))) + "hello") + v5) + v5))
                                                    }
                                                    v8 = 64
                                                    v1[fi12] = true
                                                    _ = v39
                                                }
                                            } else {
                                                switch (^(fi12)) {
                                                case 32768:
                                                    v5 = v5
                                                case 939:
                                                    v52 := GenS6{F0: 0, F1: v8}
                                                    _ = v52
                                                }
                                                v53 := GenS4{F0: ("" + v5)}
                                                _ = v53
                                            }
                                            var v54 uint64 = 1
                                            fmt.Println(v5)
                                            _ = v54
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = fi12
                                        }
                                    } else {
                                        v5 = v4
                                        v8 = -1
                                        fmt.Println(v2, v8, v2)
                                    }
                                }
                                var v55 uint = 42
                                _ = v38
                                _ = v55
                            }
                            var v56 string = ""
                            _ = v56
                        }
                    case "a":
                        v5 = (v4 + v5)
                    case "hello":
                        switch (fi34 / ((-72) | 1)) {
                        case 128:
                            switch int(v2) {
                            case 0:
                                v57 := GenS2{F0: 0}
                                _ = v57
                            case 541:
                                _ = len(v7)
                            case -2147483648:
                                switch (fi12 * 128) {
                                case 9223372036854775807:
                                    v5 = ("foo bar" + v4)
                                case 2147483647:
                                    switch ("" + v5) {
                                    case "":
                                        var v58 int8 = v8
                                        _ = v58
                                    case "hello":
                                        var v59 float64 = (float64(0.0) * ((((1e10 * 3.14159) * 0.1) / (1.5 + (3.14159 * float64((-2.25 / 3.14159))))) + -2.25))
                                        if false {
                                            var v60 int32 = v2
                                            _ = v60
                                        } else {
                                            v2 = -894
                                        }
                                        _ = v59
                                    case "foo bar":
                                        if (true && ((("" + "\n\t") != ("foo bar" + (v5 + v5))) || false)) {
                                            v61 := GenS3{F0: (v2 % ((-859) | 1))}
                                            _ = v61
                                        }
                                    default:
                                        var v62 string = (v5 + v5)
                                        _ = v62
                                    }
                                case -128:
                                    if false {
                                        var v63 int8 = -1
                                        v64 := GenS3{F0: (v2 << ((255) & 63))}
                                        _ = v63
                                        _ = v64
                                    }
                                default:
                                    v65 := v1[fi34]
                                    _ = v65
                                }
                            default:
                                if false {
                                    if (!false) {
                                        v5 = v4
                                        var v66 int64 = -1
                                        _ = v66
                                    } else {
                                        var v67 int16 = 0
                                        v7 = append(v7, (fi34 % (((fi34 * fi12)) | 1)))
                                        _ = v67
                                    }
                                    v4 = "foo bar"
                                }
                            }
                        case 32767:
                            var v68 int8 = (v8 % ((-42) | 1))
                            _ = v68
                        default:
                            var v69 int32 = v2
                            _ = v69
                        }
                    }
                    _ = fi34
                }
            }
            _ = fi12
        }
        v4 = v5
    } else {
        var v70 int32 = (v2 - 1)
        _ = v70
    }
    v71 := GenS1{F0: v2, F1: 0, F2: -9223372036854775808}
    var v72 string = "\n\t"
    var v73 uint64 = 0
    switch v72 {
    case "foo bar":
        v74 := v1[585]
        _ = v74
    case "\n\t":
        for fi75 := 0; fi75 < 20; fi75++ {
            v5 = (("hello" + "hello") + "foo bar")
            if false {
                var v76 uint32 = 2
                v77 := GenS7{F0: "\n\t", F1: "a", F2: 65535}
                _ = v76
                _ = v77
            }
            v72 = (v72 + ("a" + v4))
            _ = fi75
        }
    case "hello":
        v72 = "foo bar"
    default:
        if false {
            for fi78 := 0; fi78 < 8; fi78++ {
                if (v2 > v2) {
                    if (127 > uint8(v73)) {
                        v5 = "a"
                        fmt.Println(v4)
                    }
                    v8 += v8
                    v73 = v73
                }
                var v79 int32 = (v2 | (v2 * (^(v2))))
                for fi80 := 0; fi80 < 7; fi80++ {
                    v8 = v8
                    fmt.Println(fi78, v79)
                    v7 = append(v7, (fi78 * (fi80 ^ (fi78 / ((fi78) | 1)))))
                    _ = fi80
                }
                _ = fi78
                _ = v79
            }
            if false {
                if false {
                    if true {
                        v81 := GenS8{F0: false}
                        _ = v81
                    }
                    fmt.Println(v5)
                } else {
                    if (!(2147483647 != -2147483648)) {
                        v82 := v6[2]
                        _ = len(v1)
                        _ = v82
                    }
                    fmt.Println(v73, v2)
                }
                var v83 int16 = 0
                _ = v83
            }
        } else {
            v8 = -106
        }
    }
    var v84 bool = (!(32767 >= 32767))
    if v84 {
        v85 := GenS9{F0: 9223372036854775807, F1: (v8 | v8)}
        var v86 int32 = v2
        delete(v3, (("a" + v72) + v5))
        _ = v85
        _ = v86
    }
    for fi87 := 0; fi87 < 10; fi87++ {
        delete(v1, (fi87 | fi87))
        var v88 uint = 4294967295
        v8 -= v8
        _ = fi87
        _ = v88
    }
    var v89 uint64 = v73
    v2 = -1
    v90 := GenS4{F0: v5}
    if (v84 && (32767 > 219)) {
        _ = len(v6)
        switch ((v5 + v72) + v4) {
        case "\n\t":
            fmt.Println(v8)
        default:
            if (46 >= v8) {
                v5 = "foo bar"
            }
        }
    } else {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v8:", v8)
    fmt.Println("v84:", v84)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v72
    _ = v73
    _ = v8
    _ = v84
    _ = v89
    _ = v90
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint }
func genFn0() uint8 {
    v4 := GenS1{F0: 0, F1: -97}
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, fi5, 9223372036854775807, fi5}
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS2{F0: -1, F1: 1000}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: 127, F1: 1000}
            fmt.Println(fi9)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi9
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    var v11 string = "foo bar"
    v12 := GenS2{F0: 60, F1: 32767}
    _ = v11
    _ = v12
    _ = v4
    return 186
}

func main() {
    v1 := []uint32{182}
    var v2 bool = false
    var v3 int32 = 0
    if len(v1) > 0 {
        v1[0] = 42
    }
    if (32767 > 1000) {
        if v2 {
            v3 = v3
            switch "a" {
            case "\n\t":
                if ((v2 || (42 <= 32767)) && true) {
                    if (!v2) {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v3, v3)
                        }
                        var v13 uint8 = 2
                        var v14 uint = 0
                        _ = v13
                        _ = v14
                    }
                    if (0 != 18446744073709551615) {
                        if v2 {
                            v2 = true
                            var v15 string = "a"
                            _ = v15
                        }
                        _ = len(v1)
                    }
                }
            }
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = (642 != 42)
                _ = fi16
            }
        }
        var v17 uint8 = 1
        var v18 uint16 = 65535
        _ = v17
        _ = v18
    } else {
        var v19 uint32 = 1
        _ = v19
    }
    if v2 {
        if v2 {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 int8 = 1
                v22 := GenS2{F0: (-(v21)), F1: 104}
                switch (fi20 + (fi20 | (^(fi20)))) {
                case 127:
                    v23 := GenS3{F0: 127, F1: 165, F2: 840}
                    _ = v23
                case -407:
                    var v24 bool = false
                    _ = v24
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            var v25 uint8 = 1
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = v25
        }
        switch -188 {
        case -837:
            v26 := GenS1{F0: -364, F1: -1}
            _ = v26
        case 813:
            v27 := GenS1{F0: 366, F1: 77}
            _ = v27
        case 42:
            switch int(v3) {
            case 32767:
                if false {
                    var v28 uint16 = 42
                    if v2 {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v3 += (^(v3))
                            _ = fi29
                        }
                    }
                    v2 = (576 <= -1)
                    _ = v28
                } else {
                    v3 -= (v3 / ((v3) | 1))
                    fmt.Println(v3, v2)
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                }
            case 127:
                fmt.Println(v3, v2)
            case -1:
                switch 239 {
                case 4294967295:
                    var v30 uint8 = 1
                    _ = v30
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        default:
            switch -2147483648 {
            case 127:
                var v31 uint8 = 42
                _ = v31
            default:
                switch -1 {
                case 1:
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    }
                case 4294967295:
                    fmt.Println(v2)
                case -9223372036854775808:
                    var v32 uint8 = 127
                    _ = v32
                default:
                    switch -9223372036854775808 {
                    case 2147483647:
                        var v33 int32 = 612
                        _ = v33
                    case 474:
                        v3 = -1
                    }
                }
            }
        }
    }
    v34 := GenS4{F0: false, F1: 42}
    switch 0 {
    case 32768:
        fmt.Println(v2, v2)
    case 1:
        switch 724 {
        case 953:
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 bool = false
                for fi37 := 0; fi37 < 16; fi37++ {
                    _ = len(v1)
                    if (((((1.5 - 1.5) * 1e10) - (((float32((3.14159 * 3.14159)) * 1.5) - 3.14159) - (1e10 + (float32((((3.14159 - 1.5) / float64(0.1)) * ((((-2.25 + 1.5) + ((1e10 * (1e10 * (3.14159 - (((float64(3.14159) / 1.0) / float64(1e10)) - 0.1)))) - (-2.25 / (0.1 * (3.14159 + 1e10))))) * float64(float32((float64((0.1 + 0.1)) + 3.14159)))) * 0.0))) / 1.5)))) - ((((0.1 / 1.0) + 3.14159) / 1.0) * 0.1)) <= float32(-2.25)) {
                        switch "hello" {
                        case "":
                            v38 := GenS2{F0: 42, F1: 900}
                            _ = v38
                        default:
                            fmt.Println(v36, fi35)
                        }
                        switch (fi35 & (fi37 - (fi37 + fi35))) {
                        case -2147483648:
                            switch (("foo bar" + (("a" + "") + ("" + ""))) + "\n\t") {
                            case "foo bar":
                                v39 := GenS5{F0: 4294967295}
                                _ = v39
                            case "a":
                                fmt.Println(fi35, fi37)
                            }
                        case 2147483647:
                            v2 = (0.0 >= (3.14159 / 0.1))
                        case 42:
                            v3 = v3
                        default:
                            v40 := GenS1{F0: -818, F1: 23}
                            _ = v40
                        }
                    }
                    var v41 string = ("foo bar" + "a")
                    var v42 uint8 = 2
                    _ = fi37
                    _ = v41
                    _ = v42
                }
                _ = fi35
                _ = v36
            }
        case 42:
            if false {
                for fi43 := 0; fi43 < 14; fi43++ {
                    fmt.Println(fi43, v3)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi43
                }
            } else {
                switch 2147483647 {
                case 476:
                    v3 = (v3 << ((42) & 63))
                }
                switch "foo bar" {
                case "hello":
                    if v2 {
                        var v44 int64 = -1
                        switch "foo bar" {
                        case "":
                            var v45 int = 42
                            _ = v45
                        case "foo bar":
                            v3 = 0
                        default:
                            switch ("hello" + "") {
                            case "\n\t":
                                if (v3 != v3) {
                                    v44 = (^(v44))
                                }
                            case "a":
                                if (65535 > 382) {
                                    var v46 bool = (!v2)
                                    _ = v46
                                } else {
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        v48 := GenS5{F0: 891}
                                        if (543 < 90) {
                                            for fi49 := 0; fi49 < 10; fi49++ {
                                                fmt.Println(fi49)
                                                _ = fi49
                                            }
                                            for fi50 := 0; fi50 < 12; fi50++ {
                                                fmt.Println(fi50)
                                                fmt.Println(v3)
                                                _ = fi50
                                            }
                                            v3 = 2147483647
                                        } else {
                                            v44 += -1
                                        }
                                        _ = fi47
                                        _ = v48
                                    }
                                }
                            case "hello":
                                v2 = true
                            }
                        }
                        _ = v44
                    } else {
                        v51 := GenS1{F0: 0, F1: 1}
                        _ = v51
                    }
                default:
                    if true {
                        v2 = v2
                    }
                }
                for fi52 := 0; fi52 < 17; fi52++ {
                    v3 = 42
                    fmt.Println(fi52, fi52, fi52)
                    fmt.Println(v2, v2, fi52)
                    _ = fi52
                }
            }
        }
    case 2147483647:
        v3 = v3
    }
    var v53 int8 = 100
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi54 := 0; fi54 < 10; fi54++ {
        v55 := GenS4{F0: (v2 && v2), F1: 2}
        v3 = -898
        _ = fi54
        _ = v55
    }
    v56 := GenS1{F0: 0, F1: (v53 % (((v53 & v53)) | 1))}
    var v57 int64 = 403
    v57 += v57
    for fi58 := 0; fi58 < 6; fi58++ {
        v59 := GenS1{F0: -983, F1: v53}
        _ = fi58
        _ = v59
    }
    v57 = (v57 - v57)
    var v60 uint32 = 42
    if v2 {
        v53 += (v53 & 0)
    }
    for fi61 := 0; fi61 < 11; fi61++ {
        switch 128 {
        case 657:
            v62 := GenS4{F0: v2, F1: 957}
            _ = v62
        case -360:
            v1 = append(v1, 611)
        case 904:
            v1 = append(v1, v60)
        }
        _ = fi61
    }
    var v63 int64 = v57
    var v64 float32 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v53
    _ = v56
    _ = v57
    _ = v60
    _ = v63
    _ = v64
}

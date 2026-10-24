package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint8) uint16 {
    p0 = p0
    p0 -= p0
    if (((float64(float32((float64(-2.25) / 1.0))) / float64((3.14159 / 1.0))) != ((float64(0.0) / -2.25) / 1.0)) && (-1 < 9223372036854775807)) {
        switch -673 {
        case 127:
            fmt.Println(p0)
        case 9223372036854775807:
            for fi6 := 0; fi6 < 14; fi6++ {
                p0 = p0
                _ = fi6
            }
        }
        if false {
            fmt.Println(p0)
        } else {
            for fi7 := 0; fi7 < 1; fi7++ {
                p0 -= (p0 / ((p0) | 1))
                var v8 int64 = -9223372036854775808
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 bool = false
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: -2147483648, F1: 42, F2: p0}
            _ = v11
        }
        if false {
            p0 += 0
            if (!true) {
                v12 := GenS2{F0: -2147483648, F1: 1, F2: p0}
                p0 -= (p0 << ((2) & 63))
                fmt.Println(p0)
                _ = v12
            } else {
                fmt.Println(p0)
                var v13 int16 = 1000
                var v14 uint8 = p0
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: 32768, F1: -1, F2: 127}
            _ = v15
        } else {
            fmt.Println(p0)
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(fi16, p0)
                _ = fi16
            }
            v17 := GenS3{F0: ("hello" + "a"), F1: (p0 / (((p0 + 128)) | 1))}
            _ = v17
        }
    }
    fmt.Println(p0)
    for fi18 := 0; fi18 < 17; fi18++ {
        switch 42 {
        case -660:
            v19 := []int{}
            _ = v19
        case -260:
            v20 := GenS4{F0: (float32(0.0) == 0.1), F1: 1}
            _ = v20
        }
        switch (fi18 & 32768) {
        case 42:
            v21 := GenS5{F0: 42, F1: 4294967295, F2: 514}
            _ = v21
        }
        _ = fi18
    }
    _ = p0
    return 65535
}
func genFn1(p0 int, p1 int64, p2 int64) uint16 {
    v22 := GenS6{F0: 0}
    var v23 string = (("" + "hello") + "")
    switch p0 {
    case 32768:
        for fi24 := 0; fi24 < 7; fi24++ {
            switch (p0 | (p0 >> ((1) & 63))) {
            case 553:
                for fi25 := 0; fi25 < 17; fi25++ {
                    if false {
                        v26 := []int{4294967295, p0, 9223372036854775807, (fi24 + (fi24 ^ p0))}
                        v27 := GenS6{F0: 2147483648}
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
            }
            v28 := []int{fi24, -487, (p0 + p0), -2147483648}
            p2 = (^(p1))
            _ = fi24
            _ = v28
        }
    case 1:
        if true {
            var v29 int32 = -224
            if ((!(!(42 <= 18446744073709551615))) && (32768 == 32768)) {
                v30 := []int{p0, int(v29), p0}
                _ = v30
            } else {
                v31 := []int{int(p1), (p0 >> ((766) & 63)), p0, (p0 << ((4294967295) & 63))}
                var v32 float64 = 1e10
                if len(v31) > 0 {
                    v31[0] = (p0 % (((p0 & p0)) | 1))
                }
                _ = v31
                _ = v32
            }
            _ = v29
        } else {
            var v33 uint = 0
            for fi34 := 0; fi34 < 4; fi34++ {
                v35 := []int{fi34}
                if (!false) {
                    if false {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            switch ("a" + "hello") {
                            case "hello":
                                if len(v35) > 0 {
                                    v35[0] = fi34
                                }
                            case "foo bar":
                                var v37 int = fi36
                                _ = v37
                            case "":
                                p1 = (-(p1))
                            default:
                                if len(v35) >= 2 {
                                    v35 = v35[1:len(v35)]
                                }
                            }
                            if true {
                                _ = len(v35)
                            }
                            v38 := GenS7{F0: 42, F1: 339, F2: 4294967295}
                            _ = fi36
                            _ = v38
                        }
                    } else {
                        p1 += (p2 >> ((18446744073709551615) & 63))
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    switch ("" + "\n\t") {
                    case "":
                        fmt.Println(p2, p0)
                    case "a":
                        fmt.Println(v23)
                    default:
                        for fi39 := 0; fi39 < 12; fi39++ {
                            if len(v35) > 0 {
                                v35[0] = (fi39 ^ 500)
                            }
                            p2 += 42
                            p0 = fi39
                            _ = fi39
                        }
                    }
                }
                _ = fi34
                _ = v35
            }
            v40 := GenS3{F0: v23, F1: 1}
            _ = v33
            _ = v40
        }
    default:
        if true {
            var v41 int32 = 0
            fmt.Println(p2, p2)
            _ = v41
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 65535
}

func main() {
    var v1 int64 = -620
    v2 := make(map[string]string)
    v3 := []string{"foo bar", "\n\t"}
    v4 := make(map[string]string)
    var v5 float32 = 1.5
    if false {
        v42 := GenS8{F0: -664}
        v5 = ((v5 * (float32(((float64(1e10) / -2.25) + 1.5)) - float32(0.1))) * (v5 / v5))
        var v43 bool = (4294967295 == 960)
        _ = v42
        _ = v43
    }
    if false {
        for fi44 := 0; fi44 < 15; fi44++ {
            v45 := v2[("" + "foo bar")]
            _ = fi44
            _ = v45
        }
        var v46 uint64 = 203
        v46 = (^(v46))
        _ = v46
    } else {
        var v47 float64 = 1.5
        _ = v47
    }
    fmt.Println(v5, v5)
    if len(v3) > 0 {
        v3[0] = ("a" + ("" + "foo bar"))
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        v49 := v4["foo bar"]
        switch v49 {
        case "":
            for fi50 := 0; fi50 < 10; fi50++ {
                var v51 string = v49
                switch (("a" + "\n\t") + ((v51 + ((v51 + (v51 + "a")) + "foo bar")) + v49)) {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                }
                _ = fi50
                _ = v51
            }
        default:
            var v52 float32 = v5
            _ = v52
        }
        v5 = v5
        _ = fi48
        _ = v49
    }
    var v53 int64 = -859
    var v54 int16 = -452
    v1 -= (^(v1))
    if true {
        v1 = v53
        fmt.Println(v53)
        if ((false || false) || (2 <= 32767)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v55 uint64 = 0
            v56 := v4["\n\t"]
            _ = v55
            _ = v56
        } else {
            fmt.Println(v53, v53)
        }
    }
    v57 := v2["hello"]
    v53 = v1
    for fi58 := 0; fi58 < 18; fi58++ {
        var v59 uint32 = 575
        switch v57 {
        case "":
            v2[""] = v57
        case "a":
            v60 := GenS8{F0: 42}
            _ = v60
        }
        _ = fi58
        _ = v59
    }
    fmt.Println(v57, v5)
    if ((("foo bar" + "\n\t") + (("a" + "a") + "\n\t")) > "\n\t") {
        _ = len(v3)
        var v61 int64 = v1
        fmt.Println(v53, v5, v5)
        _ = v61
    }
    for fi62 := 0; fi62 < 4; fi62++ {
        fmt.Println(v57)
        fmt.Println(v54, v5)
        _ = len(v3)
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v57
}

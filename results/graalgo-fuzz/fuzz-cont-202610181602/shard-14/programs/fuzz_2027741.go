package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 int64; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{317, -2147483648, -2147483648, 905}
    var v3 int64 = 813
    var v4 int32 = -1
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 bool = false
        var v7 float64 = 0.0
        v6 = v6
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "a":
        var v8 int8 = 1
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: ("\n\t" + "")}
    var v10 int8 = 100
    v3 = 42
    if (true || true) {
        if false {
            v4 = int32(v10)
            if false {
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v4 -= (v4 - -877)
                        _ = fi12
                    }
                    fmt.Println(fi11, v10, v3)
                    var v13 int64 = (v3 - (v3 | v3))
                    _ = fi11
                    _ = v13
                }
            }
            if (v4 < (v4 ^ 2147483647)) {
                v10 = (v10 | -128)
            }
        } else {
            v14 := GenS2{F0: 2}
            _ = v14
        }
    } else {
        v3 = -674
    }
    var v15 bool = (v3 <= 42)
    v15 = false
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (fi16 % ((int(v3)) | 1)) {
        case 2147483647:
            v17 := GenS2{F0: 4294967295}
            _ = v17
        case 127:
            fmt.Println(fi16)
        }
        _ = fi16
    }
    var v18 uint8 = 255
    if v15 {
        switch "a" {
        case "a":
            v19 := GenS1{F0: 42, F1: ("\n\t" + "\n\t")}
            _ = v19
        case "":
            switch 1 {
            case 808:
                for fi20 := 0; fi20 < 19; fi20++ {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 int32 = int32(fi20)
                        var v23 int16 = -32768
                        var v24 float32 = 1.5
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    switch -393 {
                    case 0:
                        v18 += v18
                    case -152:
                        var v25 int64 = v3
                        _ = v25
                    case 9223372036854775807:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v2 = append(v2, fi20)
                    }
                    fmt.Println(fi20)
                    _ = fi20
                }
            case -9223372036854775808:
                v1[543] = 4294967295
            case -2147483648:
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 uint8 = 9
                    switch "hello" {
                    case "a":
                        switch "\n\t" {
                        case "a":
                            v10 = 100
                        case "\n\t":
                            fmt.Println(fi26)
                        }
                    case "\n\t":
                        v28 := GenS1{F0: 4294967295, F1: ""}
                        _ = v28
                    default:
                        v29 := GenS2{F0: 1}
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
            default:
                v30 := GenS2{F0: 0}
                _ = v30
            }
        case "hello":
            var v31 float32 = 0.1
            _ = v31
        default:
            if (!v15) {
                for fi32 := 0; fi32 < 4; fi32++ {
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v4 = v4
                        v10 = (^(v10))
                        _ = fi33
                    }
                    _ = fi32
                }
                var v34 int16 = 1000
                _ = v34
            }
        }
    }
    fmt.Println(v10)
    var v35 int16 = 0
    if v15 {
        switch -125 {
        case 32767:
            fmt.Println(v35)
        case 1:
            if ((684 >= 1) || (1.5 > (-2.25 + 1e10))) {
                var v36 float32 = ((((-2.25 - (((((-2.25 * 0.0) + (float32((0.1 + float64(((3.14159 / float32((1e10 * 3.14159))) / float32(-2.25))))) / 3.14159)) - 0.1) + float32((((float64(float32(1.5)) * 3.14159) + float64((((3.14159 + 0.1) / float32(((3.14159 / -2.25) * (1e10 - -2.25)))) - 1.5))) + 0.0))) - (0.1 / -2.25))) + (1.5 - float32(1e10))) / -2.25) + 1.5)
                _ = v36
            } else {
                var v37 uint16 = 635
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            for fi39 := 0; fi39 < 11; fi39++ {
                switch (fi39 ^ 32768) {
                case -866:
                    v35 -= (v35 & 318)
                case 32768:
                    switch fi38 {
                    case 4294967295:
                        fmt.Println(v18)
                    case 1:
                        var v40 uint = 4294967295
                        _ = v40
                    default:
                        if false {
                            fmt.Println(v10, v4, v4)
                            v2 = append(v2, 128)
                            var v41 int32 = v4
                            _ = v41
                        }
                    }
                case 42:
                    fmt.Println(fi39, v4)
                default:
                    switch "foo bar" {
                    case "hello":
                        v42 := GenS1{F0: uint(v18), F1: ("foo bar" + ("hello" + "hello"))}
                        _ = v42
                    case "a":
                        v43 := GenS1{F0: 145, F1: ""}
                        _ = v43
                    case "\n\t":
                        fmt.Println(v35, v4)
                    default:
                        var v44 float32 = 1e10
                        _ = v44
                    }
                }
                var v45 float64 = (0.1 / 1.5)
                v46 := v1[2147483647]
                _ = fi39
                _ = v45
                _ = v46
            }
            switch (fi38 | (fi38 >> ((827) & 63))) {
            case 42:
                v47 := GenS3{F0: 18446744073709551615}
                _ = v47
            }
            if v15 {
                if len(v2) > 0 {
                    v2[0] = -9223372036854775808
                }
            } else {
                var v48 float32 = (float32(3.14159) - ((1e10 + -2.25) / float32((1e10 / ((3.14159 + (0.0 - -2.25)) * 1.5)))))
                _ = v48
            }
            _ = fi38
        }
    }
    var v49 int64 = int64(v35)
    switch ("" + "hello") {
    case "\n\t":
        if ((!(!true)) && (false || (839 < 254))) {
            for fi50 := 0; fi50 < 19; fi50++ {
                var v51 float32 = 1.5
                _ = fi50
                _ = v51
            }
        }
    }
    v18 += (v18 >> ((2) & 63))
    for fi52 := 0; fi52 < 13; fi52++ {
        v4 = (-(v4))
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint = 162
            _ = len(v2)
            if v15 {
                switch (fi53 - 0) {
                case 9223372036854775807:
                    switch fi53 {
                    case 128:
                        v55 := GenS4{F0: (42 > 4294967295), F1: (v10 << ((v54) & 63))}
                        _ = v55
                    case -700:
                        var v56 uint32 = 502
                        _ = v56
                    }
                case 991:
                    var v57 int8 = v10
                    _ = v57
                case 127:
                    if v15 {
                        v3 = -9223372036854775808
                        switch fi52 {
                        case 127:
                            v35 = (^(v35))
                        }
                        fmt.Println(v35)
                    } else {
                        var v58 float32 = 3.14159
                        var v59 int64 = v49
                        if v15 {
                            if len(v2) > 0 {
                                v2[0] = fi53
                            }
                            v60 := GenS5{F0: 2147483648, F1: (v59 % ((int64(fi52)) | 1)), F2: v15}
                            v35 = v35
                            var v61 uint8 = 2
                            v62 := GenS6{F0: 1}
                            var v63 uint32 = 404
                            if false {
                                v64 := GenS4{F0: v15, F1: (v10 % (((v10 | 42)) | 1))}
                                _ = v64
                            }
                            _ = v60
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        _ = v58
                        _ = v59
                    }
                default:
                    fmt.Println(v3, v3, v15)
                }
                switch "a" {
                case "":
                    v65 := GenS7{F0: ((("" + "hello") + "a") + "foo bar")}
                    if ((v49 % ((9223372036854775807) | 1)) > (v3 * (-(v49)))) {
                        var v66 bool = v15
                        _ = v66
                    }
                    _ = v65
                default:
                    fmt.Println(fi53, v49, v4)
                }
            } else {
                _ = len(v2)
                v67 := GenS8{F0: v4}
                _ = v67
            }
            _ = fi53
            _ = v54
        }
        if false {
            switch fi52 {
            case 32768:
                v68 := GenS4{F0: v15, F1: (v10 >> ((65535) & 63))}
                _ = v68
            case 2147483647:
                var v69 string = "a"
                _ = v69
            case 42:
                var v70 int = (fi52 - int(v35))
                _ = v70
            default:
                v1[1] = (-(fi52))
            }
            for fi71 := 0; fi71 < 7; fi71++ {
                if len(v2) > 0 {
                    v2[0] = 32767
                }
                switch ((("" + "foo bar") + "") + "hello") {
                case "foo bar":
                    v35 = v35
                case "":
                    if false {
                        switch fi52 {
                        case 32767:
                            var v72 uint32 = 2147483648
                            _ = v72
                        case 4294967295:
                            fmt.Println(v18)
                        case 0:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    } else {
                        v4 += (v4 + (v4 * v4))
                    }
                case "a":
                    v73 := GenS7{F0: ("a" + "hello")}
                    _ = v73
                default:
                    var v74 uint32 = 0
                    _ = v74
                }
                _ = fi71
            }
            var v75 int8 = (v10 - (v10 + v10))
            _ = v75
        }
        _ = fi52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v49
    _ = v9
}

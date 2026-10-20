package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0(p0 int8, p1 string) uint64 {
    v5 := GenS1{F0: "\n\t"}
    if (340 > 1) {
        v6 := []int{32768, 128, -308}
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: ("a" + ((p1 + "a") + p1))}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return 55
}
func genFn1(p0 uint) string {
    var v9 int32 = 2147483647
    v10 := []int{32767, -103, -963}
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(p0, p0, fi11)
        _ = fi11
    }
    v10 = append(v10, -2147483648)
    _ = p0
    _ = v10
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 914}
    var v2 bool = (!false)
    var v3 uint64 = 18446744073709551615
    var v4 int32 = 444
    var v12 float32 = float32((-2.25 / 1e10))
    if v2 {
        v13 := GenS1{F0: "\n\t"}
        if v2 {
            if ((("" + "a") + "\n\t") == ("foo bar" + "")) {
                v4 += v4
                for fi14 := 0; fi14 < 17; fi14++ {
                    v4 = (-(v4))
                    _ = fi14
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    switch fi15 {
                    case 128:
                        v1 = append(v1, 2147483648)
                    case -932:
                        v2 = ((!false) && (-968 != 32767))
                    case -9223372036854775808:
                        var v16 float32 = (0.1 / (v12 / (v12 - 1.5)))
                        _ = v16
                    }
                    v17 := GenS2{F0: 1, F1: "a", F2: 0}
                    _ = fi15
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                v1 = append(v1, 1)
                _ = fi18
            }
        }
        var v19 uint64 = 0
        _ = v13
        _ = v19
    } else {
        v20 := GenS1{F0: "a"}
        switch "a" {
        case "\n\t":
            v1 = append(v1, 2)
        case "hello":
            for fi21 := 0; fi21 < 5; fi21++ {
                if v2 {
                    v12 = -2.25
                    var v22 int32 = v4
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v12 = (v12 / float32(0.1))
                        if ((v3 << ((65535) & 63)) <= (v3 << ((4294967295) & 63))) {
                            if v2 {
                                v24 := GenS2{F0: 0, F1: "foo bar", F2: 0}
                                if (v2 || v2) {
                                    if (42 != 42) {
                                        if v2 {
                                            var v25 float64 = -2.25
                                            v4 -= (v4 << ((778) & 63))
                                            _ = v25
                                        }
                                        var v26 int64 = 42
                                        _ = v26
                                    }
                                    v1 = append(v1, 2)
                                } else {
                                    switch -1 {
                                    case -748:
                                        if (((-1 != -522) || (true || false)) || true) {
                                            var v27 int16 = 1
                                            _ = v27
                                        } else {
                                            v2 = v2
                                            fmt.Println(v3, fi21)
                                            var v28 int8 = -1
                                            _ = v28
                                        }
                                    case -173:
                                        if len(v1) > 0 {
                                            v1[0] = 213
                                        }
                                    case -1:
                                        v29 := GenS3{F0: 0}
                                        _ = v29
                                    }
                                    v1 = append(v1, 42)
                                }
                                v12 -= v12
                                _ = v24
                            } else {
                                fmt.Println(v2, fi23, v3)
                            }
                            if (((-2.25 + (((1.5 * 1e10) * float64(v12)) + (0.1 - (float64((0.1 - v12)) - (0.1 - (1.5 - -2.25)))))) / ((1e10 - float64(float32(3.14159))) + 3.14159)) != 3.14159) {
                                v1 = append(v1, 225)
                                var v30 int32 = v4
                                v1 = append(v1, 42)
                                v2 = v2
                                _ = v30
                            }
                            switch fi21 {
                            case -1:
                                if false {
                                    v3 = (v3 * 228)
                                    v12 += (v12 + v12)
                                    switch ("a" + ("a" + "hello")) {
                                    case "a":
                                        var v31 int8 = 100
                                        _ = v31
                                    case "foo bar":
                                        fmt.Println(v4, v2)
                                    case "\n\t":
                                        var v32 uint16 = 755
                                        _ = v32
                                    }
                                } else {
                                    var v33 bool = v2
                                    _ = v33
                                }
                            default:
                                v34 := GenS1{F0: ("a" + "\n\t")}
                                _ = v34
                            }
                        } else {
                            v35 := GenS2{F0: 127, F1: "foo bar", F2: 1}
                            v36 := GenS1{F0: ""}
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS1{F0: ("" + "hello")}
                        _ = fi23
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 15; fi38++ {
                    fmt.Println(v12, v2, v4)
                    var v39 string = "foo bar"
                    fmt.Println(v39, v12)
                    _ = fi38
                    _ = v39
                }
                if (479 >= 849) {
                    fmt.Println(fi21, v3)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v4 += v4
                } else {
                    var v40 int = -893
                    _ = v40
                }
                _ = fi21
            }
        default:
            switch (("hello" + "") + "\n\t") {
            case "hello":
                v4 -= v4
            }
        }
        if (735 == 377) {
            var v41 uint16 = 19
            v1 = append(v1, 761)
            fmt.Println(v3)
            _ = v41
        } else {
            switch "a" {
            case "foo bar":
                var v42 uint16 = 255
                _ = v42
            case "":
                v2 = v2
            case "a":
                for fi43 := 0; fi43 < 8; fi43++ {
                    v2 = v2
                    var v44 string = "a"
                    v45 := GenS4{F0: 856}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            default:
                if (42 > 128) {
                    switch -692 {
                    case -959:
                        var v46 int = -1
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 19; fi47++ {
                        _ = len(v1)
                        _ = fi47
                    }
                } else {
                    fmt.Println(v2)
                    fmt.Println(v3, v12, v12)
                    if v2 {
                        v2 = v2
                    }
                }
            }
            if (v2 || v2) {
                for fi48 := 0; fi48 < 4; fi48++ {
                    var v49 int64 = 9223372036854775807
                    _ = fi48
                    _ = v49
                }
                fmt.Println(v12, v12, v12)
                switch "foo bar" {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    v3 = (^(v3))
                case "":
                    v2 = true
                default:
                    v50 := GenS4{F0: 1}
                    _ = v50
                }
            } else {
                for fi51 := 0; fi51 < 6; fi51++ {
                    v12 = v12
                    v12 = float32(3.14159)
                    if (((v4 / ((-2147483648) | 1)) != (-(v4))) || (!true)) {
                        v2 = true
                    }
                    _ = fi51
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
            }
            for fi52 := 0; fi52 < 1; fi52++ {
                fmt.Println(v12, v12, v4)
                var v53 uint16 = 522
                _ = fi52
                _ = v53
            }
        }
        _ = v20
    }
    if v2 {
        for fi54 := 0; fi54 < 19; fi54++ {
            v3 -= (^(v3))
            switch (fi54 << ((1) & 63)) {
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 765:
                fmt.Println(v3, fi54, v2)
            case 2147483647:
                v4 -= 1
            default:
                switch "hello" {
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    v1 = append(v1, 2147483647)
                default:
                    if (!v2) {
                        if v2 {
                            var v55 int64 = -9223372036854775808
                            _ = v55
                        } else {
                            var v56 uint32 = 524
                            v57 := GenS5{F0: uint16(v3), F1: 0}
                            _ = v56
                            _ = v57
                        }
                        v2 = true
                        if true {
                            v3 -= v3
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 399
                        }
                    }
                }
            }
            _ = fi54
        }
        fmt.Println(v12)
        switch 32767 {
        case 42:
            v3 += 0
        }
    } else {
        _ = len(v1)
        switch -424 {
        case -1:
            if len(v1) > 0 {
                v1[0] = 933
            }
        case 411:
            for fi58 := 0; fi58 < 10; fi58++ {
                v4 -= 42
                v12 += float32((float64((float32((float64(1e10) / ((1.5 * -2.25) * 3.14159))) + float32((3.14159 / ((1e10 - 0.1) + 1.5))))) * -2.25))
                _ = fi58
            }
        case 662:
            for fi59 := 0; fi59 < 14; fi59++ {
                if ("hello" != "\n\t") {
                    v60 := GenS2{F0: 127, F1: "\n\t", F2: 0}
                    var v61 string = "hello"
                    _ = v60
                    _ = v61
                }
                _ = fi59
            }
        default:
            var v62 uint = 18446744073709551615
            _ = v62
        }
        var v63 int16 = 829
        _ = v63
    }
    for fi64 := 0; fi64 < 10; fi64++ {
        if len(v1) > 0 {
            v1[0] = 888
        }
        switch fi64 {
        case 4294967295:
            v65 := GenS1{F0: "\n\t"}
            _ = v65
        }
        if len(v1) > 0 {
            v1[0] = 269
        }
        _ = fi64
    }
    v3 = v3
    v12 += 1.5
    switch "" {
    case "a":
        v4 = (v4 / (((v4 << ((239) & 63))) | 1))
    case "\n\t":
        v4 = v4
    case "foo bar":
        if (32767 <= 964) {
            var v66 uint32 = 42
            var v67 float64 = (1.5 - (0.0 - (0.1 - float64(float32((0.1 * (3.14159 / 1.5)))))))
            var v68 uint32 = v66
            _ = v66
            _ = v67
            _ = v68
        }
    }
    switch -2147483648 {
    case -1:
        if (1e10 == 1.5) {
            v4 += v4
        } else {
            v12 -= v12
        }
    case 2147483647:
        var v69 string = "foo bar"
        _ = v69
    }
    if ("hello" > "a") {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v70 := GenS4{F0: 1000}
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = v70
    }
    fmt.Println(v3, v12)
    var v71 string = ("a" + "hello")
    if v2 {
        v72 := GenS5{F0: 516, F1: 65535}
        for fi73 := 0; fi73 < 2; fi73++ {
            fmt.Println(v3, fi73)
            if ((!v2) || true) {
                if (v12 >= v12) {
                    var v74 uint64 = v3
                    _ = v74
                }
                fmt.Println(v4, v3)
                if (763 <= -1) {
                    var v75 string = v71
                    var v76 string = v71
                    var v77 uint64 = v3
                    _ = v75
                    _ = v76
                    _ = v77
                } else {
                    v78 := GenS3{F0: 2}
                    _ = v78
                }
            } else {
                var v79 uint = 65535
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v80 string = "\n\t"
                _ = v79
                _ = v80
            }
            _ = fi73
        }
        _ = v72
    }
    for fi81 := 0; fi81 < 5; fi81++ {
        for fi82 := 0; fi82 < 5; fi82++ {
            var v83 string = ""
            var v84 float64 = 0.1
            _ = fi82
            _ = v83
            _ = v84
        }
        v85 := GenS6{F0: (!v2), F1: (v4 | v4), F2: 42}
        for fi86 := 0; fi86 < 16; fi86++ {
            var v87 uint8 = 128
            switch ((v71 + v71) + "a") {
            case "\n\t":
                if v2 {
                    switch (fi86 & fi81) {
                    case -2147483648:
                        v71 = (("foo bar" + ("a" + v71)) + v71)
                    }
                    v87 = v87
                    var v88 uint64 = (v3 - 18446744073709551615)
                    _ = v88
                } else {
                    v89 := GenS6{F0: (2147483647 > 0), F1: v4, F2: 1}
                    _ = v89
                }
            case "hello":
                v4 = int32(v3)
            case "a":
                var v90 float64 = 3.14159
                _ = v90
            }
            _ = fi86
            _ = v87
        }
        _ = fi81
        _ = v85
    }
    var v91 uint64 = genFn0(-128, v71)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v91
}

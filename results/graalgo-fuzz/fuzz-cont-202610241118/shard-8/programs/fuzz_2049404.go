package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int16 = 658
    var v2 uint64 = 0
    v3 := []string{"a", "\n\t", "a"}
    v4 := make(map[uint32]uint32)
    v5 := []string{(("a" + "\n\t") + ("foo bar" + "foo bar"))}
    switch 9223372036854775807 {
    case 42:
        if (255 >= 519) {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 bool = ((1.5 - ((0.0 * 3.14159) * 3.14159)) >= 0.1)
                if true {
                    v2 -= (v2 * v2)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (-740 > -1) {
                        _ = len(v5)
                        v8 := GenS1{F0: 422}
                        _ = v8
                    } else {
                        v1 += v1
                        switch -702 {
                        case -485:
                            if ((1.5 / 1.0) > -2.25) {
                                v7 = (783 != 2)
                            }
                        case -712:
                            for fi9 := 0; fi9 < 16; fi9++ {
                                if (v7 || true) {
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v7)
                                    default:
                                        if (v2 != v2) {
                                            var v10 int32 = -2147483648
                                            var v11 float32 = 3.14159
                                            v11 = 1.5
                                            v12 := GenS1{F0: 2}
                                            _ = v10
                                            _ = v11
                                            _ = v12
                                        }
                                    }
                                }
                                v4[620] = 290
                                v13 := GenS2{F0: (v2 ^ 1)}
                                _ = fi9
                                _ = v13
                            }
                        case -9223372036854775808:
                            fmt.Println(v7, v7)
                        default:
                            if false {
                                var v14 float32 = 1e10
                                v7 = (1 == 2)
                                for fi15 := 0; fi15 < 15; fi15++ {
                                    if (-2 <= -927) {
                                        v7 = ((!(!(v1 < (v1 / ((v1) | 1))))) && (-629 < 42))
                                        v2 += (v2 / ((374) | 1))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        fmt.Println(v14)
                                        var v16 int32 = 2147483647
                                        _ = v16
                                    }
                                    _ = fi15
                                }
                                _ = v14
                            }
                        }
                        for fi17 := 0; fi17 < 5; fi17++ {
                            fmt.Println(v7, v2)
                            var v18 float32 = (-2.25 - (0.0 + float32(1.5)))
                            v19 := GenS1{F0: 885}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                } else {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v2 -= 0
                        var v21 float32 = -2.25
                        v1 = (v1 & v1)
                        v2 -= v2
                        v21 -= v21
                        var v22 float32 = (float32((3.14159 + (1.5 + 0.1))) * (v21 - (((-2.25 * v21) + v21) * 0.1)))
                        switch (-(fi6)) {
                        case 127:
                            fmt.Println(fi20, fi20, fi20)
                        case 4294967295:
                            if (!(3.14159 < (v22 - float32(-2.25)))) {
                                var v23 float64 = 0.0
                                fmt.Println(v21, v23)
                                switch fi6 {
                                case 32767:
                                    switch 9223372036854775807 {
                                    case -890:
                                        var v24 uint = 42
                                        _ = v24
                                    case 32768:
                                        switch "hello" {
                                        case "foo bar":
                                            var v25 int64 = 1
                                            _ = v25
                                        case "a":
                                            var v26 float32 = (((v21 * v21) * (v21 - 1.5)) + ((v21 + v21) + (v21 - v22)))
                                            var v27 uint16 = 1
                                            _ = v26
                                            _ = v27
                                        default:
                                            if (!v7) {
                                                if len(v3) > 0 {
                                                    v3[0] = ((("" + "a") + ("" + (((("" + "foo bar") + "foo bar") + "hello") + "hello"))) + "foo bar")
                                                }
                                                fmt.Println(v21)
                                                var v28 uint = 18446744073709551615
                                                _ = v28
                                            } else {
                                                var v29 uint32 = 575
                                                _ = v29
                                            }
                                        }
                                    case -1:
                                        v2 = v2
                                    default:
                                        var v30 uint8 = 55
                                        _ = v30
                                    }
                                default:
                                    fmt.Println(v7, v2, fi6)
                                }
                                _ = v23
                            }
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    var v31 float64 = (0.1 * 0.1)
                    var v32 string = "a"
                    _ = v31
                    _ = v32
                }
                v1 = (-(v1))
                _ = fi6
                _ = v7
            }
            var v33 int8 = 123
            _ = v33
        } else {
            var v34 bool = false
            var v35 uint64 = v2
            var v36 float64 = 3.14159
            _ = v34
            _ = v35
            _ = v36
        }
    case 2147483647:
        v1 -= v1
    }
    var v37 uint = 0
    v38 := GenS3{F0: (-(v1)), F1: false, F2: -46}
    _ = len(v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v37, v1, v2)
    v2 = v2
    fmt.Println(v37, v2)
    v2 -= (v2 - v2)
    switch 32767 {
    case 127:
        for fi39 := 0; fi39 < 4; fi39++ {
            v37 += 2
            v40 := v4[2]
            var v41 int = 856
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    switch ("a" + "a") {
    case "hello":
        var v42 int64 = 0
        _ = v42
    case "a":
        var v43 string = (("" + "a") + (("" + "\n\t") + "foo bar"))
        _ = v43
    }
    switch 9223372036854775807 {
    case 348:
        var v44 uint32 = 4294967295
        _ = v44
    case -2147483648:
        for fi45 := 0; fi45 < 4; fi45++ {
            switch (fi45 << ((uint(v2)) & 63)) {
            case -1:
                var v46 float64 = -2.25
                _ = v46
            case 42:
                switch (fi45 / ((fi45) | 1)) {
                case 0:
                    if ((uint8(v2) > 1) && true) {
                        var v47 bool = false
                        _ = v47
                    } else {
                        var v48 int = (-(fi45))
                        _ = v48
                    }
                case -2147483648:
                    v37 -= 322
                case 2147483647:
                    v49 := GenS3{F0: (v1 << ((v37) & 63)), F1: (((0.1 + (((0.1 - -2.25) / (0.0 + 0.1)) - 0.1)) / (0.1 - 0.0)) == 0.0), F2: -128}
                    _ = v49
                }
            default:
                var v50 float32 = 1.5
                _ = v50
            }
            var v51 int16 = v1
            var v52 int64 = -1
            _ = fi45
            _ = v51
            _ = v52
        }
    case 32767:
        if false {
            fmt.Println(v37)
            v2 += (^(v2))
            switch -399 {
            case 42:
                if false {
                    if ((v1 | v1) != (v1 % (((v1 - (v1 | 0))) | 1))) {
                        for fi53 := 0; fi53 < 16; fi53++ {
                            var v54 string = ""
                            _ = len(v5)
                            v1 += (v1 | v1)
                            _ = fi53
                            _ = v54
                        }
                    }
                    for fi55 := 0; fi55 < 16; fi55++ {
                        v1 = v1
                        var v56 int64 = 9223372036854775807
                        _ = fi55
                        _ = v56
                    }
                } else {
                    v57 := GenS4{F0: false}
                    switch 131 {
                    case 4294967295:
                        v58 := v4[1]
                        _ = v58
                    case -2147483648:
                        var v59 uint64 = v2
                        _ = v59
                    case -1:
                        switch -34 {
                        case 0:
                            if (!(false && (!(true || true)))) {
                                if (v2 == 42) {
                                    v37 = (v37 >> ((v37) & 63))
                                    var v60 string = "foo bar"
                                    switch int(v1) {
                                    case 112:
                                        v37 = v37
                                    case 4294967295:
                                        for fi61 := 0; fi61 < 16; fi61++ {
                                            if false {
                                                var v62 uint16 = 32767
                                                switch (v60 + "") {
                                                case "\n\t":
                                                    _ = len(v4)
                                                case "hello":
                                                    var v63 string = ("" + ((v60 + "hello") + v60))
                                                    _ = v63
                                                }
                                                for fi64 := 0; fi64 < 15; fi64++ {
                                                    var v65 float64 = (1.5 / (1.5 - -2.25))
                                                    var v66 int = (fi64 >> ((v37) & 63))
                                                    _ = fi64
                                                    _ = v65
                                                    _ = v66
                                                }
                                                _ = v62
                                            }
                                            v2 = v2
                                            _ = fi61
                                        }
                                    case 32768:
                                        if (!(!(2147483648 != 4294967295))) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                    _ = v60
                                }
                                var v67 int8 = int8(v1)
                                v67 -= v67
                                _ = v67
                            } else {
                                if ("hello" < "") {
                                    v1 += (-(v1))
                                    var v68 uint8 = 88
                                    _ = v68
                                }
                                var v69 float64 = 0.1
                                var v70 string = "\n\t"
                                _ = v69
                                _ = v70
                            }
                        case 9223372036854775807:
                            var v71 uint32 = 2
                            _ = v71
                        case -9223372036854775808:
                            var v72 int64 = -9223372036854775808
                            _ = v72
                        default:
                            v1 = 1
                        }
                    }
                    v1 += (v1 & (v1 % ((-32768) | 1)))
                    _ = v57
                }
            case -2147483648:
                var v73 int8 = 56
                _ = v73
            default:
                v4[2147483648] = 2
            }
        } else {
            v2 -= (v2 >> ((v37) & 63))
            switch "foo bar" {
            case "":
                if false {
                    v37 += v37
                }
            case "\n\t":
                var v74 int64 = 0
                _ = v74
            default:
                v37 = v37
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    default:
        v2 -= (v2 - (^(v2)))
    }
    var v75 uint64 = v2
    var v76 int16 = (-(v1))
    for fi77 := 0; fi77 < 6; fi77++ {
        v75 -= (v2 - (^(v2)))
        fmt.Println(v37, v1)
        _ = fi77
    }
    var v78 string = "\n\t"
    if (3.14159 == 0.1) {
        v79 := GenS4{F0: false}
        _ = v79
    } else {
        var v80 string = "a"
        _ = v80
    }
    var v81 int = -551
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v78
    _ = v81
}

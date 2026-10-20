package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint8; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 bool; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 string; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 int8; F2 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    v3 := make(map[string]uint32)
    v4 := make(map[string]string)
    v5 := []string{"\n\t", "", "a", "foo bar"}
    v2 = 0.1
    var v6 float32 = 3.14159
    if false {
        fmt.Println(v6, v6)
        v2 -= (((0.1 / v2) / v2) * ((3.14159 * 1.5) + 0.0))
        delete(v4, ("a" + ((("foo bar" + "\n\t") + ("" + ("a" + "foo bar"))) + "foo bar")))
    } else {
        if true {
            if (!(!(1 > 0))) {
                v2 = (1e10 / float64(1e10))
                if false {
                    fmt.Println(v2, v6)
                    fmt.Println(v2)
                    for fi7 := 0; fi7 < 8; fi7++ {
                        v2 += v2
                        var v8 int = fi7
                        v6 += 1e10
                        _ = fi7
                        _ = v8
                    }
                }
                v9 := GenS1{F0: true}
                _ = v9
            }
            var v10 int32 = -1
            _ = v10
        } else {
            _ = len(v3)
            v11 := GenS2{F0: 42, F1: 2147483647}
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int64 = 42
                if (127 == 253) {
                    var v14 bool = false
                    fmt.Println(v2)
                    var v15 uint64 = 821
                    _ = v14
                    _ = v15
                } else {
                    if (!false) {
                        var v16 string = "a"
                        var v17 int8 = -128
                        v6 = 3.14159
                        _ = v16
                        _ = v17
                    }
                    v13 = (v13 / ((v13) | 1))
                    fmt.Println(v2)
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v6 = float32(1.5)
            _ = fi18
        }
        var v19 uint8 = 1
        _ = v19
    }
    var v20 bool = (-5 <= 1)
    var v21 int32 = 42
    if v20 {
        v2 = 1e10
    }
    v6 = 3.14159
    v6 += 0.0
    var v22 float32 = (v6 / ((-2.25 * 1.5) * float32(v2)))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v23 bool = (((!(4294967295 != 1)) && true) && v20)
    v22 += v22
    if false {
        if (((v23 || (true || (42 < -9223372036854775808))) || v23) && false) {
            v24 := GenS3{F0: -1, F1: 127, F2: v21}
            if (!true) {
                var v25 uint8 = 127
                fmt.Println(v2, v25, v22)
                v26 := GenS2{F0: 42, F1: 2147483648}
                _ = v25
                _ = v26
            } else {
                var v27 float32 = (float32(v2) * 0.0)
                _ = v27
            }
            _ = v24
        } else {
            for fi28 := 0; fi28 < 11; fi28++ {
                if v20 {
                    if false {
                        _ = len(v5)
                        v29 := GenS1{F0: v23}
                        _ = v29
                    } else {
                        fmt.Println(v2, v20, v22)
                        if v20 {
                            switch "\n\t" {
                            case "foo bar":
                                v23 = true
                            case "hello":
                                v5 = append(v5, ("" + ("\n\t" + ("foo bar" + "a"))))
                            case "":
                                if v23 {
                                    var v30 int8 = 127
                                    var v31 uint = 42
                                    v32 := GenS4{F0: v20}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS5{F0: 1, F1: false, F2: -1}
                                    fmt.Println(v23)
                                    switch "hello" {
                                    case "\n\t":
                                        if (42 <= -455) {
                                            fmt.Println(v21, v20, v23)
                                            v22 = v22
                                            if (42 < 991) {
                                                if ((v22 + (v6 + (1.5 * 3.14159))) >= v22) {
                                                    for fi34 := 0; fi34 < 5; fi34++ {
                                                        v35 := GenS4{F0: (!(9223372036854775807 != -1))}
                                                        _ = fi34
                                                        _ = v35
                                                    }
                                                    var v36 uint8 = 128
                                                    _ = v36
                                                } else {
                                                    _ = len(v5)
                                                    var v37 uint = 1
                                                    _ = v37
                                                }
                                            }
                                        }
                                    }
                                    _ = v33
                                }
                            }
                        }
                        v23 = v23
                    }
                    if (42 == 42) {
                        fmt.Println(v6, v22, v21)
                        var v38 int32 = (-(v21))
                        fmt.Println(v2, v6)
                        _ = v38
                    } else {
                        if false {
                            var v39 string = ((("\n\t" + ("\n\t" + "\n\t")) + "") + "hello")
                            fmt.Println(fi28)
                            _ = v39
                        } else {
                            v40 := GenS2{F0: 65535, F1: 65535}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v40
                        }
                    }
                } else {
                    for fi41 := 0; fi41 < 5; fi41++ {
                        var v42 uint32 = 2147483647
                        var v43 int16 = int16(fi41)
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    var v44 int64 = -9223372036854775808
                    var v45 string = ""
                    _ = v44
                    _ = v45
                }
                switch "foo bar" {
                case "":
                    if (0 > fi28) {
                        fmt.Println(v2)
                        if len(v1) > 0 {
                            v1[0] = -18
                        }
                    } else {
                        v2 -= 0.0
                        v46 := GenS6{F0: fi28}
                        switch 1 {
                        case 782:
                            v47 := GenS1{F0: true}
                            _ = v47
                        case 530:
                            for fi48 := 0; fi48 < 10; fi48++ {
                                v49 := GenS7{F0: 877}
                                if (true || v20) {
                                    v22 += v6
                                    delete(v4, "a")
                                    var v50 int64 = -578
                                    _ = v50
                                }
                                _ = fi48
                                _ = v49
                            }
                        case 128:
                            switch fi28 {
                            case -9223372036854775808:
                                for fi51 := 0; fi51 < 20; fi51++ {
                                    v6 = (-2.25 * float32(v2))
                                    _ = fi51
                                }
                            case 32768:
                                for fi52 := 0; fi52 < 15; fi52++ {
                                    switch (^(fi52)) {
                                    case 9223372036854775807:
                                        if v20 {
                                            v1 = append(v1, (fi52 / (((fi28 * fi52)) | 1)))
                                            v53 := GenS8{F0: 2, F1: (("hello" + "") + "\n\t"), F2: 42}
                                            _ = v53
                                        } else {
                                            v54 := GenS9{F0: 0}
                                            switch (fi28 >> ((0) & 63)) {
                                            case 2147483647:
                                                v6 += v6
                                                fmt.Println(v20)
                                                v55 := GenS10{F0: 262, F1: -52, F2: -71}
                                                _ = v55
                                            case -2147483648:
                                                switch fi52 {
                                                case -2147483648:
                                                    if (!(v6 <= (v6 + v22))) {
                                                        var v56 uint = 384
                                                        _ = v56
                                                    } else {
                                                        var v57 uint = 18446744073709551615
                                                        v20 = true
                                                        _ = v57
                                                    }
                                                case 2147483647:
                                                    v5 = append(v5, (("\n\t" + "") + "foo bar"))
                                                case 128:
                                                    v23 = (174 <= 2)
                                                }
                                            case -940:
                                                v5 = append(v5, "foo bar")
                                            }
                                            var v58 string = "foo bar"
                                            if ((-2.25 - (v6 + (0.0 / 0.1))) >= v6) {
                                                var v59 int16 = -1
                                                v60 := v4[v58]
                                                if v23 {
                                                    switch 1 {
                                                    case 901:
                                                        v23 = v23
                                                    case 755:
                                                        if ((fi28 / (((fi28 + 42)) | 1)) == fi28) {
                                                            _ = len(v3)
                                                            v61 := GenS8{F0: 18446744073709551615, F1: v58, F2: -1}
                                                            switch v60 {
                                                            case "hello":
                                                                if true {
                                                                    var v62 int32 = v21
                                                                    _ = v62
                                                                }
                                                            }
                                                            _ = v61
                                                        }
                                                    case 676:
                                                        if len(v1) > 0 {
                                                            v1[0] = fi28
                                                        }
                                                    default:
                                                        v20 = (128 < 127)
                                                    }
                                                }
                                                _ = v59
                                                _ = v60
                                            } else {
                                                if v20 {
                                                    var v63 string = v58
                                                    var v64 bool = v23
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    var v65 uint16 = 1
                                                    _ = len(v5)
                                                    var v66 int32 = v21
                                                    _ = v65
                                                    _ = v66
                                                }
                                                delete(v4, "\n\t")
                                                var v67 uint8 = 231
                                                _ = v67
                                            }
                                            var v68 uint64 = 42
                                            _ = v54
                                            _ = v58
                                            _ = v68
                                        }
                                    case -1:
                                        v2 -= (v2 + ((v2 - (3.14159 / 1e10)) * v2))
                                    case 127:
                                        var v69 uint16 = 1
                                        _ = v69
                                    }
                                    fmt.Println(v21, v23)
                                    var v70 string = "\n\t"
                                    _ = fi52
                                    _ = v70
                                }
                            default:
                                v20 = true
                            }
                        }
                        _ = v46
                    }
                case "foo bar":
                    for fi71 := 0; fi71 < 19; fi71++ {
                        v6 += ((v6 - float32(-2.25)) / -2.25)
                        v5 = append(v5, "foo bar")
                        _ = fi71
                    }
                case "a":
                    for fi72 := 0; fi72 < 4; fi72++ {
                        v2 = ((v2 + ((v2 / 0.1) * v2)) - v2)
                        switch fi28 {
                        case -1:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 128:
                            switch "\n\t" {
                            case "\n\t":
                                v2 = v2
                            case "":
                                v21 -= v21
                            case "a":
                                v2 -= ((((v2 - 0.0) + ((1e10 - 1.5) * v2)) + 0.0) / v2)
                            }
                        case -480:
                            if true {
                                v6 = 3.14159
                                var v73 int8 = 127
                                _ = v73
                            }
                        default:
                            v74 := GenS11{F0: 985}
                            _ = v74
                        }
                        v75 := GenS12{F0: 833, F1: 127}
                        _ = fi72
                        _ = v75
                    }
                default:
                    v76 := v4[("\n\t" + "foo bar")]
                    _ = v76
                }
                _ = fi28
            }
            v5 = append(v5, "")
        }
        var v77 bool = v20
        _ = v77
    } else {
        v6 = ((v6 * (0.0 * (float32(v2) / (v6 - 3.14159)))) - (-2.25 / -2.25))
    }
    var v78 uint8 = 62
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}

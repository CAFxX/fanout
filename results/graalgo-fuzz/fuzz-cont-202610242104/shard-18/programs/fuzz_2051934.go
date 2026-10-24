package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int8; F2 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 bool; F2 uint8 }
type GenS12 struct { F0 string; F1 string; F2 bool }

func main() {
    v1 := []int{122, 32767, -9223372036854775808}
    var v2 uint32 = 519
    v3 := make(map[uint32]string)
    var v4 uint64 = 18446744073709551615
    var v5 string = (("hello" + ("" + "\n\t")) + "")
    var v6 string = ((("foo bar" + v5) + v5) + (v5 + "foo bar"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        switch fi7 {
        case 0:
            var v8 int16 = 492
            _ = v8
        case 4294967295:
            fmt.Println(fi7)
        case -2147483648:
            v9 := GenS1{F0: ""}
            _ = v9
        }
        if true {
            v4 -= v4
        } else {
            for fi10 := 0; fi10 < 13; fi10++ {
                v6 = "a"
                if true {
                    if (274 >= 65535) {
                        v11 := GenS1{F0: v6}
                        _ = v11
                    }
                    if true {
                        switch 32768 {
                        case -528:
                            fmt.Println(fi10, fi10)
                        case 128:
                            v12 := v3[349]
                            var v13 float32 = 3.14159
                            var v14 int32 = -1
                            _ = v12
                            _ = v13
                            _ = v14
                        case -389:
                            if ((v4 * (v4 / ((869) | 1))) >= (v4 / ((1) | 1))) {
                                delete(v3, v2)
                                var v15 float32 = 0.1
                                v6 = (v6 + "foo bar")
                                _ = v15
                            }
                        }
                    } else {
                        if true {
                            delete(v3, 2147483647)
                            var v16 bool = true
                            _ = v16
                        } else {
                            var v17 int = fi10
                            v1 = append(v1, 2147483647)
                            var v18 int8 = 0
                            v19 := GenS2{F0: 961, F1: (v18 >> ((4294967295) & 63)), F2: 32767}
                            var v20 uint = 4294967295
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        v21 := v3[1]
                        v5 = ""
                        _ = v21
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = v2
                }
                var v22 int8 = 127
                fmt.Println(v22)
                if (-2147483648 <= -780) {
                    var v23 bool = (1e10 >= (0.1 - 1.5))
                    delete(v3, v2)
                    switch (^(fi10)) {
                    case 1:
                        var v24 int64 = 9223372036854775807
                        _ = v24
                    case 4294967295:
                        if true {
                            var v25 int64 = 42
                            v26 := v3[v2]
                            v23 = false
                            v27 := GenS3{F0: 1000, F1: ((((!((false && (!false)) && (42 > 22))) && (255 <= 448)) || false) && v23)}
                            v6 = ((v5 + v5) + v5)
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(fi10, fi7)
                            v6 = "a"
                            v5 = "a"
                        }
                        if (false && true) {
                            v28 := GenS1{F0: "foo bar"}
                            switch (^(fi10)) {
                            case 2147483647:
                                v22 = 42
                            case -239:
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        if ((318 < 516) && v23) {
                                            var v29 int8 = v22
                                            _ = v29
                                        }
                                    case "foo bar":
                                        v22 = v22
                                    case "a":
                                        var v30 uint8 = uint8(v2)
                                        _ = v30
                                    }
                                    v31 := GenS4{F0: 167, F1: (v22 - v22)}
                                    _ = v31
                                } else {
                                    if (float32(3.14159) > ((((1.5 - 3.14159) - 0.0) + 3.14159) / -2.25)) {
                                        if v23 {
                                            v2 += (^(v2))
                                            var v32 uint64 = v4
                                            v33 := v3[(v2 >> ((1) & 63))]
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := GenS5{F0: (fi7 - int(v2)), F1: 1}
                                        v4 -= (^(v4))
                                        var v35 uint64 = (v4 / ((v4) | 1))
                                        var v36 int = fi10
                                        v37 := GenS6{F0: 1}
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            case 32768:
                                v2 = (v2 >> ((1) & 63))
                            default:
                                v38 := GenS2{F0: 255, F1: int8(fi10), F2: 1}
                                _ = v38
                            }
                            var v39 float64 = (0.0 / 1.0)
                            _ = len(v3)
                            if ((v22 ^ (v22 | v22)) <= v22) {
                                v40 := GenS5{F0: fi10, F1: 2}
                                _ = v40
                            } else {
                                fmt.Println(v5, v22, v4)
                                if v23 {
                                    var v41 int32 = 42
                                    _ = len(v3)
                                    switch int(v41) {
                                    case -725:
                                        if (v23 || true) {
                                            var v42 bool = v23
                                            v41 -= -2147483648
                                            var v43 int = -1
                                            _ = v42
                                            _ = v43
                                        }
                                    }
                                    _ = v41
                                }
                            }
                            _ = v28
                            _ = v39
                        } else {
                            switch "hello" {
                            case "hello":
                                v44 := GenS4{F0: 2, F1: v22}
                                _ = v44
                            default:
                                v1 = append(v1, -2147483648)
                            }
                            v45 := GenS1{F0: (v5 + (v5 + "a"))}
                            var v46 int32 = int32(fi10)
                            _ = v45
                            _ = v46
                        }
                    }
                    switch (fi7 % (((fi10 + 0)) | 1)) {
                    case 52:
                        if ((((1e10 + 1e10) - 1e10) / float64(3.14159)) < 1e10) {
                            var v47 int64 = 9223372036854775807
                            if (false && true) {
                                v48 := GenS7{F0: v6}
                                _ = v48
                            } else {
                                switch (fi7 % ((749) | 1)) {
                                case 128:
                                    if v23 {
                                        v49 := v3[v2]
                                        var v50 uint16 = 0
                                        v51 := GenS4{F0: (v50 / (((v50 * (v50 | v50))) | 1)), F1: v22}
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                case 95:
                                    v47 = 42
                                case -1:
                                    fmt.Println(v6)
                                }
                            }
                            var v52 int32 = 641
                            _ = v47
                            _ = v52
                        } else {
                            switch (fi7 & (^(fi10))) {
                            case -9223372036854775808:
                                v23 = (-400 != 620)
                            default:
                                _ = len(v3)
                            }
                            _ = len(v3)
                            v53 := GenS6{F0: 0}
                            _ = v53
                        }
                    }
                    v5 = "a"
                    _ = v23
                } else {
                    if false {
                        _ = len(v3)
                        _ = len(v3)
                    } else {
                        var v54 uint16 = 32767
                        var v55 uint32 = (v2 ^ v2)
                        if false {
                            v56 := GenS6{F0: -1}
                            _ = v56
                        }
                        _ = v54
                        _ = v55
                    }
                    fmt.Println(v6)
                    if false {
                        switch v6 {
                        case "foo bar":
                            fmt.Println(v4, fi10)
                        case "a":
                            v5 = v5
                        case "\n\t":
                            var v57 bool = (0 < int32(v2))
                            _ = v57
                        default:
                            switch fi10 {
                            case 200:
                                fmt.Println(v4)
                            }
                        }
                        _ = len(v3)
                    } else {
                        var v58 uint = uint(v4)
                        if (32767 <= -995) {
                            var v59 float64 = float64(3.14159)
                            _ = v59
                        } else {
                            v1 = append(v1, fi7)
                        }
                        var v60 uint = (v58 >> ((v58) & 63))
                        _ = v58
                        _ = v60
                    }
                    v61 := GenS8{F0: -32768, F1: (^(v22)), F2: 0}
                    switch v6 {
                    case "foo bar":
                        switch fi10 {
                        case -253:
                            v5 = (v5 + (v6 + (v6 + ("\n\t" + v5))))
                        case 9223372036854775807:
                            v6 = ((v5 + "") + v5)
                        }
                    default:
                        v22 = v22
                    }
                    _ = v61
                }
                v62 := GenS5{F0: fi10, F1: 68}
                _ = fi10
                _ = v22
                _ = v62
            }
            if ((42 < 604) || ("a" == "hello")) {
                if (-32768 < 934) {
                    if true {
                        fmt.Println(v5, v2, v4)
                        v63 := GenS9{F0: (^(v4))}
                        _ = v63
                    } else {
                        v2 = v2
                        if true {
                            for fi64 := 0; fi64 < 1; fi64++ {
                                var v65 int64 = -781
                                v3[65535] = "foo bar"
                                var v66 bool = (float64(3.14159) == (3.14159 + -2.25))
                                var v67 int32 = 0
                                v5 = (((v6 + (v6 + (("" + v5) + "hello"))) + (((v5 + ("" + "")) + "a") + "hello")) + v6)
                                _ = fi64
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            for fi68 := 0; fi68 < 8; fi68++ {
                                if (-559 > 165) {
                                    var v69 int64 = 42
                                    var v70 uint64 = (v4 ^ (^(v4)))
                                    _ = v69
                                    _ = v70
                                }
                                _ = fi68
                            }
                        } else {
                            v71 := GenS10{F0: (fi7 ^ fi7)}
                            _ = len(v1)
                            _ = v71
                        }
                    }
                    switch -1 {
                    case 1:
                        v6 = v5
                    default:
                        var v72 uint64 = v4
                        _ = v72
                    }
                }
            }
            if true {
                if false {
                    v73 := GenS11{F0: 593, F1: true, F2: 128}
                    v74 := GenS12{F0: ((v6 + "a") + v5), F1: v5, F2: (!true)}
                    switch v6 {
                    case "\n\t":
                        fmt.Println(v6)
                    case "foo bar":
                        var v75 bool = true
                        _ = v75
                    }
                    _ = v73
                    _ = v74
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
            }
        }
        v2 += 2
        _ = fi7
    }
    for fi76 := 0; fi76 < 17; fi76++ {
        v5 = v5
        v6 = "hello"
        _ = fi76
    }
    v3[v2] = "hello"
    v5 = "foo bar"
    v6 = v5
    switch 1 {
    case -539:
        v2 -= 65535
    default:
        v4 = (^(v4))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v77 uint32 = (v2 % (((^(v2))) | 1))
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}

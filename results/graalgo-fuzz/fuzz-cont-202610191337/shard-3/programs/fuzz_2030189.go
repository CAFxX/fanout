package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint; F1 string; F2 uint }

func main() {
    v1 := []string{"foo bar", "\n\t", "\n\t"}
    var v2 int16 = -457
    v3 := make(map[string]uint32)
    var v4 int16 = v2
    var v5 int = 1
    var v6 uint8 = 130
    switch "\n\t" {
    case "\n\t":
        v3["a"] = 4294967295
    case "foo bar":
        var v7 bool = false
        _ = v7
    case "":
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 = (v5 << ((42) & 63))
            var v9 int = int(v6)
            _ = fi8
            _ = v9
        }
    default:
        if (0 == 255) {
            v4 = v2
        }
    }
    _ = len(v1)
    switch 128 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS1{F0: -18}
                v5 = -509
                var v13 uint = 949
                _ = fi11
                _ = v12
                _ = v13
            }
            v4 -= (v2 | (v2 >> ((65535) & 63)))
            _ = fi10
        }
    case 32768:
        var v14 uint8 = 127
        _ = v14
    default:
        v15 := GenS2{F0: v5, F1: (^(v5))}
        _ = v15
    }
    v2 = v4
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v16 := GenS2{F0: (-(v5)), F1: (v5 / (((^(v5))) | 1))}
    switch v5 {
    case 864:
        v17 := GenS1{F0: -96}
        _ = v17
    }
    fmt.Println(v4)
    if true {
        var v18 int64 = 42
        var v19 float64 = 0.0
        _ = v18
        _ = v19
    } else {
        var v20 int8 = 127
        if (2147483647 > 2147483647) {
            if (v20 > 60) {
                var v21 int32 = 42
                if (!((int(v6) > v5) && ((1e10 - -2.25) > float32(3.14159)))) {
                    var v22 uint8 = (v6 % ((v6) | 1))
                    _ = v22
                }
                _ = v21
            }
            switch v5 {
            case 1:
                var v23 uint = 717
                _ = v23
            case 9223372036854775807:
                for fi24 := 0; fi24 < 3; fi24++ {
                    switch ("a" + (("" + ((("" + ("a" + "")) + "hello") + "")) + "\n\t")) {
                    case "foo bar":
                        fmt.Println(v4, v2)
                    default:
                        fmt.Println(v2, v4, v5)
                    }
                    for fi25 := 0; fi25 < 1; fi25++ {
                        if true {
                            var v26 int64 = 0
                            var v27 uint16 = 2
                            switch "" {
                            case "hello":
                                v28 := v3["foo bar"]
                                _ = v28
                            case "\n\t":
                                switch fi25 {
                                case -1:
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        v20 = 127
                                    default:
                                        v3[""] = 2
                                    }
                                default:
                                    var v29 float32 = float32(float64(3.14159))
                                    _ = v29
                                }
                            case "":
                                fmt.Println(v4)
                            }
                            _ = v26
                            _ = v27
                        } else {
                            if (-2.25 > ((0.0 / ((float32(0.1) * 1e10) / float32(1e10))) * 0.1)) {
                                var v30 float64 = 0.0
                                switch 128 {
                                case -9223372036854775808:
                                    v31 := GenS1{F0: v20}
                                    _ = v31
                                case 9223372036854775807:
                                    switch fi25 {
                                    case 9223372036854775807:
                                        switch fi25 {
                                        case 32768:
                                            if true {
                                                v32 := v3["foo bar"]
                                                var v33 float64 = -2.25
                                                switch 32767 {
                                                case 1:
                                                    var v34 float64 = (v30 + v33)
                                                    _ = v34
                                                case 9223372036854775807:
                                                    var v35 uint64 = 4294967295
                                                    _ = v35
                                                }
                                                _ = v32
                                                _ = v33
                                            }
                                        case 16:
                                            if (!false) {
                                                var v36 uint16 = 41
                                                _ = len(v1)
                                                switch (fi24 + (^(fi24))) {
                                                case 2147483647:
                                                    if (true && ("a" == "hello")) {
                                                        if (!(!false)) {
                                                            v3["a"] = 77
                                                        }
                                                        var v37 uint = 255
                                                        _ = v37
                                                    } else {
                                                        if true {
                                                            v6 += v6
                                                            switch v5 {
                                                            case 442:
                                                                v20 = v20
                                                            case -1:
                                                                var v38 uint16 = (v36 | v36)
                                                                _ = v38
                                                            case -911:
                                                                v1 = append(v1, "\n\t")
                                                            }
                                                            if (0.1 == 1e10) {
                                                                var v39 string = ""
                                                                _ = v39
                                                            }
                                                        } else {
                                                            var v40 string = "foo bar"
                                                            fmt.Println(v36)
                                                            _ = v40
                                                        }
                                                        v41 := GenS1{F0: 1}
                                                        _ = v41
                                                    }
                                                case -2147483648:
                                                    v42 := GenS3{F0: "hello", F1: v6, F2: -9223372036854775808}
                                                    _ = v42
                                                case 32767:
                                                    if true {
                                                        v20 -= (-(v20))
                                                    } else {
                                                        v30 = v30
                                                        if (754 != v36) {
                                                            v3[""] = 863
                                                            v36 = (v36 / ((65535) | 1))
                                                        }
                                                        v5 = fi24
                                                        v6 -= (^(v6))
                                                    }
                                                }
                                                _ = v36
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v43 := GenS2{F0: fi25, F1: fi25}
                                                v44 := GenS1{F0: (^(v20))}
                                                _ = v43
                                                _ = v44
                                            }
                                        case 21:
                                            var v45 int64 = 42
                                            _ = v45
                                        }
                                        v5 += 42
                                        v46 := GenS3{F0: "", F1: v6, F2: 1}
                                        if true {
                                            v47 := GenS2{F0: fi24, F1: (v5 - (^(fi24)))}
                                            if (-1 < -221) {
                                                v5 += (fi24 | 0)
                                            }
                                            _ = v47
                                        } else {
                                            v3[("hello" + ("foo bar" + ((("a" + "a") + "hello") + "\n\t")))] = 1
                                            var v48 int8 = v20
                                            var v49 float32 = -2.25
                                            _ = v48
                                            _ = v49
                                        }
                                        var v50 uint16 = 2
                                        _ = v46
                                        _ = v50
                                    case 4294967295:
                                        var v51 bool = (true || (true || (true || (!true))))
                                        _ = v51
                                    default:
                                        v6 += v6
                                    }
                                case 32768:
                                    v52 := GenS3{F0: "foo bar", F1: (v6 - (v6 / (((v6 - v6)) | 1))), F2: 9223372036854775807}
                                    _ = v52
                                default:
                                    var v53 float64 = v30
                                    _ = v53
                                }
                                v30 = v30
                                _ = v30
                            } else {
                                var v54 float64 = 0.0
                                _ = v54
                            }
                            var v55 uint = 4294967295
                            _ = v55
                        }
                        v56 := v3[""]
                        v6 = v6
                        _ = fi25
                        _ = v56
                    }
                    _ = fi24
                }
            case -9223372036854775808:
                if ((v5 >> ((2) & 63)) == 4294967295) {
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        var v57 uint16 = 32767
                        _ = v57
                    case "a":
                        v20 += v20
                    case "\n\t":
                        fmt.Println(v4, v2, v5)
                    }
                    var v58 uint64 = 372
                    delete(v3, "a")
                    _ = v58
                } else {
                    v59 := GenS2{F0: (v5 | (-(v5))), F1: v5}
                    for fi60 := 0; fi60 < 7; fi60++ {
                        v20 -= 1
                        _ = fi60
                    }
                    _ = v59
                }
            default:
                var v61 string = "a"
                _ = v61
            }
        }
        v62 := GenS4{F0: v6, F1: -1}
        _ = v20
        _ = v62
    }
    if ((127 < 61) && true) {
        if true {
            v63 := GenS5{F0: false, F1: 4294967295}
            v64 := GenS6{F0: 18446744073709551615, F1: "", F2: 1}
            if (255 >= 0) {
                v65 := GenS2{F0: (^(v5)), F1: (v5 - v5)}
                for fi66 := 0; fi66 < 16; fi66++ {
                    v6 += v6
                    _ = len(v3)
                    var v67 bool = false
                    _ = fi66
                    _ = v67
                }
                _ = v65
            } else {
                fmt.Println(v5, v4)
                v2 += 1
            }
            _ = v63
            _ = v64
        }
        _ = len(v1)
        v4 = v4
    } else {
        var v68 uint64 = 18446744073709551615
        _ = v68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

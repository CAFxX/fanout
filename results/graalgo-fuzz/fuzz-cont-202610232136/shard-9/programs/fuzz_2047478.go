package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint }
type GenS6 struct { F0 string; F1 uint16; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int32; F1 uint; F2 uint16 }

func main() {
    v1 := []string{"\n\t", "a", ("foo bar" + "hello")}
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            v3 = v3
            fmt.Println(v3, fi6)
            switch (fi6 + fi6) {
            case -749:
                var v7 string = "foo bar"
                _ = v7
            case 42:
                v8 := GenS1{F0: 1, F1: v3, F2: int32(fi5)}
                _ = v8
            case 9223372036854775807:
                switch (-(fi6)) {
                case 9223372036854775807:
                    if (v3 == (v3 | (^(v3)))) {
                        if true {
                            v9 := v2["a"]
                            v10 := GenS1{F0: 0, F1: v3, F2: -83}
                            v11 := GenS1{F0: 1, F1: v3, F2: 2147483647}
                            if ("foo bar" < "") {
                                v9 = (v9 + v9)
                                var v12 bool = false
                                _ = v12
                            }
                            v4 -= 4294967295
                            if (!true) {
                                var v13 float32 = (1.5 - 3.14159)
                                var v14 uint32 = v9
                                var v15 uint64 = (^(v4))
                                var v16 uint16 = 32767
                                if (1000 >= -95) {
                                    v4 += v15
                                }
                                v4 -= v4
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            var v17 uint8 = 1
                            _ = v10
                            _ = v11
                            _ = v17
                            _ = v9
                        }
                        v4 = (v4 % ((2) | 1))
                        fmt.Println(fi6)
                        fmt.Println(v4)
                    }
                default:
                    v18 := GenS1{F0: 2, F1: (^(v3)), F2: 682}
                    _ = v18
                }
            }
            _ = fi6
        }
        _ = fi5
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v2[("\n\t" + "")] = 4294967295
        _ = fi19
    }
    v4 += 244
    if true {
        v4 = (v4 << (((v3 / (((v3 + v3)) | 1))) & 63))
        delete(v2, (("foo bar" + "\n\t") + ("" + "")))
    } else {
        v20 := GenS1{F0: 127, F1: v3, F2: 0}
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint64 = (v4 - (v4 % ((2) | 1)))
            var v23 float64 = (0.0 / 1.5)
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 int64 = 1
                v22 -= (^(v22))
                _ = fi24
                _ = v25
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        var v26 int = 1
        _ = v20
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        if false {
            var v28 uint64 = 18446744073709551615
            if false {
                v28 = (v28 + (v28 << ((v3) & 63)))
                fmt.Println(v3, fi27)
            }
            _ = v28
        }
        _ = fi27
    }
    switch "\n\t" {
    case "":
        v4 = 63
    case "\n\t":
        for fi29 := 0; fi29 < 4; fi29++ {
            v2[("hello" + "hello")] = 1
            v30 := GenS2{F0: 65535, F1: 126}
            v3 -= v3
            _ = fi29
            _ = v30
        }
    case "a":
        var v31 uint64 = (^(v4))
        _ = v31
    default:
        fmt.Println(v4, v4)
    }
    v32 := GenS3{F0: -944}
    if false {
        var v33 bool = false
        for fi34 := 0; fi34 < 18; fi34++ {
            v1 = append(v1, "foo bar")
            v33 = (!false)
            _ = fi34
        }
        var v35 uint64 = v4
        _ = v33
        _ = v35
    }
    var v36 uint32 = 2147483647
    v3 = 2
    if ((3.14159 * 0.1) < 1e10) {
        if false {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
    } else {
        v37 := GenS1{F0: 127, F1: (v3 - v3), F2: 42}
        for fi38 := 0; fi38 < 11; fi38++ {
            v36 += v36
            var v39 bool = (1 < (v36 << (((v3 >> ((v3) & 63))) & 63)))
            fmt.Println(v39)
            _ = fi38
            _ = v39
        }
        if true {
            if true {
                var v40 int64 = 822
                _ = len(v2)
                _ = v40
            } else {
                v41 := GenS4{F0: uint8(v4)}
                switch -2147483648 {
                case 0:
                    v3 += 2
                case -2147483648:
                    if (32767 <= 1) {
                        if (42 != -2) {
                            v42 := GenS1{F0: uint8(v36), F1: (v3 + 4294967295), F2: -2147483648}
                            _ = v42
                        } else {
                            v2["a"] = 1
                        }
                    }
                default:
                    if (-2147483648 < 448) {
                        if (false || false) {
                            if (0.0 != -2.25) {
                                v4 -= (v4 * (v4 + (v4 - (^(v4)))))
                                for fi43 := 0; fi43 < 5; fi43++ {
                                    switch (("a" + "foo bar") + "") {
                                    case "":
                                        delete(v2, (("\n\t" + "\n\t") + "foo bar"))
                                    case "a":
                                        for fi44 := 0; fi44 < 20; fi44++ {
                                            v4 += 42
                                            v45 := GenS2{F0: v36, F1: 32767}
                                            _ = fi44
                                            _ = v45
                                        }
                                    case "foo bar":
                                        if false {
                                            if (("a" + (("hello" + "hello") + "hello")) != "hello") {
                                                if len(v1) > 0 {
                                                    v1[0] = ("" + "hello")
                                                }
                                                v46 := GenS4{F0: 255}
                                                _ = v46
                                            }
                                        } else {
                                            for fi47 := 0; fi47 < 20; fi47++ {
                                                v48 := GenS4{F0: 1}
                                                _ = fi47
                                                _ = v48
                                            }
                                            var v49 int = fi43
                                            var v50 float64 = (0.0 + 1.5)
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    _ = fi43
                                }
                                v2[""] = v36
                            }
                            if true {
                                var v51 uint = v3
                                _ = v51
                            }
                        }
                    }
                }
                v52 := GenS5{F0: uint8(v3), F1: (2 <= 0), F2: v3}
                _ = v41
                _ = v52
            }
            if ((2147483647 != (v36 / ((922) | 1))) && (42 < 42)) {
                delete(v2, ("a" + ""))
                var v53 string = "\n\t"
                _ = v53
            }
            v54 := GenS6{F0: ("foo bar" + "\n\t"), F1: uint16(v4), F2: "hello"}
            _ = v54
        } else {
            v36 = 585
            if false {
                if false {
                    for fi55 := 0; fi55 < 8; fi55++ {
                        v3 += (v3 | 890)
                        v3 = 1
                        v56 := GenS7{F0: 0}
                        _ = fi55
                        _ = v56
                    }
                }
                v57 := v2["\n\t"]
                _ = v57
            } else {
                if ((("hello" + "hello") + "\n\t") >= "\n\t") {
                    var v58 float32 = 3.14159
                    if false {
                        _ = len(v2)
                        v59 := GenS8{F0: v4}
                        _ = v59
                    } else {
                        switch "a" {
                        case "":
                            switch 1 {
                            case 624:
                                for fi60 := 0; fi60 < 11; fi60++ {
                                    v61 := GenS3{F0: 0}
                                    var v62 uint16 = 0
                                    v36 += (v36 - v36)
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                            case 32768:
                                var v63 int32 = -995
                                _ = v63
                            }
                        default:
                            if true {
                                fmt.Println(v3)
                                if true {
                                    fmt.Println(v4, v4)
                                } else {
                                    for fi64 := 0; fi64 < 6; fi64++ {
                                        v2["foo bar"] = v36
                                        _ = len(v2)
                                        var v65 string = "foo bar"
                                        _ = fi64
                                        _ = v65
                                    }
                                    for fi66 := 0; fi66 < 13; fi66++ {
                                        v3 = 589
                                        _ = fi66
                                    }
                                    var v67 int32 = -850
                                    _ = v67
                                }
                            }
                        }
                        if false {
                            v68 := GenS7{F0: 1}
                            _ = v68
                        }
                        var v69 string = "a"
                        _ = v69
                    }
                    if ("a" >= ("hello" + "foo bar")) {
                        fmt.Println(v58)
                    }
                    _ = v58
                }
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
            for fi70 := 0; fi70 < 18; fi70++ {
                v36 = (^(v36))
                for fi71 := 0; fi71 < 4; fi71++ {
                    v3 -= v3
                    v4 -= uint64(v36)
                    _ = len(v2)
                    _ = fi71
                }
                _ = fi70
            }
        }
        _ = v37
    }
    fmt.Println(v4, v4)
    switch "\n\t" {
    case "":
        var v72 uint64 = (^(v4))
        _ = v72
    case "a":
        for fi73 := 0; fi73 < 15; fi73++ {
            v74 := GenS9{F0: (fi73 * (fi73 + 9223372036854775807)), F1: (v36 - (v36 | v36))}
            _ = fi73
            _ = v74
        }
    case "\n\t":
        var v75 int = -139
        _ = v75
    default:
        v76 := GenS10{F0: 341, F1: (v3 | (v3 ^ v3)), F2: 436}
        _ = v76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
}

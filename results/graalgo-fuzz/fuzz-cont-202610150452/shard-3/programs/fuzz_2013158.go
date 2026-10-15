package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 string; F2 string }
type GenS5 struct { F0 uint8; F1 int64; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 string; F1 uint8; F2 uint }
type GenS10 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"", ((("hello" + "foo bar") + (("foo bar" + "") + "")) + ("a" + "\n\t")), "a", ("\n\t" + ((("hello" + "a") + (("foo bar" + (("" + "hello") + ("\n\t" + "\n\t"))) + "a")) + ("hello" + ("" + "hello"))))}
    v3 := make(map[string]string)
    var v4 uint32 = 2
    v4 = 0
    if false {
        if true {
            switch 4294967295 {
            case -2147483648:
                if false {
                    switch 0 {
                    case -1:
                        delete(v1, "foo bar")
                    default:
                        _ = len(v1)
                    }
                    v4 = (v4 - v4)
                } else {
                    v4 -= v4
                    var v5 float64 = -2.25
                    switch (("foo bar" + "foo bar") + "hello") {
                    case "foo bar":
                        v6 := GenS1{F0: ((((0.1 / ((1e10 + ((0.1 - 0.0) + (0.0 - ((3.14159 - 1.5) / ((0.1 + (1e10 * (float32(v5) / 0.1))) * -2.25))))) - ((0.1 + float32(v5)) * ((0.0 + ((3.14159 + 1.5) / 1.0)) / 1.0)))) / 3.14159) != 1.5) && (true || (-403 < 4294967295))), F1: v4, F2: 0}
                        _ = v6
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            var v7 int = -157
                            _ = v7
                        case "a":
                            var v8 int8 = 1
                            _ = v8
                        case "hello":
                            v4 -= (v4 / (((v4 + (^(v4)))) | 1))
                        }
                    default:
                        for fi9 := 0; fi9 < 15; fi9++ {
                            for fi10 := 0; fi10 < 14; fi10++ {
                                v4 = (^(v4))
                                if false {
                                    var v11 uint32 = (v4 ^ v4)
                                    v12 := GenS2{F0: 127, F1: -567}
                                    _ = v11
                                    _ = v12
                                } else {
                                    if false {
                                        v1["foo bar"] = ("foo bar" + "foo bar")
                                        v3[(("foo bar" + "") + ("hello" + "foo bar"))] = "\n\t"
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        var v13 int = -281
                                        if false {
                                            var v14 int8 = 100
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v13)
                                            v15 := GenS3{F0: 1}
                                            v4 = 115
                                            v14 = v14
                                            fmt.Println(fi10, v4)
                                            _ = v14
                                            _ = v15
                                        } else {
                                            var v16 int32 = -729
                                            var v17 uint32 = 2147483647
                                            switch "hello" {
                                            case "":
                                                if true {
                                                    v18 := GenS3{F0: 1}
                                                    _ = v18
                                                }
                                            case "\n\t":
                                                fmt.Println(v13, fi10, v16)
                                            case "foo bar":
                                                var v19 uint8 = 2
                                                v4 += 4294967295
                                                switch fi9 {
                                                case 2147483647:
                                                    v20 := GenS4{F0: "a", F1: "", F2: ("\n\t" + ("foo bar" + "a"))}
                                                    _ = v20
                                                case 586:
                                                    v21 := GenS5{F0: v19, F1: -967, F2: 1000}
                                                    _ = v21
                                                case -930:
                                                    v22 := GenS4{F0: ("\n\t" + "a"), F1: ("a" + "\n\t"), F2: "\n\t"}
                                                    _ = v22
                                                }
                                                _ = v19
                                            default:
                                                _ = len(v2)
                                            }
                                            v4 = v4
                                            _ = v16
                                            _ = v17
                                        }
                                        if (!true) {
                                            v13 = (fi9 * (v13 / (((v13 ^ fi9)) | 1)))
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(fi9, v5, v5)
                                            case "":
                                                v5 += 0.1
                                            case "a":
                                                var v23 float32 = (((((3.14159 * 3.14159) * (0.1 * 3.14159)) - (3.14159 + ((1e10 * 0.0) + (1e10 + 0.0)))) / 1e10) + 0.1)
                                                _ = v23
                                            }
                                            if ("hello" <= "a") {
                                                switch "foo bar" {
                                                case "":
                                                    switch "" {
                                                    case "\n\t":
                                                        v3["\n\t"] = ""
                                                    default:
                                                        v4 = (^(v4))
                                                    }
                                                }
                                            } else {
                                                var v24 uint64 = 18446744073709551615
                                                _ = v24
                                            }
                                            v25 := v1[""]
                                            _ = v25
                                        }
                                        _ = v13
                                    }
                                    v26 := GenS6{F0: 42}
                                    _ = v26
                                }
                                _ = fi10
                            }
                            _ = fi9
                        }
                    }
                    _ = v5
                }
            case 2147483647:
                var v27 float64 = 1.5
                _ = v27
            default:
                for fi28 := 0; fi28 < 11; fi28++ {
                    v4 = v4
                    _ = fi28
                }
            }
            if len(v2) > 0 {
                v2[0] = "a"
            }
        } else {
            for fi29 := 0; fi29 < 10; fi29++ {
                if false {
                    var v30 int = fi29
                    _ = v30
                }
                var v31 uint8 = 149
                v4 = v4
                _ = fi29
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint8 = 242
            v1["a"] = ("" + ("" + "foo bar"))
            _ = fi32
            _ = v33
        }
    } else {
        fmt.Println(v4)
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        fmt.Println(v4)
    case "":
        for fi34 := 0; fi34 < 5; fi34++ {
            v2 = append(v2, "foo bar")
            _ = fi34
        }
    }
    var v35 int8 = 42
    if false {
        var v36 uint8 = 42
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        if (0 > 32767) {
            fmt.Println(v4, v35, fi37)
            if (1 <= 0) {
                v4 -= v4
            } else {
                var v38 uint = uint(v4)
                _ = v38
            }
            v2 = append(v2, "\n\t")
        } else {
            for fi39 := 0; fi39 < 12; fi39++ {
                if (1000 >= 42) {
                    v40 := GenS5{F0: 255, F1: 0, F2: -1}
                    _ = v40
                }
                var v41 uint64 = 874
                v41 -= (v41 << ((439) & 63))
                _ = fi39
                _ = v41
            }
            if true {
                v35 = (-(v35))
                switch "" {
                case "a":
                    v2 = append(v2, ("hello" + ("" + "")))
                case "hello":
                    if true {
                        fmt.Println(fi37)
                    }
                case "\n\t":
                    v4 -= v4
                }
            } else {
                switch -963 {
                case 2147483647:
                    var v42 uint32 = v4
                    _ = v42
                }
            }
        }
        v43 := GenS7{F0: "foo bar"}
        _ = fi37
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 uint32 = v4
        var v46 uint32 = (v4 & (v4 | (v4 + (v45 << ((585) & 63)))))
        for fi47 := 0; fi47 < 6; fi47++ {
            switch (fi47 * fi47) {
            case 128:
                v48 := v3[("a" + "\n\t")]
                _ = v48
            case 113:
                v45 = (v46 + 2147483648)
            case 32768:
                fmt.Println(fi47)
                var v49 uint64 = 26
                var v50 uint16 = 378
                v4 = 473
                _ = v49
                _ = v50
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch (fi44 & fi44) {
            case -1:
                if (!true) {
                    if false {
                        v45 = (^(v45))
                        if true {
                            var v51 int64 = 1
                            var v52 uint = 255
                            v45 += (^(v45))
                            _ = v51
                            _ = v52
                        } else {
                            var v53 float64 = (0.1 - (1e10 / (3.14159 - 1.5)))
                            if (404 < 0) {
                                if false {
                                    v4 = v46
                                } else {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        _ = len(v1)
                                    }
                                }
                                fmt.Println(v35, v53)
                                v4 += (v45 + 42)
                            } else {
                                v35 = v35
                                v54 := GenS8{F0: 0, F1: v46}
                                var v55 bool = true
                                v56 := GenS7{F0: ("\n\t" + "")}
                                v4 = v46
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            _ = v53
                        }
                        if false {
                            v45 += v46
                        }
                    } else {
                        fmt.Println(v46, v45)
                        v57 := GenS9{F0: ("a" + "\n\t"), F1: 111, F2: 2}
                        var v58 uint32 = 65535
                        _ = v57
                        _ = v58
                    }
                } else {
                    v46 = (^(v46))
                    v59 := GenS4{F0: "\n\t", F1: "hello", F2: "\n\t"}
                    fmt.Println(fi47)
                    _ = v59
                }
            case -9223372036854775808:
                v2 = append(v2, (("hello" + ("a" + "hello")) + ""))
                if false {
                    var v60 int8 = 1
                    _ = v60
                } else {
                    v4 += v46
                }
            case 1:
                v61 := GenS10{F0: (v4 * 65535), F1: (^(v45))}
                _ = v61
            default:
                var v62 uint16 = 32767
                _ = v62
            }
            _ = fi47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    v35 += v35
    v63 := GenS8{F0: 42, F1: v4}
    if (!(4294967295 == 65535)) {
        v64 := v3["a"]
        _ = v64
    } else {
        var v65 bool = false
        var v66 uint8 = 0
        var v67 uint8 = v66
        _ = v65
        _ = v66
        _ = v67
    }
    var v68 int16 = 42
    v35 -= (^(v35))
    var v69 uint32 = v4
    var v70 bool = false
    _ = len(v2)
    if v70 {
        var v71 float32 = 3.14159
        var v72 uint = 65535
        _ = v71
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v63
    _ = v68
    _ = v69
    _ = v70
}

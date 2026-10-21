package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 bool; F1 string; F2 int }
type GenS9 struct { F0 int16; F1 bool; F2 uint32 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{(("a" + "hello") + "a"), ("hello" + "foo bar")}
    var v3 int32 = 42
    var v4 int8 = -125
    var v5 uint32 = 4294967295
    v5 = 2147483647
    fmt.Println(v4)
    v4 = (v4 * v4)
    v4 = -1
    v6 := GenS1{F0: (^(v5)), F1: ("hello" + "a")}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 80
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            if true {
                var v10 uint8 = 0
                _ = v10
            } else {
                var v11 int32 = v3
                _ = v11
            }
            v4 = 51
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if ((!false) && true) {
        fmt.Println(v3, v4, v4)
        if false {
            v12 := GenS1{F0: v5, F1: "foo bar"}
            v1[("hello" + (("foo bar" + "\n\t") + ""))] = (!true)
            _ = v12
        }
        switch 42 {
        case -690:
            v13 := GenS1{F0: v5, F1: "a"}
            _ = v13
        case 1:
            if true {
                var v14 int32 = (v3 | -2147483648)
                fmt.Println(v4, v14, v5)
                _ = v14
            } else {
                v4 -= 0
                if (128 != 255) {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if (true || (-2.25 <= 1e10)) {
                            switch "a" {
                            case "a":
                                v16 := GenS1{F0: 42, F1: "a"}
                                _ = v16
                            case "foo bar":
                                if false {
                                    for fi17 := 0; fi17 < 17; fi17++ {
                                        v18 := GenS2{F0: (fi17 << ((2) & 63))}
                                        var v19 uint8 = 0
                                        if (812 > 65535) {
                                            v20 := GenS3{F0: (v5 * (v5 / (((v5 / ((uint32(v19)) | 1))) | 1))), F1: 1000, F2: 42}
                                            if (0 == -413) {
                                                var v21 uint32 = v5
                                                v22 := GenS1{F0: uint32(v19), F1: "hello"}
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if true {
                                                    v23 := GenS1{F0: (v5 / ((v5) | 1)), F1: "a"}
                                                    v24 := GenS4{F0: 922, F1: (v19 & 128), F2: 42}
                                                    _ = v23
                                                    _ = v24
                                                }
                                                v25 := GenS5{F0: 2}
                                                v26 := GenS3{F0: 2147483647, F1: -467, F2: -408}
                                                _ = v25
                                                _ = v26
                                            }
                                            _ = v20
                                        } else {
                                            if false {
                                                v4 = v4
                                                if ((v5 < v5) || true) {
                                                    v19 = v19
                                                    v3 += int32(v19)
                                                }
                                            } else {
                                                var v27 uint64 = uint64(v19)
                                                _ = v27
                                            }
                                        }
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                    }
                                    switch fi15 {
                                    case 446:
                                        var v28 float64 = 1e10
                                        _ = v28
                                    case -654:
                                        fmt.Println(v3, v3)
                                    case 4294967295:
                                        v4 += -88
                                    }
                                }
                            default:
                                for fi29 := 0; fi29 < 3; fi29++ {
                                    var v30 int8 = 42
                                    v5 += (v5 * v5)
                                    if ((1.5 > (float32((0.1 * ((-2.25 / float64(1e10)) / (1.5 + 1e10)))) / (-2.25 + -2.25))) && true) {
                                        if false {
                                            switch "foo bar" {
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                            case "":
                                                v31 := GenS3{F0: (v5 | 0), F1: 880, F2: 907}
                                                _ = v31
                                            case "hello":
                                                if true {
                                                    if false {
                                                        v4 = -128
                                                        var v32 float64 = float64(float32((((0.1 / 3.14159) * float64(3.14159)) - (float64(0.0) + float64((0.1 * 0.0))))))
                                                        fmt.Println(v3)
                                                        _ = v32
                                                    }
                                                }
                                            default:
                                                v33 := GenS3{F0: (v5 + (v5 + (v5 + (v5 << ((1) & 63))))), F1: -1, F2: -775}
                                                _ = v33
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (!(!false)) {
                                                    v30 += v30
                                                    var v34 float32 = (-2.25 * float32(1e10))
                                                    var v35 int = 1
                                                    _ = v34
                                                    _ = v35
                                                }
                                            default:
                                                v36 := GenS4{F0: 1, F1: 114, F2: 22}
                                                _ = v36
                                            }
                                            if ("a" > ("a" + "")) {
                                                switch fi29 {
                                                case 0:
                                                    if (!(false || true)) {
                                                        v30 = v4
                                                        v37 := GenS6{F0: "foo bar"}
                                                        _ = v37
                                                    }
                                                case -190:
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v38 float64 = 1.5
                                                        _ = v38
                                                    case "":
                                                        v4 = (v30 >> ((633) & 63))
                                                    case "\n\t":
                                                        var v39 float64 = (float64(((3.14159 + 3.14159) * 0.0)) - 0.0)
                                                        _ = v39
                                                    }
                                                case -1:
                                                    v40 := GenS7{F0: (v30 ^ v4), F1: v5}
                                                    _ = v40
                                                default:
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = "\n\t"
                                                        }
                                                    }
                                                }
                                                var v41 float32 = (1e10 + 0.0)
                                                _ = v41
                                            }
                                            var v42 uint32 = (v5 << ((255) & 63))
                                            _ = v42
                                        }
                                        v3 += v3
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            }
                        }
                        switch ((("a" + ("hello" + "foo bar")) + "a") + "") {
                        case "":
                            if ((v4 % ((v4) | 1)) > (v4 >> ((18446744073709551615) & 63))) {
                                fmt.Println(fi15, fi15, v3)
                            } else {
                                v43 := GenS8{F0: false, F1: "\n\t", F2: fi15}
                                fmt.Println(v4)
                                if true {
                                    v3 += v3
                                }
                                _ = v43
                            }
                        }
                        v44 := GenS9{F0: 316, F1: true, F2: v5}
                        _ = fi15
                        _ = v44
                    }
                    if ("foo bar" > ("" + "\n\t")) {
                        fmt.Println(v4)
                    }
                }
                switch 32767 {
                case 626:
                    v45 := v1["hello"]
                    _ = v45
                case -1:
                    v5 = 1
                default:
                    v4 -= v4
                }
            }
        }
    } else {
        v46 := GenS6{F0: ("a" + "a")}
        fmt.Println(v5)
        _ = v46
    }
    v3 = v3
    switch "foo bar" {
    case "foo bar":
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            if true {
                for fi47 := 0; fi47 < 20; fi47++ {
                    fmt.Println(v3)
                    var v48 uint16 = 159
                    for fi49 := 0; fi49 < 1; fi49++ {
                        if false {
                            var v50 int64 = 1
                            var v51 uint = 18446744073709551615
                            _ = v50
                            _ = v51
                        } else {
                            v3 = v3
                            var v52 int = (-(fi47))
                            _ = v52
                        }
                        if false {
                            switch fi49 {
                            case 42:
                                v1[""] = false
                            case 2147483647:
                                delete(v1, "hello")
                            }
                        } else {
                            fmt.Println(v5, v3)
                            if true {
                                fmt.Println(v4, v48, fi49)
                                v53 := GenS6{F0: "\n\t"}
                                _ = v53
                            } else {
                                v4 -= v4
                                if false {
                                    if true {
                                        var v54 int = 4294967295
                                        v3 = v3
                                        _ = v54
                                    } else {
                                        v4 += (v4 ^ int8(v5))
                                    }
                                }
                                var v55 int = 127
                                var v56 string = ("foo bar" + "a")
                                _ = v55
                                _ = v56
                            }
                        }
                        _ = fi49
                    }
                    _ = fi47
                    _ = v48
                }
                v3 = v3
            }
        }
    case "":
        var v57 int64 = 0
        _ = v57
    }
    v3 = v3
    v58 := GenS10{F0: v3, F1: 128, F2: 32}
    fmt.Println(v3, v5)
    fmt.Println(v4, v5)
    for fi59 := 0; fi59 < 11; fi59++ {
        v4 -= 100
        v1["hello"] = true
        _ = fi59
    }
    v3 = (^(v3))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}

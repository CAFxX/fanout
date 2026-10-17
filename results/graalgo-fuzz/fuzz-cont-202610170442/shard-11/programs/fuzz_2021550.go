package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint64, p1 uint) int32 {
    var v6 uint8 = 128
    fmt.Println(p0)
    v7 := []int{-655}
    var v8 float32 = 1.5
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 uint64 = 282
    v2 := make(map[string]string)
    v3 := []string{"foo bar", ("foo bar" + (((("foo bar" + "foo bar") + (("a" + ((((((("a" + "") + "hello") + "") + "hello") + "hello") + "a") + "foo bar")) + (("" + ("a" + "a")) + (("hello" + "") + "")))) + "foo bar") + (("a" + "a") + "a")))}
    var v4 uint16 = 1
    var v5 int8 = 127
    v9 := v2[""]
    v4 -= v4
    var v10 int16 = -32768
    switch 128 {
    case -658:
        v3 = append(v3, (v9 + (("\n\t" + v9) + v9)))
    }
    fmt.Println(v4, v4)
    fmt.Println(v10, v4)
    var v11 int32 = -277
    if true {
        fmt.Println(v4, v11)
        if (479 < 747) {
            v11 -= v11
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int = (fi12 / (((fi12 + fi12)) | 1))
                fmt.Println(v13)
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: "foo bar", F1: 2147483647}
            _ = v14
        } else {
            if true {
                if false {
                    var v15 float32 = 0.1
                    if (v9 == "") {
                        v10 = v10
                        if true {
                            if (!(false && (!false))) {
                                if (false && false) {
                                    var v16 bool = false
                                    for fi17 := 0; fi17 < 8; fi17++ {
                                        for fi18 := 0; fi18 < 3; fi18++ {
                                            v19 := GenS1{F0: "a", F1: 42}
                                            _ = fi18
                                            _ = v19
                                        }
                                        var v20 int16 = int16(v1)
                                        switch fi17 {
                                        case 0:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = fi17
                                        _ = v20
                                    }
                                    switch v9 {
                                    case "a":
                                        v21 := GenS1{F0: (v9 + "foo bar"), F1: 72}
                                        _ = v21
                                    }
                                    _ = v16
                                }
                                v10 -= v10
                            }
                            v4 = (v4 >> ((uint(v4)) & 63))
                            _ = len(v2)
                        }
                    } else {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            var v23 int32 = v11
                            switch v9 {
                            case "a":
                                for fi24 := 0; fi24 < 12; fi24++ {
                                    v25 := GenS1{F0: (("\n\t" + "") + ("hello" + v9)), F1: 268}
                                    _ = fi24
                                    _ = v25
                                }
                            case "hello":
                                v23 = v11
                            default:
                                for fi26 := 0; fi26 < 19; fi26++ {
                                    v27 := GenS1{F0: ((v9 + v9) + ("\n\t" + ((("hello" + (v9 + ("foo bar" + v9))) + (v9 + "\n\t")) + (v9 + "a")))), F1: 1}
                                    v28 := GenS2{F0: 4294967295}
                                    v9 = "foo bar"
                                    var v29 uint16 = v4
                                    v9 = (v9 + ((v9 + ((("\n\t" + v9) + ((v9 + v9) + (v9 + v9))) + v9)) + v9))
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            }
                            if false {
                                v3 = append(v3, v9)
                            } else {
                                var v30 uint = 2
                                _ = v30
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    fmt.Println(v15)
                    _ = v15
                } else {
                    v31 := GenS1{F0: v9, F1: 918}
                    _ = v31
                }
                v32 := GenS3{F0: (v1 - (v1 % ((635) | 1))), F1: (^(v4))}
                v5 = (v5 % (((v5 - v5)) | 1))
                _ = v32
            } else {
                if (!(!true)) {
                    var v33 uint16 = (v4 | v4)
                    if (1 == (v5 * (^(v5)))) {
                        v33 = (v33 * v33)
                        var v34 uint16 = v33
                        var v35 int64 = 0
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS3{F0: v1, F1: 65535}
                    _ = v33
                    _ = v36
                }
            }
            v37 := GenS2{F0: 42}
            _ = v37
        }
        var v38 int64 = 42
        _ = v38
    } else {
        _ = len(v2)
        v1 = 4294967295
    }
    if false {
        v10 += (v10 / ((32767) | 1))
        if ((("" + v9) + v9) <= v9) {
            fmt.Println(v9)
            switch (v9 + "foo bar") {
            case "":
                v39 := GenS1{F0: "", F1: 4294967295}
                _ = v39
            case "foo bar":
                var v40 uint8 = 127
                _ = v40
            default:
                for fi41 := 0; fi41 < 13; fi41++ {
                    v5 += v5
                    _ = fi41
                }
            }
            var v42 string = v9
            _ = v42
        }
        v43 := GenS4{F0: 851}
        _ = v43
    } else {
        for fi44 := 0; fi44 < 9; fi44++ {
            fmt.Println(v11)
            v2[v9] = v9
            if false {
                var v45 float32 = 0.0
                for fi46 := 0; fi46 < 1; fi46++ {
                    var v47 string = v9
                    _ = fi46
                    _ = v47
                }
                var v48 int8 = (v5 ^ (v5 ^ (-(v5))))
                _ = v45
                _ = v48
            } else {
                fmt.Println(v9, v10)
                v11 -= 42
                for fi49 := 0; fi49 < 15; fi49++ {
                    switch v9 {
                    case "hello":
                        v10 = v10
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "\n\t":
                        if true {
                            v2[v9] = v9
                            var v50 uint = 0
                            v51 := GenS3{F0: (v1 - (v1 | (v1 << ((v50) & 63)))), F1: 65535}
                            _ = v50
                            _ = v51
                        }
                    }
                    v52 := GenS1{F0: v9, F1: 65535}
                    _ = fi49
                    _ = v52
                }
            }
            _ = fi44
        }
        fmt.Println(v5)
    }
    v53 := GenS3{F0: v1, F1: v4}
    var v54 uint = 255
    var v55 int8 = (v5 << ((v54) & 63))
    if true {
        switch 9223372036854775807 {
        case 4294967295:
            var v56 uint8 = 255
            _ = v56
        case -867:
            switch -2147483648 {
            case 32767:
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "foo bar")
                }
            default:
                v3 = append(v3, "")
            }
        case 465:
            if (!false) {
                for fi57 := 0; fi57 < 2; fi57++ {
                    for fi58 := 0; fi58 < 4; fi58++ {
                        fmt.Println(v11, v5, v10)
                        _ = fi58
                    }
                    switch v9 {
                    case "hello":
                        for fi59 := 0; fi59 < 19; fi59++ {
                            var v60 uint32 = 469
                            _ = fi59
                            _ = v60
                        }
                    default:
                        for fi61 := 0; fi61 < 14; fi61++ {
                            switch (v9 + (v9 + v9)) {
                            case "\n\t":
                                fmt.Println(v5, fi61, v11)
                            }
                            var v62 float64 = 1e10
                            var v63 int64 = 42
                            v64 := GenS5{F0: false, F1: 42}
                            v1 += (v1 % (((v1 / ((v1) | 1))) | 1))
                            if (genFn0((v1 & v1), v54) != v11) {
                                if (!(true || true)) {
                                    if ((!true) && false) {
                                        v11 -= v11
                                    } else {
                                        if (v54 > v54) {
                                            if false {
                                                v65 := v2[v9]
                                                _ = v65
                                            } else {
                                                var v66 bool = true
                                                _ = v66
                                            }
                                        }
                                        v62 = -2.25
                                        v63 -= (v63 & v63)
                                    }
                                    switch (fi57 / (((fi57 * (fi57 + fi57))) | 1)) {
                                    case 42:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        v62 = ((v62 + ((v62 + 1.5) + v62)) * v62)
                                    }
                                    v67 := v2[v9]
                                    var v68 int32 = genFn0((v1 ^ (v1 & v1)), v54)
                                    _ = v67
                                    _ = v68
                                } else {
                                    v69 := GenS5{F0: (float32(v62) >= 0.0), F1: 933}
                                    var v70 int8 = v55
                                    v71 := v2["foo bar"]
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                            }
                            _ = fi61
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                    }
                    v10 = -800
                    _ = fi57
                }
            }
        default:
            delete(v2, v9)
        }
    }
    v72 := GenS1{F0: v9, F1: 0}
    for fi73 := 0; fi73 < 8; fi73++ {
        for fi74 := 0; fi74 < 14; fi74++ {
            var v75 uint16 = uint16(v1)
            _ = fi74
            _ = v75
        }
        _ = fi73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v72
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 int8) string {
    p0 -= p0
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return "a"
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{100, 42, 2147483648}
    var v3 bool = false
    var v4 uint16 = 280
    if (true && ((v1 >> ((0) & 63)) != 100)) {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 int8 = v1
            var v7 float32 = 3.14159
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        var v8 int8 = v1
        v1 -= 40
        _ = v8
    }
    fmt.Println(v4)
    if (18446744073709551615 == 0) {
        var v9 int = 4294967295
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            var v11 int8 = (v1 >> ((0) & 63))
            _ = v11
        } else {
            v12 := GenS1{F0: 42, F1: v3, F2: 32767}
            var v13 uint16 = (^(v4))
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 int32 = 1
    v15 := GenS1{F0: 4294967295, F1: true, F2: -1}
    for fi16 := 0; fi16 < 19; fi16++ {
        _ = len(v2)
        _ = fi16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int16 = -32768
        fmt.Println(v1)
        var v19 string = "a"
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v3)
    var v20 int32 = 0
    var v21 int32 = -601
    switch ("foo bar" + "") {
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            switch "foo bar" {
            case "hello":
                v23 := GenS1{F0: 0, F1: (1.5 != 1e10), F2: -32768}
                _ = v23
            case "foo bar":
                v24 := GenS2{F0: "a"}
                _ = v24
            }
            v2 = append(v2, 946)
            _ = fi22
        }
    case "\n\t":
        var v25 uint64 = 685
        _ = v25
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v2 = append(v2, 949)
    }
    var v26 int = 32767
    for fi27 := 0; fi27 < 2; fi27++ {
        if len(v2) > 0 {
            v2[0] = 42
        }
        v3 = (((0.1 - (float64(float32(1.5)) / float64((0.1 / 1.0)))) != 3.14159) || v3)
        _ = fi27
    }
    switch (v26 << ((7) & 63)) {
    case 42:
        var v28 float32 = 0.1
        _ = v28
    case 127:
        v29 := GenS1{F0: 274, F1: (v1 > -1), F2: -32768}
        _ = v29
    default:
        v30 := GenS1{F0: 18446744073709551615, F1: (42 > 18446744073709551615), F2: 58}
        _ = v30
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(true || (32768 >= v26))) {
        var v31 int = (^(v26))
        for fi32 := 0; fi32 < 17; fi32++ {
            switch (("foo bar" + ("a" + "\n\t")) + "foo bar") {
            case "foo bar":
                v33 := GenS1{F0: 124, F1: v3, F2: 739}
                _ = v33
            }
            _ = fi32
        }
        _ = v31
    } else {
        switch v26 {
        case 128:
            for fi34 := 0; fi34 < 1; fi34++ {
                if ((42 <= 2) && true) {
                    var v35 float64 = -2.25
                    var v36 uint8 = 110
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            if v3 {
                                switch "foo bar" {
                                case "a":
                                    var v38 int32 = -261
                                    _ = v38
                                case "":
                                    var v39 uint64 = 1
                                    _ = v39
                                case "foo bar":
                                    v2 = append(v2, 982)
                                default:
                                    fmt.Println(v14)
                                }
                                v40 := GenS3{F0: 94, F1: v4}
                                var v41 int64 = 0
                                v42 := GenS1{F0: uint64(v4), F1: v3, F2: 42}
                                _ = v40
                                _ = v41
                                _ = v42
                            } else {
                                var v43 int8 = (v1 - (v1 >> ((1) & 63)))
                                if (!false) {
                                    var v44 uint32 = 1
                                    var v45 uint64 = 504
                                    var v46 uint64 = 4294967295
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                switch v26 {
                                case 127:
                                    if false {
                                        var v47 uint8 = 127
                                        fmt.Println(v43, v3, v14)
                                        v48 := GenS1{F0: uint64(v47), F1: v3, F2: -32768}
                                        _ = v47
                                        _ = v48
                                    } else {
                                        fmt.Println(v3, v3, v1)
                                        var v49 string = ""
                                        _ = v49
                                    }
                                default:
                                    if (42 > -745) {
                                        var v50 int8 = 115
                                        if v3 {
                                            var v51 int32 = (-(v14))
                                            if true {
                                                var v52 bool = (v4 < v4)
                                                var v53 int8 = int8(v51)
                                                _ = len(v2)
                                                _ = v52
                                                _ = v53
                                            } else {
                                                switch fi34 {
                                                case 32767:
                                                    if (!true) {
                                                        fmt.Println(v26, v20, v3)
                                                    }
                                                case 4294967295:
                                                    v54 := GenS4{F0: "hello"}
                                                    _ = v54
                                                }
                                                v55 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar"))}
                                                var v56 uint16 = v4
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v55
                                                _ = v56
                                            }
                                            v51 = (^(v51))
                                            _ = v51
                                        }
                                        _ = v50
                                    } else {
                                        v14 -= 0
                                        var v57 int32 = -1
                                        v14 += (-(v20))
                                        _ = v57
                                    }
                                }
                                switch fi34 {
                                case -301:
                                    v20 = (v14 ^ (v14 / ((int32(v4)) | 1)))
                                default:
                                    v20 -= (-(v20))
                                }
                                _ = len(v2)
                                _ = v43
                            }
                            v26 = -9223372036854775808
                            _ = fi37
                        }
                        v58 := GenS1{F0: 633, F1: v3, F2: -32768}
                        v59 := GenS4{F0: "hello"}
                        _ = v58
                        _ = v59
                    } else {
                        fmt.Println(v14, v1, v1)
                    }
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    if v3 {
                        var v61 string = ("foo bar" + "a")
                        switch -90 {
                        case 2147483647:
                            v62 := GenS5{F0: v1}
                            _ = v62
                        case 32768:
                            if (255 != 255) {
                                if (!v3) {
                                    v63 := GenS2{F0: "hello"}
                                    var v64 int32 = (^(v20))
                                    v3 = v3
                                    _ = v63
                                    _ = v64
                                } else {
                                    v65 := GenS4{F0: (("\n\t" + (("" + ("\n\t" + v61)) + v61)) + (v61 + v61))}
                                    _ = v65
                                }
                                if len(v2) > 0 {
                                    v2[0] = 807
                                }
                                var v66 float32 = (float32(((float64((0.1 - (1.5 - 0.1))) * 3.14159) + 1.5)) - float32(((0.0 - (((0.0 / 1.0) * (((1e10 / 1.0) * 0.0) - 1e10)) - 1.5)) - (0.1 - (-2.25 - 0.0)))))
                                _ = v66
                            }
                        }
                        _ = v61
                    }
                    v3 = (((1.5 / ((0.0 - float64(1.5)) + float64(3.14159))) <= (-2.25 - (((0.1 * 0.1) / 1.0) + 1e10))) && v3)
                    _ = fi60
                }
                _ = fi34
            }
        }
        var v67 int64 = 264
        _ = v67
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        _ = len(v2)
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}

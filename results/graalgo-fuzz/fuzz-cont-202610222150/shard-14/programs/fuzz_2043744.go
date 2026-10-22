package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint8; F2 int8 }
type GenS6 struct { F0 int16; F1 string; F2 bool }
type GenS7 struct { F0 int16 }
func genFn0() int32 {
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            v5 -= 2
        }
        _ = fi6
    }
    if (-587 == 1000) {
        v7 := GenS1{F0: 2}
        var v8 int64 = 42
        _ = v7
        _ = v8
    } else {
        switch 42 {
        case 1:
            var v9 int16 = 0
            _ = v9
        case 9223372036854775807:
            var v10 int8 = 100
            _ = v10
        case -1:
            v5 = v5
        }
    }
    if true {
        if (65535 > 2147483647) {
            switch (("foo bar" + ("" + "foo bar")) + (("a" + "hello") + ("" + "foo bar"))) {
            case "":
                fmt.Println(v5)
            default:
                var v11 int = -977
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int64 = 1
                fmt.Println(v5, v13)
                var v14 uint = 0
                _ = fi12
                _ = v13
                _ = v14
            }
            if (!false) {
                v15 := []int{1}
                var v16 float32 = (1.5 * -2.25)
                v16 -= 0.0
                _ = v15
                _ = v16
            }
        }
    } else {
        v17 := []int{1}
        if len(v17) > 0 {
            v17[0] = 253
        }
        _ = v17
    }
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    v3 := []string{}
    v4 := make(map[int]int)
    switch -97 {
    case 32767:
        v18 := GenS2{F0: 255, F1: 703, F2: -856}
        _ = v18
    case -906:
        switch "" {
        case "":
            v19 := GenS3{F0: "foo bar"}
            _ = v19
        }
    case -9223372036854775808:
        if ((("a" + "hello") != "a") && (487 > 2)) {
            for fi20 := 0; fi20 < 9; fi20++ {
                if ((1e10 / 1.5) < (-2.25 - 0.1)) {
                    fmt.Println(fi20)
                    var v21 int64 = 42
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 string = ("a" + "")
                        v21 = v21
                        if false {
                            var v24 uint8 = 241
                            v23 = (v23 + v23)
                            var v25 uint64 = 838
                            _ = v24
                            _ = v25
                        } else {
                            var v26 uint = 42
                            if len(v3) > 0 {
                                v3[0] = v23
                            }
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                } else {
                    v27 := GenS1{F0: 2}
                    _ = len(v2)
                    v28 := GenS2{F0: 255, F1: 4294967295, F2: 42}
                    _ = v27
                    _ = v28
                }
                var v29 uint64 = 195
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 int64 = -1
                    var v32 float64 = float64(0.1)
                    v32 = v32
                    v33 := GenS3{F0: "hello"}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi20
                _ = v29
            }
            var v34 int = 32768
            if false {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            }
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint = 501
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 uint = v36
            var v39 uint16 = 145
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 int32 = 2147483647
        _ = fi40
        _ = v41
    }
    var v42 float64 = 0.1
    var v43 int64 = 9223372036854775807
    fmt.Println(v42, v42)
    fmt.Println(v43, v43)
    var v44 int = -9223372036854775808
    v2[(v44 % ((v44) | 1))] = (3.14159 > (((-2.25 - 1.5) + (0.1 / float32((1e10 + 0.0)))) * (1.5 * 0.0)))
    var v45 int32 = 881
    for fi46 := 0; fi46 < 11; fi46++ {
        switch "" {
        case "a":
            v2[-4] = true
        case "hello":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
        var v47 bool = false
        v47 = true
        _ = fi46
        _ = v47
    }
    var v48 int16 = 0
    for fi49 := 0; fi49 < 2; fi49++ {
        if false {
            if (((42 <= 2) && (true || true)) || ((v45 & v45) < (^(v45)))) {
                v43 = v43
                if false {
                    for fi50 := 0; fi50 < 9; fi50++ {
                        switch "hello" {
                        case "a":
                            var v51 int32 = 42
                            _ = v51
                        default:
                            if (255 < 1) {
                                var v52 int64 = v43
                                var v53 string = ("a" + "")
                                _ = v52
                                _ = v53
                            } else {
                                fmt.Println(v45, v42)
                            }
                        }
                        if (0.0 != -2.25) {
                            v54 := GenS2{F0: 127, F1: 4294967295, F2: 0}
                            v42 = v42
                            _ = v54
                        }
                        _ = fi50
                    }
                    var v55 int16 = (v48 / ((int16(v43)) | 1))
                    switch "hello" {
                    case "hello":
                        if true {
                            if (v42 >= v42) {
                                var v56 uint64 = 42
                                _ = v56
                            } else {
                                if (568 > 18446744073709551615) {
                                    for fi57 := 0; fi57 < 11; fi57++ {
                                        fmt.Println(v48, v43, v55)
                                        v44 = 128
                                        v58 := v1[(-(fi57))]
                                        _ = fi57
                                        _ = v58
                                    }
                                    for fi59 := 0; fi59 < 17; fi59++ {
                                        v60 := GenS2{F0: 128, F1: 18446744073709551615, F2: 42}
                                        _ = fi59
                                        _ = v60
                                    }
                                    var v61 string = ("hello" + (("a" + ("\n\t" + "\n\t")) + (("" + "hello") + "a")))
                                    _ = v61
                                }
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    switch (^(v44)) {
                                    case -840:
                                        if ("hello" == ("hello" + "\n\t")) {
                                            v48 = (v55 >> ((0) & 63))
                                        } else {
                                            v44 = 32768
                                        }
                                    case 42:
                                        var v63 int16 = 1000
                                        _ = v63
                                    case 629:
                                        var v64 float64 = v42
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        v45 = 1
                                        _ = v64
                                    default:
                                        v65 := GenS4{F0: 169, F1: 1}
                                        _ = v65
                                    }
                                    _ = fi62
                                }
                            }
                            v48 += v48
                        }
                    }
                    _ = v55
                }
            } else {
                if true {
                    if true {
                        fmt.Println(v48, v48, v44)
                        if false {
                            var v66 bool = (!true)
                            _ = v66
                        }
                    } else {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v42, v44, v42)
                        case "hello":
                            fmt.Println(v43, v45, v48)
                        default:
                            var v67 uint64 = 189
                            _ = v67
                        }
                        for fi68 := 0; fi68 < 6; fi68++ {
                            switch "hello" {
                            case "":
                                var v69 uint = 1
                                _ = v69
                            case "hello":
                                v70 := GenS2{F0: 18, F1: 0, F2: 9223372036854775807}
                                _ = v70
                            }
                            if (2 <= 42) {
                                var v71 float64 = (v42 + (v42 - float64((0.0 * (3.14159 * 0.1)))))
                                _ = v71
                            } else {
                                fmt.Println(v48, v44)
                                v3 = append(v3, (("a" + "a") + "foo bar"))
                            }
                            _ = fi68
                        }
                    }
                } else {
                    fmt.Println(v45)
                }
                v72 := GenS5{F0: v45, F1: 46, F2: 42}
                _ = v72
            }
            v44 += v44
        }
        _ = fi49
    }
    if false {
        if (68 < 255) {
            switch v44 {
            case -2147483648:
                if false {
                    v73 := GenS6{F0: v48, F1: ("\n\t" + ("" + "hello")), F2: false}
                    for fi74 := 0; fi74 < 11; fi74++ {
                        switch "hello" {
                        case "foo bar":
                            var v75 uint = 0
                            _ = v75
                        case "":
                            fmt.Println(fi74, v45, v43)
                        case "hello":
                            v45 += 923
                        default:
                            fmt.Println(v45, v48, v48)
                        }
                        var v76 float32 = float32(v42)
                        _ = fi74
                        _ = v76
                    }
                    _ = v73
                }
            default:
                for fi77 := 0; fi77 < 2; fi77++ {
                    v45 = -1
                    _ = fi77
                }
            }
        } else {
            fmt.Println(v45, v42, v45)
            v78 := GenS7{F0: (v48 ^ (v48 >> ((785) & 63)))}
            _ = v78
        }
        v79 := GenS6{F0: 42, F1: "a", F2: true}
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v48
}

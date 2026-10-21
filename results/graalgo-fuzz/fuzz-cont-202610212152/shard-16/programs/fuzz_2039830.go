package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 bool; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint; F2 string }
type GenS8 struct { F0 string; F1 int16; F2 int64 }
type GenS9 struct { F0 int16; F1 uint64; F2 int8 }
type GenS10 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    var v7 float32 = 1.5
    p0 += 285
    _ = p0
    _ = v7
    return 128
}
func genFn1() uint16 {
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{(fi8 + fi8), fi8}
        v10 := GenS1{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 127 {
    case 128:
        v11 := GenS1{F0: 0}
        _ = v11
    case 32768:
        v12 := []int{310, -1}
        _ = v12
    }
    var v13 int32 = 2147483647
    _ = v13
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t", "\n\t"}
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    var v5 int8 = 127
    v6 := make(map[string]uint32)
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS1{F0: -433}
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: -9223372036854775808}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int64 = -9223372036854775808
        fmt.Println(fi18)
        _ = fi18
        _ = v19
    }
    if (1 > 29) {
        if (53 == 2) {
            v20 := GenS2{F0: 96, F1: ((("hello" + "\n\t") + "hello") + "\n\t")}
            _ = v20
        } else {
            if false {
                v3 -= v3
                if false {
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "hello")
                    }
                    v5 -= 0
                    if true {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi21
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    } else {
                        var v22 int8 = v5
                        if true {
                            if true {
                                var v23 string = "foo bar"
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    for fi25 := 0; fi25 < 3; fi25++ {
                                        switch "a" {
                                        case "a":
                                            switch fi24 {
                                            case 16:
                                                switch ((v23 + ((v23 + "") + ("" + ("hello" + (v23 + (v23 + (v23 + (v23 + "\n\t")))))))) + "a") {
                                                case "foo bar":
                                                    var v26 float32 = -2.25
                                                    _ = v26
                                                case "":
                                                    if true {
                                                        var v27 bool = (!false)
                                                        var v28 uint64 = 947
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        var v29 int16 = 32767
                                                        v30 := GenS1{F0: 0}
                                                        switch (v3 + fi24) {
                                                        case 128:
                                                            var v31 uint64 = 0
                                                            _ = v31
                                                        case -475:
                                                            v32 := GenS3{F0: 1, F1: 2}
                                                            _ = v32
                                                        case 42:
                                                            v3 += fi24
                                                        }
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    v33 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615}
                                                    _ = v33
                                                default:
                                                    v22 += v5
                                                }
                                            }
                                        default:
                                            switch ((("hello" + v23) + ("foo bar" + v23)) + v23) {
                                            case "\n\t":
                                                var v34 string = v23
                                                _ = v34
                                            case "foo bar":
                                                v5 += 1
                                            case "hello":
                                                if false {
                                                    var v35 int64 = -9223372036854775808
                                                    _ = v35
                                                }
                                                v22 = (v22 & v22)
                                                fmt.Println(v22)
                                            }
                                        }
                                        if (!(!(-1 > 42))) {
                                            var v36 string = ("foo bar" + "\n\t")
                                            _ = v36
                                        } else {
                                            fmt.Println(fi25, fi24, fi25)
                                            fmt.Println(v5, v3, fi25)
                                            v37 := GenS4{F0: -594}
                                            _ = v37
                                        }
                                        _ = fi25
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    _ = fi24
                                }
                                _ = v23
                            }
                        } else {
                            for fi38 := 0; fi38 < 16; fi38++ {
                                v39 := GenS5{F0: "\n\t", F1: false, F2: "foo bar"}
                                if ("hello" >= "") {
                                    v5 += 55
                                    var v40 uint = 4294967295
                                    _ = v40
                                } else {
                                    for fi41 := 0; fi41 < 19; fi41++ {
                                        v6[("a" + "foo bar")] = 646
                                        _ = fi41
                                    }
                                    v5 += (v22 & (v22 - v5))
                                }
                                _ = fi38
                                _ = v39
                            }
                            v22 -= (v5 & 100)
                        }
                        _ = v22
                    }
                } else {
                    var v42 float64 = -2.25
                    _ = v42
                }
                if false {
                    var v43 uint8 = 2
                    v43 -= (v43 / ((42) | 1))
                    _ = v43
                } else {
                    delete(v4, 170)
                }
            }
            v3 = v3
        }
        fmt.Println(v5, v3)
        v44 := GenS6{F0: 1}
        _ = v44
    } else {
        v45 := GenS7{F0: 332, F1: 2, F2: "foo bar"}
        fmt.Println(v3, v5)
        var v46 float64 = 1e10
        _ = v45
        _ = v46
    }
    v47 := GenS3{F0: 2, F1: 0}
    v48 := GenS3{F0: 0, F1: 42}
    for fi49 := 0; fi49 < 15; fi49++ {
        if (!true) {
            switch "" {
            case "foo bar":
                v1[v3] = ((832 >= 473) || (((0.1 - 0.1) - float64((0.0 / 1.0))) >= (-2.25 - 3.14159)))
            case "\n\t":
                v50 := GenS5{F0: "", F1: ((2 == 65535) || (599 <= 491)), F2: "a"}
                _ = v50
            default:
                v51 := GenS8{F0: "", F1: 520, F2: -137}
                _ = v51
            }
            switch (^(fi49)) {
            case 32768:
                v3 -= v3
            case 9:
                for fi52 := 0; fi52 < 17; fi52++ {
                    if true {
                        v1[(fi49 >> ((18446744073709551615) & 63))] = true
                        v53 := GenS9{F0: -1, F1: 1, F2: (v5 | (v5 & v5))}
                        v3 += v3
                        _ = v53
                    } else {
                        v5 += v5
                        v54 := GenS4{F0: 237}
                        _ = v54
                    }
                    _ = fi52
                }
            case 9223372036854775807:
                v55 := GenS2{F0: genFn0(18446744073709551615), F1: ""}
                _ = v55
            }
            if false {
                v3 = (v3 & (-(v3)))
                v56 := GenS4{F0: 0}
                v4[65535] = 42
                _ = v56
            }
        }
        _ = fi49
    }
    var v57 uint = 868
    v58 := v6[""]
    if (false && ((v5 | 127) == v5)) {
        v59 := GenS8{F0: ((("foo bar" + "a") + "a") + ""), F1: 1, F2: 9223372036854775807}
        v60 := GenS10{F0: 333, F1: "a", F2: "foo bar"}
        v5 = int8(v57)
        _ = v59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}

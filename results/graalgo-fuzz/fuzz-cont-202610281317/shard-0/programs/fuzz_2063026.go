package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32; F2 string }
type GenS7 struct { F0 string; F1 int64; F2 int64 }
func genFn0() int {
    var v6 int8 = 42
    v7 := GenS1{F0: 795, F1: "foo bar"}
    for fi8 := 0; fi8 < 19; fi8++ {
        if (true && true) {
            for fi9 := 0; fi9 < 5; fi9++ {
                v6 = (-(v6))
                v10 := GenS1{F0: 315, F1: "hello"}
                _ = fi9
                _ = v10
            }
        } else {
            var v11 float32 = (0.1 - -2.25)
            var v12 int8 = (v6 << ((65535) & 63))
            fmt.Println(v11, v12, fi8)
            _ = v11
            _ = v12
        }
        fmt.Println(v6, v6)
        _ = fi8
    }
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int32) uint64 {
    if ("a" == "\n\t") {
        switch genFn0() {
        case 42:
            v13 := []int{genFn0(), genFn0(), 32767}
            _ = v13
        case 32768:
            for fi14 := 0; fi14 < 16; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := []int{(fi15 << ((0) & 63)), fi14, (fi15 & (fi15 / ((fi15) | 1)))}
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS2{F0: (fi17 - genFn0())}
                    _ = fi17
                    _ = v18
                }
                var v19 uint32 = 2147483647
                _ = fi14
                _ = v19
            }
        case 153:
            v20 := []int{0}
            _ = v20
        }
    }
    p0 = p0
    var v21 int = 4294967295
    v22 := GenS1{F0: 4294967295, F1: "hello"}
    _ = p0
    _ = v21
    _ = v22
    return 405
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := []uint32{65535, (v1 >> ((255) & 63))}
    v5 := make(map[string]string)
    fmt.Println(v1)
    switch "" {
    case "foo bar":
        v1 = 65535
    case "\n\t":
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int32 = -2147483648
            for fi25 := 0; fi25 < 5; fi25++ {
                if (!false) {
                    _ = len(v5)
                    v4 = append(v4, v1)
                    if (2 != 0) {
                        switch -2147483648 {
                        case 2147483647:
                            if false {
                                var v26 bool = false
                                var v27 bool = (false && (605 != 4294967295))
                                switch (("a" + "a") + "a") {
                                case "\n\t":
                                    fmt.Println(v1, v27, v1)
                                case "hello":
                                    v3["foo bar"] = 2147483647
                                case "foo bar":
                                    var v28 uint = 4294967295
                                    _ = v28
                                default:
                                    var v29 uint16 = 32767
                                    v30 := GenS2{F0: (fi23 | fi23)}
                                    _ = len(v4)
                                    _ = v29
                                    _ = v30
                                }
                                _ = v26
                                _ = v27
                            }
                        case 4294967295:
                            var v31 int = -9223372036854775808
                            _ = v31
                        case 1:
                            if true {
                                v32 := GenS3{F0: true}
                                _ = v32
                            }
                        default:
                            switch "\n\t" {
                            case "":
                                v33 := GenS1{F0: 22, F1: (("a" + ("" + "foo bar")) + "")}
                                _ = v33
                            case "foo bar":
                                var v34 uint32 = 703
                                _ = v34
                            case "a":
                                v24 = (v24 * (v24 * 0))
                            }
                        }
                        v24 = v24
                        if (42 >= v1) {
                            var v35 string = ""
                            _ = v35
                        } else {
                            switch "\n\t" {
                            case "a":
                                switch genFn0() {
                                case 0:
                                    v1 = v1
                                case -9223372036854775808:
                                    switch (-(fi25)) {
                                    case 2147483647:
                                        if false {
                                            v24 -= v24
                                            switch (fi25 ^ (-(fi25))) {
                                            case 4294967295:
                                                var v36 int = fi25
                                                _ = v36
                                            }
                                        }
                                    case -76:
                                        v37 := GenS1{F0: 0, F1: (("\n\t" + "foo bar") + "\n\t")}
                                        _ = v37
                                    case -225:
                                        v24 = 0
                                    default:
                                        if false {
                                            v1 = v1
                                            var v38 int = (fi25 >> ((42) & 63))
                                            v1 -= (v1 - v1)
                                            _ = v38
                                        }
                                    }
                                }
                            case "hello":
                                if ((("foo bar" + "hello") + "\n\t") <= "\n\t") {
                                    v24 += (-(v24))
                                }
                            default:
                                var v39 uint = 0
                                _ = v39
                            }
                        }
                    } else {
                        var v40 uint64 = 1
                        _ = v40
                    }
                }
                var v41 int32 = (v24 + -2147483648)
                v41 = (v41 + (v41 | v24))
                var v42 int32 = 1
                var v43 float64 = (float64((0.1 - 3.14159)) * ((0.0 / 1e10) + 1e10))
                _ = fi25
                _ = v41
                _ = v42
                _ = v43
            }
            var v44 bool = false
            _ = fi23
            _ = v24
            _ = v44
        }
    }
    var v45 int = 42
    if false {
        if true {
            switch (v45 & v45) {
            case 9223372036854775807:
                v45 = v45
            case 32768:
                var v46 uint32 = 451
                _ = v46
            }
        } else {
            v47 := GenS3{F0: (((0.0 * 1e10) * 0.1) == 0.0)}
            _ = v47
        }
        for fi48 := 0; fi48 < 12; fi48++ {
            fmt.Println(v45)
            for fi49 := 0; fi49 < 13; fi49++ {
                switch (-(fi49)) {
                case -910:
                    fmt.Println(fi49, fi48)
                case 866:
                    v45 = 127
                }
                _ = fi49
            }
            for fi50 := 0; fi50 < 16; fi50++ {
                v51 := GenS4{F0: 809, F1: 124}
                _ = fi50
                _ = v51
            }
            _ = fi48
        }
    }
    if true {
        for fi52 := 0; fi52 < 2; fi52++ {
            var v53 uint = 811
            v54 := GenS5{F0: (("" + ("a" + "a")) + "a"), F1: 255}
            _ = fi52
            _ = v53
            _ = v54
        }
    } else {
        var v55 int = -718
        _ = v55
    }
    switch "" {
    case "\n\t":
        v4 = append(v4, (v1 | v1))
    case "a":
        v56 := GenS6{F0: false, F1: v1, F2: ("" + "")}
        _ = v56
    case "foo bar":
        var v57 bool = false
        _ = v57
    default:
        v58 := v5[""]
        _ = v58
    }
    v1 += v1
    v45 = 32767
    delete(v5, "\n\t")
    v59 := GenS2{F0: v45}
    if false {
        v45 = (v45 - -2147483648)
    } else {
        v60 := GenS7{F0: "", F1: -1, F2: 121}
        v4 = append(v4, 1)
        _ = v60
    }
    for fi61 := 0; fi61 < 17; fi61++ {
        v1 += v1
        _ = fi61
    }
    v1 -= 2147483647
    if len(v4) > 0 {
        v4[0] = (v1 % (((^(v1))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v59
}

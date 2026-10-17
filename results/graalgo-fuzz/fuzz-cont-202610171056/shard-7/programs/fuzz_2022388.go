package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0(p0 uint16, p1 uint16, p2 uint16) int8 {
    v7 := []int{}
    var v8 uint8 = uint8(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 100
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 int = -824
    var v5 int16 = 27
    var v6 int32 = 2147483647
    fmt.Println(v5)
    if (!false) {
        fmt.Println(v6)
        fmt.Println(v4)
    }
    var v9 int = 433
    if (1 > 127) {
        v6 += (v6 >> ((963) & 63))
        var v10 float64 = (1e10 - ((1.5 - ((((-2.25 / ((1.5 / (3.14159 * 1.5)) - 1e10)) + 0.1) / (0.1 / -2.25)) * 0.0)) / (1e10 - ((3.14159 / 1.5) * (float64(1e10) / 1e10)))))
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            v5 -= v5
            _ = fi11
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v6 = (v6 << ((255) & 63))
        var v13 uint8 = 255
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v9, v4)
    v4 = v4
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := v2["hello"]
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS1{F0: 0, F1: -1}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    if (false || true) {
        var v18 int32 = v6
        _ = len(v2)
        _ = v18
    } else {
        switch 9223372036854775807 {
        case 162:
            if false {
                switch "foo bar" {
                case "":
                    var v19 string = "foo bar"
                    _ = v19
                case "hello":
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int32 = (v6 % ((v6) | 1))
                        _ = fi20
                        _ = v21
                    }
                case "\n\t":
                    v5 += v5
                }
                if (v5 < v5) {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        if (42 <= (v5 << ((255) & 63))) {
                            if true {
                                v9 += (v4 ^ v9)
                                v23 := GenS2{F0: 100}
                                v24 := GenS3{F0: (v4 & fi22), F1: 255, F2: "hello"}
                                _ = v23
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    switch (v9 - 9223372036854775807) {
                                    case 1:
                                        var v26 float32 = (((float32((-2.25 / 1.0)) * 1e10) - 1.5) + (3.14159 * (1.5 - 3.14159)))
                                        _ = v26
                                    case 2147483647:
                                        if false {
                                            v27 := []int{v4, fi22, v9}
                                            var v28 uint = 0
                                            var v29 uint8 = 2
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                    fmt.Println(fi25)
                                    _ = fi25
                                }
                            }
                        } else {
                            var v30 uint = 4294967295
                            _ = v30
                        }
                        var v31 uint64 = 4294967295
                        _ = fi22
                        _ = v31
                    }
                    if ((false || (!true)) || (678 == 275)) {
                        var v32 int32 = v6
                        fmt.Println(v5)
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v4 += fi34
                                _ = fi34
                            }
                            for fi35 := 0; fi35 < 2; fi35++ {
                                v36 := GenS1{F0: 2, F1: (^(v5))}
                                v4 = v4
                                delete(v3, ("foo bar" + (((("" + ("hello" + (("foo bar" + ("" + "hello")) + ("a" + "\n\t")))) + ("\n\t" + ("" + (((("a" + "hello") + ("" + "\n\t")) + "foo bar") + "\n\t")))) + "\n\t") + ("\n\t" + "\n\t"))))
                                fmt.Println(v9)
                                _ = fi35
                                _ = v36
                            }
                            _ = fi33
                        }
                        for fi37 := 0; fi37 < 12; fi37++ {
                            for fi38 := 0; fi38 < 2; fi38++ {
                                var v39 int8 = -128
                                _ = fi38
                                _ = v39
                            }
                            _ = fi37
                        }
                        v9 -= v9
                    }
                }
                v40 := []int{(v9 >> ((4294967295) & 63)), v4, (v9 >> ((516) & 63)), v4}
                _ = v40
            }
        case -1:
            var v41 string = ""
            _ = v41
        case 535:
            switch "foo bar" {
            case "a":
                fmt.Println(v9)
            case "\n\t":
                for fi42 := 0; fi42 < 19; fi42++ {
                    for fi43 := 0; fi43 < 12; fi43++ {
                        v9 = v4
                        var v44 bool = (!(true && ((((v6 + v6) > (^(v6))) && (false && false)) || true)))
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
            case "foo bar":
                v4 -= (v9 | -1)
            }
        }
        delete(v3, "a")
        delete(v2, "foo bar")
    }
    var v45 float32 = 0.1
    var v46 int64 = 9223372036854775807
    for fi47 := 0; fi47 < 9; fi47++ {
        if true {
            for fi48 := 0; fi48 < 4; fi48++ {
                var v49 uint = 87
                v50 := GenS4{F0: (-(v5)), F1: (v6 * (v6 + (v6 | (-(v6)))))}
                _ = fi48
                _ = v49
                _ = v50
            }
            if false {
                var v51 float32 = -2.25
                var v52 bool = true
                _ = v51
                _ = v52
            } else {
                v53 := v3["\n\t"]
                delete(v1, 253)
                _ = v53
            }
            v46 += v46
        } else {
            v46 = (v46 + (v46 >> ((1) & 63)))
            v54 := []int{}
            _ = v54
        }
        for fi55 := 0; fi55 < 2; fi55++ {
            var v56 int16 = (^(v5))
            var v57 int32 = v6
            if (int16(v6) != v5) {
                v58 := GenS4{F0: (^(v5)), F1: v6}
                v59 := GenS2{F0: 0}
                _ = v58
                _ = v59
            }
            _ = fi55
            _ = v56
            _ = v57
        }
        _ = fi47
    }
    var v60 int64 = v46
    if (65535 == 4294967295) {
        if false {
            v61 := []int{-898, v9}
            v62 := GenS1{F0: 255, F1: 1000}
            _ = v61
            _ = v62
        } else {
            if (!false) {
                v46 += (v46 | v60)
            }
            if false {
                _ = len(v3)
                for fi63 := 0; fi63 < 4; fi63++ {
                    switch -1 {
                    case 127:
                        v64 := v2["hello"]
                        _ = v64
                    }
                    var v65 int64 = (v60 * (v46 >> ((42) & 63)))
                    _ = fi63
                    _ = v65
                }
                var v66 uint64 = 42
                _ = v66
            } else {
                switch "a" {
                case "a":
                    var v67 uint8 = 235
                    _ = v67
                default:
                    switch (v9 * (v9 * (v4 + int(v60)))) {
                    case 609:
                        var v68 uint64 = 2
                        _ = v68
                    default:
                        fmt.Println(v45)
                    }
                }
            }
            var v69 uint32 = 1
            _ = v69
        }
        v70 := GenS3{F0: v4, F1: 42, F2: ""}
        var v71 int = 127
        _ = v70
        _ = v71
    }
    v72 := []int{v9, v9}
    var v73 int8 = 42
    if len(v72) > 0 {
        v72[0] = (v4 | -56)
    }
    v74 := v2["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v60
    _ = v72
    _ = v73
    _ = v74
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 uint64; F1 uint64; F2 int }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint32, p2 int) uint32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: (p2 & 0), F1: 32767}
        v8 := GenS1{F0: fi6, F1: uint16(p0)}
        if false {
            v9 := []int{p2}
            fmt.Println(fi6, p1)
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v10 := GenS1{F0: p2, F1: 1}
    v11 := []int{p2, p2, 32768, (-(p2))}
    var v12 uint16 = 42
    if len(v11) > 0 {
        v11[0] = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (p1 >> ((2) & 63))
}

func main() {
    var v1 int32 = -1
    var v2 int = 1
    v3 := make(map[string]string)
    var v4 float64 = 0.1
    v5 := []uint32{}
    var v13 float64 = v4
    if false {
        var v14 float64 = (v13 + v4)
        for fi15 := 0; fi15 < 10; fi15++ {
            if false {
                fmt.Println(v13, v14)
                _ = len(v5)
                var v16 bool = (2 < 4294967295)
                _ = v16
            } else {
                fmt.Println(v4, fi15, v4)
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 uint8 = 16
                    _ = fi17
                    _ = v18
                }
            }
            v1 += 0
            v13 += ((v14 - 1.5) - (-2.25 / 0.1))
            _ = fi15
        }
        v4 = -2.25
        _ = v14
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS1{F0: (fi20 >> ((455) & 63)), F1: 679}
            v13 -= float64(((0.1 + 1e10) + float32(1e10)))
            var v22 float32 = ((((0.1 + (0.1 * -2.25)) / ((1e10 / float32(((v4 / float64(0.1)) - (1.5 + v4)))) / 1.0)) * (1e10 / ((((0.1 - 0.1) + float32(1e10)) + 1.5) - 1.5))) + 3.14159)
            _ = fi20
            _ = v21
            _ = v22
        }
        if false {
            v1 = 2147483647
            v4 += v4
        }
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        _ = fi19
    }
    var v23 bool = true
    v24 := GenS2{F0: 255, F1: "a", F2: ((("a" + "hello") + ("hello" + ("a" + ""))) + "\n\t")}
    if (v2 != (v2 >> ((255) & 63))) {
        v1 = -2147483648
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            if v23 {
                v3[("" + ("" + ("hello" + "")))] = "a"
                v26 := GenS2{F0: 0, F1: "foo bar", F2: "\n\t"}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 7; fi27++ {
                    fmt.Println(v1)
                    _ = fi27
                }
            }
            _ = fi25
        }
    }
    fmt.Println(v23)
    var v28 int = (v2 + v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v29 float32 = 0.1
    var v30 uint = 255
    fmt.Println(v1, v2)
    var v31 uint32 = 739
    if true {
        if v23 {
            switch "" {
            case "a":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                v23 = true
            case "":
                var v32 uint64 = 2
                _ = v32
            default:
                switch -772 {
                case 4294967295:
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if v23 {
                            switch "hello" {
                            case "\n\t":
                                var v34 int64 = -9223372036854775808
                                _ = v34
                            default:
                                v35 := GenS3{F0: (^(v31)), F1: (-(v1)), F2: ("" + ((("\n\t" + "\n\t") + "a") + "a"))}
                                _ = v35
                            }
                            v36 := GenS4{F0: 4294967295, F1: 2, F2: (v2 - v28)}
                            if len(v5) > 0 {
                                v5[0] = v31
                            }
                            _ = v36
                        }
                        var v37 uint = 42
                        v4 = v4
                        _ = fi33
                        _ = v37
                    }
                case -505:
                    v38 := v3[("hello" + ("a" + "hello"))]
                    _ = v38
                case 32767:
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        var v39 uint64 = 858
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 int64 = 0
                        _ = fi40
                        _ = v41
                    }
                }
            }
            if (false || (v1 == v1)) {
                v13 -= (v13 - float64((v29 - (3.14159 - v29))))
                var v42 int8 = 1
                for fi43 := 0; fi43 < 14; fi43++ {
                    v42 = v42
                    if false {
                        v29 += 1.5
                        v44 := GenS4{F0: 296, F1: 0, F2: v28}
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 2; fi45++ {
                            switch 0 {
                            case 42:
                                if (v1 >= (v1 ^ v1)) {
                                    var v46 float32 = 1.5
                                    v3["a"] = "\n\t"
                                    _ = v46
                                } else {
                                    var v47 int = int(v31)
                                    _ = v47
                                }
                            case 474:
                                v48 := v3[(("foo bar" + "") + ("hello" + "foo bar"))]
                                _ = v48
                            }
                            v5 = append(v5, 2)
                            if ((int(v1) != v28) && v23) {
                                if (255 < uint8(v30)) {
                                    if (((v42 * 0) != v42) && v23) {
                                        var v49 uint8 = 42
                                        v31 = (v31 + 1)
                                        switch ("\n\t" + "") {
                                        case "":
                                            var v50 int = fi45
                                            _ = v50
                                        case "a":
                                            if v23 {
                                                v51 := GenS2{F0: 42, F1: ("\n\t" + "a"), F2: "hello"}
                                                _ = v51
                                            }
                                        default:
                                            v49 += (v49 ^ (v49 - (^(v49))))
                                        }
                                        _ = v49
                                    } else {
                                        var v52 uint = v30
                                        if (12 == 2) {
                                            var v53 int32 = (v1 >> ((v52) & 63))
                                            _ = v53
                                        }
                                        _ = v52
                                    }
                                    var v54 int32 = v1
                                    _ = v54
                                } else {
                                    var v55 int16 = -1
                                    switch ((("\n\t" + ("a" + "a")) + ("foo bar" + "a")) + ("foo bar" + "a")) {
                                    case "":
                                        switch (fi45 | 1) {
                                        case -1:
                                            v30 -= (v30 ^ 255)
                                        case 718:
                                            if v23 {
                                                v56 := GenS4{F0: 18446744073709551615, F1: 885, F2: fi45}
                                                _ = v56
                                            } else {
                                                v28 += (fi43 ^ fi43)
                                                var v57 int = fi43
                                                _ = v57
                                            }
                                        }
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v55
                                }
                                v58 := GenS5{F0: 42}
                                v59 := v3["\n\t"]
                                _ = v58
                                _ = v59
                            }
                            if v23 {
                                v60 := v3["\n\t"]
                                _ = v60
                            } else {
                                var v61 int32 = (v1 - v1)
                                switch fi45 {
                                case 32767:
                                    var v62 string = ("hello" + (("\n\t" + "\n\t") + "hello"))
                                    _ = v62
                                }
                                v29 -= v29
                                _ = v61
                            }
                            _ = fi45
                        }
                    }
                    for fi63 := 0; fi63 < 6; fi63++ {
                        v3[("hello" + ("\n\t" + "foo bar"))] = "foo bar"
                        v2 += fi63
                        _ = fi63
                    }
                    _ = fi43
                }
                _ = v42
            }
            fmt.Println(v1)
        } else {
            switch "foo bar" {
            case "":
                var v64 int16 = 32767
                _ = v64
            }
            for fi65 := 0; fi65 < 16; fi65++ {
                var v66 int = 32767
                var v67 int = (v66 * (v28 * (fi65 & (v28 % ((v28) | 1)))))
                _ = fi65
                _ = v66
                _ = v67
            }
        }
        fmt.Println(v4)
        var v68 int16 = 0
        _ = v68
    }
    var v69 uint32 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v69
}

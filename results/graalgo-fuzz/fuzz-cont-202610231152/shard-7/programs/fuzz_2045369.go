package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 int16; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 uint = 497
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := []uint32{42}
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    var v7 int64 = 719
    v1 = (v1 | v1)
    v8 := GenS1{F0: (v1 ^ 548), F1: ("\n\t" + "a"), F2: 4294967295}
    if true {
        var v9 int32 = -1
        fmt.Println(v7, v7)
        if true {
            for fi10 := 0; fi10 < 19; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := GenS1{F0: 255, F1: "a", F2: (fi10 ^ (-(fi10)))}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            v1 = (^(v1))
            _ = fi13
        }
    }
    var v14 int32 = -1
    if false {
        var v15 uint64 = 4294967295
        v1 += 0
        if (false && (true && (127 != 69))) {
            if ((!((v14 ^ (v14 * v14)) < v14)) && false) {
                for fi16 := 0; fi16 < 4; fi16++ {
                    fmt.Println(v15, v14, v1)
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if (false || (899 != -42)) {
                            v14 += 513
                            v1 -= v1
                        } else {
                            v18 := GenS1{F0: (^(v1)), F1: (("foo bar" + ("\n\t" + ("\n\t" + "a"))) + "\n\t"), F2: fi16}
                            _ = v18
                        }
                        if ((255 > 851) || false) {
                            if (1 == -803) {
                                var v19 float32 = (((0.1 + 1.5) - 3.14159) * 1.5)
                                fmt.Println(fi16, fi16)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if true {
                                    if false {
                                        v1 += (v1 * 2)
                                    } else {
                                        if (42 == 2147483648) {
                                            var v20 uint64 = v15
                                            _ = v20
                                        }
                                        var v21 uint32 = 4294967295
                                        _ = v21
                                    }
                                    v22 := GenS2{F0: 767, F1: 65535, F2: 0}
                                    if false {
                                        v23 := GenS1{F0: 44, F1: "a", F2: (fi17 | (fi16 ^ 1))}
                                        _ = v23
                                    } else {
                                        v24 := GenS2{F0: 4294967295, F1: 805, F2: 255}
                                        var v25 int16 = 0
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = v22
                                }
                                if ((0.1 > (((v19 * (float32((-2.25 * 3.14159)) / v19)) / v19) + v19)) && ((!false) || (!false))) {
                                    if (157 >= 2) {
                                        var v26 bool = false
                                        _ = v26
                                    } else {
                                        switch fi16 {
                                        case 0:
                                            v15 = (v15 << ((4294967295) & 63))
                                        case -15:
                                            if true {
                                                fmt.Println(v14, fi16)
                                                v27 := GenS3{F0: fi17, F1: 42, F2: ""}
                                                _ = v27
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    v28 := GenS1{F0: (v1 >> (((^(v1))) & 63)), F1: "", F2: 128}
                                                    _ = v28
                                                case "":
                                                    v29 := GenS3{F0: fi17, F1: 42, F2: ("hello" + "hello")}
                                                    _ = v29
                                                default:
                                                    if true {
                                                        if (0 >= 112) {
                                                            v30 := GenS1{F0: (v1 % ((2) | 1)), F1: ("hello" + ""), F2: (-(fi17))}
                                                            _ = v30
                                                        } else {
                                                            var v31 uint = 0
                                                            if (!((1e10 == 0.0) || (0 != 2147483647))) {
                                                                var v32 float64 = (0.1 - (float64(3.14159) / -2.25))
                                                                v15 = v15
                                                                var v33 uint8 = 42
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                            _ = v31
                                                        }
                                                        if (!(true && false)) {
                                                            if true {
                                                                var v34 int16 = 1000
                                                                _ = v34
                                                            }
                                                            v35 := GenS4{F0: 1000, F1: -32768, F2: (!(false || false))}
                                                            v36 := GenS3{F0: (^(fi17)), F1: 1, F2: "a"}
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        var v37 bool = (42 <= 65535)
                                                        v4 = append(v4, 1)
                                                        if v37 {
                                                            var v38 float32 = (v19 * 3.14159)
                                                            _ = v38
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            delete(v6, 720)
                                                        }
                                                        _ = v37
                                                    }
                                                }
                                            }
                                        case 813:
                                            var v39 uint32 = 1
                                            _ = v39
                                        }
                                    }
                                }
                                if true {
                                    if ((v1 - (^(v1))) < (v1 & (v1 * v1))) {
                                        var v40 int32 = v14
                                        if (1.5 != 0.1) {
                                            var v41 bool = true
                                            v15 += 2
                                            var v42 float64 = 0.1
                                            _ = v41
                                            _ = v42
                                        } else {
                                            fmt.Println(v7, v14)
                                            if true {
                                                fmt.Println(fi16)
                                            } else {
                                                var v43 string = ("a" + "foo bar")
                                                if (!false) {
                                                    v19 -= -2.25
                                                    v44 := GenS3{F0: (^(fi16)), F1: -32768, F2: v43}
                                                    _ = v44
                                                } else {
                                                    var v45 uint32 = 2
                                                    if false {
                                                        var v46 int32 = (^(v40))
                                                        var v47 int16 = 0
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        fmt.Println(v43, v14)
                                                    }
                                                    _ = v45
                                                }
                                                _ = len(v4)
                                                _ = v43
                                            }
                                        }
                                        fmt.Println(v15, v40)
                                        _ = v40
                                    } else {
                                        v48 := GenS2{F0: 4294967295, F1: 42, F2: 0}
                                        v14 = (v14 - v14)
                                        _ = v48
                                    }
                                    var v49 string = "\n\t"
                                    var v50 bool = (2 >= 255)
                                    _ = v49
                                    _ = v50
                                }
                                switch (-(fi17)) {
                                case -1:
                                    fmt.Println(v19, v7, fi16)
                                case 1:
                                    v15 = v15
                                case 380:
                                    var v51 float64 = (1e10 + -2.25)
                                    _ = v51
                                default:
                                    var v52 uint8 = 255
                                    _ = v52
                                }
                                _ = v19
                            } else {
                                v1 -= v1
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if (("a" + "foo bar") <= "") {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v53 := GenS3{F0: 820, F1: 1000, F2: "hello"}
                                    _ = v53
                                } else {
                                    v54 := GenS5{F0: v15}
                                    var v55 float32 = 0.1
                                    v14 = (v14 | (v14 | 2147483647))
                                    v56 := GenS6{F0: v7}
                                    v15 -= 18446744073709551615
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            case "hello":
                                if false {
                                    var v57 uint16 = 255
                                    v15 -= (^(v15))
                                    var v58 uint64 = v15
                                    _ = v57
                                    _ = v58
                                } else {
                                    v15 = (v15 | (^(v15)))
                                    v59 := GenS2{F0: 4294967295, F1: 1, F2: 128}
                                    var v60 float64 = (-2.25 + 3.14159)
                                    _ = v59
                                    _ = v60
                                }
                                fmt.Println(v1)
                                _ = len(v3)
                            case "":
                                switch fi17 {
                                case -85:
                                    v7 += v7
                                }
                            default:
                                var v61 uint32 = 2147483648
                                _ = v61
                            }
                        } else {
                            v62 := GenS6{F0: (v7 >> ((2) & 63))}
                            _ = v62
                        }
                        _ = fi17
                    }
                    _ = len(v5)
                    _ = fi16
                }
                var v63 uint = v1
                for fi64 := 0; fi64 < 18; fi64++ {
                    v65 := GenS5{F0: 28}
                    if false {
                        _ = len(v4)
                        var v66 float64 = 0.1
                        _ = v66
                    }
                    _ = fi64
                    _ = v65
                }
                _ = v63
            }
            var v67 float32 = 3.14159
            fmt.Println(v15, v67)
            _ = v67
        }
        _ = v15
    }
    v14 = (v14 % ((-863) | 1))
    var v68 int32 = v14
    var v69 int = 2147483647
    v7 += v7
    for fi70 := 0; fi70 < 4; fi70++ {
        var v71 float64 = 0.1
        _ = fi70
        _ = v71
    }
    var v72 int = (v69 | (-(v69)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v7
    _ = v72
    _ = v8
}

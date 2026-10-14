package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 uint; F1 int8; F2 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := GenS1{F0: "", F1: "", F2: (42 < -128)}
    v5 := GenS2{F0: 2147483647, F1: "\n\t"}
    var v6 int = 798
    v7 := GenS3{F0: 593}
    v8 := GenS4{F0: -2147483648}
    v6 = (v6 << ((1) & 63))
    v9 := []int{v6}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float64 = 1.5
        _ = fi10
        _ = v11
    }
    if false {
        switch v6 {
        case 365:
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: "a", F1: "", F2: (true && true)}
                _ = fi12
                _ = v13
            }
        case 127:
            v6 = v6
        default:
            if true {
                switch v6 {
                case 42:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v6 += 702
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 int8 = 42
                            _ = fi15
                            _ = v16
                        }
                        _ = fi14
                    }
                default:
                    var v17 uint8 = 1
                    _ = v17
                }
                if (!true) {
                    if false {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            var v19 string = "hello"
                            var v20 int32 = -2147483648
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        v21 := GenS5{F0: (v6 & v6)}
                        v22 := v3[-1]
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS6{F0: 1, F1: 966}
                        _ = v23
                    }
                    if true {
                        var v24 string = ""
                        if false {
                            if (!false) {
                                for fi25 := 0; fi25 < 6; fi25++ {
                                    if (!false) {
                                        _ = len(v9)
                                        if true {
                                            v6 -= (fi25 ^ (fi25 / ((fi25) | 1)))
                                        } else {
                                            var v26 uint32 = 1
                                            _ = v26
                                        }
                                    }
                                    if true {
                                        for fi27 := 0; fi27 < 6; fi27++ {
                                            var v28 float64 = ((0.0 + -2.25) - 1.5)
                                            if true {
                                                v6 += fi25
                                            } else {
                                                var v29 int = v6
                                                v30 := GenS7{F0: 2147483648, F1: 65535}
                                                v29 -= (v29 % ((fi25) | 1))
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = fi27
                                            _ = v28
                                        }
                                    } else {
                                        var v31 string = ((((("" + v24) + "\n\t") + v24) + "foo bar") + v24)
                                        _ = v31
                                    }
                                    _ = fi25
                                }
                                v6 = 128
                            } else {
                                if (!(255 < 32767)) {
                                    switch v24 {
                                    case "\n\t":
                                        _ = len(v9)
                                    case "hello":
                                        v6 = (v6 >> ((0) & 63))
                                    }
                                }
                            }
                            _ = len(v9)
                        } else {
                            v1[v6] = false
                        }
                        v32 := GenS8{F0: (2 == 2), F1: 42}
                        _ = v24
                        _ = v32
                    }
                    v6 += (v6 / (((v6 * -635)) | 1))
                } else {
                    var v33 uint16 = 65535
                    v6 = (v6 % (((-(v6))) | 1))
                    if len(v9) > 0 {
                        v9[0] = (v6 >> ((0) & 63))
                    }
                    _ = v33
                }
            } else {
                var v34 float64 = (0.0 + ((0.0 * (1.5 / ((0.1 * 1.5) - ((0.0 / float64(1e10)) * 1e10)))) * ((3.14159 / -2.25) * float64((3.14159 + (0.0 - (1e10 - (0.1 - 0.0))))))))
                _ = v34
            }
        }
        switch 355 {
        case 2147483647:
            if false {
                v6 = v6
                if (!("hello" == "foo bar")) {
                    switch "a" {
                    case "":
                        for fi35 := 0; fi35 < 6; fi35++ {
                            v6 = (v6 % (((fi35 << ((1) & 63))) | 1))
                            _ = fi35
                        }
                    case "hello":
                        fmt.Println(v6)
                    case "a":
                        switch v6 {
                        case 9223372036854775807:
                            var v36 string = "a"
                            _ = v36
                        case 127:
                            fmt.Println(v6)
                        case 432:
                            if false {
                                v6 = 630
                                switch 32768 {
                                case 128:
                                    v6 -= v6
                                }
                                var v37 uint = 1
                                _ = v37
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = v6
                                }
                                if true {
                                    if true {
                                        v6 = (v6 * v6)
                                        var v38 bool = false
                                        _ = v38
                                    }
                                    v39 := v2[("" + "\n\t")]
                                    if false {
                                        fmt.Println(v6, v6)
                                        if false {
                                            var v40 float32 = ((-2.25 * 1.5) / 1.0)
                                            _ = v40
                                        } else {
                                            for fi41 := 0; fi41 < 7; fi41++ {
                                                switch "hello" {
                                                case "\n\t":
                                                    v42 := GenS1{F0: "hello", F1: "a", F2: false}
                                                    _ = v42
                                                case "a":
                                                    for fi43 := 0; fi43 < 10; fi43++ {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v44 int = fi41
                                                            v45 := GenS9{F0: 4294967295, F1: -1, F2: 728}
                                                            _ = v44
                                                            _ = v45
                                                        default:
                                                            v46 := GenS6{F0: -58, F1: 1}
                                                            _ = v46
                                                        }
                                                        var v47 int8 = 100
                                                        v48 := GenS5{F0: (fi43 % ((v6) | 1))}
                                                        _ = fi43
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                case "foo bar":
                                                    v3[fi41] = (1.5 != -2.25)
                                                default:
                                                    var v49 uint8 = 255
                                                    _ = v49
                                                }
                                                var v50 int32 = 2147483647
                                                _ = fi41
                                                _ = v50
                                            }
                                            v6 += v6
                                        }
                                    }
                                    _ = v39
                                } else {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    v6 = (v6 - (v6 ^ v6))
                                }
                            }
                        default:
                            if true {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                        }
                    default:
                        if true {
                            fmt.Println(v6)
                            var v51 bool = false
                            var v52 uint32 = 1
                            _ = v51
                            _ = v52
                        } else {
                            v53 := GenS6{F0: 360, F1: 4294967295}
                            _ = v53
                        }
                    }
                }
                delete(v1, -9223372036854775808)
            }
        }
        for fi54 := 0; fi54 < 10; fi54++ {
            var v55 float32 = 1e10
            var v56 uint64 = 461
            fmt.Println(v6, v6)
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    if (true && ("foo bar" == (("" + "foo bar") + "foo bar"))) {
        var v57 uint64 = 4294967295
        _ = v57
    } else {
        for fi58 := 0; fi58 < 1; fi58++ {
            v6 = v6
            v9 = append(v9, (-(v6)))
            _ = fi58
        }
        v6 = (v6 / (((-(v6))) | 1))
    }
    v59 := GenS5{F0: v6}
    var v60 float32 = (((-2.25 / 1.5) + 0.0) / 1.5)
    v61 := v1[v6]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v7
    _ = v8
    _ = v9
}

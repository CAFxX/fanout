package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 int }

func main() {
    var v1 int8 = 67
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]uint32)
    v4 := []uint32{2, 0}
    var v5 int64 = v2
    v6 := GenS1{F0: 18446744073709551615, F1: 40, F2: v5}
    if (true && false) {
        if false {
            if false {
                switch "a" {
                case "\n\t":
                    switch ("a" + "hello") {
                    case "a":
                        _ = len(v4)
                    case "":
                        var v7 string = "\n\t"
                        _ = v7
                    }
                case "hello":
                    v5 = v2
                case "":
                    switch 128 {
                    case -246:
                        switch 0 {
                        case 4294967295:
                            v4 = append(v4, 577)
                        case -2147483648:
                            if true {
                                if true {
                                    if (!(613 == 42)) {
                                        v8 := GenS1{F0: 892, F1: 2, F2: 1}
                                        _ = v8
                                    }
                                    var v9 int8 = -1
                                    for fi10 := 0; fi10 < 10; fi10++ {
                                        v1 = 42
                                        _ = fi10
                                    }
                                    _ = v9
                                } else {
                                    if ((true || false) || false) {
                                        v11 := GenS2{F0: 18, F1: 0, F2: 802}
                                        for fi12 := 0; fi12 < 7; fi12++ {
                                            if true {
                                                v4 = append(v4, 445)
                                            }
                                            _ = fi12
                                        }
                                        _ = v11
                                    } else {
                                        v5 -= v2
                                        v13 := GenS1{F0: 2, F1: 556, F2: v2}
                                        v5 -= 9223372036854775807
                                        _ = v13
                                    }
                                    var v14 string = "hello"
                                    var v15 int64 = v2
                                    _ = v14
                                    _ = v15
                                }
                                if (255 > 276) {
                                    v3["hello"] = 97
                                    var v16 uint = 1
                                    _ = v16
                                } else {
                                    v17 := GenS1{F0: 457, F1: 32767, F2: (-(v5))}
                                    fmt.Println(v2)
                                    _ = v17
                                }
                                var v18 bool = false
                                _ = v18
                            }
                        case 9223372036854775807:
                            for fi19 := 0; fi19 < 3; fi19++ {
                                v20 := GenS3{F0: (-(v5)), F1: 1, F2: 42}
                                v21 := v3[(("\n\t" + "\n\t") + "\n\t")]
                                var v22 bool = (v1 != v1)
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        default:
                            if (1 >= 0) {
                                for fi23 := 0; fi23 < 10; fi23++ {
                                    v24 := GenS2{F0: 18446744073709551615, F1: 128, F2: 1}
                                    switch (^(fi23)) {
                                    case 9223372036854775807:
                                        var v25 float64 = 1.5
                                        _ = v25
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v3[("\n\t" + "a")] = 331
                                        }
                                    }
                                    delete(v3, ("foo bar" + ""))
                                    _ = fi23
                                    _ = v24
                                }
                                if false {
                                    fmt.Println(v5, v1, v5)
                                } else {
                                    v26 := GenS4{F0: 128}
                                    _ = v26
                                }
                            } else {
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    v28 := GenS5{F0: false}
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    var v29 uint16 = 2
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                v1 += (v1 * (v1 + v1))
                                if false {
                                    v30 := GenS6{F0: "foo bar"}
                                    if false {
                                        if true {
                                            var v31 float32 = 3.14159
                                            _ = v31
                                        } else {
                                            if (0 == (-(v1))) {
                                                for fi32 := 0; fi32 < 1; fi32++ {
                                                    for fi33 := 0; fi33 < 16; fi33++ {
                                                        switch "hello" {
                                                        case "a":
                                                            var v34 uint64 = 582
                                                            _ = v34
                                                        }
                                                        _ = len(v3)
                                                        v1 = v1
                                                        fmt.Println(v1, v2, v2)
                                                        _ = fi33
                                                    }
                                                    fmt.Println(v2, v1)
                                                    v35 := GenS7{F0: 0}
                                                    _ = fi32
                                                    _ = v35
                                                }
                                            }
                                            var v36 uint32 = 42
                                            _ = v36
                                        }
                                    }
                                    if (true || false) {
                                        switch "foo bar" {
                                        case "":
                                            var v37 bool = (2 > 625)
                                            _ = v37
                                        case "a":
                                            for fi38 := 0; fi38 < 19; fi38++ {
                                                v5 += v2
                                                v1 += -1
                                                if ((fi38 | (fi38 + fi38)) == (^(fi38))) {
                                                    switch (fi38 - 597) {
                                                    case -791:
                                                        v1 += (v1 / (((v1 & v1)) | 1))
                                                    case 32767:
                                                        fmt.Println(v1)
                                                    default:
                                                        delete(v3, "foo bar")
                                                    }
                                                    if true {
                                                        if false {
                                                            v5 = int64(v1)
                                                            v39 := GenS2{F0: 0, F1: 2, F2: -2147483648}
                                                            v40 := GenS6{F0: ("hello" + "")}
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            v1 = (v1 % (((v1 >> ((669) & 63))) | 1))
                                                            v41 := GenS8{F0: 128}
                                                            _ = v41
                                                        }
                                                    } else {
                                                        fmt.Println(v1, fi38)
                                                    }
                                                } else {
                                                    var v42 uint64 = 4294967295
                                                    if true {
                                                        v43 := GenS9{F0: 331, F1: fi38}
                                                        _ = v43
                                                    }
                                                    _ = v42
                                                }
                                                _ = fi38
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                        v44 := v3["a"]
                                        v2 += 0
                                        _ = v44
                                    } else {
                                        v5 += (-(v2))
                                        v45 := GenS8{F0: 0}
                                        _ = v45
                                    }
                                    _ = v30
                                }
                            }
                        }
                    case 128:
                        var v46 uint32 = 2147483648
                        _ = v46
                    case -634:
                        var v47 float32 = 1e10
                        _ = v47
                    default:
                        v2 = v2
                    }
                default:
                    var v48 uint16 = 226
                    _ = v48
                }
                var v49 int8 = 127
                _ = len(v3)
                _ = v49
            } else {
                if true {
                    var v50 bool = false
                    fmt.Println(v2, v1)
                    v51 := GenS9{F0: 255, F1: 32768}
                    _ = v50
                    _ = v51
                }
                var v52 string = ("" + "hello")
                v2 -= (v2 * (v5 << ((2) & 63)))
                _ = v52
            }
            v1 = v1
            var v53 uint16 = 255
            _ = v53
        } else {
            if (1 < v1) {
                v1 -= (v1 % (((-(v1))) | 1))
            } else {
                for fi54 := 0; fi54 < 8; fi54++ {
                    v1 = v1
                    for fi55 := 0; fi55 < 6; fi55++ {
                        if (!(32767 == 280)) {
                            v2 = (v2 | v5)
                            var v56 uint = 42
                            _ = v56
                        }
                        switch fi55 {
                        case 128:
                            var v57 int32 = -9
                            v58 := GenS8{F0: 128}
                            _ = v57
                            _ = v58
                        }
                        var v59 string = ("\n\t" + "a")
                        _ = len(v3)
                        v60 := GenS9{F0: 67, F1: -1}
                        v2 -= (v2 - (v2 | (v5 | v2)))
                        _ = fi55
                        _ = v59
                        _ = v60
                    }
                    fmt.Println(v5, v2)
                    _ = fi54
                }
            }
            var v61 int16 = -1
            v2 += 146
            _ = v61
        }
        v2 = (v5 * (v5 - (^(v2))))
        for fi62 := 0; fi62 < 12; fi62++ {
            v5 -= 42
            var v63 float32 = 1.5
            v64 := GenS5{F0: true}
            _ = fi62
            _ = v63
            _ = v64
        }
    } else {
        fmt.Println(v2, v1, v2)
    }
    if false {
        switch (((("hello" + ("\n\t" + ((("foo bar" + "\n\t") + ("hello" + ("\n\t" + ("a" + "foo bar")))) + "a"))) + "\n\t") + ("" + (((("a" + "") + "a") + "\n\t") + ""))) + ("foo bar" + "\n\t")) {
        case "foo bar":
            _ = len(v4)
        default:
            var v65 float64 = 1.5
            _ = v65
        }
    }
    if false {
        for fi66 := 0; fi66 < 17; fi66++ {
            v1 -= v1
            v1 = v1
            _ = fi66
        }
        var v67 string = "\n\t"
        _ = v67
    }
    var v68 int32 = -2147483648
    if false {
        fmt.Println(v68)
        var v69 string = "foo bar"
        _ = v69
    } else {
        var v70 int64 = v2
        var v71 int32 = v68
        _ = v70
        _ = v71
    }
    fmt.Println(v2, v5)
    var v72 uint32 = 625
    v72 -= v72
    var v73 int16 = -359
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v72
    _ = v73
}

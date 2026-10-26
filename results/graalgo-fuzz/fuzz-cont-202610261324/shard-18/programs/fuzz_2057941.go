package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 string; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint }
func genFn0(p0 uint16, p1 uint16, p2 uint8) int {
    fmt.Println(p2, p0, p0)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 2}
        _ = fi6
        _ = v7
    }
    switch -861 {
    case 128:
        var v8 uint = 65535
        _ = v8
    case -692:
        var v9 int64 = 603
        _ = v9
    case 727:
        fmt.Println(p1, p2)
    }
    v10 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint32) int8 {
    v11 := []int{-527, genFn0(42, 2, 42), 135}
    if (false && false) {
        var v12 uint = 738
        var v13 uint32 = 2
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 59
}

func main() {
    var v1 uint64 = 42
    var v2 uint8 = 128
    var v3 int = 9223372036854775807
    var v4 uint32 = 764
    v5 := make(map[string]int)
    if ((("" + ("foo bar" + "")) + ("\n\t" + "foo bar")) == "hello") {
        switch v3 {
        case 44:
            v4 += 2147483648
        default:
            v14 := GenS1{F0: 979}
            _ = v14
        }
        v15 := GenS3{F0: v4}
        _ = v15
    }
    if false {
        switch ("\n\t" + (("foo bar" + "hello") + "\n\t")) {
        case "hello":
            v16 := []int{genFn0(60, 609, v2)}
            _ = v16
        case "\n\t":
            if ((632 != 42) || false) {
                var v17 uint = 0
                for fi18 := 0; fi18 < 8; fi18++ {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v2 -= v2
                        v4 = v4
                        _ = fi19
                    }
                    var v20 uint16 = 42
                    _ = fi18
                    _ = v20
                }
                v21 := GenS4{F0: 127}
                _ = v17
                _ = v21
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    if (-2.25 != (1.5 / float64(0.1))) {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            if false {
                                switch (fi23 % ((v3) | 1)) {
                                case 42:
                                    if true {
                                        var v24 uint16 = 255
                                        _ = v24
                                    } else {
                                        var v25 int32 = 0
                                        _ = v25
                                    }
                                case 9223372036854775807:
                                    if true {
                                        if (fi22 < 2147483647) {
                                            var v26 uint32 = 4294967295
                                            v2 = (v2 - (^(v2)))
                                            _ = v26
                                        } else {
                                            v1 -= v1
                                        }
                                        switch 2147483647 {
                                        case -83:
                                            v27 := GenS3{F0: (^(v4))}
                                            v28 := v5[""]
                                            _ = v27
                                            _ = v28
                                        case 9223372036854775807:
                                            v29 := GenS2{F0: (v3 ^ fi23)}
                                            _ = v29
                                        }
                                        v2 -= 127
                                    }
                                }
                                switch fi23 {
                                case 0:
                                    fmt.Println(v2)
                                case -565:
                                    var v30 int16 = 42
                                    var v31 uint = 18446744073709551615
                                    _ = v30
                                    _ = v31
                                case -123:
                                    v3 -= fi22
                                }
                                v32 := GenS4{F0: 100}
                                if (fi22 < int(v1)) {
                                    if (true || true) {
                                        var v33 uint8 = (v2 << ((0) & 63))
                                        var v34 uint8 = v33
                                        _ = v33
                                        _ = v34
                                    }
                                    fmt.Println(fi23)
                                    var v35 int64 = 0
                                    _ = v35
                                }
                                if (true && (!(v1 != v1))) {
                                    switch v3 {
                                    case 4294967295:
                                        fmt.Println(fi22, v4)
                                    }
                                    v3 += -158
                                } else {
                                    var v36 uint = 1
                                    fmt.Println(v2, v36)
                                    fmt.Println(fi22, v2, v2)
                                    _ = v36
                                }
                                v2 -= v2
                                _ = v32
                            }
                            var v37 float32 = -2.25
                            _ = fi23
                            _ = v37
                        }
                    } else {
                        v4 = (v4 ^ v4)
                    }
                    _ = fi22
                }
                if false {
                    v2 = 253
                } else {
                    v38 := GenS5{F0: (!(-128 <= 43)), F1: ("foo bar" + ("" + "a")), F2: 65535}
                    _ = v38
                }
            }
        case "":
            if (true && true) {
                var v39 string = "hello"
                switch ("a" + (v39 + v39)) {
                case "foo bar":
                    v4 = 42
                case "a":
                    switch "foo bar" {
                    case "hello":
                        v40 := GenS1{F0: 296}
                        _ = v40
                    }
                }
                _ = v39
            } else {
                v1 -= (v1 % ((uint64(v4)) | 1))
            }
        }
    } else {
        if true {
            switch "a" {
            case "foo bar":
                v1 += uint64(v2)
            case "\n\t":
                fmt.Println(v4)
            case "a":
                for fi41 := 0; fi41 < 17; fi41++ {
                    var v42 int32 = -135
                    var v43 float32 = 0.1
                    var v44 int32 = v42
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
            }
            switch -144 {
            case 128:
                for fi45 := 0; fi45 < 8; fi45++ {
                    switch 2147483647 {
                    case 2147483647:
                        v46 := []int{(v3 ^ -437)}
                        _ = v46
                    case 1:
                        fmt.Println(v4, v3)
                    }
                    fmt.Println(fi45, v3, v4)
                    v47 := GenS1{F0: uint16(v2)}
                    _ = fi45
                    _ = v47
                }
            case 127:
                if (((((-(v3)) > 128) && (1e10 <= float32(3.14159))) || (32767 > 0)) && true) {
                    for fi48 := 0; fi48 < 2; fi48++ {
                        var v49 bool = false
                        _ = fi48
                        _ = v49
                    }
                    v1 += v1
                    var v50 uint16 = 2
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 16; fi51++ {
                        v1 = (v1 * (v1 + 1))
                        var v52 int = (-(v3))
                        v53 := GenS5{F0: (-32768 >= 1), F1: "hello", F2: 42}
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 int8 = genFn1(v4)
                    switch ("\n\t" + "a") {
                    case "hello":
                        v55 := []int{v3}
                        _ = v55
                    default:
                        v4 += v4
                    }
                    _ = v54
                }
            case 4294967295:
                fmt.Println(v2, v1)
            }
        } else {
            var v56 bool = (-388 > -266)
            var v57 int8 = -25
            if v56 {
                v2 = (^(v2))
                v2 += (v2 | (v2 ^ 42))
                if v56 {
                    fmt.Println(v57, v56)
                    switch v3 {
                    case 128:
                        var v58 float64 = (((3.14159 / ((3.14159 * 1.5) + (0.0 / (3.14159 * ((-2.25 * 3.14159) - (1e10 + 0.0)))))) * 0.1) * 1.5)
                        _ = v58
                    case -362:
                        var v59 int = genFn0(716, 255, 0)
                        _ = v59
                    case 554:
                        for fi60 := 0; fi60 < 17; fi60++ {
                            v61 := GenS3{F0: (^(v4))}
                            for fi62 := 0; fi62 < 8; fi62++ {
                                v63 := GenS5{F0: v56, F1: ("hello" + ""), F2: 65535}
                                _ = fi62
                                _ = v63
                            }
                            var v64 string = "\n\t"
                            _ = fi60
                            _ = v61
                            _ = v64
                        }
                    }
                } else {
                    v56 = v56
                }
            }
            _ = v56
            _ = v57
        }
        switch genFn0(1, 567, v2) {
        case 32768:
            var v65 uint32 = (v4 << ((255) & 63))
            _ = v65
        case 4294967295:
            for fi66 := 0; fi66 < 10; fi66++ {
                v67 := GenS4{F0: 0}
                for fi68 := 0; fi68 < 2; fi68++ {
                    v69 := GenS6{F0: v2}
                    v70 := []int{(v3 ^ genFn0(277, 65535, 42)), fi68, (v3 ^ v3)}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                _ = fi66
                _ = v67
            }
        }
    }
    fmt.Println(v1, v1)
    var v71 uint32 = v4
    fmt.Println(v71, v1, v1)
    for fi72 := 0; fi72 < 19; fi72++ {
        if false {
            fmt.Println(fi72)
            v2 += v2
            delete(v5, "a")
        } else {
            fmt.Println(v2, v4, fi72)
            v2 = (^(v2))
        }
        _ = fi72
    }
    var v73 int8 = 1
    v71 = 42
    for fi74 := 0; fi74 < 16; fi74++ {
        v71 -= (v71 ^ 4294967295)
        v73 = genFn1(v4)
        _ = fi74
    }
    switch ("hello" + ("\n\t" + ("a" + "\n\t"))) {
    case "foo bar":
        v75 := GenS5{F0: ((-2.25 < 3.14159) && true), F1: "", F2: 18446744073709551615}
        _ = v75
    case "\n\t":
        var v76 uint64 = v1
        _ = v76
    }
    if true {
        for fi77 := 0; fi77 < 10; fi77++ {
            if (false || (2 >= 12)) {
                for fi78 := 0; fi78 < 13; fi78++ {
                    var v79 float32 = 1.5
                    _ = fi78
                    _ = v79
                }
                v73 -= v73
                v3 -= v3
            } else {
                v80 := GenS7{F0: 0}
                _ = v80
            }
            v81 := GenS8{F0: -1, F1: 255, F2: 0}
            _ = fi77
            _ = v81
        }
        v5[""] = v3
        var v82 bool = (32767 > 798)
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v73
}

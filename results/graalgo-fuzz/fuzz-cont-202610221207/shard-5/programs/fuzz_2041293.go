package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 int16; F2 bool }

func main() {
    var v1 uint32 = 582
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v4 -= v4
            var v7 uint8 = v4
            switch "\n\t" {
            case "hello":
                v8 := GenS1{F0: false, F1: 65535, F2: 2}
                _ = v8
            case "\n\t":
                var v9 int = fi5
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: false, F1: 0, F2: 1}
                _ = v10
            default:
                if false {
                    var v11 float64 = 1.5
                    _ = v11
                }
                var v12 int32 = -2147483648
                _ = len(v3)
                v4 += (v4 * v4)
                var v13 int32 = (v12 + (-(v12)))
                _ = v12
                _ = v13
            }
            _ = fi6
            _ = v7
        }
        var v14 string = "hello"
        _ = fi5
        _ = v14
    }
    if false {
        fmt.Println(v4)
        v15 := GenS2{F0: ((603 < -484) || (73 < 100)), F1: 65535, F2: ""}
        v16 := v2[((("" + ("a" + "\n\t")) + "a") + "\n\t")]
        _ = v15
        _ = v16
    } else {
        if (!false) {
            for fi17 := 0; fi17 < 11; fi17++ {
                v4 = (^(v4))
                fmt.Println(v1, v4)
                if (42 == 1) {
                    v18 := []int{(fi17 + (fi17 / ((fi17) | 1))), (fi17 << ((4294967295) & 63))}
                    v1 = (v1 + (v1 - v1))
                    _ = v18
                }
                _ = fi17
            }
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := []int{}
                    var v22 int8 = 42
                    if (18446744073709551615 != 18446744073709551615) {
                        v23 := GenS3{F0: 387}
                        var v24 uint16 = 2
                        _ = v23
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi25 := 0; fi25 < 12; fi25++ {
                    v4 -= (v4 * 36)
                    if (v1 >= v1) {
                        var v26 string = "hello"
                        _ = v26
                    } else {
                        if (1.5 >= 3.14159) {
                            fmt.Println(v1, fi25)
                            if false {
                                v27 := v3[(fi19 + fi25)]
                                v28 := GenS4{F0: 4294967295, F1: true, F2: (fi25 * fi19)}
                                var v29 int = 31
                                if false {
                                    var v30 bool = (v27 && (!v27))
                                    v31 := GenS3{F0: 0}
                                    _ = v30
                                    _ = v31
                                }
                                v32 := GenS2{F0: false, F1: 111, F2: ""}
                                v33 := GenS5{F0: 42}
                                var v34 int = (fi25 % (((fi19 >> ((58) & 63))) | 1))
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            var v35 uint32 = (v1 & v1)
                            if false {
                                _ = len(v3)
                            }
                            var v36 int8 = 100
                            var v37 float64 = 1e10
                            var v38 uint = 1
                            var v39 int32 = 2147483647
                            v37 -= 3.14159
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = fi25
                }
                _ = fi19
            }
            v40 := GenS6{F0: false}
            _ = v40
        }
        v41 := GenS1{F0: (!(42 <= -69)), F1: 845, F2: 267}
        _ = v41
    }
    v1 = v1
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v42 int8 = 1
        _ = v42
    case "foo bar":
        v43 := GenS1{F0: (42 != 0), F1: 926, F2: 65}
        _ = v43
    case "hello":
        fmt.Println(v1, v1)
    }
    if ((!(42 <= 1000)) || false) {
        v44 := v3[42]
        switch ("" + "") {
        case "foo bar":
            v44 = false
        case "":
            fmt.Println(v44, v4, v44)
        }
        if v44 {
            if v44 {
                var v45 bool = v44
                v46 := GenS3{F0: 238}
                v47 := GenS3{F0: 32767}
                _ = v45
                _ = v46
                _ = v47
            }
        } else {
            var v48 uint16 = 42
            fmt.Println(v48, v4)
            var v49 float32 = 3.14159
            _ = v48
            _ = v49
        }
        _ = v44
    } else {
        v4 -= v4
        v50 := GenS6{F0: false}
        fmt.Println(v4, v1)
        _ = v50
    }
    v1 = 100
    if (3.14159 >= float32(float64(-2.25))) {
        v51 := v2[("" + "foo bar")]
        if (0 <= 0) {
            v52 := GenS1{F0: (-9223372036854775808 > -9223372036854775808), F1: 569, F2: 18446744073709551615}
            v53 := GenS7{F0: 128, F1: -1, F2: v51}
            for fi54 := 0; fi54 < 19; fi54++ {
                var v55 uint64 = 1
                _ = fi54
                _ = v55
            }
            _ = v52
            _ = v53
        } else {
            if v51 {
                var v56 int8 = -1
                _ = v56
            } else {
                v57 := []int{-481, 42}
                _ = v57
            }
            var v58 int64 = -1
            v59 := v2[""]
            _ = v58
            _ = v59
        }
        _ = v51
    } else {
        var v60 uint32 = v1
        _ = v60
    }
    v61 := GenS2{F0: true, F1: 255, F2: (((("foo bar" + ("" + "hello")) + ((("foo bar" + "\n\t") + "a") + "foo bar")) + "a") + "a")}
    for fi62 := 0; fi62 < 9; fi62++ {
        var v63 int8 = 100
        fmt.Println(fi62, v63, v63)
        for fi64 := 0; fi64 < 12; fi64++ {
            delete(v3, (fi64 ^ fi62))
            v3[(fi64 / ((42) | 1))] = true
            v65 := []int{128, -861, int(v4), (fi62 << ((0) & 63))}
            v1 = (v1 << ((255) & 63))
            _ = fi64
            _ = v65
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
}

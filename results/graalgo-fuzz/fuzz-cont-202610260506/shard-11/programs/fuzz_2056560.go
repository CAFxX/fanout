package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
type GenS4 struct { F0 int64; F1 int16; F2 int64 }
type GenS5 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 int16, p1 uint) int {
    var v6 uint8 = 184
    var v7 int16 = p0
    var v8 string = "foo bar"
    switch ("foo bar" + (("" + "") + v8)) {
    case "\n\t":
        v8 = ((v8 + ("hello" + v8)) + v8)
    case "foo bar":
        v7 -= (v7 - (p0 ^ v7))
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                switch fi10 {
                case -882:
                    v11 := GenS1{F0: (v6 | v6), F1: p0, F2: false}
                    p1 += p1
                    _ = v11
                }
                var v12 int = 9223372036854775807
                v13 := []int{}
                _ = fi10
                _ = v12
                _ = v13
            }
            _ = fi9
        }
    }
    v14 := []int{128}
    _ = p0
    _ = p1
    _ = v14
    _ = v6
    _ = v7
    _ = v8
    return int(p1)
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    var v3 int32 = 1
    var v4 uint = 65535
    var v5 uint = v4
    var v15 int8 = -1
    v15 = (v15 >> ((42) & 63))
    if false {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS2{F0: 1, F1: 42, F2: 0}
            var v18 int = (fi16 + (-(fi16)))
            fmt.Println(fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        v15 -= v15
    }
    v19 := GenS2{F0: 4294967295, F1: v15, F2: 2}
    if (false && false) {
        var v20 int8 = -121
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float64 = (1e10 + ((1.5 - ((float64(1e10) - ((-2.25 * (0.0 * (-2.25 + 3.14159))) + 0.1)) - 0.1)) / float64(0.1)))
            if true {
                var v23 uint8 = 255
                v24 := GenS3{F0: genFn0((-(v2)), (v4 & (v5 | 65535))), F1: (42 <= (v3 | v3)), F2: 42}
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := []int{fi25}
                    var v27 float32 = 0.0
                    var v28 float64 = (-2.25 + 1.5)
                    var v29 bool = true
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = v23
                _ = v24
            } else {
                _ = len(v1)
                fmt.Println(v2)
                v30 := []int{(fi21 << ((v5) & 63)), fi21}
                _ = v30
            }
            _ = fi21
            _ = v22
        }
        v4 = (^(v5))
        _ = v20
    } else {
        for fi31 := 0; fi31 < 7; fi31++ {
            switch "\n\t" {
            case "\n\t":
                v32 := GenS2{F0: 0, F1: v15, F2: 2}
                _ = v32
            case "foo bar":
                fmt.Println(fi31)
            case "a":
                if (!true) {
                    var v33 uint8 = 140
                    _ = v33
                } else {
                    v5 += v4
                    v34 := GenS3{F0: genFn0((v2 * (-(v2))), v4), F1: true, F2: 2147483647}
                    var v35 uint8 = 255
                    _ = v34
                    _ = v35
                }
            default:
                fmt.Println(v3)
            }
            v15 += (v15 - v15)
            _ = fi31
        }
    }
    v36 := []int{int(v15), 128, 324}
    switch 32767 {
    case 9223372036854775807:
        v5 = v5
    case 42:
        v37 := v1[539]
        _ = v37
    }
    if (!true) {
        v2 -= v2
        for fi38 := 0; fi38 < 4; fi38++ {
            var v39 int32 = v3
            for fi40 := 0; fi40 < 15; fi40++ {
                v5 += v4
                _ = fi40
            }
            v41 := GenS4{F0: 1, F1: v2, F2: -861}
            _ = fi38
            _ = v39
            _ = v41
        }
    }
    var v42 bool = false
    var v43 int16 = v2
    switch "hello" {
    case "hello":
        var v44 float64 = ((1.5 + ((1e10 - (0.1 / (3.14159 + 3.14159))) * (1e10 / (-2.25 / float64(-2.25))))) + -2.25)
        _ = v44
    case "a":
        switch "hello" {
        case "":
            fmt.Println(v5, v3)
        case "a":
            for fi45 := 0; fi45 < 17; fi45++ {
                var v46 uint = v5
                var v47 uint64 = 584
                if len(v36) > 0 {
                    v36[0] = -942
                }
                _ = fi45
                _ = v46
                _ = v47
            }
        default:
            for fi48 := 0; fi48 < 9; fi48++ {
                if false {
                    fmt.Println(v15)
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v50 := GenS5{F0: v3, F1: (fi48 << ((v4) & 63)), F2: (("" + (("\n\t" + "a") + "hello")) + "foo bar")}
                        _ = fi49
                        _ = v50
                    }
                } else {
                    fmt.Println(v42)
                }
                _ = fi48
            }
        }
    case "\n\t":
        v51 := GenS1{F0: 127, F1: v2, F2: v42}
        _ = v51
    default:
        var v52 int16 = v2
        _ = v52
    }
    _ = len(v1)
    if v42 {
        v3 = v3
        switch "foo bar" {
        case "":
            if (v2 >= v43) {
                v42 = v42
                v53 := GenS4{F0: int64(v15), F1: 32767, F2: 1}
                if false {
                    for fi54 := 0; fi54 < 19; fi54++ {
                        for fi55 := 0; fi55 < 18; fi55++ {
                            var v56 int32 = (v3 << (((^(v5))) & 63))
                            var v57 uint = (v5 ^ 255)
                            var v58 uint = (v57 & v57)
                            _ = fi55
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        v59 := GenS4{F0: -214, F1: v43, F2: 9223372036854775807}
                        _ = len(v36)
                        _ = fi54
                        _ = v59
                    }
                    var v60 uint16 = 2
                    _ = v60
                }
                _ = v53
            }
        case "a":
            v61 := GenS3{F0: -602, F1: (v42 && ((0.1 * (1.5 - (1e10 * 0.0))) < 1.5)), F2: 4294967295}
            _ = v61
        default:
            var v62 bool = true
            _ = v62
        }
        switch -1 {
        case 0:
            v4 += v4
        }
    }
    var v63 uint = (v5 ^ v4)
    fmt.Println(v5, v63, v63)
    fmt.Println(v15, v4)
    if len(v36) > 0 {
        v36[0] = genFn0(v43, 2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v63
}

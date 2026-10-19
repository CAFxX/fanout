package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint; F1 bool }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int32 = 50
        if false {
            v8 := []int{230}
            _ = v8
        }
        var v9 int8 = 125
        _ = fi6
        _ = v7
        _ = v9
    }
    if (65535 == 65535) {
        switch "foo bar" {
        case "hello":
            var v10 float32 = float32(0.0)
            _ = v10
        case "foo bar":
            v11 := GenS1{F0: 18446744073709551615}
            _ = v11
        }
        if true {
            switch ("foo bar" + "a") {
            case "a":
                if (1e10 < (0.0 * (float32(3.14159) / 0.1))) {
                    var v12 uint64 = 742
                    if true {
                        var v13 uint = 243
                        var v14 int32 = -2147483648
                        _ = v13
                        _ = v14
                    }
                    var v15 bool = false
                    _ = v12
                    _ = v15
                }
            case "\n\t":
                var v16 float64 = 1.5
                _ = v16
            }
        }
    }
    var v17 uint = 744
    var v18 uint64 = 702
    v19 := []int{42, 1}
    _ = v17
    _ = v18
    _ = v19
    return 127
}
func genFn1(p0 int32) string {
    p0 = 0
    var v20 bool = true
    fmt.Println(p0)
    v21 := GenS1{F0: 18446744073709551615}
    if (42 <= -792) {
        switch "\n\t" {
        case "\n\t":
            v22 := []int{127, -101}
            _ = v22
        case "a":
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int64 = 1
                for fi25 := 0; fi25 < 4; fi25++ {
                    if (!false) {
                        p0 = (-(p0))
                        switch ("\n\t" + "") {
                        case "a":
                            fmt.Println(fi23)
                        case "foo bar":
                            var v26 int16 = 69
                            _ = v26
                        default:
                            v20 = false
                        }
                        var v27 int8 = 100
                        _ = v27
                    }
                    switch ("hello" + "hello") {
                    case "hello":
                        v20 = (1 > 933)
                    case "foo bar":
                        v28 := GenS2{F0: 4294967295, F1: 2}
                        _ = v28
                    case "a":
                        fmt.Println(fi25)
                    }
                    if v20 {
                        var v29 uint32 = 327
                        if (814 >= 2) {
                            fmt.Println(v20, p0, v29)
                            p0 -= 42
                            var v30 string = ("a" + "")
                            _ = v30
                        }
                        var v31 uint32 = (^(v29))
                        _ = v29
                        _ = v31
                    }
                    _ = fi25
                }
                _ = fi23
                _ = v24
            }
        default:
            v32 := GenS3{F0: 100, F1: -2, F2: -128}
            _ = v32
        }
        var v33 int16 = 265
        _ = v33
    }
    _ = p0
    _ = v20
    _ = v21
    return "a"
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v34 uint = 0
    v35 := []int{127}
    var v36 int8 = 0
    fmt.Println(v36, v34, v34)
    v1 = v1
    fmt.Println(v36, v36, v34)
    var v37 uint64 = 189
    fmt.Println(v1, v1)
    var v38 int32 = 1
    if false {
        if false {
            if (1e10 >= 3.14159) {
                v39 := GenS4{F0: 194, F1: 0, F2: (^(v36))}
                for fi40 := 0; fi40 < 11; fi40++ {
                    var v41 uint8 = 255
                    var v42 float32 = 3.14159
                    v5[(v1 + "foo bar")] = (1.5 != float64(float32(3.14159)))
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                if (!false) {
                    var v43 uint64 = v37
                    v37 -= v43
                    if ((^(v36)) >= v36) {
                        v37 += v37
                        switch 4294967295 {
                        case -263:
                            var v44 int = 42
                            _ = v44
                        default:
                            v45 := GenS4{F0: 109, F1: 32767, F2: v36}
                            _ = v45
                        }
                        if (!true) {
                            v35 = append(v35, 4294967295)
                            var v46 uint16 = 1
                            var v47 uint = (v34 * 65535)
                            _ = v46
                            _ = v47
                        }
                    }
                    _ = v43
                }
                _ = v39
            } else {
                fmt.Println(v38, v38, v36)
                v48 := GenS4{F0: 781, F1: 65535, F2: v36}
                _ = v48
            }
            var v49 int8 = (v36 << ((v34) & 63))
            _ = v49
        } else {
            var v50 uint16 = 1
            _ = v50
        }
    } else {
        var v51 uint64 = v37
        if true {
            switch "" {
            case "foo bar":
                v37 -= (v37 / ((v51) | 1))
            case "":
                if ((float32(0.0) + ((3.14159 * float32(3.14159)) + (1.5 + 0.1))) >= 0.0) {
                    for fi52 := 0; fi52 < 18; fi52++ {
                        v53 := GenS3{F0: v36, F1: 0, F2: 0}
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        v34 += v34
                        _ = fi52
                        _ = v53
                    }
                }
            case "hello":
                if (!false) {
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v5["\n\t"] = true
                        v34 = (v34 + (v34 * v34))
                        fmt.Println(fi54, fi54, v36)
                        _ = fi54
                    }
                    var v55 uint16 = uint16(v51)
                    _ = v55
                } else {
                    v56 := GenS5{F0: (v34 + (v34 * (v34 | v34))), F1: (-9223372036854775808 < -9223372036854775808)}
                    _ = v56
                }
            default:
                for fi57 := 0; fi57 < 20; fi57++ {
                    fmt.Println(v38, v1)
                    var v58 int32 = v38
                    v1 = v1
                    _ = fi57
                    _ = v58
                }
            }
        }
        _ = v51
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        delete(v3, 938)
        switch (-(fi59)) {
        case 2147483647:
            switch (("\n\t" + v1) + (("foo bar" + v1) + v1)) {
            case "":
                v60 := v3[65535]
                _ = v60
            case "\n\t":
                for fi61 := 0; fi61 < 18; fi61++ {
                    fmt.Println(v36)
                    var v62 int8 = v36
                    _ = fi61
                    _ = v62
                }
            default:
                fmt.Println(v37, v36)
            }
        default:
            v38 = (v38 >> ((v34) & 63))
        }
        _ = len(v35)
        _ = fi59
    }
    if true {
        if true {
            for fi63 := 0; fi63 < 14; fi63++ {
                v64 := GenS4{F0: 42, F1: 0, F2: v36}
                _ = fi63
                _ = v64
            }
            var v65 int32 = (v38 % ((v38) | 1))
            _ = v65
        }
        var v66 int16 = -1
        _ = v66
    } else {
        var v67 int64 = 55
        _ = v67
    }
    var v68 float32 = 0.0
    v69 := GenS1{F0: 465}
    v38 -= v38
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v68
    _ = v69
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 int64, p1 int16, p2 int16) int32 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    var v6 string = "foo bar"
    var v7 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 bool = true
    fmt.Println(v3)
    if (1e10 != (3.14159 + 3.14159)) {
        var v8 int16 = 908
        var v9 int = 128
        _ = v8
        _ = v9
    } else {
        v3 = (0 == genFn0(42, -302, -1))
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "foo bar":
            fmt.Println(v3)
        default:
            switch ("a" + ("\n\t" + ("a" + "a"))) {
            case "hello":
                v3 = (false || v3)
            case "a":
                v10 := []int{}
                _ = v10
            case "":
                for fi11 := 0; fi11 < 11; fi11++ {
                    if v3 {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            v3 = (fi11 <= (fi11 >> ((2) & 63)))
                            if (false || (!((v3 || (42 != 127)) && true))) {
                                v13 := []int{fi11}
                                var v14 bool = v3
                                var v15 uint16 = 42
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            var v16 uint8 = 128
                            _ = fi12
                            _ = v16
                        }
                        if false {
                            switch (-(fi11)) {
                            case 2147483647:
                                for fi17 := 0; fi17 < 6; fi17++ {
                                    var v18 string = (("hello" + ("foo bar" + "")) + ("a" + "hello"))
                                    var v19 float64 = (((1e10 / 0.1) + -2.25) - 0.0)
                                    v18 = v18
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            case -9223372036854775808:
                                v20 := GenS1{F0: 1, F1: -2147483648, F2: -128}
                                _ = v20
                            case 79:
                                switch (-(fi11)) {
                                case 1:
                                    var v21 bool = (42 >= 2)
                                    _ = v21
                                case 0:
                                    if v3 {
                                        for fi22 := 0; fi22 < 16; fi22++ {
                                            var v23 int8 = 42
                                            fmt.Println(v23)
                                            _ = fi22
                                            _ = v23
                                        }
                                        v3 = (956 != 243)
                                        for fi24 := 0; fi24 < 18; fi24++ {
                                            if v3 {
                                                v3 = (v3 && v3)
                                                var v25 uint16 = 965
                                                switch fi24 {
                                                case -758:
                                                    var v26 bool = v3
                                                    _ = v26
                                                case -2147483648:
                                                    v27 := GenS2{F0: -1, F1: (v25 >> ((65535) & 63))}
                                                    _ = v27
                                                case 32768:
                                                    v25 = (v25 * v25)
                                                }
                                                _ = v25
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(fi11, v3)
                                                default:
                                                    if (!(-1 >= -2147483648)) {
                                                        var v28 int16 = 465
                                                        var v29 uint16 = 62
                                                        v3 = v3
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        v30 := GenS3{F0: 9223372036854775807, F1: "", F2: 42}
                                                        if v3 {
                                                            v31 := []int{}
                                                            _ = v31
                                                        }
                                                        _ = v30
                                                    }
                                                }
                                            }
                                            var v32 uint64 = 449
                                            _ = fi24
                                            _ = v32
                                        }
                                    } else {
                                        for fi33 := 0; fi33 < 15; fi33++ {
                                            v3 = (943 < 0)
                                            _ = fi33
                                        }
                                        var v34 uint64 = 18446744073709551615
                                        _ = v34
                                    }
                                case -441:
                                    for fi35 := 0; fi35 < 15; fi35++ {
                                        if (!v3) {
                                            var v36 int = (fi35 - fi11)
                                            if (((3.14159 + (0.0 - 0.1)) - (3.14159 / (1.5 - 3.14159))) != float32(1e10)) {
                                                if v3 {
                                                    switch fi35 {
                                                    case 1:
                                                        v36 += (fi35 % ((v36) | 1))
                                                    case 251:
                                                        var v37 float32 = 0.0
                                                        _ = v37
                                                    case 128:
                                                        v38 := GenS4{F0: 559}
                                                        _ = v38
                                                    default:
                                                        if (42 > 42) {
                                                            fmt.Println(v36)
                                                        }
                                                    }
                                                    if (748 != 42) {
                                                        switch ("a" + "foo bar") {
                                                        case "foo bar":
                                                            v39 := GenS5{F0: 2147483648, F1: "foo bar"}
                                                            _ = v39
                                                        }
                                                        if (genFn0(-9223372036854775808, -1, 0) >= 1) {
                                                            v40 := v1[75]
                                                            var v41 string = v40
                                                            v42 := []int{fi11, 127, v36, (fi35 & fi11)}
                                                            _ = v40
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            fmt.Println(v3, v3)
                                                            switch "foo bar" {
                                                            case "":
                                                                v3 = v3
                                                                v36 -= v36
                                                                var v43 int8 = 127
                                                                _ = v43
                                                            }
                                                            var v44 int = (^(fi11))
                                                            fmt.Println(v3, fi35, fi35)
                                                            v45 := []int{(v44 | (fi35 + v36)), v44}
                                                            v44 = fi11
                                                            var v46 int16 = 32767
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                    switch (v36 / ((v36) | 1)) {
                                                    case -464:
                                                        if (!v3) {
                                                            switch "a" {
                                                            case "\n\t":
                                                                v47 := []int{(fi35 & (fi35 * 336)), fi11, 4294967295}
                                                                _ = v47
                                                            default:
                                                                v48 := GenS5{F0: 0, F1: ("hello" + "foo bar")}
                                                                _ = v48
                                                            }
                                                            v36 = fi35
                                                        } else {
                                                            v36 -= fi11
                                                            v36 = fi11
                                                        }
                                                    case -892:
                                                        v36 = 116
                                                    case 42:
                                                        v49 := v1[0]
                                                        _ = v49
                                                    default:
                                                        var v50 uint8 = 144
                                                        _ = v50
                                                    }
                                                } else {
                                                    var v51 float32 = float32(1e10)
                                                    _ = v51
                                                }
                                                fmt.Println(v3, v3)
                                                fmt.Println(v3, fi11, v36)
                                            }
                                            v52 := []int{v36, (fi35 | (fi35 * (fi35 * (fi35 | (fi35 * 0))))), (v36 & fi35)}
                                            _ = v36
                                            _ = v52
                                        }
                                        _ = fi35
                                    }
                                }
                            default:
                                for fi53 := 0; fi53 < 3; fi53++ {
                                    switch "foo bar" {
                                    case "hello":
                                        v1[65535] = "foo bar"
                                    }
                                    _ = fi53
                                }
                            }
                        }
                    } else {
                        v54 := GenS6{F0: 2, F1: -32768, F2: 230}
                        if v3 {
                            for fi55 := 0; fi55 < 10; fi55++ {
                                if (v3 && (369 == 42)) {
                                    fmt.Println(fi55, fi11, v3)
                                }
                                var v56 uint32 = 238
                                _ = fi55
                                _ = v56
                            }
                        } else {
                            var v57 uint16 = 1
                            var v58 float32 = (1e10 / (1.5 * 1.5))
                            if v3 {
                                v58 = v58
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    fmt.Println(v58, v58, fi59)
                                    v60 := GenS5{F0: 320, F1: "a"}
                                    _ = fi59
                                    _ = v60
                                }
                                var v61 float64 = -2.25
                                _ = v61
                            }
                            _ = v57
                            _ = v58
                        }
                        for fi62 := 0; fi62 < 9; fi62++ {
                            v63 := GenS6{F0: 42, F1: 1, F2: 255}
                            _ = fi62
                            _ = v63
                        }
                        _ = v54
                    }
                    _ = fi11
                }
            default:
                v3 = (1 == 0)
            }
        }
        v64 := GenS6{F0: 0, F1: 1, F2: 128}
        _ = v64
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v3 = true
    var v65 uint16 = 115
    var v66 uint64 = 42
    if (255 < 2) {
        if (42 >= -922) {
            var v67 uint = 42
            _ = v67
        }
        var v68 uint16 = 42
        fmt.Println(v68, v68)
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
}

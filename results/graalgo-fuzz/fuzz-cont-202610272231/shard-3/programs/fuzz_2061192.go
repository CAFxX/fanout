package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() uint16 {
    if false {
        v6 := GenS1{F0: 100, F1: -483, F2: 2147483648}
        v7 := []int{42}
        _ = v6
        _ = v7
    }
    var v8 uint32 = 2147483647
    if (255 >= 0) {
        v8 += 0
        switch "hello" {
        case "":
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    v11 := []int{fi10, (fi9 + fi10), 1, fi10}
                    switch "hello" {
                    case "":
                        fmt.Println(v8, fi10)
                    case "foo bar":
                        v8 = v8
                    default:
                        v8 = v8
                    }
                    if false {
                        var v12 uint16 = 32767
                        var v13 uint64 = 18446744073709551615
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                }
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 uint64 = 432
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: 127, F1: 1, F2: 0}
                _ = fi9
                _ = v16
            }
        }
        if false {
            var v17 int16 = 606
            if (false || (3.14159 != 0.0)) {
                v18 := []int{42}
                if true {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        fmt.Println(fi19, v8)
                        v8 += 883
                        switch fi19 {
                        case 0:
                            for fi20 := 0; fi20 < 18; fi20++ {
                                var v21 float64 = float64(0.0)
                                var v22 uint16 = 0
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case 127:
                            if true {
                                if (!false) {
                                    v8 += 65535
                                } else {
                                    if len(v18) > 0 {
                                        v18[0] = (fi19 - -716)
                                    }
                                    fmt.Println(v17)
                                }
                            }
                        case -250:
                            fmt.Println(v17)
                        }
                        _ = fi19
                    }
                    v17 = v17
                } else {
                    _ = len(v18)
                    if true {
                        v18 = append(v18, 42)
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v17 = (v17 % (((-(v17))) | 1))
                            _ = fi23
                        }
                    } else {
                        fmt.Println(v17, v8)
                    }
                    switch 2147483647 {
                    case 32767:
                        for fi24 := 0; fi24 < 7; fi24++ {
                            v25 := GenS1{F0: 108, F1: 0, F2: v8}
                            _ = fi24
                            _ = v25
                        }
                    case -2147483648:
                        var v26 uint16 = 478
                        _ = v26
                    case -1:
                        v8 += v8
                    default:
                        v18 = append(v18, 1)
                    }
                }
                _ = v18
            } else {
                switch "a" {
                case "foo bar":
                    v27 := GenS2{F0: 255, F1: 42}
                    _ = v27
                case "hello":
                    v28 := GenS2{F0: 170, F1: 1}
                    _ = v28
                default:
                    var v29 uint8 = 54
                    _ = v29
                }
                v30 := GenS3{F0: 42, F1: -128}
                var v31 int32 = -2147483648
                _ = v30
                _ = v31
            }
            v32 := GenS4{F0: 188, F1: ""}
            _ = v17
            _ = v32
        }
    }
    var v33 uint16 = 1
    _ = v33
    _ = v8
    return v33
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 1
    var v3 uint = 255
    v4 := make(map[string]uint32)
    var v5 float64 = (0.1 / float64(3.14159))
    v3 += (v3 * v3)
    v2 = 417
    if (!true) {
        v34 := []int{910, 32768, 915}
        for fi35 := 0; fi35 < 5; fi35++ {
            if false {
                v2 += 1
            }
            switch fi35 {
            case 42:
                v36 := GenS4{F0: 2, F1: ""}
                _ = v36
            }
            var v37 uint16 = 380
            _ = fi35
            _ = v37
        }
        _ = v34
    } else {
        switch ("" + "foo bar") {
        case "hello":
            switch 238 {
            case 128:
                for fi38 := 0; fi38 < 1; fi38++ {
                    if false {
                        v3 = v3
                        for fi39 := 0; fi39 < 13; fi39++ {
                            v40 := GenS5{F0: 2, F1: v3}
                            var v41 float32 = (1e10 + 1.5)
                            var v42 uint32 = 840
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        v3 = 879
                    }
                    v5 = v5
                    var v43 uint = (^(v3))
                    _ = fi38
                    _ = v43
                }
            case 31:
                if ((float32((v5 / float64(3.14159))) / -2.25) <= 0.1) {
                    var v44 int32 = 2147483647
                    for fi45 := 0; fi45 < 7; fi45++ {
                        var v46 int16 = 0
                        v44 = 1
                        _ = fi45
                        _ = v46
                    }
                    v44 -= v2
                    _ = v44
                }
            case -1:
                for fi47 := 0; fi47 < 3; fi47++ {
                    v2 = v2
                    _ = fi47
                }
            }
        case "foo bar":
            for fi48 := 0; fi48 < 19; fi48++ {
                if (0 == 42) {
                    fmt.Println(v2, v3, v3)
                }
                v49 := []int{}
                v2 = (v2 + (v2 * v2))
                _ = fi48
                _ = v49
            }
        }
    }
    var v50 int8 = 0
    var v51 string = (((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + ("foo bar" + "hello"))) + "a")
    var v52 int16 = -1
    v2 -= 42
    var v53 float64 = v5
    var v54 bool = (1 < (^(v3)))
    v55 := GenS3{F0: -2147483648, F1: 1}
    v56 := v4[v51]
    if ((v52 - (v52 + (v52 * v52))) > v52) {
        if true {
            var v57 uint32 = (v56 / (((^(v56))) | 1))
            _ = v57
        }
        for fi58 := 0; fi58 < 11; fi58++ {
            var v59 int = 9223372036854775807
            _ = fi58
            _ = v59
        }
        var v60 string = (((v51 + v51) + "hello") + "\n\t")
        _ = v60
    } else {
        var v61 uint32 = (v56 + 2147483648)
        fmt.Println(v52, v2, v54)
        var v62 int16 = -692
        _ = v61
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}

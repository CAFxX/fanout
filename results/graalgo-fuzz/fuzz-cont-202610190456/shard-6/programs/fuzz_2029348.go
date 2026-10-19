package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
func genFn0(p0 int64) int8 {
    p0 = p0
    p0 = p0
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: v6}
        switch "a" {
        case "":
            if false {
                if (65535 < (^(v6))) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        v10 := []int{9223372036854775807, fi9, (^(fi9)), (fi9 / ((fi9) | 1))}
                        _ = len(v10)
                        var v11 int = fi9
                        var v12 uint64 = 678
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS2{F0: -2147483648, F1: 1000}
                    var v14 uint = 0
                    _ = v13
                    _ = v14
                }
                var v15 float64 = 0.1
                if false {
                    if (((614 == -2147483648) && true) && false) {
                        fmt.Println(fi7, v6)
                        switch "" {
                        case "\n\t":
                            var v16 uint = 42
                            _ = v16
                        case "a":
                            var v17 int16 = 42
                            _ = v17
                        case "":
                            v18 := GenS2{F0: 1, F1: 32767}
                            _ = v18
                        default:
                            switch fi7 {
                            case 4294967295:
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    v15 += (v15 / (((v15 + v15) + v15) + v15))
                                    var v20 float32 = 1.5
                                    v21 := GenS1{F0: 405}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                        }
                        p0 -= p0
                    }
                } else {
                    v6 -= v6
                }
                _ = v15
            }
        }
        _ = fi7
        _ = v8
    }
    var v22 string = (("hello" + "\n\t") + "a")
    _ = p0
    _ = v22
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (65535 != 0)
    var v3 uint64 = 640
    v4 := []int{9223372036854775807, 127, 0, -9223372036854775808}
    var v5 uint16 = 0
    var v23 string = ("" + "a")
    _ = len(v1)
    var v24 uint32 = uint32(v3)
    for fi25 := 0; fi25 < 19; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v5, v23, v24)
            var v27 int32 = 42
            var v28 uint16 = v5
            var v29 int64 = 0
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        if (42 <= -46) {
            v5 = v5
            v30 := GenS1{F0: (v5 & 940)}
            for fi31 := 0; fi31 < 13; fi31++ {
                if v2 {
                    fmt.Println(v24)
                    var v32 uint64 = 0
                    if v2 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v2 {
                            v2 = (44 == 127)
                        }
                    } else {
                        if false {
                            if (v2 && v2) {
                                var v33 bool = false
                                fmt.Println(v32)
                                _ = v33
                            }
                            v23 = "foo bar"
                            if (v23 >= "hello") {
                                v34 := GenS3{F0: 0, F1: (v3 | v3), F2: 108}
                                _ = v34
                            } else {
                                var v35 bool = false
                                _ = v35
                            }
                        } else {
                            fmt.Println(v2, v23)
                        }
                        fmt.Println(v5, v5, fi25)
                        if false {
                            v2 = (0.0 <= 3.14159)
                            var v36 uint16 = 42
                            _ = v36
                        }
                        var v37 float64 = (0.1 / 0.1)
                        var v38 int = fi25
                        v39 := GenS4{F0: (v24 | (^(v24))), F1: (^(v24))}
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = v32
                }
                if (false && v2) {
                    delete(v1, (fi25 / ((-101) | 1)))
                    v40 := GenS1{F0: v5}
                    v24 = v24
                    _ = v40
                }
                _ = fi31
            }
            _ = v30
        }
        for fi41 := 0; fi41 < 9; fi41++ {
            switch "a" {
            case "\n\t":
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v23, v23, v5)
                case "":
                    v3 += uint64(v24)
                default:
                    var v42 bool = true
                    _ = v42
                }
            case "":
                v1[(fi25 + (fi25 >> ((2) & 63)))] = 9223372036854775807
            case "hello":
                var v43 uint16 = 42
                _ = v43
            default:
                var v44 int8 = 0
                _ = v44
            }
            if v2 {
                if v2 {
                    delete(v1, fi25)
                }
            }
            var v45 float32 = 0.0
            switch v23 {
            case "":
                _ = len(v1)
            }
            if v2 {
                var v46 int = fi41
                _ = v46
            }
            if len(v4) > 0 {
                v4[0] = (^(fi25))
            }
            _ = fi41
            _ = v45
        }
        _ = fi25
    }
    _ = len(v4)
    var v47 uint16 = v5
    var v48 uint16 = (v47 ^ (v47 | (v5 - 65535)))
    for fi49 := 0; fi49 < 8; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := GenS5{F0: 0}
            var v52 uint64 = v3
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi49
    }
    v53 := GenS6{F0: -122}
    v5 = v5
    fmt.Println(v2)
    fmt.Println(v48, v23)
    for fi54 := 0; fi54 < 19; fi54++ {
        for fi55 := 0; fi55 < 11; fi55++ {
            var v56 bool = false
            _ = fi55
            _ = v56
        }
        _ = fi54
    }
    _ = len(v4)
    v57 := GenS2{F0: int32(v3), F1: -1}
    for fi58 := 0; fi58 < 11; fi58++ {
        v59 := GenS7{F0: ("" + ((v23 + v23) + (v23 + v23)))}
        var v60 uint8 = 255
        for fi61 := 0; fi61 < 14; fi61++ {
            v2 = (v47 < 255)
            _ = fi61
        }
        _ = fi58
        _ = v59
        _ = v60
    }
    var v62 float32 = 3.14159
    v48 += v48
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v53
    _ = v57
    _ = v62
}

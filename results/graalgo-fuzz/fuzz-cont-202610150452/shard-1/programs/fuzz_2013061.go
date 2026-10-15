package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int16; F2 uint }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 bool; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 string = "\n\t"
    var v3 int32 = 2147483647
    v4 := make(map[uint32]int)
    v5 := []int{726, 1}
    v6 := GenS1{F0: 128, F1: -1, F2: 0}
    v3 = v3
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 string = v2
        var v9 float32 = float32(1.5)
        if true {
            fmt.Println(v2, v3, v9)
            var v10 int = (fi7 % ((fi7) | 1))
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = v3
    v2 = "\n\t"
    switch "a" {
    case "a":
        var v11 uint64 = 2
        _ = v11
    case "\n\t":
        if (!true) {
            if true {
                fmt.Println(v2)
            }
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint8 = 19
                var v14 int = (^(fi12))
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint32 = 2
            _ = v15
        }
    case "hello":
        _ = len(v5)
    }
    var v16 float64 = (3.14159 - float64(0.1))
    for fi17 := 0; fi17 < 4; fi17++ {
        switch fi17 {
        case -900:
            var v18 float64 = ((1.5 - v16) - v16)
            _ = v18
        case 42:
            v19 := v1[(fi17 / (((fi17 * fi17)) | 1))]
            _ = v19
        case 189:
            v3 -= (^(v3))
        default:
            v20 := GenS2{F0: -128, F1: 2147483647}
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := GenS3{F0: 134}
            var v23 float64 = v16
            var v24 int64 = -572
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v25 uint = 4294967295
    var v26 uint64 = 901
    v5 = append(v5, 9223372036854775807)
    if false {
        for fi27 := 0; fi27 < 17; fi27++ {
            switch v2 {
            case "foo bar":
                if true {
                    v26 = (v26 % ((v26) | 1))
                }
            case "\n\t":
                for fi28 := 0; fi28 < 20; fi28++ {
                    v2 = ""
                    if len(v5) > 0 {
                        v5[0] = fi27
                    }
                    if false {
                        var v29 int16 = 455
                        var v30 int = (fi27 / ((fi28) | 1))
                        _ = v29
                        _ = v30
                    } else {
                        if len(v5) > 0 {
                            v5[0] = (fi27 ^ -2147483648)
                        }
                        var v31 float64 = v16
                        v3 -= (v3 + 0)
                        _ = v31
                    }
                    _ = fi28
                }
            case "hello":
                v2 = v2
            }
            _ = fi27
        }
    } else {
        switch 4294967295 {
        case -2147483648:
            _ = len(v4)
        case 42:
            v16 = v16
        case 2147483647:
            for fi32 := 0; fi32 < 20; fi32++ {
                v2 = v2
                switch (fi32 % ((fi32) | 1)) {
                case 32767:
                    v33 := GenS4{F0: -520, F1: 1000, F2: (v25 & (v25 & v25))}
                    _ = v33
                default:
                    if (2 > 2) {
                        if false {
                            switch (v2 + "a") {
                            case "hello":
                                var v34 float32 = 3.14159
                                _ = v34
                            }
                        } else {
                            if false {
                                for fi35 := 0; fi35 < 20; fi35++ {
                                    var v36 string = ((v2 + ((v2 + v2) + v2)) + ("a" + "foo bar"))
                                    var v37 string = "foo bar"
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                switch v2 {
                                case "\n\t":
                                    for fi38 := 0; fi38 < 4; fi38++ {
                                        var v39 int16 = 411
                                        switch (v2 + (v2 + v2)) {
                                        case "foo bar":
                                            if (!(false && false)) {
                                                v40 := GenS5{F0: (v39 & 0), F1: (v3 >> (((v25 - (v25 * v25))) & 63)), F2: (v2 + v2)}
                                                v41 := GenS3{F0: 161}
                                                v42 := v4[0]
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v4[2147483648] = 1
                                                var v43 int32 = -646
                                                fmt.Println(fi38, fi32)
                                                v25 += v25
                                                if ((v39 % ((v39) | 1)) >= v39) {
                                                    var v44 float64 = 0.1
                                                    _ = v44
                                                } else {
                                                    fmt.Println(v2)
                                                    var v45 float64 = v16
                                                    _ = v45
                                                }
                                                _ = v43
                                            }
                                        case "\n\t":
                                            v25 -= v25
                                        case "hello":
                                            _ = len(v5)
                                        default:
                                            if (!false) {
                                                v26 += v26
                                                switch fi38 {
                                                case 32768:
                                                    fmt.Println(v3, fi38)
                                                case 2147483647:
                                                    var v46 float64 = v16
                                                    _ = v46
                                                case 4294967295:
                                                    var v47 string = "hello"
                                                    _ = v47
                                                }
                                            } else {
                                                v48 := GenS3{F0: 1}
                                                v49 := GenS6{F0: v2, F1: true}
                                                switch 127 {
                                                case 4294967295:
                                                    switch v2 {
                                                    case "hello":
                                                        v50 := GenS7{F0: 127}
                                                        _ = v50
                                                    }
                                                }
                                                if (false || (-2147483648 <= v3)) {
                                                    v4[195] = (-(fi32))
                                                    v5 = append(v5, 32768)
                                                    v26 = (v26 | (v26 ^ v26))
                                                    if false {
                                                        var v51 uint32 = 1
                                                        var v52 int16 = (v39 >> ((v25) & 63))
                                                        var v53 bool = (true || true)
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        v54 := v1[(fi32 ^ fi32)]
                                                        _ = v54
                                                    }
                                                }
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        switch ("" + v2) {
                                        case "a":
                                            switch (v2 + v2) {
                                            case "":
                                                if false {
                                                    v39 += 0
                                                } else {
                                                    fmt.Println(v16, v16, v25)
                                                    if (float32(0.1) < 1e10) {
                                                        v26 = v26
                                                        v25 -= 1
                                                    }
                                                }
                                            case "hello":
                                                v16 = float64(3.14159)
                                            case "a":
                                                var v55 uint = v25
                                                _ = v55
                                            }
                                        default:
                                            fmt.Println(fi32, fi38, v26)
                                        }
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            }
                            for fi56 := 0; fi56 < 20; fi56++ {
                                delete(v1, (^(fi56)))
                                if ((true || (42 <= 42)) && (false || (!true))) {
                                    var v57 float32 = 0.0
                                    _ = v57
                                }
                                fmt.Println(v26, v26)
                                switch "foo bar" {
                                case "a":
                                    if false {
                                        v26 -= v26
                                        var v58 bool = ("foo bar" < ("foo bar" + v2))
                                        _ = v58
                                    } else {
                                        v59 := GenS1{F0: 42, F1: -128, F2: 0}
                                        _ = v59
                                    }
                                default:
                                    v60 := v4[2147483647]
                                    _ = v60
                                }
                                _ = fi56
                            }
                            switch int(v26) {
                            case -675:
                                var v61 int16 = -174
                                _ = v61
                            default:
                                fmt.Println(v2, v16)
                            }
                        }
                        var v62 float64 = v16
                        v63 := GenS8{F0: (fi32 / ((fi32) | 1)), F1: true, F2: (false || ((!false) && false))}
                        _ = v62
                        _ = v63
                    } else {
                        v3 = -666
                        var v64 bool = false
                        _ = v64
                    }
                }
                _ = fi32
            }
        default:
            v26 += (v26 / (((v26 | v26)) | 1))
        }
    }
    v3 -= (^(v3))
    var v65 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
}

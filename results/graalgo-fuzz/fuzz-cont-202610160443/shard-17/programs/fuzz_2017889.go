package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 string; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int64 }
func genFn0(p0 bool, p1 uint32, p2 uint32) int64 {
    p1 = (p1 + (^(p2)))
    if (-1 <= 1) {
        var v4 int32 = 2147483647
        v5 := []int{345, 128, 9223372036854775807, -583}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 42}
            fmt.Println(p2, p1)
            _ = fi6
            _ = v7
        }
        if (1 >= -1) {
            if p0 {
                p1 = p1
                v8 := GenS1{F0: 42}
                _ = v8
            }
        } else {
            if p0 {
                var v9 string = "\n\t"
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS2{F0: (0.0 < 1e10), F1: 623, F2: 127}
                    for fi12 := 0; fi12 < 18; fi12++ {
                        if (!true) {
                            switch (((("a" + (("" + ((v9 + v9) + v9)) + v9)) + ("hello" + "")) + v9) + ("a" + v9)) {
                            case "\n\t":
                                v13 := []int{}
                                _ = v13
                            case "a":
                                var v14 string = v9
                                _ = v14
                            default:
                                var v15 uint32 = p2
                                _ = v15
                            }
                            if p0 {
                                v16 := GenS3{F0: 42, F1: fi10, F2: 2}
                                var v17 uint32 = p2
                                _ = v16
                                _ = v17
                            }
                            var v18 float64 = ((float64(float32(float64(0.1))) + (-2.25 / 0.1)) * 1e10)
                            _ = v18
                        } else {
                            if p0 {
                                fmt.Println(fi12)
                            }
                            fmt.Println(p2, p1)
                            var v19 float32 = ((1e10 + (0.0 / (0.0 - -2.25))) + 0.0)
                            _ = v19
                        }
                        _ = fi12
                    }
                    _ = fi10
                    _ = v11
                }
                var v20 int64 = -1
                _ = v20
                _ = v9
            } else {
                v21 := GenS2{F0: false, F1: 18446744073709551615, F2: 0}
                _ = v21
            }
            switch 4294967295 {
            case 878:
                switch 551 {
                case 946:
                    if (!(!(int32(p2) <= -872))) {
                        v22 := []int{-2147483648}
                        _ = v22
                    } else {
                        var v23 bool = true
                        v24 := GenS1{F0: -128}
                        _ = v23
                        _ = v24
                    }
                default:
                    p1 = 126
                }
            case 41:
                if p0 {
                    v25 := GenS2{F0: p0, F1: 699, F2: 42}
                    _ = v25
                } else {
                    p2 += 694
                    var v26 int8 = -1
                    var v27 uint = 255
                    _ = v26
                    _ = v27
                }
            default:
                p2 += p1
            }
            if false {
                var v28 uint64 = 18446744073709551615
                if (!false) {
                    var v29 string = (((("a" + ((("" + "a") + "foo bar") + ("" + ""))) + ("a" + ("" + ("" + "foo bar")))) + (("a" + "foo bar") + "")) + ("" + "foo bar"))
                    _ = v29
                } else {
                    fmt.Println(p0, p0)
                }
                switch "hello" {
                case "\n\t":
                    var v30 bool = p0
                    _ = v30
                }
                _ = v28
            }
        }
        switch "a" {
        case "\n\t":
            var v31 uint32 = p2
            _ = v31
        case "":
            var v32 string = "foo bar"
            _ = v32
        case "hello":
            if p0 {
                v33 := GenS4{F0: true, F1: 128}
                var v34 uint32 = (p2 & (p2 << ((524) & 63)))
                switch 127 {
                case -67:
                    v35 := GenS1{F0: -1}
                    _ = v35
                case 480:
                    var v36 uint8 = 127
                    _ = v36
                default:
                    v37 := []int{127, 32767, 1, 32768}
                    _ = v37
                }
                _ = v33
                _ = v34
            }
        default:
            for fi38 := 0; fi38 < 3; fi38++ {
                v39 := GenS5{F0: "", F1: 1, F2: 103}
                if true {
                    v40 := GenS1{F0: 42}
                    _ = v40
                } else {
                    if (-477 <= 836) {
                        p1 = p1
                    }
                    var v41 int16 = -32768
                    for fi42 := 0; fi42 < 10; fi42++ {
                        p0 = (!((571 == 645) || false))
                        _ = fi42
                    }
                    _ = v41
                }
                if false {
                    fmt.Println(p0, p1)
                    v43 := GenS5{F0: "", F1: 255, F2: 1}
                    _ = v43
                }
                _ = fi38
                _ = v39
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}
func genFn1(p0 uint32) int64 {
    p0 -= (p0 >> ((27) & 63))
    v44 := GenS6{F0: (p0 * (p0 * p0)), F1: 2, F2: 9223372036854775807}
    _ = p0
    _ = v44
    return genFn0((true && (1 >= 100)), 0, (p0 & (p0 % (((p0 / (((p0 / ((2147483648) | 1))) | 1))) | 1))))
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 795
    var v3 uint16 = uint16(v2)
    v45 := []int{}
    var v46 int64 = 453
    var v47 uint32 = 436
    v47 = 209
    v3 = (v3 / (((v3 | v3)) | 1))
    v47 -= (v47 * v2)
    if false {
        fmt.Println(v46, v3)
        v1 = (v1 / float32(float64(v1)))
    } else {
        fmt.Println(v46, v3, v47)
    }
    if false {
        if true {
            fmt.Println(v3)
            if true {
                v3 -= (v3 << ((8) & 63))
            } else {
                var v48 uint32 = (^(v47))
                _ = v48
            }
        } else {
            for fi49 := 0; fi49 < 11; fi49++ {
                for fi50 := 0; fi50 < 13; fi50++ {
                    if (!((0 != 42) || true)) {
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        fmt.Println(v3, v2, v47)
                        switch (fi50 - fi50) {
                        case -2147483648:
                            var v51 float64 = (1e10 * (3.14159 * (1e10 - (1.5 * float64((1.5 / (v1 + 1e10)))))))
                            fmt.Println(v1, fi49, fi49)
                            _ = v51
                        case 32768:
                            var v52 uint32 = 2147483647
                            _ = v52
                        }
                    } else {
                        switch fi49 {
                        case -757:
                            v47 -= (v47 / ((371) | 1))
                        }
                    }
                    _ = fi50
                }
                var v53 float32 = v1
                _ = fi49
                _ = v53
            }
            if true {
                if (v3 > (v3 ^ (v3 << ((65535) & 63)))) {
                    fmt.Println(v47)
                    var v54 string = "a"
                    switch 939 {
                    case -2147483648:
                        for fi55 := 0; fi55 < 3; fi55++ {
                            _ = len(v45)
                            v45 = append(v45, int(v3))
                            _ = fi55
                        }
                    case 32767:
                        switch 4294967295 {
                        case 42:
                            v56 := GenS2{F0: true, F1: uint(v47), F2: 74}
                            _ = v56
                        default:
                            v57 := GenS1{F0: -71}
                            _ = v57
                        }
                    case 2147483647:
                        v58 := GenS1{F0: 15}
                        _ = v58
                    }
                    _ = v54
                } else {
                    v59 := GenS6{F0: (v2 % ((v47) | 1)), F1: 42, F2: (v46 * v46)}
                    _ = len(v45)
                    _ = v59
                }
            }
            v60 := GenS5{F0: "\n\t", F1: 42, F2: 0}
            _ = v60
        }
        if (false && false) {
            for fi61 := 0; fi61 < 4; fi61++ {
                v3 = (v3 >> ((0) & 63))
                if len(v45) > 0 {
                    v45[0] = -829
                }
                _ = fi61
            }
            var v62 int = -1
            v46 = (v46 + 1)
            _ = v62
        }
    }
    v47 -= (^(v47))
    if len(v45) > 0 {
        v45[0] = 9223372036854775807
    }
    v46 = (v46 / (((v46 + v46)) | 1))
    var v63 uint = 2
    if ((v63 * v63) != 100) {
        switch 644 {
        case 32768:
            if len(v45) >= 2 {
                v45 = v45[1:len(v45)]
            }
        case 144:
            switch "foo bar" {
            case "a":
                fmt.Println(v63)
            }
        case 1:
            var v64 int8 = 42
            _ = v64
        }
    } else {
        v45 = append(v45, int(v46))
        _ = len(v45)
        v65 := GenS2{F0: true, F1: 2, F2: 67}
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v63
}

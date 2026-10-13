package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 int64 }
func genFn0(p0 int32) string {
    v6 := GenS1{F0: 154, F1: false, F2: 982}
    fmt.Println(p0)
    if ("a" < (("" + "foo bar") + "")) {
        v7 := []int{-782, 32768, 127}
        var v8 int64 = -1
        p0 = 0
        _ = v7
        _ = v8
    }
    p0 = p0
    if (912 >= 0) {
        v9 := []int{-437, -2147483648, 32767}
        if true {
            p0 = -253
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int64 = 9
                    var v13 int = (fi11 % ((fi10) | 1))
                    if true {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v12 += 570
                    } else {
                        var v14 float32 = 0.0
                        if (false || (true && (444 >= (fi10 ^ v13)))) {
                            var v15 int64 = (v12 / ((42) | 1))
                            v16 := GenS1{F0: 1, F1: (32767 >= 1000), F2: fi10}
                            v17 := GenS2{F0: 9223372036854775807, F1: "foo bar"}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                    }
                    v13 = (fi11 * -278)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v20 := GenS3{F0: 4294967295, F1: 126, F2: 4294967295}
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 float32 = 0.0
                    var v23 uint = 289
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi18
            }
        } else {
            var v24 int8 = 97
            _ = v24
        }
        _ = v9
    }
    _ = p0
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!(7 < 128))
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 int = 957
    v25 := []int{}
    for fi26 := 0; fi26 < 16; fi26++ {
        if len(v25) > 0 {
            v25[0] = (^(fi26))
        }
        switch (-(v5)) {
        case 4294967295:
            v5 = v5
        case 9223372036854775807:
            _ = len(v4)
        case 1:
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint8 = 2
                _ = fi27
                _ = v28
            }
        default:
            v5 = (v5 - (v5 * v5))
        }
        _ = fi26
    }
    v2 = (v2 && (0 != 42))
    var v29 bool = (true || (488 == 0))
    v5 = 4294967295
    fmt.Println(v5)
    if v2 {
        var v30 bool = ((!(18446744073709551615 > 776)) || (v2 || ((-2.25 - (float64((float32(1.5) / 3.14159)) + 0.0)) == 1e10)))
        var v31 uint32 = 0
        switch "a" {
        case "":
            if true {
                v25 = append(v25, 42)
                v29 = true
                v2 = (!(!v2))
            }
        case "hello":
            var v32 bool = (!(false && (false && v2)))
            _ = v32
        default:
            if v29 {
                var v33 int32 = -142
                _ = v33
            }
        }
        _ = v30
        _ = v31
    } else {
        for fi34 := 0; fi34 < 14; fi34++ {
            v35 := GenS3{F0: 2147483647, F1: 822, F2: 65535}
            _ = fi34
            _ = v35
        }
        if (3.14159 >= (float32(float64((1e10 / 0.1))) * -2.25)) {
            _ = len(v25)
            if v29 {
                v29 = true
                v36 := v3[32768]
                _ = v36
            } else {
                for fi37 := 0; fi37 < 19; fi37++ {
                    var v38 int64 = 0
                    fmt.Println(v2)
                    _ = fi37
                    _ = v38
                }
                switch v5 {
                case -9223372036854775808:
                    switch "hello" {
                    case "":
                        v29 = v2
                    case "foo bar":
                        switch v5 {
                        case 42:
                            switch v5 {
                            case 32768:
                                fmt.Println(v5, v5, v29)
                            }
                        case 96:
                            v39 := GenS4{F0: (-32768 != 42), F1: 915, F2: ""}
                            _ = v39
                        }
                    case "hello":
                        v29 = v29
                    }
                case 9223372036854775807:
                    fmt.Println(v2)
                default:
                    if (1000 < 32767) {
                        switch v5 {
                        case -2147483648:
                            var v40 uint64 = 437
                            _ = v40
                        case -235:
                            fmt.Println(v29, v2, v29)
                        }
                        switch "hello" {
                        case "":
                            if v29 {
                                switch "a" {
                                case "a":
                                    var v41 string = ("foo bar" + "hello")
                                    _ = v41
                                case "hello":
                                    for fi42 := 0; fi42 < 17; fi42++ {
                                        v43 := GenS5{F0: 558}
                                        if true {
                                            var v44 uint16 = 32767
                                            switch -230 {
                                            case 0:
                                                if false {
                                                    switch fi42 {
                                                    case 9223372036854775807:
                                                        var v45 uint64 = 18446744073709551615
                                                        _ = v45
                                                    case 32768:
                                                        _ = len(v25)
                                                    default:
                                                        var v46 uint32 = 4294967295
                                                        _ = v46
                                                    }
                                                }
                                            default:
                                                var v47 uint8 = 255
                                                _ = v47
                                            }
                                            v2 = (((v2 && v29) || ((^(v44)) <= (v44 / (((v44 >> ((4294967295) & 63))) | 1)))) || (!((1 <= 0) || (!v2))))
                                            _ = v44
                                        }
                                        v25 = append(v25, (fi42 + fi42))
                                        _ = fi42
                                        _ = v43
                                    }
                                case "foo bar":
                                    if (v2 || (1 >= 0)) {
                                        fmt.Println(v2, v29)
                                        v29 = false
                                    }
                                }
                            } else {
                                fmt.Println(v29, v2)
                                fmt.Println(v5)
                                if ((!false) || (112 < 255)) {
                                    if v2 {
                                        v1[1] = 1
                                        v48 := GenS6{F0: 9223372036854775807, F1: 1000}
                                        _ = v48
                                    }
                                    _ = len(v25)
                                } else {
                                    v5 = (v5 & (^(v5)))
                                    if false {
                                        v5 += v5
                                    }
                                    if v29 {
                                        if v2 {
                                            switch 32768 {
                                            case 128:
                                                v2 = v2
                                            case -9223372036854775808:
                                                v29 = (2 == 370)
                                            case 2147483647:
                                                var v49 uint8 = 109
                                                _ = v49
                                            default:
                                                for fi50 := 0; fi50 < 11; fi50++ {
                                                    var v51 float32 = -2.25
                                                    if (!(771 == 4294967295)) {
                                                        v51 += ((((v51 / v51) + (1e10 * ((1.5 * 3.14159) + 3.14159))) * 1.5) * float32(-2.25))
                                                    } else {
                                                        var v52 string = "hello"
                                                        v29 = (!true)
                                                        var v53 uint8 = 128
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    _ = fi50
                                                    _ = v51
                                                }
                                            }
                                            var v54 int64 = 1
                                            _ = v54
                                        } else {
                                            var v55 uint64 = 4294967295
                                            _ = len(v25)
                                            var v56 string = "hello"
                                            _ = v55
                                            _ = v56
                                        }
                                    }
                                }
                            }
                        default:
                            var v57 int = (v5 + v5)
                            _ = v57
                        }
                        v58 := GenS4{F0: v2, F1: 394, F2: "hello"}
                        _ = v58
                    }
                }
                v59 := GenS7{F0: 42}
                _ = v59
            }
        } else {
            v60 := GenS2{F0: 469, F1: "hello"}
            _ = len(v25)
            var v61 uint8 = 2
            _ = v60
            _ = v61
        }
    }
    v25 = append(v25, 1)
    for fi62 := 0; fi62 < 18; fi62++ {
        v63 := GenS8{F0: 605, F1: 1}
        var v64 uint = 2
        _ = fi62
        _ = v63
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

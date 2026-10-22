package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 string; F1 string; F2 int64 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint32; F1 int }
func genFn0() uint16 {
    v5 := []int{-752, 4294967295, 0}
    if (652 >= 42) {
        _ = len(v5)
    }
    if true {
        for fi6 := 0; fi6 < 16; fi6++ {
            _ = len(v5)
            var v7 string = ""
            _ = fi6
            _ = v7
        }
    } else {
        var v8 uint16 = 255
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    v8 += 1
                    v8 += v8
                } else {
                    switch "a" {
                    case "\n\t":
                        var v11 int32 = -1
                        var v12 float64 = 3.14159
                        var v13 int = (fi9 / (((fi9 + (fi10 << ((558) & 63)))) | 1))
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                if true {
                    var v14 float32 = (-2.25 + (3.14159 - 0.1))
                    if true {
                        var v15 uint64 = 591
                        _ = v15
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v14 -= (((float32(((1.5 / (-2.25 / float64(float32(1e10)))) * ((float64(float32(3.14159)) + (0.1 / float64(((v14 / 3.14159) / (1.5 * ((v14 / 0.1) / v14)))))) * 0.0))) + (float32((0.1 / 1.0)) / v14)) / v14) * ((0.1 / float32(3.14159)) / v14))
                        }
                    }
                    v14 = v14
                    _ = v14
                }
                _ = fi10
            }
            _ = fi9
        }
        _ = v8
    }
    _ = v5
    return 0
}

func main() {
    var v1 bool = ((("hello" + (((("\n\t" + "hello") + "hello") + "a") + "\n\t")) + "foo bar") >= (("" + "foo bar") + ""))
    var v2 float64 = 0.0
    var v3 uint32 = 2
    var v4 string = "foo bar"
    v16 := []int{}
    v3 -= (v3 ^ (v3 % (((v3 | v3)) | 1)))
    fmt.Println(v2)
    v2 += v2
    switch "foo bar" {
    case "\n\t":
        var v17 uint32 = 75
        _ = v17
    case "a":
        switch "a" {
        case "hello":
            v3 = (v3 & v3)
        case "":
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint8 = 0
                _ = fi18
                _ = v19
            }
        default:
            var v20 float32 = (1e10 / 1.5)
            _ = v20
        }
    }
    var v21 float64 = 0.1
    v3 = (v3 & v3)
    switch (v4 + (v4 + "")) {
    case "a":
        fmt.Println(v21)
    case "":
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS1{F0: genFn0(), F1: v4, F2: 177}
            var v24 string = "hello"
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        v16 = append(v16, 128)
    }
    v21 -= v2
    fmt.Println(v4)
    v3 = (v3 ^ v3)
    var v25 uint32 = (v3 % ((v3) | 1))
    if (445 != (v25 | v3)) {
        switch v4 {
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            fmt.Println(v21, v21, v21)
        case "foo bar":
            v26 := GenS2{F0: 0}
            _ = v26
        default:
            if v1 {
                var v27 uint8 = 127
                for fi28 := 0; fi28 < 11; fi28++ {
                    if v1 {
                        fmt.Println(v4)
                    }
                    _ = fi28
                }
                var v29 bool = (!false)
                _ = v27
                _ = v29
            } else {
                for fi30 := 0; fi30 < 1; fi30++ {
                    var v31 uint16 = 32767
                    _ = fi30
                    _ = v31
                }
                v3 -= 716
            }
        }
        v32 := GenS3{F0: v25, F1: v25}
        v16 = append(v16, 4294967295)
        _ = v32
    } else {
        if v1 {
            if len(v16) > 0 {
                v16[0] = 42
            }
            switch "a" {
            case "hello":
                var v33 uint64 = 4294967295
                _ = v33
            case "":
                v34 := GenS2{F0: 32767}
                _ = v34
            }
            var v35 uint32 = (v3 >> ((42) & 63))
            _ = v35
        } else {
            if (-986 == -527) {
                if len(v16) > 0 {
                    v16[0] = 127
                }
                if (v1 || false) {
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v1 = (((v2 / float64(1.5)) - v2) > ((v21 * ((float64((3.14159 - 0.0)) * (-2.25 - (1.5 - v2))) / -2.25)) - -2.25))
                        var v37 int = fi36
                        v38 := GenS4{F0: ("\n\t" + ((("foo bar" + ("" + (v4 + v4))) + "hello") + (v4 + v4))), F1: v4, F2: 0}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                for fi39 := 0; fi39 < 6; fi39++ {
                    fmt.Println(v25, v21)
                    var v40 bool = true
                    switch 342 {
                    case -9223372036854775808:
                        if true {
                            switch 43 {
                            case 719:
                                var v41 string = ((v4 + "") + v4)
                                _ = v41
                            case 127:
                                v42 := GenS5{F0: v25, F1: 2147483647}
                                _ = v42
                            }
                            var v43 uint8 = 128
                            if len(v16) > 0 {
                                v16[0] = -1
                            }
                            _ = v43
                        }
                    case -412:
                        if v1 {
                            for fi44 := 0; fi44 < 8; fi44++ {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                v3 = (v25 << ((18446744073709551615) & 63))
                                v45 := GenS6{F0: 0, F1: 128, F2: 0}
                                if v40 {
                                    var v46 int8 = -128
                                    _ = v46
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v1 = (42 > 0)
                                    }
                                    if (!(((1e10 * (v2 - ((v21 + 1e10) - (0.1 + 1.5)))) + (v2 + 0.1)) == float64(((1e10 / 1.0) + (-2.25 + 0.1))))) {
                                        var v47 uint64 = 1
                                        if v40 {
                                            if false {
                                                var v48 uint8 = 202
                                                fmt.Println(v2, v48, v25)
                                                v2 -= v2
                                                _ = v48
                                            }
                                            var v49 string = (v4 + v4)
                                            var v50 uint16 = 2
                                            _ = v49
                                            _ = v50
                                        } else {
                                            var v51 int8 = 0
                                            v52 := GenS2{F0: 1}
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v47
                                    }
                                }
                                _ = fi44
                                _ = v45
                            }
                            v16 = append(v16, (-(fi39)))
                            if v40 {
                                switch fi39 {
                                case 127:
                                    v53 := GenS4{F0: (v4 + v4), F1: ("hello" + "hello"), F2: -1}
                                    _ = v53
                                case 4294967295:
                                    v54 := GenS7{F0: 2147483647, F1: v25}
                                    _ = v54
                                case 9223372036854775807:
                                    var v55 int64 = 42
                                    _ = v55
                                }
                            } else {
                                var v56 uint16 = 65535
                                _ = v56
                            }
                        }
                    case -369:
                        if v1 {
                            fmt.Println(v3, v25)
                            if v40 {
                                var v57 int32 = -1
                                var v58 bool = ((42 == int64(fi39)) && v1)
                                _ = v57
                                _ = v58
                            } else {
                                v59 := GenS7{F0: 2147483647, F1: v3}
                                var v60 int32 = 2147483647
                                _ = v59
                                _ = v60
                            }
                        } else {
                            for fi61 := 0; fi61 < 19; fi61++ {
                                var v62 uint16 = uint16(v3)
                                v4 = v4
                                var v63 int8 = 127
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                        }
                    default:
                        if false {
                            for fi64 := 0; fi64 < 12; fi64++ {
                                v65 := GenS3{F0: v25, F1: (v25 + 65535)}
                                var v66 int16 = 32767
                                v3 = 65535
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        }
                    }
                    _ = fi39
                    _ = v40
                }
            }
        }
        if (false && v1) {
            switch 9223372036854775807 {
            case 4294967295:
                v67 := GenS8{F0: 42, F1: 42}
                _ = v67
            case 888:
                var v68 uint32 = 4294967295
                _ = v68
            case -63:
                var v69 bool = false
                _ = v69
            default:
                var v70 uint = 235
                _ = v70
            }
        }
        fmt.Println(v4, v1, v4)
    }
    var v71 uint32 = v25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v71
}

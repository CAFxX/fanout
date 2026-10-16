package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 int32; F1 uint64; F2 string }
func genFn0(p0 uint, p1 int32, p2 uint32) int64 {
    var v6 float32 = 0.1
    if true {
        if false {
            var v7 bool = ((float32((-2.25 + 3.14159)) * (0.1 - (1.5 + float32(float64(v6))))) < 1e10)
            _ = v7
        } else {
            if ((((true && true) || (59 <= 100)) && (false || ((false && (1 <= 128)) && ((0.0 >= v6) || true)))) && (42 <= 1)) {
                v8 := GenS1{F0: 65535}
                _ = v8
            }
            v9 := GenS2{F0: 42, F1: p2}
            v10 := GenS3{F0: p0, F1: "foo bar"}
            _ = v10
            _ = v9
        }
    } else {
        v11 := GenS4{F0: 1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}
func genFn1() bool {
    var v12 uint32 = 0
    if (!true) {
        v12 -= (^(v12))
    } else {
        var v13 float64 = 1e10
        _ = v13
    }
    v14 := []int{-2147483648, -1, 42, 32768}
    for fi15 := 0; fi15 < 7; fi15++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v12 -= 2147483648
        _ = fi15
    }
    _ = v12
    _ = v14
    return (-128 <= -128)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 487
    var v3 uint = 42
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 bool = ((128 <= 127) || true)
        v2 = (v2 % ((v2) | 1))
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: (v2 * 32767)}
    for fi19 := 0; fi19 < 2; fi19++ {
        if false {
            v20 := GenS4{F0: -1}
            v21 := GenS5{F0: 128}
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS1{F0: v2}
                _ = fi22
                _ = v23
            }
            v24 := GenS5{F0: 1}
            _ = v24
        }
        if (false || (!(true && (false && false)))) {
            switch fi19 {
            case 614:
                switch "hello" {
                case "a":
                    switch fi19 {
                    case 2147483647:
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v26 := GenS6{F0: v3}
                            _ = fi25
                            _ = v26
                        }
                    default:
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := GenS2{F0: 190, F1: 0}
                            _ = fi27
                            _ = v28
                        }
                    }
                default:
                    v29 := v1[2147483647]
                    _ = v29
                }
            case -1:
                switch fi19 {
                case 2147483647:
                    var v30 int32 = 42
                    _ = v30
                case 128:
                    var v31 uint8 = 72
                    _ = v31
                case 72:
                    fmt.Println(fi19, v2, v2)
                }
            }
            if (0 != 100) {
                var v32 float64 = 0.1
                fmt.Println(v32)
                var v33 int8 = 100
                _ = v32
                _ = v33
            }
        } else {
            fmt.Println(v2, v2)
            var v34 bool = true
            _ = v34
        }
        _ = fi19
    }
    v3 = v3
    var v35 uint = v3
    fmt.Println(v2)
    var v36 int = 4294967295
    var v37 string = "hello"
    v37 = "hello"
    if false {
        delete(v1, 42)
        v37 = "a"
    } else {
        switch ((v37 + (v37 + "\n\t")) + "hello") {
        case "foo bar":
            v38 := GenS5{F0: 1}
            _ = v38
        case "hello":
            switch "foo bar" {
            case "foo bar":
                if true {
                    switch v36 {
                    case -9223372036854775808:
                        v1[42] = 2147483647
                    case 2147483647:
                        for fi39 := 0; fi39 < 14; fi39++ {
                            var v40 int16 = 32767
                            fmt.Println(fi39, v35, v40)
                            _ = fi39
                            _ = v40
                        }
                    default:
                        if false {
                            var v41 uint = (v35 | v3)
                            _ = v41
                        } else {
                            var v42 float64 = float64(float32(1e10))
                            for fi43 := 0; fi43 < 16; fi43++ {
                                for fi44 := 0; fi44 < 8; fi44++ {
                                    v45 := v1[1]
                                    if true {
                                        v42 = float64(1.5)
                                        if true {
                                            v42 += 0.1
                                            var v46 uint16 = 0
                                            delete(v1, 818)
                                            _ = v46
                                        } else {
                                            var v47 uint32 = 65535
                                            var v48 int64 = 42
                                            fmt.Println(v37, v45, v35)
                                            _ = v47
                                            _ = v48
                                        }
                                        var v49 int16 = -124
                                        _ = v49
                                    } else {
                                        var v50 bool = (v45 > 0)
                                        _ = v50
                                    }
                                    if true {
                                        v5[2147483647] = true
                                        v51 := GenS3{F0: (v35 - v35), F1: (v37 + "")}
                                        _ = v51
                                    } else {
                                        var v52 int64 = 1
                                        if (!(-32768 != -32768)) {
                                            var v53 int16 = 42
                                            v2 = v2
                                            var v54 int8 = 127
                                            _ = v53
                                            _ = v54
                                        }
                                        var v55 string = ""
                                        _ = v52
                                        _ = v55
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                fmt.Println(v37, v42, v36)
                                _ = fi43
                            }
                            _ = v42
                        }
                    }
                    var v56 uint64 = 778
                    _ = v56
                } else {
                    var v57 int8 = 100
                    v58 := v5[619]
                    if (!(114 == uint16(v35))) {
                        v36 += v36
                        v3 = (v35 | v35)
                        if false {
                            if false {
                                v59 := GenS2{F0: 127, F1: 705}
                                _ = v59
                            }
                            v60 := v4[(("" + (v37 + v37)) + v37)]
                            _ = v60
                        }
                    } else {
                        var v61 float64 = 3.14159
                        fmt.Println(v58, v36, v36)
                        _ = v61
                    }
                    _ = v57
                    _ = v58
                }
            case "a":
                var v62 int32 = 1
                _ = v62
            }
        default:
            v36 += v36
        }
    }
    v3 += (v35 | (v3 / ((v35) | 1)))
    v35 += v35
    var v63 float64 = (1.5 - (0.1 / 0.1))
    v64 := GenS2{F0: 77, F1: 202}
    v65 := GenS7{F0: (v3 ^ v35), F1: 128}
    v66 := GenS8{F0: 812, F1: 0, F2: v37}
    var v67 int16 = -1
    if false {
        v63 += ((v63 - (-2.25 * float64((3.14159 - 0.1)))) * 0.0)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}

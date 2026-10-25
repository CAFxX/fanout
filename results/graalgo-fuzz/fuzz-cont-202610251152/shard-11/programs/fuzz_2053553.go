package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint8) uint {
    v5 := GenS1{F0: -19, F1: -2147483648}
    v6 := []int{42, 0}
    var v7 uint64 = 0
    var v8 uint32 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() int32 {
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: fi9, F1: 1}
        _ = fi9
        _ = v10
    }
    if (3.14159 < (3.14159 - (3.14159 * (-2.25 / 1e10)))) {
        v11 := []int{}
        var v12 bool = (genFn0(42) != 875)
        _ = v11
        _ = v12
    } else {
        var v13 int32 = 42
        if (!(2 < 75)) {
            v13 -= v13
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int8 = 1
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int16 = -5
                v17 += -78
                v17 -= (v17 % ((v17) | 1))
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    v18 := GenS2{F0: 128, F1: -1, F2: 42}
    _ = v18
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", (("" + "hello") + "a"), ("a" + (("\n\t" + "a") + ""))}
    var v3 string = "foo bar"
    var v4 int64 = 42
    v4 = v4
    var v19 string = v3
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS3{F0: (("a" + "foo bar") + ("\n\t" + "hello")), F1: true}
        switch fi20 {
        case -195:
            v22 := GenS4{F0: 127, F1: (^(fi20)), F2: 18446744073709551615}
            _ = v22
        default:
            fmt.Println(v3)
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            _ = len(v2)
            v19 = v3
            _ = fi23
        }
        _ = fi20
        _ = v21
    }
    v19 = (((v19 + (v3 + (((v3 + "hello") + (v19 + (v3 + "\n\t"))) + v3))) + ((v3 + (v3 + "a")) + "hello")) + v19)
    fmt.Println(v19, v4)
    v4 -= (v4 ^ v4)
    v3 = (v3 + "")
    v24 := GenS1{F0: 128, F1: -1}
    if true {
        v25 := GenS5{F0: 4294967295}
        var v26 int64 = v4
        _ = v25
        _ = v26
    }
    var v27 uint16 = 322
    v3 = ("foo bar" + "hello")
    if ((!(-2147483648 <= 0)) && (0 == -1)) {
        if (v3 != (v3 + "")) {
            v4 = v4
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 uint64 = 881
                    _ = len(v1)
                    _ = fi29
                    _ = v30
                }
                var v31 int16 = 0
                _ = fi28
                _ = v31
            }
            var v32 string = ((v19 + (v3 + (v19 + "a"))) + "\n\t")
            _ = v32
        }
        v33 := GenS5{F0: 4294967295}
        _ = v33
    } else {
        if true {
            if false {
                if false {
                    v19 = v3
                    v4 += 1
                } else {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        fmt.Println(fi34, fi34, v19)
                        _ = fi34
                    }
                    switch 2147483647 {
                    case 1:
                        switch 632 {
                        case -427:
                            if false {
                                v27 = (v27 - 65535)
                            }
                        case 1:
                            var v35 int = -9223372036854775808
                            _ = v35
                        case 683:
                            fmt.Println(v3, v4, v19)
                        default:
                            var v36 uint8 = 128
                            _ = v36
                        }
                    case 32767:
                        if true {
                            var v37 int16 = 10
                            var v38 uint64 = 2
                            for fi39 := 0; fi39 < 9; fi39++ {
                                v40 := GenS2{F0: 45, F1: 0, F2: v38}
                                _ = fi39
                                _ = v40
                            }
                            _ = v37
                            _ = v38
                        }
                    case 32768:
                        for fi41 := 0; fi41 < 6; fi41++ {
                            var v42 uint64 = 42
                            switch fi41 {
                            case 32767:
                                if true {
                                    for fi43 := 0; fi43 < 10; fi43++ {
                                        _ = len(v1)
                                        v42 += v42
                                        var v44 int32 = genFn1()
                                        _ = fi43
                                        _ = v44
                                    }
                                    v45 := GenS5{F0: (v42 % ((v42) | 1))}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v45
                                } else {
                                    v46 := GenS6{F0: (v4 % ((1) | 1)), F1: 0}
                                    fmt.Println(v3, v42, v27)
                                    v27 = v27
                                    _ = v46
                                }
                            case -1:
                                switch "" {
                                case "hello":
                                    var v47 uint = 485
                                    _ = v47
                                default:
                                    var v48 int8 = 42
                                    _ = v48
                                }
                            case 398:
                                if (v27 > v27) {
                                    if true {
                                        var v49 int8 = 42
                                        _ = v49
                                    } else {
                                        if true {
                                            v27 = v27
                                        }
                                    }
                                } else {
                                    var v50 int = 128
                                    v51 := GenS7{F0: 56}
                                    _ = v50
                                    _ = v51
                                }
                            }
                            var v52 int64 = (v4 % ((int64(fi41)) | 1))
                            _ = fi41
                            _ = v42
                            _ = v52
                        }
                    default:
                        fmt.Println(v3)
                    }
                }
                var v53 bool = (((v4 % ((v4) | 1)) == 0) && false)
                _ = v53
            } else {
                v19 = v19
            }
            v54 := GenS4{F0: 255, F1: -464, F2: 1}
            _ = v54
        } else {
            delete(v1, 42)
            v55 := GenS3{F0: "a", F1: (0 <= 0)}
            var v56 int = 9223372036854775807
            _ = v55
            _ = v56
        }
        _ = len(v2)
        v4 += (-(v4))
    }
    var v57 uint32 = 891
    for fi58 := 0; fi58 < 17; fi58++ {
        fmt.Println(v57, v57)
        _ = fi58
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        v60 := GenS8{F0: (42 == (v57 / (((v57 | (^(v57)))) | 1)))}
        _ = fi59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v57
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int8 }
func genFn0() uint64 {
    if false {
        var v4 int8 = 0
        for fi5 := 0; fi5 < 20; fi5++ {
            v4 -= 100
            _ = fi5
        }
        _ = v4
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 32767}
            if (false || true) {
                v8 := GenS1{F0: (-(fi6))}
                _ = v8
            } else {
                var v9 uint16 = 2
                v10 := []int{(fi6 ^ -1), 666, (-(fi6))}
                var v11 float32 = 0.0
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 string = "\n\t"
            _ = fi6
            _ = v12
            _ = v7
        }
        var v13 string = "hello"
        if true {
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS2{F0: 42, F1: 1}
                _ = fi14
                _ = v15
            }
        } else {
            v13 = (("foo bar" + "a") + (("hello" + v13) + v13))
        }
        _ = v13
    }
    var v16 int = 1
    v16 = 0
    v17 := GenS3{F0: 803, F1: 828, F2: 0}
    _ = v16
    _ = v17
    return 2
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = -2.25
    var v3 uint32 = 1
    var v18 uint8 = 128
    v2 = 3.14159
    if true {
        v19 := GenS4{F0: 1, F1: "foo bar"}
        if false {
            var v20 uint8 = (v18 >> ((2) & 63))
            v21 := []int{4294967295}
            v1 = (v1 | (v1 / ((-1) | 1)))
            _ = v20
            _ = v21
        } else {
            v18 += v18
        }
        fmt.Println(v1, v1, v3)
        _ = v19
    } else {
        v22 := GenS5{F0: 0}
        v23 := GenS6{F0: ("" + "\n\t"), F1: int8(v18)}
        _ = v22
        _ = v23
    }
    var v24 int8 = -1
    v25 := []int{0, 359}
    if (true && true) {
        fmt.Println(v24, v1, v2)
        var v26 int16 = int16(v3)
        var v27 int16 = -1
        _ = v26
        _ = v27
    }
    if false {
        var v28 uint32 = v3
        var v29 int32 = (v1 | (v1 ^ (v1 % ((2147483647) | 1))))
        if false {
            switch 890 {
            case 684:
                var v30 float64 = v2
                _ = v30
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 float64 = -2.25
                    var v33 uint8 = 128
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
            fmt.Println(v24)
            if ((v24 | v24) == v24) {
                v34 := GenS5{F0: int64(v3)}
                if true {
                    var v35 int16 = 32767
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        switch 128 {
                        case 667:
                            v28 -= 1
                        case 758:
                            v37 := GenS4{F0: 155, F1: ("foo bar" + (("hello" + "") + "\n\t"))}
                            _ = v37
                        case -2147483648:
                            var v38 int32 = v29
                            _ = v38
                        default:
                            var v39 int = fi36
                            _ = v39
                        }
                        _ = fi36
                    }
                    switch 1 {
                    case 127:
                        var v40 bool = true
                        _ = v40
                    case 128:
                        var v41 string = "foo bar"
                        _ = v41
                    case 42:
                        v2 = ((((float64((-2.25 / 1e10)) / v2) + float64(1.5)) + 1e10) * 0.1)
                    default:
                        if (false || (128 > 634)) {
                            if (609 != 786) {
                                v2 += v2
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v3 = v3
                                    case "":
                                        v42 := GenS5{F0: 0}
                                        _ = v42
                                    default:
                                        var v43 uint8 = v18
                                        _ = v43
                                    }
                                }
                            }
                            if (true && (v2 < 0.1)) {
                                v44 := GenS7{F0: 0, F1: (^(v3))}
                                _ = v44
                            }
                            if false {
                                v24 = v24
                                if (!(!false)) {
                                    for fi45 := 0; fi45 < 14; fi45++ {
                                        var v46 int64 = 0
                                        if len(v25) > 0 {
                                            v25[0] = (fi45 * fi45)
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                    v47 := GenS8{F0: v24}
                                    if true {
                                        if true {
                                            v1 += v1
                                            var v48 float32 = float32(3.14159)
                                            _ = v48
                                        }
                                    } else {
                                        v29 = v1
                                    }
                                    _ = v47
                                }
                                var v49 float32 = 1.5
                                _ = v49
                            }
                        }
                    }
                    switch "foo bar" {
                    case "a":
                        var v50 int8 = v24
                        _ = v50
                    }
                }
                _ = v34
            } else {
                var v51 int8 = int8(v28)
                var v52 int16 = 1
                _ = v51
                _ = v52
            }
        }
        _ = v28
        _ = v29
    }
    v18 -= (^(v18))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}

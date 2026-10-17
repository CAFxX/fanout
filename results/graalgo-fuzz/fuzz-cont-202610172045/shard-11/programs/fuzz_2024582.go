package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 bool; F1 int; F2 uint }
func genFn0(p0 string) uint8 {
    v5 := []int{-2147483648, 32767, 943}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 42
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: -75}
    v9 := GenS1{F0: -232, F1: 63}
    if (!(!true)) {
        p0 = p0
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(fi11, p0, fi11)
                    p0 = ("a" + "a")
                    var v12 uint16 = 65535
                    _ = fi11
                    _ = v12
                }
                var v13 int = 9223372036854775807
                _ = v13
            }
            var v14 uint64 = 0
            _ = fi10
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 string = "foo bar"
            _ = fi15
            _ = v16
        }
    }
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return 236
}

func main() {
    var v1 uint = 0
    var v2 float32 = 1.5
    v3 := make(map[uint32]string)
    var v4 uint64 = 170
    if (true && false) {
        var v17 int = 2147483647
        var v18 float32 = 0.1
        _ = v17
        _ = v18
    }
    var v19 uint32 = 2147483648
    var v20 int8 = 127
    if (3.14159 <= float32(0.0)) {
        v19 -= (^(v19))
        v21 := []int{685}
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v19)
            if (!true) {
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(v19, fi23)
                    v24 := GenS1{F0: 1000, F1: -628}
                    if true {
                        if ((v4 | (^(v4))) != (v4 - v4)) {
                            if true {
                                if (1 <= 880) {
                                    var v25 float32 = ((0.0 - v2) / 1e10)
                                    _ = v25
                                }
                            }
                            v26 := GenS2{F0: v1, F1: (^(v1))}
                            _ = v26
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    } else {
                        v27 := GenS1{F0: int16(v4), F1: 486}
                        _ = v27
                    }
                    _ = fi23
                    _ = v24
                }
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 int = 4294967295
                    v30 := GenS3{F0: (("foo bar" + "hello") + "a")}
                    switch (("hello" + "") + "a") {
                    case "foo bar":
                        v29 -= (^(v29))
                    default:
                        var v31 int = -125
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                if (true || (((("foo bar" + ("hello" + "a")) + "foo bar") > ("hello" + "hello")) || ((false && (-620 >= -32768)) || true))) {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        var v33 int = (fi32 ^ fi32)
                        var v34 float32 = (0.0 - v2)
                        if true {
                            switch (("\n\t" + "a") + ("\n\t" + "a")) {
                            case "\n\t":
                                v19 = (v19 + v19)
                            default:
                                fmt.Println(v1, fi32)
                            }
                            var v35 int32 = 1
                            if ((^(v35)) <= v35) {
                                switch 2147483647 {
                                case 312:
                                    var v36 bool = false
                                    _ = v36
                                case -871:
                                    if false {
                                        if ((v4 << ((v1) & 63)) != v4) {
                                            if (255 >= 744) {
                                                var v37 bool = false
                                                v38 := GenS4{F0: (v1 & (v1 % (((v1 & v1)) | 1))), F1: "a"}
                                                _ = v37
                                                _ = v38
                                            }
                                            v34 -= v34
                                            v39 := GenS5{F0: false, F1: v4}
                                            _ = v39
                                        }
                                    } else {
                                        v19 = (v19 / ((v19) | 1))
                                    }
                                case 128:
                                    if true {
                                        if false {
                                            v40 := GenS1{F0: 0, F1: v35}
                                            v41 := GenS6{F0: v4, F1: (^(v20))}
                                            _ = v40
                                            _ = v41
                                        } else {
                                            switch fi22 {
                                            case 482:
                                                v2 = 3.14159
                                            case 794:
                                                v4 = 4294967295
                                            case 128:
                                                var v42 uint = (^(v1))
                                                _ = v42
                                            }
                                            var v43 float64 = 1e10
                                            if false {
                                                v44 := GenS7{F0: (v33 | (fi32 & fi32)), F1: -826}
                                                delete(v3, (v19 + (v19 % (((v19 & 323)) | 1))))
                                                _ = v44
                                            } else {
                                                v4 -= (^(v4))
                                                var v45 int = v33
                                                var v46 uint = (v1 * (^(v1)))
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v43
                                        }
                                        var v47 bool = (true || (int32(v4) == (^(v35))))
                                        var v48 string = "hello"
                                        v49 := v3[v19]
                                        var v50 uint = (^(v1))
                                        var v51 int64 = -691
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 int8 = (-(v20))
                                        switch "\n\t" {
                                        case "foo bar":
                                            if false {
                                                v35 -= (v35 % ((v35) | 1))
                                                var v53 bool = (!(32767 > 0))
                                                _ = v53
                                            } else {
                                                v34 += v2
                                            }
                                        }
                                        _ = v52
                                    }
                                default:
                                    fmt.Println(fi22, fi22, v19)
                                }
                                fmt.Println(v34, v19)
                            } else {
                                _ = len(v3)
                                delete(v3, v19)
                                var v54 uint64 = (v4 | v4)
                                _ = v54
                            }
                            _ = v35
                        }
                        switch "" {
                        case "":
                            if (!(0 >= 0)) {
                                var v55 uint32 = v19
                                v33 = fi22
                                v1 = v1
                                _ = v55
                            } else {
                                var v56 uint = 255
                                v56 = (v56 ^ 65535)
                                _ = v56
                            }
                        default:
                            if (fi22 == (fi22 | v33)) {
                                if true {
                                    var v57 int8 = v20
                                    _ = v57
                                }
                                _ = len(v21)
                                fmt.Println(v34, v19)
                                fmt.Println(v19)
                                v58 := v3[(v19 - (v19 % ((42) | 1)))]
                                _ = v58
                            }
                        }
                        v33 += (fi32 + (-(fi32)))
                        var v59 uint8 = 127
                        fmt.Println(v34)
                        v20 = 42
                        var v60 uint64 = v4
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 19; fi61++ {
                        var v62 int = -993
                        v63 := GenS8{F0: (-808 != 42), F1: int(v4), F2: (v1 ^ v1)}
                        v20 += v20
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    delete(v3, v19)
                } else {
                    if (!true) {
                        switch fi22 {
                        case 42:
                            if (2147483647 != 26) {
                                _ = len(v3)
                                v19 += v19
                            }
                        }
                    }
                    for fi64 := 0; fi64 < 3; fi64++ {
                        fmt.Println(fi64)
                        fmt.Println(fi22, v4)
                        v4 -= v4
                        var v65 int32 = 42
                        v65 = (v65 << ((646) & 63))
                        _ = fi64
                        _ = v65
                    }
                    var v66 string = "foo bar"
                    _ = v66
                }
            } else {
                var v67 float32 = v2
                _ = v67
            }
            switch ("hello" + "hello") {
            case "a":
                if true {
                    v68 := v3[(^(v19))]
                    _ = v68
                } else {
                    var v69 float32 = ((v2 + float32((0.0 * 1e10))) - 0.0)
                    v69 -= (v69 / float32(-2.25))
                    _ = v69
                }
            default:
                v2 = v2
            }
            _ = fi22
        }
        _ = v21
    } else {
        v70 := GenS1{F0: -32768, F1: 1}
        _ = v70
    }
    fmt.Println(v19)
    v71 := GenS4{F0: 18446744073709551615, F1: "\n\t"}
    v20 += v20
    switch -964 {
    case -588:
        for fi72 := 0; fi72 < 19; fi72++ {
            switch ("" + "foo bar") {
            case "a":
                v1 += v1
            default:
                fmt.Println(v1, v2)
            }
            if ((!false) || (true && true)) {
                v2 += (v2 - 0.0)
            }
            fmt.Println(v2, v2, fi72)
            _ = fi72
        }
    case 0:
        _ = len(v3)
    case 174:
        switch "a" {
        case "":
            v73 := v3[v19]
            _ = v73
        }
    }
    v4 = v4
    var v74 bool = (255 == 1)
    var v75 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v71
    _ = v74
    _ = v75
}

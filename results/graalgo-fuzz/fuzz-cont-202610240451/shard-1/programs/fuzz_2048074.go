package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint32; F2 int }
type GenS5 struct { F0 int64; F1 bool; F2 int16 }
type GenS6 struct { F0 string; F1 int8; F2 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int64 }
func genFn0(p0 int16) bool {
    v6 := GenS1{F0: 2, F1: 146}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint8 = 255
            v10 := GenS2{F0: "a", F1: v9, F2: int32(v9)}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if false {
        var v11 float32 = 0.0
        v12 := []int{502}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v6
    return (!true)
}

func main() {
    var v1 int64 = -206
    v2 := make(map[int]int)
    var v3 float64 = 0.0
    v4 := make(map[int]int)
    v5 := []uint32{4294967295}
    v13 := GenS3{F0: 127}
    switch 927 {
    case -333:
        for fi14 := 0; fi14 < 12; fi14++ {
            v1 += (v1 | v1)
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int32 = 0
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v17 := GenS3{F0: 127}
    v5 = append(v5, 517)
    var v18 int = -2147483648
    switch (-(v18)) {
    case 32767:
        v19 := GenS2{F0: "a", F1: 2, F2: -1}
        _ = v19
    case 835:
        switch "foo bar" {
        case "":
            var v20 int64 = (v1 * v1)
            _ = v20
        case "a":
            v1 -= (v1 & -17)
        default:
            v21 := GenS2{F0: "\n\t", F1: 239, F2: -1}
            _ = v21
        }
    case 868:
        fmt.Println(v3, v18, v3)
    default:
        fmt.Println(v18, v1)
    }
    if false {
        for fi22 := 0; fi22 < 10; fi22++ {
            v3 = ((v3 * (float64((0.1 / -2.25)) - float64(3.14159))) + float64(float32(float64(3.14159))))
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 uint32 = 395
                v25 := GenS4{F0: v18, F1: v24, F2: v18}
                v1 = (v1 / ((v1) | 1))
                v26 := GenS3{F0: -1}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
        switch "" {
        case "a":
            if false {
                var v27 int32 = -1
                _ = v27
            } else {
                if true {
                    var v28 int64 = v1
                    var v29 int16 = 32767
                    _ = v28
                    _ = v29
                } else {
                    if true {
                        var v30 uint = 42
                        var v31 string = "\n\t"
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int = v18
                        v3 += v3
                        var v33 uint = 323
                        _ = v32
                        _ = v33
                    }
                }
                var v34 int16 = 0
                _ = v34
            }
        case "":
            v1 += v1
        }
        var v35 int = v18
        _ = v35
    } else {
        v18 = (-(v18))
        if true {
            var v36 float32 = (-2.25 * (1.5 * 0.1))
            if (!true) {
                v37 := GenS2{F0: (("a" + ("a" + ("hello" + ""))) + "foo bar"), F1: 127, F2: 0}
                switch ("a" + "foo bar") {
                case "hello":
                    var v38 int8 = 127
                    _ = v38
                default:
                    for fi39 := 0; fi39 < 20; fi39++ {
                        var v40 int8 = 100
                        v41 := GenS4{F0: (fi39 / (((fi39 & 127)) | 1)), F1: 2, F2: (fi39 & fi39)}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
                switch "\n\t" {
                case "hello":
                    for fi42 := 0; fi42 < 9; fi42++ {
                        if (-39 < -117) {
                            if (944 < -2147483648) {
                                var v43 int64 = v1
                                v43 = 9223372036854775807
                                _ = v43
                            } else {
                                var v44 float32 = (v36 * ((v36 + (((v36 - (v36 + (v36 + (float32(v3) * 1.5)))) + (v36 * (((0.1 + float32((v3 * 3.14159))) * v36) / (v36 + 3.14159)))) - float32((v3 + v3)))) * v36))
                                var v45 float32 = (v36 + (0.1 / (v44 / v36)))
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    v47 := GenS5{F0: (-(v1)), F1: (4294967295 >= 2), F2: 42}
                                    v44 -= 0.0
                                    if (2 == 42) {
                                        fmt.Println(v44)
                                        switch (v18 ^ v18) {
                                        case -516:
                                            fmt.Println(v45)
                                        case 1:
                                            fmt.Println(v1, v36, fi46)
                                        case 719:
                                            fmt.Println(v1)
                                        default:
                                            _ = len(v2)
                                        }
                                        v48 := GenS3{F0: 100}
                                        fmt.Println(v1, fi42)
                                        _ = v48
                                    }
                                    v36 = v45
                                    _ = fi46
                                    _ = v47
                                }
                                _ = v44
                                _ = v45
                            }
                            if false {
                                if (42 > 1) {
                                    v1 = (v1 | (^(v1)))
                                } else {
                                    switch ("hello" + ("\n\t" + ("\n\t" + ("hello" + "\n\t")))) {
                                    case "hello":
                                        v49 := v2[128]
                                        _ = v49
                                    }
                                    for fi50 := 0; fi50 < 20; fi50++ {
                                        switch ("" + "") {
                                        case "a":
                                            v51 := GenS2{F0: "hello", F1: 0, F2: 2147483647}
                                            _ = v51
                                        case "foo bar":
                                            var v52 bool = true
                                            _ = v52
                                        default:
                                            delete(v2, v18)
                                        }
                                        _ = fi50
                                    }
                                }
                                var v53 bool = ((fi42 < (-(v18))) || false)
                                var v54 int32 = -919
                                _ = v53
                                _ = v54
                            }
                        } else {
                            var v55 uint16 = 987
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v56 uint8 = 127
                            _ = v55
                            _ = v56
                        }
                        _ = fi42
                    }
                default:
                    var v57 int8 = -1
                    _ = v57
                }
                _ = v37
            } else {
                fmt.Println(v3, v18, v18)
                var v58 float64 = ((v3 + -2.25) - 3.14159)
                _ = v58
            }
            for fi59 := 0; fi59 < 20; fi59++ {
                v60 := GenS6{F0: "hello", F1: -75, F2: ""}
                _ = fi59
                _ = v60
            }
            _ = v36
        }
        if (-70 != 42) {
            v1 = v1
            switch v18 {
            case 2147483647:
                if (!true) {
                    v61 := GenS7{F0: 65535}
                    _ = v61
                } else {
                    switch ("hello" + "hello") {
                    case "":
                        var v62 int32 = 1
                        _ = v62
                    }
                    v3 = v3
                    var v63 int64 = (v1 | (v1 + v1))
                    _ = v63
                }
            case -1:
                var v64 uint32 = 740
                _ = v64
            case 0:
                v3 += (1e10 / 1.0)
            }
        } else {
            v65 := GenS3{F0: 23}
            _ = v65
        }
    }
    var v66 int16 = int16(v18)
    v1 = (v1 ^ -191)
    v67 := GenS3{F0: 127}
    if (("" > "") && true) {
        if true {
            v68 := GenS8{F0: (v66 + v66)}
            for fi69 := 0; fi69 < 14; fi69++ {
                v66 = v66
                v4[fi69] = 9223372036854775807
                _ = fi69
            }
            v1 += v1
            _ = v68
        } else {
            var v70 int64 = -881
            v71 := GenS9{F0: (v66 / ((32767) | 1)), F1: (v1 | 42)}
            _ = v70
            _ = v71
        }
        v3 = 1e10
        var v72 uint32 = 206
        _ = v72
    } else {
        _ = len(v2)
        v73 := GenS7{F0: 65535}
        fmt.Println(v1)
        _ = v73
    }
    var v74 float64 = v3
    var v75 string = "a"
    v76 := GenS9{F0: v66, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v74
    _ = v75
    _ = v76
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
func genFn0(p0 int16, p1 uint, p2 int) int64 {
    p2 -= (p2 & 42)
    v6 := GenS1{F0: 42}
    var v7 float32 = (0.0 * (-2.25 / (0.1 * -2.25)))
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 float64 = 1e10
            v11 := []int{}
            var v12 int = -477
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        if false {
            v13 := GenS1{F0: 0}
            _ = v13
        } else {
            v14 := []int{}
            v15 := GenS1{F0: 32767}
            _ = v14
            _ = v15
        }
        v16 := []int{(fi8 + (fi8 % ((p2) | 1))), fi8, p2}
        _ = fi8
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 97
    var v3 uint64 = 896
    var v4 int8 = (v2 / (((v2 % ((v2) | 1))) | 1))
    v5 := make(map[int]string)
    var v17 int64 = -909
    v18 := []int{260, 32767, 1}
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int16 = -32768
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                if len(v18) > 0 {
                    v18[0] = (fi21 / ((128) | 1))
                }
                v22 := GenS2{F0: 537}
                switch (("hello" + "") + "hello") {
                case "":
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                default:
                    var v23 uint64 = v3
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            var v24 uint64 = v3
            _ = v24
        } else {
            v3 += v3
            var v25 uint8 = 2
            v25 -= v25
            _ = v25
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = fi19
        _ = v20
    }
    delete(v5, -793)
    for fi26 := 0; fi26 < 4; fi26++ {
        fmt.Println(v17, fi26, v3)
        for fi27 := 0; fi27 < 16; fi27++ {
            fmt.Println(v3)
            v4 -= (-(v2))
            switch "hello" {
            case "":
                if false {
                    v17 += (v17 | v17)
                    v4 -= v2
                    v2 = (v4 - (v4 / (((-(v2))) | 1)))
                } else {
                    var v28 string = "hello"
                    switch (fi27 + -568) {
                    case 0:
                        v29 := GenS3{F0: (4294967295 <= 1), F1: 68}
                        _ = v29
                    default:
                        var v30 uint64 = v3
                        _ = v30
                    }
                    if false {
                        switch ("foo bar" + v28) {
                        case "foo bar":
                            if true {
                                if (3.14159 != (((0.0 * ((1.5 - 0.1) + (((float32(0.0) + (float32((((float64((0.0 - (1.5 * 3.14159))) - float64(-2.25)) + 1e10) * 0.1)) / 3.14159)) * 0.1) + (-2.25 / 0.1)))) / (-2.25 + 0.1)) + (1.5 / 1.5))) {
                                    fmt.Println(v17, fi26, fi27)
                                    v3 = (v3 >> ((18446744073709551615) & 63))
                                } else {
                                    var v31 uint64 = 2
                                    _ = v31
                                }
                            } else {
                                var v32 int32 = 0
                                delete(v5, (fi27 | fi27))
                                _ = v32
                            }
                        case "\n\t":
                            delete(v1, v28)
                        }
                        var v33 uint32 = 0
                        switch ("a" + ((v28 + v28) + (v28 + ((v28 + (v28 + v28)) + v28)))) {
                        case "hello":
                            v4 -= (^(v4))
                        case "\n\t":
                            if (false && (v28 >= "foo bar")) {
                                fmt.Println(v2)
                                if ((false && true) || true) {
                                    switch fi27 {
                                    case 42:
                                        v34 := v5[(-(fi26))]
                                        _ = v34
                                    case 0:
                                        fmt.Println(fi27, v33, fi27)
                                    default:
                                        v28 = "\n\t"
                                        fmt.Println(fi27, fi27, fi27)
                                        v3 = v3
                                    }
                                    _ = len(v18)
                                    var v35 int = fi26
                                    _ = v35
                                }
                            }
                        default:
                            fmt.Println(fi27, v33, v28)
                        }
                        v4 = (v2 / ((v2) | 1))
                        var v36 int64 = 544
                        _ = v33
                        _ = v36
                    } else {
                        var v37 uint64 = (v3 & 42)
                        switch v28 {
                        case "\n\t":
                            delete(v5, 9223372036854775807)
                        default:
                            var v38 float32 = (((0.0 * float32((float64(0.1) - -2.25))) / 1e10) - 1.5)
                            _ = v38
                        }
                        v4 -= (v4 % ((67) | 1))
                        _ = v37
                    }
                    _ = v28
                }
                v1["foo bar"] = 949
            case "\n\t":
                v39 := GenS4{F0: 865, F1: (v4 * v2)}
                _ = v39
            }
            _ = fi27
        }
        _ = fi26
    }
    fmt.Println(v2)
    var v40 int16 = 1
    v41 := GenS5{F0: "\n\t", F1: -481}
    v42 := GenS6{F0: (v2 % ((int8(v17)) | 1))}
    fmt.Println(v4, v40)
    v5[505] = "foo bar"
    fmt.Println(v4)
    if false {
        v3 = (v3 ^ 18446744073709551615)
    } else {
        if false {
            for fi43 := 0; fi43 < 1; fi43++ {
                var v44 bool = (true || (0 < v2))
                delete(v5, (fi43 % ((fi43) | 1)))
                v3 += v3
                _ = fi43
                _ = v44
            }
        }
        v3 = (v3 << ((0) & 63))
    }
    switch ("a" + "") {
    case "a":
        var v45 uint16 = 65535
        _ = v45
    }
    var v46 int8 = (v2 % ((-1) | 1))
    var v47 uint8 = 2
    switch -9223372036854775808 {
    case 128:
        for fi48 := 0; fi48 < 4; fi48++ {
            v17 = genFn0(v40, 929, 127)
            for fi49 := 0; fi49 < 11; fi49++ {
                v17 = v17
                var v50 uint8 = 255
                v51 := GenS5{F0: (("a" + "a") + "foo bar"), F1: -2147483648}
                fmt.Println(v4, v4)
                if (0.1 <= 0.1) {
                    if ((float64(float32(0.0)) + 1e10) == 0.0) {
                        if true {
                            v18 = append(v18, fi49)
                            var v52 int = fi49
                            _ = v52
                        }
                        v53 := GenS4{F0: (v3 / ((2) | 1)), F1: -1}
                        var v54 uint8 = 0
                        _ = v53
                        _ = v54
                    } else {
                        v55 := GenS7{F0: (-(fi48))}
                        var v56 int = 127
                        _ = v55
                        _ = v56
                    }
                }
                fmt.Println(v4, v2, v46)
                switch fi48 {
                case 660:
                    if ((v50 + v47) < v50) {
                        v50 -= v50
                        if (v50 == (v47 - 65)) {
                            var v57 float64 = float64(((float32(0.0) * (0.0 * 0.1)) * float32(((((float64(1.5) / -2.25) * float64(float32(3.14159))) + 3.14159) / -2.25))))
                            _ = v57
                        }
                        fmt.Println(fi48)
                    } else {
                        v47 -= 127
                        v58 := GenS7{F0: int(v3)}
                        _ = v58
                    }
                case 478:
                    v50 = v47
                default:
                    v18 = append(v18, fi49)
                }
                _ = fi49
                _ = v50
                _ = v51
            }
            var v59 uint64 = v3
            _ = fi48
            _ = v59
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v5
}

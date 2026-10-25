package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int; F2 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32) int16 {
    v5 := GenS1{F0: 832, F1: -10}
    var v6 int16 = 733
    var v7 string = ("a" + ("\n\t" + ""))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int8, p1 int) int8 {
    v8 := []int{p1, 128, (p1 >> ((255) & 63)), (p1 - (p1 + p1))}
    v9 := GenS2{F0: 641, F1: 32767}
    var v10 float64 = (float64(0.1) - ((1.5 / 1e10) * -2.25))
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]bool)
    var v3 float64 = 1e10
    v4 := make(map[string]bool)
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int8 = (v1 << ((18446744073709551615) & 63))
            fmt.Println(v1, v3)
            _ = fi12
            _ = v13
        }
        switch -303 {
        case 889:
            for fi14 := 0; fi14 < 15; fi14++ {
                if false {
                    switch fi14 {
                    case -726:
                        v3 -= float64(((-2.25 * (1.5 * ((((0.0 - (3.14159 * 1e10)) / 3.14159) - -2.25) + (1.5 + 3.14159)))) * (1.5 + (0.0 - 0.1))))
                    case -1:
                        switch fi11 {
                        case 32767:
                            v15 := GenS2{F0: 0, F1: 59}
                            _ = v15
                        case 128:
                            if false {
                                var v16 uint8 = 127
                                _ = v16
                            } else {
                                var v17 uint8 = 128
                                if (true && true) {
                                    v3 = v3
                                    delete(v4, "\n\t")
                                    var v18 int8 = v1
                                    _ = v18
                                } else {
                                    switch (fi14 ^ -9223372036854775808) {
                                    case 128:
                                        var v19 bool = (!("hello" <= "\n\t"))
                                        _ = v19
                                    default:
                                        var v20 int32 = 730
                                        var v21 bool = false
                                        v17 += (v17 % ((v17) | 1))
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 string = "a"
                                    if false {
                                        switch fi14 {
                                        case 1:
                                            v17 = 127
                                        default:
                                            v23 := []int{0, (fi14 & fi14), fi14, 128}
                                            _ = v23
                                        }
                                    } else {
                                        switch 118 {
                                        case 9223372036854775807:
                                            v17 -= v17
                                        case 127:
                                            v22 = "foo bar"
                                        case 4294967295:
                                            v17 = (v17 >> ((1) & 63))
                                        default:
                                            var v24 uint32 = 1
                                            _ = v24
                                        }
                                    }
                                    var v25 string = "foo bar"
                                    v26 := GenS1{F0: 2, F1: int8(fi14)}
                                    var v27 uint = 1
                                    _ = v22
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v17
                            }
                        default:
                            fmt.Println(v3, v1, fi14)
                        }
                    case 2147483647:
                        var v28 int64 = 9223372036854775807
                        _ = v28
                    }
                } else {
                    var v29 int64 = 1
                    v30 := GenS3{F0: -145, F1: (fi11 + (fi11 - 13))}
                    _ = v29
                    _ = v30
                }
                fmt.Println(v3, fi11)
                v3 = float64(0.1)
                _ = fi14
            }
        default:
            var v31 uint = 145
            _ = v31
        }
        _ = fi11
    }
    fmt.Println(v3, v3)
    v3 = 0.1
    v32 := GenS4{F0: ""}
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int64 = 9223372036854775807
        var v35 float32 = 1.5
        v34 = (v34 << ((0) & 63))
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v1 += (v1 >> ((307) & 63))
        _ = fi36
    }
    fmt.Println(v1, v1)
    switch "hello" {
    case "\n\t":
        switch -57 {
        case 32768:
            v3 = 3.14159
        case 4294967295:
            v37 := GenS5{F0: 4294967295, F1: 1, F2: 2147483647}
            _ = v37
        case 9223372036854775807:
            if ("hello" < "\n\t") {
                switch -1 {
                case -9223372036854775808:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        var v39 float64 = ((v3 / 0.1) * (float64(0.0) + ((v3 / float64(float32(v3))) + float64(1.5))))
                        v40 := GenS4{F0: ""}
                        var v41 int = -746
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                case 9223372036854775807:
                    v42 := []int{-13, 2147483647, 9223372036854775807}
                    _ = v42
                default:
                    v1 -= (v1 * genFn1((v1 >> ((18446744073709551615) & 63)), -9223372036854775808))
                }
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := []int{}
                    if (((v3 > v3) && (false || (65535 >= 0))) && false) {
                        var v45 uint8 = 2
                        _ = v45
                    }
                    if true {
                        var v46 float64 = 3.14159
                        v3 -= v3
                        fmt.Println(v46, v46, fi43)
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                for fi47 := 0; fi47 < 14; fi47++ {
                    var v48 int16 = 42
                    var v49 float64 = 0.0
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            } else {
                var v50 int16 = -1
                v2[814] = (734 > 242)
                var v51 uint32 = 2
                _ = v50
                _ = v51
            }
        default:
            var v52 uint8 = 2
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 2; fi53++ {
            for fi54 := 0; fi54 < 1; fi54++ {
                v55 := GenS5{F0: 1, F1: fi53, F2: 588}
                v56 := v4[("foo bar" + ("foo bar" + "foo bar"))]
                v57 := GenS5{F0: 4294967295, F1: (-(fi53)), F2: 887}
                v3 -= v3
                v58 := GenS3{F0: genFn0(-1), F1: 42}
                _ = fi54
                _ = v55
                _ = v56
                _ = v57
                _ = v58
            }
            v59 := GenS6{F0: (v1 + (v1 / (((v1 | (v1 - (v1 * v1)))) | 1)))}
            _ = fi53
            _ = v59
        }
    }
    v1 = (v1 ^ 1)
    v3 = 0.1
    v1 -= genFn1((v1 / ((genFn1(v1, 4294967295)) | 1)), -2147483648)
    v60 := GenS2{F0: -1, F1: 545}
    var v61 uint64 = 374
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v60
    _ = v61
}

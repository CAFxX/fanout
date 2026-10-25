package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0(p0 int32, p1 int32, p2 uint) int16 {
    p0 += -1
    p1 -= -2147483648
    v7 := []int{9223372036854775807, 42, 460}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 42
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]int)
    var v3 uint64 = 615
    var v4 float64 = ((1e10 / 0.1) + 1.5)
    var v5 uint16 = 113
    var v6 string = "foo bar"
    var v10 uint16 = (v5 - 255)
    var v11 int32 = 679
    delete(v2, v6)
    v5 -= (^(v5))
    if (false || (true || true)) {
        if (v4 != ((((1e10 / float64(-2.25)) + (1.5 - v4)) * ((-2.25 / v4) / 0.1)) / (v4 - v4))) {
            fmt.Println(v4, v5)
            v12 := GenS1{F0: ("foo bar" + v6)}
            _ = v12
        } else {
            v6 = v6
        }
    }
    var v13 uint32 = 42
    var v14 uint = 4294967295
    v6 = "\n\t"
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(v5, v14, fi15)
        _ = fi15
    }
    v1 = (v1 * (v1 * v1))
    v4 = (v4 / (((v4 - (-2.25 * v4)) - float64((1e10 + (0.1 + (-2.25 + 1e10))))) - v4))
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := v2[(v6 + "hello")]
        v18 := []int{int(v3)}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint = (v14 << ((v14) & 63))
    if ((^(v5)) < v5) {
        var v20 bool = (v11 > -230)
        for fi21 := 0; fi21 < 13; fi21++ {
            if v20 {
                fmt.Println(v5)
            } else {
                switch v6 {
                case "\n\t":
                    switch fi21 {
                    case -319:
                        var v22 uint16 = 952
                        _ = v22
                    case 266:
                        fmt.Println(v3, v5)
                    case 564:
                        if v20 {
                            var v23 uint8 = 127
                            for fi24 := 0; fi24 < 18; fi24++ {
                                v25 := []int{(v1 | 9223372036854775807), -9223372036854775808}
                                if (!v20) {
                                    v26 := GenS2{F0: fi24, F1: v6, F2: 65535}
                                    _ = v26
                                } else {
                                    if ((v5 ^ 42) >= 936) {
                                        switch v6 {
                                        case "":
                                            delete(v2, v6)
                                        case "a":
                                            var v27 int8 = 0
                                            _ = v27
                                        }
                                        var v28 uint64 = 2
                                        v29 := GenS3{F0: 18446744073709551615, F1: v13, F2: -2147483648}
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int32 = (v11 / (((-(v11))) | 1))
                                    fmt.Println(v3)
                                    _ = v30
                                }
                                _ = fi24
                                _ = v25
                            }
                            fmt.Println(v23, v23, v10)
                            _ = v23
                        } else {
                            fmt.Println(v19, v11)
                        }
                    }
                case "a":
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v19 = (v19 << (((v14 + (^(v19)))) & 63))
                        _ = fi31
                    }
                case "":
                    switch "" {
                    case "hello":
                        if ((v10 > v10) || ((v11 * 0) < 42)) {
                            v32 := v2[("" + "hello")]
                            delete(v2, v6)
                            switch v6 {
                            case "foo bar":
                                v3 -= v3
                            case "a":
                                var v33 uint64 = uint64(v5)
                                _ = v33
                            }
                            _ = v32
                        }
                    case "\n\t":
                        v34 := GenS2{F0: int(v13), F1: v6, F2: (v5 / ((uint16(v19)) | 1))}
                        _ = v34
                    }
                default:
                    v3 += v3
                }
                var v35 uint32 = (v13 % (((v13 + (v13 * (^(v13))))) | 1))
                v13 = (v13 / (((v35 | 42)) | 1))
                _ = v35
            }
            var v36 int8 = 1
            if v20 {
                var v37 uint32 = (v13 | (^(v13)))
                v38 := GenS4{F0: (v10 << ((18446744073709551615) & 63)), F1: 18446744073709551615, F2: v36}
                _ = v37
                _ = v38
            } else {
                switch ("hello" + ("\n\t" + v6)) {
                case "":
                    v39 := GenS3{F0: (^(v3)), F1: (^(v13)), F2: (v11 % (((v11 / ((v11) | 1))) | 1))}
                    _ = v39
                case "a":
                    for fi40 := 0; fi40 < 9; fi40++ {
                        v41 := []int{}
                        _ = fi40
                        _ = v41
                    }
                default:
                    if (371 == 1) {
                        switch v1 {
                        case -181:
                            v42 := []int{}
                            _ = v42
                        case 128:
                            for fi43 := 0; fi43 < 6; fi43++ {
                                if (("\n\t" + ((v6 + (v6 + ("\n\t" + (v6 + v6)))) + "a")) != v6) {
                                    v44 := GenS5{F0: false, F1: genFn0(-206, (v11 / (((-(v11))) | 1)), (v14 / (((v19 | 1)) | 1))), F2: (v13 ^ v13)}
                                    var v45 bool = (v20 && (v14 != (v14 & v14)))
                                    v46 := []int{}
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    switch 2147483647 {
                                    case 9223372036854775807:
                                        v1 -= 2147483647
                                    default:
                                        fmt.Println(v3, v5)
                                    }
                                    v47 := []int{(fi43 & (fi21 * 32767)), (fi21 ^ v1)}
                                    _ = v47
                                }
                                v6 = (v6 + v6)
                                var v48 uint64 = (v3 & (v3 << ((uint(v3)) & 63)))
                                _ = fi43
                                _ = v48
                            }
                        case -322:
                            v2[v6] = v1
                        }
                    }
                }
                v13 += (v13 >> (((v19 / (((^(v19))) | 1))) & 63))
                switch (fi21 % (((fi21 ^ v1)) | 1)) {
                case 9223372036854775807:
                    switch (^(fi21)) {
                    case 32767:
                        if v20 {
                            var v49 uint = (v14 ^ v14)
                            var v50 int64 = 0
                            _ = v49
                            _ = v50
                        }
                    case -9223372036854775808:
                        for fi51 := 0; fi51 < 4; fi51++ {
                            v11 = v11
                            var v52 uint = v19
                            _ = fi51
                            _ = v52
                        }
                    }
                case -1:
                    fmt.Println(v19, v11)
                case -130:
                    var v53 float32 = (1e10 * float32(float64(0.1)))
                    _ = v53
                default:
                    for fi54 := 0; fi54 < 17; fi54++ {
                        v10 = v10
                        v55 := GenS1{F0: (v6 + v6)}
                        _ = fi54
                        _ = v55
                    }
                }
            }
            _ = fi21
            _ = v36
        }
        var v56 int = (v1 % ((v1) | 1))
        _ = v20
        _ = v56
    } else {
        v2[((v6 + "hello") + "hello")] = v1
        v57 := []int{}
        _ = v57
    }
    if true {
        var v58 bool = (false || true)
        v6 = (("foo bar" + (v6 + v6)) + "a")
        var v59 int32 = v11
        _ = v58
        _ = v59
    } else {
        if false {
            v60 := GenS6{F0: (v14 | v19), F1: v13}
            fmt.Println(v3)
            _ = v60
        } else {
            v2["\n\t"] = (v1 - v1)
            for fi61 := 0; fi61 < 11; fi61++ {
                if false {
                    switch (v6 + v6) {
                    case "hello":
                        _ = len(v2)
                    default:
                        for fi62 := 0; fi62 < 18; fi62++ {
                            if false {
                                switch 427 {
                                case 9223372036854775807:
                                    v19 += (v19 * v19)
                                case 32768:
                                    delete(v2, "a")
                                default:
                                    var v63 uint32 = v13
                                    var v64 string = ((v6 + "") + "\n\t")
                                    _ = v63
                                    _ = v64
                                }
                                v5 += (v10 % ((v10) | 1))
                                var v65 uint = (v14 / (((v19 - (v14 | (v14 << ((v14) & 63))))) | 1))
                                _ = v65
                            }
                            v66 := []int{(fi61 + fi61), (v1 << ((v19) & 63)), (fi61 - -197), (fi62 & fi61)}
                            var v67 uint = (v19 % (((v14 + v14)) | 1))
                            _ = fi62
                            _ = v66
                            _ = v67
                        }
                    }
                }
                v2[("hello" + "a")] = 2147483647
                _ = fi61
            }
        }
        v68 := GenS6{F0: v19, F1: v13}
        v69 := []int{(v1 >> ((0) & 63)), 431, v1}
        _ = v68
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

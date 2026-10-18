package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 string; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
func genFn0(p0 bool, p1 bool) int32 {
    var v7 int = 1
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 6; fi8++ {
            p0 = (493 >= 65535)
            if (p1 || ((1.5 <= float32(1e10)) && p1)) {
                for fi9 := 0; fi9 < 13; fi9++ {
                    p1 = false
                    _ = fi9
                }
                fmt.Println(p0, p1, v7)
                if true {
                    v10 := []int{}
                    for fi11 := 0; fi11 < 19; fi11++ {
                        p1 = (false || (p0 && false))
                        _ = fi11
                    }
                    _ = v10
                } else {
                    var v12 uint16 = 2
                    v13 := GenS1{F0: 2, F1: 2, F2: 1000}
                    if true {
                        fmt.Println(v7)
                        v14 := GenS1{F0: 1, F1: 2147483648, F2: 529}
                        v15 := []int{(fi8 + fi8), fi8, 0}
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint32 = 645
                        v17 := []int{}
                        var v18 uint = 255
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v12
                    _ = v13
                }
            }
            _ = fi8
        }
    case "hello":
        fmt.Println(p1, p0, p0)
    case "foo bar":
        if (int8(v7) >= 100) {
            v7 = 9223372036854775807
        }
    }
    v19 := []int{v7, v7}
    _ = p0
    _ = p1
    _ = v19
    _ = v7
    return -1
}

func main() {
    var v1 uint32 = 27
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 2
    var v4 uint32 = (v1 + v1)
    v5 := []uint32{v1, 2147483648, (^(v1))}
    v6 := []uint32{v4, (v1 << ((18446744073709551615) & 63))}
    for fi20 := 0; fi20 < 17; fi20++ {
        v1 += v4
        fmt.Println(fi20, v4)
        if (false || true) {
            v21 := GenS1{F0: 2147483648, F1: 4294967295, F2: -32768}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 17; fi22++ {
                if (false && (true && (-9223372036854775808 > 9223372036854775807))) {
                    v23 := v2[(v1 << ((18446744073709551615) & 63))]
                    _ = v23
                } else {
                    var v24 int32 = -1
                    var v25 uint32 = v1
                    var v26 uint16 = 1
                    _ = v24
                    _ = v25
                    _ = v26
                }
                switch "a" {
                case "foo bar":
                    if false {
                        v3 = v3
                        switch "\n\t" {
                        case "foo bar":
                            var v27 string = ("foo bar" + "hello")
                            _ = v27
                        }
                        if ((((float64(1.5) == (0.1 * (float64(0.0) * -2.25))) && ((fi22 != fi20) && (0.1 < 3.14159))) || (-2.25 == ((3.14159 * 3.14159) / 1.0))) || true) {
                            v3 = (v3 / ((v3) | 1))
                            fmt.Println(v4)
                            switch 32767 {
                            case 0:
                                switch fi22 {
                                case 127:
                                    v5 = append(v5, (^(v4)))
                                case 0:
                                    if false {
                                        switch (fi20 ^ 793) {
                                        case 603:
                                            v1 = v4
                                        case 127:
                                            var v28 int64 = 42
                                            _ = v28
                                        }
                                    } else {
                                        var v29 string = (("a" + "foo bar") + "")
                                        fmt.Println(v1, fi20)
                                        v1 += v1
                                        _ = v29
                                    }
                                }
                            case -883:
                                fmt.Println(v4)
                            case 127:
                                var v30 float32 = ((0.0 + 1e10) - (3.14159 * 1e10))
                                _ = v30
                            default:
                                v4 -= 2147483647
                            }
                        }
                    } else {
                        v1 += (v4 | v4)
                    }
                case "a":
                    v1 = v1
                default:
                    v4 += (v1 + v4)
                }
                _ = fi22
            }
            v3 = v3
        }
        _ = fi20
    }
    v1 = (^(v4))
    v31 := GenS2{F0: "hello", F1: "hello"}
    _ = len(v5)
    v32 := GenS2{F0: "hello", F1: "foo bar"}
    if false {
        var v33 uint64 = 2
        _ = v33
    } else {
        v34 := GenS3{F0: 42, F1: 278}
        v35 := GenS4{F0: -1, F1: 1000}
        _ = v34
        _ = v35
    }
    switch 4294967295 {
    case 880:
        v2[42] = 2
    case 2147483647:
        v36 := GenS1{F0: v4, F1: (v1 + (v4 ^ (v1 / ((v4) | 1)))), F2: 32767}
        _ = v36
    }
    v3 = (v3 << ((82) & 63))
    var v37 int64 = 873
    switch 32767 {
    case 4294967295:
        var v38 uint64 = uint64(v4)
        _ = v38
    case -133:
        fmt.Println(v37, v4, v1)
    default:
        if false {
            fmt.Println(v37, v37)
        } else {
            v5 = append(v5, (v1 ^ 1))
            v39 := GenS4{F0: v37, F1: 1}
            var v40 uint16 = 32767
            _ = v39
            _ = v40
        }
    }
    if (2 == 4294967295) {
        v1 += 0
        for fi41 := 0; fi41 < 17; fi41++ {
            for fi42 := 0; fi42 < 8; fi42++ {
                switch (fi41 / (((fi41 ^ 479)) | 1)) {
                case -1:
                    v3 += (v3 >> ((uint(v4)) & 63))
                case 9223372036854775807:
                    switch ("hello" + "") {
                    case "hello":
                        var v43 int = int(v3)
                        _ = v43
                    case "":
                        var v44 string = (("hello" + "hello") + "\n\t")
                        _ = v44
                    case "a":
                        var v45 uint = 987
                        _ = v45
                    default:
                        v37 = 0
                    }
                default:
                    v46 := GenS3{F0: 1000, F1: 744}
                    _ = v46
                }
                _ = fi42
            }
            v3 += (v3 * (v3 % (((v3 / ((42) | 1))) | 1)))
            _ = fi41
        }
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        if (!(0 != -541)) {
            v3 = uint8(v1)
            v1 = (^(v1))
        } else {
            var v48 int32 = 2147483647
            _ = v48
        }
        v49 := GenS3{F0: -32768, F1: 510}
        if (false || ("\n\t" == ("hello" + ""))) {
            v50 := v2[(^(v1))]
            _ = v50
        } else {
            v3 = (v3 + 127)
            if (-1 != 100) {
                for fi51 := 0; fi51 < 3; fi51++ {
                    v52 := GenS5{F0: ("\n\t" + "hello"), F1: 2147483647, F2: 42}
                    var v53 int64 = v37
                    v3 += 0
                    v54 := GenS1{F0: (v1 | 556), F1: v1, F2: 42}
                    _ = fi51
                    _ = v52
                    _ = v53
                    _ = v54
                }
                v37 = (v37 * (v37 - (v37 << ((1) & 63))))
                v55 := GenS6{F0: 65535, F1: 4294967295}
                _ = v55
            }
            fmt.Println(v4, v4, v4)
        }
        _ = fi47
        _ = v49
    }
    delete(v2, (^(v4)))
    v1 = v4
    v56 := GenS7{F0: (^(v3)), F1: 963, F2: 42}
    if false {
        var v57 uint32 = (v1 % (((v4 % (((v4 % ((v1) | 1))) | 1))) | 1))
        if (!(42 > 972)) {
            v37 = -1
            var v58 int64 = v37
            v2[42] = (v4 - (v57 & 1))
            _ = v58
        }
        _ = v57
    } else {
        var v59 float32 = (float32(3.14159) * (1.5 / (3.14159 + (1e10 / 0.1))))
        var v60 float32 = float32(0.0)
        _ = v59
        _ = v60
    }
    var v61 int16 = 598
    fmt.Println(v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v61
}

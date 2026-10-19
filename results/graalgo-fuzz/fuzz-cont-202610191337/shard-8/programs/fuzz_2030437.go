package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int16; F2 bool }
type GenS6 struct { F0 uint16; F1 string }
func genFn0() int16 {
    var v7 int16 = -32768
    var v8 uint = 18446744073709551615
    v8 += v8
    v9 := []int{-345, 1}
    if (true || (!(2147483647 > 1))) {
        v10 := GenS1{F0: uint8(v8)}
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return (v7 << (((v8 - 37)) & 63))
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 int64 = -518
    var v5 uint64 = v1
    v6 := make(map[int]string)
    fmt.Println(v4, v4, v4)
    v11 := GenS1{F0: 255}
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    delete(v3, 32768)
    if (-805 <= -404) {
        if (!(("a" <= ("hello" + "")) || true)) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v1 -= v1
                v1 = v1
                _ = fi13
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint64 = (^(v1))
                if (fi14 < 9223372036854775807) {
                    switch (fi14 | (fi14 ^ (fi14 & (fi14 << ((348) & 63))))) {
                    case 365:
                        var v16 float32 = (0.1 - (0.0 * 0.0))
                        _ = v16
                    case 0:
                        v2[fi14] = fi14
                    case 9223372036854775807:
                        var v17 string = "\n\t"
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v19 := []int{-1, (-(fi18)), fi18, 128}
                            _ = fi18
                            _ = v19
                        }
                    }
                    switch int(v15) {
                    case -209:
                        v3[fi14] = true
                    case 4294967295:
                        switch fi14 {
                        case -1:
                            switch "a" {
                            case "\n\t":
                                v5 -= v15
                            }
                        case 0:
                            var v20 uint32 = 0
                            _ = v20
                        default:
                            v1 -= (^(v15))
                        }
                    default:
                        v1 -= (v15 % (((v15 % (((^(v15))) | 1))) | 1))
                    }
                }
                v21 := GenS1{F0: 62}
                _ = fi14
                _ = v15
                _ = v21
            }
        } else {
            v6[387] = "a"
            v22 := GenS1{F0: 42}
            switch "\n\t" {
            case "a":
                if (2 == 63) {
                    fmt.Println(v5, v4, v1)
                } else {
                    switch "a" {
                    case "\n\t":
                        v4 = v4
                    case "a":
                        v4 -= (v4 << ((18446744073709551615) & 63))
                    default:
                        fmt.Println(v5)
                    }
                    v5 = (v1 << ((42) & 63))
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v24 := GenS1{F0: 255}
                        v4 = v4
                        _ = fi23
                        _ = v24
                    }
                }
            case "hello":
                v25 := GenS1{F0: 42}
                _ = v25
            case "":
                var v26 int16 = 37
                _ = v26
            default:
                if (false || (151 >= 600)) {
                    v5 = (v5 % ((v5) | 1))
                }
            }
            _ = v22
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            v5 -= (v1 | 146)
            _ = fi27
        }
        v28 := GenS2{F0: 2}
        _ = v28
    } else {
        var v29 uint8 = 0
        if false {
            v30 := []int{-9223372036854775808, -9223372036854775808}
            for fi31 := 0; fi31 < 13; fi31++ {
                fmt.Println(v5, fi31, v1)
                switch (fi31 ^ int(v29)) {
                case -9223372036854775808:
                    var v32 uint32 = 65535
                    _ = v32
                case 0:
                    if false {
                        v2[2147483647] = 32768
                    } else {
                        fmt.Println(v4, v5, v1)
                        fmt.Println(v5)
                    }
                case -1:
                    fmt.Println(fi31)
                default:
                    switch "\n\t" {
                    case "a":
                        v29 += v29
                    case "":
                        v1 = v5
                    case "foo bar":
                        var v33 uint8 = 204
                        _ = v33
                    }
                }
                _ = len(v3)
                _ = fi31
            }
            _ = v30
        }
        v5 = v1
        _ = v29
    }
    var v34 uint64 = 18446744073709551615
    var v35 int = -876
    for fi36 := 0; fi36 < 8; fi36++ {
        v4 += v4
        if (((v34 ^ (v34 + (v34 | (v34 >> ((0) & 63))))) == (v34 >> ((390) & 63))) || true) {
            v35 = 88
            fmt.Println(fi36, v35, fi36)
            v35 = v35
        } else {
            if false {
                var v37 string = ""
                if false {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        if ((fi36 - v35) > fi38) {
                            v35 = fi36
                            fmt.Println(fi38)
                            v4 -= 610
                        }
                        var v39 int8 = 1
                        if true {
                            if (v39 != (v39 % (((v39 ^ (v39 << ((717) & 63)))) | 1))) {
                                v40 := GenS1{F0: 2}
                                if (0 < 127) {
                                    var v41 float64 = 3.14159
                                    _ = v41
                                } else {
                                    v42 := GenS3{F0: 2, F1: int8(v5), F2: v37}
                                    _ = v42
                                }
                                v4 -= 9223372036854775807
                                _ = v40
                            } else {
                                v43 := []int{v35, fi38}
                                var v44 uint64 = 4294967295
                                _ = v43
                                _ = v44
                            }
                            v45 := []int{}
                            _ = v45
                        } else {
                            v46 := []int{(fi38 + (v35 - v35))}
                            _ = v46
                        }
                        _ = fi38
                        _ = v39
                    }
                    if true {
                        var v47 int32 = -610
                        for fi48 := 0; fi48 < 2; fi48++ {
                            v49 := []int{fi36, int(v47), (fi48 ^ fi36)}
                            _ = fi48
                            _ = v49
                        }
                        _ = v47
                    } else {
                        if false {
                            var v50 int64 = v4
                            _ = v50
                        }
                        if false {
                            v35 = (fi36 * (v35 & -14))
                            fmt.Println(v37, v1, v4)
                        }
                    }
                } else {
                    for fi51 := 0; fi51 < 17; fi51++ {
                        if false {
                            var v52 uint = 405
                            v53 := GenS4{F0: v5, F1: 42, F2: 394}
                            var v54 uint32 = 460
                            _ = v52
                            _ = v53
                            _ = v54
                        } else {
                            if (0 <= -32768) {
                                fmt.Println(fi51, v4, v35)
                                fmt.Println(fi51, fi51)
                            }
                            if ((!("a" == ((v37 + "a") + (v37 + ((v37 + (v37 + "foo bar")) + "foo bar"))))) && ((true || true) && false)) {
                                if ((851 <= 255) || true) {
                                    fmt.Println(v5, v5, v4)
                                    var v55 uint32 = 2
                                    _ = v55
                                }
                                v56 := GenS1{F0: 0}
                                _ = v56
                            } else {
                                v57 := []int{fi36, (fi36 << ((4294967295) & 63))}
                                v5 -= 4294967295
                                switch 2147483647 {
                                case -835:
                                    v34 += v5
                                }
                                if ("foo bar" > "\n\t") {
                                    v58 := GenS5{F0: 1, F1: genFn0(), F2: true}
                                    var v59 int16 = int16(v34)
                                    var v60 int32 = 136
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    fmt.Println(v5, v1, v1)
                                }
                                switch (v37 + "") {
                                case "a":
                                    var v61 uint = 4294967295
                                    _ = v61
                                case "\n\t":
                                    v5 = 2
                                case "foo bar":
                                    var v62 uint8 = 42
                                    if ((("a" + v37) + (("hello" + v37) + v37)) != (v37 + ("hello" + v37))) {
                                        v62 = 71
                                        var v63 float32 = 3.14159
                                        fmt.Println(v4)
                                        v64 := GenS6{F0: 32767, F1: ((v37 + "hello") + (v37 + v37))}
                                        _ = v63
                                        _ = v64
                                    }
                                    var v65 bool = false
                                    _ = v62
                                    _ = v65
                                default:
                                    var v66 float32 = float32(1e10)
                                    _ = v66
                                }
                                _ = v57
                            }
                        }
                        _ = fi51
                    }
                    var v67 bool = false
                    if v67 {
                        var v68 int64 = -899
                        var v69 uint64 = v5
                        _ = v68
                        _ = v69
                    }
                    _ = v67
                }
                _ = v37
            }
            var v70 uint64 = (v1 * (v1 & v1))
            fmt.Println(v34)
            _ = v70
        }
        var v71 int64 = (v4 & -273)
        _ = fi36
        _ = v71
    }
    v4 += (v4 % ((1) | 1))
    v34 = v1
    var v72 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}

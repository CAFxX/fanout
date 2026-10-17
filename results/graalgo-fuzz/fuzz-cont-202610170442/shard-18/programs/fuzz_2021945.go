package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS7 struct { F0 int8; F1 bool }
func genFn0(p0 uint8, p1 int, p2 int) bool {
    p1 = (p1 & (p2 * 636))
    v7 := []int{p1, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (true || false)
}
func genFn1(p0 int, p1 string) uint32 {
    var v8 string = p1
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint16 = 255
    var v4 uint8 = 255
    var v5 uint8 = 66
    var v6 int16 = -32768
    fmt.Println(v3)
    switch "hello" {
    case "\n\t":
        fmt.Println(v6, v4, v3)
    case "":
        v6 = 42
    case "hello":
        v5 -= (v5 >> ((176) & 63))
    }
    var v10 string = ("foo bar" + ("hello" + (((("hello" + "hello") + "\n\t") + "\n\t") + "foo bar")))
    var v11 uint = 292
    v11 -= 4294967295
    var v12 float32 = 1e10
    var v13 bool = (v11 <= v11)
    if false {
        delete(v1, genFn1(127, (v10 + v10)))
        if v13 {
            v14 := GenS1{F0: -2147483648}
            fmt.Println(v11, v6, v13)
            _ = v14
        }
    } else {
        v15 := []int{-474, -458, -2147483648, 9223372036854775807}
        _ = v15
    }
    fmt.Println(v13, v13, v3)
    var v16 uint = v11
    var v17 string = ("foo bar" + "hello")
    if true {
        v2[436] = 311
        v13 = (true || v13)
        v18 := GenS1{F0: -1}
        _ = v18
    } else {
        v19 := []int{-1, 128, 2147483647, -68}
        _ = v19
    }
    v20 := GenS2{F0: (v6 >> (((v16 ^ v16)) & 63)), F1: 2147483647, F2: 42}
    v21 := []int{42, -244}
    switch 593 {
    case 32767:
        v3 -= 42
    case -1:
        var v22 int64 = 182
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        if (!(!v13)) {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            v24 := GenS2{F0: 359, F1: 0, F2: (^(v11))}
            _ = v24
        } else {
            switch (fi23 & fi23) {
            case -951:
                fmt.Println(v6, v17)
            case -1:
                if (0 >= v3) {
                    _ = len(v21)
                    v5 = 1
                    v21 = append(v21, (-(fi23)))
                } else {
                    if (!false) {
                        v25 := GenS2{F0: 1, F1: 2147483647, F2: v16}
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v12 = float32(0.0)
                            _ = fi26
                        }
                        for fi27 := 0; fi27 < 19; fi27++ {
                            if (!((v3 <= (v3 ^ v3)) && (248 >= 489))) {
                                v6 += (^(v6))
                                var v28 uint8 = 127
                                if (v12 <= (v12 - (float32(1e10) * (v12 + ((0.1 / (v12 * (v12 - -2.25))) / v12))))) {
                                    v28 = v4
                                } else {
                                    if false {
                                        fmt.Println(v13)
                                    } else {
                                        v12 += 3.14159
                                        v29 := GenS1{F0: 42}
                                        _ = v29
                                    }
                                    var v30 uint16 = (v3 + uint16(v4))
                                    _ = v30
                                }
                                _ = v28
                            } else {
                                v31 := GenS3{F0: (v3 << ((672) & 63))}
                                if (v17 <= v17) {
                                    v6 -= v6
                                } else {
                                    var v32 int = 32767
                                    _ = v32
                                }
                                _ = v31
                            }
                            var v33 uint16 = (v3 / (((v3 & 0)) | 1))
                            var v34 float32 = 1.5
                            v33 -= v33
                            if (!v13) {
                                v35 := v1[uint32(v11)]
                                v36 := GenS4{F0: v5, F1: -2147483648, F2: (-(v6))}
                                v37 := GenS5{F0: v16, F1: 4294967295}
                                v34 = 3.14159
                                if (1 != 42) {
                                    _ = len(v21)
                                    var v38 bool = (!((uint8(v11) < 80) && ((fi23 % (((fi27 << ((18446744073709551615) & 63))) | 1)) >= 2147483647)))
                                    if (((uint32(v16) == 214) && false) && true) {
                                        v6 = v6
                                    }
                                    _ = v38
                                } else {
                                    var v39 uint64 = 123
                                    if (!v13) {
                                        v40 := v1[1]
                                        var v41 int64 = 1
                                        v35 = (v40 + "foo bar")
                                        _ = v40
                                        _ = v41
                                    } else {
                                        fmt.Println(fi23, v6, v5)
                                        v11 += v16
                                    }
                                    v17 = v10
                                    _ = v39
                                }
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                fmt.Println(v6, v34)
                                var v42 int32 = 2147483647
                                var v43 uint = (v11 | (^(v11)))
                                _ = v42
                                _ = v43
                            }
                            if v13 {
                                var v44 float64 = 0.1
                                if (42 < fi23) {
                                    v45 := GenS5{F0: (v11 + v11), F1: 0}
                                    v44 += float64(float32((((v44 - v44) - v44) + 1e10)))
                                    fmt.Println(v11, v3)
                                    fmt.Println(fi23)
                                    v46 := GenS6{F0: (^(v33)), F1: v6, F2: 1}
                                    _ = v45
                                    _ = v46
                                }
                                _ = v44
                            }
                            _ = fi27
                            _ = v33
                            _ = v34
                        }
                        _ = v25
                    } else {
                        switch 397 {
                        case 0:
                            v21 = append(v21, fi23)
                        case 32768:
                            var v47 uint = 255
                            _ = v47
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v48 int64 = 0
                            _ = v48
                        case "\n\t":
                            fmt.Println(v12, v11, v13)
                        case "":
                            if (false || (v17 != "")) {
                                v49 := GenS7{F0: 100, F1: (-2147483648 <= -2147483648)}
                                _ = v49
                            }
                        default:
                            v50 := GenS4{F0: v4, F1: -2147483648, F2: v6}
                            _ = v50
                        }
                        v2[65535] = -892
                    }
                    if v13 {
                        for fi51 := 0; fi51 < 9; fi51++ {
                            var v52 string = (v10 + "hello")
                            v5 = v4
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = fi51
                            _ = v52
                        }
                        v6 += (v6 + (v6 ^ (^(v6))))
                    }
                }
            case 128:
                if (1 != 4294967295) {
                    switch "a" {
                    case "\n\t":
                        if false {
                            for fi53 := 0; fi53 < 10; fi53++ {
                                var v54 string = "\n\t"
                                fmt.Println(v17, v54)
                                v55 := GenS4{F0: uint8(v16), F1: -2147483648, F2: v6}
                                switch fi53 {
                                case 32767:
                                    switch v17 {
                                    case "a":
                                        if (!v13) {
                                            v5 -= 128
                                        }
                                    case "\n\t":
                                        var v56 int64 = -9223372036854775808
                                        _ = v56
                                    case "":
                                        switch v17 {
                                        case "a":
                                            switch v10 {
                                            case "\n\t":
                                                delete(v2, 0)
                                                v11 += 703
                                            case "":
                                                var v57 uint64 = 1
                                                _ = v57
                                            case "a":
                                                v12 += v12
                                            default:
                                                if true {
                                                    if len(v21) > 0 {
                                                        v21[0] = -2147483648
                                                    }
                                                } else {
                                                    fmt.Println(v3, fi53, v5)
                                                    v12 += 1e10
                                                }
                                            }
                                        }
                                    default:
                                        v10 = (((v54 + ("" + v10)) + v10) + "a")
                                    }
                                }
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                            switch "\n\t" {
                            case "a":
                                v58 := GenS6{F0: (v3 - 1), F1: (v6 + 1000), F2: 0}
                                _ = v58
                            default:
                                v59 := GenS6{F0: 858, F1: v6, F2: 208}
                                _ = v59
                            }
                            var v60 bool = v13
                            _ = v60
                        }
                    case "hello":
                        var v61 uint32 = 4294967295
                        _ = v61
                    }
                    v10 = (v10 + (v10 + v17))
                }
            }
        }
        if v13 {
            v13 = v13
        } else {
            delete(v2, uint32(v16))
        }
        _ = fi23
    }
    for fi62 := 0; fi62 < 18; fi62++ {
        if v13 {
            fmt.Println(v5, v11)
        }
        _ = fi62
    }
    var v63 float32 = float32(float64(float32(3.14159)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}

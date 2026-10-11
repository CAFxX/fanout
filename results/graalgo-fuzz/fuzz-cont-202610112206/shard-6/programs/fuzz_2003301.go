package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 uint; F2 int64 }
type GenS6 struct { F0 int8; F1 int8; F2 uint64 }
type GenS7 struct { F0 int32; F1 string; F2 int64 }
func genFn0() uint8 {
    switch "\n\t" {
    case "":
        var v6 int64 = -9223372036854775808
        _ = v6
    case "hello":
        v7 := GenS1{F0: 808}
        _ = v7
    default:
        v8 := GenS1{F0: 483}
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{fi9, 42, (fi9 * fi9), (-(fi9))}
        var v11 int16 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float32 = (float32(0.1) * float32(3.14159))
    fmt.Println(v12)
    var v13 bool = (false && true)
    _ = v12
    _ = v13
    return 42
}
func genFn1(p0 int16, p1 int16, p2 int64) int {
    var v14 int = 32768
    p2 = p2
    v15 := []int{v14, -372, (v14 & 927), v14}
    _ = len(v15)
    v16 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return v14
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    v4 := []string{("hello" + (("foo bar" + ("\n\t" + "\n\t")) + "a"))}
    var v5 int16 = 1000
    fmt.Println(v3, v5, v3)
    v5 += -1
    var v17 string = ""
    if false {
        var v18 int8 = 42
        fmt.Println(v1, v17)
        _ = v18
    } else {
        var v19 int8 = 1
        _ = v19
    }
    if (v17 == (v17 + ("" + "foo bar"))) {
        v17 = "hello"
        v20 := GenS2{F0: (^(v1)), F1: (v1 << ((737) & 63)), F2: 2}
        v17 = (v17 + (v17 + "hello"))
        _ = v20
    } else {
        v3 = -32768
        fmt.Println(v17)
        if (v5 == (v5 * 1000)) {
            fmt.Println(v1)
            if true {
                v3 = v5
                v21 := GenS3{F0: 1, F1: 2, F2: 32767}
                v1 = 2147483647
                _ = v21
            }
            v22 := GenS1{F0: 1}
            _ = v22
        }
    }
    fmt.Println(v3)
    for fi23 := 0; fi23 < 2; fi23++ {
        if (654 > 2) {
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := GenS4{F0: 1000, F1: 745, F2: true}
                _ = fi24
                _ = v25
            }
        }
        _ = fi23
    }
    if true {
        switch ("foo bar" + "\n\t") {
        case "hello":
            switch -9223372036854775808 {
            case 0:
                switch v17 {
                case "\n\t":
                    for fi26 := 0; fi26 < 11; fi26++ {
                        _ = len(v2)
                        var v27 bool = (127 == 128)
                        _ = fi26
                        _ = v27
                    }
                }
            case -1:
                var v28 uint64 = uint64(v1)
                _ = v28
            case 20:
                v5 -= v5
            default:
                if (63 > 18446744073709551615) {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 int16 = int16(fi29)
                        for fi31 := 0; fi31 < 15; fi31++ {
                            v5 += (v30 / ((v30) | 1))
                            v5 = v3
                            v5 += v3
                            _ = len(v2)
                            _ = fi31
                        }
                        _ = fi29
                        _ = v30
                    }
                    v32 := v2[(v1 << ((500) & 63))]
                    _ = v32
                } else {
                    switch v17 {
                    case "hello":
                        var v33 int64 = 0
                        _ = v33
                    case "":
                        v34 := GenS5{F0: 127, F1: 42, F2: -658}
                        _ = v34
                    case "a":
                        v1 -= (v1 << ((4294967295) & 63))
                    default:
                        switch "" {
                        case "a":
                            v35 := GenS2{F0: (v1 << ((866) & 63)), F1: 65535, F2: 104}
                            _ = v35
                        case "\n\t":
                            v3 -= 1
                        }
                    }
                }
            }
        case "foo bar":
            v3 = v5
        case "a":
            v1 = v1
        default:
            v5 += (v3 & v5)
        }
        v36 := GenS1{F0: 32767}
        for fi37 := 0; fi37 < 9; fi37++ {
            if false {
                if (!false) {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        if (!false) {
                            if true {
                                v17 = "a"
                                if (true || ((((v17 + "") + (v17 + "foo bar")) + v17) == v17)) {
                                    fmt.Println(v17, v1, fi37)
                                    v1 -= (v1 >> ((0) & 63))
                                    v17 = v17
                                } else {
                                    delete(v2, (v1 + (v1 + (v1 & (v1 | 42)))))
                                    var v39 uint = 790
                                    var v40 uint16 = uint16(v1)
                                    if true {
                                        v17 = "a"
                                        v3 = (v5 / (((-(v3))) | 1))
                                        v41 := GenS6{F0: -128, F1: 100, F2: 0}
                                        switch (fi37 / ((int(v1)) | 1)) {
                                        case 42:
                                            var v42 string = v17
                                            _ = v42
                                        case -946:
                                            if (4294967295 == 174) {
                                                delete(v2, (v1 << ((497) & 63)))
                                            }
                                        case 2147483647:
                                            delete(v2, v1)
                                        default:
                                            if ((float64(0.1) / float64(float32((1e10 - 1.5)))) <= (0.1 * 3.14159)) {
                                                v43 := GenS7{F0: -838, F1: (v17 + "hello"), F2: 1}
                                                _ = v43
                                            }
                                        }
                                        _ = v41
                                    }
                                    _ = v39
                                    _ = v40
                                }
                                var v44 string = ("a" + ("a" + v17))
                                var v45 uint32 = (v1 - (^(v1)))
                                _ = v44
                                _ = v45
                            }
                            switch v17 {
                            case "":
                                fmt.Println(fi38)
                            case "foo bar":
                                var v46 uint16 = 1
                                _ = v46
                            }
                            var v47 uint8 = 128
                            _ = v47
                        }
                        _ = fi38
                    }
                    v4 = append(v4, v17)
                } else {
                    for fi48 := 0; fi48 < 13; fi48++ {
                        var v49 int = (-(fi48))
                        var v50 uint64 = 903
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 15; fi51++ {
                        if (!(!(((true && false) || true) || true))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v52 bool = (-625 == 164)
                            _ = v52
                        }
                        _ = len(v4)
                        _ = fi51
                    }
                    if false {
                        v3 = (v5 % ((v3) | 1))
                    }
                }
            }
            _ = fi37
        }
        _ = v36
    } else {
        v3 += 1
        var v53 int8 = 42
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

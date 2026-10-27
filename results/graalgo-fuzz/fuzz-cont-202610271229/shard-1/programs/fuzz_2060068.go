package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 int8; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int8 }
func genFn0(p0 uint) uint16 {
    var v6 uint32 = 0
    if (uint64(p0) >= 2) {
        if (!true) {
            v7 := GenS1{F0: 1, F1: 1, F2: v6}
            _ = v7
        }
    } else {
        if true {
            p0 = p0
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    if false {
                        var v10 float32 = (0.1 + float32(1.5))
                        var v11 int16 = 1
                        v12 := []int{(fi9 / (((fi8 ^ (fi9 % ((fi9) | 1)))) | 1)), (^(fi8))}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        p0 = p0
                    }
                    p0 = (^(p0))
                    fmt.Println(p0, fi8, p0)
                    if (false && true) {
                        v13 := GenS2{F0: 32767, F1: 32767, F2: v6}
                        var v14 int8 = 38
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    if (!(-1 < 1000)) {
                        if true {
                            if (-2147483648 <= 1) {
                                v16 := []int{}
                                if (!(42 >= 1)) {
                                    var v17 bool = (true && false)
                                    v17 = ((v17 || v17) && v17)
                                    _ = v17
                                }
                                _ = v16
                            } else {
                                if (!(-2.25 != 0.0)) {
                                    p0 -= 42
                                }
                            }
                        }
                    } else {
                        p0 = p0
                    }
                    v18 := GenS3{F0: 1, F1: 42, F2: ("a" + ("" + "hello"))}
                    v19 := GenS4{F0: 26, F1: 32767, F2: 2}
                    _ = fi15
                    _ = v18
                    _ = v19
                }
                _ = fi8
            }
            v20 := GenS4{F0: -1, F1: 764, F2: 18446744073709551615}
            _ = v20
        } else {
            v21 := []int{32767}
            if false {
                if false {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 int16 = -1
                        _ = fi22
                        _ = v23
                    }
                } else {
                    fmt.Println(p0, v6)
                    if len(v21) > 0 {
                        v21[0] = 0
                    }
                }
                var v24 bool = (!false)
                switch 32767 {
                case 127:
                    v25 := GenS5{F0: 1, F1: "foo bar"}
                    _ = v25
                case 132:
                    fmt.Println(v6)
                case 0:
                    var v26 int16 = 1
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        fmt.Println(v6)
                        v24 = (1 < 42)
                        v28 := GenS2{F0: 42, F1: 646, F2: v6}
                        _ = fi27
                        _ = v28
                    }
                }
                _ = v24
            } else {
                var v29 float64 = 0.1
                _ = v29
            }
            var v30 int32 = 1
            _ = v21
            _ = v30
        }
    }
    fmt.Println(v6, p0)
    var v31 uint64 = 496
    if (((p0 == (p0 / ((p0) | 1))) || true) || ((!(-2147483648 <= 169)) || false)) {
        if false {
            var v32 float32 = ((0.0 * 1.5) * 0.1)
            v33 := []int{-367, 142, -572, 122}
            v6 = 19
            _ = v32
            _ = v33
        }
    }
    _ = p0
    _ = v31
    _ = v6
    return 285
}
func genFn1(p0 int8) uint64 {
    var v34 string = "\n\t"
    v35 := GenS4{F0: (p0 % (((p0 % ((p0) | 1))) | 1)), F1: -32768, F2: 686}
    for fi36 := 0; fi36 < 18; fi36++ {
        v34 = v34
        _ = fi36
    }
    _ = p0
    _ = v34
    _ = v35
    return 18446744073709551615
}

func main() {
    var v1 int64 = 44
    var v2 bool = true
    var v3 uint = 4294967295
    v4 := []uint32{42, uint32(v3), 861}
    var v5 int8 = 100
    v2 = (!(float32(float64(((0.0 - -2.25) - -2.25))) <= -2.25))
    v37 := GenS1{F0: genFn0(1), F1: 1, F2: 65535}
    var v38 float32 = 3.14159
    for fi39 := 0; fi39 < 2; fi39++ {
        if (63 < v5) {
            v1 += -9223372036854775808
            v3 = 266
        } else {
            switch fi39 {
            case -306:
                switch fi39 {
                case 4294967295:
                    switch (fi39 >> (((v3 << ((v3) & 63))) & 63)) {
                    case 32768:
                        if ("hello" <= "hello") {
                            if false {
                                if true {
                                    var v40 bool = v2
                                    _ = v40
                                } else {
                                    v2 = (v2 || ((18446744073709551615 > genFn1(36)) || v2))
                                    v38 += 3.14159
                                    v41 := GenS6{F0: 42}
                                    _ = v41
                                }
                            } else {
                                var v42 uint64 = 0
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v43 int32 = -337
                                _ = v42
                                _ = v43
                            }
                        } else {
                            v3 -= 0
                        }
                    default:
                        v3 += v3
                    }
                default:
                    v44 := GenS6{F0: 662}
                    _ = v44
                }
            case 32768:
                v45 := GenS2{F0: 42, F1: 42, F2: 2147483647}
                _ = v45
            case 128:
                _ = len(v4)
            default:
                v1 = (v1 >> ((v3) & 63))
            }
            var v46 uint16 = 70
            var v47 int32 = -985
            _ = v46
            _ = v47
        }
        _ = fi39
    }
    v1 -= -9223372036854775808
    var v48 uint32 = 607
    v49 := GenS6{F0: genFn1(100)}
    v4 = append(v4, (v48 & 24))
    fmt.Println(v48, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = v3
    v3 = v3
    var v50 int32 = 42
    var v51 string = "hello"
    if (true || (!v2)) {
        fmt.Println(v51, v51, v50)
    } else {
        if ((v51 + "") >= ("hello" + (("hello" + v51) + ("" + v51)))) {
            var v52 uint32 = v48
            _ = v52
        }
        if ((9223372036854775807 != (v1 >> ((v3) & 63))) || v2) {
            v5 = 1
            if len(v4) > 0 {
                v4[0] = v48
            }
            for fi53 := 0; fi53 < 8; fi53++ {
                if ((v48 == 2147483647) || false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v54 int64 = (v1 << (((v3 + uint(v48))) & 63))
                    if (((2147483648 <= 2147483647) && (1.5 != 3.14159)) || false) {
                        switch -1 {
                        case -581:
                            if v2 {
                                v55 := GenS6{F0: 266}
                                _ = v55
                            } else {
                                fmt.Println(v5, v1)
                                v4 = append(v4, v48)
                            }
                        case 127:
                            _ = len(v4)
                        case 197:
                            v56 := GenS7{F0: v5, F1: v5}
                            _ = v56
                        }
                    }
                    v5 += v5
                    _ = v54
                }
                _ = fi53
            }
        }
        switch 42 {
        case 2147483647:
            v57 := GenS3{F0: v5, F1: int64(v3), F2: (v51 + "hello")}
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}

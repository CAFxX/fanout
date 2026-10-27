package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0(p0 int8, p1 int) int {
    p1 = (p1 - 1)
    v7 := GenS1{F0: 42}
    var v8 string = "a"
    var v9 float64 = 0.0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p1 * 32767)
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.1
    v4 := []int{1, 861}
    var v5 uint32 = 2
    var v6 uint16 = 1
    var v10 int = 653
    v1 = (-(v1))
    if false {
        v5 = v5
        var v11 bool = true
        var v12 float32 = (1e10 * (((0.0 * (1.5 / 3.14159)) / 1.5) / 1.0))
        _ = v11
        _ = v12
    }
    v4 = append(v4, 32768)
    if true {
        var v13 int64 = v1
        v5 -= (v5 % ((v5) | 1))
        for fi14 := 0; fi14 < 8; fi14++ {
            switch "\n\t" {
            case "hello":
                v5 = v5
            case "":
                v10 = fi14
            case "\n\t":
                switch "\n\t" {
                case "":
                    var v15 bool = (0.0 != 3.14159)
                    _ = v15
                case "hello":
                    v16 := GenS1{F0: (v6 >> ((42) & 63))}
                    _ = v16
                default:
                    if len(v4) > 0 {
                        v4[0] = fi14
                    }
                }
            }
            _ = fi14
        }
        _ = v13
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v2[v5] = false
        v1 = (v1 + int64(v6))
        v18 := GenS2{F0: ("hello" + "\n\t")}
        _ = fi17
        _ = v18
    }
    switch int(v5) {
    case 42:
        var v19 int16 = -32768
        _ = v19
    case -9223372036854775808:
        v4 = append(v4, v10)
    case -943:
        v3 += v3
    default:
        var v20 int64 = v1
        _ = v20
    }
    v10 -= v10
    if true {
        v2[42] = true
        v21 := v2[v5]
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint64 = 4294967295
                v24 -= v24
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        var v25 int32 = 1
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 uint32 = v5
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v5 = v5
    }
    if true {
        v1 = (v1 - (-(v1)))
        if ((42 >= -1) && true) {
            if (831 >= (^(v6))) {
                var v28 bool = false
                var v29 int32 = 42
                switch v10 {
                case -2147483648:
                    var v30 uint64 = 315
                    _ = v30
                case -9223372036854775808:
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v32 := GenS1{F0: 255}
                        for fi33 := 0; fi33 < 15; fi33++ {
                            if v28 {
                                var v34 int32 = 0
                                _ = v34
                            } else {
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    v35 := v2[v5]
                                    _ = v35
                                default:
                                    v3 = v3
                                    v5 = (v5 + v5)
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi33
                                }
                                if true {
                                    var v36 bool = (int8(fi33) >= 42)
                                    v5 = (v5 | (v5 >> ((18446744073709551615) & 63)))
                                    if v28 {
                                        v37 := GenS1{F0: (v6 / (((v6 >> ((18446744073709551615) & 63))) | 1))}
                                        v38 := v2[v5]
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if false {
                                            var v39 int = fi33
                                            _ = v39
                                        } else {
                                            delete(v2, v5)
                                            if v28 {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    v2[4294967295] = false
                                                    v40 := GenS2{F0: "a"}
                                                    if v28 {
                                                        v3 -= v3
                                                        v28 = v36
                                                        v41 := GenS3{F0: 145, F1: (v1 / ((v1) | 1)), F2: v29}
                                                        if len(v4) > 0 {
                                                            v4[0] = genFn0(127, (fi33 * (fi33 % ((42) | 1))))
                                                        }
                                                        var v42 int32 = v29
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        var v43 string = ""
                                                        _ = v43
                                                    }
                                                    _ = v40
                                                }
                                            }
                                            if true {
                                                v5 = 20
                                                fmt.Println(v3, v5)
                                                v36 = v28
                                            }
                                        }
                                        if (42 < (v5 / (((v5 | v5)) | 1))) {
                                            var v44 float32 = 0.1
                                            _ = v44
                                        } else {
                                            v45 := GenS1{F0: (^(v6))}
                                            _ = v45
                                        }
                                    }
                                    if (!(v28 && (v28 || false))) {
                                        if len(v4) > 0 {
                                            v4[0] = v10
                                        }
                                        v46 := GenS1{F0: v6}
                                        fmt.Println(v10, v29)
                                        v47 := v2[v5]
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := GenS2{F0: "hello"}
                                    v49 := v2[v5]
                                    _ = v36
                                    _ = v48
                                    _ = v49
                                } else {
                                    var v50 bool = (42 >= 127)
                                    _ = v50
                                }
                                v28 = (false || true)
                                v51 := GenS4{F0: 0, F1: 547, F2: 1}
                                _ = v51
                            }
                            var v52 float64 = ((((v3 + (0.1 * ((v3 / v3) + v3))) + 0.0) - v3) / v3)
                            _ = fi33
                            _ = v52
                        }
                        var v53 uint64 = 2
                        _ = fi31
                        _ = v32
                        _ = v53
                    }
                case 127:
                    switch v10 {
                    case 32767:
                        var v54 float64 = 0.1
                        _ = v54
                    case -2147483648:
                        switch v10 {
                        case -2147483648:
                            var v55 bool = (-1 == v1)
                            _ = v55
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = (v10 << ((42) & 63))
                        }
                    }
                }
                _ = v28
                _ = v29
            } else {
                _ = len(v2)
                var v56 uint16 = (v6 + 32767)
                _ = v56
            }
        }
    }
    v1 = (v1 ^ (v1 / (((v1 / (((v1 / ((1) | 1))) | 1))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 int; F2 int32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int32; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0(p0 uint16, p1 int32, p2 uint64) uint16 {
    var v7 bool = (255 <= 197)
    if (uint(p2) > uint(p0)) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint = 18446744073709551615
            if (1.5 != 0.0) {
                if v7 {
                    v10 := GenS1{F0: (v9 + v9)}
                    _ = v10
                }
                v9 += 179
            }
            _ = fi8
            _ = v9
        }
        var v11 int64 = 9223372036854775807
        _ = v11
    } else {
        var v12 int32 = (p1 / (((p1 + int32(p2))) | 1))
        p1 -= v12
        for fi13 := 0; fi13 < 4; fi13++ {
            v12 = v12
            if v7 {
                p2 -= 206
            }
            _ = fi13
        }
        _ = v12
    }
    var v14 int32 = p1
    v15 := GenS2{F0: v7, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v7
    return (p0 | p0)
}

func main() {
    var v1 uint8 = 42
    var v2 int = -261
    var v3 float32 = 1.5
    var v4 int32 = -2147483648
    v5 := []int{v2, (v2 + (v2 % (((v2 / (((v2 << ((741) & 63))) | 1))) | 1))), v2, v2}
    v6 := make(map[string]string)
    v1 += v1
    v1 = v1
    var v16 uint32 = 821
    for fi17 := 0; fi17 < 8; fi17++ {
        if true {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 string = ""
                if (42 <= (v1 << ((42) & 63))) {
                    if true {
                        var v20 uint64 = 758
                        v1 = 128
                        fmt.Println(v4)
                        v4 -= v4
                        v6[(((v19 + v19) + "hello") + (((v19 + "foo bar") + "hello") + ((v19 + (("\n\t" + (v19 + "a")) + ("" + "a"))) + ("hello" + ("" + v19)))))] = "\n\t"
                        var v21 int16 = 1
                        if false {
                            v4 = 244
                        } else {
                            v22 := GenS1{F0: 1}
                            if len(v5) > 0 {
                                v5[0] = -43
                            }
                            _ = v22
                        }
                        v21 = 0
                        _ = v20
                        _ = v21
                    } else {
                        v5 = append(v5, (v2 % (((v2 * (fi18 % ((fi17) | 1)))) | 1)))
                        if true {
                            v23 := GenS3{F0: -2147483648, F1: 32767, F2: (v4 * (v4 % ((-1) | 1)))}
                            if (0.0 >= (3.14159 + (float64((0.0 * 1e10)) - 3.14159))) {
                                v24 := GenS4{F0: false, F1: v1}
                                if (true && false) {
                                    if (false || (!true)) {
                                        v25 := GenS5{F0: 1, F1: (v4 << ((2) & 63)), F2: 740}
                                        _ = v25
                                    } else {
                                        fmt.Println(v16, v1, v3)
                                        delete(v6, ((v19 + v19) + v19))
                                        if len(v5) > 0 {
                                            v5[0] = (^(fi17))
                                        }
                                        switch fi18 {
                                        case 42:
                                            v16 += (v16 * (v16 * 189))
                                        default:
                                            if (v16 < v16) {
                                                var v26 uint8 = v1
                                                _ = v26
                                            } else {
                                                v3 += v3
                                                v19 = (((v19 + "foo bar") + ("" + (v19 + "hello"))) + v19)
                                            }
                                        }
                                    }
                                    _ = len(v6)
                                    switch v19 {
                                    case "":
                                        v27 := GenS6{F0: 846, F1: 18446744073709551615, F2: true}
                                        _ = v27
                                    }
                                } else {
                                    v19 = "\n\t"
                                    if (4294967295 > 4294967295) {
                                        v4 += v4
                                        var v28 int64 = -693
                                        v1 = (v1 | v1)
                                        var v29 uint8 = v1
                                        switch v19 {
                                        case "":
                                            var v30 int = int(v4)
                                            _ = v30
                                        case "foo bar":
                                            v4 = (v4 * (v4 * (v4 | v4)))
                                        case "\n\t":
                                            var v31 int8 = 42
                                            _ = v31
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                v32 := GenS3{F0: (fi17 % ((4294967295) | 1)), F1: fi18, F2: v4}
                                v19 = "hello"
                                var v33 float32 = v3
                                var v34 int64 = 1
                                _ = v24
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                if (3.14159 < v3) {
                                    v5 = append(v5, (^(v2)))
                                    v35 := GenS7{F0: 255}
                                    v6["foo bar"] = ""
                                    _ = v35
                                } else {
                                    v6[v19] = "\n\t"
                                }
                                var v36 uint16 = genFn0(402, (v4 >> ((65535) & 63)), 42)
                                _ = v36
                            }
                            _ = v23
                        }
                    }
                    var v37 string = v19
                    _ = v37
                } else {
                    var v38 float32 = 1.5
                    var v39 int8 = -128
                    v40 := GenS6{F0: 42, F1: 808, F2: ((v39 / ((81) | 1)) >= v39)}
                    v39 = v39
                    switch v19 {
                    case "":
                        var v41 uint32 = (v16 << ((255) & 63))
                        _ = v41
                    case "foo bar":
                        if false {
                            v4 -= -2147483648
                        }
                    default:
                        v42 := GenS7{F0: 306}
                        _ = v42
                    }
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = len(v5)
                _ = len(v6)
                var v43 uint16 = 1
                _ = fi18
                _ = v19
                _ = v43
            }
            fmt.Println(v1, v3, v16)
            fmt.Println(v4, v4, v4)
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            var v45 int16 = -1
            v2 = (v2 | fi17)
            _ = fi44
            _ = v45
        }
        v46 := GenS5{F0: v4, F1: v4, F2: 255}
        _ = fi17
        _ = v46
    }
    v3 = ((v3 - v3) / v3)
    switch ("" + (("a" + ("hello" + (((("\n\t" + ("hello" + "foo bar")) + "a") + ((("a" + ("hello" + "a")) + "hello") + (((("" + "a") + ("" + ((("foo bar" + (("" + ("" + ("" + ("\n\t" + "a")))) + "foo bar")) + "a") + ((((("hello" + ("\n\t" + "a")) + "\n\t") + "\n\t") + ("a" + "a")) + "\n\t")))) + ("foo bar" + "")) + ""))) + ""))) + ("hello" + ("\n\t" + "hello")))) {
    case "":
        if (("foo bar" >= ("\n\t" + "foo bar")) || ((v4 ^ v4) != v4)) {
            if len(v5) > 0 {
                v5[0] = (v2 ^ v2)
            }
            var v47 string = "a"
            if ((uint64(v16) < uint64(v1)) || true) {
                v6[(("\n\t" + v47) + v47)] = v47
            }
            _ = v47
        }
    case "foo bar":
        var v48 int32 = (v4 * 0)
        _ = v48
    default:
        var v49 uint64 = 620
        _ = v49
    }
    var v50 uint16 = 311
    for fi51 := 0; fi51 < 5; fi51++ {
        for fi52 := 0; fi52 < 6; fi52++ {
            if false {
                v53 := GenS3{F0: fi52, F1: (^(fi51)), F2: v4}
                v50 -= 368
                var v54 uint32 = 2147483648
                _ = v53
                _ = v54
            } else {
                switch 468 {
                case 275:
                    var v55 int = (fi51 + (v2 ^ (v2 * fi51)))
                    _ = v55
                case -9223372036854775808:
                    v2 = 4294967295
                case 4294967295:
                    fmt.Println(v4, v3, fi51)
                }
            }
            fmt.Println(fi52, v50)
            _ = fi52
        }
        v56 := GenS8{F0: 255, F1: (v4 / ((v4) | 1))}
        _ = fi51
        _ = v56
    }
    v4 = (v4 * v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}

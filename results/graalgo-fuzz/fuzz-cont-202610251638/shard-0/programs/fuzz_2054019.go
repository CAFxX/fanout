package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 uint64) int8 {
    fmt.Println(p1, p2)
    v5 := GenS1{F0: "", F1: (p0 * p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint, p2 uint64) bool {
    p0 = p0
    var v6 int64 = 0
    fmt.Println(v6)
    var v7 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v7 != 2)
}

func main() {
    var v1 float64 = (0.0 - (float64(1.5) * (1e10 + 0.1)))
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := make(map[string]bool)
    v8 := []int{-999, -1, 409}
    v9 := GenS1{F0: "\n\t", F1: 4294967295}
    v3 = false
    _ = len(v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            v3 = v3
            if (32767 == 941) {
                var v12 int32 = -2147483648
                _ = v12
            }
            v3 = false
            v4["hello"] = (0 != 32767)
            _ = fi11
        }
        v3 = v3
        fmt.Println(v3, fi10, fi10)
        _ = fi10
    }
    var v13 float32 = -2.25
    if ((0.0 + v1) > v1) {
        v14 := GenS2{F0: 448, F1: (v13 < v13), F2: -32768}
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(fi15, fi16, fi16)
                    v17 := GenS3{F0: 34, F1: 4294967295}
                    if v3 {
                        delete(v2, 0)
                        if v3 {
                            var v18 float32 = (v13 + 0.0)
                            v19 := v4[(("hello" + "\n\t") + ("foo bar" + "\n\t"))]
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(v1, fi15)
                            switch ((("foo bar" + "hello") + "foo bar") + "") {
                            case "":
                                var v20 int8 = genFn0(65535, 42, 42)
                                _ = v20
                            case "hello":
                                var v21 string = "hello"
                                if true {
                                    var v22 int32 = 0
                                    switch v21 {
                                    case "hello":
                                        var v23 uint16 = 255
                                        _ = v23
                                    case "foo bar":
                                        v3 = (-9223372036854775808 <= int64(fi16))
                                    default:
                                        v21 = (v21 + ((v21 + v21) + "\n\t"))
                                    }
                                    v24 := GenS2{F0: 553, F1: (42 != v22), F2: 1000}
                                    _ = v22
                                    _ = v24
                                } else {
                                    v25 := GenS2{F0: 85, F1: v3, F2: 1000}
                                    if (-1 > -1) {
                                        v3 = true
                                    }
                                    _ = v25
                                }
                                var v26 int16 = 42
                                _ = v21
                                _ = v26
                            case "foo bar":
                                if (!((fi16 + (fi15 - fi16)) > (fi15 << ((2) & 63)))) {
                                    fmt.Println(fi16)
                                }
                            default:
                                fmt.Println(v1)
                            }
                        }
                        v8 = append(v8, -2147483648)
                    } else {
                        fmt.Println(v13, fi16)
                    }
                    delete(v4, "hello")
                    _ = fi16
                    _ = v17
                }
            } else {
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 string = ""
                    v28 = v28
                    fmt.Println(fi27, v13, fi15)
                    _ = fi27
                    _ = v28
                }
            }
            v13 -= 0.0
            for fi29 := 0; fi29 < 17; fi29++ {
                fmt.Println(fi15, v13)
                var v30 uint64 = 546
                _ = fi29
                _ = v30
            }
            _ = fi15
        }
        _ = v14
    } else {
        fmt.Println(v1, v3, v3)
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS1{F0: (("foo bar" + "") + "a"), F1: 2147483647}
            _ = fi31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS4{F0: 32767, F1: (("\n\t" + "foo bar") + (("\n\t" + "a") + ("a" + "hello")))}
        if (1 < 2) {
            var v35 uint8 = 0
            v1 -= (-2.25 - v1)
            if (422 <= -9223372036854775808) {
                v13 += (v13 - v13)
                var v36 bool = false
                _ = v36
            }
            _ = v35
        }
        v37 := GenS5{F0: 1, F1: 245, F2: 4294967295}
        _ = fi33
        _ = v34
        _ = v37
    }
    if false {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := v2[1]
        switch ("foo bar" + "foo bar") {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v13, v39, v39)
            case "hello":
                for fi40 := 0; fi40 < 7; fi40++ {
                    switch (-(fi38)) {
                    case -423:
                        if ("foo bar" <= "foo bar") {
                            var v41 int64 = 383
                            if ((!((!v3) || v3)) && ((^(v41)) > -9223372036854775808)) {
                                var v42 int64 = -421
                                switch (fi38 >> ((430) & 63)) {
                                case 32767:
                                    v39 = "hello"
                                case 298:
                                    var v43 float32 = float32((((0.0 - (v1 + v1)) / v1) * 1e10))
                                    _ = v43
                                default:
                                    v3 = false
                                }
                                switch (v39 + v39) {
                                case "":
                                    var v44 int = (fi38 % (((fi38 << ((1) & 63))) | 1))
                                    _ = v44
                                }
                                _ = v42
                            }
                            if v3 {
                                var v45 int = (fi38 << ((891) & 63))
                                delete(v4, "\n\t")
                                v8 = append(v8, (fi40 % ((v45) | 1)))
                                if (84 <= 12) {
                                    fmt.Println(fi40, v13, v3)
                                    _ = len(v4)
                                    var v46 uint64 = 586
                                    _ = v46
                                } else {
                                    if (32767 != -860) {
                                        v47 := GenS6{F0: 1, F1: -912}
                                        v41 = (v41 % ((571) | 1))
                                        _ = v47
                                    } else {
                                        fmt.Println(v45, v13)
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        var v48 int64 = (^(v41))
                                        _ = v48
                                    }
                                }
                                _ = v45
                            } else {
                                var v49 int32 = 42
                                _ = v49
                            }
                            v3 = false
                            if v3 {
                                if v3 {
                                    if len(v8) > 0 {
                                        v8[0] = 127
                                    }
                                } else {
                                    var v50 uint64 = 2
                                    var v51 bool = v3
                                    v50 = v50
                                    _ = v50
                                    _ = v51
                                }
                                if (!(48 < 42)) {
                                    v3 = (197 == 473)
                                }
                            }
                            _ = v41
                        } else {
                            if v3 {
                                switch ("\n\t" + (v39 + (v39 + v39))) {
                                case "":
                                    v1 -= v1
                                case "a":
                                    var v52 uint16 = 837
                                    _ = v52
                                case "hello":
                                    fmt.Println(v13, v3)
                                default:
                                    v53 := GenS7{F0: 127}
                                    _ = v53
                                }
                                _ = len(v2)
                                if (42 < 2) {
                                    var v54 int64 = -314
                                    switch int(v54) {
                                    case -982:
                                        fmt.Println(fi38)
                                    }
                                    _ = v54
                                }
                                var v55 int32 = 0
                                v1 -= v1
                                _ = v55
                            }
                        }
                    default:
                        v3 = ((189 != 668) || true)
                    }
                    switch "\n\t" {
                    case "":
                        var v56 bool = v3
                        _ = v56
                    }
                    _ = fi40
                }
            default:
                v39 = ((v39 + "foo bar") + v39)
            }
        case "foo bar":
            fmt.Println(fi38, v1)
        default:
            var v57 string = v39
            _ = v57
        }
        _ = fi38
        _ = v39
    }
    _ = len(v4)
    var v58 bool = v3
    var v59 float64 = v1
    var v60 bool = v3
    if (v59 > v59) {
        for fi61 := 0; fi61 < 9; fi61++ {
            v8 = append(v8, (fi61 | (fi61 | (fi61 ^ (fi61 & (fi61 | fi61))))))
            v3 = (56 == 0)
            _ = fi61
        }
    } else {
        if len(v8) > 0 {
            v8[0] = 4294967295
        }
        v62 := GenS6{F0: 65535, F1: 1000}
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v59
    _ = v60
    _ = v8
    _ = v9
}

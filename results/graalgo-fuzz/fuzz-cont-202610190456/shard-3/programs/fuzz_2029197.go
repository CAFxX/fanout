package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint32 }
func genFn0(p0 string, p1 int16, p2 int) uint32 {
    var v6 int = (-(p2))
    var v7 int64 = 9223372036854775807
    p1 -= 0
    v8 := GenS1{F0: 0, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 404
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v9 int32 = 180
    if true {
        if (!false) {
            fmt.Println(v9)
            v10 := GenS2{F0: 127, F1: 2, F2: -1}
            _ = v10
        } else {
            var v11 int32 = (v9 + (^(v9)))
            _ = v11
        }
    }
    if false {
        v12 := GenS3{F0: -9223372036854775808, F1: 42, F2: 42}
        for fi13 := 0; fi13 < 16; fi13++ {
            v9 -= (v9 >> ((42) & 63))
            v14 := GenS4{F0: 65535, F1: 2, F2: 774}
            var v15 float32 = 1e10
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        if (2147483647 < 4294967295) {
            for fi16 := 0; fi16 < 3; fi16++ {
                if ((753 < -32768) || (!false)) {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS5{F0: -1}
                        v9 += int32(fi17)
                        _ = fi17
                        _ = v18
                    }
                    var v19 uint16 = 473
                    v20 := GenS5{F0: -1}
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS2{F0: -1, F1: 2147483647, F2: -2147483648}
                    _ = v21
                }
                _ = fi16
            }
            if (42 <= 32768) {
                v9 -= -39
                var v22 int = 0
                _ = v22
            }
            switch "hello" {
            case "foo bar":
                if true {
                    if (!false) {
                        v9 += v9
                    } else {
                        var v23 float64 = 1.5
                        for fi24 := 0; fi24 < 2; fi24++ {
                            fmt.Println(v23)
                            var v25 string = ("" + (("foo bar" + "\n\t") + ""))
                            fmt.Println(v23)
                            _ = fi24
                            _ = v25
                        }
                        if ((69 == v9) && (!(!(1 != 0)))) {
                            var v26 uint16 = 837
                            _ = v26
                        } else {
                            if ((-470 == 0) || true) {
                                delete(v2, (("a" + ("a" + "a")) + ("\n\t" + "a")))
                                var v27 uint = 42
                                v28 := v4[-1]
                                _ = v27
                                _ = v28
                            } else {
                                v2[("" + "hello")] = (42 >= 1000)
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    var v30 float32 = 1e10
                                    _ = fi29
                                    _ = v30
                                }
                                v31 := GenS3{F0: -261, F1: 4294967295, F2: 0}
                                _ = v31
                            }
                        }
                        _ = v23
                    }
                    v2["foo bar"] = (!true)
                } else {
                    if (false || (!((!((-888 >= 0) && (32767 <= 484))) && ((false && (127 <= 46)) || false)))) {
                        v9 = (v9 / ((-2147483648) | 1))
                    }
                    var v32 uint = 255
                    if (127 == 255) {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            v32 = (v32 % (((v32 | (v32 << ((v32) & 63)))) | 1))
                            switch (fi33 | 910) {
                            case -9223372036854775808:
                                v32 = v32
                            case 759:
                                fmt.Println(fi33, v32, v32)
                            case 127:
                                if (32767 == -32768) {
                                    v34 := GenS2{F0: 1, F1: 42, F2: -657}
                                    switch (("hello" + "hello") + "a") {
                                    case "a":
                                        v35 := []int{(fi33 - (fi33 >> (((v32 * v32)) & 63))), fi33}
                                        _ = v35
                                    }
                                    v9 = v9
                                    _ = v34
                                }
                            default:
                                var v36 float32 = ((3.14159 * 1e10) - 1.5)
                                _ = v36
                            }
                            fmt.Println(v32)
                            _ = fi33
                        }
                    } else {
                        for fi37 := 0; fi37 < 15; fi37++ {
                            v32 -= v32
                            fmt.Println(v32, fi37)
                            if true {
                                switch "hello" {
                                case "foo bar":
                                    if true {
                                        v32 = (v32 * (v32 - 4294967295))
                                        var v38 uint32 = genFn0("foo bar", 1, (fi37 / ((fi37) | 1)))
                                        _ = v38
                                    }
                                default:
                                    v39 := []int{}
                                    _ = v39
                                }
                            }
                            _ = fi37
                        }
                        for fi40 := 0; fi40 < 4; fi40++ {
                            v9 -= (-(v9))
                            v41 := []int{(fi40 | (fi40 & fi40)), fi40, 5}
                            fmt.Println(fi40, v32)
                            _ = fi40
                            _ = v41
                        }
                        if (!((127 <= 255) && (!(!((0.1 / ((0.1 / 1.5) * (-2.25 - ((1e10 * (((1e10 * ((0.0 + float64(-2.25)) * (1e10 * (0.1 + 1.5)))) / float64(0.1)) - (3.14159 + (((-2.25 * 3.14159) * (1e10 - 1.5)) - (0.1 / -2.25))))) / (1e10 - 0.0))))) > -2.25))))) {
                            fmt.Println(v9)
                        } else {
                            if false {
                                v42 := GenS2{F0: 100, F1: 2, F2: 0}
                                v9 = v9
                                delete(v5, 1)
                                _ = v42
                            } else {
                                var v43 uint = v32
                                fmt.Println(v32, v43)
                                switch "" {
                                case "foo bar":
                                    var v44 uint16 = 65535
                                    _ = v44
                                case "a":
                                    v45 := []int{}
                                    _ = v45
                                default:
                                    v46 := GenS6{F0: 42, F1: 391}
                                    _ = v46
                                }
                                _ = v43
                            }
                            var v47 float64 = float64(3.14159)
                            _ = v47
                        }
                    }
                    _ = v32
                }
            default:
                var v48 int = 42
                _ = v48
            }
        } else {
            v49 := []int{760, 4294967295, 42, -497}
            v9 -= -2147483648
            _ = v49
        }
    }
    var v50 int32 = (v9 << ((687) & 63))
    var v51 uint32 = 2
    var v52 uint64 = 2
    v53 := v3[("\n\t" + "\n\t")]
    v52 = 174
    v50 += 2147483647
    var v54 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v9
}

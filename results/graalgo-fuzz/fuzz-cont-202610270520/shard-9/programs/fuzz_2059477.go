package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 int32) uint16 {
    v4 := []int{32767, 128, -1, 32767}
    if false {
        p0 = 0
    } else {
        if false {
            if ((0 > 4294967295) && true) {
                v5 := GenS1{F0: 2}
                switch "foo bar" {
                case "a":
                    fmt.Println(p0, p0)
                case "":
                    var v6 int32 = 424
                    _ = v6
                case "\n\t":
                    for fi7 := 0; fi7 < 14; fi7++ {
                        for fi8 := 0; fi8 < 19; fi8++ {
                            if (false || (40 <= 658)) {
                                v9 := GenS1{F0: 255}
                                var v10 int64 = 800
                                fmt.Println(p0, p1, p1)
                                _ = v10
                                _ = v9
                            } else {
                                _ = len(v4)
                                p1 += p1
                                p1 += (p0 + p0)
                            }
                            _ = fi8
                        }
                        for fi11 := 0; fi11 < 18; fi11++ {
                            switch "foo bar" {
                            case "a":
                                if true {
                                    var v12 float64 = -2.25
                                    v12 -= (float64(3.14159) * (1.5 / 0.1))
                                    _ = len(v4)
                                    _ = v12
                                }
                            case "foo bar":
                                var v13 uint16 = 1
                                _ = v13
                            }
                            _ = fi11
                        }
                        if ((true || (!true)) && false) {
                            var v14 float64 = (0.0 * -2.25)
                            var v15 bool = (891 == 1)
                            _ = v14
                            _ = v15
                        }
                        _ = fi7
                    }
                default:
                    fmt.Println(p1, p1)
                }
                _ = v5
            } else {
                switch "foo bar" {
                case "\n\t":
                    v16 := GenS2{F0: 9223372036854775807}
                    _ = v16
                case "a":
                    p0 += p1
                }
            }
            switch "a" {
            case "hello":
                var v17 int = 903
                _ = v17
            default:
                var v18 uint16 = 42
                _ = v18
            }
            p1 -= (p1 * (p0 ^ 1))
        } else {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 string = "a"
                _ = fi19
                _ = v20
            }
            var v21 bool = true
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 31
}
func genFn1() uint64 {
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := []int{fi22, fi22, fi22, (fi22 + (fi22 * 32767))}
        for fi24 := 0; fi24 < 15; fi24++ {
            if false {
                var v25 int32 = 1
                var v26 int = fi24
                _ = v25
                _ = v26
            }
            _ = fi24
        }
        fmt.Println(fi22)
        _ = fi22
        _ = v23
    }
    v27 := GenS3{F0: -609}
    var v28 uint8 = 2
    _ = v27
    _ = v28
    return 4294967295
}

func main() {
    var v1 bool = (true || true)
    var v2 int64 = -283
    var v3 uint32 = 0
    if (42 <= 0) {
        v1 = v1
    }
    if v1 {
        v2 += (-(v2))
        var v29 int8 = -128
        _ = v29
    } else {
        var v30 uint8 = 142
        v31 := GenS2{F0: (v2 << ((0) & 63))}
        _ = v30
        _ = v31
    }
    v2 = v2
    v3 += (v3 * 2)
    v3 = (v3 % ((v3) | 1))
    if ((193 == 18446744073709551615) && (!true)) {
        if true {
            var v32 uint16 = 414
            v3 += (^(v3))
            _ = v32
        }
        v1 = v1
    } else {
        fmt.Println(v1, v3)
        var v33 uint32 = (v3 >> ((2) & 63))
        _ = v33
    }
    var v34 int32 = 747
    v35 := []int{}
    fmt.Println(v2, v34)
    var v36 int = 32768
    if (!false) {
        v35 = append(v35, v36)
    }
    switch "a" {
    case "hello":
        if true {
            switch ("" + "\n\t") {
            case "a":
                if (!true) {
                    var v37 int16 = 0
                    _ = v37
                } else {
                    fmt.Println(v1, v1)
                }
            }
            v3 += (v3 % (((v3 & (v3 + 4294967295))) | 1))
            switch 363 {
            case -1:
                _ = len(v35)
            default:
                var v38 string = ((("foo bar" + "hello") + "hello") + "hello")
                _ = v38
            }
        } else {
            var v39 int32 = (v34 ^ (v34 ^ -2147483648))
            _ = v39
        }
    case "a":
        v34 += 0
    case "":
        var v40 bool = (948 != 65535)
        _ = v40
    default:
        var v41 string = (("hello" + "foo bar") + "")
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}

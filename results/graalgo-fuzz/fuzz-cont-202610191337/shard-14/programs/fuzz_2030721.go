package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 uint64, p1 bool) int {
    p0 = (p0 / (((^(p0))) | 1))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, (fi7 | fi7), -1, fi7}
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 -= 486
            _ = fi9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            _ = len(v8)
            p0 = (p0 * p0)
            p0 = (p0 ^ p0)
            fmt.Println(p0)
            _ = fi10
        }
        _ = fi7
        _ = v8
    }
    if ((p0 > 36) && p1) {
        var v11 string = ""
        _ = v11
    }
    var v12 bool = p1
    v13 := GenS1{F0: 0, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -9223372036854775808
}
func genFn1(p0 uint64) bool {
    v14 := GenS1{F0: 0, F1: 0}
    v15 := GenS1{F0: 1, F1: 0}
    v16 := []int{genFn0((p0 / ((p0) | 1)), false), -9223372036854775808, 1}
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return false
}

func main() {
    var v1 bool = (32767 == 65535)
    v2 := make(map[string]uint32)
    var v3 int32 = -714
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v6 int16 = 230
    var v17 bool = true
    fmt.Println(v1, v6)
    fmt.Println(v1, v3, v6)
    v4[42] = ("foo bar" + ("" + "hello"))
    for fi18 := 0; fi18 < 17; fi18++ {
        switch fi18 {
        case 4294967295:
            if v17 {
                v19 := GenS2{F0: 93, F1: ("" + "hello"), F2: -1}
                _ = v19
            } else {
                v20 := GenS3{F0: "\n\t"}
                _ = v20
            }
        case 9223372036854775807:
            fmt.Println(v6, v3)
        case -1:
            var v21 int16 = -1
            _ = v21
        }
        _ = fi18
    }
    v17 = true
    if v1 {
        v22 := GenS4{F0: 32767, F1: 0}
        var v23 int32 = -2147483648
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS5{F0: 32767, F1: "hello", F2: 18446744073709551615}
                _ = fi25
                _ = v26
            }
            switch (("foo bar" + "foo bar") + ("hello" + "\n\t")) {
            case "\n\t":
                var v27 int32 = 0
                _ = v27
            case "":
                fmt.Println(fi24, v3, fi24)
            case "foo bar":
                for fi28 := 0; fi28 < 5; fi28++ {
                    if (((fi28 ^ (fi24 >> ((65535) & 63))) <= genFn0(0, v1)) || true) {
                        switch ("" + ("foo bar" + ("\n\t" + ("a" + "foo bar")))) {
                        case "a":
                            fmt.Println(fi24, v3, v17)
                        default:
                            v3 = (v3 >> ((2) & 63))
                        }
                        v29 := []int{fi24}
                        _ = v29
                    }
                    if false {
                        if (0.1 > 0.0) {
                            if (!v1) {
                                switch (fi28 * genFn0(18446744073709551615, v17)) {
                                case 2147483647:
                                    var v30 int64 = 0
                                    _ = v30
                                case 42:
                                    v3 = v3
                                default:
                                    fmt.Println(v3, v1, fi28)
                                }
                                v3 = 0
                                var v31 uint32 = 80
                                _ = v31
                            }
                        } else {
                            var v32 uint = 249
                            if v17 {
                                var v33 int8 = 1
                                _ = v33
                            } else {
                                if (false || false) {
                                    var v34 float64 = ((0.0 + (1.5 / (3.14159 - (float64(0.1) * float64(-2.25))))) / 3.14159)
                                    v35 := v4[(fi28 & (fi24 >> (((v32 + v32)) & 63)))]
                                    _ = v34
                                    _ = v35
                                }
                                var v36 float32 = 0.1
                                if v17 {
                                    v17 = true
                                    v3 = -2147483648
                                } else {
                                    fmt.Println(v1, fi28, v1)
                                    var v37 int16 = 0
                                    _ = v37
                                }
                                _ = v36
                            }
                            var v38 int16 = (v6 | (v6 >> ((65535) & 63)))
                            var v39 int8 = 86
                            _ = v32
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS2{F0: 1, F1: "foo bar", F2: -38}
                        _ = v40
                    } else {
                        var v41 float32 = (1.5 / 1e10)
                        _ = v41
                    }
                    _ = fi28
                }
            default:
                if v17 {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        fmt.Println(v1, fi24)
                        fmt.Println(v3)
                        _ = fi42
                    }
                    if v1 {
                        for fi43 := 0; fi43 < 15; fi43++ {
                            fmt.Println(fi43)
                            if (!v1) {
                                var v44 uint32 = 2147483647
                                var v45 uint8 = 117
                                _ = v44
                                _ = v45
                            }
                            _ = fi43
                        }
                        for fi46 := 0; fi46 < 15; fi46++ {
                            var v47 uint64 = 18446744073709551615
                            delete(v4, (fi24 ^ fi24))
                            _ = fi46
                            _ = v47
                        }
                    }
                }
            }
            v6 -= (v6 / ((v6) | 1))
            _ = fi24
        }
        if (127 < 128) {
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 uint16 = 1
                _ = fi48
                _ = v49
            }
            for fi50 := 0; fi50 < 11; fi50++ {
                for fi51 := 0; fi51 < 18; fi51++ {
                    v3 -= (v3 ^ v3)
                    _ = fi51
                }
                for fi52 := 0; fi52 < 3; fi52++ {
                    var v53 uint = 255
                    _ = fi52
                    _ = v53
                }
                _ = fi50
            }
            v54 := GenS5{F0: 1, F1: ("\n\t" + "hello"), F2: 1}
            _ = v54
        }
        v6 -= v6
    }
    var v55 uint8 = 127
    v56 := GenS4{F0: genFn0(0, v17), F1: 42}
    var v57 int64 = -252
    v58 := GenS6{F0: v6, F1: 605, F2: 18446744073709551615}
    var v59 int16 = 1
    if v17 {
        var v60 float32 = (0.1 - 0.1)
        var v61 int16 = v59
        _ = v60
        _ = v61
    } else {
        for fi62 := 0; fi62 < 6; fi62++ {
            v63 := GenS7{F0: v57, F1: int(v55)}
            _ = fi62
            _ = v63
        }
        for fi64 := 0; fi64 < 14; fi64++ {
            v57 = int64(v55)
            _ = fi64
        }
    }
    v65 := []int{}
    if len(v65) >= 2 {
        v65 = v65[1:len(v65)]
    }
    v3 -= -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v65
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint16; F1 string }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    v7 := []int{0}
    var v8 int = -106
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := []string{(("\n\t" + "a") + "a"), "", "foo bar"}
    var v2 bool = false
    var v3 bool = (42 != 1)
    var v4 int64 = -9223372036854775808
    var v5 uint16 = 399
    var v6 int64 = -1
    v5 = (v5 | (v5 + (v5 << ((4294967295) & 63))))
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = 101
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v6, fi11)
            var v12 uint = 1
            var v13 int8 = v10
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(fi9)
            }
            _ = fi14
        }
        _ = fi9
        _ = v10
    }
    var v15 uint = 4294967295
    var v16 int8 = -14
    if len(v1) > 0 {
        v1[0] = ("a" + ("\n\t" + ((((("" + ("\n\t" + ("a" + ("" + "hello")))) + "\n\t") + "a") + "") + "a")))
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int8 = v16
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        switch (("\n\t" + "foo bar") + "\n\t") {
        case "hello":
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS1{F0: -201, F1: v5, F2: 1}
                _ = fi20
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 19; fi22++ {
                v4 = 42
                var v23 uint32 = 664
                _ = fi22
                _ = v23
            }
        }
        var v24 uint8 = 128
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS2{F0: 2147483647, F1: false, F2: int64(v24)}
            if v3 {
                var v27 uint8 = 255
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 += (v4 - (v4 * (v6 / ((v4) | 1))))
                _ = v27
            }
            v28 := GenS3{F0: 100}
            _ = fi25
            _ = v26
            _ = v28
        }
        _ = fi19
        _ = v24
    }
    if true {
        v6 += -1
        v29 := GenS4{F0: 321, F1: 2}
        _ = v29
    }
    v16 += (v16 - (v16 + int8(v5)))
    v30 := GenS5{F0: -1, F1: 513}
    if true {
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 uint32 = 4294967295
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 float64 = (((0.1 * (0.0 + -2.25)) + -2.25) - 3.14159)
                _ = fi33
                _ = v34
            }
            var v35 int = (fi31 | 32767)
            _ = fi31
            _ = v32
            _ = v35
        }
        var v36 int16 = int16(v4)
        switch int(v5) {
        case 2147483647:
            if true {
                for fi37 := 0; fi37 < 11; fi37++ {
                    _ = len(v1)
                    if ((v36 ^ -32768) < -1) {
                        v38 := GenS3{F0: (v16 << ((v15) & 63))}
                        v1 = append(v1, ("" + "hello"))
                        for fi39 := 0; fi39 < 2; fi39++ {
                            v40 := GenS2{F0: -2147483648, F1: (!v3), F2: (v4 * (v6 << (((v15 << (((v15 % ((65535) | 1))) & 63))) & 63)))}
                            switch "hello" {
                            case "\n\t":
                                v5 = v5
                            case "a":
                                v2 = v3
                            case "foo bar":
                                fmt.Println(v3, fi39)
                            default:
                                fmt.Println(v15, v6, fi39)
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = v38
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if v3 {
                            _ = len(v1)
                            v5 += 32767
                        } else {
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 uint64 = 1
                                v1 = append(v1, "hello")
                                _ = fi41
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 20; fi43++ {
                                v15 = 621
                                if ((v2 || false) && v2) {
                                    v36 += 356
                                    var v44 int16 = (v36 | v36)
                                    _ = len(v1)
                                    _ = v44
                                } else {
                                    v45 := GenS3{F0: v16}
                                    if v2 {
                                        v46 := GenS6{F0: (v5 ^ 32767), F1: "hello"}
                                        _ = v46
                                    }
                                    v1 = append(v1, ("foo bar" + ((("a" + ("\n\t" + "")) + "hello") + "foo bar")))
                                    _ = v45
                                }
                                _ = fi43
                            }
                        }
                    }
                    _ = fi37
                }
            } else {
                var v47 uint8 = 127
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v47
            }
        case 127:
            if len(v1) > 0 {
                v1[0] = ("\n\t" + (("\n\t" + "a") + ("foo bar" + (("foo bar" + "\n\t") + "a"))))
            }
        default:
            switch genFn0(0) {
            case 4294967295:
                v36 += v36
            default:
                for fi48 := 0; fi48 < 2; fi48++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi48
                }
            }
        }
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

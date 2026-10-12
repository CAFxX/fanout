package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8) uint64 {
    for fi4 := 0; fi4 < 20; fi4++ {
        if (-128 >= 42) {
            if (389 >= 4294967295) {
                if (!true) {
                    p0 = (^(p0))
                }
            }
            v5 := []int{}
            if (2 != 862) {
                switch 127 {
                case 42:
                    v6 := GenS1{F0: 255, F1: "foo bar"}
                    _ = v6
                }
            }
            _ = v5
        }
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{(^(fi4)), fi4, (fi7 ^ fi4)}
            p0 -= p0
            fmt.Println(p0, fi7, fi4)
            _ = fi7
            _ = v8
        }
        _ = fi4
    }
    fmt.Println(p0)
    if ((0 > 0) && (-1 < 32767)) {
        p0 = p0
        fmt.Println(p0)
        p0 = (p0 * p0)
    }
    _ = p0
    return 42
}

func main() {
    v1 := []int{-17}
    var v2 uint32 = 1
    var v3 uint = 42
    if true {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: v3, F1: "\n\t"}
            if (!(!true)) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    v2 = uint32(v3)
                    _ = fi11
                }
            }
            var v12 int32 = 312
            _ = fi9
            _ = v10
            _ = v12
        }
    }
    switch "" {
    case "hello":
        v3 -= (v3 - v3)
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 0
            var v15 int64 = -188
            var v16 string = ("foo bar" + (("\n\t" + ((((("" + ((("foo bar" + "a") + ("hello" + "\n\t")) + (("foo bar" + ("\n\t" + ("a" + "hello"))) + ("a" + "a")))) + ("hello" + "\n\t")) + "a") + "hello") + ("\n\t" + ("\n\t" + "hello")))) + "hello"))
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    case "foo bar":
        switch "" {
        case "a":
            v3 = (v3 >> ((v3) & 63))
        case "foo bar":
            if false {
                switch int(v2) {
                case -82:
                    if true {
                        var v17 int64 = -597
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if (genFn0(uint8(v3)) > 2) {
                                v19 := GenS1{F0: 65535, F1: "hello"}
                                v2 = v2
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    var v21 uint32 = (v2 - (v2 - v2))
                                    _ = fi20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                v17 = (v17 % ((v17) | 1))
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi18 << ((479) & 63))
                            }
                            var v22 uint8 = 0
                            _ = fi18
                            _ = v22
                        }
                        v23 := GenS1{F0: v3, F1: ""}
                        _ = v17
                        _ = v23
                    }
                case 958:
                    switch 42 {
                    case 9223372036854775807:
                        if (genFn0(85) <= 685) {
                            if (-20 > -45) {
                                v24 := GenS2{F0: "", F1: 0, F2: 100}
                                if ((2147483647 == int(v3)) || true) {
                                    if len(v1) > 0 {
                                        v1[0] = -1
                                    }
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        v3 = (v3 / (((^(v3))) | 1))
                                        v2 = v2
                                        v3 = v3
                                        _ = fi25
                                    }
                                    var v26 uint32 = (^(v2))
                                    _ = v26
                                }
                                _ = v24
                            } else {
                                v2 = 886
                                var v27 int32 = -2147483648
                                _ = v27
                            }
                        }
                    case 869:
                        if (1 == -504) {
                            if (42 <= uint64(v3)) {
                                _ = len(v1)
                                if (-1 <= 2147483647) {
                                    v28 := GenS3{F0: 4294967295}
                                    var v29 int64 = 0
                                    _ = v28
                                    _ = v29
                                } else {
                                    v3 = (v3 ^ v3)
                                }
                                v3 -= 42
                            }
                        }
                    }
                }
                var v30 int16 = 657
                if (v2 > uint32(v3)) {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        fmt.Println(v3, fi31, v3)
                        _ = fi31
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v30
            } else {
                if true {
                    v3 = v3
                    fmt.Println(v2, v2)
                } else {
                    var v32 float64 = float64(3.14159)
                    _ = v32
                }
                v33 := GenS3{F0: 0}
                _ = v33
            }
        default:
            v34 := GenS4{F0: 4294967295, F1: 255}
            _ = v34
        }
    }
    fmt.Println(v2, v3)
    if (-1 <= -32768) {
        v35 := GenS1{F0: (v3 & v3), F1: ("\n\t" + ("" + "\n\t"))}
        var v36 uint32 = (v2 - (^(v2)))
        _ = v35
        _ = v36
    } else {
        v3 += v3
        v2 = (v2 / ((1) | 1))
    }
    v3 -= v3
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 uint64 = uint64(v2)
        var v39 int32 = -642
        _ = fi37
        _ = v38
        _ = v39
    }
    v3 = (v3 & v3)
    v1 = append(v1, -1)
    for fi40 := 0; fi40 < 3; fi40++ {
        v2 += uint32(v3)
        for fi41 := 0; fi41 < 11; fi41++ {
            var v42 uint8 = 42
            _ = fi41
            _ = v42
        }
        v43 := GenS5{F0: 0, F1: 137}
        _ = fi40
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}

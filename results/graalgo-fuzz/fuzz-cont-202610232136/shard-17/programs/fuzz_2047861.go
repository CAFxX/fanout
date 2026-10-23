package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int16; F2 int }
type GenS5 struct { F0 bool; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int }
func genFn0(p0 int64, p1 bool) uint8 {
    if false {
        var v5 float32 = (1.5 * 3.14159)
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: -107}
        fmt.Println(fi6)
        if p1 {
            v8 := GenS2{F0: 100, F1: 42}
            _ = v8
        } else {
            fmt.Println(p1, p0)
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return 223
}
func genFn1(p0 uint16, p1 uint64, p2 uint) uint16 {
    if (128 >= uint8(p1)) {
        v9 := GenS3{F0: ("\n\t" + ("\n\t" + "\n\t"))}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(p2)
            case "hello":
                v11 := []int{fi10, fi10, fi10, fi10}
                _ = v11
            }
            switch ("foo bar" + "") {
            case "":
                if (!true) {
                    p2 = (p2 + (p2 >> ((493) & 63)))
                } else {
                    var v12 uint8 = 128
                    var v13 uint32 = 2147483648
                    _ = v12
                    _ = v13
                }
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    if false {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            var v15 uint64 = p1
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        v16 := GenS3{F0: ("a" + (("" + "") + "a"))}
                        _ = v16
                    }
                default:
                    var v17 uint16 = 111
                    _ = v17
                }
            default:
                var v18 uint8 = 0
                _ = v18
            }
            _ = fi10
        }
    }
    if ("" < "\n\t") {
        v19 := GenS1{F0: 39}
        var v20 int8 = 0
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 int = 32767
    v1 = (v1 * v1)
    v21 := []int{(v4 - (v4 - 2147483647)), (v4 / ((v4) | 1)), (v4 + -1), v4}
    for fi22 := 0; fi22 < 11; fi22++ {
        if false {
            switch "foo bar" {
            case "\n\t":
                var v23 int64 = 743
                _ = v23
            default:
                if false {
                    if len(v21) > 0 {
                        v21[0] = (v4 * v4)
                    }
                    v1 -= v1
                    switch "foo bar" {
                    case "":
                        if (!true) {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                fmt.Println(fi24)
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    v1 = v1
                                case "a":
                                    v21 = append(v21, (-(fi22)))
                                case "\n\t":
                                    fmt.Println(v1, v4, fi22)
                                }
                                _ = fi24
                            }
                            fmt.Println(fi22, v1)
                            if ((!((!false) || ((((v4 | fi22) < fi22) && (true && false)) && false))) || false) {
                                var v25 float32 = (float32((((-2.25 / (float64(3.14159) - 1.5)) - 0.0) - ((1.5 - 0.0) * (0.1 - 1.5)))) + 1e10)
                                _ = v25
                            } else {
                                fmt.Println(v1)
                                var v26 float64 = (0.0 + 0.0)
                                _ = v26
                            }
                        }
                    case "\n\t":
                        v1 = (v1 & (v1 & 127))
                    case "a":
                        var v27 uint = 4294967295
                        _ = v27
                    }
                } else {
                    var v28 int8 = 127
                    if true {
                        if (2147483647 <= 41) {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                _ = len(v3)
                                v1 += 97
                                var v30 int32 = 0
                                _ = fi29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 20; fi31++ {
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                _ = fi31
                            }
                        }
                    } else {
                        if ((v4 & fi22) >= v4) {
                            var v32 bool = true
                            var v33 string = ("hello" + "\n\t")
                            _ = v32
                            _ = v33
                        }
                        var v34 int32 = -2147483648
                        var v35 uint32 = 694
                        _ = v34
                        _ = v35
                    }
                    _ = v28
                }
            }
        }
        fmt.Println(v4)
        _ = fi22
    }
    if (("" <= "\n\t") && (true && (!(!true)))) {
        switch v4 {
        case 1:
            var v36 int = v4
            _ = v36
        }
    }
    v37 := GenS1{F0: -128}
    var v38 int64 = -1
    if ("\n\t" != "foo bar") {
        fmt.Println(v1, v38)
        switch (v4 << ((42) & 63)) {
        case -965:
            switch v4 {
            case -659:
                v38 += v38
            case 42:
                v39 := GenS4{F0: "\n\t", F1: -1, F2: int(v1)}
                _ = v39
            default:
                v1 = (v1 ^ 1)
            }
        case 1:
            var v40 float64 = float64(1.5)
            _ = v40
        case 32768:
            if (2147483647 < 277) {
                v2[381] = 881
                if false {
                    var v41 bool = false
                    var v42 int = -2147483648
                    v42 = 387
                    _ = v41
                    _ = v42
                }
                v43 := GenS5{F0: (0 != (v38 + v38)), F1: 814, F2: -2147483648}
                _ = v43
            } else {
                var v44 uint16 = 794
                for fi45 := 0; fi45 < 18; fi45++ {
                    v46 := GenS6{F0: v38, F1: (v1 & genFn0(348, (0.1 >= 0.0)))}
                    _ = fi45
                    _ = v46
                }
                _ = v44
            }
        }
        v47 := GenS7{F0: (v4 / ((v4) | 1))}
        _ = v47
    } else {
        if (uint64(v1) >= 18446744073709551615) {
            var v48 int8 = -15
            _ = v48
        }
        v4 = (v4 >> ((1) & 63))
    }
    v38 = (v38 + v38)
    v1 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}

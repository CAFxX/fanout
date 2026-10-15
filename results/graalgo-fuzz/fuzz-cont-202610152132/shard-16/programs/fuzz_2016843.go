package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
func genFn0(p0 int16, p1 int16, p2 int16) string {
    p1 -= (^(p2))
    if false {
        v5 := []int{}
        switch 0 {
        case 0:
            v6 := GenS1{F0: 0, F1: 1, F2: 0}
            _ = v6
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v5
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 int64 = -152
                    p1 = -583
                    _ = fi8
                    _ = v9
                }
                if false {
                    fmt.Println(p1, p2, p1)
                } else {
                    if ((true && true) || false) {
                        p1 = (p0 >> ((178) & 63))
                    }
                }
            } else {
                p0 = p1
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int64 = -9223372036854775808
                    p1 -= (p0 + 32767)
                    _ = fi10
                    _ = v11
                }
                var v12 uint = 65535
                _ = v12
            }
            switch ("a" + "foo bar") {
            case "foo bar":
                for fi13 := 0; fi13 < 16; fi13++ {
                    fmt.Println(p1)
                    _ = fi13
                }
            default:
                switch ("\n\t" + "foo bar") {
                case "":
                    for fi14 := 0; fi14 < 4; fi14++ {
                        var v15 int32 = 1
                        v16 := GenS1{F0: 42, F1: -1, F2: 179}
                        fmt.Println(fi14, p2)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case "a":
                    v17 := GenS2{F0: 151, F1: 953}
                    _ = v17
                case "\n\t":
                    var v18 float64 = 0.0
                    _ = v18
                default:
                    var v19 uint32 = 1
                    _ = v19
                }
            }
            v20 := GenS2{F0: 899, F1: 65535}
            _ = fi7
            _ = v20
        }
        var v21 float32 = -2.25
        _ = v21
    }
    fmt.Println(p0, p2, p0)
    fmt.Println(p0, p1)
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{}
    v3 := []uint32{4294967295, 1, 2147483648}
    v4 := make(map[string]string)
    v22 := GenS2{F0: 18446744073709551615, F1: 549}
    for fi23 := 0; fi23 < 11; fi23++ {
        if len(v3) > 0 {
            v3[0] = 2
        }
        if false {
            var v24 float64 = 1.5
            var v25 int16 = 1000
            _ = v24
            _ = v25
        }
        v3 = append(v3, 2147483647)
        _ = fi23
    }
    var v26 int16 = 32767
    v26 = (v26 & (v26 & (v26 % (((v26 - (v26 | (v26 * (v26 & 0))))) | 1))))
    v27 := GenS3{F0: 188, F1: -1, F2: 213}
    v26 -= 42
    _ = len(v4)
    var v28 uint16 = 795
    switch ("a" + "foo bar") {
    case "\n\t":
        v26 = (^(v26))
    default:
        switch ("\n\t" + (("hello" + "\n\t") + (("hello" + "hello") + (("a" + ("a" + "a")) + "a")))) {
        case "foo bar":
            var v29 uint32 = 42
            _ = v29
        case "":
            v28 = v28
        case "hello":
            var v30 string = ("" + ("" + "\n\t"))
            _ = v30
        }
    }
    if (v28 != 656) {
        for fi31 := 0; fi31 < 2; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                v28 += v28
                var v33 float64 = -2.25
                var v34 uint16 = v28
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 uint16 = 65535
            v36 -= 42
            for fi37 := 0; fi37 < 10; fi37++ {
                if false {
                    var v38 int64 = 42
                    _ = v38
                }
                v36 += v36
                if true {
                    v26 = v26
                } else {
                    if (false || true) {
                        switch "a" {
                        case "a":
                            if true {
                                var v39 int8 = 42
                                switch "foo bar" {
                                case "foo bar":
                                    v26 -= (v26 | int16(v39))
                                case "":
                                    v40 := GenS3{F0: 296, F1: 42, F2: 2147483648}
                                    _ = v40
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v41 := GenS2{F0: 1, F1: 255}
                                        _ = v41
                                    case "hello":
                                        if (!(false || true)) {
                                            v42 := v4["hello"]
                                            _ = v42
                                        }
                                    default:
                                        v39 += v39
                                    }
                                default:
                                    v26 -= v26
                                }
                                _ = v39
                            } else {
                                delete(v4, "")
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v43 string = "hello"
                                if (2147483647 != -419) {
                                    v28 += v36
                                    v26 = v26
                                } else {
                                    var v44 int64 = 1
                                    _ = v44
                                }
                                _ = v43
                            }
                        case "\n\t":
                            v45 := GenS4{F0: (("a" + (("hello" + (("" + "") + "a")) + "foo bar")) + "\n\t")}
                            _ = v45
                        default:
                            switch 2147483647 {
                            case 127:
                                var v46 uint16 = 349
                                v47 := GenS5{F0: ((!(774 <= 2)) || false)}
                                v48 := GenS6{F0: 0}
                                v36 += v28
                                _ = v46
                                _ = v47
                                _ = v48
                            case 0:
                                v26 -= (v26 << ((42) & 63))
                            default:
                                v28 = v36
                            }
                        }
                    }
                    v26 = v26
                }
                _ = fi37
            }
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}

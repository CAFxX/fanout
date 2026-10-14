package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 int16 }
func genFn0(p0 bool, p1 uint64, p2 uint16) uint {
    fmt.Println(p0, p2, p2)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v7 := GenS1{F0: 157, F1: 127, F2: 532}
        _ = v7
    case "":
        if (!(42 < 9223372036854775807)) {
            v8 := GenS1{F0: 127, F1: -111, F2: -850}
            var v9 uint16 = uint16(p1)
            if p0 {
                for fi10 := 0; fi10 < 2; fi10++ {
                    if p0 {
                        v11 := GenS2{F0: (p1 % (((p1 >> ((0) & 63))) | 1))}
                        v12 := []int{}
                        switch "hello" {
                        case "":
                            var v13 uint8 = 0
                            _ = v13
                        default:
                            v9 += p2
                        }
                        _ = v11
                        _ = v12
                    }
                    fmt.Println(p2, v9, p0)
                    if p0 {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            fmt.Println(fi10)
                            v9 += 0
                            _ = fi14
                        }
                        fmt.Println(v9, p1, p1)
                    }
                    _ = fi10
                }
                fmt.Println(p1, p0)
            }
            _ = v8
            _ = v9
        } else {
            v15 := []int{-666, 2147483647}
            _ = v15
        }
    case "\n\t":
        p2 -= (p2 & (p2 << ((282) & 63)))
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{(fi16 >> ((398) & 63)), (fi16 % ((2147483647) | 1))}
            switch fi16 {
            case -2147483648:
                var v18 int64 = 9223372036854775807
                _ = v18
            }
            var v19 int8 = 100
            _ = fi16
            _ = v17
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 884
}
func genFn1() uint8 {
    var v20 int16 = 32767
    if (float64(float32(-2.25)) <= 1.5) {
        v21 := GenS3{F0: -124}
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 string = ("hello" + ("a" + ("" + "")))
            if (true && false) {
                v24 := []int{1}
                fmt.Println(v20, v23, fi22)
                _ = v24
            } else {
                var v25 uint32 = 2
                v26 := []int{32767, -1}
                var v27 int = -1
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            v20 = 32767
            if (!("\n\t" != (("\n\t" + "hello") + "foo bar"))) {
                v20 = v20
            }
            switch fi28 {
            case 42:
                var v29 int64 = -9223372036854775808
                _ = v29
            default:
                var v30 uint32 = 65535
                _ = v30
            }
            _ = fi28
        }
        _ = v21
    } else {
        v31 := []int{940}
        _ = v31
    }
    if false {
        switch 2147483647 {
        case -330:
            if true {
                if false {
                    v20 -= (v20 * (v20 % (((v20 ^ (v20 + (v20 >> ((genFn0(false, 42, 703)) & 63))))) | 1)))
                    v20 = -32768
                }
            } else {
                if (!false) {
                    if (!((false || false) && true)) {
                        v32 := []int{}
                        _ = v32
                    }
                    v33 := []int{}
                    var v34 int16 = (v20 >> ((4294967295) & 63))
                    _ = v33
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v20 += (-(v20))
                        v20 -= (-(v20))
                        _ = fi35
                    }
                    var v36 uint8 = 1
                    _ = v36
                }
            }
        default:
            v37 := []int{42, 2147483647}
            _ = v37
        }
    }
    if true {
        fmt.Println(v20)
        if ("foo bar" == "") {
            var v38 uint16 = 42
            _ = v38
        } else {
            switch 9223372036854775807 {
            case 1:
                var v39 int64 = 42
                _ = v39
            case 127:
                switch "\n\t" {
                case "\n\t":
                    if (false && true) {
                        v20 = 32767
                        fmt.Println(v20)
                    }
                }
            case 9223372036854775807:
                for fi40 := 0; fi40 < 9; fi40++ {
                    var v41 int16 = 1
                    fmt.Println(v20, v20, v20)
                    for fi42 := 0; fi42 < 11; fi42++ {
                        v20 = (^(v41))
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
            default:
                for fi43 := 0; fi43 < 16; fi43++ {
                    if (true && ((fi43 | (fi43 / ((fi43) | 1))) >= (fi43 << ((310) & 63)))) {
                        v20 -= (v20 & (v20 * v20))
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            var v44 int32 = -2147483648
                            _ = v44
                        }
                        v45 := []int{128, (fi43 * fi43), (fi43 << ((42) & 63))}
                        _ = v45
                    }
                    if true {
                        v46 := GenS3{F0: -128}
                        _ = v46
                    }
                    _ = fi43
                }
            }
        }
    } else {
        var v47 uint8 = 85
        _ = v47
    }
    _ = v20
    return 2
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int32 = -2147483648
    var v4 int8 = 42
    var v5 string = (("" + ("a" + "")) + "a")
    var v6 int = 2147483647
    _ = len(v2)
    switch (v6 + v6) {
    case -668:
        var v48 uint32 = 65535
        _ = v48
    }
    v49 := GenS1{F0: 112, F1: (v4 + -1), F2: (v6 & -392)}
    v50 := GenS4{F0: (1.5 >= -2.25), F1: 1}
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 int8 = v4
        _ = fi51
        _ = v52
    }
    v53 := GenS5{F0: 18446744073709551615, F1: genFn0(false, 42, 562), F2: -82}
    v4 = 127
    v54 := GenS5{F0: 42, F1: 2, F2: -1}
    var v55 int64 = 9223372036854775807
    fmt.Println(v4, v4, v55)
    v56 := GenS2{F0: 42}
    for fi57 := 0; fi57 < 13; fi57++ {
        v58 := v2[v5]
        _ = fi57
        _ = v58
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}

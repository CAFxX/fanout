package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int32, p1 string, p2 bool) uint16 {
    var v4 int8 = 0
    if (!(p0 >= (p0 - (p0 >> ((192) & 63))))) {
        p2 = (!(!(0 == -32768)))
    }
    if false {
        v5 := []int{127}
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 57, F1: 0, F2: p1}
            var v8 uint64 = 1
            if (fi6 != fi6) {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v8 += 42
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi9
                }
            } else {
                p2 = false
                var v10 int32 = 1
                if (p0 != 0) {
                    v11 := GenS1{F0: (^(v4)), F1: 2, F2: p1}
                    _ = v11
                }
                _ = v10
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return 660
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 40
    var v3 float32 = (0.1 / 1.5)
    switch 474 {
    case 42:
        v13 := []int{42, 0}
        _ = v13
    case 2147483647:
        v14 := GenS1{F0: -72, F1: 137, F2: (("" + ("hello" + "foo bar")) + ("" + "\n\t"))}
        _ = v14
    default:
        var v15 bool = (-1 <= 0)
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v1[0] = ""
        var v17 uint16 = 0
        _ = fi16
        _ = v17
    }
    if true {
        switch -1 {
        case 128:
            v2 += v2
        default:
            var v18 uint32 = 232
            _ = v18
        }
        v2 -= (^(v2))
        v19 := []int{}
        _ = v19
    } else {
        v20 := []int{-640, 346, -765, -40}
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 uint16 = 65535
        _ = len(v1)
        _ = fi21
        _ = v22
    }
    var v23 float64 = float64((v3 * 1e10))
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 bool = true
        v26 := GenS2{F0: (!v25), F1: 65535, F2: 127}
        if (128 == (v2 | v2)) {
            for fi27 := 0; fi27 < 16; fi27++ {
                switch 557 {
                case -1:
                    v28 := []int{-2147483648}
                    _ = v28
                case -922:
                    switch 42 {
                    case 4294967295:
                        if true {
                            var v29 int = fi24
                            switch "\n\t" {
                            case "hello":
                                var v30 bool = v25
                                _ = v30
                            case "a":
                                v29 = fi27
                            case "\n\t":
                                v31 := GenS3{F0: 0, F1: 4294967295, F2: 2147483648}
                                _ = v31
                            default:
                                v23 = (v23 / (v23 - (float64((0.0 / 0.1)) - float64(v3))))
                            }
                            var v32 string = "\n\t"
                            _ = v29
                            _ = v32
                        } else {
                            v23 = (1e10 / (v23 * ((float64((-2.25 / v3)) / v23) * v23)))
                        }
                    case -997:
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            v33 := GenS1{F0: -1, F1: uint(v2), F2: "foo bar"}
                            _ = v33
                        }
                    case 9223372036854775807:
                        v34 := []int{}
                        _ = v34
                    }
                    var v35 int32 = 2147483647
                    v36 := GenS1{F0: 23, F1: 255, F2: "a"}
                    _ = v35
                    _ = v36
                case 166:
                    v37 := v1[2147483647]
                    _ = v37
                }
                _ = fi27
            }
        } else {
            v23 -= 1e10
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v38 uint16 = 245
    v23 = (v23 / v23)
    switch "hello" {
    case "\n\t":
        v39 := GenS4{F0: 42, F1: "hello"}
        _ = v39
    case "hello":
        v40 := GenS2{F0: true, F1: 42, F2: 1}
        _ = v40
    }
    if (false || (!(-9223372036854775808 > 9223372036854775807))) {
        for fi41 := 0; fi41 < 10; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := GenS2{F0: false, F1: 199, F2: 0}
                var v44 int8 = -59
                _ = fi42
                _ = v43
                _ = v44
            }
            var v45 float64 = float64(v3)
            _ = fi41
            _ = v45
        }
    } else {
        v46 := []int{128, -758, 0}
        if false {
            v47 := GenS5{F0: ((741 > 0) || (!true))}
            _ = v47
        }
        v38 = genFn0(42, "foo bar", (18446744073709551615 <= 764))
        _ = v46
    }
    _ = len(v1)
    v2 = v2
    switch 86 {
    case 42:
        fmt.Println(v38, v23, v23)
    case 2147483647:
        if true {
            v48 := GenS6{F0: v2}
            _ = v48
        }
    case 9223372036854775807:
        v23 = float64((v3 * 1e10))
    }
    fmt.Println(v38, v23, v2)
    v2 = (v2 | (v2 | (v2 ^ 255)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
}

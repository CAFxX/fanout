package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int64 }

func main() {
    var v1 bool = true
    var v2 uint64 = 2
    var v3 uint16 = 0
    var v4 uint = 2
    v1 = true
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int8 = 0
        v4 -= v4
        v1 = (true || (!(0 != 176)))
        _ = fi5
        _ = v6
    }
    if true {
        var v7 int = -719
        if v1 {
            for fi8 := 0; fi8 < 1; fi8++ {
                v7 = 2147483647
                _ = fi8
            }
            v2 = v2
        } else {
            v3 -= v3
        }
        _ = v7
    } else {
        if (!v1) {
            switch -946 {
            case -9223372036854775808:
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(fi9, fi9, v2)
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v2 = 709
                        switch fi9 {
                        case 42:
                            var v11 float32 = -2.25
                            _ = v11
                        case 4294967295:
                            var v12 string = "\n\t"
                            _ = v12
                        case -456:
                            v4 -= v4
                        default:
                            fmt.Println(v2, fi9, v1)
                            v2 = (v2 % ((0) | 1))
                            var v13 bool = false
                            _ = v13
                        }
                        _ = fi10
                    }
                    var v14 float64 = (0.1 * float64(-2.25))
                    _ = fi9
                    _ = v14
                }
            case -965:
                v15 := GenS1{F0: 47, F1: 1, F2: (^(v4))}
                _ = v15
            case 269:
                if true {
                    var v16 int32 = 0
                    _ = v16
                }
            }
            v17 := []int{127, 42, 9223372036854775807, 127}
            v18 := GenS1{F0: -125, F1: 806, F2: v4}
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v4 = v4
        if v1 {
            var v20 int32 = -1
            v4 += uint(v3)
            _ = v20
        }
        _ = fi19
    }
    switch ("hello" + (("foo bar" + "\n\t") + "foo bar")) {
    case "foo bar":
        switch "\n\t" {
        case "":
            if v1 {
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 uint32 = 599
                        v3 = v3
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS1{F0: 94, F1: -478, F2: v4}
                    v25 := []int{9223372036854775807}
                    _ = fi21
                    _ = v24
                    _ = v25
                }
                fmt.Println(v2, v1)
            }
        default:
            var v26 uint64 = uint64(v3)
            _ = v26
        }
    case "\n\t":
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS2{F0: (v2 + (^(v2)))}
            if (uint16(v4) > v3) {
                v2 = v2
                v1 = (v1 || (true && false))
            } else {
                if v1 {
                    fmt.Println(fi27, v1)
                    var v29 uint8 = 42
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        if (fi30 == fi30) {
                            v31 := []int{int(v2), (fi30 | (fi27 ^ -273)), (fi30 * (fi30 * (-(fi27))))}
                            if ("foo bar" > "") {
                                v32 := GenS3{F0: (v4 ^ 56)}
                                var v33 uint64 = (v2 - (^(v2)))
                                _ = v32
                                _ = v33
                            } else {
                                switch "hello" {
                                case "a":
                                    var v34 string = "\n\t"
                                    _ = v34
                                case "hello":
                                    v35 := GenS1{F0: 1, F1: 0, F2: v4}
                                    _ = v35
                                case "\n\t":
                                    if len(v31) > 0 {
                                        v31[0] = fi30
                                    }
                                }
                            }
                            _ = v31
                        }
                        v36 := []int{fi27, 127, (-(fi27)), fi27}
                        fmt.Println(v3, v1)
                        _ = fi30
                        _ = v36
                    }
                    if (false || (((!(false || (-2.25 != 1.5))) || v1) || true)) {
                        var v37 uint32 = 2
                        fmt.Println(v4, v2)
                        v38 := GenS4{F0: -777, F1: 1}
                        _ = v37
                        _ = v38
                    }
                }
                v2 -= (v2 << (((v4 ^ v4)) & 63))
                v4 = uint(v2)
            }
            _ = fi27
            _ = v28
        }
    case "hello":
        for fi39 := 0; fi39 < 3; fi39++ {
            if (2147483648 < 2147483647) {
                if ((-926 <= (fi39 % (((fi39 - (^(fi39)))) | 1))) || (!v1)) {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 int16 = 0
                        _ = fi40
                        _ = v41
                    }
                    var v42 uint = 255
                    _ = v42
                } else {
                    var v43 int8 = 42
                    v44 := []int{(fi39 ^ fi39), fi39}
                    _ = v43
                    _ = v44
                }
            }
            _ = fi39
        }
    default:
        v2 = (v2 << (((^(v4))) & 63))
    }
    v45 := GenS2{F0: 991}
    if (100 < 0) {
        var v46 int = 0
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}

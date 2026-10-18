package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 int) uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := []int{(fi5 % ((1) | 1)), (-(fi4)), (p0 * fi4)}
            _ = fi5
            _ = v6
        }
        var v7 float32 = 0.0
        if true {
            p0 += fi4
            var v8 uint16 = 885
            v7 += (1.5 - v7)
            _ = v8
        }
        _ = fi4
        _ = v7
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint32 = 758
        _ = fi9
        _ = v10
    }
    v11 := []int{(p0 + (p0 % ((p0) | 1))), (p0 - p0), p0, p0}
    _ = p0
    _ = v11
    return 148
}

func main() {
    var v1 bool = (!false)
    var v2 int32 = 42
    var v3 uint16 = 255
    fmt.Println(v2, v1)
    v12 := GenS1{F0: 370}
    switch "a" {
    case "foo bar":
        if false {
            v13 := []int{-9223372036854775808, int(v3)}
            _ = v13
        }
    case "":
        v2 += (v2 % ((-1) | 1))
    }
    if v1 {
        v14 := []int{}
        switch 598 {
        case -9223372036854775808:
            v15 := GenS2{F0: v3}
            _ = v15
        case 2147483647:
            var v16 uint = 4294967295
            _ = v16
        case -1:
            fmt.Println(v2, v1)
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint32 = 261
            switch "a" {
            case "a":
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int64 = -985
                    _ = fi19
                    _ = v20
                }
            default:
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3, v3, v2)
                case "hello":
                    v21 := GenS1{F0: 1}
                    _ = v21
                default:
                    var v22 uint16 = (v3 >> ((177) & 63))
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                switch ("foo bar" + "a") {
                case "a":
                    switch fi23 {
                    case 813:
                        var v24 int8 = -1
                        _ = v24
                    }
                case "hello":
                    v14 = append(v14, fi17)
                default:
                    v25 := GenS2{F0: uint16(v18)}
                    _ = v25
                }
                v3 -= (v3 * v3)
                v18 = v18
                _ = fi23
            }
            _ = fi17
            _ = v18
        }
        _ = v14
    }
    fmt.Println(v2, v1)
    v26 := []int{}
    for fi27 := 0; fi27 < 14; fi27++ {
        if v1 {
            if false {
                var v28 uint64 = 2
                if (0.1 > 0.1) {
                    switch ("a" + "a") {
                    case "a":
                        v1 = v1
                    case "":
                        v2 = (v2 / ((295) | 1))
                    case "\n\t":
                        if true {
                            v28 = v28
                        }
                    }
                }
                v1 = v1
                _ = v28
            }
        } else {
            v3 = (v3 + v3)
            for fi29 := 0; fi29 < 6; fi29++ {
                if false {
                    if (0 >= 42) {
                        _ = len(v26)
                        if (-9223372036854775808 > -9223372036854775808) {
                            v2 -= v2
                            var v30 int = -650
                            var v31 int = -9223372036854775808
                            _ = v30
                            _ = v31
                        }
                    }
                    switch "hello" {
                    case "hello":
                        v32 := GenS2{F0: (v3 / (((v3 ^ (v3 / ((v3) | 1)))) | 1))}
                        _ = v32
                    case "\n\t":
                        var v33 float64 = 1e10
                        _ = v33
                    case "foo bar":
                        switch ("a" + "") {
                        case "a":
                            var v34 string = "foo bar"
                            _ = v34
                        case "hello":
                            if (true || v1) {
                                v2 = v2
                                fmt.Println(fi27, v1)
                                if len(v26) > 0 {
                                    v26[0] = fi29
                                }
                                var v35 uint8 = 127
                                _ = v35
                            }
                        }
                    }
                    if v1 {
                        var v36 int32 = (-(v2))
                        _ = v36
                    }
                }
                _ = fi29
            }
            if len(v26) > 0 {
                v26[0] = fi27
            }
        }
        _ = fi27
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v37 := GenS3{F0: "\n\t", F1: -1, F2: 1}
    var v38 int = -9223372036854775808
    for fi39 := 0; fi39 < 17; fi39++ {
        _ = len(v26)
        if len(v26) > 0 {
            v26[0] = (fi39 + (fi39 & fi39))
        }
        _ = fi39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        fmt.Println(v38, v38)
        v41 := GenS4{F0: false, F1: v3}
        _ = len(v26)
        _ = fi40
        _ = v41
    }
    if (true || (47 >= 242)) {
        fmt.Println(v2, v2, v1)
    }
    if len(v26) > 0 {
        v26[0] = (-(v38))
    }
    var v42 int = v38
    v38 = int(v3)
    v3 = v3
    var v43 int32 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v42
    _ = v43
}

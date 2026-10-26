package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int16; F1 int8; F2 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16) bool {
    var v7 uint8 = 255
    v8 := []int{2147483647, 42}
    _ = p0
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 int) int8 {
    p0 = 9223372036854775807
    p0 = 128
    var v9 float32 = (float32(0.1) + (-2.25 + float32(1e10)))
    var v10 bool = (!false)
    v10 = (!(!false))
    _ = p0
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = -2.25
    var v3 string = ""
    var v4 uint8 = 1
    var v5 float64 = -2.25
    var v6 uint = 255
    switch ("hello" + v3) {
    case "foo bar":
        v11 := []int{-1, -1}
        _ = v11
    case "a":
        v5 = (3.14159 + v5)
    case "":
        var v12 uint = 18446744073709551615
        _ = v12
    }
    switch "a" {
    case "foo bar":
        v13 := GenS1{F0: 681, F1: 255}
        _ = v13
    }
    v1 = ""
    v5 = 1e10
    for fi14 := 0; fi14 < 11; fi14++ {
        if ((2 >= 42) && (73 > -31)) {
            var v15 int64 = -9223372036854775808
            var v16 int32 = 0
            var v17 int16 = -32768
            _ = v15
            _ = v16
            _ = v17
        } else {
            var v18 uint32 = 313
            _ = v18
        }
        var v19 float32 = ((v2 * (v2 * ((v2 - 1e10) / (float32(float64(1e10)) / float32(v5))))) / float32(v5))
        _ = fi14
        _ = v19
    }
    v20 := GenS1{F0: -1, F1: (v4 % ((255) | 1))}
    fmt.Println(v5)
    switch v3 {
    case "foo bar":
        if (true || (!("" < ""))) {
            if false {
                var v21 int8 = -128
                if ((false || ((452 >= 0) && true)) && ((^(v4)) > (v4 | (v4 / ((242) | 1))))) {
                    v3 = (v3 + "foo bar")
                    var v22 int16 = 0
                    _ = v22
                } else {
                    v4 += 1
                }
                for fi23 := 0; fi23 < 6; fi23++ {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v25 := GenS1{F0: -118, F1: 42}
                        v4 -= v4
                        _ = fi24
                        _ = v25
                    }
                    var v26 int8 = v21
                    _ = fi23
                    _ = v26
                }
                _ = v21
            } else {
                fmt.Println(v6, v6)
                for fi27 := 0; fi27 < 19; fi27++ {
                    v28 := GenS2{F0: 100, F1: 2}
                    if (false && (42 < 974)) {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            switch fi27 {
                            case 32768:
                                v30 := []int{42}
                                _ = v30
                            case -771:
                                if ((false || true) || false) {
                                    v31 := GenS1{F0: -193, F1: 42}
                                    if false {
                                        if (false && false) {
                                            v32 := []int{(-(fi27)), fi27}
                                            _ = v32
                                        } else {
                                            v33 := GenS2{F0: 1, F1: 0}
                                            var v34 bool = true
                                            _ = v33
                                            _ = v34
                                        }
                                        v3 = ((((v3 + v1) + (((v1 + v3) + v3) + (("\n\t" + (v3 + (("\n\t" + "foo bar") + v3))) + ""))) + v3) + "")
                                    } else {
                                        var v35 int = fi29
                                        _ = v35
                                    }
                                    _ = v31
                                }
                            default:
                                var v36 uint = v6
                                _ = v36
                            }
                            var v37 float64 = ((v5 / v5) * (((float64(1.5) * (((float64(v2) + 0.1) / (v5 - (v5 / ((v5 + 1.5) * v5)))) / (((v5 / v5) - (-2.25 - (((0.1 * 0.0) + v5) / float64(float32(v5))))) + 0.0))) * 1.5) - v5))
                            _ = fi29
                            _ = v37
                        }
                        v4 += v4
                    } else {
                        var v38 string = "hello"
                        _ = v38
                    }
                    _ = fi27
                    _ = v28
                }
            }
        } else {
            for fi39 := 0; fi39 < 7; fi39++ {
                v2 = v2
                _ = fi39
            }
        }
    default:
        var v40 int32 = int32(v6)
        _ = v40
    }
    v41 := GenS3{F0: 42, F1: 1, F2: 127}
    var v42 uint32 = 2147483648
    if (9223372036854775807 != 553) {
        var v43 float32 = v2
        _ = v43
    }
    v44 := GenS4{F0: 1}
    fmt.Println(v2, v2, v1)
    if true {
        v45 := GenS2{F0: int8(v42), F1: 716}
        _ = v45
    } else {
        switch -2147483648 {
        case 128:
            v5 = 0.0
        case 32767:
            var v46 int32 = 0
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}

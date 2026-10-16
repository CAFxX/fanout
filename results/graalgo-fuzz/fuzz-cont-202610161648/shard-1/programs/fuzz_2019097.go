package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 uint, p1 uint16, p2 uint) int64 {
    if (!true) {
        if (!(p2 >= p2)) {
            v4 := []int{42}
            if (true && true) {
                for fi5 := 0; fi5 < 13; fi5++ {
                    var v6 uint8 = 2
                    _ = fi5
                    _ = v6
                }
            } else {
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 int32 = 2147483647
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 int = fi9
                    for fi11 := 0; fi11 < 11; fi11++ {
                        fmt.Println(p0, fi9, p0)
                        if false {
                            _ = len(v4)
                        }
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v13 := GenS1{F0: 255, F1: 1}
                        var v14 string = "hello"
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                }
                var v15 float64 = ((((1e10 / 1e10) + (1.5 + (1e10 * 3.14159))) / 1.5) - -2.25)
                _ = v15
            }
            _ = v4
        }
    } else {
        v16 := []int{32767, 812}
        var v17 uint64 = 736
        p1 -= (p1 + p1)
        _ = v16
        _ = v17
    }
    v18 := []int{-1, 2147483647, -2147483648, 32768}
    var v19 bool = (0 == 32767)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return -504
}
func genFn1() uint8 {
    var v20 string = "hello"
    var v21 float64 = 1.5
    _ = v20
    _ = v21
    return 128
}

func main() {
    v1 := []int{792, 42}
    var v2 bool = false
    var v3 int64 = -25
    fmt.Println(v2, v2)
    fmt.Println(v3)
    for fi22 := 0; fi22 < 19; fi22++ {
        v2 = (!v2)
        var v23 float32 = 3.14159
        _ = fi22
        _ = v23
    }
    var v24 uint64 = 2
    v25 := GenS1{F0: 128, F1: 49}
    switch int(v24) {
    case 42:
        fmt.Println(v24, v2)
    case -193:
        var v26 int16 = -32768
        _ = v26
    }
    var v27 uint = 219
    fmt.Println(v2, v27)
    var v28 uint32 = 4294967295
    var v29 bool = (false || v2)
    for fi30 := 0; fi30 < 9; fi30++ {
        v28 += (v28 * (v28 - 0))
        if ((!v29) && false) {
            for fi31 := 0; fi31 < 12; fi31++ {
                v32 := GenS2{F0: "foo bar", F1: 32767, F2: -47}
                switch "\n\t" {
                case "hello":
                    fmt.Println(v28, v24)
                case "\n\t":
                    var v33 float64 = (((-2.25 + float64(0.1)) + ((0.0 * 0.1) - (float64(float32((3.14159 + 0.1))) / float64(((float32(-2.25) / float32(1e10)) / float32(-2.25)))))) / -2.25)
                    _ = v33
                case "":
                    v27 -= (^(v27))
                default:
                    switch ("hello" + (("\n\t" + "\n\t") + (("foo bar" + "") + (("" + "hello") + "")))) {
                    case "hello":
                        if false {
                            v3 += (v3 * v3)
                            if (3.14159 <= 0.0) {
                                v34 := GenS2{F0: ("hello" + "\n\t"), F1: -219, F2: 42}
                                switch "foo bar" {
                                case "foo bar":
                                    v24 = v24
                                }
                                _ = v34
                            }
                            if v29 {
                                v35 := GenS1{F0: genFn1(), F1: 8}
                                v36 := GenS1{F0: 223, F1: 2}
                                _ = v35
                                _ = v36
                            }
                        } else {
                            v29 = ((!v2) && true)
                            var v37 uint8 = genFn1()
                            v2 = true
                            _ = v37
                        }
                    }
                }
                v1 = append(v1, (fi31 << (((v27 + v27)) & 63)))
                _ = fi31
                _ = v32
            }
            v24 = uint64(v28)
            if v2 {
                var v38 uint16 = 65535
                v1 = append(v1, fi30)
                var v39 uint32 = 51
                _ = v38
                _ = v39
            } else {
                var v40 string = "\n\t"
                v41 := GenS3{F0: 128, F1: uint8(v24), F2: v24}
                for fi42 := 0; fi42 < 15; fi42++ {
                    var v43 float64 = (float64(3.14159) - 1.5)
                    v2 = true
                    var v44 float32 = 1.5
                    var v45 uint16 = uint16(v24)
                    fmt.Println(v3)
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                }
                _ = v40
                _ = v41
            }
        } else {
            v29 = false
            v3 += (^(v3))
        }
        _ = fi30
    }
    v46 := GenS2{F0: "", F1: int16(v3), F2: 1}
    var v47 int16 = -518
    if true {
        v24 += (v24 / ((v24) | 1))
    }
    fmt.Println(v24)
    var v48 uint = v27
    for fi49 := 0; fi49 < 11; fi49++ {
        v27 -= (v48 % ((v27) | 1))
        for fi50 := 0; fi50 < 10; fi50++ {
            var v51 float64 = float64(0.1)
            v52 := GenS4{F0: 7, F1: -677, F2: 26}
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v46
    _ = v47
    _ = v48
}

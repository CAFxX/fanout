package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int16 = 1
        var v6 int16 = (v5 - (v5 / ((32767) | 1)))
        fmt.Println(v5)
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{1, -694, 32768, 42}
    if len(v7) > 0 {
        v7[0] = 744
    }
    v8 := GenS1{F0: 0}
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int8) int16 {
    var v9 int16 = -1
    v10 := GenS2{F0: 598, F1: ((("" + "a") + "hello") + (("a" + ("\n\t" + ("" + ""))) + ("hello" + "a"))), F2: genFn0()}
    for fi11 := 0; fi11 < 18; fi11++ {
        p0 = (-(p0))
        _ = fi11
    }
    var v12 int32 = -510
    _ = p0
    _ = v10
    _ = v12
    _ = v9
    return -352
}

func main() {
    var v1 float32 = (1.5 / 3.14159)
    v2 := []string{(("hello" + "foo bar") + "")}
    var v3 int8 = 42
    var v13 string = (("\n\t" + "\n\t") + ("" + (("foo bar" + "") + "")))
    fmt.Println(v1, v13)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float32 = v1
        _ = fi14
        _ = v15
    }
    var v16 float64 = -2.25
    v17 := GenS1{F0: 0}
    if true {
        if true {
            v1 -= float32(v16)
            v18 := GenS2{F0: 18446744073709551615, F1: v13, F2: 215}
            v16 = 0.1
            _ = v18
        } else {
            var v19 uint16 = 65535
            v3 -= v3
            if (!(1 > -604)) {
                if true {
                    v20 := GenS3{F0: -406}
                    _ = v20
                } else {
                    v16 += v16
                }
                if true {
                    v16 -= v16
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v22 := GenS2{F0: 42, F1: "", F2: 255}
                        v16 = (v16 * (float64((-2.25 * (v1 * v1))) - -2.25))
                        v3 += v3
                        _ = fi21
                        _ = v22
                    }
                    var v23 uint16 = v19
                    _ = v23
                }
                if false {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 uint16 = v19
                        for fi26 := 0; fi26 < 12; fi26++ {
                            v16 += (0.1 / v16)
                            _ = fi26
                        }
                        v27 := GenS4{F0: genFn0(), F1: 42}
                        _ = fi24
                        _ = v25
                        _ = v27
                    }
                    var v28 float32 = v1
                    _ = v28
                }
            }
            _ = v19
        }
    } else {
        v16 = v16
        if true {
            for fi29 := 0; fi29 < 6; fi29++ {
                v30 := GenS1{F0: 0}
                v31 := GenS5{F0: (128 < 0), F1: 1, F2: true}
                v32 := GenS1{F0: 65535}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            if (42 < -667) {
                v16 = v16
                switch v13 {
                case "a":
                    var v33 int16 = -484
                    _ = v33
                default:
                    switch "\n\t" {
                    case "\n\t":
                        v34 := GenS5{F0: true, F1: -223, F2: true}
                        _ = v34
                    default:
                        v35 := GenS4{F0: 255, F1: 2}
                        _ = v35
                    }
                }
            } else {
                v13 = (((v13 + v13) + v13) + "")
                var v36 uint8 = 2
                v1 -= v1
                _ = v36
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        }
    }
    v2 = append(v2, (v13 + v13))
    _ = len(v2)
    var v37 int64 = 450
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0() bool {
    v4 := GenS1{F0: 1, F1: -9223372036854775808}
    v5 := []int{54, 2147483647}
    var v6 bool = (28 >= 42)
    _ = v4
    _ = v5
    _ = v6
    return (!(100 != 103))
}
func genFn1(p0 int64, p1 uint32, p2 int) uint64 {
    var v7 uint64 = 2
    v8 := []int{}
    if (42 == 400) {
        if false {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int32 = -961
                _ = fi9
                _ = v10
            }
            if true {
                v11 := GenS2{F0: true}
                _ = v11
            } else {
                if false {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        var v13 float32 = (float32(0.0) / (0.0 - float32(0.1)))
                        switch int(v7) {
                        case 1:
                            v14 := GenS1{F0: 0, F1: (p2 - (p2 / ((fi12) | 1)))}
                            _ = v14
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        if (false || (-2147483648 > 12)) {
                            var v15 uint8 = 219
                            var v16 uint = 18446744073709551615
                            _ = v15
                            _ = v16
                        } else {
                            v7 = (^(v7))
                        }
                        _ = fi12
                        _ = v13
                    }
                    p0 = 1
                    switch (("hello" + ("" + "")) + ("" + "hello")) {
                    case "":
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 uint64 = v7
                            if (!false) {
                                var v19 uint64 = 4294967295
                                var v20 float64 = 1.5
                                v19 -= (v7 + v19)
                                _ = v19
                                _ = v20
                            } else {
                                for fi21 := 0; fi21 < 6; fi21++ {
                                    if (!false) {
                                        _ = len(v8)
                                        var v22 int8 = 102
                                        _ = v22
                                    } else {
                                        if false {
                                            var v23 uint = 656
                                            _ = v23
                                        }
                                    }
                                    _ = fi21
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                    default:
                        if (false || true) {
                            if len(v8) > 0 {
                                v8[0] = 42
                            }
                        }
                    }
                }
            }
        } else {
            v24 := GenS3{F0: 100, F1: -2147483648}
            if ((74 == 127) || (true || false)) {
                var v25 bool = false
                p2 = (p2 + 32768)
                _ = v25
            }
            _ = v24
        }
    } else {
        v26 := GenS1{F0: -1, F1: (p2 / ((42) | 1))}
        _ = v26
    }
    v27 := GenS1{F0: -32768, F1: (p2 * p2)}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v7
    _ = v8
    return (^(v7))
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = false
    var v3 int64 = -262
    switch 4294967295 {
    case 128:
        fmt.Println(v2)
    }
    var v28 int32 = -2147483648
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint16 = 255
        for fi31 := 0; fi31 < 13; fi31++ {
            if v2 {
                v32 := []int{-246, fi31, fi29, fi29}
                var v33 float64 = 1e10
                var v34 uint32 = 2
                if (-44 == 100) {
                    var v35 int = -9223372036854775808
                    switch v1 {
                    case "a":
                        if v2 {
                            var v36 uint = 255
                            _ = v36
                        }
                    case "":
                        fmt.Println(v3)
                    default:
                        v37 := GenS3{F0: 1, F1: (v28 & (-(v28)))}
                        _ = v37
                    }
                    _ = v35
                }
                fmt.Println(v33, v3)
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v38 uint16 = v30
                var v39 int32 = (v28 % ((v28) | 1))
                _ = v38
                _ = v39
            }
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v1 = "a"
    v28 = (v28 & v28)
    if ((v3 / (((^(v3))) | 1)) < (v3 << ((4294967295) & 63))) {
        for fi40 := 0; fi40 < 14; fi40++ {
            v28 = v28
            var v41 uint32 = 638
            _ = fi40
            _ = v41
        }
        v1 = ((v1 + (((v1 + "hello") + v1) + v1)) + v1)
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        for fi43 := 0; fi43 < 1; fi43++ {
            if v2 {
                v44 := GenS2{F0: false}
                v1 = v1
                _ = v44
            } else {
                v28 = (v28 | v28)
                var v45 float32 = (-2.25 + 3.14159)
                var v46 uint32 = 188
                var v47 int = -779
                v48 := GenS3{F0: 42, F1: (v28 & (^(v28)))}
                _ = v45
                _ = v46
                _ = v47
                _ = v48
            }
            var v49 int8 = -124
            v28 = 42
            _ = fi43
            _ = v49
        }
        _ = fi42
    }
    fmt.Println(v28)
    v3 = 42
    var v50 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v50
}

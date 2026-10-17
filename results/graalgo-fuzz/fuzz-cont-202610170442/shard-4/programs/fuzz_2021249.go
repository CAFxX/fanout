package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int; F1 int }
func genFn0() uint64 {
    var v6 uint16 = 42
    var v7 float64 = ((-2.25 * (float64((0.0 * 1.5)) + (-2.25 - 1e10))) * float64(1.5))
    fmt.Println(v7, v6)
    v8 := []int{4294967295, 1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1() uint64 {
    var v9 int16 = 42
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: 1}
        v9 -= v9
        v12 := GenS1{F0: 0}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("foo bar" + "foo bar") {
    case "":
        for fi13 := 0; fi13 < 3; fi13++ {
            v9 = (v9 | v9)
            for fi14 := 0; fi14 < 8; fi14++ {
                v9 += v9
                v15 := GenS2{F0: -452, F1: false}
                var v16 int = 32768
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 991}
            _ = fi13
            _ = v17
        }
    case "\n\t":
        v9 = (v9 | 0)
    }
    _ = v9
    return 2
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 uint64 = 332
    v5 := make(map[uint32]bool)
    v18 := []int{2147483647, 4294967295}
    _ = len(v18)
    fmt.Println(v4)
    fmt.Println(v4)
    var v19 int32 = 1
    if ((v19 + (^(v19))) > int32(v4)) {
        v4 = v4
        v18 = append(v18, -9223372036854775808)
    }
    switch 42 {
    case -2147483648:
        var v20 uint64 = (v4 ^ v4)
        _ = v20
    case -178:
        v19 = v19
    case 1:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float64 = float64((-2.25 - 1.5))
            _ = fi21
            _ = v22
        }
    default:
        if false {
            if ((!false) || (!(((v1 <= v1) && true) && (!true)))) {
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch fi23 {
                    case 4294967295:
                        var v24 float64 = v1
                        _ = v24
                    case -281:
                        v19 -= (-(v19))
                    case 127:
                        v25 := GenS2{F0: 0, F1: ((v4 % (((v4 - 42)) | 1)) == (v4 + (^(v4))))}
                        _ = v25
                    default:
                        var v26 int64 = 0
                        _ = v26
                    }
                    var v27 bool = (-2147483648 == -401)
                    _ = fi23
                    _ = v27
                }
                fmt.Println(v1, v4)
                switch 32768 {
                case -958:
                    if true {
                        v4 -= v4
                        v28 := GenS2{F0: -9223372036854775808, F1: false}
                        _ = v28
                    } else {
                        var v29 uint32 = 65535
                        if false {
                            v30 := GenS1{F0: v19}
                            v31 := GenS3{F0: 751, F1: 255, F2: (v4 - genFn0())}
                            var v32 int16 = -1
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            v33 := v3[128]
                            for fi34 := 0; fi34 < 12; fi34++ {
                                var v35 float32 = 3.14159
                                v35 += ((3.14159 * v35) * (v35 - 0.0))
                                fmt.Println(v4)
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                        var v36 uint8 = uint8(v29)
                        _ = v29
                        _ = v36
                    }
                case -476:
                    var v37 int32 = 1
                    _ = v37
                }
            } else {
                if (9223372036854775807 != 1) {
                    _ = len(v5)
                    v1 -= v1
                }
                v38 := GenS1{F0: (v19 / ((0) | 1))}
                _ = len(v3)
                _ = v38
            }
        } else {
            var v39 float64 = (1.5 - v1)
            _ = v39
        }
    }
    v40 := GenS3{F0: -671, F1: 128, F2: (v4 >> ((1) & 63))}
    for fi41 := 0; fi41 < 12; fi41++ {
        var v42 uint8 = 255
        if len(v18) > 0 {
            v18[0] = (^(fi41))
        }
        if true {
            var v43 int8 = 94
            v44 := v5[2147483647]
            for fi45 := 0; fi45 < 3; fi45++ {
                v46 := GenS1{F0: v19}
                fmt.Println(v44)
                _ = fi45
                _ = v46
            }
            _ = v43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    if false {
        v47 := GenS4{F0: 2147483647}
        _ = v47
    } else {
        switch "\n\t" {
        case "hello":
            v48 := GenS5{F0: 244, F1: ""}
            _ = v48
        case "\n\t":
            v19 += (-(v19))
        case "a":
            if true {
                v49 := GenS6{F0: 32767, F1: -1}
                _ = v49
            } else {
                for fi50 := 0; fi50 < 18; fi50++ {
                    var v51 uint16 = 0
                    _ = fi50
                    _ = v51
                }
            }
        default:
            if ((695 < 910) && ((("hello" + (("\n\t" + ("hello" + ("hello" + "hello"))) + (("foo bar" + "hello") + "a"))) >= "a") || true)) {
                fmt.Println(v4, v4, v1)
                var v52 int8 = -128
                var v53 int64 = 1
                _ = v52
                _ = v53
            }
        }
    }
    _ = len(v3)
    fmt.Println(v19, v4, v19)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}

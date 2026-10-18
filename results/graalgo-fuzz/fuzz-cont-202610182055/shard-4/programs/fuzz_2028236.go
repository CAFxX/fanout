package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64 }
func genFn0() int32 {
    switch "a" {
    case "":
        v4 := []int{0, 42, 1000, 2147483647}
        _ = v4
    case "\n\t":
        var v5 uint32 = 42
        _ = v5
    }
    var v6 float32 = (0.1 * float32(0.1))
    if true {
        v7 := []int{}
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                if (445 < 913) {
                    v9 := GenS1{F0: 0, F1: 42}
                    var v10 uint8 = 127
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v10
                    _ = v9
                }
                var v11 int16 = 1
                v11 += v11
                _ = fi8
                _ = v11
            }
        } else {
            var v12 string = ("" + ("" + ""))
            var v13 uint32 = 313
            _ = v12
            _ = v13
        }
        _ = v7
    }
    fmt.Println(v6)
    _ = v6
    return 1
}
func genFn1() uint16 {
    var v14 bool = true
    fmt.Println(v14)
    v15 := []int{-92, 32767, -1}
    if (4294967295 == 18446744073709551615) {
        v16 := GenS2{F0: 42, F1: -2, F2: 544}
        v15 = append(v15, 42)
        v14 = (v14 || v14)
        _ = v16
    } else {
        v14 = v14
        var v17 int16 = -32768
        v18 := GenS1{F0: 18446744073709551615, F1: 0}
        _ = v17
        _ = v18
    }
    _ = v14
    _ = v15
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 0
    var v3 int16 = -1
    v19 := GenS2{F0: 173, F1: genFn0(), F2: 128}
    if true {
        var v20 int = 42
        v21 := GenS3{F0: ((18446744073709551615 <= 255) || false), F1: 1}
        _ = v20
        _ = v21
    }
    var v22 int64 = 1
    v23 := GenS4{F0: 622, F1: (743 <= 2), F2: 42}
    var v24 string = "\n\t"
    var v25 int = 844
    var v26 uint16 = 65535
    var v27 int16 = int16(v26)
    if false {
        if (!((v24 + (("hello" + ("" + v24)) + (("a" + v24) + v24))) < "hello")) {
            if (!(v25 > -2147483648)) {
                v2 += (-(v2))
                v28 := GenS1{F0: 2, F1: 536}
                _ = v28
            } else {
                switch v24 {
                case "foo bar":
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v24 = (v24 + ("" + "\n\t"))
                        var v30 uint16 = (v26 << ((0) & 63))
                        _ = fi29
                        _ = v30
                    }
                case "a":
                    fmt.Println(v27, v24)
                default:
                    if true {
                        _ = len(v1)
                        v31 := []int{(v25 >> ((255) & 63)), v25}
                        _ = v31
                    }
                }
                if false {
                    v2 -= v3
                    v32 := GenS5{F0: -1}
                    v33 := GenS4{F0: 2147483648, F1: (!(!false)), F2: 551}
                    _ = v32
                    _ = v33
                } else {
                    var v34 float64 = (float64(-2.25) / ((((0.1 + -2.25) + -2.25) * 3.14159) + 3.14159))
                    switch v24 {
                    case "":
                        if true {
                            v35 := GenS4{F0: 674, F1: (false || (true || (42 <= 4294967295))), F2: 566}
                            _ = v35
                        }
                    case "hello":
                        var v36 string = "\n\t"
                        _ = v36
                    default:
                        v26 += 780
                    }
                    _ = v34
                }
                var v37 uint = 18446744073709551615
                _ = v37
            }
            switch ((v24 + "") + (((("a" + v24) + v24) + "foo bar") + ("a" + v24))) {
            case "":
                switch v24 {
                case "foo bar":
                    v38 := GenS2{F0: 894, F1: -1, F2: 2147483647}
                    _ = v38
                default:
                    v39 := []int{v25, v25}
                    _ = v39
                }
            case "foo bar":
                delete(v1, v25)
            default:
                fmt.Println(v2)
            }
        } else {
            var v40 int = (v25 - -459)
            for fi41 := 0; fi41 < 11; fi41++ {
                v1[128] = 350
                fmt.Println(v27)
                _ = fi41
            }
            var v42 uint32 = 4294967295
            _ = v40
            _ = v42
        }
        v43 := GenS5{F0: (v22 / ((-59) | 1))}
        var v44 int8 = 27
        _ = v43
        _ = v44
    }
    v2 -= v2
    v45 := GenS4{F0: 134, F1: false, F2: 818}
    if true {
        var v46 float64 = 0.1
        _ = v46
    }
    v26 += (v26 >> ((42) & 63))
    if false {
        switch "\n\t" {
        case "\n\t":
            for fi47 := 0; fi47 < 15; fi47++ {
                fmt.Println(v24, v2, v25)
                var v48 bool = true
                _ = fi47
                _ = v48
            }
        case "":
            v49 := []int{-244, (-(v25)), (^(v25))}
            _ = v49
        case "hello":
            var v50 int64 = (^(v22))
            _ = v50
        default:
            _ = len(v1)
        }
        fmt.Println(v27, v24)
    }
    v51 := []int{(v25 * (v25 % ((v25) | 1)))}
    var v52 bool = (!(!true))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v45
    _ = v51
    _ = v52
}

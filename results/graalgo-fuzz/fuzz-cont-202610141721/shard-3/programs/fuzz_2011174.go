package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }
func genFn0(p0 bool) int16 {
    switch "a" {
    case "":
        var v5 int64 = 744
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            p0 = (!p0)
            switch ("hello" + ("" + "hello")) {
            case "a":
                var v7 int32 = 0
                _ = v7
            case "foo bar":
                var v8 int16 = -1
                _ = v8
            default:
                v9 := GenS1{F0: 91, F1: -32768}
                _ = v9
            }
            _ = fi6
        }
    }
    p0 = p0
    if p0 {
        v10 := GenS1{F0: 546, F1: 42}
        _ = v10
    } else {
        var v11 uint8 = 127
        for fi12 := 0; fi12 < 15; fi12++ {
            if (p0 || (65535 >= 1)) {
                v13 := GenS2{F0: 42}
                fmt.Println(p0, fi12)
                _ = v13
            } else {
                switch (fi12 + 32767) {
                case 190:
                    v11 = (v11 | v11)
                case 4294967295:
                    fmt.Println(v11, v11)
                }
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 bool = false
                    _ = fi14
                    _ = v15
                }
            }
            switch -784 {
            case 128:
                p0 = p0
            case 42:
                var v16 int8 = -48
                _ = v16
            case 2147483647:
                var v17 int8 = 1
                _ = v17
            }
            _ = fi12
        }
        if (42 == -287) {
            fmt.Println(v11)
        }
        _ = v11
    }
    v18 := []int{}
    var v19 float32 = 1e10
    _ = p0
    _ = v18
    _ = v19
    return 799
}
func genFn1(p0 int16, p1 uint32, p2 int8) uint {
    var v20 int = 2147483647
    var v21 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 922
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.0
    var v3 float64 = (1e10 * float64(float32(1.5)))
    var v4 float64 = float64(((float32(v3) + float32(v3)) - (v2 / (v2 - v2))))
    var v22 uint8 = 128
    fmt.Println(v3, v3, v4)
    v2 -= (((-2.25 - float32(v4)) - (3.14159 * v2)) / (float32((((v4 / v3) - v4) * (0.1 * float64(v2)))) / (v2 + -2.25)))
    v3 += -2.25
    v23 := GenS2{F0: -2147483648}
    var v24 int = -2147483648
    var v25 uint16 = 967
    v24 -= v24
    var v26 uint32 = 913
    v22 -= v22
    if true {
        if (!(42 == 1)) {
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 uint32 = 2
                fmt.Println(v25, v28, fi27)
                _ = fi27
                _ = v28
            }
            v3 += v4
            for fi29 := 0; fi29 < 10; fi29++ {
                if false {
                    switch "a" {
                    case "a":
                        switch "" {
                        case "\n\t":
                            var v30 uint = uint(v22)
                            _ = v30
                        case "a":
                            var v31 uint16 = (^(v25))
                            _ = v31
                        default:
                            v32 := GenS1{F0: 1, F1: 233}
                            _ = v32
                        }
                    case "":
                        switch v24 {
                        case 346:
                            v4 = -2.25
                        default:
                            var v33 int16 = -32768
                            _ = v33
                        }
                    case "\n\t":
                        v2 -= v2
                    default:
                        v4 += ((0.0 * 1.5) - v3)
                    }
                } else {
                    var v34 bool = false
                    _ = v34
                }
                if true {
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v25 = (v25 ^ v25)
                        _ = fi35
                    }
                    switch (^(fi29)) {
                    case -738:
                        v3 = -2.25
                    }
                    v36 := GenS1{F0: 2, F1: 1000}
                    _ = v36
                } else {
                    v37 := GenS3{F0: 192}
                    _ = v37
                }
                _ = fi29
            }
        } else {
            fmt.Println(v4, v25)
            switch "a" {
            case "\n\t":
                fmt.Println(v26, v4)
            case "hello":
                switch (v24 ^ v24) {
                case 0:
                    if (-1 <= -2147483648) {
                        var v38 uint = uint(v25)
                        v39 := v1[(v24 - -9223372036854775808)]
                        _ = v38
                        _ = v39
                    } else {
                        fmt.Println(v25, v22, v3)
                        v4 += -2.25
                        for fi40 := 0; fi40 < 4; fi40++ {
                            delete(v1, 32767)
                            switch (fi40 - 42) {
                            case -224:
                                v41 := GenS2{F0: 0}
                                _ = v41
                            case 2147483647:
                                var v42 uint = genFn1(42, (v26 / (((v26 | (v26 % ((v26) | 1)))) | 1)), int8(v26))
                                _ = v42
                            case 32768:
                                for fi43 := 0; fi43 < 13; fi43++ {
                                    v24 += (v24 / (((fi40 | (fi40 % ((v24) | 1)))) | 1))
                                    var v44 uint8 = (v22 << ((1) & 63))
                                    _ = fi43
                                    _ = v44
                                }
                            default:
                                fmt.Println(v24)
                            }
                            for fi45 := 0; fi45 < 9; fi45++ {
                                if true {
                                    v3 -= 1.5
                                    v46 := []int{}
                                    _ = v46
                                } else {
                                    v25 = v25
                                }
                                v47 := []int{fi45}
                                v1[v24] = ("a" + "")
                                _ = fi45
                                _ = v47
                            }
                            _ = fi40
                        }
                    }
                case 726:
                    var v48 int32 = -2147483648
                    _ = v48
                case 2147483647:
                    v4 = (0.0 / float64(v2))
                }
            case "foo bar":
                if true {
                    v3 -= 1.5
                    if true {
                        var v49 uint = 818
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 14; fi50++ {
                            switch 1 {
                            case 32767:
                                if (-183 > -1) {
                                    var v51 int = 42
                                    _ = v51
                                } else {
                                    var v52 uint8 = (v22 | (v22 ^ 0))
                                    var v53 uint16 = 109
                                    v2 = v2
                                    _ = v52
                                    _ = v53
                                }
                            case 321:
                                v54 := GenS2{F0: 0}
                                _ = v54
                            }
                            fmt.Println(v3, v4)
                            _ = fi50
                        }
                    }
                }
            }
        }
    } else {
        for fi55 := 0; fi55 < 20; fi55++ {
            fmt.Println(v3, v3, v3)
            _ = fi55
        }
    }
    v2 -= (0.1 - float32(3.14159))
    v4 = v4
    for fi56 := 0; fi56 < 9; fi56++ {
        var v57 uint32 = (^(v26))
        _ = fi56
        _ = v57
    }
    v4 += v3
    v25 = 2
    v58 := GenS4{F0: 805, F1: (v22 / ((v22) | 1)), F2: (v22 << ((4294967295) & 63))}
    var v59 float32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v58
    _ = v59
}

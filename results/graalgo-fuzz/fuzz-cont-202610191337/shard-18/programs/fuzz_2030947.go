package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint8; F2 int32 }
type GenS6 struct { F0 int; F1 uint; F2 bool }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int32; F1 int; F2 int64 }

func main() {
    var v1 bool = (1.5 < (0.1 / (1.5 - float64(1e10))))
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 bool = (2147483647 != 169)
    var v5 float32 = (float32(float64((1e10 * float32((1.5 - 1.5))))) / 1.0)
    var v6 int32 = 42
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v1)
        if false {
            var v8 int = 9223372036854775807
            v9 := GenS1{F0: 189}
            _ = v8
            _ = v9
        }
        switch (^(fi7)) {
        case -9223372036854775808:
            fmt.Println(v6, v5)
        case 128:
            fmt.Println(v5, v4, v5)
        case 213:
            var v10 int = fi7
            _ = v10
        default:
            v5 += v5
        }
        _ = fi7
    }
    if (809 < 1) {
        fmt.Println(v6, v4)
        v1 = v1
        if v1 {
            fmt.Println(v6, v1)
        } else {
            v5 = (v5 + (float32(1e10) - 3.14159))
            fmt.Println(v6, v6)
            var v11 float32 = (v5 / (v5 + (v5 + ((1e10 - v5) / ((float32(1e10) + v5) * v5)))))
            _ = v11
        }
    }
    if ((!v4) && ((v4 || (v5 >= v5)) || v4)) {
        v5 -= (float32(1e10) + 1e10)
        var v12 int32 = (v6 ^ v6)
        _ = v12
    } else {
        var v13 int8 = -67
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v5 -= 0.1
        if v1 {
            v15 := GenS2{F0: 0, F1: (true || ((v4 && false) || v4)), F2: "\n\t"}
            _ = v15
        } else {
            switch "foo bar" {
            case "":
                v16 := GenS1{F0: 40}
                _ = v16
            case "hello":
                if false {
                    var v17 int64 = 535
                    _ = v17
                }
            }
            var v18 bool = ((v6 * (-(v6))) <= -282)
            v6 = v6
            _ = v18
        }
        if v1 {
            switch (-(fi14)) {
            case -123:
                _ = len(v3)
            case 32767:
                v19 := GenS3{F0: 0, F1: v4, F2: -1}
                _ = v19
            case 0:
                v5 = float32(1e10)
            }
        }
        _ = fi14
    }
    var v20 uint32 = 2147483647
    v21 := v2[795]
    v2[(v20 % ((v20) | 1))] = v21
    var v22 float64 = ((0.0 / 1e10) - 1e10)
    v20 -= (^(v20))
    if v4 {
        v23 := GenS4{F0: 137}
        switch -9223372036854775808 {
        case 0:
            for fi24 := 0; fi24 < 3; fi24++ {
                v22 = v22
                for fi25 := 0; fi25 < 20; fi25++ {
                    v26 := GenS5{F0: 100, F1: 0, F2: -1}
                    v1 = (v4 || (255 >= 32767))
                    _ = fi25
                    _ = v26
                }
                v4 = (77 >= 42)
                _ = fi24
            }
        default:
            for fi27 := 0; fi27 < 15; fi27++ {
                fmt.Println(v4, v22)
                _ = fi27
            }
        }
        _ = v23
    } else {
        var v28 uint64 = 4294967295
        v1 = (!(42 == 0))
        _ = v28
    }
    if true {
        v5 -= 1e10
        delete(v2, 2147483647)
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if (-1 < -258) {
            v30 := GenS3{F0: 49, F1: v4, F2: 24}
            v1 = v1
            var v31 uint32 = 1
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            if false {
                v33 := GenS6{F0: -2147483648, F1: 0, F2: ((!true) && v1)}
                _ = v33
            } else {
                switch ("" + ("foo bar" + v21)) {
                case "\n\t":
                    v34 := GenS6{F0: fi32, F1: 1, F2: (65535 == 416)}
                    _ = v34
                case "a":
                    v35 := GenS4{F0: 42}
                    _ = v35
                default:
                    fmt.Println(v4, v5, v21)
                }
                switch fi29 {
                case -9223372036854775808:
                    if false {
                        if ((false || (v1 && ("" == v21))) && (474 == -32768)) {
                            v36 := []int{}
                            var v37 uint64 = 18446744073709551615
                            _ = v36
                            _ = v37
                        }
                        if (0 < 42) {
                            v38 := GenS7{F0: 127, F1: 270}
                            v39 := GenS6{F0: (-(fi32)), F1: 0, F2: v4}
                            _ = v38
                            _ = v39
                        } else {
                            if (float32(v22) > (0.0 - v5)) {
                                if ((!(!true)) && v4) {
                                    fmt.Println(v5, v5)
                                    var v40 int32 = v6
                                    fmt.Println(v1, v5)
                                    if v1 {
                                        var v41 float64 = 1e10
                                        switch (fi32 << ((0) & 63)) {
                                        case -394:
                                            var v42 int8 = -1
                                            _ = v42
                                        case 132:
                                            v43 := []int{fi32}
                                            _ = v43
                                        default:
                                            var v44 string = "foo bar"
                                            _ = v44
                                        }
                                        _ = v41
                                    } else {
                                        var v45 int64 = 0
                                        v46 := GenS8{F0: (^(v6)), F1: -225, F2: (-(v45))}
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v40
                                } else {
                                    var v47 int8 = -128
                                    _ = v47
                                }
                                var v48 uint16 = 880
                                _ = v48
                            }
                        }
                        var v49 float32 = v5
                        var v50 string = v21
                        v3[v50] = (((("" + "hello") + v50) + "a") + "hello")
                        _ = v49
                        _ = v50
                    } else {
                        if v4 {
                            v20 = 2
                        }
                        fmt.Println(v5)
                        v4 = ((1 == 0) && v1)
                        fmt.Println(v4, v21)
                    }
                default:
                    var v51 uint16 = 42
                    _ = v51
                }
                v22 -= v22
            }
            _ = fi32
        }
        fmt.Println(fi29)
        _ = fi29
    }
    var v52 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}

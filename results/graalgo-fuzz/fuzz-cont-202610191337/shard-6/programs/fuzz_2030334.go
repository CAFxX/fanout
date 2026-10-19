package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS8 struct { F0 uint64; F1 int32; F2 int8 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = 127
    var v3 float32 = ((3.14159 - ((1e10 - (0.1 * 1e10)) * 0.1)) / (((-2.25 * (1e10 * float32((0.1 * -2.25)))) / 3.14159) / 0.1))
    var v4 uint16 = 723
    var v5 float32 = v3
    for fi6 := 0; fi6 < 1; fi6++ {
        if false {
            v4 -= (v4 | 42)
            var v7 float64 = 0.0
            _ = v7
        }
        v8 := []int{(fi6 >> ((739) & 63))}
        _ = fi6
        _ = v8
    }
    var v9 uint16 = v4
    v2 = v2
    var v10 float64 = (float64(0.1) * 0.1)
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 1; fi12++ {
        v5 = v3
        if false {
            if false {
                v11 = ("\n\t" + v11)
                v4 += v9
            }
            switch fi12 {
            case 128:
                v13 := []int{}
                _ = v13
            default:
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 int32 = -1
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v9 += (v9 & 32767)
                _ = fi16
            }
        }
        switch (v11 + (v11 + v11)) {
        case "":
            v10 = v10
        }
        _ = fi12
    }
    v17 := []int{0}
    for fi18 := 0; fi18 < 9; fi18++ {
        v9 += v9
        v19 := GenS1{F0: 2, F1: 523, F2: -2147483648}
        _ = fi18
        _ = v19
    }
    _ = len(v17)
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(v4)
        if (((-(v2)) != v2) || (!true)) {
            var v21 uint8 = v1
            _ = v21
        } else {
            if (1 == -32768) {
                if (0.1 >= float64(0.1)) {
                    switch (fi20 << ((65535) & 63)) {
                    case 32767:
                        if true {
                            var v22 int16 = 32767
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v11 = v11
                                v11 = v11
                                v24 := GenS2{F0: 482, F1: 0, F2: (^(v1))}
                                _ = fi23
                                _ = v24
                            }
                            v22 = (-(v22))
                            _ = v22
                        } else {
                            if true {
                                var v25 int16 = 42
                                _ = v25
                            } else {
                                var v26 uint64 = 0
                                if false {
                                    v27 := GenS2{F0: 0, F1: 1, F2: v1}
                                    v28 := GenS2{F0: 0, F1: -1, F2: v1}
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 int = (fi20 / ((fi20) | 1))
                                    _ = v29
                                }
                                fmt.Println(v11, v10, v4)
                                _ = v26
                            }
                        }
                    case 2147483647:
                        switch fi20 {
                        case 323:
                            fmt.Println(v1)
                        default:
                            switch (v11 + ("foo bar" + ("a" + ("hello" + v11)))) {
                            case "a":
                                v3 += v3
                            case "\n\t":
                                for fi30 := 0; fi30 < 5; fi30++ {
                                    switch v11 {
                                    case "foo bar":
                                        v31 := GenS3{F0: -1, F1: 2}
                                        _ = v31
                                    case "\n\t":
                                        v32 := GenS2{F0: int64(v1), F1: 9223372036854775807, F2: v1}
                                        _ = v32
                                    default:
                                        v33 := GenS3{F0: (fi20 / ((fi30) | 1)), F1: 1}
                                        _ = v33
                                    }
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    _ = fi30
                                }
                            case "":
                                v34 := GenS4{F0: v11, F1: v9}
                                _ = v34
                            default:
                                v4 -= 42
                            }
                        }
                    }
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v5 = (0.0 - (v3 * -2.25))
                        var v36 bool = ((v5 - (v3 * (float32(float64(3.14159)) * v5))) != (float32(1e10) * (float32((1e10 + (v10 - ((v10 + (v10 + v10)) - v10)))) - float32(v10))))
                        v36 = false
                        _ = fi35
                        _ = v36
                    }
                }
            } else {
                fmt.Println(fi20, fi20)
                var v37 int = fi20
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 uint = 18446744073709551615
            _ = fi38
            _ = v39
        }
        _ = fi20
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: v1}
        if true {
            v42 := GenS5{F0: (^(v1)), F1: fi40}
            var v43 int32 = -615
            if false {
                v44 := GenS6{F0: (v43 / (((v43 | (-(v43)))) | 1))}
                _ = v44
            }
            _ = v42
            _ = v43
        } else {
            if false {
                var v45 int = (fi40 << ((255) & 63))
                _ = v45
            } else {
                if false {
                    for fi46 := 0; fi46 < 10; fi46++ {
                        var v47 string = v11
                        if true {
                            switch (fi40 << ((65535) & 63)) {
                            case -352:
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            case 1:
                                if (v11 > v47) {
                                    v48 := GenS6{F0: 42}
                                    v49 := GenS4{F0: v47, F1: v4}
                                    v4 = v4
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                        v11 = (((((((("" + (v11 + "a")) + (v47 + (v11 + v11))) + v11) + (v11 + "")) + v47) + "\n\t") + ("foo bar" + ((v11 + v47) + (v11 + ("" + ((v47 + v11) + v11)))))) + "hello")
                        if false {
                            if (false && false) {
                                if (742 == int64(v1)) {
                                    fmt.Println(fi46, fi46)
                                    v4 = 551
                                } else {
                                    if (!(v1 != v1)) {
                                        v50 := GenS3{F0: (fi46 - (fi46 << ((42) & 63))), F1: 4294967295}
                                        _ = v50
                                    }
                                }
                                v11 = v47
                                var v51 uint = 311
                                _ = v51
                            }
                            if false {
                                v17 = append(v17, fi40)
                                fmt.Println(v11, v11)
                            }
                            v2 += v2
                            if false {
                                var v52 string = v47
                                _ = v52
                            }
                            v4 -= v4
                        }
                        v53 := GenS7{F0: 1, F1: 1000, F2: uint8(v9)}
                        _ = fi46
                        _ = v47
                        _ = v53
                    }
                    switch v11 {
                    case "":
                        v4 = v4
                    default:
                        var v54 float32 = v5
                        _ = v54
                    }
                }
                if true {
                    var v55 string = ((v11 + ("" + ("foo bar" + "hello"))) + v11)
                    for fi56 := 0; fi56 < 3; fi56++ {
                        var v57 float64 = v10
                        var v58 uint = 2
                        v57 = 0.0
                        if len(v17) > 0 {
                            v17[0] = fi56
                        }
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    v59 := GenS2{F0: -1, F1: 0, F2: (^(v1))}
                    _ = v55
                    _ = v59
                }
                v60 := GenS8{F0: 126, F1: 1, F2: (v2 ^ -40)}
                _ = v60
            }
            if len(v17) > 0 {
                v17[0] = (fi40 * fi40)
            }
        }
        switch (fi40 + fi40) {
        case -1:
            fmt.Println(v9)
        case 42:
            if len(v17) > 0 {
                v17[0] = 42
            }
        default:
            var v61 uint32 = 498
            _ = v61
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

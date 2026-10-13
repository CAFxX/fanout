package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int8; F1 uint; F2 int32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int64; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int16; F1 bool }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 uint32 = 1
    v4 := make(map[uint32]uint32)
    if ((true || false) || (-353 < -1)) {
        v1 = 0.1
    } else {
        var v5 int32 = -933
        v6 := GenS1{F0: (v5 | v5)}
        _ = v5
        _ = v6
    }
    var v7 uint32 = 211
    fmt.Println(v1)
    var v8 bool = true
    if (!true) {
        fmt.Println(v8)
        v9 := GenS2{F0: -2147483648}
        _ = v9
    }
    fmt.Println(v3)
    switch ("" + "") {
    case "\n\t":
        var v10 uint64 = 4294967295
        _ = v10
    case "a":
        if v8 {
            for fi11 := 0; fi11 < 19; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v3 += 0
                    fmt.Println(v8)
                    v13 := GenS3{F0: (32767 >= -655)}
                    if v8 {
                        if v8 {
                            var v14 bool = false
                            v15 := GenS1{F0: -320}
                            var v16 int = -1
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        switch "hello" {
                        case "\n\t":
                            v17 := GenS3{F0: false}
                            _ = v17
                        case "":
                            v1 = v1
                        default:
                            _ = len(v2)
                        }
                        v7 -= v7
                    } else {
                        v18 := GenS4{F0: ("hello" + ("foo bar" + (("foo bar" + "a") + "a"))), F1: 42}
                        v8 = v8
                        var v19 int64 = 9223372036854775807
                        delete(v2, ("a" + ""))
                        _ = v18
                        _ = v19
                    }
                    _ = fi12
                    _ = v13
                }
                var v20 float32 = ((v1 * v1) / v1)
                _ = fi11
                _ = v20
            }
        } else {
            if (-2147483648 <= -50) {
                if (v1 < (float32((3.14159 / (float64(0.1) + 3.14159))) - (float32(0.0) + v1))) {
                    switch -558 {
                    case 32767:
                        v8 = true
                    case 4294967295:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            if false {
                                fmt.Println(v7, v7, v8)
                                _ = len(v2)
                            } else {
                                if false {
                                    v22 := []int{42, fi21, fi21, (-(fi21))}
                                    if v8 {
                                        for fi23 := 0; fi23 < 10; fi23++ {
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            _ = fi23
                                        }
                                        fmt.Println(v1, v3, v3)
                                        var v24 int32 = -708
                                        _ = v24
                                    } else {
                                        if (((v8 || v8) || v8) && (false || (-40 <= 42))) {
                                            var v25 int32 = 42
                                            if (false || ((1.5 - (((3.14159 + (1.5 * (0.1 * 3.14159))) + 0.1) - (0.0 / 3.14159))) != 0.1)) {
                                                switch (fi21 << ((65535) & 63)) {
                                                case 32767:
                                                    for fi26 := 0; fi26 < 1; fi26++ {
                                                        v2[("\n\t" + "hello")] = (fi21 >> ((1) & 63))
                                                        fmt.Println(fi26, v7, fi26)
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v27 := GenS5{F0: -58, F1: 1, F2: (v25 << ((979) & 63))}
                                                            _ = v27
                                                        }
                                                        _ = fi26
                                                    }
                                                default:
                                                    v28 := GenS5{F0: 42, F1: 2, F2: v25}
                                                    _ = v28
                                                }
                                                v29 := GenS5{F0: -82, F1: 1, F2: v25}
                                                if v8 {
                                                    for fi30 := 0; fi30 < 13; fi30++ {
                                                        var v31 uint8 = 128
                                                        var v32 float32 = ((v1 * v1) / 1.5)
                                                        fmt.Println(v3, v8, fi30)
                                                        var v33 string = "foo bar"
                                                        fmt.Println(v7, v31)
                                                        _ = fi30
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    v7 = v7
                                                    v25 -= (v25 & 2147483647)
                                                }
                                                _ = v29
                                            } else {
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                                if (!(-1 < 9223372036854775807)) {
                                                    v34 := GenS6{F0: 2, F1: v8}
                                                    fmt.Println(v1, v3, v3)
                                                    switch fi21 {
                                                    case 42:
                                                        var v35 uint8 = 55
                                                        _ = v35
                                                    case -1:
                                                        var v36 uint = 4294967295
                                                        _ = v36
                                                    case 9223372036854775807:
                                                        var v37 uint = 190
                                                        _ = v37
                                                    default:
                                                        delete(v2, ("" + "\n\t"))
                                                    }
                                                    _ = v34
                                                }
                                            }
                                            _ = v25
                                        }
                                    }
                                    _ = v22
                                }
                            }
                            _ = fi21
                        }
                    case 127:
                        if ("\n\t" > ("\n\t" + "hello")) {
                            var v38 int32 = -2147483648
                            fmt.Println(v38, v1)
                            _ = v38
                        }
                    default:
                        var v39 uint8 = 255
                        _ = v39
                    }
                    v3 += (v7 - 65535)
                }
                v40 := []int{127}
                _ = v40
            }
            if true {
                fmt.Println(v1)
            } else {
                for fi41 := 0; fi41 < 11; fi41++ {
                    v42 := GenS6{F0: 368, F1: (!(872 == -1))}
                    if (false && ((0 <= 42) && (v8 && (v8 && v8)))) {
                        fmt.Println(v1, v3)
                        var v43 uint8 = 127
                        v44 := GenS7{F0: 1, F1: (-1 > 1), F2: (^(v43))}
                        _ = v43
                        _ = v44
                    }
                    if v8 {
                        v3 += 949
                        for fi45 := 0; fi45 < 13; fi45++ {
                            var v46 string = "a"
                            _ = fi45
                            _ = v46
                        }
                        v1 += v1
                    } else {
                        var v47 int16 = -32768
                        if true {
                            if v8 {
                                if ("\n\t" < ("hello" + "foo bar")) {
                                    var v48 int16 = v47
                                    _ = v48
                                }
                                var v49 string = ("\n\t" + "hello")
                                for fi50 := 0; fi50 < 13; fi50++ {
                                    if v8 {
                                        var v51 int64 = -9223372036854775808
                                        _ = v51
                                    }
                                    switch (fi50 & (^(fi50))) {
                                    case 9223372036854775807:
                                        var v52 int32 = 42
                                        _ = v52
                                    default:
                                        v53 := GenS8{F0: 255}
                                        _ = v53
                                    }
                                    v54 := GenS9{F0: (^(v7)), F1: v7}
                                    _ = fi50
                                    _ = v54
                                }
                                _ = v49
                            }
                        } else {
                            v55 := GenS10{F0: (v47 - (-(v47))), F1: ((!true) || v8)}
                            v1 -= (v1 - (v1 + v1))
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 15; fi56++ {
                            v57 := []int{}
                            v7 = (v7 % (((^(v7))) | 1))
                            v58 := v4[0]
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        _ = v47
                    }
                    _ = fi41
                    _ = v42
                }
                fmt.Println(v8, v7)
                var v59 int8 = 7
                _ = v59
            }
        }
    case "foo bar":
        for fi60 := 0; fi60 < 18; fi60++ {
            v8 = ((v8 && v8) && v8)
            if ((0.1 / float32(-2.25)) > v1) {
                if false {
                    var v61 float32 = (v1 - 0.1)
                    _ = v61
                }
            } else {
                v62 := GenS7{F0: 0, F1: (545 > 4294967295), F2: 53}
                _ = v62
            }
            v1 += v1
            _ = fi60
        }
    }
    v7 -= v7
    v3 = (v7 & (^(v7)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

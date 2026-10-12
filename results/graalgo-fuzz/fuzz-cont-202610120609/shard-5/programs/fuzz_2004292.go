package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []uint32{2147483648, 4294967295, 1, 712}
    for fi4 := 0; fi4 < 7; fi4++ {
        _ = len(v3)
        switch fi4 {
        case -457:
            fmt.Println(fi4)
        default:
            var v5 uint = 472
            _ = v5
        }
        for fi6 := 0; fi6 < 17; fi6++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                if (478 < 991) {
                    v7 := GenS1{F0: ("hello" + "foo bar"), F1: 42}
                    _ = v7
                } else {
                    var v8 uint8 = 127
                    if (v8 >= (v8 & 149)) {
                        v8 = (v8 << ((255) & 63))
                        var v9 uint32 = 338
                        v8 -= v8
                        _ = v9
                    } else {
                        var v10 string = ("\n\t" + ("a" + "foo bar"))
                        v8 = v8
                        if (true || true) {
                            var v11 float64 = 0.1
                            _ = v11
                        } else {
                            v8 -= (v8 & (v8 | 127))
                            if (1000 < 0) {
                                var v12 uint32 = 433
                                _ = v12
                            }
                            if (146 < -1) {
                                var v13 uint16 = 2
                                _ = v13
                            }
                        }
                        _ = v10
                    }
                    if true {
                        var v14 uint16 = 42
                        v8 = (v8 & (v8 << ((255) & 63)))
                        v8 = (^(v8))
                        fmt.Println(fi4, fi4)
                        _ = v14
                    } else {
                        v15 := GenS1{F0: ((("\n\t" + "a") + "foo bar") + ("foo bar" + "foo bar")), F1: 919}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v15
                    }
                    _ = v8
                }
            }
            if (float64((float32(3.14159) - (1e10 / 0.1))) <= -2.25) {
                if len(v3) > 0 {
                    v3[0] = 325
                }
                var v16 int8 = 51
                fmt.Println(fi4, fi6, v16)
                _ = v16
            }
            var v17 float64 = 0.0
            _ = fi6
            _ = v17
        }
        _ = fi4
    }
    var v18 int64 = 1
    v18 -= v18
    switch 32768 {
    case 9223372036854775807:
        if ((!(4294967295 == 404)) || (4294967295 <= 32767)) {
            fmt.Println(v18)
            v18 += -142
        }
    case -9223372036854775808:
        v18 = 9223372036854775807
    case -294:
        v19 := GenS1{F0: "hello", F1: 1}
        _ = v19
    default:
        v18 += (v18 >> ((396) & 63))
    }
    switch -527 {
    case 32767:
        if false {
            fmt.Println(v18)
            fmt.Println(v18)
        } else {
            fmt.Println(v18)
        }
    default:
        fmt.Println(v18)
    }
    v18 += v18
    for fi20 := 0; fi20 < 19; fi20++ {
        v18 -= -9223372036854775808
        switch fi20 {
        case -1:
            switch "foo bar" {
            case "a":
                _ = len(v2)
            case "":
                fmt.Println(v18, fi20)
            case "hello":
                var v21 string = ((((("hello" + ("foo bar" + "foo bar")) + (((((("a" + "hello") + ("foo bar" + ("" + ("hello" + "a")))) + "a") + "\n\t") + "a") + "a")) + "") + "\n\t") + "a")
                _ = v21
            default:
                v22 := GenS1{F0: ("foo bar" + "hello"), F1: 722}
                _ = v22
            }
        case 32768:
            var v23 int16 = 32767
            _ = v23
        }
        _ = fi20
    }
    switch "hello" {
    case "\n\t":
        var v24 float64 = 1e10
        _ = v24
    case "a":
        var v25 int64 = v18
        _ = v25
    case "foo bar":
        var v26 int = 4294967295
        _ = v26
    default:
        delete(v2, 42)
    }
    v18 = v18
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v27 uint8 = 108
    v27 = v27
    v27 -= 253
    var v28 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}

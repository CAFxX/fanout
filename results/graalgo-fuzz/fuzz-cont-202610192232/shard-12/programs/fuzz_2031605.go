package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint8; F1 int16; F2 int32 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint64; F1 int; F2 uint64 }
type GenS8 struct { F0 int32; F1 int16; F2 uint }
func genFn0() int16 {
    v5 := GenS1{F0: 665}
    var v6 uint = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        switch (fi7 >> ((v6) & 63)) {
        case 4294967295:
            var v8 uint8 = 127
            _ = v8
        }
        _ = fi7
    }
    var v9 int8 = 1
    v10 := []int{329}
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 42
    var v3 int16 = -695
    var v4 uint64 = 623
    var v11 uint = 18446744073709551615
    switch "hello" {
    case "a":
        if (v4 != 18446744073709551615) {
            switch -800 {
            case 0:
                v12 := []int{661, 127, -1, -45}
                _ = v12
            }
        }
    case "hello":
        v13 := GenS2{F0: v4, F1: 32767}
        _ = v13
    case "\n\t":
        v2 = 42
    default:
        if false {
            var v14 int32 = 216
            v15 := []int{128, int(v3), 9223372036854775807, 341}
            _ = v14
            _ = v15
        } else {
            v16 := v1[-275]
            switch "" {
            case "foo bar":
                for fi17 := 0; fi17 < 18; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 uint32 = 4294967295
                        v11 = (v11 >> ((410) & 63))
                        if ((v2 * v2) <= (v2 + (v2 << ((255) & 63)))) {
                            if (3.14159 != (1.5 - 0.0)) {
                                v11 = v11
                                v20 := GenS2{F0: (^(v4)), F1: 2}
                                _ = v20
                            }
                            v21 := GenS3{F0: v16, F1: ((v4 << ((v11) & 63)) == (^(v4)))}
                            var v22 int32 = v2
                            if v16 {
                                var v23 int16 = (v3 | 32767)
                                _ = v23
                            }
                            fmt.Println(fi17)
                            _ = v21
                            _ = v22
                        } else {
                            v3 += v3
                            var v24 uint64 = (v4 / ((1) | 1))
                            if v16 {
                                v25 := GenS3{F0: v16, F1: false}
                                var v26 uint = v11
                                _ = v25
                                _ = v26
                            } else {
                                var v27 uint16 = 1
                                var v28 uint = (v11 + 18446744073709551615)
                                v29 := GenS4{F0: v3, F1: "\n\t"}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = v24
                        }
                        var v30 int16 = (v3 >> ((2) & 63))
                        fmt.Println(fi17, v3)
                        _ = fi18
                        _ = v19
                        _ = v30
                    }
                    _ = fi17
                }
            case "":
                if false {
                    v2 += v2
                    v11 += 65535
                }
            case "\n\t":
                fmt.Println(v2, v11)
            default:
                var v31 int64 = 1
                _ = v31
            }
            v2 = v2
            _ = v16
        }
    }
    if (336 == 858) {
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 int32 = -2147483648
            if ((!(!(1 >= 128))) || false) {
                for fi34 := 0; fi34 < 10; fi34++ {
                    v4 -= (v4 | 922)
                    if true {
                        fmt.Println(fi34, fi34)
                        var v35 uint16 = 2
                        v36 := GenS5{F0: 128, F1: (v3 >> ((v11) & 63)), F2: (v33 + v2)}
                        _ = v35
                        _ = v36
                    }
                    v3 -= 32767
                    _ = fi34
                }
            } else {
                v37 := []int{}
                _ = v37
            }
            _ = fi32
            _ = v33
        }
        var v38 float32 = ((3.14159 - (float32(1.5) * 1.5)) / 1e10)
        v39 := GenS1{F0: 9223372036854775807}
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 5; fi40++ {
            v11 = v11
            var v41 int = fi40
            _ = fi40
            _ = v41
        }
        v4 = v4
    }
    v42 := GenS6{F0: 42, F1: 0}
    v43 := GenS7{F0: (^(v4)), F1: 1, F2: v4}
    if (false || false) {
        v44 := GenS3{F0: true, F1: false}
        if (65535 < 2) {
            v1[-468] = true
            switch -193 {
            case 2147483647:
                for fi45 := 0; fi45 < 5; fi45++ {
                    v46 := GenS2{F0: (v4 * 762), F1: 536}
                    if (255 == 84) {
                        v47 := []int{}
                        fmt.Println(v4, v2)
                        _ = v47
                    } else {
                        var v48 string = ""
                        switch (("a" + "foo bar") + (v48 + "foo bar")) {
                        case "\n\t":
                            delete(v1, 2147483647)
                        case "hello":
                            var v49 string = "a"
                            _ = v49
                        default:
                            fmt.Println(v11, v11)
                        }
                        _ = v48
                    }
                    _ = fi45
                    _ = v46
                }
            }
        } else {
            var v50 int = -9223372036854775808
            var v51 int8 = -17
            _ = v50
            _ = v51
        }
        _ = v44
    }
    switch 127 {
    case 970:
        v52 := GenS8{F0: -140, F1: (v3 >> (((v11 ^ v11)) & 63)), F2: 4294967295}
        _ = v52
    }
    fmt.Println(v4, v2)
    v2 = 42
    for fi53 := 0; fi53 < 11; fi53++ {
        v4 = (v4 ^ 868)
        if false {
            for fi54 := 0; fi54 < 7; fi54++ {
                if ((false || (!(true && (false && false)))) && true) {
                    switch 42 {
                    case 32767:
                        v3 += (v3 + (v3 << ((v11) & 63)))
                    }
                    v55 := GenS1{F0: 242}
                    fmt.Println(fi54, v3)
                    _ = v55
                }
                v3 = v3
                v56 := []int{fi53, fi54, 32768, fi54}
                _ = fi54
                _ = v56
            }
            delete(v1, (^(fi53)))
        }
        var v57 float32 = (1.5 + -2.25)
        _ = fi53
        _ = v57
    }
    var v58 int8 = 106
    v58 -= v58
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v58
}

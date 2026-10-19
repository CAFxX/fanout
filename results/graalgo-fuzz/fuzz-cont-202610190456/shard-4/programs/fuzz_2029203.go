package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0() int32 {
    if false {
        for fi4 := 0; fi4 < 13; fi4++ {
            var v5 uint8 = 2
            var v6 int8 = 42
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 uint32 = 1
        _ = v7
    }
    var v8 int16 = 1
    _ = v8
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 1
    v3 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float64 = 3.14159
        _ = fi9
        _ = v10
    }
    v2 -= v2
    v11 := GenS1{F0: -32768, F1: 969, F2: 65535}
    v2 = (^(v2))
    for fi12 := 0; fi12 < 11; fi12++ {
        v2 -= 128
        _ = fi12
    }
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 12; fi14++ {
        v2 = 128
        for fi15 := 0; fi15 < 17; fi15++ {
            if false {
                v16 := GenS1{F0: 1, F1: 65535, F2: v13}
                v17 := GenS1{F0: 1, F1: 32767, F2: (v13 - (v13 >> (((v13 | 4294967295)) & 63)))}
                v18 := GenS2{F0: (v13 / (((v13 >> ((v13) & 63))) | 1)), F1: 393}
                _ = v16
                _ = v17
                _ = v18
            }
            v2 += (v2 % (((v2 >> (((v13 + (^(v13)))) & 63))) | 1))
            _ = fi15
        }
        switch fi14 {
        case -356:
            for fi19 := 0; fi19 < 12; fi19++ {
                if ((true || false) && ((false || (65535 == 759)) && true)) {
                    v13 += 65535
                    if (125 != 69) {
                        var v20 uint64 = 42
                        _ = v20
                    } else {
                        switch (fi14 - (fi14 | 128)) {
                        case 42:
                            v13 = (v13 | (v13 + v13))
                        case -809:
                            switch "" {
                            case "a":
                                var v21 int64 = -425
                                _ = v21
                            case "\n\t":
                                if false {
                                    var v22 float64 = (0.1 * 0.0)
                                    delete(v1, (fi19 ^ (-(fi19))))
                                    switch (fi19 >> ((161) & 63)) {
                                    case -9223372036854775808:
                                        if true {
                                            v13 += v13
                                            var v23 uint32 = 677
                                            switch ("hello" + (("foo bar" + "\n\t") + (("" + (("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + "hello")) + "foo bar"))) {
                                            case "":
                                                v24 := []int{fi19, 128}
                                                _ = v24
                                            case "hello":
                                                var v25 uint16 = 32767
                                                _ = v25
                                            default:
                                                v26 := GenS2{F0: v13, F1: v23}
                                                _ = v26
                                            }
                                            _ = v23
                                        }
                                    case -999:
                                        v27 := v3[2]
                                        _ = v27
                                    case 42:
                                        v2 += (v2 << ((v13) & 63))
                                    }
                                    v2 = (v2 << ((v13) & 63))
                                    _ = v22
                                }
                            }
                        case 0:
                            var v28 int32 = 1
                            _ = v28
                        }
                    }
                }
                v29 := GenS3{F0: uint(v2), F1: 887, F2: "hello"}
                v13 -= v13
                _ = fi19
                _ = v29
            }
        case -899:
            if false {
                if true {
                    v30 := GenS4{F0: v13, F1: -1}
                    if ((true && false) && false) {
                        fmt.Println(v13)
                    }
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        if (((!(4294967295 < 2)) || false) && true) {
                            var v32 float32 = (float32(-2.25) / 1.5)
                            fmt.Println(v13, v2, v13)
                            if false {
                                if (0 == 9223372036854775807) {
                                    var v33 string = "hello"
                                    if (826 >= 964) {
                                        var v34 int = fi14
                                        v35 := v3[823]
                                        v35 = v35
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v3[1] = (2 <= 65535)
                                    }
                                    _ = v33
                                } else {
                                    if (!("a" >= "hello")) {
                                        v2 -= 42
                                        var v36 int = fi31
                                        var v37 int8 = -1
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 uint8 = 128
                                        _ = v38
                                    }
                                }
                                v13 -= v13
                                var v39 string = "hello"
                                _ = v39
                            } else {
                                v2 -= (v2 ^ (^(v2)))
                                v40 := []int{(fi31 << (((v13 >> (((v13 / ((v13) | 1))) & 63))) & 63)), (fi14 >> (((v13 | (v13 << (((v13 / ((v13) | 1))) & 63)))) & 63)), (fi14 * fi14)}
                                _ = v40
                            }
                            var v41 int8 = 42
                            var v42 string = (((("hello" + "a") + "\n\t") + ("hello" + "\n\t")) + ("\n\t" + "\n\t"))
                            var v43 string = "hello"
                            _ = v32
                            _ = v41
                            _ = v42
                            _ = v43
                        } else {
                            v44 := GenS2{F0: v13, F1: 2147483648}
                            v13 += (v13 % ((42) | 1))
                            var v45 int8 = -45
                            if (false && false) {
                                v46 := GenS3{F0: (v13 % ((v13) | 1)), F1: 32767, F2: "\n\t"}
                                _ = v46
                            }
                            v47 := []int{fi31}
                            _ = v44
                            _ = v45
                            _ = v47
                        }
                        var v48 int32 = genFn0()
                        v49 := GenS1{F0: 32767, F1: 32767, F2: (v13 - (^(v13)))}
                        _ = fi31
                        _ = v48
                        _ = v49
                    }
                    v2 += 87
                    for fi50 := 0; fi50 < 19; fi50++ {
                        switch "foo bar" {
                        case "a":
                            var v51 int8 = 1
                            _ = v51
                        }
                        _ = fi50
                    }
                }
                fmt.Println(v2, v2, v2)
                var v52 bool = (18446744073709551615 < uint64(v13))
                _ = v52
            } else {
                fmt.Println(fi14)
                v13 = (^(v13))
                v13 += 255
            }
        case 9223372036854775807:
            var v53 uint16 = 1
            _ = v53
        }
        _ = fi14
    }
    switch -2147483648 {
    case -240:
        if (2 < 42) {
            v1[56] = 126
            fmt.Println(v2, v13)
        } else {
            v13 += (^(v13))
            v1[127] = 4294967295
            v54 := GenS3{F0: (v13 ^ v13), F1: 255, F2: ("hello" + ("a" + ("\n\t" + "foo bar")))}
            _ = v54
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
}

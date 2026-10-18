package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    var v1 int = 32767
    var v2 bool = true
    var v3 string = "foo bar"
    if (0 < int64(v1)) {
        for fi4 := 0; fi4 < 18; fi4++ {
            v5 := []int{v1, v1}
            for fi6 := 0; fi6 < 12; fi6++ {
                v2 = (v2 || v2)
                _ = len(v5)
                var v7 uint = 4294967295
                switch -929 {
                case 1:
                    fmt.Println(v3, v1)
                    v8 := GenS1{F0: -128, F1: ("hello" + ("a" + v3)), F2: 2}
                    _ = v8
                case -933:
                    v5 = append(v5, fi4)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v9 uint16 = 694
                    if (1 > -1) {
                        var v10 int32 = 1
                        if (!(!v2)) {
                            var v11 uint64 = 42
                            _ = v11
                        }
                        _ = v10
                    }
                    fmt.Println(v2)
                    v2 = (-1 < -2147483648)
                    _ = v9
                }
                var v12 int = (v1 << (((v7 * v7)) & 63))
                var v13 uint = (v7 + 256)
                if v2 {
                    v7 = (v7 >> ((v13) & 63))
                } else {
                    var v14 string = (v3 + (v3 + (v3 + v3)))
                    _ = v14
                }
                _ = fi6
                _ = v12
                _ = v13
                _ = v7
            }
            var v15 string = "hello"
            _ = fi4
            _ = v15
            _ = v5
        }
    }
    v2 = (!v2)
    switch 386 {
    case -176:
        var v16 uint32 = 4294967295
        _ = v16
    case -611:
        var v17 uint = 0
        _ = v17
    case 2147483647:
        v2 = v2
    }
    v1 -= 509
    if (-2147483648 >= -160) {
        if v2 {
            var v18 uint8 = 128
            v19 := GenS1{F0: 127, F1: v3, F2: 4294967295}
            v20 := []int{128, v1, (v1 ^ (v1 * v1)), v1}
            _ = v18
            _ = v19
            _ = v20
        } else {
            var v21 int32 = -2147483648
            v2 = v2
            v21 = v21
            _ = v21
        }
        v3 = (("foo bar" + v3) + v3)
    }
    if v2 {
        v22 := []int{v1, 1}
        v2 = (741 == 1)
        var v23 uint16 = 32767
        _ = v22
        _ = v23
    }
    v24 := []int{v1}
    v3 = v3
    if len(v24) > 0 {
        v24[0] = -9223372036854775808
    }
    switch "\n\t" {
    case "hello":
        if (!(1 != 198)) {
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
    case "a":
        var v25 int32 = -633
        _ = v25
    case "foo bar":
        for fi26 := 0; fi26 < 14; fi26++ {
            v2 = ((-9223372036854775808 != 1) && ((!v2) || (65535 != 2)))
            if len(v24) > 0 {
                v24[0] = -67
            }
            _ = fi26
        }
    default:
        for fi27 := 0; fi27 < 4; fi27++ {
            if (484 <= 13) {
                var v28 bool = (!v2)
                for fi29 := 0; fi29 < 5; fi29++ {
                    v3 = (v3 + "hello")
                    _ = fi29
                }
                _ = v28
            }
            if v2 {
                v24 = append(v24, (fi27 - 127))
                if len(v24) > 0 {
                    v24[0] = v1
                }
            } else {
                if true {
                    var v30 int8 = -39
                    var v31 int16 = 1
                    if true {
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        v2 = (((((1.5 / 1.5) * 3.14159) + 1.5) * ((0.1 / -2.25) / 0.1)) > (1e10 * 0.1))
                    }
                    _ = v30
                    _ = v31
                }
                if len(v24) > 0 {
                    v24[0] = 32767
                }
            }
            _ = fi27
        }
    }
    v2 = true
    v32 := GenS1{F0: 127, F1: "", F2: 377}
    v3 = "a"
    switch "\n\t" {
    case "\n\t":
        for fi33 := 0; fi33 < 4; fi33++ {
            if (-255 != 42) {
                var v34 uint32 = 65535
                _ = v34
            } else {
                switch (v3 + "a") {
                case "hello":
                    switch (^(fi33)) {
                    case 127:
                        if len(v24) > 0 {
                            v24[0] = fi33
                        }
                    case -9223372036854775808:
                        v35 := GenS2{F0: 9223372036854775807}
                        _ = v35
                    default:
                        v1 -= (v1 * v1)
                    }
                default:
                    if v2 {
                        v2 = (v2 || (1 < 702))
                    } else {
                        var v36 uint16 = 592
                        v37 := GenS3{F0: 127, F1: "foo bar"}
                        _ = v36
                        _ = v37
                    }
                }
                fmt.Println(v2)
                var v38 int8 = -1
                _ = v38
            }
            for fi39 := 0; fi39 < 1; fi39++ {
                v40 := GenS4{F0: -2147483648, F1: int32(fi39)}
                _ = fi39
                _ = v40
            }
            _ = fi33
        }
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        v2 = (v2 && (false || (2 <= 255)))
        fmt.Println(v3, fi41, fi41)
        _ = fi41
    }
    if (-2.25 == (3.14159 / -2.25)) {
        v3 = v3
        fmt.Println(v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
}

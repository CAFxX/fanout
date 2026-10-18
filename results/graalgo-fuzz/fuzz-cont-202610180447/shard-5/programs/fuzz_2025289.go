package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    var v1 uint32 = 301
    var v2 int16 = 32767
    v3 := []int{-184}
    var v4 int16 = (v2 >> ((2) & 63))
    v4 = v2
    switch 42 {
    case 9223372036854775807:
        v5 := GenS1{F0: (v1 | v1), F1: -1}
        _ = v5
    case 128:
        switch 1 {
        case -817:
            var v6 uint64 = uint64(v1)
            _ = v6
        case 32767:
            var v7 uint = 245
            _ = v7
        default:
            switch 32768 {
            case 502:
                var v8 uint32 = v1
                _ = v8
            default:
                v2 = (-(v2))
            }
        }
    case 4294967295:
        v9 := GenS1{F0: 2147483647, F1: 0}
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            v11 := GenS2{F0: 587}
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(fi12)
                _ = fi12
            }
            _ = v11
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 -= (v4 ^ v2)
                } else {
                    if (856 >= 2) {
                        v4 = 654
                    }
                    fmt.Println(fi13, v1, fi13)
                    var v14 bool = (1 != (^(v4)))
                    _ = v14
                }
                _ = fi13
            }
            var v15 string = "hello"
            var v16 bool = false
            _ = v15
            _ = v16
        }
        _ = fi10
    }
    if (v1 > 4294967295) {
        var v17 float32 = 1.5
        if (true || (127 <= -2147483648)) {
            if (true || true) {
                v17 = -2.25
            }
            var v18 int8 = -28
            _ = v18
        }
        _ = v17
    } else {
        if (899 < 0) {
            if (!false) {
                if (!true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 -= v1
                    v1 = 2147483648
                }
            } else {
                v19 := GenS3{F0: 65535, F1: v1}
                v3 = append(v3, 2147483647)
                for fi20 := 0; fi20 < 9; fi20++ {
                    v4 = v4
                    if true {
                        if false {
                            v21 := GenS4{F0: 950, F1: 1, F2: 255}
                            var v22 float64 = (1.5 - ((float64(3.14159) / -2.25) - 1e10))
                            _ = v21
                            _ = v22
                        } else {
                            v2 += v2
                            var v23 string = "foo bar"
                            _ = v23
                        }
                    } else {
                        _ = len(v3)
                    }
                    if false {
                        if (18446744073709551615 == 1) {
                            fmt.Println(fi20, v1)
                        }
                    }
                    _ = fi20
                }
                _ = v19
            }
            v4 = 651
            if len(v3) > 0 {
                v3[0] = 904
            }
        } else {
            switch ((("hello" + (("\n\t" + "hello") + "hello")) + ("" + ((((((("a" + "") + "\n\t") + "foo bar") + "foo bar") + "foo bar") + "") + (("a" + "foo bar") + "\n\t")))) + (((("" + (("hello" + "a") + "\n\t")) + "\n\t") + ("foo bar" + "hello")) + "")) {
            case "a":
                var v24 string = ""
                _ = v24
            default:
                if false {
                    var v25 bool = true
                    _ = v25
                } else {
                    var v26 int16 = v4
                    v1 -= v1
                    if (true || true) {
                        v26 = (^(v4))
                        v4 = (v26 % (((^(v2))) | 1))
                        v3 = append(v3, int(v4))
                    }
                    _ = v26
                }
            }
        }
        if true {
            switch "hello" {
            case "\n\t":
                v4 += v4
            case "":
                v27 := GenS3{F0: 822, F1: v1}
                _ = v27
            case "hello":
                v3 = append(v3, 9223372036854775807)
            }
            fmt.Println(v1)
        } else {
            v2 += v4
            switch ("" + "a") {
            case "hello":
                var v28 uint = 425
                _ = v28
            case "a":
                if (-1 >= 0) {
                    var v29 uint8 = 207
                    _ = v29
                }
            }
            if true {
                var v30 uint64 = 240
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v30
            }
        }
    }
    v2 += v4
    switch -1 {
    case 32767:
        var v31 uint16 = 42
        _ = v31
    case -581:
        for fi32 := 0; fi32 < 18; fi32++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = (fi32 & fi32)
            }
            _ = fi32
        }
    case -93:
        v33 := GenS1{F0: (v1 & (v1 ^ v1)), F1: -178}
        _ = v33
    default:
        v34 := GenS1{F0: (v1 - (v1 ^ v1)), F1: 0}
        _ = v34
    }
    var v35 int8 = 0
    switch "hello" {
    case "hello":
        if (true && true) {
            if (v4 >= (v2 + -541)) {
                v4 = (^(v2))
                _ = len(v3)
                fmt.Println(v2)
            } else {
                if (v1 >= (v1 << ((255) & 63))) {
                    switch -1 {
                    case -2147483648:
                        v2 = v2
                    case 42:
                        fmt.Println(v4, v35, v1)
                    case -1:
                        v35 = v35
                    }
                    if (42 == 1) {
                        if (42 > 2) {
                            v35 -= v35
                            fmt.Println(v35, v35, v2)
                            v36 := GenS4{F0: 255, F1: 409, F2: 102}
                            _ = v36
                        } else {
                            v37 := GenS4{F0: 2, F1: -1, F2: 2}
                            if (!((255 <= 32767) || (-1 > -1))) {
                                var v38 uint16 = 256
                                _ = v38
                            }
                            _ = v37
                        }
                        v1 = (v1 - (v1 / ((2147483647) | 1)))
                        var v39 bool = true
                        _ = v39
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        } else {
            if false {
                for fi40 := 0; fi40 < 2; fi40++ {
                    v1 -= (v1 << ((1) & 63))
                    _ = fi40
                }
                v2 = 1
            }
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            v35 = v35
        }
    default:
        if (-509 <= -2147483648) {
            switch "\n\t" {
            case "":
                for fi41 := 0; fi41 < 17; fi41++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi41
                }
            }
            v35 += (-(v35))
        }
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        switch (fi42 >> ((4294967295) & 63)) {
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            fmt.Println(v1)
        case -550:
            var v43 uint = 4294967295
            _ = v43
        }
        var v44 uint = 2
        var v45 uint16 = 973
        _ = fi42
        _ = v44
        _ = v45
    }
    v1 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}

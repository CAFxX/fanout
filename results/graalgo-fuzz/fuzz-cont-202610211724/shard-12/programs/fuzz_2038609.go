package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 uint32; F1 bool; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 string; F1 uint8; F2 uint32 }
type GenS9 struct { F0 int }
func genFn0() uint8 {
    v4 := []int{9223372036854775807}
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = (v5 >> ((42) & 63))
        _ = fi6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        _ = len(v4)
        _ = fi7
    }
    _ = v4
    _ = v5
    return 232
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 9223372036854775807
    var v3 int8 = -126
    v2 = v2
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            if true {
                var v10 bool = (799 >= 273)
                v11 := []int{(fi8 | int(v3)), fi8, fi8, fi9}
                switch (-(fi9)) {
                case 4294967295:
                    v12 := GenS1{F0: 126, F1: 433}
                    _ = v12
                case -545:
                    if ("a" > "foo bar") {
                        v13 := GenS2{F0: ((("\n\t" + "foo bar") + "hello") + "hello"), F1: (-(v3)), F2: v2}
                        _ = v13
                    } else {
                        v14 := GenS3{F0: 2, F1: 18446744073709551615, F2: -986}
                        _ = v14
                    }
                case 2147483647:
                    switch (("\n\t" + ("foo bar" + "hello")) + ("\n\t" + "foo bar")) {
                    case "\n\t":
                        v11 = append(v11, (fi9 | fi9))
                    }
                default:
                    var v15 uint8 = 232
                    _ = v15
                }
                fmt.Println(v10)
                v16 := GenS4{F0: (fi9 % (((fi9 + (fi9 ^ 58))) | 1)), F1: "a", F2: 255}
                if (!v10) {
                    var v17 uint = 194
                    _ = v17
                }
                _ = v10
                _ = v11
                _ = v16
            }
            var v18 uint16 = 436
            _ = fi9
            _ = v18
        }
        _ = fi8
    }
    v2 = (v2 ^ v2)
    if false {
        v19 := []int{-1, 0}
        v19 = append(v19, -9223372036854775808)
        switch "\n\t" {
        case "foo bar":
            var v20 int64 = 216
            _ = v20
        case "":
            var v21 bool = (!(v2 != 42))
            _ = v21
        }
        _ = v19
    }
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := GenS5{F0: v1, F1: false, F2: 6}
            var v24 uint8 = 42
            var v25 int16 = 0
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        if true {
            v1 = v1
            fmt.Println(v3)
        } else {
            v2 = 567
            switch 128 {
            case 0:
                v26 := GenS5{F0: (v1 | v1), F1: (-157 > 1), F2: -32768}
                _ = v26
            default:
                for fi27 := 0; fi27 < 18; fi27++ {
                    switch ("hello" + (("hello" + "hello") + (((("a" + ("a" + "foo bar")) + "hello") + "\n\t") + (("foo bar" + "a") + "\n\t")))) {
                    case "a":
                        var v28 uint16 = 42
                        _ = v28
                    default:
                        if true {
                            if (v1 < 42) {
                                v2 -= v2
                                var v29 int32 = -2147483648
                                _ = v29
                            } else {
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    if (871 > 42) {
                                        v31 := []int{-462, fi30, -317}
                                        fmt.Println(fi30)
                                        var v32 uint = 4294967295
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 uint = 255
                                    if true {
                                        var v34 int8 = v3
                                        v35 := []int{fi30}
                                        _ = v34
                                        _ = v35
                                    }
                                    fmt.Println(v1, fi30, fi27)
                                    var v36 uint8 = 66
                                    _ = fi30
                                    _ = v33
                                    _ = v36
                                }
                            }
                        }
                    }
                    v37 := GenS6{F0: -116}
                    _ = fi27
                    _ = v37
                }
            }
            v1 -= (v1 / ((v1) | 1))
        }
    }
    var v38 int = 4294967295
    v39 := GenS7{F0: 282, F1: 284}
    v3 += (v3 ^ (v3 / ((v3) | 1)))
    v40 := GenS8{F0: "a", F1: genFn0(), F2: v1}
    fmt.Println(v1)
    if ((!false) || (v3 == -8)) {
        v41 := []int{}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 16; fi42++ {
            v3 -= 0
            var v43 string = "foo bar"
            _ = fi42
            _ = v43
        }
        v44 := GenS9{F0: v38}
        _ = v44
    }
    v38 += (v38 / (((v38 % ((-620) | 1))) | 1))
    v1 += 836
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}

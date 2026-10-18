package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 int32 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 bool; F1 int16 }

func main() {
    var v1 int32 = 1
    var v2 string = "foo bar"
    var v3 uint = 4294967295
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint32 = 65535
        v6 := GenS1{F0: 1, F1: v3}
        v7 := []int{(fi4 & (-(fi4))), int(v1)}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    if false {
        var v8 bool = (!(true && (!(465 >= 42))))
        _ = v8
    }
    if (!false) {
        v2 = "foo bar"
        v9 := GenS2{F0: 32767}
        v10 := GenS2{F0: -1}
        _ = v10
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint32 = 773
            v13 := GenS3{F0: (v3 % (((v3 - (^(v3)))) | 1)), F1: -9223372036854775808}
            _ = fi11
            _ = v12
            _ = v13
        }
        switch (("hello" + v2) + v2) {
        case "hello":
            fmt.Println(v2, v3)
        case "":
            v3 = (v3 + (v3 - 255))
        case "foo bar":
            var v14 uint8 = 203
            _ = v14
        }
    }
    v15 := GenS4{F0: 42, F1: (v3 * (v3 * v3)), F2: 1}
    switch 42 {
    case -982:
        v1 = (v1 << (((^(v3))) & 63))
    case 2147483647:
        if (v2 != v2) {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int8 = 48
                switch fi16 {
                case 301:
                    fmt.Println(fi16, v1)
                }
                _ = fi16
                _ = v17
            }
        }
    default:
        for fi18 := 0; fi18 < 4; fi18++ {
            v2 = "hello"
            var v19 int = fi18
            switch (fi18 << ((2) & 63)) {
            case -2147483648:
                var v20 string = "hello"
                _ = v20
            case 127:
                v21 := GenS5{F0: 2, F1: v19, F2: (v1 >> ((v3) & 63))}
                _ = v21
            default:
                fmt.Println(v2, v3, fi18)
            }
            _ = fi18
            _ = v19
        }
    }
    if (!false) {
        v22 := GenS6{F0: v3, F1: 0}
        v23 := []int{-811, 712}
        v3 = v3
        _ = v22
        _ = v23
    }
    fmt.Println(v2, v1, v1)
    if true {
        v24 := []int{127, 128}
        v3 -= 42
        _ = v24
    } else {
        for fi25 := 0; fi25 < 1; fi25++ {
            switch (fi25 >> (((^(v3))) & 63)) {
            case -9223372036854775808:
                v1 -= v1
            case -247:
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 int = (fi26 - (fi25 + (fi25 % (((-(fi25))) | 1))))
                    var v28 int8 = -82
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                v30 := []int{fi25, fi29, -845}
                v1 = (-(v1))
                fmt.Println(v2, v2, fi29)
                _ = fi29
                _ = v30
            }
            _ = fi25
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int = fi31
            if true {
                for fi33 := 0; fi33 < 5; fi33++ {
                    var v34 bool = (!(2 > 42))
                    v35 := GenS2{F0: 32767}
                    if v34 {
                        v34 = v34
                        if v34 {
                            var v36 bool = (0 > 0)
                            _ = v36
                        } else {
                            if (!false) {
                                if v34 {
                                    if v34 {
                                        v3 -= v3
                                        if false {
                                            v37 := GenS7{F0: 255}
                                            if (true || v34) {
                                                v1 += (v1 / ((1) | 1))
                                                v2 = v2
                                            }
                                            v3 += (v3 >> (((v3 >> ((v3) & 63))) & 63))
                                            v3 -= v3
                                            _ = v37
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v1, fi33)
                            }
                            if (v1 >= (v1 * v1)) {
                                switch (^(v32)) {
                                case 128:
                                    v32 = fi33
                                case 4294967295:
                                    var v38 int64 = 651
                                    _ = v38
                                case 2147483647:
                                    var v39 float32 = 1e10
                                    _ = v39
                                }
                                switch v2 {
                                case "\n\t":
                                    switch "" {
                                    case "a":
                                        v40 := GenS8{F0: 4294967295, F1: (!(!(-1 > 0)))}
                                        _ = v40
                                    case "hello":
                                        v41 := GenS6{F0: (v3 + v3), F1: 0}
                                        _ = v41
                                    }
                                    var v42 uint = v3
                                    switch v32 {
                                    case 32767:
                                        var v43 uint = (v3 | v3)
                                        _ = v43
                                    default:
                                        var v44 uint32 = 464
                                        _ = v44
                                    }
                                    _ = v42
                                case "foo bar":
                                    v45 := GenS2{F0: 1}
                                    _ = v45
                                case "":
                                    v46 := []int{(fi31 << ((18446744073709551615) & 63)), (v32 / (((-(fi31))) | 1)), 9223372036854775807, 1}
                                    _ = v46
                                default:
                                    switch fi33 {
                                    case 127:
                                        if (v3 <= v3) {
                                            v47 := GenS9{F0: v34, F1: 1}
                                            _ = v47
                                        }
                                    case 2147483647:
                                        var v48 int32 = -45
                                        _ = v48
                                    case 42:
                                        var v49 float32 = (float32((1.5 + 0.1)) / ((float32((0.1 / 0.1)) * 0.1) - -2.25))
                                        _ = v49
                                    }
                                }
                                v32 -= (-(v32))
                                var v50 uint32 = 442
                                fmt.Println(v3, v2)
                                _ = v50
                            }
                            v2 = (v2 + "a")
                        }
                        v3 = v3
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                v3 += (v3 * v3)
            }
            _ = fi31
            _ = v32
        }
    }
    var v51 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v51
}

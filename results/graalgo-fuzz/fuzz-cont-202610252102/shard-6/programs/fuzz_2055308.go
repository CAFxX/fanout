package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
func genFn0() int16 {
    v4 := GenS1{F0: 1, F1: 255}
    var v5 uint16 = 2
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 1
    var v3 uint8 = 160
    v6 := []int{42}
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                fmt.Println(fi7, v1, v3)
                _ = fi8
            }
            _ = fi7
        }
    }
    if true {
        v9 := GenS1{F0: 528, F1: 1}
        _ = v9
    } else {
        if (v1 != v1) {
            for fi10 := 0; fi10 < 12; fi10++ {
                if ((-9223372036854775808 == 9223372036854775807) && false) {
                    v1 -= (v1 & v1)
                }
                fmt.Println(v2)
                fmt.Println(v3, v2)
                _ = fi10
            }
            if false {
                fmt.Println(v3, v2)
                if (705 == 255) {
                    if true {
                        v1 += v1
                        fmt.Println(v1, v1, v1)
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 uint16 = 0
                            if (!true) {
                                var v13 int16 = 42
                                _ = v13
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                    switch ("\n\t" + ("" + ("foo bar" + ("a" + "")))) {
                    case "a":
                        var v14 int32 = 42
                        _ = v14
                    case "\n\t":
                        if ((0 > uint16(v2)) && false) {
                            for fi15 := 0; fi15 < 1; fi15++ {
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    var v17 float64 = 0.1
                                    if true {
                                        fmt.Println(fi16)
                                        _ = len(v6)
                                        v2 = (v2 - 214)
                                        fmt.Println(v1, v3)
                                        var v18 float32 = (float32(v17) * 0.0)
                                        _ = v18
                                    }
                                    v19 := GenS2{F0: 0}
                                    var v20 int32 = -491
                                    _ = fi16
                                    _ = v17
                                    _ = v19
                                    _ = v20
                                }
                                v3 += 132
                                _ = fi15
                            }
                            var v21 int8 = 0
                            _ = v21
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                var v22 string = ("hello" + "\n\t")
                                _ = v22
                            case "hello":
                                var v23 float64 = -2.25
                                _ = v23
                            default:
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        fmt.Println(fi24, fi24)
                                        v1 += (v1 >> ((2) & 63))
                                        _ = fi25
                                    }
                                    var v26 uint8 = (^(v2))
                                    _ = fi24
                                    _ = v26
                                }
                            }
                        }
                    case "":
                        v27 := GenS3{F0: 42, F1: 990}
                        _ = v27
                    }
                } else {
                    switch 0 {
                    case 4294967295:
                        v28 := GenS4{F0: 1, F1: (v1 + v1)}
                        _ = v28
                    default:
                        v1 = v1
                    }
                    fmt.Println(v2, v3)
                    if ((v3 / (((v2 >> ((18446744073709551615) & 63))) | 1)) < 1) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v29 float32 = 3.14159
                        _ = v29
                    }
                }
            }
            v1 = 1
        }
        if (787 < 128) {
            v1 = (v1 ^ (v1 << ((301) & 63)))
        }
    }
    var v30 int8 = (v1 >> ((58) & 63))
    _ = len(v6)
    v1 = 100
    for fi31 := 0; fi31 < 17; fi31++ {
        v2 += (v2 * (v3 % (((v3 | v2)) | 1)))
        v1 = (v1 - (v30 >> ((255) & 63)))
        _ = fi31
    }
    v1 = (v1 * v30)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v3 = (v3 - v3)
    if ((165 <= 2147483648) && false) {
        for fi32 := 0; fi32 < 8; fi32++ {
            if (!true) {
                for fi33 := 0; fi33 < 3; fi33++ {
                    if false {
                        v34 := GenS5{F0: 42, F1: v1, F2: "a"}
                        var v35 float32 = (float32((3.14159 * 0.0)) - float32(3.14159))
                        _ = v34
                        _ = v35
                    } else {
                        v36 := GenS5{F0: v3, F1: (v1 << ((18446744073709551615) & 63)), F2: "\n\t"}
                        var v37 bool = (!(!((("" + "") < "\n\t") || false)))
                        _ = v36
                        _ = v37
                    }
                    _ = fi33
                }
            }
            _ = fi32
        }
        fmt.Println(v2)
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS3{F0: 4294967295, F1: 434}
        var v40 string = "\n\t"
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS2{F0: 42}
    var v42 uint16 = 1
    var v43 uint8 = 9
    switch -322 {
    case 2147483647:
        v2 += (v3 & (v3 ^ v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v41
    _ = v42
    _ = v43
    _ = v6
}

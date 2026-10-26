package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint16; F2 int64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint16 = 635
            v6 = 0
            _ = fi5
            _ = v6
        }
        var v7 int8 = 1
        _ = fi4
        _ = v7
    }
    var v8 uint = 2
    v9 := GenS1{F0: 18446744073709551615, F1: 0}
    v8 = 65535
    var v10 uint8 = 0
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int64 = 1
    v2 := []int{}
    var v3 int64 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("hello" + ("\n\t" + "\n\t")) {
    case "a":
        v11 := GenS2{F0: v1, F1: 523, F2: (v3 & (v3 ^ v1))}
        _ = v11
    }
    if ((3.14159 < (0.1 * 1.5)) && true) {
        var v12 int8 = 89
        var v13 uint64 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
        _ = v13
    } else {
        v2 = append(v2, -527)
        for fi14 := 0; fi14 < 14; fi14++ {
            if false {
                fmt.Println(fi14)
                v2 = append(v2, (fi14 * fi14))
            }
            v3 += -395
            _ = fi14
        }
        _ = len(v2)
    }
    v3 -= -9223372036854775808
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 -= 1
                if true {
                    if len(v2) > 0 {
                        v2[0] = (-(fi15))
                    }
                    var v17 uint32 = 0
                    var v18 uint = 42
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
            if true {
                fmt.Println(v1)
                var v19 uint8 = 127
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 string = (("foo bar" + "hello") + "\n\t")
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                var v22 int = fi15
                switch "foo bar" {
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    if ((!false) || (924 != 2)) {
                        if len(v2) > 0 {
                            v2[0] = 56
                        }
                        v1 -= -1
                    } else {
                        v1 -= v1
                    }
                case "":
                    for fi23 := 0; fi23 < 13; fi23++ {
                        v24 := GenS2{F0: v1, F1: 42, F2: v3}
                        _ = fi23
                        _ = v24
                    }
                }
                if true {
                    _ = len(v2)
                    v3 -= 0
                    var v25 int16 = 1
                    _ = v25
                } else {
                    fmt.Println(v3)
                    switch "" {
                    case "foo bar":
                        v26 := GenS3{F0: 9223372036854775807, F1: 255}
                        _ = v26
                    default:
                        v1 = (v3 + (v3 >> ((436) & 63)))
                    }
                    v3 = 42
                }
                _ = v22
            }
            v3 -= v1
        }
        _ = len(v2)
        _ = fi15
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    v27 := GenS1{F0: 613, F1: 0}
    v28 := GenS1{F0: 0, F1: -1}
    v1 = (v1 + v1)
    v1 -= v3
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint32 = 2147483647
        v3 = 42
        _ = fi29
        _ = v30
    }
    var v31 bool = (42 <= 0)
    if false {
        var v32 uint32 = 448
        var v33 int8 = 127
        v1 -= v3
        _ = v32
        _ = v33
    } else {
        var v34 int32 = -1
        if v31 {
            v34 = v34
        } else {
            if (127 == 197) {
                for fi35 := 0; fi35 < 6; fi35++ {
                    v36 := GenS1{F0: 1, F1: 600}
                    if len(v2) > 0 {
                        v2[0] = (fi35 << ((42) & 63))
                    }
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v31)
            }
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            switch ("" + "") {
            case "foo bar":
                _ = len(v2)
            case "a":
                switch fi37 {
                case 1:
                    switch (fi37 << ((18446744073709551615) & 63)) {
                    case 405:
                        v34 += v34
                    case 4294967295:
                        var v38 uint16 = 0
                        _ = v38
                    case 127:
                        var v39 int = 128
                        _ = v39
                    }
                default:
                    v31 = v31
                }
            case "":
                if len(v2) > 0 {
                    v2[0] = (-(fi37))
                }
            }
            if len(v2) > 0 {
                v2[0] = 385
            }
            _ = fi37
        }
        _ = v34
    }
    v3 += (v3 << ((584) & 63))
    fmt.Println(v31)
    v40 := GenS4{F0: 651, F1: 330, F2: genFn0()}
    if (v31 && v31) {
        fmt.Println(v31, v31)
        v31 = (975 == 2)
    } else {
        v41 := GenS5{F0: (!v31), F1: 238, F2: -128}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v40
}

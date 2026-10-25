package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int64; F2 int }
type GenS5 struct { F0 uint8; F1 int8; F2 int16 }
type GenS6 struct { F0 int32; F1 int64; F2 int }

func main() {
    var v1 uint64 = 2
    var v2 int = -9223372036854775808
    var v3 string = ((("foo bar" + "foo bar") + "\n\t") + "\n\t")
    if (!((v1 != 935) && false)) {
        var v4 int = (v2 % (((^(v2))) | 1))
        v2 = (v4 / ((v4) | 1))
        _ = v4
    } else {
        v5 := GenS1{F0: (v2 ^ -9223372036854775808)}
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v3)
            var v7 uint8 = 127
            switch (fi6 << ((255) & 63)) {
            case 2147483647:
                v8 := GenS1{F0: (v2 >> ((901) & 63))}
                _ = v8
            case -31:
                for fi9 := 0; fi9 < 4; fi9++ {
                    fmt.Println(v2)
                    _ = fi9
                }
            default:
                v3 = ("\n\t" + v3)
            }
            _ = fi6
            _ = v7
        }
        v3 = v3
        _ = v5
    }
    switch (v3 + "") {
    case "a":
        v10 := []int{v2, (^(v2)), (v2 % ((v2) | 1))}
        _ = v10
    default:
        var v11 uint = 2
        _ = v11
    }
    v12 := GenS2{F0: (577 != 42), F1: 255}
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: (42 >= 1), F1: 2}
        for fi15 := 0; fi15 < 1; fi15++ {
            switch ("a" + v3) {
            case "foo bar":
                v1 += (v1 % (((v1 & (^(v1)))) | 1))
            case "\n\t":
                var v16 float64 = 0.1
                _ = v16
            case "hello":
                v3 = ("" + v3)
            }
            v17 := []int{fi15}
            _ = fi15
            _ = v17
        }
        if false {
            v18 := GenS3{F0: v3}
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(fi19)
                switch -1 {
                case 234:
                    fmt.Println(fi13, v3)
                case -162:
                    var v20 float64 = 1e10
                    _ = v20
                default:
                    switch (fi19 - (fi13 / (((v2 & v2)) | 1))) {
                    case -331:
                        if (false && true) {
                            v21 := GenS4{F0: 100, F1: 1, F2: v2}
                            v22 := []int{-1, fi19}
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v3)
                            v23 := GenS5{F0: 128, F1: 100, F2: -590}
                            _ = v23
                        }
                    case 32767:
                        v3 = "a"
                    default:
                        fmt.Println(v2, fi13)
                    }
                }
                _ = fi19
            }
            v24 := []int{fi13, (^(fi13)), (v2 % ((1) | 1)), 9223372036854775807}
            _ = v18
            _ = v24
        } else {
            v25 := []int{(^(fi13))}
            if false {
                v26 := GenS6{F0: 42, F1: -263, F2: 9223372036854775807}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 3; fi27++ {
                    v1 = (v1 >> ((1) & 63))
                    fmt.Println(v3, v2)
                    if (32767 < 1) {
                        v2 -= (fi27 / (((v2 + 4294967295)) | 1))
                        var v28 uint64 = (^(v1))
                        var v29 uint = 1
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint32 = 98
                        _ = v30
                    }
                    _ = len(v25)
                    _ = fi27
                }
                v31 := GenS6{F0: 0, F1: 0, F2: 282}
                _ = len(v25)
                _ = v31
            }
            _ = v25
        }
        _ = fi13
        _ = v14
    }
    v2 = (v2 + v2)
    var v32 int64 = 42
    var v33 uint8 = 15
    var v34 int = (v2 ^ -9223372036854775808)
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS5{F0: (v33 << ((826) & 63)), F1: -1, F2: 42}
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            v2 += (fi38 | (-(fi38)))
            if true {
                v3 = (v3 + v3)
                v33 -= 84
            } else {
                v2 = 275
            }
            v39 := GenS1{F0: -1}
            _ = fi38
            _ = v39
        }
        v40 := []int{(v2 << ((2) & 63)), fi35, v2, (fi35 & 0)}
        _ = fi35
        _ = v40
    }
    if true {
        v41 := GenS3{F0: (v3 + "")}
        var v42 float64 = (0.1 - 0.0)
        _ = v41
        _ = v42
    } else {
        v34 = v34
        var v43 int64 = v32
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 float64 = 1.5
        fmt.Println(v1, v2)
        _ = fi44
        _ = v45
    }
    var v46 int8 = 5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v46
}

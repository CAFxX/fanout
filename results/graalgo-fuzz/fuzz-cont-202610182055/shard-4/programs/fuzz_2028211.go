package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 string; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }

func main() {
    var v1 int = 0
    var v2 uint16 = 636
    var v3 string = "hello"
    var v4 int8 = 42
    v5 := GenS1{F0: (v4 | -1), F1: (v2 & (v2 / ((110) | 1))), F2: 117}
    var v6 int8 = (v4 & v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float64 = -2.25
        v1 += (fi7 / (((v1 - -2147483648)) | 1))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch ("\n\t" + v3) {
        case "a":
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(v1, v1)
                var v11 uint8 = 128
                switch 497 {
                case 853:
                    var v12 uint = 606
                    _ = v12
                case 128:
                    v13 := []int{(fi10 ^ (fi9 % ((fi10) | 1))), (^(fi9))}
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case "":
            if true {
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(fi9, v3)
                    var v15 uint64 = 0
                    var v16 int64 = 939
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                switch 994 {
                case 128:
                    if false {
                        switch -665 {
                        case 2147483647:
                            v17 := GenS2{F0: (^(v6))}
                            _ = v17
                        case 42:
                            v18 := []int{v1, (-(v1))}
                            _ = v18
                        case 0:
                            v2 = 2
                        default:
                            v19 := GenS3{F0: v1, F1: ("" + v3), F2: -2147483648}
                            _ = v19
                        }
                        fmt.Println(v3, v6, v4)
                        v20 := GenS3{F0: 1, F1: (v3 + "hello"), F2: -1}
                        _ = v20
                    }
                case -9223372036854775808:
                    var v21 uint = 1
                    _ = v21
                }
            } else {
                if true {
                    v22 := GenS3{F0: 2147483647, F1: (("foo bar" + ("hello" + v3)) + ("" + (v3 + v3))), F2: -1}
                    v2 -= v2
                    _ = v22
                }
                v6 += v6
                if false {
                    var v23 string = "hello"
                    _ = v23
                }
            }
        default:
            fmt.Println(v1, v2)
        }
        fmt.Println(v4)
        var v24 uint8 = 128
        _ = fi9
        _ = v24
    }
    fmt.Println(v3, v6)
    v25 := []int{v1, v1, (v1 & (v1 ^ 32768))}
    var v26 int32 = -856
    var v27 int16 = 279
    if true {
        _ = len(v25)
        v28 := GenS4{F0: 251, F1: (v4 >> ((0) & 63)), F2: int64(v6)}
        _ = v28
    } else {
        var v29 uint16 = v2
        switch (v1 % ((v1) | 1)) {
        case 9223372036854775807:
            v2 = v29
        case 52:
            if false {
                v30 := GenS5{F0: v29}
                if (true && true) {
                    switch 4294967295 {
                    case 127:
                        if len(v25) > 0 {
                            v25[0] = -686
                        }
                    case 2147483647:
                        v6 += v4
                    case -1:
                        if len(v25) > 0 {
                            v25[0] = -9223372036854775808
                        }
                    }
                    var v31 uint8 = 33
                    _ = v31
                } else {
                    if len(v25) > 0 {
                        v25[0] = (-(v1))
                    }
                    v32 := GenS6{F0: v3}
                    _ = v32
                }
                _ = v30
            }
        }
        v3 = v3
        _ = v29
    }
    if len(v25) > 0 {
        v25[0] = -810
    }
    var v33 int16 = v27
    v3 = ""
    if len(v25) > 0 {
        v25[0] = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

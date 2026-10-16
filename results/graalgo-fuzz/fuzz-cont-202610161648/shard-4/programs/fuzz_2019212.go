package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    var v1 int = 356
    var v2 int8 = 105
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := []string{"foo bar", ("\n\t" + ""), ("foo bar" + ("foo bar" + ("\n\t" + ""))), ("foo bar" + "a")}
    var v6 uint32 = 469
    fmt.Println(v6, v3, v6)
    if true {
        for fi7 := 0; fi7 < 8; fi7++ {
            if false {
                var v8 bool = (1 < 1)
                var v9 uint16 = 304
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(v6, v6)
                    _ = fi10
                }
                v3 = (v3 & 0)
                for fi11 := 0; fi11 < 11; fi11++ {
                    fmt.Println(v2, fi7)
                    _ = fi11
                }
            }
            v12 := GenS1{F0: "", F1: 2}
            _ = fi7
            _ = v12
        }
        fmt.Println(v2)
    }
    if (292 != (v6 - v6)) {
        switch (-(v1)) {
        case -462:
            v2 -= (v2 * v2)
        case -2147483648:
            v1 = v1
        case 128:
            v2 += v2
        }
        fmt.Println(v6, v6, v1)
        for fi13 := 0; fi13 < 17; fi13++ {
            v1 -= 2147483647
            if ((false || (42 < 65535)) && ((42 <= 101) && (v6 != (v6 << ((18446744073709551615) & 63))))) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 int16 = -32768
                    _ = fi14
                    _ = v15
                }
                v2 = (-(v2))
                switch "hello" {
                case "":
                    var v16 uint8 = 2
                    _ = v16
                default:
                    var v17 uint64 = 4294967295
                    _ = v17
                }
            } else {
                v4["hello"] = "hello"
                if true {
                    var v18 uint16 = 0
                    _ = v18
                } else {
                    if false {
                        fmt.Println(v2, v2, fi13)
                    }
                }
                if ((v6 & v6) < 0) {
                    v19 := GenS2{F0: (("foo bar" + ("\n\t" + "foo bar")) + "\n\t"), F1: (!true), F2: 2147483648}
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if (-9223372036854775808 == 0) {
                            v2 = int8(v3)
                            var v21 float64 = 1.5
                            _ = v21
                        } else {
                            fmt.Println(v3)
                        }
                        var v22 bool = (false && ((v2 ^ v2) > (v2 << ((4294967295) & 63))))
                        v23 := GenS1{F0: "\n\t", F1: 457}
                        _ = fi20
                        _ = v22
                        _ = v23
                    }
                    _ = v19
                }
            }
            switch (fi13 + (fi13 | (fi13 % (((v1 | v1)) | 1)))) {
            case 390:
                fmt.Println(v1, fi13, v3)
            case 42:
                if false {
                    fmt.Println(v3, v6, v2)
                }
            case -985:
                switch ("hello" + "a") {
                case "a":
                    var v24 string = "foo bar"
                    _ = v24
                }
            default:
                _ = len(v4)
            }
            _ = fi13
        }
    } else {
        v25 := GenS2{F0: (("\n\t" + "") + "a"), F1: (false && (42 > (^(v6)))), F2: v6}
        if true {
            v2 = 1
            var v26 uint16 = 2
            _ = v26
        }
        _ = v25
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v27 int32 = 1
    v6 = 65535
    var v28 float32 = -2.25
    if (((-2.25 * v28) - 1e10) <= v28) {
        v29 := v4[""]
        fmt.Println(v2, v1)
        for fi30 := 0; fi30 < 14; fi30++ {
            _ = len(v5)
            v3 = v3
            _ = fi30
        }
        _ = v29
    } else {
        v31 := GenS3{F0: (v1 / (((v1 >> ((1) & 63))) | 1)), F1: 42, F2: true}
        for fi32 := 0; fi32 < 13; fi32++ {
            v6 -= v6
            _ = fi32
        }
        _ = v31
    }
    var v33 bool = true
    switch "\n\t" {
    case "foo bar":
        v34 := GenS2{F0: "hello", F1: (v6 <= 65535), F2: (v6 + (v6 ^ 0))}
        _ = v34
    case "a":
        v35 := GenS4{F0: 1, F1: true, F2: v6}
        _ = v35
    case "\n\t":
        if true {
            _ = len(v5)
        }
    }
    if ((v27 == v27) && v33) {
        for fi36 := 0; fi36 < 1; fi36++ {
            for fi37 := 0; fi37 < 7; fi37++ {
                var v38 int8 = (v2 % ((v2) | 1))
                fmt.Println(v6, v3, v2)
                v6 += 0
                _ = fi37
                _ = v38
            }
            v39 := GenS2{F0: ("foo bar" + "hello"), F1: v33, F2: 0}
            v6 = 610
            _ = fi36
            _ = v39
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 float64 = ((((1.5 - (float64(v28) + (((0.0 + 3.14159) / ((3.14159 * -2.25) * 3.14159)) * 1.5))) / 1e10) - 1e10) * 1.5)
            var v42 uint32 = v6
            v43 := GenS2{F0: "", F1: (true && true), F2: v6}
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        if v33 {
            for fi45 := 0; fi45 < 1; fi45++ {
                v46 := v4["foo bar"]
                _ = fi45
                _ = v46
            }
            _ = len(v5)
            _ = len(v5)
        } else {
            var v47 float32 = v28
            _ = v47
        }
        _ = fi44
    }
    v6 -= 351
    delete(v4, "\n\t")
    for fi48 := 0; fi48 < 20; fi48++ {
        fmt.Println(v33, v27)
        for fi49 := 0; fi49 < 9; fi49++ {
            v5 = append(v5, "foo bar")
            _ = fi49
        }
        _ = fi48
    }
    var v50 float64 = (1e10 - 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}

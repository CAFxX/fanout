package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 string; F2 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 float32 = 0.0
    var v2 uint = 42
    v3 := []uint32{}
    switch -512 {
    case -9223372036854775808:
        var v4 string = "a"
        _ = v4
    case 32768:
        var v5 bool = (-1 >= 0)
        _ = v5
    case -306:
        if false {
            if (!true) {
                var v6 float32 = (1.5 + float32(-2.25))
                fmt.Println(v1)
                v1 = ((((v1 - v6) / (1e10 - 0.0)) / v1) * 1.5)
                _ = v6
            }
        } else {
            v3 = append(v3, 0)
            _ = len(v3)
        }
    default:
        for fi7 := 0; fi7 < 19; fi7++ {
            v1 = (0.0 / float32(1e10))
            _ = fi7
        }
    }
    var v8 uint64 = 506
    v2 = 816
    fmt.Println(v8)
    var v9 bool = true
    var v10 int8 = -69
    v1 = v1
    v1 -= v1
    fmt.Println(v9)
    switch "foo bar" {
    case "foo bar":
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int = fi11
            if v9 {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v9 = v9
                    switch fi11 {
                    case -650:
                        switch fi11 {
                        case -1:
                            var v14 uint = v2
                            _ = v14
                        case 227:
                            v15 := GenS1{F0: v8, F1: v2, F2: 2147483647}
                            _ = v15
                        default:
                            if len(v3) > 0 {
                                v3[0] = 913
                            }
                        }
                    case 0:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 2147483647:
                        v1 -= v1
                    }
                    v16 := GenS2{F0: (65535 < 4294967295), F1: ("\n\t" + ("hello" + "foo bar")), F2: -1}
                    _ = fi13
                    _ = v16
                }
                v17 := GenS1{F0: v8, F1: v2, F2: -2147483648}
                _ = v17
            } else {
                v2 -= v2
                var v18 uint64 = v8
                if true {
                    var v19 uint8 = 246
                    _ = v19
                }
                _ = v18
            }
            _ = fi11
            _ = v12
        }
    case "hello":
        for fi20 := 0; fi20 < 12; fi20++ {
            v9 = ((v1 * (v1 / 0.1)) >= (v1 * v1))
            _ = fi20
        }
    }
    v21 := GenS1{F0: 18446744073709551615, F1: v2, F2: 1}
    v8 -= v8
    var v22 int64 = 164
    switch 633 {
    case 4294967295:
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS2{F0: false, F1: "\n\t", F2: v22}
            for fi25 := 0; fi25 < 8; fi25++ {
                if len(v3) > 0 {
                    v3[0] = 939
                }
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    case 2147483647:
        for fi26 := 0; fi26 < 10; fi26++ {
            switch (("" + ("\n\t" + "foo bar")) + ("a" + ("hello" + "\n\t"))) {
            case "\n\t":
                v27 := GenS3{F0: (v22 / ((v22) | 1)), F1: v22}
                _ = v27
            case "":
                var v28 int16 = 0
                _ = v28
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                fmt.Println(v1, v22)
            }
            _ = fi26
        }
    default:
        if v9 {
            v8 -= v8
            v29 := GenS4{F0: 4294967295}
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}

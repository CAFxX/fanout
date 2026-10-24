package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint, p1 uint8, p2 uint8) bool {
    var v4 float32 = 0.0
    v4 -= (v4 - 1.5)
    switch 0 {
    case 32767:
        v5 := GenS1{F0: 32767, F1: -1, F2: -41}
        _ = v5
    }
    fmt.Println(v4, p0)
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}
func genFn1() string {
    v6 := GenS1{F0: 0, F1: 2147483647, F2: 0}
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int16 = 1
            _ = fi8
            _ = v9
        }
        var v10 int64 = 42
        _ = fi7
        _ = v10
    }
    var v11 uint32 = 4294967295
    _ = v11
    _ = v6
    return "a"
}

func main() {
    v1 := []uint32{1, 42, 529}
    var v2 int16 = 969
    var v3 uint16 = 2
    v3 += 32767
    if (!true) {
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    var v13 uint64 = 859
    if (9223372036854775807 == 128) {
        fmt.Println(v3, v2, v2)
        var v14 uint64 = 146
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 string = ((("hello" + ((("\n\t" + "foo bar") + "\n\t") + "")) + "") + ("a" + "\n\t"))
            var v17 bool = false
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    fmt.Println(v13)
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    switch "foo bar" {
    case "\n\t":
        if (false || true) {
            switch 4294967295 {
            case 127:
                v2 -= v2
            default:
                switch 127 {
                case -9223372036854775808:
                    if true {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v3 = v3
                            switch "a" {
                            case "":
                                if true {
                                    var v19 uint32 = 2147483647
                                    fmt.Println(v13)
                                    _ = v19
                                }
                            }
                            _ = fi18
                        }
                    }
                case 128:
                    fmt.Println(v3, v3, v2)
                }
            }
        }
    case "foo bar":
        v20 := GenS1{F0: 0, F1: 2147483647, F2: 2147483647}
        _ = v20
    case "hello":
        v21 := GenS2{F0: 100, F1: 4294967295}
        _ = v21
    }
    var v22 int8 = 100
    if (!(true || (0 >= -456))) {
        var v23 uint8 = 1
        v1 = append(v1, 2147483648)
        _ = v23
    }
    switch 4294967295 {
    case 32768:
        if false {
            v2 -= (v2 & -1)
        } else {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = len(v1)
            switch "a" {
            case "\n\t":
                v13 = (v13 & (v13 & v13))
            case "":
                if true {
                    var v24 uint32 = 0
                    v2 += 32767
                    if (true && ((-(v22)) < v22)) {
                        v13 = (v13 % ((v13) | 1))
                        v25 := GenS1{F0: v3, F1: -2147483648, F2: -1}
                        if ((v2 >> ((18446744073709551615) & 63)) == (^(v2))) {
                            if true {
                                v22 = (-(v22))
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    v1 = append(v1, (^(v24)))
                                    _ = fi26
                                }
                            }
                            var v27 int8 = 88
                            var v28 int64 = 42
                            _ = v27
                            _ = v28
                        } else {
                            switch ((("" + (("hello" + ("foo bar" + ("hello" + ("\n\t" + "a")))) + "foo bar")) + "a") + ("a" + "a")) {
                            case "a":
                                var v29 uint32 = (v24 / ((2147483647) | 1))
                                _ = v29
                            default:
                                v24 = v24
                            }
                            v30 := GenS2{F0: (v22 & v22), F1: 334}
                            _ = v30
                        }
                        _ = v25
                    }
                    _ = v24
                } else {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        fmt.Println(v13, v13)
                        v22 -= -1
                        _ = len(v1)
                        _ = fi31
                    }
                }
            case "foo bar":
                v32 := GenS2{F0: (-(v22)), F1: 924}
                _ = v32
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
}

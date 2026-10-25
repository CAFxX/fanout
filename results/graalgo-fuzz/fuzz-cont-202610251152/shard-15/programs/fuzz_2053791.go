package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16, p1 string, p2 int64) string {
    p0 += p0
    v4 := GenS1{F0: 32767, F1: 32767, F2: 4294967295}
    v5 := GenS1{F0: 1, F1: 375, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p1 + (p1 + p1))
}

func main() {
    var v1 int64 = 200
    var v2 float64 = (1.5 - -2.25)
    v3 := make(map[uint32]uint32)
    v6 := []int{944, 980}
    v7 := v3[4294967295]
    v1 += (v1 / ((v1) | 1))
    for fi8 := 0; fi8 < 11; fi8++ {
        switch 903 {
        case 0:
            if (475 > 0) {
                fmt.Println(fi8)
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case -763:
            v9 := GenS1{F0: fi8, F1: (fi8 / (((fi8 % (((fi8 | -769)) | 1))) | 1)), F2: 2}
            _ = v9
        case -9223372036854775808:
            switch 32768 {
            case 1:
                var v10 bool = (true || (2147483647 != 138))
                _ = v10
            }
        default:
            var v11 bool = (113 > 127)
            _ = v11
        }
        if false {
            fmt.Println(v2, v2)
            v12 := GenS1{F0: fi8, F1: 871, F2: 985}
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                var v13 int8 = 127
                _ = v13
            case "":
                for fi14 := 0; fi14 < 10; fi14++ {
                    v2 += v2
                    _ = fi14
                }
            default:
                for fi15 := 0; fi15 < 18; fi15++ {
                    switch "hello" {
                    case "a":
                        v7 = 42
                    case "foo bar":
                        v7 += v7
                        v16 := GenS1{F0: (fi15 & -67), F1: fi15, F2: 901}
                        _ = v16
                    case "\n\t":
                        v7 += v7
                    default:
                        if (-395 != 183) {
                            switch (("a" + "\n\t") + "") {
                            case "hello":
                                var v17 int8 = 127
                                _ = v17
                            case "foo bar":
                                v2 = 0.1
                            default:
                                var v18 int8 = 0
                                _ = v18
                            }
                            v19 := GenS1{F0: (-(fi15)), F1: (fi8 % ((740) | 1)), F2: 1}
                            var v20 float64 = 0.0
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(fi15)
                        }
                    }
                    _ = fi15
                }
            }
            _ = v12
        }
        _ = fi8
    }
    if (2 == 65535) {
        if len(v6) > 0 {
            v6[0] = -395
        }
        if true {
            var v21 float64 = 0.0
            if true {
                var v22 int = int(v1)
                switch "\n\t" {
                case "hello":
                    v1 = -670
                case "":
                    v22 = v22
                case "\n\t":
                    v2 += 1.5
                }
                _ = v22
            }
            _ = v21
        } else {
            var v23 uint32 = (^(v7))
            if len(v6) > 0 {
                v6[0] = 188
            }
            _ = v23
        }
        var v24 float64 = float64(1e10)
        _ = v24
    } else {
        v25 := GenS2{F0: 18446744073709551615}
        _ = v25
    }
    if (-1 > 32767) {
        v2 += v2
        v7 -= v7
    }
    fmt.Println(v1, v2)
    v2 = float64((-2.25 * 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}

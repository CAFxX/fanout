package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -551
    v3 := []int{}
    var v4 int16 = (v2 >> ((255) & 63))
    fmt.Println(v2, v2)
    v4 += (-(v2))
    var v5 int = -810
    v2 += v4
    v4 = -347
    v2 = 32767
    _ = len(v1)
    var v6 int16 = v2
    v5 = (v5 % ((v5) | 1))
    v6 -= (v2 | v2)
    var v7 int16 = v2
    switch v5 {
    case -959:
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                if (0 >= 42) {
                    if len(v3) > 0 {
                        v3[0] = (fi8 << ((65535) & 63))
                    }
                    v3 = append(v3, 32767)
                    if (100 != 114) {
                        v10 := GenS1{F0: "foo bar"}
                        v7 += v4
                        _ = v10
                    } else {
                        v4 = (^(v4))
                        var v11 string = "\n\t"
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v11
                    }
                    if (4294967295 == 1) {
                        v12 := GenS1{F0: "\n\t"}
                        _ = v12
                    }
                }
                v13 := v1[("foo bar" + (("\n\t" + "a") + ""))]
                _ = fi9
                _ = v13
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: "a"}
                _ = fi14
                _ = v15
            }
            _ = fi8
        }
    default:
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: ""}
            _ = fi16
            _ = v17
        }
    }
    v1["foo bar"] = 42
    switch "\n\t" {
    case "\n\t":
        if true {
            if false {
                v2 = v6
                v5 -= 546
                fmt.Println(v7)
            }
        }
    default:
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                v19 := GenS2{F0: (fi18 / ((-354) | 1)), F1: 42}
                if (773 != 1) {
                    if len(v3) > 0 {
                        v3[0] = v5
                    }
                    var v20 int64 = 0
                    _ = v20
                }
                var v21 int64 = 1
                _ = v19
                _ = v21
            }
            _ = fi18
        }
    }
    fmt.Println(v6, v6)
    if (!(!false)) {
        var v22 int32 = -2147483648
        v5 -= v5
        _ = v22
    } else {
        fmt.Println(v2, v4)
    }
    if ((-9223372036854775808 != 9223372036854775807) || (((v5 - v5) != v5) || (true && ((127 != -128) || (false && false))))) {
        v7 -= v2
    } else {
        if ((false || (!true)) && (("foo bar" + "") >= ("hello" + "a"))) {
            v23 := GenS3{F0: 1}
            _ = v23
        }
        if (!(65535 == 18446744073709551615)) {
            v6 += (v6 % (((^(v6))) | 1))
        }
        fmt.Println(v4, v6, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

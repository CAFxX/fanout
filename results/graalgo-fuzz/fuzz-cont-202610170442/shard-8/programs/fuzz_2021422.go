package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v5 uint16 = 636
    var v6 uint8 = 27
    fmt.Println(v6, v5)
    v7 := GenS1{F0: 2, F1: true}
    fmt.Println(v5, v6)
    v5 = (v5 << ((65535) & 63))
    switch "\n\t" {
    case "a":
        v5 = (v5 ^ (v5 / (((v5 ^ v5)) | 1)))
    }
    if len(v2) > 0 {
        v2[0] = -774
    }
    v8 := v3[("" + "hello")]
    v9 := GenS1{F0: 624, F1: ((0 != 0) || (32768 >= 32767))}
    v5 = 0
    if v8 {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 += 22
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint32 = 0
        if (!(710 <= 412)) {
            if v8 {
                v12 := GenS1{F0: 42, F1: false}
                v4["foo bar"] = (v11 - (v11 / (((v11 ^ (v11 & (v11 - (^(v11)))))) | 1)))
                switch "\n\t" {
                case "\n\t":
                    if (-930 != -72) {
                        v13 := GenS2{F0: 42, F1: ("" + "a")}
                        _ = v13
                    }
                case "hello":
                    v11 += (v11 % (((v11 >> ((18446744073709551615) & 63))) | 1))
                case "foo bar":
                    v14 := GenS1{F0: 0, F1: ((false && (v8 && (4294967295 <= 1))) || v8)}
                    _ = v14
                default:
                    if (((^(v6)) <= v6) || (v8 || ((406 >= 378) || (v6 < v6)))) {
                        fmt.Println(v11)
                    } else {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v8 = ((false && ((^(fi10)) >= (fi15 ^ 32768))) && (fi15 >= fi15))
                            v5 = v5
                            _ = fi15
                        }
                    }
                }
                _ = v12
            }
            if v8 {
                fmt.Println(v8, v8)
                v6 = v6
                for fi16 := 0; fi16 < 20; fi16++ {
                    _ = len(v2)
                    _ = fi16
                }
            } else {
                var v17 string = "foo bar"
                _ = v17
            }
        }
        _ = fi10
        _ = v11
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS2{F0: (v6 - v6), F1: "foo bar"}
        switch -478 {
        case 4294967295:
            for fi20 := 0; fi20 < 7; fi20++ {
                v8 = (v8 || true)
                v6 = (v6 & 206)
                switch (-(fi20)) {
                case 42:
                    switch int(v5) {
                    case 744:
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    case 1:
                        v6 = (^(v6))
                    case 67:
                        fmt.Println(v5)
                    }
                default:
                    fmt.Println(fi20, v5, fi18)
                }
                _ = fi20
            }
        case 642:
            if v8 {
                v4[("" + ("foo bar" + ""))] = 42
                var v22 int64 = 893
                _ = v22
            } else {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int16 = 1
                    var v25 bool = false
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v8 = v8
            }
        }
        if len(v2) > 0 {
            v2[0] = fi18
        }
        _ = fi18
        _ = v19
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

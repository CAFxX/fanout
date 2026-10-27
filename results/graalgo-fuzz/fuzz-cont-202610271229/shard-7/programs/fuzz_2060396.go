package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]string)
    v3 := []string{"foo bar", ("a" + "\n\t"), "\n\t"}
    fmt.Println(v1)
    if true {
        for fi4 := 0; fi4 < 12; fi4++ {
            v1 = (v1 ^ 0)
            fmt.Println(fi4)
            _ = fi4
        }
    } else {
        _ = len(v3)
        if false {
            switch "foo bar" {
            case "\n\t":
                _ = len(v3)
            case "hello":
                var v5 int64 = 0
                _ = v5
            }
            v2[237] = ("foo bar" + "foo bar")
            if len(v3) > 0 {
                v3[0] = ("" + (((("\n\t" + "a") + "a") + "") + "hello"))
            }
        } else {
            if (false || false) {
                var v6 int64 = -1
                v7 := GenS1{F0: v1, F1: 2147483648}
                v1 = (v1 % ((v1) | 1))
                _ = v6
                _ = v7
            } else {
                var v8 bool = (!(float32(3.14159) >= (1.5 - ((1.5 * 1e10) - 0.0))))
                _ = v8
            }
            if true {
                switch -574 {
                case 4294967295:
                    v3 = append(v3, "hello")
                case 127:
                    v1 += (v1 << ((42) & 63))
                case -531:
                    v1 += (v1 + 2)
                default:
                    v2[4294967295] = ("foo bar" + "a")
                }
                var v9 int32 = -2147483648
                _ = v9
            } else {
                v10 := v2[42]
                v10 = ("" + ((v10 + "foo bar") + v10))
                switch 127 {
                case 2147483647:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        switch (v10 + "hello") {
                        case "a":
                            v1 = 128
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "hello":
                            v10 = ((v10 + v10) + "foo bar")
                        }
                        v1 = v1
                        _ = fi11
                    }
                case 1:
                    v12 := GenS1{F0: (v1 | (v1 / (((v1 + v1)) | 1))), F1: 65535}
                    _ = v12
                case -1:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        if true {
                            v10 = (v10 + "\n\t")
                            fmt.Println(v1, v1, v10)
                        }
                        v10 = "hello"
                        if (42 <= 174) {
                            var v14 int16 = -1
                            _ = v14
                        }
                        _ = fi13
                    }
                default:
                    var v15 int64 = -956
                    _ = v15
                }
                _ = v10
            }
        }
        switch "\n\t" {
        case "\n\t":
            var v16 int32 = 1
            _ = v16
        case "foo bar":
            fmt.Println(v1)
        case "":
            for fi17 := 0; fi17 < 8; fi17++ {
                if (32767 <= 0) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int8 = 0
                        var v20 bool = (float64(((0.1 / 1.5) - 0.0)) >= -2.25)
                        _ = len(v3)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi17
            }
        }
    }
    v21 := GenS1{F0: (v1 + (v1 / ((v1) | 1))), F1: 2147483648}
    v1 = (v1 << ((332) & 63))
    delete(v2, 4294967295)
    for fi22 := 0; fi22 < 7; fi22++ {
        v1 = 128
        var v23 int16 = 1000
        if ("a" < "a") {
            if (100 > 39) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v24 int32 = 2147483647
                v25 := GenS1{F0: v1, F1: uint32(v1)}
                _ = v24
                _ = v25
            }
        }
        _ = fi22
        _ = v23
    }
    v26 := GenS1{F0: (v1 | v1), F1: 789}
    fmt.Println(v1)
    delete(v2, 97)
    v1 = (v1 | (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
}

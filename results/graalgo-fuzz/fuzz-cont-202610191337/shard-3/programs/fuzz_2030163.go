package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []int{-9223372036854775808, 9223372036854775807}
    var v2 int64 = 0
    v3 := []int{32767, 803, -9223372036854775808, 32767}
    var v4 uint16 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "a":
        for fi5 := 0; fi5 < 8; fi5++ {
            v4 -= (v4 | (v4 - 32767))
            v2 -= (v2 * v2)
            _ = fi5
        }
    case "hello":
        fmt.Println(v4, v4)
    case "\n\t":
        _ = len(v3)
    }
    var v6 string = ("\n\t" + "a")
    v2 = v2
    _ = len(v1)
    switch "" {
    case "foo bar":
        var v7 int8 = 127
        _ = v7
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v8 int32 = -276
        _ = v8
    }
    if ((v2 / (((v2 | -624)) | 1)) > (v2 >> ((557) & 63))) {
        if len(v3) > 0 {
            v3[0] = -1
        }
    }
    v2 = v2
    v9 := GenS1{F0: 127, F1: 0}
    for fi10 := 0; fi10 < 7; fi10++ {
        switch v6 {
        case "foo bar":
            if true {
                var v11 int8 = 0
                _ = v11
            }
        }
        var v12 uint16 = v4
        if (!true) {
            if ("" <= v6) {
                v6 = ((((v6 + "\n\t") + "hello") + "") + v6)
                var v13 uint64 = 18446744073709551615
                fmt.Println(v2)
                _ = v13
            } else {
                var v14 uint8 = 42
                var v15 int8 = -39
                v2 += 0
                _ = v14
                _ = v15
            }
            var v16 uint32 = 4294967295
            _ = v16
        } else {
            switch v6 {
            case "":
                v17 := GenS1{F0: 0, F1: 0}
                _ = v17
            case "foo bar":
                if false {
                    v18 := GenS1{F0: 215, F1: 127}
                    _ = v18
                } else {
                    switch (fi10 << ((787) & 63)) {
                    case 9223372036854775807:
                        v4 -= v4
                    case 2147483647:
                        if (true || (128 < 42)) {
                            var v19 int64 = v2
                            for fi20 := 0; fi20 < 9; fi20++ {
                                if false {
                                    var v21 uint64 = 4294967295
                                    _ = v21
                                } else {
                                    var v22 int64 = v2
                                    var v23 uint8 = 42
                                    _ = v22
                                    _ = v23
                                }
                                switch "a" {
                                case "hello":
                                    if (195 == 74) {
                                        if (!false) {
                                            v24 := GenS1{F0: 128, F1: 2}
                                            fmt.Println(fi20)
                                            var v25 int32 = 42
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                case "foo bar":
                                    switch "hello" {
                                    case "":
                                        if false {
                                            switch fi10 {
                                            case 1:
                                                var v26 int16 = -32768
                                                _ = v26
                                            default:
                                                var v27 uint8 = 2
                                                _ = v27
                                            }
                                            if ((fi10 * -247) < fi20) {
                                                v12 = (v12 ^ (v4 ^ (v4 % (((^(v12))) | 1))))
                                                v19 = 146
                                                v2 = 0
                                            }
                                        } else {
                                            v28 := GenS1{F0: 127, F1: 128}
                                            _ = v28
                                        }
                                    case "hello":
                                        var v29 int8 = 74
                                        _ = v29
                                    }
                                default:
                                    var v30 uint32 = 65535
                                    _ = v30
                                }
                                _ = fi20
                            }
                            v1 = append(v1, fi10)
                            _ = v19
                        } else {
                            if len(v1) > 0 {
                                v1[0] = fi10
                            }
                            v4 = v4
                        }
                    default:
                        v6 = v6
                    }
                }
            case "hello":
                if true {
                    v2 = v2
                } else {
                    var v31 uint = 255
                    _ = v31
                }
            }
        }
        _ = fi10
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}

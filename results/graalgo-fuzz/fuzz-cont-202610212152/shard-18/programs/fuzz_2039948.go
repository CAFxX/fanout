package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{617, 42, 2147483648}
    var v2 int = -433
    v3 := []int{}
    var v4 float32 = 0.0
    v2 = 4294967295
    if false {
        if true {
            var v5 bool = true
            v5 = ("" != "\n\t")
            _ = v5
        } else {
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v6 uint32 = 254
                    v1 = append(v1, v6)
                    _ = v6
                }
            case "hello":
                if (false && (true || false)) {
                    var v7 int16 = 0
                    v2 -= v2
                    _ = v7
                }
            default:
                for fi8 := 0; fi8 < 10; fi8++ {
                    if ((2147483647 != -227) || ("foo bar" <= ("foo bar" + ""))) {
                        for fi9 := 0; fi9 < 2; fi9++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                if (!true) {
                                    fmt.Println(fi8)
                                    fmt.Println(v4, v2)
                                    var v10 int8 = -108
                                    _ = v10
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v4 -= ((v4 / v4) * 0.1)
                            _ = fi9
                        }
                    }
                    _ = fi8
                }
            }
            switch v2 {
            case 1:
                _ = len(v1)
            case 32768:
                switch (v2 >> ((454) & 63)) {
                case 42:
                    if true {
                        fmt.Println(v4, v4)
                    } else {
                        var v11 uint8 = 127
                        if (uint(v11) >= 42) {
                            var v12 uint64 = 1
                            _ = len(v3)
                            _ = v12
                        }
                        v13 := GenS1{F0: 1, F1: v2, F2: v2}
                        _ = v11
                        _ = v13
                    }
                }
            case 640:
                fmt.Println(v2, v2)
            }
            switch (v2 >> ((2) & 63)) {
            case 42:
                v14 := GenS1{F0: -1, F1: 1, F2: 1}
                _ = v14
            case -574:
                _ = len(v3)
            case -614:
                var v15 int8 = 0
                _ = v15
            default:
                switch -797 {
                case -1:
                    v16 := GenS1{F0: -3, F1: 0, F2: (v2 % ((v2) | 1))}
                    _ = v16
                case 1:
                    if ((711 > 0) || (238 < 2)) {
                        if (false && ((1.5 / 1.0) == 3.14159)) {
                            if (4294967295 != 2) {
                                v4 = (v4 / (v4 / v4))
                                var v17 int32 = 2147483647
                                _ = v17
                            } else {
                                var v18 int16 = 337
                                fmt.Println(v18, v2, v4)
                                _ = v18
                            }
                            switch (("\n\t" + ("hello" + (("foo bar" + ("" + "\n\t")) + "\n\t"))) + (("foo bar" + "a") + (("a" + "\n\t") + ""))) {
                            case "":
                                v2 = v2
                            case "foo bar":
                                v19 := GenS1{F0: 47, F1: v2, F2: v2}
                                _ = v19
                            }
                            v20 := GenS2{F0: 255}
                            _ = v20
                        }
                    } else {
                        v2 = (^(v2))
                    }
                default:
                    var v21 int32 = 2147483647
                    _ = v21
                }
            }
        }
        fmt.Println(v2, v2)
        fmt.Println(v4, v4)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v22 uint8 = 128
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v22
    }
    var v23 int32 = 42
    var v24 float64 = 3.14159
    var v25 uint = 878
    if (v4 <= v4) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v25 = v25
        _ = len(v1)
    }
    v26 := GenS1{F0: 48, F1: (v2 & v2), F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}

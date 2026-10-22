package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }

func main() {
    v1 := []string{""}
    v2 := make(map[string]bool)
    v3 := []uint32{2, 2147483647, 1}
    v4 := make(map[uint32]int)
    switch 1 {
    case -161:
        var v5 float32 = (0.1 + 0.0)
        _ = v5
    case 128:
        var v6 int8 = 13
        _ = v6
    case 9223372036854775807:
        var v7 int8 = 127
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    var v9 uint8 = 71
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int8 = 103
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 int = 4294967295
    v8 -= (-(v8))
    var v14 uint8 = v9
    v9 = 1
    var v15 uint8 = v14
    var v16 uint16 = 136
    if false {
        for fi17 := 0; fi17 < 17; fi17++ {
            v13 = fi17
            v18 := GenS1{F0: v9, F1: (^(v16)), F2: 65535}
            _ = fi17
            _ = v18
        }
        var v19 uint16 = 32767
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS1{F0: 128, F1: (^(v19)), F2: 855}
            var v22 float64 = 1e10
            if false {
                if true {
                    var v23 int8 = -1
                    if true {
                        if (!(774 > 522)) {
                            v15 = v9
                            v22 = 0.1
                            delete(v2, "\n\t")
                        } else {
                            v22 = (-2.25 / 1e10)
                        }
                        _ = len(v4)
                        var v24 int32 = -1
                        _ = v24
                    } else {
                        switch "a" {
                        case "":
                            if ("a" <= "\n\t") {
                                var v25 uint64 = 1
                                if (v16 <= (^(v16))) {
                                    v23 += (v23 << ((1) & 63))
                                    var v26 bool = true
                                    _ = v26
                                }
                                var v27 string = ("a" + "foo bar")
                                _ = v25
                                _ = v27
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if (!true) {
                            switch (v13 + (v13 / (((fi20 * 128)) | 1))) {
                            case 127:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 9223372036854775807:
                                var v28 int64 = v8
                                _ = v28
                            default:
                                v29 := GenS1{F0: 2, F1: v19, F2: 42}
                                _ = v29
                            }
                            var v30 bool = false
                            var v31 bool = (264 != 722)
                            _ = v30
                            _ = v31
                        }
                        if ((0 <= -1) || (2147483647 != 42)) {
                            v32 := GenS1{F0: 42, F1: uint16(v9), F2: 1}
                            if ((v15 + (v9 >> ((2) & 63))) < 127) {
                                switch "" {
                                case "hello":
                                    v23 = (v23 * (v23 ^ 31))
                                case "":
                                    if (!((true && false) && false)) {
                                        if false {
                                            _ = len(v1)
                                        } else {
                                            for fi33 := 0; fi33 < 13; fi33++ {
                                                v34 := GenS1{F0: v15, F1: uint16(v15), F2: 875}
                                                _ = fi33
                                                _ = v34
                                            }
                                        }
                                    }
                                case "\n\t":
                                    var v35 uint8 = (v15 >> ((19) & 63))
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 19; fi36++ {
                                    var v37 float64 = (v22 - 1.5)
                                    var v38 uint16 = 0
                                    switch "" {
                                    case "":
                                        v13 -= 42
                                    default:
                                        fmt.Println(v16, fi36, v14)
                                    }
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                fmt.Println(v23)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v39 int16 = 42
                                _ = v39
                            }
                            _ = v32
                        }
                    }
                    _ = v23
                } else {
                    var v40 uint8 = v9
                    _ = v40
                }
                fmt.Println(v15, fi20, v8)
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

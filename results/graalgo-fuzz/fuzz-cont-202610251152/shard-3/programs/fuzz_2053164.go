package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int64 = 0
    var v2 bool = true
    var v3 uint = 18446744073709551615
    var v4 int8 = 127
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint64 = 4294967295
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 224
    v8 := []int{4294967295, -504, 42}
    _ = len(v8)
    for fi9 := 0; fi9 < 15; fi9++ {
        if (1 > 4294967295) {
            var v10 int64 = 594
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if len(v8) > 0 {
                v8[0] = (fi9 ^ (^(fi9)))
            }
            _ = v10
        } else {
            var v11 int8 = (v4 + (v4 << ((v3) & 63)))
            var v12 int32 = 2147483647
            v13 := GenS1{F0: 42}
            _ = v11
            _ = v12
            _ = v13
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
    }
    _ = len(v8)
    if len(v8) > 0 {
        v8[0] = 97
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!true) {
                switch ("\n\t" + "a") {
                case "a":
                    var v16 int = fi15
                    _ = v16
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        var v17 uint = (v3 / (((v3 ^ (v3 % (((v3 ^ v3)) | 1)))) | 1))
                        _ = v17
                    case "foo bar":
                        v1 = v1
                    default:
                        _ = len(v8)
                    }
                default:
                    var v18 float64 = ((float64(0.0) + 3.14159) + 3.14159)
                    _ = v18
                }
                if (!v2) {
                    v19 := GenS1{F0: (v4 - 1)}
                    if v2 {
                        v3 -= v3
                        v7 = v7
                    }
                    var v20 int32 = 348
                    _ = v19
                    _ = v20
                } else {
                    _ = len(v8)
                    if (!(0 != 0)) {
                        v7 = v7
                        fmt.Println(fi14, v4, v4)
                    } else {
                        var v21 int16 = -933
                        var v22 string = (((("\n\t" + "\n\t") + ("hello" + "a")) + "") + "")
                        var v23 int32 = 42
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    switch "a" {
                    case "hello":
                        var v24 uint16 = 2
                        _ = v24
                    default:
                        switch (fi14 / ((fi15) | 1)) {
                        case -1:
                            v3 = (v3 - 42)
                        case 32767:
                            switch "" {
                            case "a":
                                fmt.Println(v4)
                            case "foo bar":
                                v25 := GenS2{F0: v3}
                                _ = v25
                            case "":
                                switch fi15 {
                                case 42:
                                    v3 = v3
                                case -2147483648:
                                    v7 += v7
                                }
                            }
                        case -2147483648:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        default:
                            var v26 int8 = (v4 >> ((v3) & 63))
                            _ = v26
                        }
                    }
                }
                v1 = v1
            } else {
                var v27 uint = (v3 % (((^(v3))) | 1))
                v28 := GenS3{F0: true, F1: true}
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: 0}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi15
            _ = v29
        }
        _ = fi14
    }
    fmt.Println(v1)
    if (32767 < 325) {
        v1 -= -373
        switch 128 {
        case 127:
            for fi30 := 0; fi30 < 9; fi30++ {
                v31 := GenS4{F0: "", F1: 2147483648, F2: -228}
                for fi32 := 0; fi32 < 2; fi32++ {
                    v8 = append(v8, (fi30 << ((v3) & 63)))
                    if true {
                        if (-32768 < -32768) {
                            switch fi32 {
                            case 32768:
                                switch ("\n\t" + "hello") {
                                case "a":
                                    _ = len(v8)
                                }
                            case 42:
                                var v33 string = ("\n\t" + (("" + "\n\t") + "hello"))
                                _ = v33
                            default:
                                fmt.Println(fi32, v4)
                            }
                            var v34 uint = 255
                            _ = v34
                        }
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v4)
                        }
                    }
                    _ = fi32
                }
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 uint64 = 4294967295
                    _ = fi35
                    _ = v36
                }
                _ = fi30
                _ = v31
            }
        case 9223372036854775807:
            if ((^(v3)) == (^(v3))) {
                var v37 uint = v3
                _ = v37
            } else {
                v38 := GenS1{F0: v4}
                _ = v38
            }
        default:
            v39 := GenS5{F0: v7, F1: 975}
            _ = v39
        }
        var v40 int64 = v1
        _ = v40
    } else {
        switch "\n\t" {
        case "":
            var v41 float32 = (3.14159 / -2.25)
            _ = v41
        }
        v42 := GenS6{F0: (v4 / (((v4 & v4)) | 1)), F1: 0}
        _ = v42
    }
    if len(v8) > 0 {
        v8[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 string; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }
func genFn0(p0 uint64) uint16 {
    v5 := GenS1{F0: 308, F1: 207}
    fmt.Println(p0)
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int16 = 1000
            var v9 int = 128
            v10 := GenS2{F0: 2, F1: p0}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        p0 += 4294967295
        _ = fi6
    }
    var v11 uint8 = 19
    _ = p0
    _ = v11
    _ = v5
    return 42
}
func genFn1(p0 uint64) int8 {
    var v12 bool = ((!(128 > uint8(p0))) || false)
    v13 := GenS3{F0: "hello", F1: 2, F2: 42}
    _ = p0
    _ = v12
    _ = v13
    return -117
}

func main() {
    v1 := []int{-536, -214}
    var v2 int8 = 42
    v3 := []string{}
    var v4 bool = (979 >= 65535)
    var v14 int32 = -3
    switch "" {
    case "hello":
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int16 = 0
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(v2, v2)
                switch 13 {
                case -911:
                    var v18 int64 = -9223372036854775808
                    _ = v18
                case -839:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 515:
                    fmt.Println(v14)
                }
                fmt.Println(fi15, v16)
                var v19 int32 = 589
                var v20 int = (fi17 + fi17)
                _ = fi17
                _ = v19
                _ = v20
            }
            if v4 {
                if (0.0 <= 0.0) {
                    if v4 {
                        v1 = append(v1, int(v16))
                        v2 -= v2
                    }
                    var v21 uint64 = 0
                    _ = v21
                }
                if (58 != 32767) {
                    v22 := GenS1{F0: 42, F1: 2}
                    v23 := GenS4{F0: (v14 | (-(v14))), F1: (^(v2))}
                    v4 = (1 == 0)
                    _ = v22
                    _ = v23
                } else {
                    v2 = 42
                    if v4 {
                        var v24 uint64 = 42
                        _ = v24
                    }
                }
            } else {
                _ = len(v1)
            }
            _ = fi15
            _ = v16
        }
    default:
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS3{F0: "hello", F1: 1, F2: 2}
            _ = fi25
            _ = v26
        }
    }
    if (18446744073709551615 >= 113) {
        v4 = (true && v4)
        switch ("" + "") {
        case "":
            v14 += -1
        case "hello":
            var v27 uint16 = 2
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 10; fi28++ {
                switch (fi28 ^ 9223372036854775807) {
                case -9223372036854775808:
                    var v29 float32 = float32((float64(0.1) - 3.14159))
                    _ = v29
                case -1:
                    var v30 uint64 = 4294967295
                    _ = v30
                }
                _ = fi28
            }
        default:
            _ = len(v3)
        }
    }
    v2 = (v2 ^ 46)
    var v31 uint64 = 960
    v32 := GenS3{F0: "", F1: genFn0((v31 + (^(v31)))), F2: 127}
    for fi33 := 0; fi33 < 11; fi33++ {
        if true {
            var v34 float32 = 1e10
            fmt.Println(fi33, fi33)
            if v4 {
                v35 := GenS3{F0: ("foo bar" + "hello"), F1: 32767, F2: 163}
                for fi36 := 0; fi36 < 11; fi36++ {
                    switch "\n\t" {
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "foo bar":
                        v37 := GenS5{F0: -153, F1: (^(v14)), F2: -2147483648}
                        _ = v37
                    case "":
                        var v38 int16 = 42
                        _ = v38
                    default:
                        v39 := GenS1{F0: 9223372036854775807, F1: 369}
                        _ = v39
                    }
                    _ = fi36
                }
                if (!v4) {
                    for fi40 := 0; fi40 < 13; fi40++ {
                        var v41 uint8 = 212
                        if false {
                            var v42 uint = 255
                            _ = v42
                        }
                        v41 += 228
                        v4 = v4
                        _ = fi40
                        _ = v41
                    }
                    v43 := GenS4{F0: (v14 + (^(v14))), F1: genFn1(1)}
                    fmt.Println(v4)
                    _ = v43
                }
                _ = v35
            } else {
                v2 += 42
                v4 = false
                v44 := GenS4{F0: v14, F1: 100}
                _ = v44
            }
            _ = v34
        }
        _ = fi33
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        if true {
            _ = len(v3)
            var v46 uint8 = 143
            _ = v46
        }
        v2 -= (v2 - (v2 / (((v2 | genFn1(v31))) | 1)))
        _ = fi45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}

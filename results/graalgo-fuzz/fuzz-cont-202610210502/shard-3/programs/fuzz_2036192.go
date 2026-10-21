package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 int, p1 int32) int64 {
    switch p0 {
    case 127:
        var v7 int = (^(p0))
        _ = v7
    case 42:
        p1 -= (p1 % ((-2147483648) | 1))
    default:
        p1 = (p1 >> ((237) & 63))
    }
    var v8 uint64 = 853
    fmt.Println(p0, p0, p0)
    var v9 uint16 = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -429
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    v4 := []uint32{461, 2147483648, 1, 65535}
    var v5 float64 = 0.0
    var v6 bool = (false || true)
    if v6 {
        var v10 uint32 = 263
        if (746 != 65535) {
            v11 := GenS1{F0: 42}
            switch "hello" {
            case "hello":
                for fi12 := 0; fi12 < 11; fi12++ {
                    if v6 {
                        v5 -= (v5 - v5)
                        var v13 string = "\n\t"
                        _ = v13
                    } else {
                        var v14 int = (fi12 | -1)
                        v15 := GenS1{F0: 2}
                        var v16 int32 = 2147483647
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 string = ""
                    v18 := GenS2{F0: 42, F1: v1}
                    _ = fi12
                    _ = v17
                    _ = v18
                }
            }
            v3 = v3
            _ = v11
        } else {
            v2[v10] = (false || v6)
            for fi19 := 0; fi19 < 5; fi19++ {
                for fi20 := 0; fi20 < 17; fi20++ {
                    if ((fi19 % (((fi20 ^ fi19)) | 1)) < 2147483647) {
                        v3 = v3
                    }
                    v21 := GenS3{F0: 633}
                    v22 := GenS4{F0: -678, F1: v6, F2: 947}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                if v6 {
                    v5 += v5
                    v3 += v3
                }
                _ = fi19
            }
        }
        _ = v10
    }
    switch ("\n\t" + "hello") {
    case "hello":
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 8; fi24++ {
                switch fi24 {
                case -2147483648:
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        if v6 {
                            var v25 int8 = (v1 + (v1 | v1))
                            var v26 uint8 = 242
                            _ = v25
                            _ = v26
                        } else {
                            var v27 float64 = (v5 - 3.14159)
                            v1 -= v1
                            _ = v27
                        }
                    case "":
                        var v28 uint = 255
                        _ = v28
                    }
                }
                v6 = false
                _ = fi24
            }
            v29 := GenS5{F0: 42, F1: (fi23 % ((fi23) | 1)), F2: 1}
            if true {
                var v30 float64 = -2.25
                _ = v30
            }
            _ = fi23
            _ = v29
        }
    case "foo bar":
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int = (-(fi31))
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := GenS1{F0: 0}
                v35 := GenS2{F0: 42, F1: v1}
                if ((v6 && ((false || (v6 && v6)) && false)) && (false && v6)) {
                    switch 4294967295 {
                    case -840:
                        v2[65535] = v6
                    case 665:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    case 2147483647:
                        var v36 float32 = -2.25
                        switch "a" {
                        case "":
                            if (v6 || false) {
                                v2[65535] = ((^(fi33)) > 32767)
                                v37 := GenS3{F0: 18446744073709551615}
                                _ = v37
                            }
                        case "foo bar":
                            v38 := GenS2{F0: 42, F1: v1}
                            var v39 uint = 4294967295
                            _ = v38
                            _ = v39
                        case "\n\t":
                            v5 -= 0.1
                        }
                        _ = v36
                    }
                    var v40 uint8 = v3
                    _ = v40
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            var v41 float32 = -2.25
            _ = fi31
            _ = v32
            _ = v41
        }
    default:
        v42 := GenS6{F0: 0}
        _ = v42
    }
    v43 := GenS7{F0: -33, F1: 65535, F2: -32768}
    var v44 float64 = (-2.25 / 1.0)
    if len(v4) > 0 {
        v4[0] = 65535
    }
    v1 += (^(v1))
    v5 = (0.0 + (1e10 / v44))
    fmt.Println(v44)
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 int32 = 912
        delete(v2, 0)
        v4 = append(v4, 2147483648)
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 uint }
type GenS3 struct { F0 int; F1 uint32; F2 string }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 282
    v3 := []string{"\n\t", "foo bar", ("hello" + "hello"), ("" + "\n\t")}
    v4 := make(map[int]string)
    var v5 uint = 2
    var v6 float32 = 1.5
    fmt.Println(v2, v6)
    var v7 int16 = 0
    var v8 float32 = v6
    for fi9 := 0; fi9 < 14; fi9++ {
        v1 -= fi9
        if (!false) {
            var v10 uint = (v5 % (((v5 % (((^(v5))) | 1))) | 1))
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int16 = int16(v10)
                var v13 bool = (543 != (v2 - (v2 ^ (v2 | 552))))
                _ = fi11
                _ = v12
                _ = v13
            }
            if false {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v10 = v10
                    var v15 float64 = 1.5
                    if ((1 <= -1) && (!(0 == 0))) {
                        switch v1 {
                        case 271:
                            _ = len(v3)
                        case -2147483648:
                            if true {
                                fmt.Println(v8, fi14, v6)
                                v2 += (^(v2))
                            }
                        }
                        v10 = 42
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v10 += (v10 - v10)
                    _ = fi16
                }
            }
            _ = v10
        } else {
            fmt.Println(fi9, v2)
            var v17 float32 = float32(3.14159)
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            switch "" {
            case "foo bar":
                var v19 uint8 = 128
                _ = v19
            default:
                if (false && true) {
                    var v20 int64 = -9223372036854775808
                    if true {
                        v8 += (v8 / v6)
                        var v21 uint = (v5 << (((^(v5))) & 63))
                        v2 -= 42
                        _ = v21
                    }
                    v22 := v4[(^(v1))]
                    v23 := GenS1{F0: int32(v7), F1: 2}
                    v2 -= v2
                    _ = v20
                    _ = v22
                    _ = v23
                }
            }
            var v24 float32 = ((v6 - 1e10) - (v8 - 1e10))
            var v25 int = (fi18 * -2147483648)
            _ = fi18
            _ = v24
            _ = v25
        }
        _ = fi9
    }
    var v26 float32 = (((0.1 / (v8 / 1e10)) - ((float32(3.14159) - v6) - v8)) - 1e10)
    v2 -= (v2 | v2)
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(v5, v6, v1)
        var v28 uint32 = 2
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if false {
            var v30 uint8 = uint8(v5)
            v30 = 0
            _ = v30
        } else {
            v4[127] = "foo bar"
            v5 -= (v5 - (v5 + v5))
        }
        var v31 uint16 = v2
        _ = fi29
        _ = v31
    }
    v32 := v4[v1]
    for fi33 := 0; fi33 < 17; fi33++ {
        if false {
            if false {
                if (fi33 >= (fi33 - fi33)) {
                    v2 += (v2 >> (((^(v5))) & 63))
                    v6 = (1.5 / v8)
                } else {
                    if (!false) {
                        v8 -= float32(0.0)
                        var v34 bool = (!(127 < 128))
                        _ = v34
                    } else {
                        var v35 uint = v5
                        _ = v35
                    }
                }
                v5 -= v5
                v36 := v4[fi33]
                _ = v36
            }
        }
        _ = fi33
    }
    if true {
        if (42 >= 741) {
            switch v1 {
            case -611:
                v37 := GenS1{F0: -29, F1: 150}
                _ = v37
            }
            fmt.Println(v1, v5)
        } else {
            v38 := GenS1{F0: -2147483648, F1: 4294967295}
            fmt.Println(v26)
            _ = v38
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            if ((false || true) && true) {
                v40 := GenS1{F0: 2147483647, F1: uint32(v5)}
                v41 := GenS1{F0: 0, F1: 13}
                _ = len(v3)
                _ = v40
                _ = v41
            }
            v42 := GenS2{F0: (!false), F1: 617, F2: v5}
            for fi43 := 0; fi43 < 8; fi43++ {
                fmt.Println(v7, v6, v2)
                v44 := GenS3{F0: (fi39 & (fi43 + fi43)), F1: 4294967295, F2: "\n\t"}
                _ = fi43
                _ = v44
            }
            _ = fi39
            _ = v42
        }
        if (!(2 < 2)) {
            _ = len(v3)
        } else {
            v45 := GenS2{F0: true, F1: v1, F2: v5}
            v1 = (v1 + (v1 - v1))
            _ = v45
        }
    } else {
        v46 := GenS1{F0: -2147483648, F1: 668}
        _ = v46
    }
    var v47 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

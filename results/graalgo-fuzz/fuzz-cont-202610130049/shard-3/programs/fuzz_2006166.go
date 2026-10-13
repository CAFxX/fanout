package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
type GenS4 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 505
    v3 := make(map[int]int)
    var v4 uint32 = v2
    var v5 uint8 = 127
    var v6 bool = (!false)
    v2 -= (v4 / (((v4 - v4)) | 1))
    fmt.Println(v2, v6)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (-(fi7)) {
        case -293:
            v5 += (v5 - 8)
        case 1:
            var v8 uint64 = 4294967295
            _ = v8
        case -1:
            if (1000 > 1000) {
                var v9 uint = 944
                switch 1 {
                case -866:
                    var v10 bool = ((v4 / (((v2 * (v2 ^ (v2 + v4)))) | 1)) == v4)
                    _ = v10
                case -678:
                    fmt.Println(fi7, v4)
                }
                var v11 float64 = 3.14159
                _ = v11
                _ = v9
            }
        }
        v12 := GenS1{F0: "a", F1: (fi7 * int(v5))}
        _ = fi7
        _ = v12
    }
    var v13 int8 = 6
    if v6 {
        switch "a" {
        case "hello":
            v14 := GenS2{F0: (v4 / ((v2) | 1)), F1: v13}
            _ = v14
        case "foo bar":
            var v15 int32 = -740
            _ = v15
        default:
            switch 1 {
            case -2147483648:
                v6 = v6
            }
        }
        if false {
            v2 -= 2147483647
        } else {
            delete(v3, -553)
            v16 := GenS3{F0: (v4 | v4), F1: 159, F2: -830}
            _ = v16
        }
    }
    var v17 int64 = 0
    if false {
        var v18 int64 = v17
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS4{F0: (^(v2)), F1: ""}
                fmt.Println(v4)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        var v22 int = int(v2)
        _ = v18
        _ = v22
    } else {
        v23 := GenS1{F0: "", F1: 382}
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS3{F0: v4, F1: 42, F2: 1}
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v5 = v5
    for fi26 := 0; fi26 < 11; fi26++ {
        if (!(0 != v5)) {
            var v27 uint = 65535
            for fi28 := 0; fi28 < 18; fi28++ {
                if ((true || ((!false) && false)) || ((342 > -2147483648) && (!v6))) {
                    fmt.Println(v5, v4)
                    v4 += (v2 / (((v4 ^ (v2 ^ 2147483648))) | 1))
                }
                _ = fi28
            }
            if (v17 < v17) {
                v29 := []int{750}
                _ = len(v29)
                _ = v29
            }
            _ = v27
        }
        var v30 uint64 = 18446744073709551615
        _ = fi26
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint8 = 0
        _ = fi31
        _ = v32
    }
    v5 = (v5 & (v5 & 127))
    switch 2147483647 {
    case -90:
        switch "" {
        case "a":
            fmt.Println(v5, v6)
        case "hello":
            switch 187 {
            case 0:
                if (int32(v4) >= 42) {
                    _ = len(v3)
                    var v33 int16 = 1
                    if v6 {
                        v34 := GenS4{F0: 429, F1: (((("\n\t" + "\n\t") + (("a" + "foo bar") + (((("" + "a") + "foo bar") + "hello") + ("foo bar" + "")))) + "foo bar") + "a")}
                        fmt.Println(v17, v33)
                        var v35 int = 42
                        _ = v34
                        _ = v35
                    } else {
                        v5 -= v5
                        for fi36 := 0; fi36 < 7; fi36++ {
                            for fi37 := 0; fi37 < 3; fi37++ {
                                v38 := v3[fi37]
                                v2 -= (^(v2))
                                fmt.Println(v17, v13, v2)
                                _ = fi37
                                _ = v38
                            }
                            _ = fi36
                        }
                        var v39 uint8 = 2
                        _ = v39
                    }
                    _ = v33
                } else {
                    v40 := GenS1{F0: "\n\t", F1: -714}
                    var v41 int = 127
                    _ = v40
                    _ = v41
                }
            }
        case "foo bar":
            v2 = v2
        }
    default:
        fmt.Println(v17, v17)
    }
    var v42 uint16 = 204
    if v6 {
        v42 += v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 int64, p1 uint16) int8 {
    var v7 bool = ((2147483647 != -2147483648) && (("a" + "\n\t") < "foo bar"))
    var v8 uint16 = (p1 & 987)
    var v9 uint32 = 1
    if ((p0 | (p0 | (p0 | p0))) < (p0 << ((148) & 63))) {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch "\n\t" {
            case "":
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS1{F0: 0}
                    _ = fi11
                    _ = v12
                }
            case "foo bar":
                for fi13 := 0; fi13 < 17; fi13++ {
                    p0 = (p0 | (^(p0)))
                    _ = fi13
                }
            case "a":
                if v7 {
                    p0 = p0
                    var v14 bool = false
                    _ = v14
                }
            default:
                var v15 uint16 = v8
                _ = v15
            }
            _ = fi10
        }
        v8 += (p1 / (((^(p1))) | 1))
    } else {
        if (((v8 | v8) < 7) || v7) {
            p0 -= p0
            switch "foo bar" {
            case "\n\t":
                var v16 uint8 = 1
                _ = v16
            case "a":
                p0 = -9223372036854775808
            default:
                switch "" {
                case "":
                    v17 := GenS2{F0: v9, F1: (^(v9))}
                    _ = v17
                case "\n\t":
                    var v18 float64 = (0.0 / 1.0)
                    _ = v18
                case "a":
                    var v19 int = 9223372036854775807
                    _ = v19
                default:
                    v8 = (v8 - v8)
                }
            }
        } else {
            if (-367 == 1) {
                v9 = (v9 >> ((255) & 63))
            }
            p1 -= 65535
            if false {
                v20 := GenS3{F0: 4294967295, F1: uint8(v9)}
                v21 := []int{-678, -531, -9223372036854775808}
                _ = v20
                _ = v21
            } else {
                p0 = p0
                var v22 int32 = -338
                _ = v22
            }
        }
        if (-1 > -2147483648) {
            for fi23 := 0; fi23 < 11; fi23++ {
                if (42 <= 1) {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        fmt.Println(p1, p0)
                        v9 -= v9
                        var v25 int8 = 1
                        _ = fi24
                        _ = v25
                    }
                    v26 := []int{42, (fi23 - (fi23 & 127)), (fi23 ^ fi23), (fi23 << ((731) & 63))}
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := GenS2{F0: (v9 & v9), F1: (^(v9))}
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
                v29 := []int{(fi23 | fi23), (-(fi23))}
                var v30 uint8 = 0
                _ = fi23
                _ = v29
                _ = v30
            }
        } else {
            var v31 int32 = -2147483648
            if v7 {
                var v32 bool = (45 == 127)
                _ = v32
            }
            _ = v31
        }
        var v33 uint = 65535
        _ = v33
    }
    var v34 int8 = -1
    _ = p0
    _ = p1
    _ = v34
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 int32 = -2147483648
    v3 := []uint32{1}
    var v4 uint32 = 310
    var v5 uint = 350
    var v6 int64 = 9223372036854775807
    var v35 int64 = (v6 + (-(v6)))
    for fi36 := 0; fi36 < 10; fi36++ {
        v4 = v4
        v37 := GenS3{F0: (v5 % ((v5) | 1)), F1: (v1 / ((v1) | 1))}
        var v38 float64 = 3.14159
        _ = fi36
        _ = v37
        _ = v38
    }
    v3 = append(v3, 2)
    v6 = -1
    v3 = append(v3, (v4 << ((1) & 63)))
    fmt.Println(v5, v35)
    var v39 int16 = 1
    v1 = v1
    fmt.Println(v35, v39)
    switch "foo bar" {
    case "\n\t":
        switch 565 {
        case 9223372036854775807:
            for fi40 := 0; fi40 < 7; fi40++ {
                v3 = append(v3, v4)
                fmt.Println(v4, v2)
                _ = fi40
            }
        default:
            if true {
                for fi41 := 0; fi41 < 3; fi41++ {
                    var v42 float64 = (0.0 + ((((-2.25 / -2.25) * -2.25) + (float64(3.14159) * 1e10)) + 0.1))
                    v43 := GenS2{F0: v4, F1: v4}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                var v44 int64 = -9223372036854775808
                for fi45 := 0; fi45 < 13; fi45++ {
                    v46 := GenS3{F0: (^(v5)), F1: 128}
                    var v47 uint32 = (v4 >> ((v5) & 63))
                    var v48 bool = (true || (uint16(v5) > 902))
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
                _ = v44
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    v49 := GenS4{F0: (v2 >> (((v5 ^ v5)) & 63)), F1: ((true || (true && ("foo bar" <= "hello"))) || (1.5 > (1e10 / float64(1e10)))), F2: v39}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}

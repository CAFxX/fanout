package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 string }
func genFn0(p0 uint32, p1 int8, p2 uint64) uint {
    v5 := GenS1{F0: p1}
    switch 128 {
    case 535:
        var v6 int = 128
        _ = v6
    }
    if false {
        p0 -= 586
        fmt.Println(p0)
    } else {
        p0 += 2
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(p1, p2)
            var v8 string = "foo bar"
            for fi9 := 0; fi9 < 20; fi9++ {
                switch fi7 {
                case 9223372036854775807:
                    var v10 uint16 = 400
                    _ = v10
                case 2147483647:
                    v11 := GenS1{F0: -107}
                    _ = v11
                case -498:
                    var v12 int8 = (^(p1))
                    _ = v12
                }
                v8 = v8
                switch (fi9 << ((42) & 63)) {
                case -1:
                    v13 := GenS1{F0: (p1 - p1)}
                    _ = v13
                case 341:
                    var v14 uint = 2
                    _ = v14
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        var v15 int32 = 2147483647
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-443, 4294967295, -844}
    var v3 uint16 = 273
    var v4 uint64 = 775
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint32 = uint32(v3)
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 uint32 = (v17 / ((v17) | 1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v17, v1, v17)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS2{F0: 1, F1: 1, F2: v17}
            var v22 uint8 = 0
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi16
        _ = v17
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 uint32 = 693
        if false {
            if (true && (false && false)) {
                var v25 uint16 = uint16(v4)
                v3 = uint16(v4)
                v2 = append(v2, int(v4))
                _ = v25
            }
        } else {
            switch "" {
            case "\n\t":
                fmt.Println(v24)
            case "foo bar":
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 uint32 = v24
                    var v28 float64 = (0.1 / -2.25)
                    v24 = v24
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                v29 := GenS3{F0: v4, F1: 255}
                _ = v29
            }
            v30 := GenS2{F0: -863, F1: 161, F2: v24}
            if (true || true) {
                v4 = v4
                var v31 int32 = 536
                if (false && false) {
                    fmt.Println(v3, v31)
                    v32 := GenS4{F0: (880 > v3), F1: v1}
                    _ = v32
                }
                _ = v31
            }
            _ = v30
        }
        var v33 int64 = (v1 * (v1 | 0))
        _ = fi23
        _ = v24
        _ = v33
    }
    v34 := GenS2{F0: v1, F1: 4294967295, F2: 689}
    v35 := GenS4{F0: (false && true), F1: (v1 % (((-(v1))) | 1))}
    for fi36 := 0; fi36 < 3; fi36++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            fmt.Println(v1)
        }
        _ = fi36
    }
    var v37 int = 42
    var v38 uint = 263
    if true {
        switch 972 {
        case -734:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            for fi39 := 0; fi39 < 8; fi39++ {
                var v40 float64 = (1e10 * 1e10)
                v1 = (v1 << ((65535) & 63))
                switch -52 {
                case 450:
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        for fi41 := 0; fi41 < 9; fi41++ {
                            var v42 uint16 = (v3 - (v3 % ((738) | 1)))
                            _ = fi41
                            _ = v42
                        }
                    case "\n\t":
                        v38 += (^(v38))
                    }
                case 1:
                    var v43 int16 = 254
                    _ = v43
                case 0:
                    if true {
                        v44 := GenS3{F0: (v4 ^ (v4 % ((0) | 1))), F1: 1}
                        _ = v44
                    }
                default:
                    if ((v4 ^ (v4 << ((genFn0(87, -1, v4)) & 63))) == 1) {
                        var v45 int32 = int32(fi39)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v45 += (v45 | v45)
                            _ = fi46
                        }
                        v37 -= (v37 / ((-45) | 1))
                        _ = v45
                    }
                }
                _ = fi39
                _ = v40
            }
        }
        v38 += v38
    } else {
        fmt.Println(v3)
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 = 9223372036854775807
        _ = fi47
    }
    if len(v2) > 0 {
        v2[0] = (v37 + (v37 | 128))
    }
    var v48 float32 = float32((3.14159 * 3.14159))
    v49 := GenS5{F0: "hello"}
    for fi50 := 0; fi50 < 17; fi50++ {
        var v51 int16 = 0
        v48 += v48
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v48
    _ = v49
}

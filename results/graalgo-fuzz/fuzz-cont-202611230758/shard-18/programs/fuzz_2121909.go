package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 uint16, p1 uint16, p2 uint) int16 {
    v5 := []int{298}
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 0
    var v3 uint32 = 4294967295
    v4 := []int{-1}
    if false {
        if false {
            v6 := GenS1{F0: v3}
            _ = len(v4)
            _ = v6
        }
        v7 := GenS1{F0: (v3 >> ((1) & 63))}
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: v3}
            switch ((((("hello" + "\n\t") + "foo bar") + ("a" + (((("foo bar" + "foo bar") + "\n\t") + "foo bar") + ""))) + "\n\t") + "hello") {
            case "hello":
                _ = len(v4)
            }
            if len(v4) > 0 {
                v4[0] = (-(fi8))
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        if len(v4) > 0 {
            v4[0] = 646
        }
        var v10 uint8 = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v3 -= v3
        _ = fi11
    }
    v4 = append(v4, 2147483647)
    for fi12 := 0; fi12 < 17; fi12++ {
        if false {
            v2 += 0
        }
        _ = fi12
    }
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 721 {
    case 688:
        v3 -= v3
    default:
        v13 := GenS1{F0: v3}
        _ = v13
    }
    var v14 float32 = 1.5
    switch "a" {
    case "":
        v15 := GenS1{F0: 1}
        _ = v15
    case "\n\t":
        var v16 bool = (false || (((0 < v1) || false) && (0 == 128)))
        _ = v16
    }
    if (v1 == 1) {
        var v17 float32 = ((((v14 * 1e10) * -2.25) - (v14 * -2.25)) * v14)
        if (v14 > float32(0.0)) {
            if true {
                v18 := GenS2{F0: 42}
                v2 += (v2 + v2)
                _ = v18
            } else {
                switch 856 {
                case 836:
                    switch -2147483648 {
                    case 957:
                        switch 128 {
                        case 32768:
                            fmt.Println(v3, v14, v2)
                        default:
                            v19 := GenS2{F0: uint16(v3)}
                            _ = v19
                        }
                    case -815:
                        switch "" {
                        case "":
                            fmt.Println(v17, v3, v17)
                        case "hello":
                            var v20 int8 = 93
                            _ = v20
                        }
                    case 595:
                        var v21 bool = ("\n\t" < "foo bar")
                        _ = v21
                    default:
                        v4 = append(v4, 9223372036854775807)
                    }
                case 128:
                    v22 := GenS1{F0: (v3 >> ((v1) & 63))}
                    _ = v22
                case -755:
                    fmt.Println(v1, v1)
                default:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v4 = append(v4, 426)
                        _ = fi23
                    }
                }
                v24 := GenS3{F0: "", F1: 42}
                _ = v24
            }
        } else {
            var v25 int = 2147483647
            v2 -= 2
            for fi26 := 0; fi26 < 10; fi26++ {
                if (64 > 167) {
                    var v27 uint32 = (v3 / (((v3 * (v3 % (((v3 | v3)) | 1)))) | 1))
                    if len(v4) > 0 {
                        v4[0] = fi26
                    }
                    _ = v27
                } else {
                    switch v25 {
                    case 127:
                        v28 := GenS4{F0: -1, F1: int(v3)}
                        _ = v28
                    case 768:
                        var v29 uint64 = (v2 + (v2 / (((v2 * (^(v2)))) | 1)))
                        _ = v29
                    }
                    if len(v4) > 0 {
                        v4[0] = (v25 << (((v1 % ((2) | 1))) & 63))
                    }
                }
                _ = fi26
            }
            _ = v25
        }
        switch 32768 {
        case -9223372036854775808:
            v4 = append(v4, 1)
        }
        _ = v17
    } else {
        fmt.Println(v1, v1)
        _ = len(v4)
    }
    var v30 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}

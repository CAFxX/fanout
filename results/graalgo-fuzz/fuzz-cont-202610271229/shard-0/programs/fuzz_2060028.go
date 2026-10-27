package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    v5 := v4[1]
    for fi6 := 0; fi6 < 18; fi6++ {
        switch ("hello" + (("" + "a") + "foo bar")) {
        case "":
            switch fi6 {
            case 4294967295:
                v7 := GenS1{F0: 2, F1: (v3 - (v3 | 564))}
                _ = v7
            }
        case "a":
            delete(v1, (v5 + (v5 ^ 504)))
        default:
            v4[(fi6 + -2147483648)] = v5
        }
        _ = fi6
    }
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 19; fi8++ {
            if false {
                for fi9 := 0; fi9 < 19; fi9++ {
                    v5 -= (v5 + uint32(v3))
                    _ = fi9
                }
                fmt.Println(v3)
            }
            v10 := v4[(fi8 / ((fi8) | 1))]
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(fi8, fi11, v3)
                v12 := GenS1{F0: 65535, F1: (v3 / (((v3 >> ((908) & 63))) | 1))}
                switch "a" {
                case "":
                    var v13 string = ("foo bar" + "")
                    _ = v13
                case "foo bar":
                    if false {
                        v3 -= (v3 >> ((4294967295) & 63))
                    }
                default:
                    switch -2147483648 {
                    case 1:
                        if true {
                            v3 += v3
                        }
                    }
                    v3 = v3
                    delete(v4, fi11)
                }
                var v14 int = fi8
                v5 = v10
                var v15 uint8 = 128
                v15 += (v15 << ((65535) & 63))
                _ = fi11
                _ = v12
                _ = v14
                _ = v15
            }
            _ = fi8
            _ = v10
        }
    }
    v3 -= v3
    var v16 float32 = (3.14159 + (-2.25 / ((1.5 + (-2.25 - -2.25)) / (((1.5 * (float32((-2.25 / 3.14159)) - (1.5 - -2.25))) * 1.5) / -2.25))))
    v17 := []int{}
    if true {
        for fi18 := 0; fi18 < 13; fi18++ {
            switch "\n\t" {
            case "a":
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int16 = 32767
                    v21 := GenS1{F0: 16, F1: (v3 >> ((0) & 63))}
                    var v22 uint16 = 42
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    var v23 float32 = ((float32(-2.25) * (float32(float64(v16)) - ((v16 / 3.14159) + v16))) * v16)
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            case "hello":
                v5 += v5
            default:
                v4[1] = (v5 | v5)
            }
            v16 = (float32(0.0) - 0.0)
            _ = fi18
        }
    } else {
        if true {
            var v24 uint64 = 934
            for fi25 := 0; fi25 < 14; fi25++ {
                fmt.Println(v24)
                _ = fi25
            }
            fmt.Println(v3)
            _ = v24
        }
        v26 := GenS1{F0: 1, F1: (v3 + (v3 ^ v3))}
        _ = v26
    }
    v16 = (0.0 / (-2.25 * float32((0.1 - (0.0 / -2.25)))))
    for fi27 := 0; fi27 < 16; fi27++ {
        if false {
            v16 += 0.1
        } else {
            switch fi27 {
            case 4294967295:
                v3 = (v3 - 18446744073709551615)
            case 2147483647:
                for fi28 := 0; fi28 < 15; fi28++ {
                    v5 -= (v5 | v5)
                    var v29 uint64 = v3
                    fmt.Println(fi28, v16)
                    _ = fi28
                    _ = v29
                }
            case -9223372036854775808:
                v30 := GenS2{F0: 148}
                _ = v30
            }
        }
        v31 := GenS1{F0: 546, F1: v3}
        var v32 string = ""
        _ = fi27
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() string {
    var v4 int = 32768
    var v5 uint16 = 2
    var v6 bool = (0.0 == 0.0)
    _ = v4
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v7 uint64 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 42
        _ = fi8
        _ = v9
    }
    if false {
        var v10 uint64 = 385
        fmt.Println(v10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint8 = 42
            if (1000 < -745) {
                if true {
                    var v13 int8 = -1
                    v7 = 382
                    _ = v13
                }
                _ = len(v3)
            }
            _ = fi11
            _ = v12
        }
        v14 := GenS1{F0: 1}
        for fi15 := 0; fi15 < 3; fi15++ {
            if (1 <= -451) {
                v7 = (v7 % (((^(v7))) | 1))
                switch "foo bar" {
                case "\n\t":
                    if true {
                        var v16 int = fi15
                        _ = v16
                    }
                default:
                    if true {
                        v7 = 42
                    }
                }
            } else {
                v7 -= v7
                switch fi15 {
                case 4294967295:
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v18 := GenS2{F0: 180, F1: 2147483648}
                        _ = fi17
                        _ = v18
                    }
                case 32768:
                    var v19 int16 = 32767
                    _ = v19
                case 543:
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v7 += 4294967295
                        fmt.Println(v7)
                        var v21 int8 = 42
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 float32 = (-2.25 * (float32(1e10) / (3.14159 + ((0.1 * (1e10 * -2.25)) + 1e10))))
                _ = v22
            }
            _ = fi15
        }
        _ = v14
    }
    var v23 float32 = (1e10 / 0.1)
    v23 = (v23 * (v23 * v23))
    var v24 uint32 = 0
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            if true {
                if false {
                    v27 := GenS2{F0: 2, F1: (^(v24))}
                    v28 := GenS2{F0: 917, F1: (v24 * v24)}
                    _ = v27
                    _ = v28
                } else {
                    var v29 float32 = v23
                    delete(v3, (fi26 << ((953) & 63)))
                    _ = v29
                }
                var v30 uint8 = 1
                fmt.Println(v30, fi25)
                if len(v1) > 0 {
                    v1[0] = (v24 << ((2) & 63))
                }
                var v31 bool = true
                _ = v30
                _ = v31
            } else {
                var v32 bool = (68 >= -1)
                v7 = v7
                switch "\n\t" {
                case "foo bar":
                    var v33 uint64 = 4294967295
                    _ = v33
                case "hello":
                    var v34 float64 = float64(v23)
                    _ = v34
                case "a":
                    if (true || ("\n\t" > "")) {
                        v32 = (!((!true) || ((v32 || (-9223372036854775808 < -464)) && false)))
                    }
                default:
                    var v35 bool = (2 != 255)
                    _ = v35
                }
                fmt.Println(v7, fi26, v23)
                var v36 uint = 4294967295
                _ = v32
                _ = v36
            }
            v37 := v2[(^(v24))]
            v38 := GenS2{F0: 2, F1: v24}
            fmt.Println(v24, v7)
            switch fi25 {
            case 1:
                v37 = "foo bar"
            case 127:
                v37 = v37
            }
            _ = fi26
            _ = v37
            _ = v38
        }
        _ = fi25
    }
    var v39 bool = (false || (uint(v24) < 188))
    v24 -= (v24 | (v24 << ((625) & 63)))
    var v40 uint32 = (v24 ^ (v24 - (v24 * (v24 >> ((0) & 63)))))
    v23 = float32(0.0)
    var v41 float64 = (1e10 + -2.25)
    var v42 float64 = (1e10 * float64((((-2.25 * float32((v41 / float64((v23 / (v23 + 1.5)))))) * ((-2.25 + (v23 / v23)) * v23)) / (((((((1.5 - 0.0) - float32(v41)) / v23) / v23) * ((((1.5 * ((((v23 * v23) + 0.1) * v23) * 3.14159)) - v23) * ((v23 - v23) + (-2.25 / 3.14159))) / 1.5)) - v23) + v23))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v7
}

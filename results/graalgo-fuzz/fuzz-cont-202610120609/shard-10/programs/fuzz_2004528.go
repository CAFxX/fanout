package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 int32 = -360
    var v2 int = -2147483648
    var v3 float64 = (0.1 + (((0.1 * ((0.0 / -2.25) / float64(1e10))) * 1.5) - 0.1))
    v4 := make(map[uint32]string)
    var v5 uint32 = 0
    v5 -= v5
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            delete(v4, (v5 - (v5 | v5)))
            var v7 bool = (84 != 255)
            v8 := GenS1{F0: (fi6 != (fi6 | (^(fi6))))}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 2; fi9++ {
                if (true || true) {
                    var v10 uint32 = 0
                    _ = v10
                } else {
                    var v11 int64 = -177
                    v5 += v5
                    if (!(false && (!(42 <= -1)))) {
                        if (!false) {
                            v12 := GenS1{F0: (!(!true))}
                            v13 := GenS1{F0: (!(4294967295 <= 2))}
                            var v14 int64 = v11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v3 += (3.14159 + v3)
                    }
                    _ = v11
                }
                v15 := []int{-342, fi9, -1, v2}
                _ = fi9
                _ = v15
            }
            v3 -= v3
        }
        _ = fi6
    }
    v5 = v5
    var v16 int64 = 0
    if (290 > 511) {
        if true {
            v1 = (v1 ^ (v1 / (((^(v1))) | 1)))
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(fi17)
                var v18 bool = (881 <= 2)
                _ = fi17
                _ = v18
            }
        }
        v1 = int32(v5)
        var v19 int8 = -1
        _ = v19
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5, v16)
    case "":
        v20 := GenS1{F0: (0 >= 197)}
        _ = v20
    default:
        if true {
            if false {
                v21 := v4[(^(v5))]
                _ = v21
            } else {
                var v22 float32 = ((float32(v3) + 3.14159) - -2.25)
                v5 += (^(v5))
                _ = v22
            }
        } else {
            v23 := GenS1{F0: (v3 <= (1e10 / v3))}
            var v24 string = "a"
            _ = v23
            _ = v24
        }
    }
    v25 := []int{v2, v2}
    switch ("foo bar" + "") {
    case "\n\t":
        if (float64(1e10) != v3) {
            for fi26 := 0; fi26 < 7; fi26++ {
                switch fi26 {
                case 0:
                    var v27 float64 = v3
                    _ = v27
                default:
                    var v28 uint16 = 81
                    _ = v28
                }
                _ = fi26
            }
        } else {
            v2 = (-(v2))
        }
    case "hello":
        var v29 uint32 = (v5 >> ((2) & 63))
        _ = v29
    case "foo bar":
        fmt.Println(v16, v16, v3)
    default:
        var v30 int16 = 1000
        _ = v30
    }
    fmt.Println(v2, v3, v16)
    for fi31 := 0; fi31 < 6; fi31++ {
        switch fi31 {
        case 32767:
            var v32 int16 = 42
            _ = v32
        default:
            v16 = -1
        }
        _ = fi31
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 int = (fi33 + fi33)
        v16 = int64(v1)
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: ((1 > (^(v5))) && false)}
    fmt.Println(v5, v2, v1)
    if false {
        var v36 float64 = v3
        _ = v36
    }
    v37 := GenS2{F0: 1, F1: v16}
    fmt.Println(v5)
    v38 := GenS3{F0: (v2 + (v2 * v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}

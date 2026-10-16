package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
func genFn0(p0 uint8, p1 int16, p2 uint32) uint {
    p0 = (p0 & 235)
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 58
}
func genFn1(p0 int16, p1 uint64, p2 int8) uint32 {
    var v5 uint16 = 362
    v6 := []int{42, 42}
    p0 = 747
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 42
    v4 = 138
    fmt.Println(v4, v1)
    var v7 string = "a"
    if false {
        var v8 int8 = 42
        v1 = v1
        _ = v8
    } else {
        v1 += 2147483647
        _ = len(v3)
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = (v1 | (v1 ^ 268))
        switch "hello" {
        case "hello":
            if true {
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: 273}
                    _ = fi10
                    _ = v11
                }
            } else {
                v4 += (^(v4))
                var v12 int8 = 0
                switch v7 {
                case "":
                    var v13 float64 = float64((0.1 - 0.1))
                    _ = v13
                }
                _ = v12
            }
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch v7 {
        case "\n\t":
            v3[42] = 42
        case "a":
            v4 -= 128
        case "foo bar":
            if (!(!((fi14 + (fi14 ^ (fi14 >> ((49) & 63)))) != (^(fi14))))) {
                v4 += 42
                var v15 uint16 = 65535
                _ = v15
            } else {
                if true {
                    switch v7 {
                    case "\n\t":
                        var v16 int32 = 1
                        _ = v16
                    case "":
                        v4 += (v4 >> ((0) & 63))
                    case "foo bar":
                        switch (("foo bar" + v7) + v7) {
                        case "foo bar":
                            v1 = (v1 ^ v1)
                        }
                    default:
                        v4 -= v4
                    }
                }
                var v17 uint64 = 178
                _ = v17
            }
        default:
            v7 = "a"
        }
        if (((v7 + v7) + (v7 + v7)) <= v7) {
            v7 = "\n\t"
            if true {
                v1 -= (-(v1))
                if ("a" >= v7) {
                    v18 := []int{128}
                    _ = v18
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := GenS1{F0: uint(v4)}
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = fi14
    }
    var v21 float32 = (((0.0 * 0.1) - ((0.1 + 1e10) - 1.5)) * (-2.25 / (-2.25 * (-2.25 * float32(3.14159)))))
    v3[562] = 1
    v7 = v7
    var v22 int32 = (v1 * v1)
    v21 += v21
    v21 = 0.1
    fmt.Println(v1, v7, v7)
    if (1 <= 42) {
        v4 -= 255
    } else {
        var v23 string = ((v7 + "\n\t") + v7)
        var v24 int32 = v1
        if (!("foo bar" < v23)) {
            v25 := GenS1{F0: 2}
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v24, fi26, v22)
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 float32 = 1e10
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            v4 -= (v4 + (v4 ^ (v4 / ((v4) | 1))))
            _ = v25
        }
        _ = v23
        _ = v24
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS2{F0: 2, F1: 32767}
        for fi31 := 0; fi31 < 6; fi31++ {
            switch 0 {
            case -14:
                switch fi29 {
                case 127:
                    v22 = v22
                case 32768:
                    fmt.Println(fi31, fi31, v22)
                case -928:
                    v32 := GenS1{F0: 394}
                    _ = v32
                default:
                    v33 := v2[v7]
                    _ = v33
                }
            default:
                v1 += (v1 >> ((1) & 63))
            }
            var v34 int64 = 9223372036854775807
            v35 := GenS3{F0: (118 > 42), F1: fi31, F2: 32767}
            _ = fi31
            _ = v34
            _ = v35
        }
        v36 := []int{fi29, (fi29 - 127)}
        _ = fi29
        _ = v30
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}

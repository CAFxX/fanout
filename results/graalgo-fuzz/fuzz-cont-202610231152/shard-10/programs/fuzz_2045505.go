package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int16, p1 string, p2 uint32) uint8 {
    var v6 uint = 1
    var v7 int16 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 112
}

func main() {
    var v1 int16 = 1
    var v2 int8 = -1
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint16 = 235
    switch 795 {
    case -2147483648:
        var v8 int64 = -270
        _ = v8
    case 4294967295:
        v5 -= (v5 & v5)
    case 127:
        v2 = 103
    }
    v5 = v5
    v1 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        v5 -= v5
        v5 += (v5 & (v5 >> ((65535) & 63)))
        v1 -= v1
        _ = fi9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float32 = float32(((1.5 / -2.25) * 0.1))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int32 = -2147483648
        _ = fi12
        _ = v13
    }
    fmt.Println(v5, v5)
    v14 := GenS1{F0: false}
    var v15 float64 = -2.25
    v2 = v2
    var v16 int32 = -1
    switch "foo bar" {
    case "a":
        v1 -= (v1 | 42)
    case "foo bar":
        for fi17 := 0; fi17 < 17; fi17++ {
            v16 += (^(v16))
            for fi18 := 0; fi18 < 10; fi18++ {
                if (true || false) {
                    v19 := []int{(fi17 << ((961) & 63)), (^(fi17)), (fi18 << ((2) & 63))}
                    var v20 uint8 = 127
                    _ = v19
                    _ = v20
                } else {
                    v21 := []int{127, (-(fi18)), (-(fi18))}
                    var v22 int64 = 1
                    _ = len(v21)
                    _ = v21
                    _ = v22
                }
                _ = fi18
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS2{F0: v1, F1: (^(v16)), F2: v1}
                var v25 int32 = v16
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi17
        }
    default:
        for fi26 := 0; fi26 < 12; fi26++ {
            switch 0 {
            case -68:
                var v27 int64 = -1
                _ = v27
            case 128:
                v28 := []int{}
                _ = v28
            default:
                v29 := GenS3{F0: true, F1: 9223372036854775807}
                _ = v29
            }
            switch "" {
            case "hello":
                fmt.Println(v15, v16, v15)
            default:
                for fi30 := 0; fi30 < 7; fi30++ {
                    v31 := GenS4{F0: (v5 - v5), F1: 42}
                    _ = fi30
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := []int{fi26, (fi26 - -9223372036854775808), (fi32 % ((int(v1)) | 1)), -554}
                switch (fi32 >> ((1) & 63)) {
                case 510:
                    fmt.Println(fi26, v16, v1)
                case -1:
                    var v34 int = fi32
                    _ = v34
                case -9223372036854775808:
                    v15 = v15
                default:
                    if (-2.25 == (1.5 - (1.5 + float32(((((v15 / (v15 * 3.14159)) - v15) * v15) + v15))))) {
                        var v35 int8 = (v2 ^ (v2 / ((-1) | 1)))
                        _ = v35
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            switch (fi32 & fi32) {
                            case -184:
                                v36 := GenS4{F0: (v5 << ((746) & 63)), F1: 130}
                                _ = v36
                            }
                        case "a":
                            v1 -= v1
                        }
                    }
                }
                fmt.Println(v15, v15)
                v33 = append(v33, int(v16))
                _ = fi32
                _ = v33
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

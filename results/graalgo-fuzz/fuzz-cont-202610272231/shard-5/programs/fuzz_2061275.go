package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int64; F2 int32 }
type GenS5 struct { F0 int16; F1 uint8; F2 int }
func genFn0(p0 uint8, p1 int) bool {
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 += p0
        _ = fi6
    }
    v7 := GenS1{F0: p0, F1: 225, F2: -2147483648}
    var v8 int32 = 103
    var v9 uint = 255
    var v10 int16 = 520
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p1 < (p1 + (p1 | (p1 - (p1 * 298)))))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 4294967295
    v5 := make(map[string]uint32)
    if true {
        var v11 int8 = 100
        v12 := GenS2{F0: 136}
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 1}
    var v14 float64 = float64(1e10)
    v2 = v2
    if true {
        if (503 <= 42) {
            v15 := v1[0]
            if ((false || (!false)) && true) {
                v2 -= (v2 % ((v2) | 1))
                var v16 uint8 = v2
                _ = v16
            }
            for fi17 := 0; fi17 < 13; fi17++ {
                v1[(fi17 | (fi17 | (fi17 >> ((4294967295) & 63))))] = ""
                var v18 bool = false
                var v19 bool = (255 > 0)
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v15
        } else {
            v20 := GenS3{F0: 65535}
            _ = v20
        }
        var v21 float64 = ((float64(1e10) / (((v14 - ((1e10 * v14) / -2.25)) + v14) - -2.25)) / float64(((1e10 - ((1e10 - (0.1 + (0.1 + ((1e10 - (0.1 / (((3.14159 * float32(v14)) / 1.0) + -2.25))) - -2.25)))) - (float32(float64(3.14159)) + 0.0))) / (0.1 / 0.1))))
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int32 = 107
        for fi24 := 0; fi24 < 12; fi24++ {
            switch 42 {
            case -2147483648:
                switch 685 {
                case 4294967295:
                    v4 += v4
                case -1:
                    switch (((("a" + "") + "hello") + ("\n\t" + ("hello" + (("a" + "hello") + (("\n\t" + "foo bar") + "hello"))))) + "") {
                    case "\n\t":
                        v25 := []int{fi22, fi22}
                        _ = v25
                    case "foo bar":
                        v26 := v1[(-(fi24))]
                        _ = v26
                    case "hello":
                        v2 += 0
                    }
                case 0:
                    v27 := []int{(fi24 / (((fi22 + 0)) | 1)), (fi22 & -137), (-(fi22)), fi24}
                    _ = v27
                }
            case 0:
                var v28 int64 = -9223372036854775808
                _ = v28
            default:
                v14 += ((v14 * 0.1) / 1.0)
            }
            fmt.Println(v4)
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    switch 1 {
    case -9223372036854775808:
        if (!(2 < 2)) {
            v2 = v2
        } else {
            var v29 uint16 = 899
            v30 := GenS1{F0: (v2 * v2), F1: -9223372036854775808, F2: 0}
            _ = len(v5)
            _ = v29
            _ = v30
        }
    default:
        fmt.Println(v2, v14)
    }
    v4 = 717
    var v31 string = "\n\t"
    v32 := GenS2{F0: uint64(v4)}
    var v33 uint64 = 440
    var v34 int8 = 113
    v31 = (((v31 + "\n\t") + "\n\t") + ("a" + v31))
    for fi35 := 0; fi35 < 4; fi35++ {
        switch (fi35 + (fi35 % ((fi35) | 1))) {
        case 1:
            var v36 int16 = 638
            _ = v36
        case 993:
            for fi37 := 0; fi37 < 20; fi37++ {
                v34 += v34
                if true {
                    var v38 float32 = ((0.0 / (1.5 - ((0.0 * 1e10) / 1.0))) + (1e10 + (0.0 / (0.1 * 0.1))))
                    v39 := GenS4{F0: 18446744073709551615, F1: -613, F2: -970}
                    v40 := GenS5{F0: -566, F1: (^(v2)), F2: (fi35 * (^(fi35)))}
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi37
            }
        case 32768:
            _ = len(v3)
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}

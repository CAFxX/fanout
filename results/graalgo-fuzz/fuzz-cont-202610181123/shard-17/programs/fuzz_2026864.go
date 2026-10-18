package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int64; F2 int16 }
func genFn0(p0 uint32, p1 bool) int16 {
    p0 = (p0 % ((62) | 1))
    var v7 int64 = -1
    var v8 int32 = 0
    var v9 uint32 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 640
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-491}
    var v3 uint8 = 37
    v4 := []uint32{2}
    v5 := make(map[string]uint32)
    var v6 uint32 = 743
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 string = "a"
        v12 := GenS1{F0: -128}
        _ = fi10
        _ = v11
        _ = v12
    }
    if (false || (65535 == 2)) {
        for fi13 := 0; fi13 < 10; fi13++ {
            v3 -= v3
            var v14 uint8 = (v3 ^ v3)
            _ = fi13
            _ = v14
        }
    } else {
        if ((!true) || false) {
            if ("" < "\n\t") {
                var v15 int32 = 42
                _ = v15
            }
            if (!false) {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := GenS1{F0: 36}
                    var v18 int = 4294967295
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            } else {
                var v19 float32 = (((0.1 - ((-2.25 + (1e10 * 0.1)) + (0.1 * 0.0))) * 3.14159) * (0.0 / ((1e10 / 1.5) - 1e10)))
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := GenS2{F0: 490}
                    _ = fi20
                    _ = v21
                }
                switch -466 {
                case 128:
                    v22 := GenS1{F0: 100}
                    _ = v22
                case -900:
                    var v23 int64 = 9223372036854775807
                    _ = v23
                default:
                    v4 = append(v4, v6)
                }
                _ = v19
            }
        } else {
            v24 := GenS3{F0: -604, F1: 4294967295}
            v6 -= (v6 >> ((0) & 63))
            _ = v24
        }
        var v25 string = "foo bar"
        _ = v25
    }
    v26 := GenS4{F0: (v6 | v6), F1: (v6 * (v6 * v6)), F2: 1}
    v6 -= (v6 << ((0) & 63))
    v6 = v6
    if (42 > -77) {
        switch 317 {
        case -9223372036854775808:
            v27 := GenS1{F0: 1}
            _ = v27
        case 32767:
            v3 -= v3
        default:
            if (v3 >= v3) {
                var v28 int8 = 127
                _ = v28
            } else {
                v29 := GenS5{F0: 2}
                var v30 float64 = -2.25
                v3 = v3
                _ = v29
                _ = v30
            }
        }
    }
    switch 696 {
    case 42:
        var v31 int8 = 100
        _ = v31
    case -9223372036854775808:
        var v32 bool = true
        _ = v32
    case 741:
        for fi33 := 0; fi33 < 18; fi33++ {
            v6 = 774
            _ = fi33
        }
    default:
        if (!true) {
            v1[(v6 & 151)] = -967
            switch ("a" + ("hello" + "")) {
            case "hello":
                var v34 uint16 = 635
                _ = v34
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            fmt.Println(v6, v6)
            var v35 float64 = (1.5 - (0.0 * 3.14159))
            _ = v35
        }
    }
    switch "" {
    case "hello":
        if (!(-18 < -31)) {
            if ((!true) && true) {
                if ((false || (42 != 42)) || (1000 != -1)) {
                    v36 := GenS1{F0: 100}
                    v3 += (^(v3))
                    _ = v36
                }
                switch "hello" {
                case "a":
                    v37 := GenS6{F0: (int16(v6) >= 1), F1: 0, F2: 42}
                    _ = v37
                default:
                    v6 = 42
                }
                v6 = (v6 << ((406) & 63))
            }
        } else {
            if true {
                v3 += (v3 / (((v3 - v3)) | 1))
                v3 = 255
            } else {
                v6 += (v6 << ((917) & 63))
                if len(v4) > 0 {
                    v4[0] = v6
                }
            }
        }
    }
    v6 -= (v6 >> ((0) & 63))
    v38 := GenS6{F0: ((1e10 != float32(3.14159)) || (2 < 18446744073709551615)), F1: 42, F2: -624}
    var v39 int = -105
    var v40 uint64 = 1
    v39 += v39
    var v41 int8 = 88
    var v42 uint8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}

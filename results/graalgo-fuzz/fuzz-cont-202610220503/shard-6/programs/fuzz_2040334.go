package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 int32 = -973
    v5 := []int{2147483647, -410, -616, -1}
    v6 := []string{""}
    v7 := GenS1{F0: 42}
    v4 -= v4
    switch -1 {
    case 9223372036854775807:
        v8 := GenS1{F0: -415}
        _ = v8
    case -1:
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 = -1
                v2[32767] = ""
                _ = fi10
            }
            _ = fi9
        }
    }
    if (1 == 4294967295) {
        if true {
            v4 = (-(v4))
            var v11 uint64 = 42
            var v12 string = "a"
            _ = v11
            _ = v12
        } else {
            fmt.Println(v4)
            v13 := v2[434]
            _ = v13
        }
        v4 = v4
    }
    v4 += 42
    var v14 float64 = ((-2.25 - 0.1) * 1.5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 9223372036854775807 {
    case 4294967295:
        v4 += (v4 << ((29) & 63))
    case -9223372036854775808:
        if false {
            v4 = (v4 + 1)
            var v15 uint8 = 42
            var v16 float32 = 0.1
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                if true {
                    v4 -= v4
                }
                _ = fi17
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                if (9223372036854775807 <= 42) {
                    v14 -= float64(3.14159)
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 uint = 0
                        v21 := GenS2{F0: -961}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                } else {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        delete(v1, 2147483647)
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 bool = (!(865 == 18446744073709551615))
                        v14 = (float64(0.1) + v14)
                        var v25 float64 = v14
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                v4 = -2147483648
                _ = fi18
            }
        }
    case 42:
        v14 = 1.5
    default:
        switch "a" {
        case "":
            var v26 int64 = -822
            _ = v26
        case "foo bar":
            v4 = (v4 % ((v4) | 1))
        case "hello":
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
        default:
            switch 142 {
            case 9223372036854775807:
                var v27 uint16 = 255
                _ = v27
            }
        }
    }
    var v28 int8 = 0
    v28 = (v28 * 1)
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            v4 = 346
            fmt.Println(v14, v14, v14)
            _ = fi30
        }
        var v31 int32 = (v4 * (v4 + -1))
        switch "foo bar" {
        case "\n\t":
            if (353 < 173) {
                fmt.Println(v4, v4, v4)
                v4 -= 0
                v32 := GenS2{F0: 516}
                _ = v32
            } else {
                switch fi29 {
                case -9223372036854775808:
                    var v33 uint16 = 42
                    _ = v33
                case -2147483648:
                    v5 = append(v5, (fi29 << ((4294967295) & 63)))
                default:
                    v3[(fi29 & (fi29 & 127))] = 643
                }
            }
        case "foo bar":
            v31 = (v31 * 1)
        }
        _ = fi29
        _ = v31
    }
    v28 = (v28 & v28)
    v14 = float64((float32((v14 - (v14 / v14))) + (float32(float64(float32(((float64((1e10 - (float32(0.1) + 0.0))) - v14) - 1e10)))) - ((1e10 + 3.14159) * 1e10))))
    fmt.Println(v4)
    if (175 > 4294967295) {
        v28 = v28
        v4 = (v4 * (v4 - (v4 - -303)))
        v34 := GenS3{F0: v28, F1: 457, F2: -100}
        _ = v34
    }
    fmt.Println(v14)
    if (v4 < (v4 ^ v4)) {
        var v35 uint32 = 821
        var v36 float32 = 1e10
        var v37 int8 = v28
        _ = v35
        _ = v36
        _ = v37
    } else {
        v2[127] = ""
        if ((((float64(float32((v14 * (v14 + (v14 * 1.5))))) - v14) * v14) / v14) >= 1.5) {
            var v38 string = "hello"
            fmt.Println(v4, v4)
            _ = v38
        }
        v14 = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

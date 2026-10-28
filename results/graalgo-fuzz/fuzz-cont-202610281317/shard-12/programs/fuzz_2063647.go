package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    if ((false || true) || (1.5 < 0.1)) {
        if true {
            fmt.Println(v1, v1)
        } else {
            switch "" {
            case "\n\t":
                v1 += (v1 & v1)
            default:
                var v5 bool = (((-1 != 42) || (!false)) && (!false))
                _ = v5
            }
        }
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 string = ("foo bar" + ("a" + "\n\t"))
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: 42, F1: "foo bar", F2: -32768}
                v11 := GenS1{F0: 127, F1: "\n\t", F2: 1000}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: -9223372036854775808}
            if true {
                fmt.Println(fi12)
                v14 := v3["foo bar"]
                var v15 float64 = (1e10 / 1e10)
                _ = v14
                _ = v15
            } else {
                v16 := []int{(^(fi12)), fi12, (fi12 / ((906) | 1))}
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        if true {
            var v17 string = "a"
            var v18 int64 = 9223372036854775807
            v4[(v1 & v1)] = 841
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: 80, F1: "", F2: -1}
            fmt.Println(v1, v1)
            var v20 float64 = 0.1
            _ = v19
            _ = v20
        }
    }
    var v21 int32 = 0
    var v22 uint64 = uint64(v2)
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            v1 = v1
            if false {
                if true {
                    switch fi24 {
                    case 32768:
                        if (-2.25 >= (1.5 * ((((0.1 + (0.0 - 1.5)) - (-2.25 + (0.1 - 1e10))) + (0.1 + -2.25)) - 0.1))) {
                            v25 := v4[0]
                            var v26 uint = 2
                            _ = v25
                            _ = v26
                        }
                    case -984:
                        v27 := GenS2{F0: -9223372036854775808}
                        _ = v27
                    case 127:
                        var v28 float64 = ((1e10 * 3.14159) * (0.1 - -2.25))
                        _ = v28
                    }
                } else {
                    if (v22 <= v22) {
                        var v29 int8 = -128
                        var v30 uint8 = v2
                        v31 := GenS3{F0: v1, F1: 742, F2: (v21 >> ((42) & 63))}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        var v32 float32 = 0.0
                        _ = v32
                    }
                }
            }
            v22 -= (^(v22))
            v2 = 0
            v2 = (v2 * 128)
            if (false || true) {
                delete(v4, v1)
                v33 := []int{fi24, fi24}
                v1 -= (v1 + v1)
                _ = v33
            }
            var v34 float32 = 3.14159
            _ = fi24
            _ = v34
        }
        v35 := v4[2147483647]
        _ = fi23
        _ = v35
    }
    var v36 float32 = 0.0
    if true {
        var v37 uint16 = 255
        v22 += uint64(v37)
        v38 := GenS1{F0: 1, F1: "\n\t", F2: 0}
        _ = v37
        _ = v38
    } else {
        v1 += v1
    }
    switch "foo bar" {
    case "hello":
        v39 := []int{-9223372036854775808, -528, 32768}
        _ = v39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        switch ("foo bar" + "\n\t") {
        case "hello":
            fmt.Println(fi40, v2)
        case "\n\t":
            v41 := GenS4{F0: (1e10 < (1e10 * (((1e10 / (3.14159 - ((((3.14159 - float64(v36)) + float64(v36)) - 0.1) + 3.14159))) * (1e10 - (((-2.25 + (float64(((float32(1e10) - -2.25) * v36)) / 1.5)) - (-2.25 + -2.25)) / 1.5))) + float64((((v36 / v36) + v36) - (v36 / v36)))))), F1: v21}
            _ = v41
        default:
            v1 = v1
        }
        var v42 float64 = (0.0 * ((-2.25 + 1e10) + (((0.0 / 1e10) / float64((float32((((-2.25 / -2.25) + 0.1) + -2.25)) - 1e10))) + float64(float32((-2.25 - (1e10 * 0.0)))))))
        _ = fi40
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        v44 := []int{}
        for fi45 := 0; fi45 < 8; fi45++ {
            v1 -= v1
            _ = fi45
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 float64 = -2.25
            v47 = 0.1
            _ = fi46
            _ = v47
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v4
}

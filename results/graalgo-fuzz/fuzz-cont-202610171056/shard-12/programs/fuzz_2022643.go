package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int64 }
func genFn0() uint16 {
    var v5 float64 = 0.0
    v5 += v5
    fmt.Println(v5)
    _ = v5
    return 42
}
func genFn1(p0 int32, p1 int16, p2 uint8) int16 {
    if (458 > 42) {
        v6 := []int{1, -1, -1, -721}
        _ = v6
    }
    fmt.Println(p2, p1, p2)
    var v7 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    if (4294967295 >= -77) {
        v8 := GenS1{F0: 0, F1: 676}
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 float32 = (((float32(((((float64(-2.25) + 0.1) / 0.1) + ((-2.25 / (1.5 / 1.0)) + (3.14159 - -2.25))) - -2.25)) - float32(0.1)) + ((0.1 * (1.5 + (0.1 * (0.0 + (-2.25 / 0.1))))) + 0.1)) * (-2.25 / (((0.1 - (0.1 + (0.0 + 1e10))) - (3.14159 * 1.5)) / 3.14159)))
            switch (fi9 % ((-832) | 1)) {
            case 0:
                var v11 uint32 = 65535
                _ = v11
            }
            switch ("hello" + "foo bar") {
            case "":
                v12 := GenS1{F0: 1, F1: 65535}
                _ = v12
            case "foo bar":
                v13 := v4[881]
                _ = v13
            case "hello":
                v14 := []int{fi9}
                _ = v14
            default:
                if true {
                    fmt.Println(v10, v10)
                }
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v15 uint = 4294967295
    fmt.Println(v2, v15)
    fmt.Println(v15, v2)
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v2, fi16)
        _ = fi16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int64 = 1
        if v2 {
            for fi19 := 0; fi19 < 13; fi19++ {
                if v2 {
                    v18 += (v18 >> (((v15 % (((^(v15))) | 1))) & 63))
                    v20 := GenS1{F0: 255, F1: 558}
                    _ = v20
                }
                _ = fi19
            }
            switch fi17 {
            case 744:
                var v21 float32 = (1.5 + 3.14159)
                _ = v21
            case 1:
                var v22 int64 = v18
                _ = v22
            case 0:
                v23 := GenS1{F0: 42, F1: (^(v15))}
                _ = v23
            default:
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 string = "hello"
                    v25 = v25
                    if v2 {
                        delete(v3, "")
                    }
                    _ = fi24
                    _ = v25
                }
            }
            _ = len(v1)
        } else {
            v26 := GenS2{F0: v18}
            _ = v26
        }
        switch "\n\t" {
        case "hello":
            v18 = v18
        case "a":
            v18 = (v18 - (v18 / ((v18) | 1)))
        default:
            switch "" {
            case "foo bar":
                if (0 < 0) {
                    if (!(1.5 > ((1.5 + (1.5 * (0.0 + -2.25))) * 3.14159))) {
                        var v27 uint8 = 78
                        _ = v27
                    } else {
                        v28 := v1[1]
                        _ = v28
                    }
                    v2 = (0.0 == 0.0)
                }
            default:
                v29 := GenS2{F0: v18}
                _ = v29
            }
        }
        _ = fi17
        _ = v18
    }
    var v30 uint32 = 4294967295
    switch -9223372036854775808 {
    case -9223372036854775808:
        v31 := v1[(v30 % ((4294967295) | 1))]
        _ = v31
    case 1:
        var v32 string = "\n\t"
        _ = v32
    }
    _ = len(v4)
    v2 = v2
    v33 := []int{-9223372036854775808, 9223372036854775807}
    v15 = 319
    switch "" {
    case "":
        fmt.Println(v2)
    default:
        _ = len(v33)
    }
    _ = len(v1)
    for fi34 := 0; fi34 < 13; fi34++ {
        if len(v33) > 0 {
            v33[0] = int(v15)
        }
        _ = fi34
    }
    if (v2 || v2) {
        if (false && (130 >= 98)) {
            var v35 int64 = 42
            _ = v35
        }
        var v36 uint8 = 1
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}

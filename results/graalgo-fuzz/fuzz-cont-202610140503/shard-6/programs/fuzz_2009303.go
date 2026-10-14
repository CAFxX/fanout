package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0() bool {
    v5 := GenS1{F0: 0, F1: 4294967295, F2: 4294967295}
    var v6 int64 = 0
    v7 := []int{}
    if (469 < 4294967295) {
        v6 = (^(v6))
    }
    var v8 bool = false
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 0
    var v2 float32 = -2.25
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v9 bool = (true || (42 == 695))
    var v10 int8 = 122
    if (v9 || true) {
        if (945 == 65535) {
            v11 := GenS2{F0: (-(v10)), F1: 4294967295}
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 int64 = 1
                var v14 float64 = (0.1 + (float64((float32(1e10) - (v2 / v2))) * (((float64(((v2 * (float32(0.0) / (v2 / float32(float64((v2 * v2)))))) + 1e10)) / float64((v2 * (v2 - 0.1)))) * 1.5) * 3.14159)))
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (!v9) {
                        switch 32767 {
                        case -1:
                            var v16 int16 = -32768
                            _ = v16
                        case 32767:
                            if (42 < 255) {
                                var v17 string = "\n\t"
                                switch ("foo bar" + v17) {
                                case "foo bar":
                                    v18 := GenS2{F0: 87, F1: 42}
                                    _ = v18
                                case "\n\t":
                                    var v19 uint8 = (v1 | v1)
                                    _ = v19
                                default:
                                    v20 := GenS2{F0: (v10 | v10), F1: 214}
                                    _ = v20
                                }
                                v21 := GenS1{F0: 4294967295, F1: 923, F2: 901}
                                _ = v17
                                _ = v21
                            }
                        case 1:
                            fmt.Println(v2, fi15, v13)
                        default:
                            v13 = v13
                        }
                    }
                    delete(v3, "\n\t")
                    fmt.Println(v2)
                    _ = fi15
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v22 float32 = 3.14159
            _ = v11
            _ = v22
        } else {
            switch "" {
            case "\n\t":
                var v23 int64 = -952
                _ = v23
            case "":
                v24 := []int{128, 563, 4294967295, -982}
                _ = v24
            }
        }
        fmt.Println(v9, v2)
        v4[(("foo bar" + "\n\t") + "a")] = 42
    }
    switch -609 {
    case 730:
        v25 := GenS2{F0: 0, F1: 2}
        _ = v25
    case -2147483648:
        v26 := GenS3{F0: -152}
        _ = v26
    case 32767:
        switch 32767 {
        case -265:
            v27 := []int{}
            _ = v27
        case -536:
            v28 := GenS3{F0: 773}
            _ = v28
        case -694:
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 uint32 = 2147483648
                var v31 int8 = -23
                _ = fi29
                _ = v30
                _ = v31
            }
        default:
            v1 -= 255
        }
    default:
        var v32 uint16 = 743
        _ = v32
    }
    switch 9223372036854775807 {
    case -498:
        var v33 int32 = 467
        _ = v33
    case -2147483648:
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := v4["foo bar"]
            _ = fi34
            _ = v35
        }
    case -272:
        v36 := GenS4{F0: 4294967295, F1: (!(3 > 637)), F2: (((!v9) && (-52 == -87)) || (!(!false)))}
        _ = v36
    default:
        _ = len(v4)
    }
    if ("\n\t" > "foo bar") {
        v2 = v2
        fmt.Println(v9)
    }
    v37 := []int{}
    v1 = 1
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    _ = len(v37)
    v38 := GenS1{F0: 568, F1: 429, F2: 1}
    v39 := GenS3{F0: 2147483647}
    v40 := GenS5{F0: (v1 - v1), F1: v9}
    _ = len(v3)
    for fi41 := 0; fi41 < 17; fi41++ {
        v2 = (((float32(1.5) - 3.14159) / 1e10) - 3.14159)
        _ = fi41
    }
    v2 -= v2
    v10 += (v10 & (v10 + v10))
    v10 = (-(v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v9
}

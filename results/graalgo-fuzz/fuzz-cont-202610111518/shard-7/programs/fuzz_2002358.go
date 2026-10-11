package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0(p0 uint8, p1 string, p2 uint32) string {
    var v7 uint32 = 420
    v8 := GenS1{F0: 0, F1: v7, F2: p1}
    for fi9 := 0; fi9 < 2; fi9++ {
        if (733 == fi9) {
            switch ("a" + (p1 + p1)) {
            case "hello":
                var v10 int = -1
                _ = v10
            default:
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := []int{(fi9 * fi11)}
                    fmt.Println(p2)
                    p0 = (p0 * p0)
                    _ = fi11
                    _ = v12
                }
            }
        }
        _ = fi9
    }
    var v13 int32 = 0
    var v14 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    _ = v8
    return (p1 + "")
}

func main() {
    var v1 int64 = -102
    var v2 uint16 = 291
    v3 := make(map[uint32]int)
    var v4 uint16 = v2
    v5 := make(map[string]bool)
    var v6 uint8 = 127
    v15 := []int{}
    var v16 float64 = 1.5
    for fi17 := 0; fi17 < 4; fi17++ {
        switch "" {
        case "a":
            var v18 float64 = v16
            _ = v18
        case "hello":
            switch "\n\t" {
            case "":
                var v19 float64 = (v16 - (float64(-2.25) - v16))
                _ = v19
            case "\n\t":
                v20 := GenS1{F0: -2147483648, F1: 2147483648, F2: "\n\t"}
                _ = v20
            }
        case "":
            if ((v6 + 29) == v6) {
                for fi21 := 0; fi21 < 9; fi21++ {
                    switch (^(fi17)) {
                    case 128:
                        _ = len(v3)
                    case -1:
                        v22 := GenS1{F0: 917, F1: 861, F2: ("hello" + (("foo bar" + "a") + "a"))}
                        _ = v22
                    }
                    var v23 uint = 65535
                    _ = fi21
                    _ = v23
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    switch "" {
                    case "hello":
                        v16 = (0.1 - 0.0)
                    case "a":
                        v2 = (v4 - (v2 - v4))
                    case "foo bar":
                        switch fi17 {
                        case -674:
                            v16 = float64((((-2.25 - 1.5) / 1e10) * (0.0 * 1.5)))
                        case 636:
                            v25 := GenS1{F0: 262, F1: 2147483647, F2: "foo bar"}
                            var v26 int8 = -1
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            if (!(3.14159 < 0.1)) {
                                fmt.Println(v16, fi24, v16)
                                switch fi24 {
                                case 15:
                                    v26 = (v26 % (((v26 * (v26 | (v26 << ((42) & 63))))) | 1))
                                default:
                                    switch (fi17 & (^(fi17))) {
                                    case 128:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v6 += 127
                                        default:
                                            v6 += v6
                                        }
                                    case 32767:
                                        var v27 uint8 = (v6 / (((v6 * v6)) | 1))
                                        _ = v27
                                    }
                                }
                                v1 = int64(v26)
                            }
                            _ = v25
                            _ = v26
                        case 42:
                            _ = len(v15)
                        default:
                            var v28 int8 = 100
                            _ = v28
                        }
                    }
                    v29 := GenS2{F0: 65535, F1: v6}
                    var v30 float64 = (3.14159 * (float64(((-2.25 - 0.0) - 0.1)) + v16))
                    _ = fi24
                    _ = v29
                    _ = v30
                }
            } else {
                v31 := GenS3{F0: 42, F1: -1, F2: 835}
                _ = v31
            }
        default:
            var v32 float64 = ((-2.25 - (float64((1e10 - (1e10 + 0.1))) - v16)) - v16)
            _ = v32
        }
        var v33 uint16 = uint16(v6)
        _ = fi17
        _ = v33
    }
    if len(v15) > 0 {
        v15[0] = -1
    }
    switch int(v6) {
    case 1:
        v34 := GenS4{F0: "a", F1: 127, F2: (v2 + (v4 | (v2 ^ v2)))}
        _ = v34
    default:
        var v35 uint8 = (v6 >> ((4294967295) & 63))
        _ = v35
    }
    v6 -= (v6 >> ((uint(v4)) & 63))
    var v36 uint32 = 42
    v37 := GenS5{F0: 549, F1: 4294967295, F2: 4294967295}
    var v38 float64 = 3.14159
    v4 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

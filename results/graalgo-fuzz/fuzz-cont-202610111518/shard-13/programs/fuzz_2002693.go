package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint32 }
func genFn0(p0 uint64) int64 {
    v4 := GenS1{F0: 65535, F1: 1000}
    if (2 > 1) {
        fmt.Println(p0)
    } else {
        v5 := []int{-56, 128, 127}
        _ = v5
    }
    var v6 int = -9223372036854775808
    v7 := []int{v6, v6, v6}
    _ = p0
    _ = v4
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 330
    v1 = ((((((float64((1.5 + (((3.14159 - 3.14159) + (0.1 / (0.1 - -2.25))) - 1e10))) - ((v1 + 3.14159) - ((v1 * v1) - v1))) + (v1 + v1)) / 1.5) + v1) / 0.1) - v1)
    v1 = float64(((1.5 + (1.5 / ((((float32(1.5) + 0.0) / float32(1.5)) + (1e10 - 0.0)) / 3.14159))) / 1.0))
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            v3 -= (v3 & (v3 ^ 93))
        }
        _ = fi8
    }
    var v9 int16 = -32768
    var v10 int8 = -1
    v3 += (v3 / ((v3) | 1))
    v11 := GenS1{F0: 2, F1: 32767}
    if false {
        v3 = (v3 | v3)
        for fi12 := 0; fi12 < 9; fi12++ {
            switch "\n\t" {
            case "":
                v13 := GenS2{F0: (!false), F1: 255, F2: ((v10 & (^(v10))) > (v10 / ((v10) | 1)))}
                _ = v13
            case "a":
                var v14 uint32 = 0
                _ = v14
            default:
                if (false || ("\n\t" > "foo bar")) {
                    v3 += 32767
                    v15 := GenS3{F0: int(v10), F1: -92}
                    _ = v15
                } else {
                    switch fi12 {
                    case -617:
                        v9 += 1
                    default:
                        v10 = -128
                    }
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v10 -= (v10 << ((2) & 63))
                        v10 = (v10 << ((2) & 63))
                        _ = fi16
                    }
                }
            }
            switch (fi12 & -1) {
            case 0:
                var v17 int64 = -9223372036854775808
                _ = v17
            }
            var v18 int = -9223372036854775808
            _ = fi12
            _ = v18
        }
        v19 := GenS1{F0: 2147483647, F1: (v9 * -32768)}
        _ = v19
    }
    switch -9223372036854775808 {
    case -192:
        switch -1 {
        case -256:
            if ("hello" >= ("a" + ("" + "a"))) {
                switch "a" {
                case "a":
                    if (v3 < v3) {
                        v20 := GenS3{F0: 837, F1: v10}
                        _ = v20
                    } else {
                        v21 := GenS2{F0: (float64((1e10 - 1.5)) > v1), F1: (v3 - v3), F2: (false || false)}
                        var v22 int8 = v10
                        _ = v21
                        _ = v22
                    }
                case "hello":
                    var v23 float32 = 1.5
                    _ = v23
                case "":
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if (!false) {
                            switch "\n\t" {
                            case "foo bar":
                                v1 = (v1 / 1.0)
                            case "hello":
                                var v25 uint = 0
                                _ = v25
                            }
                            v26 := []int{(fi24 % ((602) | 1)), (fi24 % (((fi24 ^ 128)) | 1)), fi24, fi24}
                            var v27 int16 = (v9 & (v9 - v9))
                            _ = v26
                            _ = v27
                        }
                        _ = fi24
                    }
                }
            } else {
                v10 -= (v10 >> ((65535) & 63))
            }
        case 32767:
            v1 -= 3.14159
        }
    case -1:
        var v28 uint = 255
        _ = v28
    case 775:
        v29 := GenS4{F0: 65535, F1: 2147483647, F2: 1}
        _ = v29
    }
    var v30 int = 4294967295
    var v31 bool = false
    switch "" {
    case "":
        if false {
            if true {
                var v32 uint32 = 0
                _ = v32
            } else {
                var v33 bool = v31
                switch ("hello" + ("foo bar" + ("\n\t" + "hello"))) {
                case "":
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v35 := []int{}
                        for fi36 := 0; fi36 < 9; fi36++ {
                            v33 = (!(fi36 <= 4294967295))
                            v37 := GenS5{F0: true, F1: -9223372036854775808}
                            v3 = v3
                            v38 := v2[927]
                            fmt.Println(v38, fi36, v30)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = fi34
                        _ = v35
                    }
                case "\n\t":
                    v39 := GenS3{F0: (v30 >> ((0) & 63)), F1: v10}
                    _ = v39
                }
                _ = v33
            }
            var v40 uint64 = 819
            v9 -= (v9 * 1000)
            _ = v40
        }
    case "\n\t":
        v41 := GenS6{F0: 4294967295, F1: 1, F2: 4294967295}
        _ = v41
    default:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v9
}

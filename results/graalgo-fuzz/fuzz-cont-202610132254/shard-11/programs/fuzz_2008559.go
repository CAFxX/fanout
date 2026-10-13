package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
type GenS4 struct { F0 uint16; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 uint = 42
    var v4 uint64 = 18446744073709551615
    var v5 int16 = 0
    if (9223372036854775807 == 11) {
        var v6 uint8 = 152
        _ = v6
    }
    if (0 > 255) {
        var v7 uint16 = 129
        _ = v7
    } else {
        v5 = v5
        switch "a" {
        case "\n\t":
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS1{F0: 0, F1: 0, F2: v3}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                if true {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 float64 = (0.0 - 0.1)
                        if false {
                            var v14 float64 = (v13 - v13)
                            var v15 uint16 = 32767
                            v14 += v14
                            _ = v14
                            _ = v15
                        } else {
                            var v16 float64 = 1.5
                            v17 := []int{353, 32767, -545, (fi11 % ((int(v3)) | 1))}
                            v13 += v16
                            if len(v17) > 0 {
                                v17[0] = 9223372036854775807
                            }
                            _ = v16
                            _ = v17
                        }
                        _ = fi12
                        _ = v13
                    }
                } else {
                    if true {
                        v4 += (v4 | 42)
                        fmt.Println(v5)
                    }
                    v18 := GenS1{F0: 423, F1: 325, F2: v3}
                    fmt.Println(v3, v5, v5)
                    _ = v18
                }
                v19 := GenS1{F0: 255, F1: 197, F2: (v3 | v3)}
                var v20 uint8 = 21
                _ = fi11
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: 42, F1: 2, F2: (v3 >> (((^(v3))) & 63))}
    var v22 uint32 = 558
    var v23 uint32 = (^(v22))
    v5 = 1
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS1{F0: 42, F1: v23, F2: v3}
        v2["\n\t"] = ""
        v26 := GenS1{F0: 678, F1: 0, F2: 715}
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        if (!((int(v4) < (^(fi27))) || false)) {
            fmt.Println(v4, v3, v23)
        }
        var v28 uint64 = (v4 + v4)
        if true {
            var v29 int = 128
            switch ("hello" + "foo bar") {
            case "a":
                switch (fi27 >> (((v3 * v3)) & 63)) {
                case 32768:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        delete(v2, "a")
                        v31 := GenS1{F0: 32767, F1: (v23 - (v22 ^ v23)), F2: v3}
                        _ = fi30
                        _ = v31
                    }
                }
            case "foo bar":
                v32 := GenS1{F0: 394, F1: (v23 % ((v23) | 1)), F2: v3}
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 int = v29
                    _ = len(v1)
                    if false {
                        v35 := GenS2{F0: (v4 * uint64(v22)), F1: "\n\t"}
                        _ = v35
                    } else {
                        var v36 int32 = -5
                        _ = v36
                    }
                    switch "\n\t" {
                    case "hello":
                        var v37 uint64 = v4
                        fmt.Println(fi33, v23)
                        _ = v37
                    case "a":
                        v5 = 32767
                    default:
                        fmt.Println(v3, v23)
                        if false {
                            v28 = (v28 % ((v4) | 1))
                            switch 1 {
                            case 127:
                                v3 -= (^(v3))
                            }
                            switch (fi33 + fi33) {
                            case 123:
                                var v38 uint32 = (v23 % ((v22) | 1))
                                _ = v38
                            default:
                                v39 := []int{(fi33 ^ (v29 >> (((v3 - v3)) & 63))), (fi33 / ((v34) | 1)), (^(fi33))}
                                _ = v39
                            }
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            }
            fmt.Println(v22)
            _ = v29
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                switch (fi27 / ((int(v28)) | 1)) {
                case 32767:
                    if true {
                        v41 := GenS1{F0: 340, F1: v23, F2: (v3 << ((4294967295) & 63))}
                        v28 += (^(v4))
                        _ = v41
                    } else {
                        switch fi40 {
                        case -1:
                            if (false && ((!true) && (1 < 0))) {
                                var v42 uint32 = v23
                                v28 -= 2
                                _ = v42
                            }
                        case -526:
                            _ = len(v2)
                        default:
                            v5 = v5
                        }
                    }
                }
                v23 = 2
                var v43 float32 = 0.0
                _ = fi40
                _ = v43
            }
            if (!(int16(v4) <= (v5 | v5))) {
                for fi44 := 0; fi44 < 18; fi44++ {
                    v45 := GenS3{F0: v28, F1: 128, F2: fi27}
                    var v46 uint8 = 160
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                v47 := GenS2{F0: (v28 % ((42) | 1)), F1: "foo bar"}
                if (-1 < -13) {
                    v5 += v5
                }
                _ = v47
            }
            var v48 uint16 = 1
            _ = v48
        }
        _ = fi27
        _ = v28
    }
    v49 := GenS1{F0: 559, F1: v23, F2: (v3 >> ((v3) & 63))}
    v5 -= 32767
    v50 := GenS4{F0: 255, F1: "a", F2: (^(v22))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}

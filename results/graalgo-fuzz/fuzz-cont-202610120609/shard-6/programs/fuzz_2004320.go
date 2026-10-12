package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 int32; F1 string; F2 int32 }

func main() {
    var v1 int32 = 900
    var v2 float64 = 1e10
    var v3 uint16 = 2
    if true {
        if false {
            for fi4 := 0; fi4 < 4; fi4++ {
                for fi5 := 0; fi5 < 19; fi5++ {
                    v6 := GenS1{F0: "a"}
                    v7 := GenS2{F0: 2147483647, F1: 4294967295, F2: int16(fi4)}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                v8 := GenS3{F0: 308, F1: v1}
                _ = fi4
                _ = v8
            }
            v1 -= 1
            v9 := []int{-723, 4294967295}
            _ = v9
        } else {
            v2 = v2
        }
        v10 := []int{-920}
        for fi11 := 0; fi11 < 4; fi11++ {
            v1 = -363
            _ = fi11
        }
        _ = v10
    }
    var v12 int16 = -32768
    var v13 float64 = 0.1
    if true {
        switch 1 {
        case 42:
            v12 = -32768
        case 1:
            switch "" {
            case "hello":
                v14 := []int{32767}
                _ = v14
            case "\n\t":
                switch 546 {
                case 72:
                    v1 = 42
                }
            }
        case 9223372036854775807:
            fmt.Println(v13, v1)
        }
    } else {
        v2 = ((float64(3.14159) * (v2 - float64(1e10))) - 1e10)
        if false {
            var v15 float32 = float32(-2.25)
            var v16 int32 = v1
            if false {
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 int32 = (v16 >> ((uint(v3)) & 63))
                    if (!(false && (((1 > 485) || true) && (true || true)))) {
                        var v19 int16 = v12
                        _ = v19
                    } else {
                        if (42 != 255) {
                            v20 := []int{1}
                            _ = len(v20)
                            _ = v20
                        } else {
                            switch fi17 {
                            case 9223372036854775807:
                                var v21 float32 = v15
                                _ = v21
                            case 1:
                                v15 -= v15
                            default:
                                v22 := GenS2{F0: 65535, F1: 2, F2: v12}
                                _ = v22
                            }
                            v1 -= (v16 >> ((945) & 63))
                        }
                    }
                    v23 := GenS3{F0: 65535, F1: (v1 + (v18 % (((v1 * v1)) | 1)))}
                    _ = fi17
                    _ = v18
                    _ = v23
                }
            } else {
                var v24 int32 = v16
                var v25 int = -625
                _ = v24
                _ = v25
            }
            _ = v15
            _ = v16
        } else {
            var v26 float64 = 0.0
            _ = v26
        }
        v27 := []int{2147483647, 0}
        _ = v27
    }
    v3 = (^(v3))
    if false {
        var v28 uint = 801
        if (32767 <= 0) {
            v29 := []int{32768, 9223372036854775807, -176}
            v29 = append(v29, 32767)
            v30 := GenS2{F0: 2147483648, F1: 65535, F2: -741}
            _ = v29
            _ = v30
        }
        _ = v28
    }
    v31 := GenS4{F0: 580}
    if false {
        v32 := GenS1{F0: "foo bar"}
        var v33 int64 = 9223372036854775807
        _ = v32
        _ = v33
    } else {
        switch -296 {
        case 109:
            v34 := GenS3{F0: 65535, F1: (v1 - v1)}
            _ = v34
        case 458:
            if (471 == (-(v1))) {
                v35 := GenS5{F0: 701, F1: -9223372036854775808}
                fmt.Println(v3, v13, v13)
                _ = v35
            }
        case 127:
            v2 += v2
        }
    }
    if (!false) {
        fmt.Println(v12)
        fmt.Println(v3, v1, v3)
        v36 := GenS6{F0: (v1 * v1), F1: "", F2: (v1 % ((-1) | 1))}
        _ = v36
    }
    if (!(!false)) {
        var v37 int8 = -1
        var v38 int32 = 42
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
}

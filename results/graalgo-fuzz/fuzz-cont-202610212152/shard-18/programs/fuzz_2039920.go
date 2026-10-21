package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint64; F2 int }
func genFn0(p0 int64) uint32 {
    var v6 uint64 = 2
    switch "a" {
    case "a":
        v6 = 2
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, fi7)
        v8 := []int{}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    if (((!((true || false) && (!(-344 != 1)))) && (true && true)) && (32767 <= 32767)) {
        var v9 uint = 42
        v10 := []int{2147483647, int(v6)}
        _ = v10
        _ = v9
    }
    v11 := []int{-97}
    _ = p0
    _ = v11
    _ = v6
    return 81
}

func main() {
    var v1 string = "a"
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 int16 = 25
    v12 := GenS1{F0: ("hello" + "hello"), F1: 255, F2: 0}
    v5 += v5
    v5 = -314
    v1 = "\n\t"
    if ((!false) || true) {
        if (457 < 0) {
            switch 32767 {
            case 128:
                v13 := GenS2{F0: "foo bar", F1: 1}
                _ = v13
            case -605:
                var v14 uint8 = 230
                _ = v14
            default:
                if false {
                    v15 := GenS3{F0: 28, F1: 1}
                    _ = v15
                }
            }
            var v16 uint8 = 2
            v17 := GenS3{F0: 18446744073709551615, F1: (v16 % ((128) | 1))}
            _ = v16
            _ = v17
        }
        _ = len(v4)
    }
    if true {
        var v18 int8 = 100
        for fi19 := 0; fi19 < 14; fi19++ {
            delete(v4, 229)
            var v20 int = -757
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    if (1.5 != 1e10) {
        var v21 int32 = -408
        if false {
            v1 = "foo bar"
            var v22 uint64 = 2
            v23 := GenS4{F0: 1, F1: (v2 & (v2 ^ v2)), F2: (v5 % ((v5) | 1))}
            _ = v22
            _ = v23
        } else {
            _ = len(v4)
            if (((0 >= 2) && (!(255 < 42))) && false) {
                switch v1 {
                case "foo bar":
                    v24 := GenS2{F0: v1, F1: 0}
                    _ = v24
                case "hello":
                    for fi25 := 0; fi25 < 1; fi25++ {
                        fmt.Println(v1, v2, v5)
                        _ = fi25
                    }
                }
            }
        }
        var v26 float64 = float64(0.0)
        _ = v21
        _ = v26
    }
    var v27 int8 = 1
    for fi28 := 0; fi28 < 4; fi28++ {
        if (true || (true && false)) {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int32 = -752
                v30 += (-(v30))
                if (!true) {
                    v2 += v2
                    v31 := GenS5{F0: 462}
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            var v32 int32 = -536
            _ = v32
        }
        var v33 uint64 = 203
        _ = fi28
        _ = v33
    }
    if (18446744073709551615 >= 1) {
        v34 := GenS6{F0: (false || (v2 > v2)), F1: 42, F2: -2147483648}
        v1 = v1
        if (0 == 0) {
            switch 4294967295 {
            case 1:
                v35 := []int{-241, 1}
                _ = v35
            case 32768:
                var v36 int = 841
                _ = v36
            default:
                switch v1 {
                case "\n\t":
                    fmt.Println(v5, v2, v27)
                case "hello":
                    if false {
                        v37 := v3[75]
                        v38 := []int{42}
                        if len(v38) >= 2 {
                            v38 = v38[1:len(v38)]
                        }
                        _ = v37
                        _ = v38
                    } else {
                        v39 := []int{216, -930}
                        v5 += (^(v5))
                        fmt.Println(v1)
                        _ = v39
                    }
                case "a":
                    if false {
                        v1 = v1
                        v5 += (v5 * v5)
                        v27 += (v27 * v27)
                    } else {
                        v40 := GenS3{F0: 255, F1: 42}
                        _ = v40
                    }
                }
            }
            fmt.Println(v5, v27)
            switch ("foo bar" + v1) {
            case "hello":
                v41 := []int{}
                _ = v41
            }
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                v1 = v1
                v43 := GenS4{F0: 1, F1: 1, F2: (^(v5))}
                _ = fi42
                _ = v43
            }
            for fi44 := 0; fi44 < 13; fi44++ {
                fmt.Println(v1)
                _ = fi44
            }
        }
        _ = v34
    }
    v45 := GenS1{F0: "", F1: 1, F2: 772}
    fmt.Println(v2, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}

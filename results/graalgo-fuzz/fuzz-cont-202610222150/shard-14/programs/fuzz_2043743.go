package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint8, p2 uint8) uint {
    var v5 int16 = -32768
    v6 := GenS1{F0: true, F1: 951, F2: 274}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 271
}

func main() {
    v1 := []uint32{2147483647, 923, 0}
    var v2 int64 = -1
    var v3 uint = 65535
    v4 := []string{}
    v3 = (v3 | genFn0(v2, 0, 222))
    switch "\n\t" {
    case "\n\t":
        v1 = append(v1, 65535)
    case "a":
        var v7 uint32 = 781
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int8 = 1
            v9 += v9
            _ = fi8
            _ = v9
        }
    default:
        fmt.Println(v2)
    }
    var v10 uint = (v3 | (v3 / (((v3 << ((v3) & 63))) | 1)))
    v2 = v2
    switch "foo bar" {
    case "hello":
        fmt.Println(v3, v3)
    }
    v11 := GenS2{F0: "foo bar"}
    v12 := GenS3{F0: 4294967295}
    v3 = genFn0(v2, 127, 255)
    var v13 bool = (65535 != 255)
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            if (1e10 > (-2.25 * -2.25)) {
                fmt.Println(fi15, v13)
                fmt.Println(v13, v13)
            } else {
                v16 := GenS3{F0: 117}
                _ = v16
            }
            v17 := GenS4{F0: 0, F1: 128, F2: 2}
            if v13 {
                switch fi14 {
                case 9223372036854775807:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 68:
                    v18 := GenS1{F0: v13, F1: 846, F2: 0}
                    _ = v18
                case -522:
                    v19 := GenS2{F0: "\n\t"}
                    _ = v19
                }
                var v20 float64 = 0.1
                fmt.Println(v2, fi15, fi15)
                _ = v20
            } else {
                var v21 uint32 = 1
                _ = v21
            }
            _ = fi15
            _ = v17
        }
        v22 := GenS5{F0: 2}
        v10 += genFn0(v2, uint8(v3), 42)
        _ = fi14
        _ = v22
    }
    fmt.Println(v13)
    v23 := GenS1{F0: false, F1: 0, F2: 42}
    if (-1 != -1) {
        v3 = genFn0((v2 & 0), 1, 222)
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 13; fi25++ {
                switch 32768 {
                case 786:
                    v26 := GenS5{F0: 347}
                    _ = v26
                }
                _ = fi25
            }
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 float64 = 1e10
                _ = fi27
                _ = v28
            }
            _ = fi24
        }
    } else {
        if v13 {
            v29 := GenS4{F0: 255, F1: 70, F2: 2}
            v13 = v13
            v13 = (65535 != 2147483647)
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            switch fi30 {
            case -2147483648:
                if (-128 >= -128) {
                    var v31 float32 = 3.14159
                    switch int(v3) {
                    case 32768:
                        v10 += (v10 / ((v10) | 1))
                    }
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        var v33 bool = (4294967295 > 2)
                        fmt.Println(v3)
                        _ = fi32
                        _ = v33
                    }
                }
            case 0:
                var v34 int8 = 0
                _ = v34
            }
            _ = fi30
        }
    }
    fmt.Println(v3)
    var v35 int8 = 36
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v4
}

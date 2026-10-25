package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }
type GenS4 struct { F0 uint; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 string }
func genFn0(p0 string, p1 int64) string {
    var v4 int = 826
    v5 := []int{v4}
    v4 = 2147483647
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint64 = 18446744073709551615
    v6 := GenS1{F0: (v3 & (v3 | 18446744073709551615))}
    v1[350] = "hello"
    v7 := GenS2{F0: (^(v3)), F1: 1, F2: 0}
    var v8 int64 = 0
    fmt.Println(v3)
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 bool = true
            if true {
                v12 := GenS3{F0: 136, F1: 373, F2: (128 == 255)}
                v8 = (v8 >> ((0) & 63))
                _ = v12
            }
            v3 += v3
            _ = fi10
            _ = v11
        }
        var v13 int8 = 1
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: (^(v3)), F1: (v8 * (v8 % ((v8) | 1))), F2: 27}
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 128
    v16 += 42
    if (((v3 >> ((710) & 63)) <= 18446744073709551615) && (false || ((false || (!false)) || false))) {
        fmt.Println(v16, v3, v16)
        v3 += (v3 & v3)
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if false {
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS2{F0: (^(v3)), F1: v8, F2: 2}
                switch "a" {
                case "\n\t":
                    fmt.Println(v16, v8, v8)
                case "foo bar":
                    v20 := GenS1{F0: v3}
                    _ = v20
                default:
                    v8 += -1
                }
                switch ((("\n\t" + "foo bar") + "a") + "\n\t") {
                case "":
                    var v21 int8 = 82
                    _ = v21
                }
                if false {
                    if (v16 < (v16 - v16)) {
                        v16 = 116
                        fmt.Println(v16, fi17)
                    } else {
                        if false {
                            v8 += v8
                            var v22 uint32 = 65535
                            _ = v22
                        }
                        v3 += (v3 * 0)
                    }
                    var v23 string = "\n\t"
                    var v24 string = v23
                    var v25 bool = (305 >= 255)
                    v3 -= (v3 << ((983) & 63))
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v26 := GenS4{F0: 4294967295, F1: true, F2: v16}
                _ = fi18
                _ = v19
                _ = v26
            }
            v16 += v16
            switch fi17 {
            case -36:
                var v27 uint16 = 1
                _ = v27
            }
        }
        _ = fi17
    }
    var v28 uint = 822
    if true {
        _ = len(v2)
    } else {
        var v29 uint16 = 255
        for fi30 := 0; fi30 < 2; fi30++ {
            switch "hello" {
            case "a":
                var v31 float64 = 1e10
                _ = v31
            case "foo bar":
                v32 := GenS5{F0: uint64(v16), F1: ""}
                _ = v32
            }
            var v33 uint64 = 4294967295
            _ = fi30
            _ = v33
        }
        v34 := GenS6{F0: 42, F1: -1}
        _ = v29
        _ = v34
    }
    if (((v3 / (((v3 % (((^(v3))) | 1))) | 1)) > 1) && ((346 > 1) || ((v3 + 42) < 42))) {
        v8 += 0
        v16 -= v16
        v35 := []int{-1, -9223372036854775808, 0}
        _ = v35
    }
    v36 := GenS7{F0: "foo bar"}
    var v37 uint16 = 522
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v6
    _ = v7
    _ = v8
}

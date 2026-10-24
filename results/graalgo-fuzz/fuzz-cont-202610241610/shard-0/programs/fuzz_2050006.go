package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint }
type GenS5 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS6 struct { F0 string; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS8 struct { F0 int16; F1 int64; F2 int }
func genFn0(p0 int16, p1 uint) uint8 {
    p0 = int16(p1)
    p0 = (^(p0))
    var v5 uint64 = 355
    if (9223372036854775807 >= -865) {
        v6 := GenS1{F0: -9223372036854775808, F1: v5}
        var v7 float32 = (((1e10 * (0.0 * (0.0 + 1.5))) - 1.5) - ((1e10 / 1e10) + 0.0))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -942, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v9 := GenS2{F0: -1}
    if false {
        var v10 bool = false
        switch "a" {
        case "foo bar":
            v11 := GenS2{F0: 100}
            _ = v11
        case "a":
            v2 -= genFn0(1, 65535)
        }
        _ = v10
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        v2 = v2
        v13 := GenS3{F0: 39, F1: 2147483647, F2: true}
        _ = v13
    }
    delete(v1, -512)
    var v14 uint8 = v2
    var v15 uint32 = 2147483648
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := GenS2{F0: 0}
        _ = fi16
        _ = v17
    }
    switch (("\n\t" + "\n\t") + "hello") {
    case "hello":
        if true {
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v2)
                v14 += (v2 | 2)
                switch "" {
                case "hello":
                    v19 := GenS1{F0: -99, F1: 1}
                    _ = v19
                case "":
                    v2 -= 8
                default:
                    v20 := GenS4{F0: 2, F1: v15, F2: 18446744073709551615}
                    _ = v20
                }
                _ = fi18
            }
            if (1 == 9223372036854775807) {
                switch ("a" + "foo bar") {
                case "\n\t":
                    v21 := []int{-873, 128}
                    _ = v21
                case "hello":
                    switch 32767 {
                    case -1:
                        v2 = 255
                    case 1:
                        var v22 uint8 = 130
                        _ = v22
                    default:
                        var v23 uint = 295
                        _ = v23
                    }
                default:
                    v15 = (v15 | v15)
                }
                v14 = (v2 ^ v2)
            } else {
                _ = len(v3)
                if false {
                    v24 := GenS2{F0: -1}
                    if true {
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v15 -= (v15 ^ (^(v15)))
                            _ = fi25
                        }
                        v26 := GenS5{F0: 667, F1: 42, F2: 4294967295}
                        v27 := v1[725]
                        _ = v26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v29 := GenS6{F0: "\n\t", F1: v2, F2: int8(fi28)}
                        var v30 int8 = -78
                        switch (fi28 / (((fi28 >> ((18446744073709551615) & 63))) | 1)) {
                        case 4294967295:
                            var v31 float64 = (((1e10 / 1e10) / (1.5 - -2.25)) * 1.5)
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = v24
                }
                v32 := GenS5{F0: 1, F1: 18446744073709551615, F2: 42}
                _ = v32
            }
        }
    default:
        v1[2147483647] = -864
    }
    if ((!true) || (((0.0 / 0.1) + ((0.1 * 1e10) + (float32(1e10) * (0.0 - (0.0 * (1.5 * 0.1)))))) > (0.0 + 1.5))) {
        var v33 int64 = 0
        v2 -= (v14 & 1)
        var v34 uint8 = 143
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := GenS6{F0: "hello", F1: v2, F2: 0}
        v2 = (v14 * (v2 % (((v14 / (((v14 - (v14 / ((v2) | 1)))) | 1))) | 1)))
        for fi37 := 0; fi37 < 8; fi37++ {
            v38 := GenS7{F0: 2, F1: 620, F2: 255}
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    var v39 float64 = -2.25
    var v40 bool = true
    for fi41 := 0; fi41 < 11; fi41++ {
        v40 = (true || v40)
        if (v40 || v40) {
            if (255 <= 42) {
                var v42 uint16 = 32767
                v43 := GenS3{F0: 83, F1: 42, F2: (-520 == 42)}
                switch (fi41 << ((uint(v2)) & 63)) {
                case -2147483648:
                    var v44 uint16 = 65535
                    _ = v44
                case -119:
                    if v40 {
                        if false {
                            for fi45 := 0; fi45 < 16; fi45++ {
                                var v46 int32 = 179
                                v47 := GenS8{F0: 0, F1: 1, F2: int(v15)}
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        } else {
                            v48 := v3[(^(v15))]
                            _ = v48
                        }
                        _ = len(v3)
                    }
                }
                _ = v42
                _ = v43
            }
            var v49 int32 = 0
            v2 -= (v2 ^ 1)
            _ = v49
        }
        fmt.Println(fi41)
        _ = fi41
    }
    v50 := GenS8{F0: -482, F1: 0, F2: 0}
    if (v40 || false) {
        v51 := v1[4294967295]
        _ = v51
    } else {
        switch "hello" {
        case "a":
            var v52 int = 1
            _ = v52
        case "":
            var v53 uint64 = 706
            _ = v53
        }
        v39 = (v39 / 1e10)
    }
    v54 := GenS5{F0: 77, F1: 42, F2: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v50
    _ = v54
    _ = v9
}

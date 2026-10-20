package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int16; F2 string }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 uint16 = 0
    v5 := GenS1{F0: 187, F1: 858, F2: 0}
    v6 := []int{451, 4294967295}
    if false {
        fmt.Println(v3)
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                v4 = v3
                _ = fi8
            }
            _ = fi7
        }
    }
    v3 = v3
    switch "a" {
    case "hello":
        v9 := GenS1{F0: 175, F1: 0, F2: 506}
        _ = v9
    default:
        var v10 string = "\n\t"
        _ = v10
    }
    if (50 <= 42) {
        if false {
            if false {
                if false {
                    v4 -= (v3 * (v4 >> ((42) & 63)))
                } else {
                    v4 = 0
                    v11 := GenS2{F0: 42, F1: v4}
                    v4 = (^(v4))
                    _ = v11
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS3{F0: 48, F1: fi12, F2: "hello"}
                    _ = fi12
                    _ = v13
                }
                var v14 int64 = 42
                _ = v14
            }
            if true {
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 uint32 = 0
                    v17 := GenS4{F0: 42}
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v19 := GenS1{F0: 1, F1: 899, F2: -970}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v20 uint8 = 0
                _ = v20
            }
        } else {
            var v21 float32 = float32(((float64(-2.25) - 0.0) / 1.5))
            fmt.Println(v3, v4, v4)
            if false {
                fmt.Println(v4, v21, v3)
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch (fi22 | (fi22 * fi22)) {
                    case 265:
                        v23 := v2[-1]
                        _ = v23
                    default:
                        fmt.Println(v3, v3, fi22)
                    }
                    fmt.Println(v21, v3)
                    _ = fi22
                }
                var v24 int16 = 444
                _ = v24
            } else {
                v4 = 283
            }
            _ = v21
        }
        switch "foo bar" {
        case "":
            var v25 uint8 = 0
            _ = v25
        default:
            if true {
                var v26 uint8 = 127
                if false {
                    fmt.Println(v3)
                }
                if (1.5 < float32((-2.25 + (-2.25 * 3.14159)))) {
                    var v27 uint = 255
                    _ = v27
                }
                _ = v26
            } else {
                fmt.Println(v4)
                var v28 uint64 = 18446744073709551615
                var v29 uint = 1
                _ = v28
                _ = v29
            }
        }
    } else {
        fmt.Println(v3, v3)
    }
    if (((true || true) || true) || ((false && false) || (210 < 2))) {
        var v30 bool = false
        fmt.Println(v30)
        _ = v30
    } else {
        v3 = (^(v4))
        var v31 uint16 = v3
        if (4294967295 > 255) {
            if (!(2147483648 != 396)) {
                fmt.Println(v4, v3)
            }
        }
        _ = v31
    }
    if true {
        v6 = append(v6, 32767)
        fmt.Println(v3, v3)
    } else {
        fmt.Println(v4)
        v32 := GenS5{F0: -754, F1: 42, F2: "\n\t"}
        _ = v32
    }
    v33 := v1[498]
    var v34 int16 = 32767
    var v35 int16 = v34
    if v33 {
        v3 -= (v3 >> ((65535) & 63))
        v36 := GenS6{F0: 1}
        fmt.Println(v33, v33)
        _ = v36
    } else {
        v33 = false
        _ = len(v2)
        v34 = -485
    }
    var v37 uint64 = uint64(v3)
    v38 := GenS4{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

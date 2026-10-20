package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{32768, 0, -12}
    var v3 uint32 = 109
    var v4 float32 = (1e10 / (0.0 - float32(0.1)))
    var v5 int = 0
    var v6 uint32 = v3
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v6, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint64 = 42
        var v10 uint32 = v3
        v6 -= v3
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 110, F1: 2147483647}
    var v12 int64 = -1
    switch ("hello" + "\n\t") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v12 = v12
    case "hello":
        for fi13 := 0; fi13 < 15; fi13++ {
            if len(v2) > 0 {
                v2[0] = (v5 | (^(fi13)))
            }
            fmt.Println(v6)
            if (!false) {
                if false {
                    var v14 float64 = (-2.25 * -2.25)
                    v12 = (v12 - (v12 % ((v12) | 1)))
                    if (false || false) {
                        fmt.Println(v5)
                    }
                    _ = v14
                } else {
                    v3 -= (^(v6))
                }
                var v15 uint32 = (^(v6))
                _ = v15
            } else {
                v6 = v6
                var v16 uint8 = 242
                switch "\n\t" {
                case "":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        _ = len(v2)
                        _ = fi17
                    }
                case "hello":
                    v5 = (^(v5))
                }
                _ = v16
            }
            _ = fi13
        }
    }
    v5 += (-(v5))
    v18 := GenS2{F0: (-(v5)), F1: ((!true) || (!false))}
    for fi19 := 0; fi19 < 19; fi19++ {
        if len(v2) > 0 {
            v2[0] = v5
        }
        v20 := GenS1{F0: -1, F1: 0}
        _ = fi19
        _ = v20
    }
    if true {
        if ((v3 * (v6 | v6)) < (^(v6))) {
            var v21 uint = 2
            switch "" {
            case "hello":
                if (0 < 1) {
                    fmt.Println(v3, v21)
                    if len(v2) > 0 {
                        v2[0] = (v5 | v5)
                    }
                } else {
                    v3 -= v3
                }
            default:
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := GenS3{F0: 7}
                    _ = fi22
                    _ = v23
                }
            }
            _ = v21
        } else {
            fmt.Println(v12, v5)
            fmt.Println(v3)
            v4 = (-2.25 - float32(1.5))
        }
        v12 -= v12
    }
    v12 -= (^(v12))
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS4{F0: (v12 % ((332) | 1)), F1: (^(v6)), F2: (v12 & v12)}
        var v26 uint64 = 0
        _ = fi24
        _ = v25
        _ = v26
    }
    if ((v5 % ((128) | 1)) > 148) {
        v2 = append(v2, v5)
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 uint16 = 65535
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 uint8 = 128
                if false {
                    _ = len(v2)
                }
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        var v31 int32 = -1
        _ = v31
    } else {
        v32 := GenS5{F0: 42}
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 uint16 = 42
        switch -835 {
        case 9223372036854775807:
            fmt.Println(v6)
        case 0:
            if (v34 >= v34) {
                var v35 uint8 = 127
                v4 = v4
                v36 := GenS6{F0: (v12 ^ (v12 | 42))}
                _ = v35
                _ = v36
            }
        }
        _ = fi33
        _ = v34
    }
    if false {
        fmt.Println(v5, v12, v5)
        v4 += (1e10 / v4)
        for fi37 := 0; fi37 < 19; fi37++ {
            v2 = append(v2, v5)
            for fi38 := 0; fi38 < 1; fi38++ {
                fmt.Println(v12, fi37, v12)
                fmt.Println(fi38, fi37, v12)
                _ = fi38
            }
            v2 = append(v2, fi37)
            _ = fi37
        }
    }
    var v39 int8 = -18
    fmt.Println(v5, v6)
    if (65535 == 65535) {
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 int8 = 0
            _ = fi40
            _ = v41
        }
        _ = len(v1)
        for fi42 := 0; fi42 < 13; fi42++ {
            v39 -= v39
            _ = fi42
        }
    }
    v1[(v3 / ((v3) | 1))] = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

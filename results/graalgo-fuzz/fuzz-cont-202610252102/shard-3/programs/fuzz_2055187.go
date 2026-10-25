package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 int = -803
    var v2 uint16 = 65535
    var v3 int8 = 1
    var v4 int8 = v3
    v5 := []int{(v1 << ((42) & 63))}
    v6 := GenS1{F0: 65535}
    if true {
        if true {
            switch "\n\t" {
            case "hello":
                fmt.Println(v2, v3, v1)
            default:
                fmt.Println(v1, v1, v2)
            }
            v5 = append(v5, v1)
            if true {
                var v7 int64 = 292
                _ = v7
            } else {
                if false {
                    var v8 string = "hello"
                    v9 := GenS2{F0: 1, F1: (v8 + v8), F2: (v3 % (((-(v3))) | 1))}
                    if true {
                        v10 := GenS2{F0: 65535, F1: v8, F2: (v4 - v3)}
                        _ = v10
                    }
                    _ = v8
                    _ = v9
                }
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        switch fi11 {
                        case 1:
                            v13 := GenS2{F0: 432, F1: "foo bar", F2: (v3 * 127)}
                            _ = v13
                        }
                        var v14 uint32 = 65535
                        v15 := GenS1{F0: v2}
                        _ = fi12
                        _ = v14
                        _ = v15
                    }
                    v4 = 0
                    _ = fi11
                }
            }
        }
        if false {
            v16 := GenS3{F0: v2, F1: 1, F2: 1}
            fmt.Println(v1)
            _ = v16
        } else {
            switch "hello" {
            case "a":
                v17 := GenS1{F0: 32767}
                _ = v17
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                var v18 uint = 18446744073709551615
                _ = v18
            default:
                v3 -= v3
            }
            v19 := GenS2{F0: 455, F1: ("foo bar" + "hello"), F2: (v3 ^ v3)}
            _ = v19
        }
        var v20 uint8 = 2
        _ = v20
    } else {
        if false {
            v1 -= v1
        } else {
            switch "foo bar" {
            case "\n\t":
                switch "foo bar" {
                case "hello":
                    v5 = append(v5, 728)
                case "a":
                    for fi21 := 0; fi21 < 8; fi21++ {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            switch v1 {
                            case 595:
                                v23 := GenS4{F0: 127}
                                _ = v23
                            case -9223372036854775808:
                                var v24 int8 = 127
                                _ = v24
                            case 531:
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v4, v2, fi21)
                                default:
                                    v2 = (v2 & v2)
                                }
                            default:
                                var v25 float32 = 0.1
                                _ = v25
                            }
                            var v26 uint16 = (v2 & 65535)
                            v3 = (^(v3))
                            _ = fi22
                            _ = v26
                        }
                        switch (fi21 & -1) {
                        case 4294967295:
                            var v27 bool = true
                            _ = v27
                        }
                        var v28 uint8 = 1
                        _ = fi21
                        _ = v28
                    }
                case "foo bar":
                    var v29 uint64 = 2
                    _ = v29
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 32768
                }
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 bool = true
                for fi32 := 0; fi32 < 2; fi32++ {
                    v2 -= (v2 >> ((255) & 63))
                    if len(v5) > 0 {
                        v5[0] = -945
                    }
                    var v33 float32 = 0.1
                    v34 := GenS2{F0: 545, F1: "\n\t", F2: 100}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = fi30
                _ = v31
            }
        }
        v35 := GenS5{F0: (v2 / ((2) | 1))}
        v1 = v1
        _ = v35
    }
    v5 = append(v5, v1)
    _ = len(v5)
    v4 -= v4
    if true {
        var v36 uint16 = v2
        _ = v36
    } else {
        switch (v1 & (v1 * 967)) {
        case 127:
            for fi37 := 0; fi37 < 10; fi37++ {
                v4 -= (v4 << ((42) & 63))
                var v38 uint16 = 1
                _ = fi37
                _ = v38
            }
        }
    }
    if true {
        if len(v5) > 0 {
            v5[0] = (v1 ^ 127)
        }
        v39 := GenS6{F0: 4294967295, F1: 207}
        _ = len(v5)
        _ = v39
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        if true {
            var v41 uint = 672
            v42 := GenS7{F0: (v2 << (((v41 | v41)) & 63))}
            switch (fi40 | v1) {
            case -396:
                if ((2 >= 2) && ((-1 < -32768) && ((float64(3.14159) + 0.1) != -2.25))) {
                    var v43 int32 = 42
                    v1 = fi40
                    _ = v43
                }
            default:
                if len(v5) > 0 {
                    v5[0] = (fi40 % ((v1) | 1))
                }
            }
            _ = v41
            _ = v42
        }
        switch "\n\t" {
        case "hello":
            var v44 int64 = -499
            _ = v44
        case "a":
            var v45 uint32 = 818
            _ = v45
        }
        v1 -= fi40
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS6 struct { F0 int16; F1 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 631
    var v3 uint = 42
    if true {
        v1 = (-(v1))
        var v4 string = "a"
        _ = v4
    } else {
        var v5 int8 = -6
        var v6 float64 = 3.14159
        _ = v5
        _ = v6
    }
    if true {
        fmt.Println(v3, v1, v3)
        if false {
            switch ("hello" + ("hello" + "\n\t")) {
            case "a":
                var v7 int = 4294967295
                _ = v7
            case "foo bar":
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := GenS1{F0: 4294967295, F1: 42}
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v1 -= (v1 / ((9223372036854775807) | 1))
                        _ = fi10
                    }
                    var v11 float64 = 1.5
                    _ = fi8
                    _ = v11
                    _ = v9
                }
            case "hello":
                var v12 uint8 = 2
                _ = v12
            }
            var v13 uint = uint(v2)
            _ = v13
        } else {
            v14 := GenS2{F0: 18446744073709551615}
            v15 := GenS3{F0: uint(v2), F1: -32768}
            _ = v14
            _ = v15
        }
    }
    if ((v3 & v3) < v3) {
        v16 := []int{2147483647, -246}
        fmt.Println(v3)
        _ = v16
    } else {
        v17 := []int{127, -1, 700}
        switch 1 {
        case 844:
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint8 = 1
                _ = fi18
                _ = v19
            }
        case 32768:
            var v20 uint64 = 1
            _ = v20
        case 42:
            v2 = v2
        }
        v1 += v1
        _ = v17
    }
    v3 = v3
    if false {
        v2 += v2
    } else {
        var v21 float64 = 3.14159
        _ = v21
    }
    v2 = (v2 | v2)
    v22 := GenS3{F0: (^(v3)), F1: 1}
    fmt.Println(v3, v2)
    var v23 uint16 = 851
    if (-501 >= -32768) {
        v24 := []int{}
        v25 := GenS4{F0: 226, F1: 0, F2: 1}
        _ = len(v24)
        _ = v24
        _ = v25
    }
    fmt.Println(v1)
    if false {
        var v26 float64 = ((float64(-2.25) / 1.0) - (0.0 / -2.25))
        if (0 != 0) {
            if true {
                v27 := []int{2147483647}
                _ = v27
            } else {
                if false {
                    if true {
                        v28 := []int{-280, 32767, 728, 127}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 uint64 = 1
                        _ = fi29
                        _ = v30
                    }
                }
                fmt.Println(v23, v2)
                var v31 uint = (^(v3))
                _ = v31
            }
        }
        _ = v26
    }
    v32 := []int{-2147483648}
    var v33 int8 = 0
    for fi34 := 0; fi34 < 6; fi34++ {
        for fi35 := 0; fi35 < 5; fi35++ {
            if ((!false) || true) {
                fmt.Println(v2)
                if (!false) {
                    v36 := GenS5{F0: 77, F1: (v2 << ((uint(v23)) & 63)), F2: 618}
                    _ = v36
                }
                var v37 uint32 = 1
                v3 += 255
                switch (fi34 - 946) {
                case 188:
                    v38 := GenS1{F0: 1, F1: v37}
                    _ = v38
                default:
                    if (793 > (v23 / ((42) | 1))) {
                        fmt.Println(fi34, v37)
                        if (-2.25 >= 0.0) {
                            v2 = (v23 / ((v2) | 1))
                            switch fi35 {
                            case 506:
                                var v39 int32 = 1
                                _ = v39
                            case 9223372036854775807:
                                switch "hello" {
                                case "foo bar":
                                    _ = len(v32)
                                }
                            }
                        } else {
                            v40 := GenS2{F0: 65535}
                            v41 := GenS6{F0: -1, F1: 1}
                            _ = v40
                            _ = v41
                        }
                        if false {
                            v42 := GenS6{F0: 1, F1: 705}
                            var v43 bool = (v3 != (v3 + 18446744073709551615))
                            _ = v42
                            _ = v43
                        }
                    }
                }
                _ = v37
            }
            var v44 int32 = 462
            v45 := GenS3{F0: v3, F1: -1}
            v1 += v1
            var v46 uint = (v3 ^ v3)
            v33 += (-(v33))
            _ = fi35
            _ = v44
            _ = v45
            _ = v46
        }
        _ = fi34
    }
    if len(v32) > 0 {
        v32[0] = -274
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
}

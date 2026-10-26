package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2147483648
    var v3 int32 = -2147483648
    var v4 uint8 = 2
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch (fi6 * 481) {
        case -593:
            v4 = v4
        case 605:
            var v7 uint32 = v2
            _ = v7
        case -265:
            switch "hello" {
            case "foo bar":
                switch "a" {
                case "":
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v4 = (v4 ^ v4)
                        _ = fi8
                    }
                default:
                    fmt.Println(fi6)
                }
            case "\n\t":
                v9 := []int{(fi6 >> ((1) & 63)), -1}
                _ = v9
            default:
                if (476 != 4294967295) {
                    if true {
                        if true {
                            if (!false) {
                                v10 := GenS1{F0: -4, F1: 1}
                                v11 := GenS1{F0: 42, F1: fi6}
                                _ = v10
                                _ = v11
                            } else {
                                _ = len(v1)
                                if true {
                                    var v12 float32 = 1e10
                                    v13 := GenS1{F0: -1, F1: (fi6 * (fi6 | fi6))}
                                    _ = v12
                                    _ = v13
                                }
                                v14 := v1["hello"]
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 1; fi15++ {
                                var v16 uint16 = uint16(v4)
                                _ = fi15
                                _ = v16
                            }
                        }
                        var v17 uint8 = (^(v4))
                        if true {
                            var v18 float64 = ((0.0 + (3.14159 * (((float64((((((((0.1 * -2.25) / (0.1 / -2.25)) + (float32(float64(float32(float64(3.14159)))) - 1e10)) * 3.14159) / float32(float64(3.14159))) * (1e10 - 0.0)) + 0.0)) - 3.14159) / 3.14159) / 1.0))) - 0.1)
                            var v19 uint16 = 657
                            v20 := []int{(fi6 / ((4294967295) | 1)), (fi6 * (fi6 + fi6)), fi6, fi6}
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 18; fi21++ {
                                var v22 float64 = -2.25
                                _ = fi21
                                _ = v22
                            }
                            var v23 int64 = 1
                            _ = v23
                        }
                        _ = v17
                    } else {
                        v24 := GenS1{F0: 0, F1: (fi6 << ((4294967295) & 63))}
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v26 := GenS2{F0: 42}
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                    if true {
                        v27 := []int{(fi6 | fi6)}
                        _ = v27
                    }
                    v28 := []int{}
                    _ = v28
                }
            }
        }
        v3 = (v3 << ((42) & 63))
        _ = fi6
    }
    var v29 int64 = 561
    var v30 int16 = 32767
    var v31 uint16 = 65535
    delete(v5, v2)
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            switch "\n\t" {
            case "a":
                fmt.Println(v29, v29)
                v31 = 1
                var v34 int64 = 1
                if (!true) {
                    v3 = (v3 >> ((255) & 63))
                    v35 := GenS3{F0: 0, F1: 42}
                    _ = v35
                } else {
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        fmt.Println(v34, v2, v2)
                    case "foo bar":
                        var v36 uint64 = 0
                        _ = v36
                    case "hello":
                        switch "" {
                        case "":
                            v30 += v30
                        }
                    default:
                        var v37 uint64 = 0
                        _ = v37
                    }
                    if (!false) {
                        v38 := GenS2{F0: -614}
                        switch (("a" + ("\n\t" + "foo bar")) + "foo bar") {
                        case "hello":
                            var v39 uint = 752
                            _ = v39
                        }
                        delete(v1, "foo bar")
                        var v40 int32 = v3
                        _ = v38
                        _ = v40
                    }
                    fmt.Println(fi32, fi33)
                }
                v41 := GenS4{F0: 241, F1: -157}
                _ = v34
                _ = v41
            case "\n\t":
                v29 -= v29
            default:
                if (127 == -1) {
                    switch 32767 {
                    case 4294967295:
                        if true {
                            var v42 float64 = (3.14159 - float64(0.1))
                            _ = v42
                        }
                    }
                }
            }
            var v43 uint32 = (v2 * v2)
            _ = fi33
            _ = v43
        }
        if true {
            var v44 uint = 0
            if false {
                var v45 bool = true
                v46 := GenS2{F0: v29}
                _ = v45
                _ = v46
            } else {
                var v47 int = fi32
                for fi48 := 0; fi48 < 6; fi48++ {
                    v49 := v5[(v2 - (v2 + 42))]
                    if (-1 > -128) {
                        v4 += uint8(v44)
                    } else {
                        v30 = int16(v29)
                        v5[v2] = (v47 & fi32)
                        v47 = (fi32 << ((uint(v31)) & 63))
                    }
                    _ = fi48
                    _ = v49
                }
                _ = v47
            }
            _ = v44
        } else {
            v2 += 4294967295
        }
        _ = fi32
    }
    v3 = v3
    v50 := []int{}
    for fi51 := 0; fi51 < 11; fi51++ {
        var v52 uint = 65535
        _ = fi51
        _ = v52
    }
    v3 -= (v3 * (-(v3)))
    v29 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v50
}

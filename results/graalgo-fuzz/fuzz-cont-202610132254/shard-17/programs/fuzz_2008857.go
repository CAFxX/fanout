package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    var v5 int = -89
    _ = len(v4)
    v5 = v5
    var v6 float32 = float32((0.1 - -2.25))
    v7 := GenS1{F0: "hello"}
    var v8 int32 = -1
    v9 := []int{2147483647, 42}
    fmt.Println(v5, v8)
    fmt.Println(v6)
    var v10 int64 = -1
    v11 := GenS1{F0: "a"}
    if (0.0 < v6) {
        v8 = v8
        if false {
            var v12 uint = 18446744073709551615
            switch v5 {
            case -9223372036854775808:
                for fi13 := 0; fi13 < 15; fi13++ {
                    v14 := GenS2{F0: 1}
                    var v15 uint64 = 0
                    v16 := GenS3{F0: -9223372036854775808}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case 9223372036854775807:
                for fi17 := 0; fi17 < 2; fi17++ {
                    _ = len(v9)
                    _ = fi17
                }
            }
            v18 := GenS4{F0: 2, F1: 4294967295, F2: -128}
            _ = v12
            _ = v18
        }
        var v19 int32 = 0
        _ = v19
    } else {
        var v20 int16 = 0
        var v21 int8 = 1
        for fi22 := 0; fi22 < 10; fi22++ {
            switch "a" {
            case "foo bar":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            case "":
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS4{F0: 1, F1: 571, F2: 1}
                    var v25 uint64 = 906
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            default:
                if (0.0 < 1.5) {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 uint64 = 42
                        v10 -= 42
                        v1[v5] = v5
                        _ = fi26
                        _ = v27
                    }
                    var v28 uint16 = 595
                    var v29 int64 = v10
                    _ = v28
                    _ = v29
                } else {
                    v10 = (v10 / (((v10 - (v10 - (^(v10))))) | 1))
                }
            }
            v6 = 1e10
            _ = fi22
        }
        _ = v20
        _ = v21
    }
    v6 = -2.25
    if false {
        if true {
            fmt.Println(v6, v6)
            var v30 uint64 = 207
            _ = v30
        }
        v31 := GenS3{F0: 1}
        _ = len(v9)
        _ = v31
    } else {
        v32 := GenS1{F0: "foo bar"}
        var v33 string = "a"
        _ = v32
        _ = v33
    }
    if (0 <= 2) {
        if true {
            v5 -= (v5 >> ((2) & 63))
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 int16 = 0
                if true {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        if false {
                            switch "hello" {
                            case "":
                                switch "" {
                                case "":
                                    v35 -= -1
                                }
                            case "a":
                                v10 = (v10 + (v10 + (v10 << ((199) & 63))))
                            case "foo bar":
                                switch "foo bar" {
                                case "foo bar":
                                    v5 = v5
                                }
                            }
                            v5 = 284
                            _ = len(v3)
                        }
                        _ = fi36
                    }
                    switch "a" {
                    case "a":
                        for fi37 := 0; fi37 < 2; fi37++ {
                            var v38 uint64 = 4294967295
                            v10 = (v10 ^ v10)
                            var v39 int8 = 100
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    default:
                        v40 := GenS4{F0: 702, F1: 852, F2: -32}
                        _ = v40
                    }
                } else {
                    for fi41 := 0; fi41 < 4; fi41++ {
                        fmt.Println(fi34)
                        v9 = append(v9, (v5 - (fi41 & (fi34 - (fi41 & int(v10))))))
                        _ = fi41
                    }
                }
                if ((127 < -26) && false) {
                    _ = len(v1)
                    v42 := GenS5{F0: 42}
                    _ = v42
                }
                _ = fi34
                _ = v35
            }
        } else {
            switch v5 {
            case 402:
                v9 = append(v9, 9223372036854775807)
            case 127:
                v2[561] = 127
            default:
                var v43 float64 = ((0.0 / 1.0) * (3.14159 - 0.1))
                _ = v43
            }
            if (!true) {
                var v44 int32 = (v8 - v8)
                v4[(-(v5))] = 777
                _ = v44
            }
            var v45 int64 = (v10 / ((v10) | 1))
            _ = v45
        }
    } else {
        for fi46 := 0; fi46 < 18; fi46++ {
            switch (fi46 / ((v5) | 1)) {
            case 1:
                v8 = 42
            default:
                var v47 float64 = 3.14159
                _ = v47
            }
            var v48 string = "a"
            for fi49 := 0; fi49 < 2; fi49++ {
                v5 += (v5 >> ((504) & 63))
                if (!true) {
                    if len(v9) > 0 {
                        v9[0] = -784
                    }
                    v6 = (v6 / 1.5)
                }
                _ = fi49
            }
            _ = fi46
            _ = v48
        }
        var v50 uint64 = 1
        _ = v50
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        if (255 >= 65535) {
            if true {
                var v52 int64 = (v10 & 0)
                delete(v4, fi51)
                _ = v52
            }
            v53 := GenS6{F0: -109}
            _ = v53
        }
        for fi54 := 0; fi54 < 13; fi54++ {
            var v55 uint32 = 1
            if true {
                switch ("hello" + "foo bar") {
                case "hello":
                    var v56 bool = true
                    _ = v56
                case "foo bar":
                    v55 += (v55 - v55)
                case "\n\t":
                    v57 := GenS7{F0: 2147483647, F1: 4294967295}
                    _ = v57
                }
                _ = len(v9)
                var v58 float32 = ((v6 - (v6 - v6)) - (float32(1e10) / v6))
                if false {
                    if len(v9) > 0 {
                        v9[0] = 4294967295
                    }
                    var v59 int64 = v10
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        fmt.Println(v59, v59)
                    default:
                        v58 = (-2.25 + v58)
                    }
                    _ = v59
                } else {
                    var v60 uint32 = (v55 * (v55 | (v55 >> ((0) & 63))))
                    fmt.Println(v60, v5, v60)
                    switch (fi51 >> ((4294967295) & 63)) {
                    case 0:
                        var v61 int8 = 127
                        _ = v61
                    case 1:
                        var v62 uint8 = 128
                        _ = v62
                    case -104:
                        var v63 uint64 = 42
                        _ = v63
                    default:
                        var v64 uint16 = 32767
                        _ = v64
                    }
                    v60 = 42
                    var v65 uint8 = 2
                    _ = v60
                    _ = v65
                }
                _ = v58
            } else {
                delete(v2, (^(v55)))
            }
            _ = fi54
            _ = v55
        }
        if false {
            v10 += v10
        } else {
            v10 -= (v10 ^ -1)
            v5 += 128
            v8 -= v8
        }
        _ = fi51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

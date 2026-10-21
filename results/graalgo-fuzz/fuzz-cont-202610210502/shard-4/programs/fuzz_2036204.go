package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0() uint8 {
    var v5 int8 = 1
    if true {
        switch -38 {
        case 9223372036854775807:
            v6 := []int{32767, -2147483648}
            _ = v6
        }
    }
    v5 += v5
    _ = v5
    return 2
}
func genFn1(p0 uint) uint16 {
    v7 := []int{127, -1, 1, 9223372036854775807}
    v7 = append(v7, -2147483648)
    p0 = p0
    var v8 int8 = 100
    if (4294967295 > -142) {
        v8 -= 127
        switch 42 {
        case -9223372036854775808:
            p0 += (p0 + (p0 + 255))
        case -2147483648:
            var v9 uint64 = 479
            _ = v9
        case -1:
            switch 2147483647 {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 12; fi10++ {
                    if ("a" >= "a") {
                        fmt.Println(p0)
                        var v11 bool = true
                        _ = v11
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = fi10
                }
            case 32767:
                p0 = (^(p0))
            case 4294967295:
                v8 = v8
            }
        }
        var v12 string = "\n\t"
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 248
}

func main() {
    v1 := []uint32{1, 2, 671}
    var v2 float32 = ((0.0 * 1e10) * 3.14159)
    var v3 uint8 = 128
    var v4 int16 = -32768
    v2 = (v2 + 1.5)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int32 = 0
            var v16 int64 = -9223372036854775808
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = len(v1)
        if (fi13 > fi13) {
            v3 -= (v3 - v3)
            v4 = (v4 >> ((2) & 63))
            v17 := GenS1{F0: -128}
            _ = v17
        }
        _ = fi13
    }
    var v18 string = "hello"
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if (uint(v3) < 18446744073709551615) {
            v20 := GenS2{F0: 2}
            _ = v20
        }
        _ = len(v1)
        if (335 <= 42) {
            var v21 float32 = 0.1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint16 = 65535
                _ = fi22
                _ = v23
            }
            switch "" {
            case "":
                var v24 int16 = -1
                _ = v24
            case "a":
                var v25 uint64 = 883
                _ = v25
            }
        }
        _ = fi19
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v26 := GenS3{F0: genFn0(), F1: false, F2: (v3 & (v3 >> ((2) & 63)))}
    if (2147483648 == 2147483647) {
        v18 = v18
        if true {
            switch "a" {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                v27 := GenS1{F0: -1}
                _ = v27
            case "hello":
                if (!false) {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 int64 = -542
                        v30 := GenS2{F0: 42}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS1{F0: 127}
                    if true {
                        var v32 int = -9223372036854775808
                        var v33 int8 = 1
                        var v34 uint32 = 65535
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        switch v18 {
                        case "foo bar":
                            v35 := GenS2{F0: 0}
                            _ = v35
                        case "\n\t":
                            switch 128 {
                            case 32767:
                                var v36 int64 = -9223372036854775808
                                _ = v36
                            case 127:
                                if false {
                                    var v37 uint64 = 2
                                    var v38 bool = ((v3 ^ (v3 - (v3 << ((0) & 63)))) < genFn0())
                                    var v39 int32 = -581
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    if (int(v4) != 825) {
                                        var v40 float64 = -2.25
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            if (v18 != v18) {
                                                var v42 uint8 = (v3 + 0)
                                                _ = v42
                                            }
                                            if false {
                                                fmt.Println(fi41, v2)
                                            } else {
                                                var v43 string = v18
                                                _ = v43
                                            }
                                            _ = fi41
                                        }
                                        var v44 int = -868
                                        _ = v40
                                        _ = v44
                                    }
                                    var v45 bool = true
                                    _ = v45
                                }
                            case 1:
                                var v46 uint32 = 1
                                _ = v46
                            default:
                                var v47 int64 = -286
                                _ = v47
                            }
                        case "hello":
                            if true {
                                fmt.Println(v3, v3, v3)
                                _ = len(v1)
                            }
                        default:
                            v18 = "\n\t"
                        }
                        fmt.Println(v4, v2, v18)
                        fmt.Println(v4, v3)
                    }
                    _ = v31
                } else {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        v4 += v4
                        _ = fi48
                    }
                    for fi49 := 0; fi49 < 9; fi49++ {
                        fmt.Println(fi49, v2)
                        _ = fi49
                    }
                    v50 := GenS4{F0: 369}
                    _ = v50
                }
            }
            v51 := GenS5{F0: 42, F1: 922, F2: -9223372036854775808}
            _ = v51
        }
    } else {
        _ = len(v1)
    }
    v18 = v18
    var v52 int32 = -600
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 uint = 18446744073709551615
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 4; fi55++ {
        var v56 uint = 0
        fmt.Println(v3)
        v57 := GenS4{F0: (v52 << ((v56) & 63))}
        _ = fi55
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        if (v18 >= "a") {
            fmt.Println(fi58, v4, v2)
        } else {
            switch v18 {
            case "a":
                v52 = (v52 / ((1) | 1))
            }
            fmt.Println(v18, v18)
        }
        var v59 uint32 = 4294967295
        v2 += v2
        _ = fi58
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v52
}

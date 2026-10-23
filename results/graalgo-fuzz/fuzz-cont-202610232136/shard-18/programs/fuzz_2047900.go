package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 int32) uint64 {
    switch 32767 {
    case -1:
        v4 := GenS1{F0: 128}
        _ = v4
    }
    if true {
        v5 := GenS2{F0: 9223372036854775807}
        _ = v5
    }
    _ = p0
    return 4294967295
}
func genFn1(p0 uint) uint64 {
    if false {
        v6 := GenS3{F0: true, F1: 527, F2: (127 <= 1)}
        var v7 int64 = 0
        _ = v6
        _ = v7
    } else {
        if true {
            v8 := []int{9, 9223372036854775807, 127, -258}
            fmt.Println(p0)
            _ = v8
        }
    }
    var v9 float32 = 1e10
    if false {
        var v10 float32 = v9
        fmt.Println(p0, p0)
        var v11 float64 = 1.5
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v9
    return 4294967295
}

func main() {
    var v1 int16 = 1
    v2 := []string{"\n\t", ("foo bar" + "foo bar"), "hello", "foo bar"}
    v3 := []uint32{2}
    switch "\n\t" {
    case "a":
        v12 := GenS1{F0: 2}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS4{F0: (fi13 - fi13)}
            v1 += v1
            v1 -= 376
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v1 = v1
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        switch "\n\t" {
        case "a":
            switch fi15 {
            case -9223372036854775808:
                for fi16 := 0; fi16 < 2; fi16++ {
                    fmt.Println(v1, fi15)
                    var v17 string = "\n\t"
                    v17 = (((((("" + (v17 + v17)) + "\n\t") + "\n\t") + v17) + "a") + v17)
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    v17 = v17
                    _ = fi16
                    _ = v17
                }
            case 955:
                v1 = v1
            default:
                for fi18 := 0; fi18 < 14; fi18++ {
                    if (!(42 >= 4294967295)) {
                        var v19 int16 = (v1 + (v1 | v1))
                        v20 := GenS5{F0: 0}
                        v1 -= (v19 + int16(fi18))
                        _ = v19
                        _ = v20
                    } else {
                        v2 = append(v2, "\n\t")
                    }
                    if false {
                        if false {
                            v21 := GenS3{F0: (0 <= 0), F1: 1, F2: true}
                            fmt.Println(v1)
                            _ = v21
                        } else {
                            var v22 uint = 306
                            var v23 uint32 = 0
                            _ = v22
                            _ = v23
                        }
                        v1 -= (v1 << ((0) & 63))
                        _ = len(v3)
                    } else {
                        var v24 float32 = -2.25
                        fmt.Println(v1)
                        if false {
                            var v25 float64 = 0.0
                            if (true || true) {
                                fmt.Println(fi18, fi15, v24)
                                if true {
                                    var v26 float32 = ((v24 + v24) - 0.1)
                                    _ = v26
                                } else {
                                    var v27 float32 = v24
                                    if false {
                                        v3 = append(v3, 627)
                                        var v28 string = ("foo bar" + "")
                                        _ = v28
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    _ = v27
                                }
                                var v29 int32 = 2147483647
                                _ = v29
                            }
                            _ = v25
                        }
                        _ = v24
                    }
                    _ = fi18
                }
            }
        case "hello":
            for fi30 := 0; fi30 < 20; fi30++ {
                fmt.Println(fi30, fi30, v1)
                v1 += (v1 << ((708) & 63))
                fmt.Println(fi15)
                _ = fi30
            }
        }
        _ = len(v3)
        v1 = (v1 % ((v1) | 1))
        _ = fi15
    }
    v1 -= v1
    v31 := GenS6{F0: 42, F1: 2, F2: 2}
    if true {
        if false {
            var v32 uint16 = 1
            _ = v32
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v2)
    }
    var v33 float64 = 0.1
    var v34 float32 = 1.5
    var v35 float32 = 1.5
    if true {
        for fi36 := 0; fi36 < 3; fi36++ {
            v34 += 0.0
            v33 = ((1.5 / (v33 + (((v33 + v33) * v33) / (v33 * float64((float32(0.1) - v35)))))) / -2.25)
            _ = fi36
        }
        var v37 uint16 = 243
        _ = v37
    } else {
        switch "" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            fmt.Println(v35)
        }
        v38 := GenS6{F0: 255, F1: 915, F2: 1}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
}

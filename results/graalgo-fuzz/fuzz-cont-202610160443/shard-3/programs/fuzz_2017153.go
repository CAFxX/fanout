package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 int }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 411
    v3 := []uint32{42}
    for fi4 := 0; fi4 < 15; fi4++ {
        if (("" + "") == "hello") {
            var v5 int16 = 32767
            v2 = v2
            for fi6 := 0; fi6 < 17; fi6++ {
                v5 = (v5 + v5)
                if true {
                    v7 := GenS1{F0: 2, F1: 5}
                    var v8 int8 = 61
                    _ = v7
                    _ = v8
                } else {
                    var v9 int64 = 1
                    _ = v9
                }
                switch (fi6 / ((859) | 1)) {
                case 4294967295:
                    var v10 float64 = (1.5 + 1.5)
                    _ = v10
                case -481:
                    var v11 int = 2147483647
                    _ = v11
                case 322:
                    var v12 uint16 = 1
                    _ = v12
                }
                _ = fi6
            }
            _ = v5
        } else {
            v13 := GenS1{F0: 1, F1: 47}
            var v14 uint8 = 2
            _ = v13
            _ = v14
        }
        var v15 int16 = 1000
        _ = fi4
        _ = v15
    }
    var v16 int32 = 42
    if true {
        v16 += v16
    }
    v17 := GenS2{F0: -109}
    var v18 float32 = ((0.1 / 1.0) - float32(0.1))
    v18 -= v18
    var v19 float32 = -2.25
    fmt.Println(v18)
    var v20 uint32 = 268
    switch -9223372036854775808 {
    case -1:
        v21 := GenS3{F0: 0}
        _ = v21
    case 673:
        if len(v3) > 0 {
            v3[0] = (v20 - v20)
        }
    case 32768:
        if (0 < 18446744073709551615) {
            v20 -= 0
            v16 += (v16 - v16)
        } else {
            _ = len(v1)
            if (!((-33 <= -63) && ((v18 > float32(0.1)) || (("foo bar" + ("a" + (("" + "hello") + (((("foo bar" + "\n\t") + "\n\t") + "hello") + ("foo bar" + "hello"))))) != ((("\n\t" + "") + "\n\t") + ""))))) {
                if false {
                    switch 32767 {
                    case 0:
                        v16 += v16
                    case -461:
                        if true {
                            _ = len(v1)
                        }
                    default:
                        if true {
                            v1[(v20 + v20)] = v20
                        } else {
                            if false {
                                switch (("foo bar" + ("hello" + ("a" + ""))) + "a") {
                                case "\n\t":
                                    var v22 bool = true
                                    _ = v22
                                default:
                                    v19 = v19
                                }
                            } else {
                                var v23 uint64 = 42
                                var v24 uint32 = 2
                                _ = v23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 16; fi25++ {
                                v26 := GenS4{F0: 0, F1: 49, F2: (fi25 / (((^(fi25))) | 1))}
                                v20 = v20
                                _ = fi25
                                _ = v26
                            }
                            fmt.Println(v16)
                        }
                    }
                }
                for fi27 := 0; fi27 < 9; fi27++ {
                    v28 := GenS5{F0: v2, F1: 1}
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v16 = (-(v16))
                        _ = fi29
                    }
                    for fi30 := 0; fi30 < 4; fi30++ {
                        var v31 uint32 = 508
                        _ = fi30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            }
        }
    default:
        delete(v1, v20)
    }
    if len(v3) > 0 {
        v3[0] = v20
    }
    v3 = append(v3, v20)
    delete(v1, (^(v20)))
    var v32 int64 = 621
    var v33 int64 = v32
    if false {
        for fi34 := 0; fi34 < 4; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                v36 := GenS3{F0: 42}
                _ = fi35
                _ = v36
            }
            if false {
                v1[v20] = (v20 | (v20 >> (((^(v2))) & 63)))
            }
            v37 := GenS6{F0: -328}
            _ = fi34
            _ = v37
        }
        var v38 int64 = v32
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
}

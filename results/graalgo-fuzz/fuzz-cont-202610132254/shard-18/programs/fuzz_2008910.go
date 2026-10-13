package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string; F2 int64 }

func main() {
    var v1 int64 = -180
    v2 := []uint32{}
    var v3 uint32 = 0
    v1 = (-(v1))
    var v4 int8 = 100
    v4 -= int8(v1)
    var v5 string = "\n\t"
    if (((v1 | v1) < -9223372036854775808) || (!false)) {
        v3 -= v3
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint32 = v3
            var v8 uint32 = v3
            _ = len(v2)
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int16 = int16(v1)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            if (250 >= 127) {
                v1 += 42
                v4 -= (v4 - v4)
                if true {
                    v4 = -1
                    var v11 float64 = 0.0
                    var v12 uint64 = 497
                    _ = v11
                    _ = v12
                } else {
                    _ = len(v2)
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 int8 = v4
                        var v15 int64 = (v1 | int64(v4))
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 uint64 = 4294967295
                        fmt.Println(fi10, v3)
                        _ = fi16
                        _ = v17
                    }
                }
            } else {
                _ = len(v2)
                var v18 uint8 = 127
                _ = v18
            }
            _ = fi10
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int = 42
            _ = fi19
            _ = v20
        }
    }
    if (-1 > 42) {
        v1 -= (v1 << ((18446744073709551615) & 63))
    }
    if (!(v1 >= -9223372036854775808)) {
        if false {
            fmt.Println(v3)
            v5 = "\n\t"
        }
        fmt.Println(v3, v4)
        if (65535 >= 836) {
            fmt.Println(v3, v5, v4)
            v1 = (v1 / ((v1) | 1))
            for fi21 := 0; fi21 < 17; fi21++ {
                if false {
                    var v22 uint = uint(v3)
                    var v23 int16 = int16(fi21)
                    _ = v22
                    _ = v23
                }
                var v24 int64 = (v1 * 436)
                _ = fi21
                _ = v24
            }
        }
    } else {
        v25 := GenS1{F0: true, F1: 838}
        var v26 float64 = float64(0.1)
        _ = v25
        _ = v26
    }
    var v27 int32 = 912
    if (42 == 145) {
        v28 := GenS2{F0: 255, F1: ("\n\t" + "")}
        _ = v28
    } else {
        if (-2.25 <= (1.5 - 1e10)) {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 int8 = 127
                if true {
                    switch v5 {
                    case "":
                        v31 := GenS1{F0: true, F1: 845}
                        _ = v31
                    case "hello":
                        v2 = append(v2, v3)
                    case "a":
                        fmt.Println(fi29, v1)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = fi29
                _ = v30
            }
        } else {
            var v32 uint8 = 255
            var v33 int8 = (v4 + v4)
            _ = v32
            _ = v33
        }
        v34 := GenS3{F0: 290}
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 int16 = 1
            var v37 uint8 = 42
            v38 := GenS4{F0: (v4 | 127), F1: ("a" + ("a" + "\n\t")), F2: (^(v1))}
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = v34
    }
    var v39 int64 = 150
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}

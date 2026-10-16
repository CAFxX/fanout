package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16) int32 {
    v4 := GenS1{F0: -2147483648, F1: int8(p0), F2: 0}
    p0 -= 32767
    _ = p0
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 42
    var v3 uint64 = v2
    _ = len(v1)
    v5 := GenS1{F0: 32767, F1: -1, F2: -32768}
    var v6 uint8 = 1
    v3 -= (v3 % ((733) | 1))
    v6 = 255
    v2 = (^(v3))
    v6 = (v6 * 42)
    var v7 uint16 = uint16(v6)
    for fi8 := 0; fi8 < 4; fi8++ {
        v3 += 528
        v2 = v2
        _ = fi8
    }
    v9 := GenS2{F0: true, F1: -2147483648}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10)
        v2 += v3
        var v11 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
    }
    var v12 float64 = 0.1
    switch 128 {
    case 0:
        var v13 float32 = (1.5 + 1.5)
        _ = v13
    case 4294967295:
        v14 := GenS2{F0: ((!true) || true), F1: 32767}
        _ = v14
    case 32768:
        switch -1 {
        case 127:
            if false {
                var v15 float64 = 1.5
                if true {
                    v3 += 0
                    if true {
                        v16 := GenS2{F0: (!(int64(v6) < 9223372036854775807)), F1: 1}
                        var v17 uint64 = (v3 | (v2 - v3))
                        _ = v16
                        _ = v17
                    }
                    var v18 int16 = 1000
                    _ = v18
                }
                if (42 <= genFn0(v7)) {
                    v19 := GenS2{F0: (!true), F1: -9223372036854775808}
                    _ = v19
                } else {
                    switch "foo bar" {
                    case "hello":
                        v20 := GenS3{F0: "hello", F1: 4294967295, F2: false}
                        _ = v20
                    case "":
                        fmt.Println(v6)
                    }
                    if (!false) {
                        if false {
                            var v21 int64 = 0
                            _ = len(v1)
                            _ = v21
                        } else {
                            v7 -= v7
                            for fi22 := 0; fi22 < 19; fi22++ {
                                v23 := GenS1{F0: 42, F1: 127, F2: 0}
                                var v24 float32 = (1e10 + float32((0.0 + v15)))
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    v26 := v1[fi22]
                                    v27 := GenS2{F0: (float32(float64(v24)) > (v24 * (v24 - (v24 + float32(((-2.25 * -2.25) * 0.0)))))), F1: 42}
                                    if true {
                                        var v28 uint64 = uint64(v7)
                                        if false {
                                            delete(v1, v26)
                                            v6 += v6
                                        } else {
                                            v29 := v1[(fi25 - -721)]
                                            v26 = (v29 >> ((2) & 63))
                                            _ = v29
                                        }
                                        var v30 float32 = v24
                                        _ = v28
                                        _ = v30
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                        var v31 uint32 = 2147483648
                        _ = v31
                    }
                }
                _ = v15
            }
        case -269:
            v7 += v7
        }
    default:
        for fi32 := 0; fi32 < 10; fi32++ {
            switch 9223372036854775807 {
            case -9223372036854775808:
                v33 := GenS2{F0: true, F1: fi32}
                _ = v33
            case 32767:
                for fi34 := 0; fi34 < 16; fi34++ {
                    if false {
                        v6 -= v6
                    }
                    v6 = 128
                    _ = fi34
                }
            default:
                v35 := GenS4{F0: 422}
                _ = v35
            }
            _ = fi32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

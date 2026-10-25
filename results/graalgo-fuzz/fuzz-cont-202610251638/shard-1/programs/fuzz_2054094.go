package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint64 }

func main() {
    v1 := []int{127, 127, -21, 644}
    var v2 string = "foo bar"
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        switch 32767 {
        case 127:
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 float32 = ((1.5 / 1.5) / (-2.25 - 1e10))
                v8 := GenS1{F0: 32767, F1: 2147483647, F2: 127}
                delete(v4, 854)
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 bool = false
            switch 4294967295 {
            case 127:
                v11 := GenS2{F0: (1e10 > (0.1 * (((((0.0 * (3.14159 - (0.0 - float32(1.5)))) / 1.5) * 0.1) * 0.0) + (3.14159 + 1.5))))}
                _ = v11
            case 4294967295:
                switch (v2 + (((v2 + v2) + "foo bar") + (v2 + "a"))) {
                case "a":
                    v12 := GenS3{F0: -1, F1: v2}
                    _ = v12
                case "":
                    if true {
                        var v13 uint8 = 161
                        if v10 {
                            var v14 string = (v2 + (v2 + v2))
                            v15 := GenS4{F0: -159, F1: 607, F2: -443}
                            var v16 uint8 = (v13 ^ v13)
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    } else {
                        v3 -= (v3 % ((v3) | 1))
                        if v10 {
                            if v10 {
                                var v17 int16 = 854
                                _ = v17
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v18 uint16 = 32767
                                _ = v18
                            }
                            if (!v10) {
                                var v19 float64 = 0.0
                                _ = v19
                            } else {
                                fmt.Println(v3)
                            }
                        } else {
                            var v20 int8 = 0
                            v3 += (v3 >> ((1) & 63))
                            _ = v20
                        }
                        if (4294967295 < 18446744073709551615) {
                            v21 := GenS2{F0: false}
                            v22 := GenS5{F0: 1, F1: 2, F2: (-(v3))}
                            v23 := GenS1{F0: 42, F1: -2147483648, F2: 1}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                case "foo bar":
                    v24 := GenS5{F0: 936, F1: 295, F2: v3}
                    _ = v24
                default:
                    v2 = v2
                }
            }
            v3 = v3
            _ = fi9
            _ = v10
        }
        _ = fi5
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v2 = "foo bar"
        v26 := GenS4{F0: 0, F1: 705, F2: -1}
        _ = fi25
        _ = v26
    }
    var v27 uint32 = 922
    for fi28 := 0; fi28 < 17; fi28++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            var v29 int16 = -37
            _ = v29
        } else {
            switch fi28 {
            case 128:
                v3 = (v3 | 549)
            }
        }
        _ = fi28
    }
    v1 = append(v1, -2147483648)
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int64 = 42
        var v32 float32 = 3.14159
        var v33 int = 9223372036854775807
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    v2 = (v2 + "a")
    for fi34 := 0; fi34 < 3; fi34++ {
        v2 = (v2 + v2)
        _ = fi34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            v27 = v27
            var v37 uint64 = 0
            var v38 bool = false
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    var v39 uint32 = 0
    switch -274 {
    case -1:
        v40 := GenS1{F0: -156, F1: -1, F2: 127}
        _ = v40
    default:
        var v41 int64 = (^(v3))
        _ = v41
    }
    v42 := GenS6{F0: 142, F1: 829, F2: 564}
    for fi43 := 0; fi43 < 12; fi43++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v44 string = v2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi43
        _ = v44
    }
    v2 = v2
    v27 -= 681
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v42
}

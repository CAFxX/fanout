package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 bool; F1 string; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 int16; F2 string }
func genFn0() uint32 {
    var v6 float64 = 0.0
    if true {
        v6 = v6
    } else {
        v7 := []int{}
        v6 = 0.0
        switch "a" {
        case "\n\t":
            var v8 int32 = -505
            _ = v8
        default:
            for fi9 := 0; fi9 < 9; fi9++ {
                for fi10 := 0; fi10 < 11; fi10++ {
                    if false {
                        v11 := GenS1{F0: 851, F1: -381, F2: 127}
                        switch (("hello" + "") + (("a" + "") + "foo bar")) {
                        case "foo bar":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        v12 := GenS1{F0: 111, F1: -1, F2: 0}
                        v13 := GenS1{F0: 255, F1: 398, F2: 255}
                        var v14 uint = 548
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS1{F0: 42, F1: -1, F2: 215}
                    var v16 uint16 = 42
                    switch fi10 {
                    case 1:
                        v17 := GenS2{F0: 6, F1: (2 <= 42), F2: 59}
                        _ = v17
                    case 166:
                        v16 = (v16 ^ v16)
                    default:
                        _ = len(v7)
                    }
                    if false {
                        if (0 > 42) {
                            switch (fi10 + (fi9 ^ 73)) {
                            case 0:
                                var v18 int16 = 42
                                _ = v18
                            case 9223372036854775807:
                                var v19 int32 = 1
                                _ = v19
                            }
                            var v20 bool = true
                            fmt.Println(v20)
                            var v21 int = fi10
                            var v22 bool = (-463 < int64(fi10))
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi10
                    _ = v15
                    _ = v16
                }
                if (915 > 2147483647) {
                    var v23 uint64 = 2
                    _ = len(v7)
                    var v24 int64 = 202
                    _ = v23
                    _ = v24
                }
                v6 = v6
                _ = fi9
            }
        }
        _ = v7
    }
    _ = v6
    return 74
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{575, 0, 2}
    var v3 int64 = 0
    var v4 uint16 = 0
    v5 := make(map[int]bool)
    v4 -= 32767
    for fi25 := 0; fi25 < 6; fi25++ {
        v4 = 2
        var v26 uint8 = 249
        for fi27 := 0; fi27 < 18; fi27++ {
            if false {
                var v28 int8 = 42
                v29 := GenS3{F0: (1000 != 0), F1: "\n\t", F2: (float32(0.1) < (0.1 * (1.5 / 3.14159)))}
                _ = v28
                _ = v29
            } else {
                v30 := GenS4{F0: (v4 >> ((1) & 63))}
                _ = v30
            }
            v31 := GenS3{F0: true, F1: "foo bar", F2: true}
            v32 := GenS4{F0: 32767}
            _ = fi27
            _ = v31
            _ = v32
        }
        _ = fi25
        _ = v26
    }
    v4 -= (v4 << ((18446744073709551615) & 63))
    v4 = 1
    for fi33 := 0; fi33 < 14; fi33++ {
        v4 = 665
        _ = fi33
    }
    v34 := GenS5{F0: false}
    for fi35 := 0; fi35 < 20; fi35++ {
        if false {
            for fi36 := 0; fi36 < 2; fi36++ {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                _ = len(v2)
                if (154 >= 25) {
                    fmt.Println(v4)
                    var v37 int16 = -271
                    _ = v37
                } else {
                    _ = len(v1)
                    var v38 int8 = -50
                    v3 = v3
                    _ = v38
                }
                var v39 float32 = (3.14159 - (0.1 * 0.1))
                var v40 int64 = (v3 / ((v3) | 1))
                v4 = v4
                _ = fi36
                _ = v39
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 16; fi41++ {
                v42 := GenS6{F0: 701, F1: -76, F2: ("" + "hello")}
                v4 = v4
                v43 := GenS5{F0: true}
                v3 -= 0
                _ = fi41
                _ = v42
                _ = v43
            }
            fmt.Println(v4)
        }
        _ = fi35
    }
    v4 = 2
    v5[-236] = (-471 < 42)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]int)
    var v3 uint = 4294967295
    if (1.5 < 1e10) {
        switch -9223372036854775808 {
        case 32767:
            var v4 uint64 = 4294967295
            _ = v4
        case 9223372036854775807:
            if (255 < 255) {
                for fi5 := 0; fi5 < 12; fi5++ {
                    switch fi5 {
                    case -1:
                        for fi6 := 0; fi6 < 6; fi6++ {
                            v7 := []int{(fi5 / ((fi5) | 1)), 32768}
                            var v8 bool = (false && false)
                            if v8 {
                                if false {
                                    var v9 uint32 = 42
                                    var v10 int32 = 908
                                    _ = v10
                                    _ = v9
                                }
                                if (!v8) {
                                    switch 0 {
                                    case 32767:
                                        var v11 int32 = -2147483648
                                        v1 += (v1 >> ((v3) & 63))
                                        fmt.Println(v3)
                                        _ = v11
                                    case 9223372036854775807:
                                        v1 += (v1 >> ((v3) & 63))
                                    default:
                                        v8 = v8
                                    }
                                    var v12 uint = v3
                                    v2[4294967295] = (fi6 / ((fi5) | 1))
                                    _ = v12
                                }
                                var v13 uint = v3
                                _ = v13
                            } else {
                                var v14 string = ("" + "")
                                v3 -= (v3 + 0)
                                _ = v14
                            }
                            _ = fi6
                            _ = v7
                            _ = v8
                        }
                    case 4294967295:
                        v2[4294967295] = (fi5 + fi5)
                    default:
                        if ((65535 == 65535) || (!false)) {
                            switch (^(fi5)) {
                            case 42:
                                v15 := GenS1{F0: 0}
                                _ = v15
                            }
                            var v16 uint16 = 0
                            v3 += v3
                            _ = v16
                        } else {
                            var v17 uint32 = 2147483648
                            _ = v17
                        }
                    }
                    var v18 uint8 = 127
                    fmt.Println(v3)
                    _ = fi5
                    _ = v18
                }
                v19 := GenS2{F0: 126, F1: uint64(v3)}
                _ = v19
            } else {
                v20 := []int{9223372036854775807, -793}
                fmt.Println(v1, v1)
                fmt.Println(v3, v3)
                _ = v20
            }
        }
    } else {
        if false {
            v21 := v2[673]
            _ = v21
        }
        for fi22 := 0; fi22 < 17; fi22++ {
            v1 = 42
            v23 := GenS3{F0: 32767}
            var v24 float32 = 1.5
            _ = fi22
            _ = v23
            _ = v24
        }
        fmt.Println(v3, v3)
    }
    fmt.Println(v3)
    if false {
        v1 = (v1 << ((258) & 63))
    } else {
        v25 := v2[65535]
        var v26 int8 = -96
        _ = v25
        _ = v26
    }
    var v27 bool = false
    if (v27 && (1 == v1)) {
        var v28 uint16 = 65535
        v29 := GenS3{F0: -1}
        _ = v28
        _ = v29
    }
    switch 839 {
    case -1:
        if (65535 != 588) {
            v30 := GenS4{F0: 128, F1: (^(v3))}
            _ = v30
        } else {
            var v31 int8 = -1
            fmt.Println(v27)
            _ = v31
        }
    case 128:
        v1 = v1
    default:
        v27 = (!false)
    }
    v32 := []int{42}
    v27 = (434 > 544)
    for fi33 := 0; fi33 < 15; fi33++ {
        switch "a" {
        case "foo bar":
            var v34 int64 = 9223372036854775807
            _ = v34
        case "\n\t":
            v1 = (v1 % ((v1) | 1))
        }
        if len(v32) > 0 {
            v32[0] = fi33
        }
        v32 = append(v32, 1)
        _ = fi33
    }
    v27 = (v27 && v27)
    var v35 int16 = 0
    v36 := GenS2{F0: 127, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v35
    _ = v36
}

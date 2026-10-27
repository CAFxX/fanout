package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint; F2 int32 }

func main() {
    var v1 int8 = 112
    v2 := []uint32{604, 446, 42}
    v3 := make(map[int]uint32)
    switch -900 {
    case 32768:
        fmt.Println(v1)
    case 128:
        for fi4 := 0; fi4 < 14; fi4++ {
            fmt.Println(v1, v1)
            _ = fi4
        }
    case 127:
        v1 -= (v1 - 127)
    default:
        for fi5 := 0; fi5 < 9; fi5++ {
            if true {
                v1 += v1
                v1 = v1
            }
            _ = fi5
        }
    }
    if len(v2) > 0 {
        v2[0] = 167
    }
    if (-258 <= 4294967295) {
        if false {
            if false {
                var v6 uint8 = 0
                _ = v6
            } else {
                v1 = 100
                switch 1 {
                case 9223372036854775807:
                    if ((!(140 >= 4294967295)) && true) {
                        var v7 bool = false
                        v3[42] = 916
                        _ = v7
                    }
                case 32768:
                    var v8 int16 = 0
                    _ = v8
                default:
                    v9 := GenS1{F0: v1}
                    _ = v9
                }
                var v10 bool = false
                _ = v10
            }
            v1 -= -1
            var v11 int16 = 755
            _ = v11
        } else {
            fmt.Println(v1)
            if false {
                v12 := GenS2{F0: 922, F1: 128}
                var v13 uint16 = 0
                if false {
                    v13 = v13
                    v14 := GenS3{F0: 128, F1: 2, F2: 2147483647}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS3{F0: 0, F1: 44, F2: 573}
                        switch fi15 {
                        case 127:
                            _ = len(v2)
                        }
                        var v17 int32 = -352
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v1 += (-(v1))
                        var v19 uint64 = 849
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS3{F0: 2, F1: 750, F2: -2147483648}
                    _ = v20
                }
                _ = v12
                _ = v13
            }
            v1 = v1
        }
        var v21 int = -9223372036854775808
        _ = v21
    }
    var v22 int64 = 42
    v23 := GenS2{F0: 4294967295, F1: 255}
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 bool = true
        _ = fi24
        _ = v25
    }
    var v26 string = "hello"
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            v22 = 974
            _ = fi28
        }
        _ = fi27
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 -= v1
        v22 -= (v22 >> ((42) & 63))
        _ = fi29
    }
    var v30 int16 = -258
    var v31 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v31
}

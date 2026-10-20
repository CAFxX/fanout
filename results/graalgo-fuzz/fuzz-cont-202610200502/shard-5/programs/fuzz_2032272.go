package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 string; F2 int8 }
func genFn0(p0 uint16, p1 int64, p2 int32) int16 {
    if true {
        var v5 int = int(p1)
        for fi6 := 0; fi6 < 11; fi6++ {
            switch -415 {
            case -839:
                switch v5 {
                case 127:
                    p2 = 610
                }
            case 127:
                v5 += (-(fi6))
            case 2147483647:
                p1 -= p1
            }
            v7 := GenS1{F0: 2147483648}
            var v8 bool = (!(82 > -1))
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint64 = 2
        _ = v5
        _ = v9
    } else {
        var v10 uint8 = 199
        v11 := GenS1{F0: 176}
        _ = v10
        _ = v11
    }
    var v12 uint64 = 4294967295
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 590
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{2147483648, 127, 2, 633}
    var v3 int64 = 220
    var v4 int32 = v1
    switch -2147483648 {
    case 128:
        if false {
            fmt.Println(v4, v1)
            v3 = 42
            var v14 uint16 = 65535
            _ = v14
        }
    }
    fmt.Println(v1, v4)
    v15 := GenS1{F0: 2147483647}
    if false {
        v16 := GenS2{F0: v3, F1: v3}
        v2 = append(v2, 687)
        _ = v16
    }
    fmt.Println(v1)
    v1 = v1
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int16 = 712
        fmt.Println(v18, v18, v18)
        if false {
            var v19 float32 = 1e10
            var v20 uint32 = 1
            _ = v19
            _ = v20
        } else {
            if true {
                for fi21 := 0; fi21 < 18; fi21++ {
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                    _ = fi21
                }
            }
            fmt.Println(v1, v3)
            v22 := GenS2{F0: int64(v4), F1: (v3 ^ v3)}
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    if (true && (158 >= 2147483647)) {
        _ = len(v2)
        var v23 string = "\n\t"
        switch 413 {
        case 32768:
            fmt.Println(v3)
        case -469:
            if false {
                for fi24 := 0; fi24 < 20; fi24++ {
                    if (!false) {
                        v25 := GenS3{F0: 1, F1: int32(fi24)}
                        v26 := GenS4{F0: 2}
                        _ = v25
                        _ = v26
                    }
                    if true {
                        switch (fi24 << ((977) & 63)) {
                        case -1:
                            var v27 int64 = (v3 | (^(v3)))
                            _ = v27
                        }
                        var v28 float64 = (1.5 / 1.0)
                        _ = v28
                    }
                    _ = fi24
                }
            } else {
                v4 = (v1 << ((2) & 63))
                fmt.Println(v1, v3, v1)
            }
        default:
            if false {
                for fi29 := 0; fi29 < 19; fi29++ {
                    var v30 int32 = v1
                    var v31 int = fi29
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            } else {
                v4 = -1
            }
        }
        _ = v23
    }
    switch "" {
    case "a":
        var v32 int8 = 42
        _ = v32
    case "":
        v33 := GenS5{F0: 0, F1: "", F2: 60}
        _ = v33
    case "\n\t":
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 uint8 = 42
            _ = fi34
            _ = v35
        }
    default:
        var v36 int16 = 170
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}

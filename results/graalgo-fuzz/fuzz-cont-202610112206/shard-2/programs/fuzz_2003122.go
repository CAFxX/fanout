package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int = 649
    var v2 uint64 = 42
    v3 := make(map[int]uint32)
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 6; fi5++ {
        v3[(-(fi5))] = uint32(v2)
        v1 -= -882
        v6 := []int{218, fi5, (fi5 & (-(v1))), fi5}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -647, F1: 65535}
    var v8 int64 = 9223372036854775807
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            switch 2147483647 {
            case 32768:
                if (true || (true || (1 != 26))) {
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v11 := []int{}
                        switch 42 {
                        case 42:
                            var v12 uint16 = 32767
                            _ = v12
                        default:
                            v13 := GenS1{F0: (v8 * -42), F1: 65535}
                            _ = v13
                        }
                        switch -1 {
                        case 1:
                            v14 := v3[1]
                            v4 = uint32(v2)
                            var v15 int = fi10
                            _ = v14
                            _ = v15
                        case 9223372036854775807:
                            v16 := GenS2{F0: -67}
                            _ = v16
                        case 0:
                            v17 := GenS3{F0: 0, F1: 86}
                            _ = v17
                        }
                        var v18 int64 = 1
                        var v19 uint64 = v2
                        _ = fi10
                        _ = v11
                        _ = v18
                        _ = v19
                    }
                }
            case -863:
                var v20 int32 = -1
                _ = v20
            case -669:
                fmt.Println(fi9, v2, v2)
            default:
                var v21 int64 = (v8 & -1)
                _ = v21
            }
            v22 := []int{(fi9 % ((v1) | 1)), v1, fi9}
            _ = fi9
            _ = v22
        }
    } else {
        switch v1 {
        case -310:
            switch 0 {
            case 627:
                v1 += v1
            }
        case 1:
            if (3.14159 > (0.1 * (1.5 + 3.14159))) {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v1 = (-(fi23))
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v25 := GenS2{F0: (v8 / ((v8) | 1))}
                        v2 = (v2 - (v2 & (v2 % ((v2) | 1))))
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
            }
        }
        var v26 uint32 = 694
        delete(v3, v1)
        _ = v26
    }
    fmt.Println(v4, v2, v2)
    var v27 uint = 4294967295
    v28 := []int{}
    var v29 uint64 = v2
    var v30 int16 = 1
    var v31 int16 = (v30 | 42)
    v8 += (v8 % (((^(v8))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
    _ = v8
}

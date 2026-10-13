package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint = 65535
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 15; fi5++ {
        if false {
            var v6 int64 = -9223372036854775808
            v4[2147483647] = 127
            _ = v6
        }
        var v7 int16 = int16(v1)
        v8 := v4[65535]
        _ = fi5
        _ = v7
        _ = v8
    }
    v9 := []int{42, -151, 50}
    v1 += 668
    if true {
        var v10 uint64 = 566
        if false {
            if false {
                v11 := GenS1{F0: v1, F1: -2147483648, F2: v2}
                v12 := GenS1{F0: (v1 >> ((v1) & 63)), F1: 50, F2: v2}
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v2 = "foo bar"
                    var v14 int8 = 127
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 int = (fi13 + (fi13 << (((v1 & 65535)) & 63)))
                        v17 := GenS2{F0: -32768, F1: v1}
                        fmt.Println(v2)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
            }
        } else {
            var v18 uint = uint(v10)
            fmt.Println(v1, v1, v2)
            _ = v18
        }
        _ = v10
    } else {
        var v19 uint = (v1 & (^(v1)))
        _ = v19
    }
    var v20 string = "\n\t"
    var v21 uint = v1
    v1 = 195
    v22 := GenS1{F0: (v1 & v21), F1: -1, F2: v20}
    if ((1e10 - 0.0) > 1.5) {
        v20 = v2
        for fi23 := 0; fi23 < 2; fi23++ {
            _ = len(v9)
            _ = fi23
        }
        var v24 int32 = 291
        _ = v24
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            v2 = ""
            var v26 int16 = -164
            if false {
                if (0 <= 1) {
                    v27 := GenS3{F0: 574}
                    _ = v27
                }
                v26 += (v26 & v26)
                var v28 int8 = 15
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        var v29 uint64 = 0
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v1, v20)
            _ = fi30
        }
        _ = v29
    }
    fmt.Println(v20, v2)
    switch -1 {
    case 200:
        fmt.Println(v21, v1, v21)
    default:
        v31 := GenS4{F0: -1, F1: 65535}
        _ = v31
    }
    var v32 uint8 = 128
    switch 32767 {
    case -342:
        v33 := v3[127]
        _ = v33
    case -853:
        v21 -= (v21 - 4294967295)
    case 32768:
        v34 := GenS2{F0: 32767, F1: 478}
        _ = v34
    }
    v21 += 255
    if (true && (1 != 0)) {
        v32 = v32
        var v35 uint = v21
        if (2147483647 > 2) {
            v9 = append(v9, 127)
            v21 = (^(v1))
        }
        _ = v35
    } else {
        v36 := GenS4{F0: 1, F1: 643}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}

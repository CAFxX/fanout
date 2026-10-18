package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() int32 {
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            v7 := GenS1{F0: 2147483647, F1: -1}
            _ = v7
        }
        var v8 int32 = 42
        _ = fi6
        _ = v8
    }
    _ = v5
    return 378
}
func genFn1(p0 string, p1 int8, p2 uint16) bool {
    v9 := []int{}
    var v10 float32 = (0.0 + 1.5)
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            v10 = ((3.14159 / float32(3.14159)) / (v10 / (((v10 - v10) - v10) / float32(-2.25))))
            _ = fi11
        }
        fmt.Println(p0, p0)
        v12 := GenS2{F0: 146}
        _ = v12
    } else {
        switch 626 {
        case 32768:
            p2 = p2
        case 42:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint = 1
                fmt.Println(p0, fi13)
                v9 = append(v9, fi13)
                _ = fi13
                _ = v14
            }
        }
        var v15 uint64 = 1
        var v16 uint8 = 0
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (!false)
}

func main() {
    var v1 int16 = -1
    v2 := []int{127, int(v1), 9223372036854775807, 32767}
    var v3 int64 = 42
    v4 := make(map[uint32]int)
    var v17 uint8 = 255
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 uint16 = 1
                v21 := GenS3{F0: uint64(v20), F1: -258, F2: 4294967295}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
    case "hello":
        if false {
            v4[65535] = 876
        } else {
            v4[0] = -9223372036854775808
            fmt.Println(v3, v3, v1)
            v22 := GenS4{F0: 65}
            _ = v22
        }
    case "":
        switch 32767 {
        case -640:
            v17 -= (v17 | (v17 >> ((2) & 63)))
        case 0:
            v1 = 975
        }
    }
    delete(v4, 52)
    for fi23 := 0; fi23 < 9; fi23++ {
        if (1 <= 42) {
            var v24 uint64 = 567
            _ = v24
        } else {
            v25 := GenS5{F0: 4294967295, F1: 71}
            v3 += v3
            _ = v25
        }
        v17 += (v17 % (((^(v17))) | 1))
        v1 -= (v1 | 32767)
        _ = fi23
    }
    var v26 uint16 = 255
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if true {
        v3 -= 578
    }
    fmt.Println(v3)
    if false {
        v27 := GenS5{F0: 173, F1: 769}
        _ = v27
    } else {
        var v28 uint32 = 0
        v26 = v26
        switch "a" {
        case "a":
            fmt.Println(v28)
        case "foo bar":
            for fi29 := 0; fi29 < 10; fi29++ {
                v17 += v17
                v30 := GenS2{F0: (^(v17))}
                var v31 uint32 = v28
                _ = fi29
                _ = v30
                _ = v31
            }
        case "\n\t":
            _ = len(v4)
        }
        _ = v28
    }
    v32 := GenS5{F0: 982, F1: 523}
    v26 -= 1
    v33 := GenS1{F0: int32(v1), F1: (v1 ^ v1)}
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}

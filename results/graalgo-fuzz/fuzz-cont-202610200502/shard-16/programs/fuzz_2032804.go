package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 0
    v3 := make(map[string]bool)
    var v4 int16 = -32768
    var v5 int32 = -984
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{}
        v4 += (v4 >> ((v2) & 63))
        _ = fi7
        _ = v8
    }
    if true {
        v9 := []int{-2147483648, 9223372036854775807, 127, -1}
        _ = v9
    } else {
        if true {
            v10 := GenS1{F0: 0, F1: 65535, F2: 65535}
            fmt.Println(v5, v4, v6)
            v11 := GenS2{F0: 32767, F1: 903, F2: 865}
            _ = v10
            _ = v11
        }
    }
    var v12 int32 = (v5 >> ((2) & 63))
    v13 := GenS3{F0: 32767, F1: v2, F2: 86}
    for fi14 := 0; fi14 < 13; fi14++ {
        switch fi14 {
        case 0:
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 int8 = 0
                var v17 int8 = (v16 << ((v2) & 63))
                v18 := []int{-1}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := []int{fi19}
        var v21 int16 = (v4 * -654)
        _ = fi19
        _ = v20
        _ = v21
    }
    if (true && false) {
        if (4294967295 == 1) {
            var v22 int8 = -23
            _ = v22
        }
        var v23 uint = (v2 - 0)
        _ = v23
    }
    if (int8(v5) < 42) {
        var v24 uint64 = v1
        switch int(v2) {
        case -929:
            v24 += 1
        case 924:
            v12 -= (v12 >> ((v2) & 63))
        case 32768:
            v5 = v12
        default:
            v25 := GenS3{F0: 65535, F1: v2, F2: 42}
            _ = v25
        }
        v24 = (^(v24))
        _ = v24
    } else {
        if false {
            v26 := []int{278, -365, 4294967295, 32768}
            _ = v26
        } else {
            if ((!(!(v1 != (v1 * 165)))) && (267 < 294)) {
                v2 += (v2 / ((v2) | 1))
            }
            var v27 int16 = 499
            for fi28 := 0; fi28 < 20; fi28++ {
                fmt.Println(v27, v27, v2)
                var v29 uint8 = 2
                var v30 int8 = -83
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
    }
    v4 += (v4 >> ((0) & 63))
    v12 += (v5 >> (((v2 ^ (v2 - (^(v2))))) & 63))
    if false {
        fmt.Println(v5, v4)
    } else {
        for fi31 := 0; fi31 < 4; fi31++ {
            v2 -= v2
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 string = (("hello" + "a") + "")
                delete(v3, v33)
                var v34 uint32 = 2147483647
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    switch -382 {
    case 42:
        delete(v3, "")
    case -644:
        if (false && ((v12 % ((-1) | 1)) >= v12)) {
            if (!false) {
                var v35 int = 128
                v5 = 1
                _ = v35
            } else {
                v36 := GenS4{F0: (v4 | -853)}
                v37 := GenS2{F0: (v4 * v4), F1: -1, F2: 4294967295}
                _ = v36
                _ = v37
            }
        }
    default:
        switch int(v2) {
        case 32767:
            var v38 uint64 = (v1 | v1)
            _ = v38
        case 127:
            for fi39 := 0; fi39 < 20; fi39++ {
                v6 = v6
                _ = fi39
            }
        default:
            v4 -= (v4 * v4)
        }
    }
    var v40 uint8 = 42
    v3[""] = true
    if ((v40 ^ (v40 & v40)) >= (^(v40))) {
        v4 -= (v4 ^ v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

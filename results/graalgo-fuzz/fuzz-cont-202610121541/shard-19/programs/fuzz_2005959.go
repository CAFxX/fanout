package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
func genFn0(p0 uint64, p1 uint64) uint8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int32 = 795
        p1 = p0
        _ = fi5
        _ = v6
    }
    if (235 < 9223372036854775807) {
        v7 := []int{-9223372036854775808, 0, 4294967295}
        _ = v7
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 1
    v3 := make(map[int]int)
    var v4 uint = 494
    v8 := GenS1{F0: false, F1: 0, F2: 204}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint32 = 281
            _ = fi10
            _ = v11
        }
        fmt.Println(v1)
        fmt.Println(fi9, v4)
        _ = fi9
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v4 += v4
        var v13 uint16 = 467
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            if ((!((v4 * v4) == 2)) && false) {
                v16 := []int{(^(fi14)), (fi15 - -103), -1}
                _ = v16
            }
            _ = fi15
        }
        v4 = 1
        v17 := GenS1{F0: (false && true), F1: 0, F2: (v2 - v2)}
        _ = fi14
        _ = v17
    }
    v1 -= v1
    v18 := GenS2{F0: (v1 & (v1 - v1)), F1: 18446744073709551615, F2: (("a" + "foo bar") > ("" + "foo bar"))}
    v19 := GenS2{F0: v1, F1: 2, F2: (934 >= 65535)}
    if true {
        var v20 string = "a"
        _ = v20
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS3{F0: 255, F1: 42, F2: 2147483648}
            var v23 uint16 = 1
            var v24 uint64 = 2
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        fmt.Println(v1, v1, v2)
        v1 -= v1
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            fmt.Println(v2, v1)
            _ = fi26
        }
        v27 := GenS2{F0: 328, F1: 18446744073709551615, F2: true}
        v28 := GenS4{F0: "foo bar", F1: (v4 + 2), F2: 1}
        _ = fi25
        _ = v27
        _ = v28
    }
    if false {
        v29 := []int{-2147483648, 32767}
        _ = v29
    } else {
        v1 = (v1 ^ v1)
    }
    if true {
        v30 := v3[32768]
        if true {
            var v31 int8 = -93
            fmt.Println(v4)
            v32 := GenS2{F0: 0, F1: 18446744073709551615, F2: true}
            _ = v31
            _ = v32
        } else {
            var v33 uint8 = 42
            _ = len(v3)
            if false {
                v2 -= v2
                delete(v3, (v30 >> ((v4) & 63)))
            }
            _ = v33
        }
        _ = v30
    }
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

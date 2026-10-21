package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint32; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 bool) uint16 {
    var v5 int = 2147483647
    var v6 int = (v5 - -1)
    if (!p0) {
        v7 := []int{v5, (v6 ^ (v6 * (-(v6)))), v5}
        v5 -= v5
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 968
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v8 uint8 = 194
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint16 = 255
        if true {
            if ((false && false) && false) {
                switch (fi9 + fi9) {
                case -1:
                    v8 -= (^(v8))
                case 791:
                    v11 := GenS1{F0: 2147483648, F1: 1}
                    _ = v11
                }
                v12 := []int{9223372036854775807, (fi9 & fi9), (fi9 | 32768), fi9}
                if len(v12) > 0 {
                    v12[0] = 32768
                }
                _ = v12
            } else {
                v13 := GenS1{F0: 828, F1: 2147483647}
                var v14 int64 = 469
                var v15 int16 = -82
                _ = v13
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: 2, F1: 1}
            _ = len(v4)
            _ = v16
        } else {
            var v17 int64 = 9223372036854775807
            if false {
                v17 += (-(v17))
            }
            _ = v17
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v8)
    switch "" {
    case "":
        var v18 int32 = 1
        _ = v18
    case "\n\t":
        v19 := GenS1{F0: 2147483647, F1: 1}
        _ = v19
    }
    var v20 uint32 = 433
    if false {
        if (936 > 0) {
            v8 -= v8
            v21 := []int{-738, -784, 1, 1}
            _ = v21
        }
        fmt.Println(v2, v20, v8)
        if (!(2147483647 >= 1)) {
            switch ("" + "\n\t") {
            case "\n\t":
                v22 := GenS1{F0: 322, F1: 2147483647}
                _ = v22
            default:
                v23 := []int{}
                _ = v23
            }
            v24 := []int{42, 43}
            v8 = 2
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
    }
    if true {
        v26 := []int{-799}
        _ = v26
    } else {
        v27 := GenS2{F0: (v20 << ((545) & 63)), F1: 32767, F2: (v20 > v20)}
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS1{F0: (^(v20)), F1: -1}
        var v30 uint32 = v20
        v31 := GenS1{F0: (v30 >> ((65535) & 63)), F1: 1}
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 bool = false
    v32 = (!v32)
    var v33 int32 = -2147483648
    var v34 uint32 = v20
    var v35 int8 = 1
    if (967 == 32767) {
        var v36 int32 = (v33 >> ((65535) & 63))
        _ = v36
    } else {
        fmt.Println(v33, v33, v32)
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        switch fi37 {
        case 1:
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS3{F0: false, F1: 42}
                _ = fi38
                _ = v39
            }
        case 9223372036854775807:
            for fi40 := 0; fi40 < 8; fi40++ {
                v32 = v32
                v41 := GenS4{F0: 1000, F1: uint64(v8)}
                v2 -= 3.14159
                _ = fi40
                _ = v41
            }
        default:
            var v42 int8 = v35
            _ = v42
        }
        var v43 uint64 = 1
        _ = fi37
        _ = v43
    }
    switch (("a" + "a") + "\n\t") {
    case "\n\t":
        for fi44 := 0; fi44 < 14; fi44++ {
            v45 := []int{}
            var v46 uint16 = 32767
            _ = len(v1)
            _ = fi44
            _ = v45
            _ = v46
        }
    case "a":
        v33 += (v33 | 1)
    }
    v47 := v1[-474]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v47
    _ = v8
}

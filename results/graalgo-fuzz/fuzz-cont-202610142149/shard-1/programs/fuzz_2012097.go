package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int64 }
func genFn0() bool {
    var v6 int32 = 1
    if (65535 >= 2) {
        v6 = (v6 ^ 2147483647)
        for fi7 := 0; fi7 < 1; fi7++ {
            v6 -= (^(v6))
            v8 := []int{}
            var v9 int8 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{}
        _ = v10
    }
    var v11 int8 = -57
    switch "" {
    case "a":
        v6 -= v6
    case "":
        v12 := GenS1{F0: 1000, F1: v6, F2: v11}
        _ = v12
    case "\n\t":
        v6 = (v6 << ((42) & 63))
    default:
        switch ("foo bar" + "foo bar") {
        case "hello":
            v13 := GenS1{F0: -1, F1: v6, F2: -80}
            _ = v13
        case "foo bar":
            v6 -= 42
        case "":
            if true {
                v11 += (^(v11))
                if (false || false) {
                    var v14 float32 = (3.14159 / 1.5)
                    v15 := GenS1{F0: 42, F1: (v6 & v6), F2: (^(v11))}
                    _ = v14
                    _ = v15
                } else {
                    if (0 != 0) {
                        var v16 int64 = -1
                        fmt.Println(v11)
                        _ = v16
                    } else {
                        var v17 int64 = 808
                        _ = v17
                    }
                }
            }
        }
    }
    _ = v11
    _ = v6
    return true
}

func main() {
    v1 := []uint32{4294967295, 42}
    v2 := []uint32{457, 2, 372, 24}
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v18 uint16 = 32767
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint16 = (^(v18))
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 896, F1: v18, F2: 1}
    v22 := GenS1{F0: -779, F1: 2147483647, F2: 42}
    switch "\n\t" {
    case "a":
        v18 += v18
    }
    fmt.Println(v18)
    v18 = 42
    v18 = 32767
    v23 := GenS2{F0: 907, F1: v18, F2: 65535}
    switch "" {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v18 = v18
    switch "a" {
    case "hello":
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 float32 = 1.5
            if true {
                if (!true) {
                    v18 = v18
                } else {
                    fmt.Println(v25, v18)
                    v26 := GenS1{F0: 42, F1: 0, F2: 100}
                    if false {
                        fmt.Println(v25)
                        v27 := GenS1{F0: 1000, F1: 0, F2: 100}
                        var v28 uint64 = 2
                        _ = v27
                        _ = v28
                    }
                    _ = v26
                }
            } else {
                if (!false) {
                    var v29 bool = true
                    v30 := GenS3{F0: 94, F1: 4294967295}
                    var v31 int32 = 2147483647
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    v18 = 65535
                    var v32 uint = 901
                    var v33 int16 = int16(v32)
                    _ = v32
                    _ = v33
                }
                var v34 int8 = 111
                _ = v34
            }
            v35 := GenS3{F0: -128, F1: 1}
            _ = fi24
            _ = v25
            _ = v35
        }
    case "a":
        v18 += (v18 ^ v18)
    case "\n\t":
        fmt.Println(v18)
    }
    if true {
        if (("foo bar" != "\n\t") || false) {
            v18 += v18
        } else {
            v18 += (v18 << ((662) & 63))
            v18 = v18
        }
    }
    if (42 >= 579) {
        var v36 int64 = 0
        if false {
            for fi37 := 0; fi37 < 3; fi37++ {
                v36 += v36
                v18 = (v18 - v18)
                v38 := GenS4{F0: v36}
                _ = fi37
                _ = v38
            }
            v36 = (v36 % ((964) | 1))
            v36 -= 9223372036854775807
        }
        var v39 uint = 65535
        _ = v36
        _ = v39
    } else {
        v40 := v4[127]
        _ = v40
    }
    var v41 int32 = 42
    var v42 int = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}

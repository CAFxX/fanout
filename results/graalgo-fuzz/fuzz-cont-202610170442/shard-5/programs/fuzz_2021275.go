package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint }
type GenS7 struct { F0 uint8; F1 int }
func genFn0(p0 bool) uint8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := []int{}
        switch "\n\t" {
        case "a":
            if p0 {
                if len(v5) > 0 {
                    v5[0] = (-(fi4))
                }
                fmt.Println(p0)
            } else {
                switch fi4 {
                case -328:
                    if len(v5) > 0 {
                        v5[0] = fi4
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "foo bar":
            if false {
                if (p0 || p0) {
                    for fi6 := 0; fi6 < 7; fi6++ {
                        p0 = true
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v7 := GenS1{F0: 54, F1: 952}
                        p0 = p0
                        _ = fi6
                        _ = v7
                    }
                    fmt.Println(fi4)
                }
                v8 := GenS1{F0: 42, F1: 0}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 7; fi9++ {
                    p0 = (((2 <= 778) || p0) || p0)
                    var v10 int = -9223372036854775808
                    _ = fi9
                    _ = v10
                }
            }
        }
        switch (-(fi4)) {
        case 0:
            v11 := GenS2{F0: -586, F1: 0}
            _ = v11
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                if (-645 < 0) {
                    v13 := GenS3{F0: ("foo bar" + "foo bar"), F1: 851, F2: fi12}
                    if p0 {
                        v14 := GenS4{F0: 1, F1: -497}
                        var v15 uint64 = 42
                        _ = len(v5)
                        _ = v14
                        _ = v15
                    }
                    _ = len(v5)
                    var v16 int = (fi4 ^ 32768)
                    var v17 uint = 4294967295
                    _ = v13
                    _ = v16
                    _ = v17
                }
                _ = fi12
            }
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    switch 4294967295 {
    case 128:
        var v18 uint = 0
        _ = v18
    default:
        var v19 uint32 = 907
        _ = v19
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(p0)
    case "hello":
        p0 = p0
    case "":
        var v20 float32 = ((1e10 - 1.5) / 1.0)
        _ = v20
    }
    _ = p0
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]bool)
    var v3 uint64 = 0
    v21 := v2[2]
    fmt.Println(v21)
    v22 := v2[929]
    fmt.Println(v22, v21)
    fmt.Println(v3, v22, v3)
    v23 := GenS2{F0: -2147483648, F1: 0}
    var v24 int64 = -1
    v25 := []int{-2147483648}
    if false {
        var v26 int64 = (v24 ^ (v24 << ((42) & 63)))
        v27 := GenS5{F0: (v26 >> ((661) & 63)), F1: 2}
        v22 = ((v3 - 261) >= (v3 / ((v3) | 1)))
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 float32 = (3.14159 + 1e10)
            _ = fi28
            _ = v29
        }
    }
    delete(v2, 2147483647)
    switch 9223372036854775807 {
    case -2147483648:
        switch -2147483648 {
        case 0:
            v30 := GenS4{F0: v1, F1: -462}
            _ = v30
        case 110:
            v22 = v22
        default:
            v1 = 664
        }
    default:
        v31 := v2[2]
        _ = v31
    }
    v32 := GenS3{F0: "\n\t", F1: 459, F2: 42}
    v33 := GenS6{F0: -1, F1: 2147483647, F2: 0}
    v34 := GenS7{F0: 2, F1: 9223372036854775807}
    var v35 bool = ((3.14159 / 1.0) != 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}

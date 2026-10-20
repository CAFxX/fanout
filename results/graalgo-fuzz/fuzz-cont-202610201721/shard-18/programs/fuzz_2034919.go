package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 bool }
func genFn0() int8 {
    var v7 uint = 65535
    var v8 uint8 = 226
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 0, F1: 255, F2: -32768}
        switch "" {
        case "foo bar":
            fmt.Println(v7, v8)
        case "\n\t":
            for fi11 := 0; fi11 < 11; fi11++ {
                v8 = v8
                v12 := GenS1{F0: 2, F1: 42, F2: 731}
                if true {
                    v8 = v8
                    var v13 float64 = -2.25
                    _ = v13
                } else {
                    v7 += v7
                    var v14 int8 = -67
                    _ = v14
                }
                var v15 int64 = -182
                _ = fi11
                _ = v12
                _ = v15
            }
        }
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 uint64) uint32 {
    var v16 float64 = 0.0
    v17 := GenS1{F0: p1, F1: 65535, F2: 459}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 int16 = -479
    var v4 float64 = 0.1
    var v5 uint8 = 113
    v6 := make(map[string]int)
    var v18 float64 = 0.1
    var v19 int = -923
    if false {
        if ("" <= "foo bar") {
            if true {
                switch (v19 ^ (v19 | (v19 - (v19 >> ((2) & 63))))) {
                case -9223372036854775808:
                    v4 = (v4 - 3.14159)
                case 135:
                    if false {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            _ = len(v2)
                            _ = fi20
                        }
                    }
                }
            }
            var v21 int8 = 1
            v22 := GenS2{F0: -2147483648}
            _ = v21
            _ = v22
        }
        var v23 uint32 = 1
        _ = v23
    } else {
        if false {
            _ = len(v2)
        }
        var v24 int8 = -99
        fmt.Println(v5, v24, v18)
        _ = v24
    }
    v19 = (-(v19))
    fmt.Println(v5, v19)
    switch "hello" {
    case "\n\t":
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint64 = 2
            v26 = v26
            if false {
                v5 = (^(v5))
            }
            _ = fi25
            _ = v26
        }
    case "foo bar":
        fmt.Println(v4, v4)
    }
    v27 := GenS3{F0: 42, F1: 0}
    if false {
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(fi29, v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi29
            }
            _ = fi28
        }
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        switch fi30 {
        case -643:
            v31 := GenS2{F0: 1}
            _ = v31
        }
        _ = fi30
    }
    v32 := GenS4{F0: 136, F1: 415, F2: (true || (1.5 < 3.14159))}
    v4 += (((v18 / (v4 * v18)) + (v4 + v18)) - (0.0 - v18))
    var v33 int64 = 891
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

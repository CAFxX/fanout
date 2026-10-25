package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    var v4 float64 = -2.25
    v3 -= (v3 * (^(v3)))
    var v5 int = 42
    fmt.Println(v5, v3, v5)
    var v6 uint8 = 128
    switch (v5 / ((v5) | 1)) {
    case 9223372036854775807:
        if (-854 >= 1) {
            if (!false) {
                v7 := GenS1{F0: 1, F1: 548, F2: (v5 & v5)}
                _ = v7
            }
        }
    default:
        v6 = (v6 >> ((411) & 63))
    }
    fmt.Println(v4, v4, v4)
    v8 := []int{v5, v5, (v5 + (v5 * 1)), 42}
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := v1[19]
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                if (!false) {
                    v4 += 0.0
                    v5 += (v5 % ((fi12) | 1))
                } else {
                    v5 = fi12
                }
                switch ("" + (("" + "hello") + "foo bar")) {
                case "":
                    v3 += uint32(v6)
                case "hello":
                    v13 := GenS2{F0: 1, F1: 1}
                    _ = v13
                case "\n\t":
                    var v14 uint = 274
                    _ = v14
                default:
                    v3 -= (v3 / ((v3) | 1))
                }
                if (v4 < ((0.1 + float64(3.14159)) * v4)) {
                    v6 = (v6 + 42)
                } else {
                    if true {
                        var v15 int = fi12
                        v16 := GenS1{F0: 2147483647, F1: 609, F2: (fi12 % ((fi12) | 1))}
                        v5 = 9223372036854775807
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int8 = -57
                        var v18 uint32 = v3
                        v18 += 2
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v4, v5, v3)
                    var v19 float64 = v4
                    v20 := v2[42]
                    _ = v19
                    _ = v20
                }
            }
            var v21 uint32 = v3
            var v22 uint32 = (^(v21))
            _ = fi12
            _ = v21
            _ = v22
        }
        _ = fi9
    }
    v23 := GenS2{F0: 0, F1: -402}
    var v24 int32 = 834
    fmt.Println(v3, v4)
    switch v5 {
    case -646:
        var v25 uint8 = 42
        _ = v25
    case 128:
        v26 := GenS1{F0: (v24 << ((65535) & 63)), F1: v24, F2: (-(v5))}
        _ = v26
    case -866:
        if true {
            var v27 bool = true
            for fi28 := 0; fi28 < 19; fi28++ {
                v6 = v6
                _ = fi28
            }
            _ = v27
        }
    default:
        v29 := GenS1{F0: (v24 & (-(v24))), F1: v24, F2: (v5 % ((128) | 1))}
        _ = v29
    }
    if true {
        v8 = append(v8, (v5 | 1))
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 float32 = 0.0
            fmt.Println(v6, v31, v24)
            var v32 uint32 = (v3 << ((65535) & 63))
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    _ = len(v1)
    switch ((("hello" + "foo bar") + "a") + "foo bar") {
    case "hello":
        v24 = -2147483648
    case "foo bar":
        v3 = v3
    case "":
        fmt.Println(v24)
    }
    fmt.Println(v24, v6)
    for fi33 := 0; fi33 < 3; fi33++ {
        if true {
            var v34 bool = false
            v35 := GenS1{F0: v24, F1: -957, F2: 0}
            _ = v34
            _ = v35
        } else {
            v24 = 2147483647
            v24 += v24
            switch "a" {
            case "a":
                v36 := GenS1{F0: v24, F1: (v24 >> ((646) & 63)), F2: (-(fi33))}
                _ = v36
            case "\n\t":
                switch "hello" {
                case "a":
                    var v37 int16 = 1000
                    _ = v37
                case "\n\t":
                    v38 := GenS3{F0: false}
                    _ = v38
                default:
                    v3 = v3
                }
            case "foo bar":
                if false {
                    var v39 int = v5
                    if true {
                        fmt.Println(v3)
                        v40 := GenS4{F0: 2}
                        fmt.Println(v4, fi33)
                        _ = v40
                    }
                    _ = v39
                }
            }
        }
        v24 -= -1
        for fi41 := 0; fi41 < 10; fi41++ {
            v24 = 2147483647
            fmt.Println(v24)
            var v42 int = (fi41 * (fi33 ^ v5))
            var v43 int = (fi41 | (v5 & 102))
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi33
    }
    fmt.Println(v5, v4, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

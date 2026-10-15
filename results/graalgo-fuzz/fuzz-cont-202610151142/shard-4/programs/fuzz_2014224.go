package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint8) int64 {
    var v6 uint8 = 1
    var v7 bool = true
    v8 := GenS1{F0: 1}
    v9 := GenS1{F0: 0}
    p0 -= 0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 425
}

func main() {
    v1 := []string{"a", (("foo bar" + "a") + "a"), ("a" + "\n\t")}
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 uint8 = 42
    _ = len(v3)
    v1 = append(v1, "a")
    var v10 uint64 = 315
    var v11 string = ("foo bar" + "a")
    var v12 int32 = 39
    var v13 int32 = 42
    v2 = v2
    v5 = (v5 ^ v5)
    v2 -= v2
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: v10}
        _ = fi14
        _ = v15
    }
    switch int(v10) {
    case -497:
        if true {
            if true {
                v10 = v10
                if (v11 > v11) {
                    v2 += v2
                    v16 := GenS3{F0: v12, F1: (^(v5)), F2: 444}
                    if true {
                        switch "\n\t" {
                        case "hello":
                            v12 = (^(v12))
                        case "a":
                            var v17 int32 = v12
                            _ = v17
                        }
                        fmt.Println(v10, v12, v10)
                        v12 = (^(v12))
                    }
                    _ = v16
                } else {
                    var v18 float32 = float32(v2)
                    _ = v18
                }
                if true {
                    switch ((v11 + v11) + ("hello" + (v11 + "hello"))) {
                    case "\n\t":
                        v4 = append(v4, 1)
                    case "hello":
                        var v19 int16 = 42
                        _ = v19
                    }
                    fmt.Println(v5, v2)
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v5 = (^(v5))
                        v13 -= (-(v12))
                        _ = fi20
                    }
                } else {
                    v10 -= (v10 >> ((255) & 63))
                    switch 32767 {
                    case -9223372036854775808:
                        v21 := GenS4{F0: 1, F1: 42}
                        _ = v21
                    }
                }
            }
            fmt.Println(v10)
            if (v11 > (v11 + ((v11 + v11) + v11))) {
                v5 = (v5 << ((2) & 63))
                if (0 != 221) {
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    v2 = -2.25
                    var v22 uint8 = (^(v5))
                    _ = v22
                } else {
                    var v23 uint32 = 370
                    _ = v23
                }
            }
        }
    case 127:
        var v24 float64 = v2
        _ = v24
    case -241:
        var v25 int8 = -1
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

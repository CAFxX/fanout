package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 bool; F1 int; F2 int32 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 uint64 = 278
    var v4 int32 = -587
    var v5 int = 32767
    var v6 int32 = v4
    var v7 float64 = 1.5
    v8 := GenS1{F0: 42}
    if (((v3 * v3) > v3) && (887 >= 84)) {
        var v9 int32 = (^(v6))
        if (!false) {
            var v10 int16 = -32768
            _ = v10
        }
        v3 += v3
        _ = v9
    }
    fmt.Println(v6, v4)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint = 0
        fmt.Println(v5, v3, v5)
        if true {
            var v13 uint64 = v3
            _ = v13
        } else {
            v14 := GenS2{F0: (v3 + (v3 % ((v3) | 1)))}
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 uint64 = v3
    v16 := GenS3{F0: -1, F1: v5}
    v3 = (v15 << ((576) & 63))
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(v3, v6)
            var v19 uint64 = (v3 * v15)
            var v20 float64 = v7
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    var v21 uint = 4294967295
    if (false || (1.5 <= (((((0.1 - 0.1) - 1.5) * 1.5) - 1.5) * 0.0))) {
        v7 -= (v7 + v7)
        v22 := GenS4{F0: (false || (!false)), F1: -1, F2: v4}
        v5 = v5
        _ = v22
    } else {
        if (!(!("a" == "\n\t"))) {
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 5; fi24++ {
                    if true {
                        v4 += v6
                        switch "hello" {
                        case "a":
                            v7 -= (3.14159 + 3.14159)
                        case "\n\t":
                            v7 = v7
                        default:
                            var v25 float64 = (float64(1.5) / (v7 + (float64(1e10) + (v7 + (v7 - 1.5)))))
                            v26 := GenS5{F0: 199, F1: 0}
                            _ = len(v2)
                            _ = v25
                            _ = v26
                        }
                    }
                    v27 := GenS6{F0: 42}
                    switch fi23 {
                    case 172:
                        fmt.Println(v7, v21)
                    }
                    _ = fi24
                    _ = v27
                }
                for fi28 := 0; fi28 < 19; fi28++ {
                    v6 -= (v6 ^ (-(v6)))
                    if true {
                        v4 += v6
                        var v29 uint8 = uint8(v3)
                        switch v5 {
                        case 128:
                            var v30 int8 = 1
                            _ = v30
                        case 32768:
                            v31 := GenS2{F0: v15}
                            _ = v31
                        }
                        _ = v29
                    } else {
                        var v32 int32 = (v6 + (^(v4)))
                        v33 := GenS7{F0: fi23}
                        fmt.Println(v15, v32, v15)
                        _ = v32
                        _ = v33
                    }
                    v5 = (fi28 - fi28)
                    v34 := GenS8{F0: 42, F1: (^(fi23))}
                    var v35 string = "a"
                    var v36 float64 = 1.5
                    v36 -= 3.14159
                    _ = fi28
                    _ = v34
                    _ = v35
                    _ = v36
                }
                _ = fi23
            }
        } else {
            var v37 float32 = 3.14159
            v6 += -1
            _ = v37
        }
        var v38 uint64 = v15
        _ = v38
    }
    fmt.Println(v4, v7, v7)
    v3 += (v15 + (v3 << ((1) & 63)))
    _ = len(v1)
    v6 -= (^(v4))
    v39 := GenS4{F0: (!true), F1: -61, F2: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -718
    var v3 int = 506
    v4 := make(map[string]int)
    var v5 uint32 = 4294967295
    fmt.Println(v2, v3, v3)
    v6 := GenS1{F0: v2, F1: (v2 ^ v2), F2: -258}
    v7 := []int{v3, (v3 << ((0) & 63))}
    v5 -= (v5 / (((v5 & (v5 | (v5 | 1)))) | 1))
    switch (v3 + -9223372036854775808) {
    case 0:
        fmt.Println(v5)
    case -199:
        var v8 float32 = (0.1 - ((0.0 * 1.5) * (1.5 + 0.0)))
        _ = v8
    }
    v9 := GenS2{F0: 869, F1: 128, F2: "foo bar"}
    if true {
        if len(v7) > 0 {
            v7[0] = -492
        }
        var v10 int16 = 32767
        if true {
            fmt.Println(v3, v2)
            if false {
                var v11 bool = ((v2 << ((1) & 63)) >= v2)
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 int16 = v10
                    var v15 uint64 = 290
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int = (v3 & (-(fi12)))
                    var v18 int32 = -1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v10 = 35
                _ = fi12
            }
        }
        _ = v10
    } else {
        var v19 uint32 = v5
        var v20 int8 = -1
        var v21 float32 = 1.5
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: -10, F1: 191, F2: "hello"}
    v5 = 1
    if false {
        if ((((-2.25 / float32(1e10)) * float32((3.14159 - (1.5 * float64(((3.14159 + -2.25) + 0.1)))))) / 1.0) <= (float32(float64(((3.14159 / (0.0 + -2.25)) + 1e10))) + (3.14159 / (0.1 - 1.5)))) {
            v5 += (v5 / ((65535) | 1))
            for fi23 := 0; fi23 < 11; fi23++ {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v5 = 65535
                    switch "\n\t" {
                    case "":
                        var v25 float64 = (float64(float32(1e10)) / 1e10)
                        _ = v25
                    case "a":
                        delete(v1, 2147483647)
                    }
                    v26 := GenS2{F0: -219, F1: 95, F2: (("hello" + "hello") + ("\n\t" + "hello"))}
                    _ = fi24
                    _ = v26
                }
                v2 += v2
                v27 := GenS1{F0: (v2 / ((-1) | 1)), F1: v2, F2: 32767}
                _ = fi23
                _ = v27
            }
            switch 127 {
            case 0:
                for fi28 := 0; fi28 < 5; fi28++ {
                    var v29 float32 = (0.1 / (0.0 + (1.5 + (3.14159 + 0.1))))
                    _ = fi28
                    _ = v29
                }
            case 32767:
                for fi30 := 0; fi30 < 17; fi30++ {
                    v31 := GenS1{F0: v2, F1: 1, F2: -32768}
                    for fi32 := 0; fi32 < 19; fi32++ {
                        if (float32(0.0) != 0.0) {
                            if len(v7) > 0 {
                                v7[0] = fi32
                            }
                        }
                        fmt.Println(v2)
                        _ = fi32
                    }
                    _ = fi30
                    _ = v31
                }
            }
        }
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 float64 = ((float64(float32((1e10 + 3.14159))) / 1.0) - 0.0)
                v36 := GenS2{F0: -109, F1: 189, F2: ("hello" + "foo bar")}
                _ = fi34
                _ = v35
                _ = v36
            }
            if true {
                _ = len(v7)
                v3 = 32767
                var v37 int32 = 641
                _ = v37
            }
            _ = fi33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

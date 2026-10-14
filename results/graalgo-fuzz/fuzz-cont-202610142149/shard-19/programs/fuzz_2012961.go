package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 uint32 = 1
    v5 := make(map[string]uint32)
    if true {
        v6 := GenS1{F0: v4, F1: 302, F2: int16(v4)}
        v4 += v4
        v4 -= (^(v4))
        _ = v6
    } else {
        if (true || false) {
            fmt.Println(v4, v1)
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = -466
            v4 -= (^(v4))
            _ = fi7
            _ = v8
        }
    }
    v4 = 65535
    v9 := v3[127]
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v9, fi10)
        var v11 int8 = -128
        _ = fi10
        _ = v11
    }
    var v12 int16 = -1
    switch 32767 {
    case -324:
        if (int32(v12) == 42) {
            for fi13 := 0; fi13 < 15; fi13++ {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v12 = v12
                    v1 += 255
                    if false {
                        fmt.Println(v1, v4, v4)
                        var v15 uint8 = 127
                        v1 -= (v15 & (^(v15)))
                        _ = v15
                    } else {
                        v16 := GenS1{F0: uint32(v1), F1: 35, F2: -32768}
                        _ = v16
                    }
                    _ = fi14
                }
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 float32 = 0.1
                    _ = fi17
                    _ = v18
                }
                _ = fi13
            }
        } else {
            var v19 uint32 = v9
            _ = v19
        }
    case 32767:
        fmt.Println(v4, v12)
    case -1:
        v20 := []int{-829, 640, 1}
        _ = v20
    }
    var v21 int32 = 1
    if ((!(!(3.14159 < 3.14159))) || false) {
        var v22 string = "foo bar"
        _ = v22
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS2{F0: v21, F1: (fi23 + (fi24 << ((42) & 63)))}
                v21 += (v21 / (((v21 - (v21 | (v21 % ((-1) | 1))))) | 1))
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        if (("hello" + "\n\t") >= "") {
            if true {
                v21 = v21
            } else {
                var v26 uint8 = 0
                v1 -= (v26 - 63)
                v9 = (v9 << ((65535) & 63))
                _ = v26
            }
            if true {
                v27 := []int{180, 2147483647}
                v12 = v12
                _ = v27
            } else {
                v9 = v4
                v4 += v9
            }
            switch 490 {
            case 721:
                fmt.Println(v4, v21)
            case 32767:
                v1 = 42
            case -2147483648:
                var v28 int64 = -1
                _ = v28
            }
        }
        var v29 uint64 = 4294967295
        _ = v29
    }
    v30 := GenS3{F0: v12, F1: v12}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}

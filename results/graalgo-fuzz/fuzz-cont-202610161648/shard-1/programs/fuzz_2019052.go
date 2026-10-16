package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 int32 = -2147483648
    var v5 uint64 = 42
    var v6 int16 = 1000
    v7 := []int{-170, 32768}
    var v8 string = ""
    v6 += (v6 / (((^(v6))) | 1))
    if (true || (2 <= 0)) {
        switch 32767 {
        case 187:
            v1 += 65535
        }
    } else {
        fmt.Println(v1)
        if ((!true) && true) {
            var v9 uint = 255
            switch -9223372036854775808 {
            case 2147483647:
                for fi10 := 0; fi10 < 19; fi10++ {
                    if false {
                        _ = len(v3)
                    }
                    v4 = (v4 + (-(v4)))
                    if (int64(fi10) != -9223372036854775808) {
                        v11 := GenS1{F0: (fi10 << ((v1) & 63)), F1: 2147483647, F2: v6}
                        var v12 uint = v1
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 uint64 = (v5 - v5)
                            _ = fi13
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
            }
            if (65535 >= 1) {
                _ = len(v7)
                if (((!(255 == 58)) && (2 > 0)) || false) {
                    var v15 int64 = 0
                    v9 = (v9 / ((65535) | 1))
                    _ = v15
                } else {
                    v9 -= (^(v9))
                }
                if false {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        fmt.Println(v5, v5, v4)
                        v9 = 117
                        if false {
                            v17 := GenS1{F0: (^(fi16)), F1: 2147483647, F2: (^(v6))}
                            fmt.Println(v6)
                            _ = v17
                        } else {
                            var v18 int64 = 1
                            _ = v18
                        }
                        _ = fi16
                    }
                }
            } else {
                var v19 float32 = (((-2.25 * (((-2.25 / ((((1e10 * float32((((float64(0.1) / 1.0) - 3.14159) / 1e10))) / 1.0) - 0.0) / float32(-2.25))) * 3.14159) + 0.0)) + (float32((((1.5 * ((0.0 - 1e10) * (float64(float32(0.1)) + (3.14159 - 1.5)))) / float64(3.14159)) * 3.14159)) * (-2.25 - 3.14159))) - 3.14159)
                v20 := GenS2{F0: 2, F1: 24, F2: (v5 ^ 42)}
                _ = v19
                _ = v20
            }
            _ = v9
        } else {
            var v21 uint = 434
            v1 = v21
            _ = v21
        }
        fmt.Println(v6)
    }
    v7 = append(v7, -1)
    v5 = 359
    switch (((("hello" + v8) + v8) + "hello") + v8) {
    case "foo bar":
        var v22 int16 = (v6 + v6)
        _ = v22
    case "":
        v6 -= 284
    }
    v8 = v8
    var v23 bool = (2147483647 != 127)
    var v24 uint64 = (v5 >> ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

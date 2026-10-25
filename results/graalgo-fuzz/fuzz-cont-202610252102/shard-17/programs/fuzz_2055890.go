package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    var v4 uint64 = 4294967295
    if (("a" <= "a") || (!((v4 * (^(v4))) <= 0))) {
        var v5 uint16 = 606
        _ = v5
    }
    if false {
        v6 := GenS1{F0: 100, F1: (v2 & -1), F2: 1}
        _ = v6
    }
    var v7 uint32 = 1
    var v8 uint8 = 255
    v8 += v8
    v9 := GenS2{F0: (v2 >> ((65535) & 63)), F1: 978}
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            if (-645 < -1) {
                if (!true) {
                    v12 := GenS1{F0: -128, F1: (v2 % ((32767) | 1)), F2: v8}
                    v13 := GenS3{F0: v4, F1: -970}
                    _ = v12
                    _ = v13
                } else {
                    var v14 bool = ((v7 | uint32(v4)) != (v7 + v7))
                    fmt.Println(v14, fi11)
                    if len(v1) > 0 {
                        v1[0] = (fi11 >> ((18446744073709551615) & 63))
                    }
                    var v15 uint16 = 2
                    _ = v14
                    _ = v15
                }
                if false {
                    v3[("hello" + "\n\t")] = true
                    var v16 int32 = -1
                    _ = v16
                } else {
                    v17 := GenS4{F0: v4, F1: (v2 | -622)}
                    delete(v3, "foo bar")
                    _ = v17
                }
                _ = len(v3)
                var v18 uint = 72
                _ = v18
            }
            switch ("hello" + "a") {
            case "foo bar":
                switch fi11 {
                case -474:
                    fmt.Println(v4, v4, v7)
                }
            }
            _ = fi11
        }
        v19 := GenS1{F0: 127, F1: v2, F2: v8}
        v20 := GenS3{F0: 0, F1: -1}
        _ = fi10
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS1{F0: 8, F1: (v2 / ((v2) | 1)), F2: (v8 * (v8 & v8))}
        _ = fi21
        _ = v22
    }
    var v23 uint8 = (v8 & v8)
    var v24 float64 = (0.1 - (1e10 - 3.14159))
    v25 := GenS4{F0: v4, F1: (v2 ^ v2)}
    fmt.Println(v4)
    v26 := GenS2{F0: 0, F1: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

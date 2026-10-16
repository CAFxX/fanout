package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64; F2 bool }
type GenS4 struct { F0 uint; F1 string; F2 uint }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 690
    var v3 float32 = 3.14159
    v4 := []int{-2147483648, 913, -734}
    v5 := GenS1{F0: 255}
    v2 = (v2 >> ((4294967295) & 63))
    var v6 int16 = -1
    v2 = v2
    var v7 int = 2147483647
    var v8 uint64 = 1
    if len(v4) > 0 {
        v4[0] = v7
    }
    v9 := GenS2{F0: 22}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                v12 := GenS2{F0: 880}
                fmt.Println(v3)
                if true {
                    if true {
                        v13 := GenS3{F0: v2, F1: -9223372036854775808, F2: ((false && (!true)) || (42 > -2))}
                        _ = v13
                    }
                }
                fmt.Println(v3, v6, v1)
                v14 := GenS4{F0: (^(v1)), F1: "\n\t", F2: 0}
                if ((v3 / ((v3 + v3) / (0.1 + v3))) > 1.5) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                fmt.Println(fi10)
                v8 += 2
                if (v1 < v1) {
                    v15 := GenS5{F0: int64(v7)}
                    _ = v15
                }
                var v16 uint = 0
                var v17 float64 = ((float64(0.1) - ((3.14159 / -2.25) / ((((1e10 + (1e10 - 1e10)) * 1.5) * 3.14159) + -2.25))) - float64(v3))
                _ = v12
                _ = v14
                _ = v16
                _ = v17
            }
            fmt.Println(v3, v2, fi10)
            if true {
                var v18 uint16 = 2
                if ((0.0 - 0.0) == (v3 / float32((1.5 + 1.5)))) {
                    if ((v1 | (^(v1))) == v1) {
                        if len(v4) > 0 {
                            v4[0] = -881
                        }
                    } else {
                        var v19 uint8 = 127
                        var v20 float64 = 0.1
                        _ = v19
                        _ = v20
                    }
                    switch (v7 >> ((uint(v8)) & 63)) {
                    case 32767:
                        var v21 uint8 = 1
                        _ = v21
                    case 2147483647:
                        if ((v8 - (^(v8))) == v8) {
                            var v22 int64 = 1
                            _ = v22
                        }
                    case 534:
                        var v23 bool = true
                        _ = v23
                    }
                    var v24 int32 = 1
                    _ = v24
                }
                _ = v18
            }
            _ = fi11
        }
        var v25 uint32 = 2
        _ = fi10
        _ = v25
    }
    v6 += int16(v8)
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            switch 42 {
            case -9223372036854775808:
                v4 = append(v4, (fi27 & fi27))
            case -69:
                fmt.Println(fi27, v6)
                var v28 float32 = float32(((float64((v3 * 3.14159)) - 3.14159) * float64((-2.25 + (v3 + 1e10)))))
                _ = v28
            default:
                var v29 int32 = 510
                _ = v29
            }
            var v30 string = "foo bar"
            var v31 bool = ((125 > 255) && false)
            _ = fi27
            _ = v30
            _ = v31
        }
        v4 = append(v4, fi26)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

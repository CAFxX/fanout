package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = -1
    var v2 uint8 = 94
    var v3 float64 = -2.25
    v4 := []int{-417, -737, 127, 1}
    v5 := GenS1{F0: 512}
    v4 = append(v4, 2147483647)
    var v6 uint64 = 18446744073709551615
    var v7 uint8 = v2
    v7 -= v7
    switch -1 {
    case -1:
        var v8 int32 = 42
        _ = v8
    case -810:
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS2{F0: 362}
                _ = fi10
                _ = v11
            }
            switch (fi9 / ((fi9) | 1)) {
            case -1:
                v7 = (v7 | uint8(v6))
            case -769:
                _ = len(v4)
            default:
                var v12 int8 = 42
                _ = v12
            }
            v7 = v7
            _ = fi9
        }
    case -2147483648:
        v2 += 255
    default:
        v2 -= v2
    }
    if (true || false) {
        var v13 uint8 = (v7 % (((^(v2))) | 1))
        fmt.Println(v1, v1)
        v13 = 0
        _ = v13
    } else {
        v1 = (v1 | v1)
    }
    var v14 float64 = 3.14159
    switch -303 {
    case -9223372036854775808:
        v2 -= v2
    }
    var v15 float32 = 0.0
    fmt.Println(v14)
    if true {
        var v16 uint = 40
        if (!(-2147483648 <= 127)) {
            fmt.Println(v7)
            var v17 uint8 = v7
            if false {
                var v18 int16 = (v1 | (v1 ^ v1))
                for fi19 := 0; fi19 < 7; fi19++ {
                    fmt.Println(v14)
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if len(v4) > 0 {
                            v4[0] = 32768
                        }
                        var v21 int64 = -9223372036854775808
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v14 += v14
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                if len(v4) > 0 {
                    v4[0] = 645
                }
                _ = v18
            } else {
                if ((v16 + v16) == 2) {
                    v22 := GenS2{F0: v16}
                    var v23 int = 2147483647
                    fmt.Println(v2)
                    _ = v22
                    _ = v23
                }
            }
            _ = v17
        } else {
            v6 += v6
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v16
    } else {
        v6 = (v6 ^ (^(v6)))
    }
    var v24 uint32 = 65535
    v25 := GenS1{F0: (^(v24))}
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    default:
        v24 = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

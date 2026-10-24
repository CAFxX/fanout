package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 0
    var v3 int64 = -9223372036854775808
    v4 := []uint32{(^(v2))}
    v5 := make(map[uint32]int)
    var v6 float64 = v1
    var v7 string = "foo bar"
    if (127 >= -49) {
        v7 = v7
        if ((1 == -873) || (1 > 255)) {
            if true {
                switch -453 {
                case -9223372036854775808:
                    v8 := GenS1{F0: v2, F1: (v2 % (((v2 ^ v2)) | 1))}
                    _ = v8
                case -211:
                    var v9 int32 = -283
                    _ = v9
                default:
                    var v10 uint64 = 0
                    _ = v10
                }
                var v11 int16 = -913
                _ = v11
            }
            if true {
                v12 := GenS1{F0: 2147483647, F1: v2}
                var v13 int = -2147483648
                var v14 uint64 = 579
                _ = v12
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: v2, F1: v2}
                _ = v15
            }
        }
    }
    if true {
        switch 4294967295 {
        case -264:
            var v16 uint8 = 0
            _ = v16
        default:
            v17 := GenS1{F0: (v2 | (v2 >> ((777) & 63))), F1: (^(v2))}
            _ = v17
        }
    }
    var v18 bool = false
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int8 = 0
        v6 += float64((1.5 + 1.5))
        var v21 int16 = -32768
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v1)
    v22 := GenS1{F0: v2, F1: (v2 + v2)}
    var v23 uint16 = 42
    switch v7 {
    case "foo bar":
        v1 = 3.14159
    case "a":
        v24 := GenS2{F0: "foo bar", F1: 215}
        _ = v24
    }
    switch v7 {
    case "\n\t":
        v5[(^(v2))] = 42
    default:
        v25 := GenS1{F0: v2, F1: v2}
        _ = v25
    }
    var v26 bool = v18
    var v27 uint8 = 127
    if v18 {
        if ((378 < v23) || v26) {
            switch 4294967295 {
            case 9223372036854775807:
                var v28 int16 = -32768
                _ = v28
            case 486:
                var v29 int64 = 765
                _ = v29
            case 32767:
                fmt.Println(v6, v2)
            default:
                var v30 bool = v26
                _ = v30
            }
            var v31 int32 = -701
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v31
        } else {
            v32 := GenS3{F0: -9223372036854775808, F1: v23}
            switch 1 {
            case 128:
                for fi33 := 0; fi33 < 18; fi33++ {
                    if false {
                        v1 = (-2.25 + 3.14159)
                    }
                    v26 = ((v3 / ((9223372036854775807) | 1)) != (v3 - v3))
                    _ = fi33
                }
            case -468:
                v23 = (v23 | v23)
            }
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

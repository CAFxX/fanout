package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16) bool {
    v5 := []int{458}
    switch -868 {
    case 4294967295:
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 string = "\n\t"
            for fi8 := 0; fi8 < 17; fi8++ {
                v7 = (v7 + "\n\t")
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    default:
        var v9 int8 = 1
        _ = v9
    }
    var v10 int32 = -29
    if false {
        if len(v5) > 0 {
            v5[0] = 546
        }
    }
    _ = p0
    _ = v10
    _ = v5
    return false
}

func main() {
    var v1 int = -502
    var v2 int8 = 127
    var v3 float64 = float64(-2.25)
    var v4 int64 = 42
    var v11 uint = 2
    var v12 int32 = -318
    var v13 string = "foo bar"
    v14 := []int{}
    v15 := GenS1{F0: 2, F1: true, F2: 108}
    if (uint32(v11) <= 42) {
        if (1 <= 0) {
            v4 -= -9223372036854775808
            v16 := GenS1{F0: 40, F1: false, F2: 42}
            v17 := GenS2{F0: -4, F1: 255}
            _ = v16
            _ = v17
        } else {
            v14 = append(v14, (v1 % (((v1 * (^(v1)))) | 1)))
            if false {
                if (0 > (v2 - -1)) {
                    if len(v14) > 0 {
                        v14[0] = v1
                    }
                    var v18 string = v13
                    _ = v18
                } else {
                    _ = len(v14)
                    fmt.Println(v11)
                }
                fmt.Println(v13, v11)
                v11 -= v11
            }
        }
    }
    if false {
        if false {
            v2 = (v2 >> ((290) & 63))
            v14 = append(v14, v1)
            v19 := GenS3{F0: v12, F1: 1, F2: 2147483647}
            _ = v19
        } else {
            var v20 int64 = (v4 + -630)
            _ = v20
        }
        v21 := GenS4{F0: v2, F1: 69}
        _ = v21
    }
    if ("a" >= v13) {
        if ((-(v2)) != -128) {
            v22 := GenS4{F0: 100, F1: 65535}
            _ = v22
        } else {
            v23 := GenS5{F0: 179}
            v4 += (v4 & (v4 / ((v4) | 1)))
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v11 += v11
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi24
        }
    } else {
        fmt.Println(v2)
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS1{F0: 105, F1: false, F2: 42}
            var v28 int16 = -121
            _ = fi26
            _ = v27
            _ = v28
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi25
    }
    switch v13 {
    case "foo bar":
        var v29 bool = (77 == 2)
        _ = v29
    case "":
        var v30 uint = (v11 ^ 42)
        _ = v30
    case "\n\t":
        v31 := GenS3{F0: (v12 | (v12 & (v12 | v12))), F1: 0, F2: 84}
        _ = v31
    default:
        v1 -= v1
    }
    switch v13 {
    case "a":
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 float32 = float32((v3 + -2.25))
            _ = fi32
            _ = v33
        }
    }
    if true {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v4 += -1
    } else {
        v14 = append(v14, 4294967295)
        v3 += 3.14159
        v13 = ((v13 + (((("a" + v13) + ("foo bar" + v13)) + v13) + ((v13 + "\n\t") + (v13 + v13)))) + "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}

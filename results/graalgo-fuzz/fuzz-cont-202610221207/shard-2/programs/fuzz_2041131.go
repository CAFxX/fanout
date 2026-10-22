package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16) uint32 {
    v4 := GenS1{F0: -36, F1: 1}
    for fi5 := 0; fi5 < 13; fi5++ {
        p0 -= 189
        if true {
            var v6 uint = 42
            _ = v6
        } else {
            var v7 float64 = -2.25
            if (789 >= 4294967295) {
                var v8 uint16 = p0
                var v9 string = ("foo bar" + "foo bar")
                _ = v8
                _ = v9
            }
            switch "a" {
            case "a":
                v10 := GenS2{F0: (fi5 & (fi5 & 9223372036854775807)), F1: 18446744073709551615, F2: (4294967295 < 42)}
                _ = v10
            }
            _ = v7
        }
        v11 := GenS2{F0: fi5, F1: 2, F2: (394 < (fi5 + fi5))}
        _ = fi5
        _ = v11
    }
    var v12 int32 = -1
    _ = p0
    _ = v12
    _ = v4
    return 65535
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 1
    v3 := make(map[int]string)
    var v13 uint32 = (v2 / ((1) | 1))
    if (42 <= 2) {
        v13 -= (v13 * (v13 % ((65535) | 1)))
        v1 = v1
    } else {
        if true {
            if true {
                v14 := v3[-552]
                if false {
                    v15 := GenS2{F0: -2147483648, F1: 2, F2: (true && false)}
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v1 += 62
                        _ = fi16
                    }
                    var v17 int = 2147483647
                    _ = v15
                    _ = v17
                } else {
                    v18 := []int{1, -72}
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if len(v18) > 0 {
                            v18[0] = -868
                        }
                        _ = fi19
                    }
                    switch (v14 + v14) {
                    case "\n\t":
                        v1 = (v1 - (^(v1)))
                    case "":
                        if len(v18) > 0 {
                            v18[0] = -1
                        }
                    default:
                        var v20 int32 = 42
                        _ = v20
                    }
                    _ = v18
                }
                v13 -= (v13 * (v13 % (((v2 + 4294967295)) | 1)))
                _ = v14
            } else {
                var v21 uint = 893
                _ = v21
            }
            v13 -= (v2 % ((genFn0(255)) | 1))
        } else {
            v22 := GenS2{F0: 1, F1: 1, F2: false}
            _ = v22
        }
        var v23 int = 934
        _ = v23
    }
    var v24 float64 = 3.14159
    if (((1.5 / 1.5) - 1e10) < (float32((((1.5 + -2.25) + ((0.0 - ((v24 / v24) / (((v24 / float64((0.1 - 3.14159))) / v24) + ((1e10 + 3.14159) - v24)))) * ((v24 + 0.1) - v24))) - v24)) - (((1e10 - (float32((v24 * -2.25)) - 1.5)) - ((3.14159 * 0.0) + -2.25)) + 3.14159))) {
        v1 = (v1 % ((v1) | 1))
        fmt.Println(v24, v13, v1)
    } else {
        var v25 int64 = 0
        _ = v25
    }
    if (v24 < float64((1e10 * ((3.14159 - 3.14159) / 1.5)))) {
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 uint8 = 127
            v28 := []int{32768, fi26}
            v24 = 0.1
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 float64 = v24
        v30 := []int{154, 32767}
        _ = v29
        _ = v30
    } else {
        v31 := GenS3{F0: 757}
        _ = v31
    }
    v13 += (v13 % (((v2 + (v13 + (v13 >> ((18446744073709551615) & 63))))) | 1))
    var v32 bool = ((!(!false)) || false)
    fmt.Println(v32, v2, v32)
    v2 = (v13 | v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v32
}

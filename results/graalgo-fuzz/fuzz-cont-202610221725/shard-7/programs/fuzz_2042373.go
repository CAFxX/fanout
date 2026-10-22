package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint }
func genFn0() uint16 {
    if true {
        var v5 uint = 145
        var v6 uint8 = 255
        _ = v5
        _ = v6
    }
    var v7 int8 = -1
    _ = v7
    return 1
}
func genFn1(p0 int, p1 uint32, p2 uint8) int16 {
    var v8 int64 = -836
    if (0.1 == 3.14159) {
        p2 = 1
    } else {
        var v9 int = 704
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -32768
}

func main() {
    var v1 int64 = -799
    var v2 uint32 = 624
    var v3 uint = 1
    v4 := make(map[string]uint32)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int32 = 1
        _ = fi10
        _ = v11
    }
    var v12 float32 = 0.0
    var v13 int32 = -2147483648
    v14 := GenS1{F0: ("foo bar" + "foo bar"), F1: 265}
    var v15 int16 = 0
    v16 := []int{int(v1), 498, 55}
    v3 = v3
    v17 := GenS2{F0: v3}
    v13 += (^(v13))
    fmt.Println(v3, v13)
    fmt.Println(v3, v2, v1)
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v4[("" + "hello")] = v2
            var v20 int64 = 1
            var v21 uint8 = 147
            if (0.1 != 0.0) {
                v22 := GenS1{F0: ((("" + ((("\n\t" + "\n\t") + "\n\t") + "")) + "a") + "\n\t"), F1: 1}
                _ = v22
            } else {
                v23 := GenS2{F0: v3}
                fmt.Println(v15, v3)
                _ = v23
            }
            if false {
                var v24 uint16 = 42
                _ = v24
            } else {
                v25 := GenS2{F0: 65535}
                switch (fi19 << (((v3 - (v3 + (v3 ^ (^(v3)))))) & 63)) {
                case -2147483648:
                    v15 += v15
                    fmt.Println(fi19, v1, v12)
                    var v26 uint8 = (v21 + 0)
                    _ = v26
                case 4294967295:
                    switch "hello" {
                    case "a":
                        var v27 string = "foo bar"
                        _ = v27
                    case "hello":
                        var v28 int32 = (v13 / (((v13 & -821)) | 1))
                        _ = v28
                    }
                default:
                    if false {
                        switch fi18 {
                        case 118:
                            if true {
                                if len(v16) > 0 {
                                    v16[0] = (fi19 % ((fi19) | 1))
                                }
                                v2 = 166
                            }
                        case -9223372036854775808:
                            v16 = append(v16, fi19)
                            v12 += v12
                        case -1:
                            fmt.Println(v1)
                        default:
                            v12 -= v12
                        }
                    } else {
                        v21 = 127
                        if len(v16) > 0 {
                            v16[0] = (fi18 ^ (fi19 >> ((v3) & 63)))
                        }
                        fmt.Println(v20)
                    }
                }
                v15 += genFn1(fi18, v2, (v21 + 244))
                _ = v25
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        v13 += v13
        fmt.Println(v12, v12)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}

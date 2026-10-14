package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 42
    v3 := make(map[string]uint32)
    var v4 uint32 = 65535
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: (^(fi6))}
        _ = fi6
        _ = v7
    }
    if (1 == 2) {
        fmt.Println(v4, v2, v4)
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 += 247
            v9 := GenS2{F0: 650, F1: "foo bar", F2: true}
            fmt.Println(fi8, fi8, v4)
            _ = fi8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int16 = 0
                if true {
                    var v12 int8 = (v2 / (((v2 ^ (v2 ^ (v2 - 127)))) | 1))
                    _ = v12
                } else {
                    v5 = v5
                    v13 := GenS3{F0: 255}
                    var v14 int8 = int8(v11)
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
            v4 = (v4 << ((42) & 63))
        } else {
            var v15 float64 = 0.1
            _ = v15
        }
    }
    var v16 uint32 = v4
    v5 = (v5 / ((v5) | 1))
    var v17 uint16 = 255
    v16 = (^(v4))
    var v18 uint16 = 464
    var v19 float64 = (-2.25 - (0.0 * 3.14159))
    if (int8(v4) == (v2 - v2)) {
        if true {
            if (0.1 < 0.1) {
                fmt.Println(v19, v17, v4)
                var v20 uint = 1
                var v21 string = ""
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    if (false && true) {
                        fmt.Println(fi23, v17)
                        _ = len(v1)
                    }
                    v18 = (v17 / ((v17) | 1))
                    _ = fi23
                }
                _ = fi22
            }
            if true {
                v18 += 65535
                for fi24 := 0; fi24 < 12; fi24++ {
                    switch (fi24 | (fi24 * 9223372036854775807)) {
                    case 2147483647:
                        switch fi24 {
                        case -2147483648:
                            if (!true) {
                                delete(v3, "")
                            }
                        case -877:
                            fmt.Println(v16)
                        }
                    case 4294967295:
                        var v25 float32 = 0.1
                        _ = v25
                    default:
                        fmt.Println(v17, fi24)
                    }
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 int = -325
                        var v28 int64 = 1
                        var v29 uint64 = 4294967295
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    switch fi24 {
                    case 127:
                        fmt.Println(v4, v18)
                    case 128:
                        if (v4 > (v4 % (((v16 + 2147483648)) | 1))) {
                            v30 := []int{42, (fi24 / ((32767) | 1)), -2147483648}
                            _ = v30
                        }
                    }
                    _ = fi24
                }
            }
        } else {
            if true {
                v31 := GenS2{F0: (v17 ^ 752), F1: (("a" + "") + ""), F2: true}
                _ = v31
            } else {
                switch "hello" {
                case "a":
                    v32 := GenS4{F0: (v4 - v16), F1: 1}
                    _ = v32
                default:
                    var v33 uint32 = 65535
                    _ = v33
                }
                v18 += (v18 % (((^(v17))) | 1))
            }
        }
        if true {
            _ = len(v1)
            v19 -= v19
        } else {
            v34 := v3["foo bar"]
            _ = v34
        }
    }
    v18 -= uint16(v16)
    v18 = (v17 + 523)
    var v35 uint = 1
    v2 -= -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}

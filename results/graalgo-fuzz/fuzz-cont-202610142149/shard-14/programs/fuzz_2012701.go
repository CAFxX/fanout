package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"hello", "hello", "\n\t", (("hello" + "hello") + (("hello" + "\n\t") + "a"))}
    var v3 uint = 2
    v4 := make(map[uint32]uint32)
    switch "hello" {
    case "a":
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 int = -485
            v3 = 287
            _ = len(v2)
            _ = fi5
            _ = v6
        }
    case "":
        v7 := GenS1{F0: (v3 ^ v3), F1: v3}
        _ = v7
    }
    v8 := GenS2{F0: true, F1: 1000, F2: 4294967295}
    var v9 bool = false
    var v10 uint = v3
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint64 = 18446744073709551615
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS2{F0: (1.5 != 1.5), F1: -32768, F2: 246}
            v9 = v9
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if v9 {
        if v9 {
            for fi16 := 0; fi16 < 4; fi16++ {
                v3 = 65535
                fmt.Println(v9, v9)
                _ = fi16
            }
        }
    }
    if ((((1e10 / -2.25) <= -2.25) || (!false)) || (!v9)) {
        var v17 uint64 = 4294967295
        v18 := GenS3{F0: 427, F1: 42, F2: v17}
        if v9 {
            if ("foo bar" > "\n\t") {
                if (41 <= 128) {
                    switch 13 {
                    case 1:
                        v10 = (v10 ^ (v3 / ((v10) | 1)))
                    }
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if (364 > 959) {
                            v3 = 65535
                            v17 = v17
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        }
                        _ = fi19
                    }
                }
                v20 := GenS4{F0: 26}
                v21 := GenS3{F0: 38, F1: 42, F2: (^(v17))}
                _ = v20
                _ = v21
            }
        }
        _ = v17
        _ = v18
    }
    if (1 >= 162) {
        var v22 int64 = 42
        _ = v22
    } else {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = 9223372036854775807
            v2 = append(v2, "a")
            _ = fi23
            _ = v24
        }
        if (!v9) {
            var v25 uint16 = 733
            switch ("a" + "a") {
            case "a":
                v10 -= 4294967295
            default:
                v9 = (!(1e10 > float64(1e10)))
            }
            switch "a" {
            case "":
                v26 := GenS5{F0: 42}
                _ = v26
            default:
                v25 -= (v25 % ((v25) | 1))
            }
            _ = v25
        } else {
            if v9 {
                var v27 float32 = -2.25
                switch -338 {
                case 9223372036854775807:
                    v28 := v1[2147483647]
                    _ = v28
                case 127:
                    if (128 < 127) {
                        var v29 bool = true
                        _ = v29
                    }
                default:
                    fmt.Println(v3)
                }
                _ = v27
            }
            for fi30 := 0; fi30 < 20; fi30++ {
                v10 += (^(v10))
                fmt.Println(v9)
                _ = fi30
            }
        }
    }
    var v31 int8 = -1
    v3 -= (v10 | (v10 << ((v3) & 63)))
    v31 = (v31 / (((v31 % ((v31) | 1))) | 1))
    v9 = ((!false) || v9)
    var v32 int8 = -1
    if v9 {
        v3 -= 389
    } else {
        v2 = append(v2, "a")
        var v33 float32 = 1.5
        _ = v33
    }
    v31 = (v31 | 42)
    var v34 bool = ((!true) || false)
    if (!(1 > 274)) {
        var v35 int8 = 1
        _ = v35
    } else {
        delete(v4, uint32(v10))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v8
    _ = v9
}

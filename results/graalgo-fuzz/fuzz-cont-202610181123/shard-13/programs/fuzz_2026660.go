package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 string) int64 {
    p0 = p0
    var v4 int64 = 324
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 255}
        if false {
            var v7 uint = 42
            v8 := GenS2{F0: (p0 + (p0 + p0)), F1: true}
            v9 := GenS1{F0: v7}
            _ = v7
            _ = v8
            _ = v9
        } else {
            if true {
                p0 = "\n\t"
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := []int{32767, fi10, (fi10 | (fi5 * (fi5 << ((993) & 63)))), fi5}
                    _ = fi10
                    _ = v11
                }
                var v12 uint16 = 2
                _ = v12
            } else {
                fmt.Println(fi5, p0)
                if false {
                    v4 = (v4 - 175)
                    var v13 int32 = 281
                    var v14 int16 = 42
                    _ = v13
                    _ = v14
                } else {
                    var v15 string = "hello"
                    _ = v15
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    if (-51 >= 1) {
                        p0 = p0
                    } else {
                        v17 := GenS3{F0: 0, F1: 128, F2: (!(!(255 <= 42)))}
                        var v18 int32 = 994
                        _ = v17
                        _ = v18
                    }
                    if false {
                        var v19 int64 = (v4 ^ v4)
                        var v20 uint = 1
                        _ = v19
                        _ = v20
                    }
                    if (781 == 959) {
                        switch (fi16 + fi16) {
                        case 848:
                            if false {
                                var v21 float32 = (3.14159 - 0.1)
                                v22 := GenS4{F0: 18446744073709551615}
                                v23 := []int{fi5, 4294967295, (^(fi16))}
                                var v24 uint = 101
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                v25 := GenS1{F0: 671}
                                v26 := []int{(-(fi16))}
                                v4 += v4
                                _ = v25
                                _ = v26
                            }
                        default:
                            v27 := []int{fi16, fi16}
                            _ = v27
                        }
                    }
                    _ = fi16
                }
            }
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return (v4 - (v4 - v4))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]int)
    var v3 bool = ("" < "foo bar")
    var v28 uint16 = 42
    switch "a" {
    case "\n\t":
        var v29 int8 = 62
        _ = v29
    case "":
        v28 = (^(v28))
    case "foo bar":
        v30 := GenS3{F0: (v28 - v28), F1: 0, F2: (!false)}
        _ = v30
    }
    v1 = (v1 >> ((18446744073709551615) & 63))
    v31 := []int{-2147483648}
    var v32 int = 9223372036854775807
    var v33 int32 = 0
    v1 = v1
    var v34 int = v32
    v2["hello"] = (^(v32))
    if len(v31) > 0 {
        v31[0] = (v34 | (v32 >> ((2) & 63)))
    }
    var v35 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}

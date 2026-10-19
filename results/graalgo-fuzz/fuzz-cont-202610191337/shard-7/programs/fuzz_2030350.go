package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    var v1 uint8 = 156
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    if true {
        v1 = v1
        fmt.Println(v1)
    } else {
        v4 := []int{9223372036854775807, -477, 703}
        for fi5 := 0; fi5 < 14; fi5++ {
            switch "foo bar" {
            case "hello":
                v1 -= 255
            case "foo bar":
                v1 += (v1 - (v1 % ((18) | 1)))
            case "a":
                var v6 int = 127
                _ = v6
            default:
                var v7 int = fi5
                _ = v7
            }
            var v8 int = (-(fi5))
            for fi9 := 0; fi9 < 5; fi9++ {
                _ = len(v2)
                switch ("" + (("hello" + "\n\t") + "\n\t")) {
                case "a":
                    fmt.Println(fi5, fi9, v1)
                    fmt.Println(v8, v1, fi5)
                    switch "\n\t" {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        fmt.Println(fi9, fi5, v8)
                    case "":
                        if false {
                            fmt.Println(v8, fi5)
                            var v10 uint = 161
                            if (255 > (v1 | 0)) {
                                var v11 int16 = 973
                                v1 = (v1 & 2)
                                v10 -= v10
                                v12 := GenS1{F0: 1, F1: -9223372036854775808}
                                _ = v11
                                _ = v12
                            } else {
                                if false {
                                    fmt.Println(fi5)
                                } else {
                                    var v13 int8 = -92
                                    if (677 > 0) {
                                        v13 = 42
                                    }
                                    _ = v13
                                }
                            }
                            v1 += v1
                            _ = v10
                        }
                    default:
                        if true {
                            if false {
                                v1 -= 1
                                if (438 >= 0) {
                                    var v14 int = fi5
                                    _ = v14
                                }
                            } else {
                                var v15 int64 = 853
                                var v16 uint64 = 42
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v15
                                _ = v16
                            }
                            v1 -= 214
                        } else {
                            if false {
                                v1 += v1
                                if (316 > 164) {
                                    v17 := GenS2{F0: 42}
                                    v18 := GenS1{F0: 42, F1: -1}
                                    _ = v17
                                    _ = v18
                                } else {
                                    v1 = (v1 + 42)
                                    fmt.Println(v1, v1)
                                }
                                v19 := v3[2]
                                _ = v19
                            }
                        }
                        var v20 uint32 = 42
                        _ = v20
                    }
                    v8 += (v8 & 4294967295)
                case "hello":
                    var v21 int = -2147483648
                    _ = v21
                }
                _ = fi9
            }
            _ = fi5
            _ = v8
        }
        var v22 float32 = 0.1
        _ = v22
        _ = v4
    }
    v1 -= v1
    v1 += (v1 >> ((0) & 63))
    var v23 bool = true
    v24 := []int{int(v1), -9223372036854775808}
    _ = len(v24)
    var v25 string = ""
    fmt.Println(v1)
    fmt.Println(v23, v1, v1)
    var v26 int32 = 1
    v27 := v2[65535]
    var v28 int64 = -832
    _ = len(v24)
    var v29 uint64 = 403
    v30 := GenS3{F0: 255, F1: (v25 + "\n\t"), F2: v28}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []int{4294967295, -927, 2147483647, 809}
    var v2 uint64 = 790
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    v2 = (v2 & (v2 - v2))
    v2 -= v2
    for fi4 := 0; fi4 < 10; fi4++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(fi4)
        _ = fi4
    }
    if (2 > 42) {
        v5 := GenS1{F0: "", F1: -1, F2: uint16(v2)}
        if false {
            if (!((352 == 2) && ((v2 < (v2 & (v2 | v2))) || ("a" == ("a" + ""))))) {
                var v6 int16 = 0
                v7 := GenS2{F0: v2, F1: 42, F2: (false && (542 < 1))}
                v3["hello"] = 2147483647
                _ = v6
                _ = v7
            }
        }
        v2 += (^(v2))
        _ = v5
    }
    var v8 int32 = -1
    var v9 uint32 = 481
    v3["foo bar"] = (v9 << ((65535) & 63))
    var v10 uint32 = (v9 - 4294967295)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    var v11 bool = true
    v2 -= (v2 & 962)
    switch "\n\t" {
    case "":
        fmt.Println(v10, v9, v2)
    case "a":
        for fi12 := 0; fi12 < 11; fi12++ {
            switch "a" {
            case "\n\t":
                var v13 int8 = 1
                _ = v13
            }
            switch fi12 {
            case 1:
                fmt.Println(v9)
            default:
                if (42 != uint8(v9)) {
                    v14 := GenS1{F0: "a", F1: 9223372036854775807, F2: uint16(v10)}
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if false {
                            v16 := GenS1{F0: "hello", F1: -9223372036854775808, F2: 2}
                            if len(v1) > 0 {
                                v1[0] = fi12
                            }
                            fmt.Println(v10, fi15)
                            _ = v16
                        } else {
                            if v11 {
                                if ("" == ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + "foo bar")))) {
                                    switch int(v2) {
                                    case 128:
                                        fmt.Println(fi12, v8, v2)
                                    case -9223372036854775808:
                                        v11 = v11
                                    case 382:
                                        var v17 int64 = 9223372036854775807
                                        _ = v17
                                    default:
                                        switch (fi12 ^ int(v8)) {
                                        case 219:
                                            var v18 uint8 = 162
                                            _ = v18
                                        case -398:
                                            if (v11 || v11) {
                                                var v19 string = ("a" + "\n\t")
                                                fmt.Println(v11)
                                                var v20 int32 = v8
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v2 = v2
                                            }
                                        case 0:
                                            v10 += v10
                                        default:
                                            var v21 int = 127
                                            _ = v21
                                        }
                                    }
                                }
                            } else {
                                var v22 uint8 = 128
                                v23 := GenS3{F0: 734, F1: -128}
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = fi15
                    }
                    _ = v14
                }
            }
            _ = fi12
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 string; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    v3 := []int{2147483647}
    for fi4 := 0; fi4 < 15; fi4++ {
        v2 = "foo bar"
        for fi5 := 0; fi5 < 15; fi5++ {
            if true {
                var v6 int32 = 1
                _ = v6
            } else {
                switch fi4 {
                case 660:
                    switch v2 {
                    case "a":
                        v7 := GenS1{F0: 39, F1: 432, F2: -2147483648}
                        _ = v7
                    }
                case -455:
                    v3 = append(v3, fi4)
                }
                v2 = ("a" + v2)
                v8 := GenS1{F0: 1, F1: 245, F2: 2147483647}
                v9 := GenS1{F0: 42, F1: 122, F2: 42}
                _ = v8
                _ = v9
            }
            _ = fi5
        }
        v2 = (v2 + v2)
        _ = fi4
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if ((true || (1 == 2)) && (!true)) {
                switch 127 {
                case 182:
                    v2 = v2
                case -9223372036854775808:
                    if (!(-1 >= -139)) {
                        fmt.Println(fi10, v2)
                        if (42 != -1) {
                            switch (fi11 & fi11) {
                            case 98:
                                _ = len(v1)
                            }
                        }
                        switch v2 {
                        case "hello":
                            switch fi11 {
                            case -2147483648:
                                v12 := GenS1{F0: 110, F1: 434, F2: -635}
                                _ = v12
                            case 32768:
                                var v13 uint8 = 128
                                _ = v13
                            }
                        case "foo bar":
                            v14 := GenS2{F0: 2}
                            _ = v14
                        default:
                            v15 := GenS3{F0: fi10, F1: 65535}
                            _ = v15
                        }
                    } else {
                        v16 := GenS4{F0: -128}
                        v17 := GenS3{F0: -9223372036854775808, F1: 2147483648}
                        _ = v16
                        _ = v17
                    }
                case -572:
                    var v18 int8 = 84
                    _ = v18
                default:
                    fmt.Println(fi11, v2)
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    var v19 uint32 = 2147483648
    if ((((!true) && false) && (42 != 32768)) && (65535 > 2)) {
        var v20 uint8 = 143
        _ = v20
    }
    v21 := v1[428]
    v22 := GenS5{F0: 2, F1: v2, F2: (!false)}
    v23 := GenS4{F0: 127}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int16 = -1
        v19 += (v19 | (v19 >> ((49) & 63)))
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}

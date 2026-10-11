package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int; F1 uint32; F2 int }
type GenS4 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -398
    v3 := make(map[uint32]int)
    var v4 uint64 = 42
    v5 := []uint32{65535, 65535}
    v6 := []string{}
    var v7 uint32 = 2147483647
    if true {
        switch 42 {
        case -382:
            v2 -= (v2 ^ (v2 & (v2 % ((1) | 1))))
        }
        var v8 int64 = -9223372036854775808
        fmt.Println(v4, v8)
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: v4}
    v7 += (^(v7))
    v7 = v7
    if (230 <= 1) {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint = 65535
            _ = fi10
            _ = v11
        }
        switch 519 {
        case 916:
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 string = ""
                switch "foo bar" {
                case "a":
                    var v14 int32 = -371
                    _ = v14
                case "":
                    if (0 >= 128) {
                        v15 := GenS1{F0: (fi12 - fi12), F1: v4}
                        for fi16 := 0; fi16 < 20; fi16++ {
                            if false {
                                if true {
                                    var v17 uint16 = 1
                                    _ = v17
                                } else {
                                    var v18 uint32 = v7
                                    v13 = ((((("foo bar" + v13) + (v13 + "\n\t")) + (v13 + ("foo bar" + (((v13 + ("\n\t" + (v13 + (("\n\t" + v13) + ("" + (v13 + "\n\t")))))) + (v13 + "")) + v13)))) + ("\n\t" + (v13 + "a"))) + "hello")
                                    fmt.Println(v2, v13)
                                    v18 -= v18
                                    _ = v18
                                }
                            } else {
                                var v19 int = fi12
                                v7 += (v7 ^ (v7 >> ((554) & 63)))
                                _ = v19
                            }
                            fmt.Println(fi16, v7, v2)
                            _ = fi16
                        }
                        _ = v15
                    }
                default:
                    if true {
                        v3[v7] = (^(fi12))
                        v20 := GenS1{F0: fi12, F1: 42}
                        var v21 bool = true
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v23 := GenS2{F0: (78 >= 15), F1: 778}
                            _ = fi22
                            _ = v23
                        }
                        v4 += (^(v4))
                    }
                }
                v2 += (v2 - 0)
                _ = fi12
                _ = v13
            }
        case 32767:
            v24 := GenS1{F0: 9223372036854775807, F1: 2}
            _ = v24
        case -1:
            if (!true) {
                fmt.Println(v4, v7)
                v3[42] = 0
                for fi25 := 0; fi25 < 11; fi25++ {
                    v7 = v7
                    v26 := GenS3{F0: (fi25 * fi25), F1: (v7 - (v7 / ((v7) | 1))), F2: fi25}
                    _ = fi25
                    _ = v26
                }
            }
        default:
            var v27 uint = 255
            _ = v27
        }
        v6 = append(v6, "hello")
    } else {
        v4 = (v4 / (((v4 << ((2) & 63))) | 1))
    }
    if true {
        _ = len(v6)
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 bool = false
            v29 = (!v29)
            _ = fi28
            _ = v29
        }
        if false {
            fmt.Println(v7, v7)
        } else {
            switch 2147483647 {
            case 2147483647:
                var v30 int64 = v2
                _ = v30
            case -817:
                if true {
                    var v31 int64 = (v2 - (v2 + 0))
                    v2 += -1
                    _ = v31
                }
            case 32768:
                for fi32 := 0; fi32 < 8; fi32++ {
                    v4 = (v4 ^ v4)
                    var v33 int64 = v2
                    _ = fi32
                    _ = v33
                }
            default:
                if false {
                    var v34 uint64 = (v4 >> ((42) & 63))
                    v35 := GenS4{F0: 127, F1: 255}
                    _ = v34
                    _ = v35
                }
            }
        }
    }
    var v36 uint32 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

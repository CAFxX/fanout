package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 uint32 = 894
    v4 := []string{"foo bar"}
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: true, F1: ("hello" + "\n\t"), F2: 996}
        v3 = v3
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 1, F1: -1}
    v3 = v3
    switch 9223372036854775807 {
    case -942:
        fmt.Println(v3)
    default:
        v3 = (v3 % (((v3 % (((v3 - (v3 + v3))) | 1))) | 1))
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v2)
        switch (fi8 - -1) {
        case 4294967295:
            if (!true) {
                fmt.Println(fi8, v3)
                switch ("hello" + ((("" + "") + "") + "a")) {
                case "\n\t":
                    if true {
                        for fi9 := 0; fi9 < 15; fi9++ {
                            fmt.Println(fi8)
                            v10 := GenS3{F0: fi8}
                            _ = fi9
                            _ = v10
                        }
                        switch (fi8 * fi8) {
                        case 127:
                            delete(v2, (v3 * (v3 & (v3 / ((v3) | 1)))))
                        default:
                            if true {
                                var v11 bool = false
                                _ = v11
                            }
                        }
                    }
                default:
                    if false {
                        var v12 int64 = 265
                        _ = v12
                    }
                }
                v3 = v3
            }
        case 376:
            delete(v2, 2)
        case -247:
            _ = len(v1)
        default:
            v3 = 42
        }
        _ = fi8
    }
    var v13 int64 = -327
    v14 := GenS4{F0: 4294967295, F1: 255, F2: 0}
    var v15 int = -624
    for fi16 := 0; fi16 < 8; fi16++ {
        if (false || true) {
            v17 := GenS1{F0: (1.5 <= float32((1e10 - 0.0))), F1: "", F2: v3}
            v3 = v3
            _ = v17
        } else {
            var v18 float32 = 0.1
            _ = v18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            fmt.Println(fi19, v15)
            fmt.Println(v3, v3)
            _ = fi19
        }
        _ = fi16
    }
    var v20 int32 = 2147483647
    v13 = (v13 * v13)
    switch "a" {
    case "foo bar":
        var v21 int32 = v20
        _ = v21
    case "":
        v2[1] = (-(v15))
    case "a":
        var v22 int16 = -60
        _ = v22
    default:
        v23 := GenS5{F0: 751}
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}

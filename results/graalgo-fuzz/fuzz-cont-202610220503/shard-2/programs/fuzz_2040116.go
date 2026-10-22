package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 string = "hello"
    var v3 uint16 = 42
    var v4 float64 = 0.0
    if true {
        v5 := GenS1{F0: 1, F1: (v2 >= v2)}
        _ = v5
    } else {
        delete(v1, v2)
        v6 := []int{-927, -2147483648, 32768, -227}
        _ = v6
    }
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 = v7
        fmt.Println(v2, v2, v2)
        _ = fi8
    }
    var v9 int = -822
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 = -960
        _ = fi10
    }
    var v11 float64 = (1.5 + (v4 / ((v4 - 0.0) + v4)))
    var v12 uint = 42
    v9 = (v9 ^ 128)
    if ((false && ((^(v7)) == (v7 / ((uint32(v3)) | 1)))) || (100 > 1)) {
        if true {
            if (true && (!false)) {
                fmt.Println(v12, v11)
                fmt.Println(v2, v9, v7)
                fmt.Println(v9, v4, v12)
            }
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint16 = v3
                v15 := GenS1{F0: 571, F1: (!(false || ((!true) || (3.14159 != 0.1))))}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int32 = 268
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v9 = (^(v9))
        if (!(!((4294967295 >= 1) || ((v12 << ((v12) & 63)) < v12)))) {
            switch v2 {
            case "\n\t":
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := GenS1{F0: v12, F1: (128 != 128)}
                    var v21 uint16 = (v3 - 63)
                    var v22 uint8 = 27
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            case "a":
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch v2 {
                    case "a":
                        var v24 int = (fi18 % ((v9) | 1))
                        _ = v24
                    case "foo bar":
                        v25 := GenS2{F0: 0, F1: (fi18 / ((-2147483648) | 1))}
                        _ = v25
                    }
                    _ = fi23
                }
            case "hello":
                if false {
                    if false {
                        delete(v1, v2)
                        var v26 string = v2
                        _ = v26
                    } else {
                        var v27 uint = (^(v12))
                        v27 = 4294967295
                        _ = v27
                    }
                    switch (-(v9)) {
                    case 1:
                        if false {
                            fmt.Println(v4)
                            for fi28 := 0; fi28 < 7; fi28++ {
                                switch (v9 / (((fi28 - v9)) | 1)) {
                                case 488:
                                    v9 += fi18
                                case 128:
                                    var v29 int16 = 42
                                    _ = v29
                                default:
                                    var v30 uint16 = (v3 - 42)
                                    _ = v30
                                }
                                switch (v2 + v2) {
                                case "foo bar":
                                    var v31 uint8 = 127
                                    _ = v31
                                case "hello":
                                    v32 := GenS3{F0: 437}
                                    _ = v32
                                default:
                                    _ = len(v1)
                                }
                                _ = fi28
                            }
                        }
                    case -658:
                        v33 := []int{v9, v9, (v9 + fi18), (fi18 - 32767)}
                        _ = v33
                    case -732:
                        v11 = ((1.5 / v11) - v11)
                    default:
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 uint32 = (v7 * v7)
                            _ = fi34
                            _ = v35
                        }
                    }
                    v36 := []int{127, v9}
                    _ = v36
                }
            }
            if ((!true) || true) {
                switch "hello" {
                case "":
                    v37 := []int{v9, fi18, v9}
                    _ = v37
                case "foo bar":
                    var v38 int = 263
                    _ = v38
                case "a":
                    v39 := GenS2{F0: -722, F1: (fi18 - fi18)}
                    _ = v39
                default:
                    fmt.Println(v2, v12, v7)
                }
                fmt.Println(v7)
                switch (v2 + "\n\t") {
                case "hello":
                    fmt.Println(v4)
                default:
                    var v40 int32 = 2147483647
                    _ = v40
                }
            }
        }
        var v41 int16 = 42
        _ = fi18
        _ = v41
    }
    v42 := GenS2{F0: -1, F1: (v9 | -860)}
    v1[v2] = (v9 * 42)
    delete(v1, (v2 + (((v2 + "foo bar") + v2) + v2)))
    delete(v1, v2)
    v43 := GenS2{F0: 32767, F1: (v9 - (v9 << ((1) & 63)))}
    for fi44 := 0; fi44 < 19; fi44++ {
        if true {
            v45 := []int{}
            var v46 float32 = 0.1
            _ = v45
            _ = v46
        }
        _ = fi44
    }
    v47 := []int{v9, (v9 / (((v9 & (v9 | (v9 ^ (v9 << ((689) & 63)))))) | 1)), (v9 + v9), 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v7
    _ = v9
}

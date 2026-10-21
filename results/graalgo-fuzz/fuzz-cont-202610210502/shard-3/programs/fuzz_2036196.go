package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 int16 }
func genFn0() uint {
    switch "foo bar" {
    case "a":
        var v5 uint16 = 0
        _ = v5
    case "\n\t":
        v6 := GenS1{F0: -32768, F1: false}
        _ = v6
    case "foo bar":
        if false {
            switch "" {
            case "":
                v7 := GenS2{F0: "hello", F1: 32767}
                _ = v7
            default:
                for fi8 := 0; fi8 < 8; fi8++ {
                    fmt.Println(fi8)
                    _ = fi8
                }
            }
            var v9 int = 1
            _ = v9
        } else {
            if ((!true) && (!true)) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint = 0
                    fmt.Println(fi10, fi10)
                    if (0.0 <= (float64(0.1) + 1e10)) {
                        var v12 uint64 = 2
                        v11 = v11
                        _ = v12
                    } else {
                        var v13 uint8 = 61
                        if (!(((2 <= 255) && false) || (-369 <= -53))) {
                            var v14 uint16 = 292
                            switch fi10 {
                            case 128:
                                fmt.Println(fi10)
                            case 32768:
                                v11 -= v11
                            case -982:
                                if true {
                                    if (((3.14159 / 0.1) * 1.5) >= 0.0) {
                                        fmt.Println(v13, v14, fi10)
                                    }
                                }
                            }
                            _ = v14
                        } else {
                            v15 := GenS2{F0: (("\n\t" + "\n\t") + "hello"), F1: 32767}
                            var v16 uint16 = 255
                            v17 := GenS2{F0: "", F1: 32767}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
    default:
        v18 := GenS1{F0: 190, F1: (32767 != 65535)}
        _ = v18
    }
    var v19 uint8 = 13
    _ = v19
    return 42
}

func main() {
    var v1 uint32 = 131
    var v2 uint8 = 33
    v3 := []int{-9223372036854775808, -9223372036854775808, -1, 1}
    var v4 uint64 = 254
    var v20 uint = 2
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 float64 = 1e10
            _ = len(v3)
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if true {
        v20 += 958
    }
    if (!true) {
        v20 += (v20 >> ((2) & 63))
    }
    v2 = (v2 ^ 0)
    v4 = uint64(v20)
    var v24 bool = ((!((!true) || false)) && false)
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 int = (fi25 % (((fi25 + fi25)) | 1))
        var v27 uint16 = 1
        for fi28 := 0; fi28 < 18; fi28++ {
            if ((!v24) && v24) {
                v27 -= v27
            } else {
                v1 = v1
                var v29 float32 = 3.14159
                _ = v29
            }
            _ = fi28
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    if (!v24) {
        v20 = v20
        var v30 int8 = 0
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []uint32{65535, 0, 42}
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint16 = 42
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: (2 >= 952), F1: (127 <= 42)}
    v7 := GenS2{F0: 1, F1: 22, F2: 0}
    v2 = (v2 | v2)
    v2 -= 258
    fmt.Println(v2)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = -128
        switch fi8 {
        case 127:
            switch ("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                var v10 string = ""
                _ = v10
            case "foo bar":
                if (!true) {
                    var v11 string = "foo bar"
                    _ = v11
                } else {
                    switch int(v2) {
                    case -587:
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v2 += v2
                            _ = fi12
                        }
                    default:
                        if (-2.25 == 0.1) {
                            var v13 uint8 = 255
                            _ = len(v3)
                            _ = v13
                        }
                    }
                    var v14 uint16 = 103
                    _ = v14
                }
            default:
                v15 := GenS3{F0: v2, F1: 0}
                _ = v15
            }
        default:
            fmt.Println(v9)
        }
        _ = fi8
        _ = v9
    }
    v2 += v2
    v2 += (v2 << ((1) & 63))
    var v16 string = "a"
    if (26 > 204) {
        if true {
            if ((237 >= 232) || false) {
                var v17 uint32 = 65535
                _ = v17
            } else {
                v16 = "hello"
                v16 = v16
            }
        }
        v16 = v16
    } else {
        if ((false && true) || (((v2 | v2) > v2) && false)) {
            var v18 float32 = -2.25
            var v19 uint8 = 139
            var v20 float32 = ((float32(0.1) - v18) * -2.25)
            _ = v18
            _ = v19
            _ = v20
        }
    }
    v16 = v16
    if false {
        var v21 int16 = (v2 ^ (v2 / ((v2) | 1)))
        v22 := GenS4{F0: 418}
        if true {
            fmt.Println(v2)
            for fi23 := 0; fi23 < 13; fi23++ {
                if false {
                    v16 = "\n\t"
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v2 = v21
                        var v25 int16 = v21
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 352
                            }
                            v21 += (-(v2))
                        }
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi23
            }
        } else {
            v21 = (v21 + v2)
            v3[4294967295] = ("hello" + (v16 + ("" + "")))
            v26 := GenS5{F0: 4294967295}
            _ = v26
        }
        _ = v21
        _ = v22
    } else {
        v1 = append(v1, 27)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}

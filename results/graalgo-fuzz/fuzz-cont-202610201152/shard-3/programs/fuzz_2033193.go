package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = (p0 - (-(p0)))
        v8 := GenS1{F0: p0, F1: 127}
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 += -1
    switch 9223372036854775807 {
    case 54:
        var v9 string = ""
        _ = v9
    case 32767:
        p0 = 42
    default:
        p0 = 1
    }
    _ = p0
    return 65535
}

func main() {
    var v1 bool = true
    v2 := []int{0, 0}
    var v3 uint64 = 647
    v4 := make(map[int]uint32)
    var v5 uint64 = v3
    v3 = v3
    if false {
        v10 := v4[286]
        _ = v10
    }
    if true {
        v11 := GenS1{F0: 109, F1: 52}
        switch 32767 {
        case 2147483647:
            if v1 {
                v12 := GenS1{F0: -83, F1: -1}
                fmt.Println(v3, v1, v1)
                fmt.Println(v5, v1, v5)
                _ = v12
            } else {
                var v13 int = 365
                switch "a" {
                case "":
                    v13 -= v13
                default:
                    var v14 int = (v13 - 663)
                    _ = v14
                }
                _ = v13
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v15 := GenS1{F0: 127, F1: 100}
        _ = v11
        _ = v15
    } else {
        fmt.Println(v3, v5, v1)
    }
    _ = len(v4)
    if ((v1 && v1) || v1) {
        fmt.Println(v3, v3, v1)
        fmt.Println(v1, v5, v1)
    }
    var v16 uint32 = 0
    for fi17 := 0; fi17 < 19; fi17++ {
        v16 = 742
        if true {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: 127, F1: int8(v16)}
                var v20 int = fi18
                var v21 int8 = 42
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            if v1 {
                switch (fi17 & 768) {
                case 1:
                    fmt.Println(fi17, fi17)
                default:
                    var v22 int64 = -1
                    _ = v22
                }
                var v23 int32 = 0
                _ = v23
            }
        } else {
            v5 -= (v3 >> ((418) & 63))
        }
        _ = fi17
    }
    v24 := GenS2{F0: 4294967295, F1: 2147483647}
    if true {
        v4[2147483647] = (v16 + (v16 % (((v16 + (v16 / ((v16) | 1)))) | 1)))
    } else {
        if false {
            var v25 bool = false
            for fi26 := 0; fi26 < 17; fi26++ {
                if true {
                    if true {
                        fmt.Println(v25)
                        var v27 int16 = 134
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi26, fi26, v27)
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                            case "foo bar":
                                fmt.Println(v27)
                            default:
                                var v28 int32 = 2147483647
                                _ = v28
                            }
                            switch -485 {
                            case -1:
                                v29 := GenS3{F0: 102, F1: 128}
                                _ = v29
                            default:
                                var v30 uint32 = 1
                                _ = v30
                            }
                        }
                        _ = v27
                    }
                }
                _ = fi26
            }
            _ = v25
        }
        v31 := GenS3{F0: 76, F1: 255}
        if ((!((0 != 2147483647) && v1)) || false) {
            v5 = (v5 >> ((108) & 63))
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

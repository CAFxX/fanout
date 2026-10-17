package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int16 }

func main() {
    var v1 int = 32767
    v2 := make(map[int]uint32)
    var v3 int32 = 42
    v3 = (v3 << ((18446744073709551615) & 63))
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            switch fi5 {
            case 127:
                var v6 int16 = -32768
                _ = v6
            case 742:
                if false {
                    if (!true) {
                        v1 -= (-(v1))
                        switch v1 {
                        case 127:
                            var v7 uint = 255
                            _ = v7
                        case -959:
                            var v8 int16 = 354
                            _ = v8
                        }
                    }
                    if (!false) {
                        v1 += fi4
                        if ((676 >= 603) || false) {
                            var v9 int16 = 666
                            _ = v9
                        }
                    } else {
                        v10 := GenS1{F0: 0, F1: 274}
                        var v11 int16 = 0
                        _ = v10
                        _ = v11
                    }
                }
            case -9223372036854775808:
                if (("hello" + ("a" + "\n\t")) > "") {
                    var v12 int64 = 42
                    if (!false) {
                        v13 := GenS2{F0: 67}
                        if (255 >= 0) {
                            v12 = (v12 % ((v12) | 1))
                            switch 32767 {
                            case 188:
                                v1 += fi5
                            default:
                                v14 := GenS1{F0: 1000, F1: 710}
                                _ = v14
                            }
                        }
                        _ = v13
                    }
                    _ = v12
                }
            default:
                v15 := GenS3{F0: (v3 >> ((0) & 63))}
                _ = v15
            }
            if true {
                v16 := []int{(-(fi4))}
                _ = v16
            } else {
                v3 = (v3 % (((^(v3))) | 1))
            }
            _ = fi5
        }
        var v17 int32 = v3
        _ = fi4
        _ = v17
    }
    v18 := []int{(v1 | (v1 & (v1 << ((4294967295) & 63))))}
    _ = len(v2)
    if (!false) {
        for fi19 := 0; fi19 < 6; fi19++ {
            _ = len(v18)
            fmt.Println(v3, v3)
            v3 -= (v3 | v3)
            _ = fi19
        }
        v20 := GenS3{F0: v3}
        _ = v20
    }
    v1 = 534
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS4{F0: "", F1: 1}
        _ = fi21
        _ = v22
    }
    v18 = append(v18, 128)
    switch (-(v1)) {
    case 32767:
        var v23 uint = 0
        _ = v23
    default:
        v24 := GenS2{F0: 697}
        _ = v24
    }
    v25 := v2[(v1 | v1)]
    v26 := GenS3{F0: (v3 + (v3 / ((-2147483648) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}

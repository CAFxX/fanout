package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 int8, p2 uint8) uint8 {
    if false {
        p0 = (^(p0))
    }
    if (!false) {
        var v4 int64 = 0
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(p2, p1, fi5)
            fmt.Println(p1)
            _ = fi5
        }
        _ = v4
    } else {
        p0 -= p0
        v6 := GenS1{F0: 65535, F1: 1}
        _ = v6
    }
    if true {
        fmt.Println(p1)
    } else {
        p0 = p0
        switch -997 {
        case 259:
            if true {
                switch 9223372036854775807 {
                case 127:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(p2, p2)
                    default:
                        v7 := GenS2{F0: (p2 * (p2 % ((p2) | 1))), F1: 255}
                        _ = v7
                    }
                case 4294967295:
                    if (990 > -1) {
                        switch "\n\t" {
                        case "\n\t":
                            v8 := GenS3{F0: uint(p2), F1: 1, F2: p1}
                            _ = v8
                        default:
                            var v9 uint16 = uint16(p2)
                            _ = v9
                        }
                        v10 := []int{526}
                        var v11 int64 = int64(p2)
                        _ = v10
                        _ = v11
                    }
                case 553:
                    v12 := GenS1{F0: 310, F1: 814}
                    _ = v12
                }
            }
        case -456:
            v13 := GenS4{F0: 4294967295}
            _ = v13
        case 127:
            var v14 uint32 = 4294967295
            _ = v14
        default:
            if false {
                var v15 int8 = 100
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint = 0
        _ = fi16
        _ = v17
    }
    p2 += (p2 >> ((545) & 63))
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 160
    v3 := []string{"foo bar", "hello", "hello"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, "\n\t")
    v2 = 42
    if true {
        v2 += v2
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v2 = (^(v2))
        v1 -= 501
        if true {
            v3 = append(v3, "foo bar")
            var v19 int32 = -2147483648
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                v1 += (^(v1))
                v2 = v2
                _ = fi20
            }
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        _ = len(v3)
        v2 = v2
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS3{F0: (v2 / ((65535) | 1)), F1: v1, F2: 1}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v1 = (v1 << ((v2) & 63))
    v24 := GenS4{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}

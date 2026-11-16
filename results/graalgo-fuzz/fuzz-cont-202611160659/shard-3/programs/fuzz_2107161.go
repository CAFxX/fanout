package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := []uint32{268, 2147483647, 2147483647}
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 85
    switch "\n\t" {
    case "\n\t":
        for fi5 := 0; fi5 < 2; fi5++ {
            v4 = v4
            if ((false || (2147483647 >= -2147483648)) && (!(!(true && true)))) {
                v6 := GenS1{F0: 24, F1: 2147483647, F2: 45}
                _ = v6
            }
            _ = fi5
        }
    case "a":
        v2[105] = "\n\t"
    }
    var v7 string = ("a" + "\n\t")
    var v8 int8 = -75
    v7 = v7
    var v9 uint32 = 65535
    var v10 uint8 = uint8(v4)
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v9)
            v13 := GenS1{F0: (v10 - v10), F1: (v9 << ((195) & 63)), F2: v8}
            v14 := GenS1{F0: v10, F1: 65535, F2: (-(v8))}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case 674:
        if true {
            var v15 uint16 = uint16(v10)
            var v16 int64 = 1
            if true {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v9 = v9
                    if (!((v10 != v10) || false)) {
                        v18 := GenS1{F0: 1, F1: (v9 | (v9 % (((v9 * v9)) | 1))), F2: -17}
                        _ = v18
                    } else {
                        v10 = 0
                        var v19 int16 = -32768
                        delete(v2, fi17)
                        _ = v19
                    }
                    _ = fi17
                }
            }
            _ = v15
            _ = v16
        }
    case 168:
        v4 += v4
    case 0:
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int32 = -1
            var v22 int = -9223372036854775808
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        v23 := GenS2{F0: 1, F1: (v8 % ((v8) | 1))}
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int16 = -32768
        if (!false) {
            if (true && true) {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int16 = (v25 & 42)
                    var v28 int16 = 1000
                    _ = len(v1)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                switch fi24 {
                case -2147483648:
                    var v29 uint64 = (v4 ^ v4)
                    _ = v29
                case 42:
                    var v30 int = (-(fi24))
                    _ = v30
                case -9223372036854775808:
                    fmt.Println(v9, v4)
                default:
                    v31 := GenS1{F0: v10, F1: v9, F2: (^(v8))}
                    _ = v31
                }
            }
            var v32 uint = 4294967295
            _ = v32
        }
        v33 := GenS2{F0: 42, F1: (v8 & (v8 >> ((829) & 63)))}
        _ = fi24
        _ = v25
        _ = v33
    }
    if true {
        v7 = v7
        for fi34 := 0; fi34 < 3; fi34++ {
            if (!true) {
                if (v9 > (^(v9))) {
                    for fi35 := 0; fi35 < 18; fi35++ {
                        var v36 uint32 = v9
                        var v37 int16 = -702
                        v10 += v10
                        v38 := GenS1{F0: 2, F1: 2147483648, F2: (^(v8))}
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    fmt.Println(v9)
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v40 := GenS1{F0: 225, F1: (v9 >> ((uint(v10)) & 63)), F2: v8}
                        v41 := GenS1{F0: 128, F1: v9, F2: v8}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                } else {
                    var v42 uint8 = v10
                    v4 = (v4 & (v4 / ((v4) | 1)))
                    _ = v42
                }
                if true {
                    v43 := GenS3{F0: (v7 + "foo bar")}
                    var v44 int16 = 42
                    _ = v43
                    _ = v44
                }
                if true {
                    v45 := GenS1{F0: 255, F1: v9, F2: v8}
                    _ = v45
                }
            }
            _ = fi34
        }
        var v46 uint64 = v4
        _ = v46
    }
    v4 += (^(v4))
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS4{F0: uint16(v4), F1: -547}
        v49 := GenS5{F0: v10}
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

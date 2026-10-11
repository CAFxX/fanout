package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 uint32 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]uint32)
    v3 := []uint32{4294967295}
    var v4 int32 = -203
    var v5 uint64 = 4294967295
    v1 = v4
    var v6 uint = 17
    v7 := v2[42]
    if ((v6 - 649) <= v6) {
        var v8 string = ("" + ("" + "a"))
        fmt.Println(v8, v4)
        _ = v8
    } else {
        switch "foo bar" {
        case "foo bar":
            v2[v7] = (v7 - v7)
        case "a":
            if true {
                if false {
                    v9 := GenS1{F0: (^(v4)), F1: v6}
                    delete(v2, v7)
                    var v10 string = ""
                    _ = v10
                    _ = v9
                }
            }
        default:
            v11 := GenS2{F0: 127, F1: 42}
            _ = v11
        }
    }
    v3 = append(v3, v7)
    if (v7 < 65535) {
        v2[(v7 % (((^(v7))) | 1))] = v7
        v2[v7] = 781
    } else {
        delete(v2, (v7 | 299))
        v2[v7] = (v7 << ((v6) & 63))
        if true {
            for fi12 := 0; fi12 < 13; fi12++ {
                switch fi12 {
                case -9223372036854775808:
                    fmt.Println(v5, v7, v7)
                case 603:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS3{F0: 1, F1: int(v7), F2: (^(v6))}
                        fmt.Println(fi12, v6, v5)
                        _ = fi13
                        _ = v14
                    }
                case 128:
                    var v15 bool = (v6 != 42)
                    _ = v15
                default:
                    v2[v7] = v7
                }
                v5 += (^(v5))
                _ = fi12
            }
            if false {
                var v16 int8 = -1
                v7 -= v7
                _ = v16
            } else {
                fmt.Println(v7, v1)
                v7 = (v7 % ((496) | 1))
            }
        }
    }
    v4 -= (v1 | 143)
    v3 = append(v3, v7)
    fmt.Println(v7)
    v4 -= -691
    fmt.Println(v6, v6, v5)
    v1 = (v4 / ((2147483647) | 1))
    var v17 uint16 = 1
    if true {
        v6 = v6
        switch -1 {
        case 128:
            for fi18 := 0; fi18 < 9; fi18++ {
                v17 -= (v17 ^ v17)
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 uint64 = (v5 / (((v5 & 1)) | 1))
                    if len(v3) > 0 {
                        v3[0] = 154
                    }
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        case 127:
            v2[0] = v7
        case 2147483647:
            v21 := GenS4{F0: 128, F1: 255, F2: (42 <= -137)}
            _ = v21
        }
    } else {
        if false {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 uint = (v6 / ((v6) | 1))
                _ = fi22
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 19; fi24++ {
                v7 = v7
                _ = fi24
            }
            var v25 string = ("" + "a")
            var v26 uint8 = 255
            _ = v25
            _ = v26
        }
        var v27 int32 = (v1 % ((-2147483648) | 1))
        if false {
            v28 := v2[(v7 + v7)]
            v7 = (v7 >> (((v6 / ((v6) | 1))) & 63))
            if len(v3) > 0 {
                v3[0] = v7
            }
            _ = v28
        }
        _ = v27
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        if (2147483647 != (^(v7))) {
            v30 := GenS5{F0: (^(v7)), F1: 100, F2: (v7 % ((v7) | 1))}
            for fi31 := 0; fi31 < 18; fi31++ {
                var v32 uint8 = 0
                fmt.Println(v5)
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            v6 += v6
            v6 += 2
            switch ((("\n\t" + "\n\t") + "hello") + "a") {
            case "foo bar":
                v5 = uint64(v7)
            case "hello":
                if false {
                    var v33 string = "a"
                    var v34 float32 = 0.0
                    _ = v33
                    _ = v34
                }
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v7 = (v7 - v7)
            }
        }
        delete(v2, v7)
        v35 := GenS1{F0: 267, F1: 486}
        _ = fi29
        _ = v35
    }
    v3 = append(v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

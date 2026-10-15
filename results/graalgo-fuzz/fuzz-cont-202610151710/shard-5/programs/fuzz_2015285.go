package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 0
    v3 := make(map[string]bool)
    var v4 int16 = -1
    v5 := make(map[int]int)
    v6 := []uint32{v1}
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case 4294967295:
            v4 = v4
        case 839:
            for fi9 := 0; fi9 < 17; fi9++ {
                switch ("a" + "\n\t") {
                case "foo bar":
                    var v10 int32 = 1
                    v2 -= v2
                    v2 -= 508
                    _ = v10
                case "\n\t":
                    _ = len(v5)
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    fmt.Println(v2, v4, v4)
                }
                _ = fi9
            }
        case 1:
            var v11 int64 = 956
            _ = v11
        }
        v4 -= v4
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := v5[fi8]
            switch "a" {
            case "\n\t":
                v14 := GenS1{F0: false, F1: v2, F2: 981}
                _ = v14
            case "foo bar":
                v13 = fi8
            case "":
                var v15 uint16 = 65535
                _ = v15
            }
            _ = len(v6)
            _ = len(v6)
            _ = fi12
            _ = v13
        }
        _ = fi8
    }
    var v16 int32 = 0
    if false {
        for fi17 := 0; fi17 < 7; fi17++ {
            v2 -= 2
            _ = fi17
        }
        v2 = uint64(v1)
        v1 -= v1
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            v16 += -3
            v2 -= v2
            for fi19 := 0; fi19 < 18; fi19++ {
                v6 = append(v6, v1)
                _ = fi19
            }
            _ = fi18
        }
        var v20 int8 = int8(v2)
        _ = v20
    }
    v21 := GenS2{F0: 1, F1: 42}
    v7 = v7
    for fi22 := 0; fi22 < 19; fi22++ {
        v2 = 1
        if ((fi22 | (-(fi22))) != fi22) {
            switch fi22 {
            case -656:
                fmt.Println(v16)
            case -854:
                v23 := v3["\n\t"]
                _ = v23
            case 2147483647:
                fmt.Println(v1)
            }
        }
        v4 = (v4 + v4)
        _ = fi22
    }
    if (true && (((0.1 * 1e10) <= 1.5) && true)) {
        v3["\n\t"] = true
    } else {
        if (-1 != 0) {
            if true {
                v2 -= (^(v2))
            } else {
                var v24 string = "\n\t"
                switch ("a" + v24) {
                case "":
                    v4 = (v4 | v4)
                case "\n\t":
                    switch int(v4) {
                    case 2147483647:
                        v5[42] = 825
                    case -9223372036854775808:
                        fmt.Println(v16)
                    default:
                        v24 = v24
                    }
                default:
                    var v25 bool = (false || (186 >= 9223372036854775807))
                    _ = v25
                }
                _ = v24
            }
            v4 = (v4 << ((1) & 63))
        } else {
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS2{F0: 65535, F1: (v2 >> ((255) & 63))}
                for fi28 := 0; fi28 < 16; fi28++ {
                    switch "" {
                    case "\n\t":
                        v29 := GenS1{F0: (false && (-69 >= 100)), F1: v2, F2: int64(v16)}
                        _ = v29
                    }
                    _ = fi28
                }
                v7 -= v7
                _ = fi26
                _ = v27
            }
            v4 = v4
        }
        delete(v5, 563)
    }
    var v30 uint8 = 1
    v3[("a" + ("hello" + ("foo bar" + ("a" + ((("a" + ("" + "\n\t")) + "") + "\n\t")))))] = (!(!true))
    var v31 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint8) int16 {
    v4 := GenS1{F0: 433, F1: 42}
    p1 = (p1 >> ((991) & 63))
    if false {
        v5 := GenS1{F0: 42, F1: 685}
        p1 -= (p0 & p1)
        _ = v5
    } else {
        v6 := []int{32768}
        v7 := GenS1{F0: -418, F1: 42}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    v1 := []string{}
    v2 := []uint32{1, 2147483648, 1, 2}
    v3 := make(map[int]uint32)
    v8 := GenS2{F0: 2, F1: ("\n\t" + "foo bar")}
    v9 := GenS3{F0: 42}
    var v10 bool = false
    switch "" {
    case "\n\t":
        var v11 float32 = float32((1.5 - ((0.0 * -2.25) - 3.14159)))
        _ = v11
    default:
        if (v10 && (3.14159 < 3.14159)) {
            switch ("foo bar" + ((("" + "\n\t") + "foo bar") + "")) {
            case "hello":
                fmt.Println(v10)
            }
        } else {
            for fi12 := 0; fi12 < 12; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    _ = len(v3)
                    fmt.Println(fi13, fi13)
                    _ = fi13
                }
                v10 = v10
                switch fi12 {
                case -158:
                    v10 = v10
                case -1:
                    var v14 string = ""
                    _ = v14
                case 0:
                    switch "\n\t" {
                    case "":
                        if (1 >= 1) {
                            v15 := GenS4{F0: 889}
                            switch 128 {
                            case -2147483648:
                                var v16 uint32 = 2
                                _ = v16
                            case -9223372036854775808:
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    v10 = v10
                                    var v18 uint32 = 2147483647
                                    _ = fi17
                                    _ = v18
                                }
                            default:
                                var v19 bool = false
                                _ = v19
                            }
                            v20 := GenS1{F0: int32(fi12), F1: -1}
                            _ = v15
                            _ = v20
                        }
                    case "\n\t":
                        v10 = v10
                    }
                default:
                    v3[fi12] = 2147483648
                }
                _ = fi12
            }
            var v21 uint = 255
            _ = v21
        }
    }
    v10 = (813 >= 0)
    v10 = (32767 == -269)
    if (v10 && v10) {
        v10 = v10
        if false {
            v22 := GenS5{F0: 41}
            _ = v22
        }
        v10 = (51 == -128)
    } else {
        fmt.Println(v10)
    }
    fmt.Println(v10)
    switch -1 {
    case 128:
        var v23 float64 = float64(0.0)
        _ = v23
    case -812:
        switch ("hello" + "\n\t") {
        case "foo bar":
            v10 = (!true)
        default:
            v3[9223372036854775807] = 65535
        }
    case 32768:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v10)
    if (127 < 128) {
        _ = len(v1)
        var v24 uint32 = 610
        _ = v24
    } else {
        v10 = v10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}

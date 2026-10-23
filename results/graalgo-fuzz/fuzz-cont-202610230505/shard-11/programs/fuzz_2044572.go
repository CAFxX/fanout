package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint16, p1 uint8, p2 uint) uint8 {
    v4 := GenS1{F0: 188, F1: -891, F2: 65535}
    var v5 uint16 = p0
    var v6 string = "hello"
    v7 := []int{0, 0, 809, 127}
    v7 = append(v7, -36)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := []uint32{1, 364, 873, 0}
    v2 := make(map[int]uint32)
    var v3 uint32 = 405
    switch "\n\t" {
    case "foo bar":
        v3 = (v3 >> ((18446744073709551615) & 63))
    case "\n\t":
        for fi8 := 0; fi8 < 20; fi8++ {
            switch fi8 {
            case -661:
                fmt.Println(v3, v3)
            case -2147483648:
                v9 := GenS2{F0: 290, F1: 1}
                _ = v9
            case 4294967295:
                _ = len(v2)
            default:
                if true {
                    var v10 uint32 = v3
                    var v11 int32 = 988
                    if true {
                        v12 := GenS2{F0: 32767, F1: 894}
                        fmt.Println(v3, v3, v3)
                        _ = v12
                    } else {
                        v3 = v3
                        v10 += (v10 - (v10 - v10))
                        var v13 uint16 = 9
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                } else {
                    _ = len(v1)
                    v3 = v3
                    var v14 float64 = 1.5
                    _ = v14
                }
            }
            _ = fi8
        }
    case "":
        for fi15 := 0; fi15 < 1; fi15++ {
            if (!true) {
                fmt.Println(v3)
            } else {
                switch -9223372036854775808 {
                case 42:
                    _ = len(v1)
                case -90:
                    v16 := GenS2{F0: 42, F1: 107}
                    _ = v16
                case 127:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if len(v1) > 0 {
                            v1[0] = (v3 - v3)
                        }
                        fmt.Println(v3, fi15, fi17)
                        _ = fi17
                    }
                default:
                    var v18 int32 = 2147483647
                    _ = v18
                }
            }
            v3 = v3
            v19 := GenS1{F0: 172, F1: -32768, F2: v3}
            _ = fi15
            _ = v19
        }
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v3 += v3
    var v20 int8 = -1
    if true {
        v1 = append(v1, 365)
    } else {
        var v21 int8 = 42
        v20 = (v21 >> ((1) & 63))
        _ = v21
    }
    var v22 int = -449
    v22 = (v22 | v22)
    v20 = 100
    v23 := GenS3{F0: 470, F1: ("a" + ((("" + "\n\t") + "foo bar") + ((("hello" + "hello") + "") + "\n\t")))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
}
